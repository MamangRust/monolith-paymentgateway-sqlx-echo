-- +goose Up
-- +goose StatementBegin

ALTER TABLE cards
    ADD COLUMN IF NOT EXISTS event_version BIGINT NOT NULL DEFAULT 0;
ALTER TABLE merchants
    ADD COLUMN IF NOT EXISTS event_version BIGINT NOT NULL DEFAULT 0;
ALTER TABLE merchant_documents
    ADD COLUMN IF NOT EXISTS event_version BIGINT NOT NULL DEFAULT 0;

CREATE UNIQUE INDEX IF NOT EXISTS uq_card_rewards_txn_id
    ON card_rewards (txn_id)
    WHERE txn_id <> '';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS uq_card_rewards_txn_id;
ALTER TABLE merchant_documents DROP COLUMN IF EXISTS event_version;
ALTER TABLE merchants DROP COLUMN IF EXISTS event_version;
ALTER TABLE cards DROP COLUMN IF EXISTS event_version;

-- +goose StatementEnd
