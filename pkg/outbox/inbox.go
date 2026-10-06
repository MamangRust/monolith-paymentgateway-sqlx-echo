package outbox

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var ErrInvalidInboxKey = errors.New("invalid consumer inbox key")

// ConsumerInbox is the durable deduplication contract used by Kafka handlers
// (durable consumer-side deduplication). It replaces in-memory-only deduplication:
// reservations survive consumer restarts and rebalances, so at-least-once
// redelivery cannot send the same email twice.
type ConsumerInbox interface {
	Reserve(ctx context.Context, consumerName, eventKey, topic string, partition int32, offset int64) (bool, bool, int64, error)
	MarkProcessed(ctx context.Context, consumerName, eventKey string, reservationVersion int64) error
	Release(ctx context.Context, consumerName, eventKey string, reservationVersion int64, processingErr error) error
}

// InboxExecutor is the executor contract for inbox persistence. It is satisfied
// by both *sqlx.DB and *sqlx.Tx so handlers can run an inbox write inside their
// own transaction or standalone.
type InboxExecutor interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

const reserveConsumerInbox = `WITH reserved AS (
    INSERT INTO consumer_inbox (
        consumer_name, event_key, topic, partition_id, message_offset,
        status, attempts, reservation_version, lease_until, last_error, processed_at
    )
    VALUES ($1, $2, $3, $4, $5, 'processing', 1, 1,
            current_timestamp + interval '1 minute', '', NULL)
    ON CONFLICT (consumer_name, event_key) DO UPDATE
    SET status = 'processing',
        attempts = consumer_inbox.attempts + 1,
        reservation_version = consumer_inbox.reservation_version + 1,
        lease_until = current_timestamp + interval '1 minute',
        last_error = '',
        topic = EXCLUDED.topic,
        partition_id = EXCLUDED.partition_id,
        message_offset = EXCLUDED.message_offset
    WHERE consumer_inbox.status <> 'processed'
      AND consumer_inbox.lease_until <= current_timestamp
    RETURNING reservation_version
)
SELECT
    EXISTS (SELECT 1 FROM reserved) AS reserved,
    EXISTS (
        SELECT 1
        FROM consumer_inbox ci
        WHERE ci.consumer_name = $1
          AND ci.event_key = $2
          AND ci.status = 'processed'
    ) AS processed,
    COALESCE(
        (SELECT reservation_version FROM reserved),
        (SELECT ci.reservation_version FROM consumer_inbox ci WHERE ci.consumer_name = $1 AND ci.event_key = $2)
    )::BIGINT AS reservation_version`

const markConsumerInboxProcessed = `UPDATE consumer_inbox
SET status = 'processed', processed_at = current_timestamp,
    lease_until = current_timestamp, last_error = ''
WHERE consumer_name = $1 AND event_key = $2
  AND status = 'processing' AND reservation_version = $3`

const releaseConsumerInbox = `UPDATE consumer_inbox
SET status = 'pending', lease_until = current_timestamp,
    last_error = $3
WHERE consumer_name = $1 AND event_key = $2
  AND status = 'processing' AND reservation_version = $4`

// InboxReservation is the result of a Reserve call.
type InboxReservation struct {
	Reserved           bool
	Processed          bool
	ReservationVersion int64
}

// Reserve claims an event for a consumer. It returns false when the event was
// already processed. An expired processing lease may be reclaimed after a
// consumer crashes.
func Reserve(ctx context.Context, exec InboxExecutor, consumerName, eventKey, topic string, partition int32, offset int64) (bool, bool, int64, error) {
	if exec == nil || consumerName == "" || eventKey == "" {
		return false, false, 0, ErrInvalidInboxKey
	}
	var reservation InboxReservation
	if err := exec.QueryRowContext(ctx, reserveConsumerInbox,
		consumerName, eventKey, topic, partition, offset,
	).Scan(&reservation.Reserved, &reservation.Processed, &reservation.ReservationVersion); err != nil {
		return false, false, 0, err
	}
	return reservation.Reserved, reservation.Processed, reservation.ReservationVersion, nil
}

// MarkProcessed completes only the active reservation.
func MarkProcessed(ctx context.Context, exec InboxExecutor, consumerName, eventKey string, reservationVersion int64) error {
	if exec == nil || consumerName == "" || eventKey == "" {
		return ErrInvalidInboxKey
	}
	_, err := exec.ExecContext(ctx, markConsumerInboxProcessed, consumerName, eventKey, reservationVersion)
	return err
}

// Release releases only the active reservation.
func Release(ctx context.Context, exec InboxExecutor, consumerName, eventKey string, reservationVersion int64, processingErr error) error {
	if exec == nil || consumerName == "" || eventKey == "" {
		return ErrInvalidInboxKey
	}
	lastError := "consumer processing failed"
	if processingErr != nil {
		lastError = processingErr.Error()
	}
	_, err := exec.ExecContext(ctx, releaseConsumerInbox, consumerName, eventKey, lastError, reservationVersion)
	return err
}

// PostgresInbox adapts a sqlx.DB to ConsumerInbox. Reservation and completion
// are committed independently because an external side effect cannot share a
// PostgreSQL transaction with the Kafka consumer.
type PostgresInbox struct {
	db *sqlx.DB
}

func NewPostgresInbox(db *sqlx.DB) (*PostgresInbox, error) {
	if db == nil {
		return nil, errors.New("inbox db is nil")
	}
	return &PostgresInbox{db: db}, nil
}

func (i *PostgresInbox) Reserve(ctx context.Context, consumerName, eventKey, topic string, partition int32, offset int64) (bool, bool, int64, error) {
	tx, err := i.db.Beginx()
	if err != nil {
		return false, false, 0, err
	}
	defer tx.Rollback()
	reserved, processed, reservationVersion, err := Reserve(ctx, tx, consumerName, eventKey, topic, partition, offset)
	if err != nil {
		return false, false, 0, err
	}
	if err := tx.Commit(); err != nil {
		return false, false, 0, err
	}
	return reserved, processed, reservationVersion, nil
}

func (i *PostgresInbox) MarkProcessed(ctx context.Context, consumerName, eventKey string, reservationVersion int64) error {
	tx, err := i.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := MarkProcessed(ctx, tx, consumerName, eventKey, reservationVersion); err != nil {
		return err
	}
	return tx.Commit()
}

func (i *PostgresInbox) Release(ctx context.Context, consumerName, eventKey string, reservationVersion int64, processingErr error) error {
	tx, err := i.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := Release(ctx, tx, consumerName, eventKey, reservationVersion, processingErr); err != nil {
		return err
	}
	return tx.Commit()
}
