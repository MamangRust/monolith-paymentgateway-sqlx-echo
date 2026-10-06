-- +goose Up
-- +goose StatementBegin
-- This migration intentionally removes duplicate active saldo rows. Before
-- applying it in production, take a database backup: the duplicate rows are
-- historical data that cannot be reconstructed by the Down migration.
-- Keep the highest real balance for each card. If balances tie, keep the
-- newest saldo row so historical duplicates cannot block the unique index.
WITH ranked AS (
    SELECT
        saldo_id,
        ROW_NUMBER() OVER (
            PARTITION BY card_number
            ORDER BY total_balance DESC, saldo_id DESC
        ) AS row_number
    FROM saldos
    WHERE deleted_at IS NULL
)
DELETE FROM saldos
WHERE saldo_id IN (
    SELECT saldo_id
    FROM ranked
    WHERE row_number > 1
);

CREATE UNIQUE INDEX idx_saldos_card_number_active
    ON saldos (card_number)
    WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_saldos_card_number_active;
-- +goose StatementEnd
