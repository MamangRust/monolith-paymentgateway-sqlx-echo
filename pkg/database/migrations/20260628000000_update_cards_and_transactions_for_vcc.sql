-- +goose Up
-- +goose StatementBegin

-- Add new columns to cards table
ALTER TABLE cards
    ADD COLUMN IF NOT EXISTS "status" VARCHAR(20) NOT NULL DEFAULT 'active',
    ADD COLUMN IF NOT EXISTS "credit_limit" INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS "outstanding_balance" INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS "reward_points" INT NOT NULL DEFAULT 0;

-- Add fraud_score column to transactions table
ALTER TABLE transactions
    ADD COLUMN IF NOT EXISTS "fraud_score" INT NOT NULL DEFAULT 0;

-- Create billing_cycles table
CREATE TABLE IF NOT EXISTS "billing_cycles" (
    "billing_id" SERIAL PRIMARY KEY,
    "card_number" VARCHAR(16) NOT NULL REFERENCES "cards" ("card_number"),
    "cycle_start" TIMESTAMP NOT NULL,
    "cycle_end" TIMESTAMP NOT NULL,
    "amount_due" INT NOT NULL DEFAULT 0,
    "due_date" TIMESTAMP NOT NULL,
    "status" VARCHAR(20) NOT NULL DEFAULT 'unpaid',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_billing_cycles_card_number ON billing_cycles (card_number);
CREATE INDEX IF NOT EXISTS idx_billing_cycles_status ON billing_cycles (status);
CREATE INDEX IF NOT EXISTS idx_billing_cycles_due_date ON billing_cycles (due_date);
CREATE INDEX IF NOT EXISTS idx_cards_status ON cards (status);
CREATE INDEX IF NOT EXISTS idx_transactions_fraud_score ON transactions (fraud_score);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS idx_transactions_fraud_score;
DROP INDEX IF EXISTS idx_cards_status;
DROP INDEX IF EXISTS idx_billing_cycles_due_date;
DROP INDEX IF EXISTS idx_billing_cycles_status;
DROP INDEX IF EXISTS idx_billing_cycles_card_number;

DROP TABLE IF EXISTS "billing_cycles";

ALTER TABLE transactions DROP COLUMN IF EXISTS "fraud_score";

ALTER TABLE cards
    DROP COLUMN IF EXISTS "reward_points",
    DROP COLUMN IF EXISTS "outstanding_balance",
    DROP COLUMN IF EXISTS "credit_limit",
    DROP COLUMN IF EXISTS "status";

-- +goose StatementEnd
