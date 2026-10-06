package kafka

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/IBM/sarama"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertCardEventLog = `INSERT INTO card_event_logs (
    topic,
    event_type,
    card_number,
    reference_id,
    payload
)
VALUES ($1, $2, NULLIF($3::text, ''), NULLIF($4::text, ''), $5::jsonb)
ON CONFLICT DO NOTHING
RETURNING event_id, topic, event_type, card_number, reference_id, payload, received_at`

// CardEventLogHandler stores card domain events durably for audit and replay.
type CardEventLogHandler struct {
	db     *sqlx.DB
	logger logger.LoggerInterface
}

func NewCardEventLogHandler(db *sqlx.DB, logger logger.LoggerInterface) *CardEventLogHandler {
	return &CardEventLogHandler{db: db, logger: logger}
}

func (h *CardEventLogHandler) Setup(sarama.ConsumerGroupSession) error {
	h.logger.Info("card event log consumer setup")
	return nil
}

func (h *CardEventLogHandler) Cleanup(sarama.ConsumerGroupSession) error {
	h.logger.Info("card event log consumer cleanup")
	return nil
}

func (h *CardEventLogHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		event, err := parseCardEvent(msg.Topic, msg.Key, msg.Value)
		if err != nil {
			// A malformed event can never succeed on retry; acknowledge it so a
			// poison message cannot block the audit consumer partition forever.
			h.logger.Error("invalid card event audit payload", zap.Error(err), zap.String("topic", msg.Topic))
			session.MarkMessage(msg, "")
			continue
		}
		if h.db == nil {
			return fmt.Errorf("card event audit database is nil")
		}
		insertCtx, cancel := context.WithTimeout(session.Context(), 10*time.Second)
		var inserted db.CardEventLog
		err = h.db.QueryRowxContext(insertCtx, insertCardEventLog,
			event.Topic,
			event.EventType,
			event.CardNumber,
			event.ReferenceID,
			event.Payload,
		).Scan(
			&inserted.EventID,
			&inserted.Topic,
			&inserted.EventType,
			&inserted.CardNumber,
			&inserted.ReferenceID,
			&inserted.Payload,
			&inserted.ReceivedAt,
		)
		cancel()
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			// Do not acknowledge database failures. Sarama will redeliver the
			// message and ON CONFLICT keeps successful redeliveries idempotent.
			return fmt.Errorf("insert card event log: %w", err)
		}
		session.MarkMessage(msg, "")
	}
	return nil
}

type cardEvent struct {
	Topic       string
	EventType   string
	CardNumber  string
	ReferenceID string
	Payload     []byte
}

func parseCardEvent(topic string, key, payload []byte) (cardEvent, error) {
	if strings.TrimSpace(topic) == "" || !json.Valid(payload) {
		return cardEvent{}, fmt.Errorf("topic or payload is invalid")
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(payload, &fields); err != nil || fields == nil {
		return cardEvent{}, fmt.Errorf("payload must be a JSON object")
	}

	event := cardEvent{
		Topic:     topic,
		EventType: strings.TrimPrefix(topic, "card."),
		Payload:   append([]byte(nil), payload...),
	}
	event.CardNumber = rawString(fields, "card_number")
	switch topic {
	case "card.payment.posted":
		event.ReferenceID = rawString(fields, "payment_id")
	case "card.fraud.alert":
		event.ReferenceID = rawString(fields, "txn_id")
	case "card.statement.generated":
		event.ReferenceID = rawString(fields, "billing_id")
		if event.ReferenceID == "" {
			event.ReferenceID = rawString(fields, "statement_id")
		}
		if event.ReferenceID == "" {
			event.ReferenceID = rawString(fields, "generated_at")
		}
		if event.ReferenceID == "" {
			event.ReferenceID = rawString(fields, "billing_cycle_day")
		}
	case "card.limit.changed":
		// Limit changes intentionally do not use the partial unique index: one
		// card may legitimately emit many limit updates over time.
		event.ReferenceID = event.CardNumber
	}
	if event.ReferenceID == "" {
		event.ReferenceID = string(key)
	}
	return event, nil
}

func rawString(fields map[string]json.RawMessage, name string) string {
	value, ok := fields[name]
	if !ok {
		return ""
	}
	var stringValue string
	if json.Unmarshal(value, &stringValue) == nil {
		return stringValue
	}
	var numberValue json.Number
	if json.Unmarshal(value, &numberValue) == nil {
		return numberValue.String()
	}
	return ""
}

var _ sarama.ConsumerGroupHandler = (*CardEventLogHandler)(nil)
