-- +goose Up
-- +goose StatementBegin

-- Fail explicitly if legacy data contains duplicate non-empty references.
DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM card_payments
        WHERE reference_id <> ''
        GROUP BY reference_id
        HAVING COUNT(*) > 1
    ) THEN
        RAISE EXCEPTION 'cannot create card payment idempotency index: duplicate non-empty reference_id values exist';
    END IF;
END $$;

-- Fail explicitly if legacy data contains duplicate billing periods. The
-- scheduler relies on this constraint for idempotent statement generation.
DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM billing_cycles
        GROUP BY card_number, cycle_start, cycle_end
        HAVING COUNT(*) > 1
    ) THEN
        RAISE EXCEPTION 'cannot create billing cycle idempotency index: duplicate card billing periods exist';
    END IF;
END $$;

-- reference_id is the client-supplied idempotency key for card payments.
-- Empty values remain allowed for legacy/internal callers.
CREATE UNIQUE INDEX IF NOT EXISTS idx_card_payments_reference_id
    ON card_payments (reference_id)
    WHERE reference_id <> '';

CREATE UNIQUE INDEX IF NOT EXISTS idx_billing_cycles_card_period
    ON billing_cycles (card_number, cycle_start, cycle_end);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS idx_billing_cycles_card_period;
DROP INDEX IF EXISTS idx_card_payments_reference_id;

-- +goose StatementEnd
