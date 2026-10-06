-- +goose Up
-- +goose StatementBegin

ALTER TABLE outbox_events
    ADD COLUMN IF NOT EXISTS claim_version BIGINT NOT NULL DEFAULT 0;

ALTER TABLE consumer_inbox
    ADD COLUMN IF NOT EXISTS reservation_version BIGINT NOT NULL DEFAULT 0;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE consumer_inbox DROP COLUMN IF EXISTS reservation_version;
ALTER TABLE outbox_events DROP COLUMN IF EXISTS claim_version;

-- +goose StatementEnd
