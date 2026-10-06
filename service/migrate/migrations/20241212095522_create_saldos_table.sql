-- +goose Up
-- +goose StatementBegin
CREATE TABLE "saldos" (
    "saldo_id" SERIAL PRIMARY KEY,
    "card_number" VARCHAR(16) NOT NULL REFERENCES "cards" ("card_number"),
    "total_balance" INT NOT NULL,
    "withdraw_amount" INT DEFAULT 0,
    "withdraw_time" TIMESTAMP DEFAULT current_timestamp,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp,
    "deleted_at" TIMESTAMP DEFAULT NULL
);

CREATE INDEX idx_saldos_card_number ON saldos (card_number);

CREATE INDEX idx_saldos_withdraw_time ON saldos (withdraw_time);

CREATE INDEX idx_saldos_total_balance ON saldos (total_balance);

CREATE INDEX idx_saldos_card_number_withdraw_time ON saldos (card_number, withdraw_time);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_saldos_card_number;

DROP INDEX IF EXISTS idx_saldos_withdraw_time;

DROP INDEX IF EXISTS idx_saldos_total_balance;

DROP INDEX IF EXISTS idx_saldos_card_number_withdraw_time;

DROP TABLE IF EXISTS "saldos";

-- +goose StatementEnd
