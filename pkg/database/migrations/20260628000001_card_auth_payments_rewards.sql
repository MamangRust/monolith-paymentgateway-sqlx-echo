-- +goose Up
-- +goose StatementBegin

-- Create card_auth_transactions table
CREATE TABLE IF NOT EXISTS "card_auth_transactions" (
    "auth_id" SERIAL PRIMARY KEY,
    "txn_id" VARCHAR(36) NOT NULL UNIQUE,
    "card_number" VARCHAR(16) NOT NULL REFERENCES "cards" ("card_number"),
    "merchant_id" INT NOT NULL DEFAULT 0,
    "amount" BIGINT NOT NULL DEFAULT 0,
    "currency" VARCHAR(3) NOT NULL DEFAULT 'IDR',
    "mcc" VARCHAR(4) NOT NULL DEFAULT '',
    "pos_entry_mode" VARCHAR(3) NOT NULL DEFAULT '',
    "status" VARCHAR(20) NOT NULL DEFAULT 'pending',
    "idempotency_key" VARCHAR(64) NOT NULL DEFAULT '',
    "risk_score" INT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_auth_txn_card_number ON card_auth_transactions (card_number);
CREATE INDEX IF NOT EXISTS idx_auth_txn_status ON card_auth_transactions (status);
CREATE INDEX IF NOT EXISTS idx_auth_txn_idempotency ON card_auth_transactions (idempotency_key);
CREATE INDEX IF NOT EXISTS idx_auth_txn_txn_id ON card_auth_transactions (txn_id);

-- Create card_payments table
CREATE TABLE IF NOT EXISTS "card_payments" (
    "payment_id" SERIAL PRIMARY KEY,
    "payment_uuid" VARCHAR(36) NOT NULL UNIQUE,
    "card_number" VARCHAR(16) NOT NULL REFERENCES "cards" ("card_number"),
    "billing_id" INT REFERENCES "billing_cycles" ("billing_id"),
    "amount" BIGINT NOT NULL DEFAULT 0,
    "payment_channel" VARCHAR(20) NOT NULL DEFAULT 'bank_transfer',
    "reference_id" VARCHAR(64) NOT NULL DEFAULT '',
    "status" VARCHAR(20) NOT NULL DEFAULT 'completed',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_card_payments_card_number ON card_payments (card_number);
CREATE INDEX IF NOT EXISTS idx_card_payments_billing_id ON card_payments (billing_id);

-- Create card_rewards table
CREATE TABLE IF NOT EXISTS "card_rewards" (
    "reward_id" SERIAL PRIMARY KEY,
    "card_number" VARCHAR(16) NOT NULL REFERENCES "cards" ("card_number"),
    "txn_id" VARCHAR(36) NOT NULL DEFAULT '',
    "amount" BIGINT NOT NULL DEFAULT 0,
    "mcc" VARCHAR(4) NOT NULL DEFAULT '',
    "points_earned" INT NOT NULL DEFAULT 0,
    "expires_at" TIMESTAMP NOT NULL,
    "redeemed" BOOLEAN NOT NULL DEFAULT FALSE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_card_rewards_card_number ON card_rewards (card_number);
CREATE INDEX IF NOT EXISTS idx_card_rewards_redeemed ON card_rewards (redeemed);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS "card_rewards";
DROP TABLE IF EXISTS "card_payments";
DROP TABLE IF EXISTS "card_auth_transactions";

-- +goose StatementEnd
