-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS outbox_events (
    id BIGSERIAL PRIMARY KEY,
    event_key VARCHAR(128) NOT NULL,
    topic VARCHAR(255) NOT NULL,
    message_key VARCHAR(255) NOT NULL DEFAULT '',
    payload JSONB NOT NULL,
    status VARCHAR(16) NOT NULL DEFAULT 'pending'
        CHECK (status IN ('pending', 'publishing', 'published', 'failed')),
    attempts INTEGER NOT NULL DEFAULT 0 CHECK (attempts >= 0),
    next_attempt_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
    last_error TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
    published_at TIMESTAMPTZ,
    CONSTRAINT outbox_events_event_key_unique UNIQUE (event_key)
);

CREATE INDEX IF NOT EXISTS idx_outbox_events_pending
    ON outbox_events (next_attempt_at, id)
    WHERE status = 'pending';
CREATE INDEX IF NOT EXISTS idx_outbox_events_stale_publishing
    ON outbox_events (updated_at, id)
    WHERE status = 'publishing';

CREATE TABLE IF NOT EXISTS consumer_inbox (
    consumer_name VARCHAR(128) NOT NULL,
    event_key VARCHAR(255) NOT NULL,
    topic VARCHAR(255) NOT NULL DEFAULT '',
    partition_id INTEGER NOT NULL DEFAULT -1,
    message_offset BIGINT NOT NULL DEFAULT -1,
    status VARCHAR(16) NOT NULL DEFAULT 'processing'
        CHECK (status IN ('processing', 'processed', 'pending')),
    attempts INTEGER NOT NULL DEFAULT 1 CHECK (attempts >= 0),
    lease_until TIMESTAMPTZ NOT NULL DEFAULT current_timestamp + interval '1 minute',
    last_error TEXT NOT NULL DEFAULT '',
    processed_at TIMESTAMPTZ,
    PRIMARY KEY (consumer_name, event_key)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS consumer_inbox;
DROP TABLE IF EXISTS outbox_events;
-- +goose StatementEnd
