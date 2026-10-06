-- +goose Up
-- +goose StatementBegin

-- Idempotency key untuk command finansial (topup/transfer/withdraw/transaction).
-- Unique index parsial hanya untuk key yang diisi ('' diizinkan, tidak di-index)
-- dan record yang belum di-soft-delete, sehingga replay dengan key sama selalu
-- mengembalikan record asli tanpa double-settlement.

ALTER TABLE topups ADD COLUMN IF NOT EXISTS idempotency_key VARCHAR(64) NOT NULL DEFAULT '';
CREATE UNIQUE INDEX IF NOT EXISTS idx_topups_idempotency_key
    ON topups (idempotency_key)
    WHERE idempotency_key <> '';

ALTER TABLE transfers ADD COLUMN IF NOT EXISTS idempotency_key VARCHAR(64) NOT NULL DEFAULT '';
CREATE UNIQUE INDEX IF NOT EXISTS idx_transfers_idempotency_key
    ON transfers (idempotency_key)
    WHERE idempotency_key <> '';

ALTER TABLE withdraws ADD COLUMN IF NOT EXISTS idempotency_key VARCHAR(64) NOT NULL DEFAULT '';
CREATE UNIQUE INDEX IF NOT EXISTS idx_withdraws_idempotency_key
    ON withdraws (idempotency_key)
    WHERE idempotency_key <> '';

ALTER TABLE transactions ADD COLUMN IF NOT EXISTS idempotency_key VARCHAR(64) NOT NULL DEFAULT '';
CREATE UNIQUE INDEX IF NOT EXISTS idx_transactions_idempotency_key
    ON transactions (idempotency_key)
    WHERE idempotency_key <> '';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS idx_topups_idempotency_key;
ALTER TABLE topups DROP COLUMN IF EXISTS idempotency_key;

DROP INDEX IF EXISTS idx_transfers_idempotency_key;
ALTER TABLE transfers DROP COLUMN IF EXISTS idempotency_key;

DROP INDEX IF EXISTS idx_withdraws_idempotency_key;
ALTER TABLE withdraws DROP COLUMN IF EXISTS idempotency_key;

DROP INDEX IF EXISTS idx_transactions_idempotency_key;
ALTER TABLE transactions DROP COLUMN IF EXISTS idempotency_key;

-- +goose StatementEnd
