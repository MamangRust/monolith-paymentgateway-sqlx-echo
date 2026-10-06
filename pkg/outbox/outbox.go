package outbox

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	ErrInvalidEvent = errors.New("invalid outbox event")
	ErrNoPublisher  = errors.New("outbox publisher is nil")
)

// Event is a message that must be persisted in the same database transaction
// as the business mutation that produced it. EventKey is an application-level
// idempotency key and is unique in the outbox table.
type Event struct {
	EventKey string
	Topic    string
	Key      string
	Payload  []byte
}

func (e Event) Validate() error {
	if e.EventKey == "" || e.Topic == "" || len(e.Payload) == 0 || !json.Valid(e.Payload) {
		return ErrInvalidEvent
	}
	return nil
}

// Executor is satisfied by *sqlx.Tx and keeps enqueue usable inside an existing
// business transaction. The caller owns commit/rollback of the transaction.
type Executor interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

const enqueueOutboxEvent = `INSERT INTO outbox_events (event_key, topic, message_key, payload)
VALUES ($1, $2, $3, $4::jsonb)
ON CONFLICT (event_key) DO NOTHING`

// Enqueue persists an event in the caller's transaction. A duplicate event key
// is treated as success, which makes retries of a business transaction safe.
func Enqueue(ctx context.Context, tx Executor, event Event) error {
	if err := event.Validate(); err != nil {
		return err
	}
	if tx == nil {
		return errors.New("outbox executor is nil")
	}
	if _, err := tx.ExecContext(ctx, enqueueOutboxEvent,
		event.EventKey,
		event.Topic,
		event.Key,
		event.Payload,
	); err != nil {
		return err
	}
	return nil
}

// Publisher is implemented by pkg/kafka.Kafka and deliberately kept small so
// the outbox package does not depend on the Kafka implementation.
type Publisher interface {
	SendMessageWithRetry(topic, key string, value []byte, attempts int) error
}

type RelayConfig struct {
	BatchSize    int
	PollInterval time.Duration
	MaxAttempts  int
	BaseBackoff  time.Duration
	Topics       []string
}

func (c RelayConfig) withDefaults() RelayConfig {
	if c.BatchSize <= 0 {
		c.BatchSize = 50
	}
	if c.PollInterval <= 0 {
		c.PollInterval = time.Second
	}
	if c.MaxAttempts <= 0 {
		c.MaxAttempts = 10
	}
	if c.BaseBackoff <= 0 {
		c.BaseBackoff = time.Second
	}
	return c
}

type Relay struct {
	db        *sqlx.DB
	publisher Publisher
	config    RelayConfig
}

func NewRelay(db *sqlx.DB, publisher Publisher, config RelayConfig) (*Relay, error) {
	if db == nil {
		return nil, errors.New("outbox pool is nil")
	}
	if publisher == nil {
		return nil, ErrNoPublisher
	}
	return &Relay{db: db, publisher: publisher, config: config.withDefaults()}, nil
}

// Run polls until ctx is cancelled. A claimed event remains recoverable when
// a process crashes: the claim is stored as publishing and stale claims are
// returned to pending by the next claim operation.
func (r *Relay) Run(ctx context.Context) error {
	for {
		processed, err := r.runBatch(ctx)
		if errors.Is(ctx.Err(), context.Canceled) {
			return ctx.Err()
		}
		if err != nil {
			// Keep the relay alive across transient database/Kafka state errors.
			// Claimed rows remain recoverable through the stale-publishing lease.
			if !wait(ctx, r.config.PollInterval) {
				return ctx.Err()
			}
			continue
		}
		if processed == 0 && !wait(ctx, r.config.PollInterval) {
			return ctx.Err()
		}
	}
}

func wait(ctx context.Context, duration time.Duration) bool {
	timer := time.NewTimer(duration)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-timer.C:
		return true
	}
}

type claimedEvent struct {
	ID           int64
	Topic        string
	Key          string
	Payload      []byte
	Attempt      int
	ClaimVersion int64
}

// claimedRow mirrors the claim RETURNING list.
type claimedRow struct {
	ID           int64  `db:"id"`
	Topic        string `db:"topic"`
	MessageKey   string `db:"message_key"`
	Payload      []byte `db:"payload"`
	Attempts     int32  `db:"attempts"`
	ClaimVersion int64  `db:"claim_version"`
}

const claimOutboxEvents = `WITH candidates AS (
    SELECT id
    FROM outbox_events
    WHERE ((status = 'pending' AND next_attempt_at <= current_timestamp)
       OR (status = 'publishing' AND updated_at < current_timestamp - interval '1 minute'))
    ORDER BY id
    FOR UPDATE SKIP LOCKED
    LIMIT $1
)
UPDATE outbox_events e
SET status = 'publishing', attempts = e.attempts + 1,
    claim_version = e.claim_version + 1,
    updated_at = current_timestamp
FROM candidates c
WHERE e.id = c.id
RETURNING e.id, e.topic, e.message_key, e.payload, e.attempts, e.claim_version`

const claimOutboxEventsByTopics = `WITH candidates AS (
    SELECT id
    FROM outbox_events
    WHERE ((status = 'pending' AND next_attempt_at <= current_timestamp)
       OR (status = 'publishing' AND updated_at < current_timestamp - interval '1 minute'))
      AND topic = ANY($2::text[])
    ORDER BY id
    FOR UPDATE SKIP LOCKED
    LIMIT $1
)
UPDATE outbox_events e
SET status = 'publishing', attempts = e.attempts + 1,
    claim_version = e.claim_version + 1,
    updated_at = current_timestamp
FROM candidates c
WHERE e.id = c.id
RETURNING e.id, e.topic, e.message_key, e.payload, e.attempts, e.claim_version`

const markOutboxPublished = `UPDATE outbox_events
SET status = 'published', published_at = current_timestamp,
    updated_at = current_timestamp, last_error = ''
WHERE id = $1 AND status = 'publishing' AND claim_version = $2`

const markOutboxFailed = `UPDATE outbox_events
SET status = $2,
    next_attempt_at = current_timestamp + $3::interval,
    last_error = $4,
    updated_at = current_timestamp
WHERE id = $1 AND status = 'publishing' AND claim_version = $5`

func (r *Relay) runBatch(ctx context.Context) (int, error) {
	events, err := r.claim(ctx)
	if err != nil {
		return 0, err
	}
	for _, event := range events {
		if err := r.publish(ctx, event); err != nil {
			if updateErr := r.markFailed(ctx, event, err); updateErr != nil {
				return len(events), fmt.Errorf("record outbox failure: %w (publish error: %v)", updateErr, err)
			}
			continue
		}
		if err := r.markPublished(ctx, event); err != nil {
			return len(events), err
		}
	}
	return len(events), nil
}

func (r *Relay) claim(ctx context.Context) ([]claimedEvent, error) {
	var rows []*claimedRow
	var err error
	if len(r.config.Topics) == 0 {
		err = r.db.SelectContext(ctx, &rows, claimOutboxEvents,
			int32(r.config.BatchSize),
		)
	} else {
		err = r.db.SelectContext(ctx, &rows, claimOutboxEventsByTopics,
			int32(r.config.BatchSize),
			r.config.Topics,
		)
	}
	if err != nil {
		return nil, err
	}
	events := make([]claimedEvent, 0, len(rows))
	for _, row := range rows {
		events = append(events, claimedEvent{
			ID: row.ID, Topic: row.Topic, Key: row.MessageKey,
			Payload: row.Payload, Attempt: int(row.Attempts), ClaimVersion: row.ClaimVersion,
		})
	}
	return events, nil
}

func (r *Relay) publish(ctx context.Context, event claimedEvent) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return r.publisher.SendMessageWithRetry(event.Topic, event.Key, event.Payload, 1)
}

func (r *Relay) markPublished(ctx context.Context, event claimedEvent) error {
	_, err := r.db.ExecContext(ctx, markOutboxPublished,
		event.ID,
		event.ClaimVersion,
	)
	return err
}

func (r *Relay) markFailed(ctx context.Context, event claimedEvent, publishErr error) error {
	status := "pending"
	if event.Attempt >= r.config.MaxAttempts {
		status = "failed"
	}
	attempt := min(event.Attempt, 10)
	if attempt < 1 {
		attempt = 1
	}
	backoff := r.config.BaseBackoff * time.Duration(1<<(attempt-1))
	// PostgreSQL parses the interval literal cast ($3::interval); a
	// microseconds string reproduces the pgtype.Interval microsecond value.
	_, err := r.db.ExecContext(ctx, markOutboxFailed,
		event.ID,
		status,
		formatPgInterval(backoff),
		publishErr.Error(),
		event.ClaimVersion,
	)
	return err
}

// formatPgInterval renders a duration as a PostgreSQL interval literal with
// microsecond precision, matching the pgtype.Interval microsecond value used
// before the migration.
func formatPgInterval(d time.Duration) string {
	return fmt.Sprintf("%d microseconds", d.Microseconds())
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var _ = sql.ErrNoRows
