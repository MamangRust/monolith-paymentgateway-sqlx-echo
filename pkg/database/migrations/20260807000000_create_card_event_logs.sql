-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS card_event_logs (
    event_id BIGSERIAL PRIMARY KEY,
    topic VARCHAR(80) NOT NULL,
    event_type VARCHAR(80) NOT NULL,
    card_number VARCHAR(16),
    reference_id VARCHAR(64),
    payload JSONB NOT NULL,
    received_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_card_event_logs_received_at
    ON card_event_logs (received_at);
CREATE INDEX IF NOT EXISTS idx_card_event_logs_topic
    ON card_event_logs (topic);
CREATE INDEX IF NOT EXISTS idx_card_event_logs_card_number
    ON card_event_logs (card_number)
    WHERE card_number IS NOT NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_card_event_logs_topic_reference
    ON card_event_logs (topic, reference_id)
    WHERE topic <> 'card.limit.changed' AND reference_id IS NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_card_event_logs_topic_reference;
DROP INDEX IF EXISTS idx_card_event_logs_card_number;
DROP INDEX IF EXISTS idx_card_event_logs_topic;
DROP INDEX IF EXISTS idx_card_event_logs_received_at;
DROP TABLE IF EXISTS card_event_logs;
-- +goose StatementEnd
