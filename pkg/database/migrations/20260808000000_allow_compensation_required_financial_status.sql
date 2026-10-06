-- +goose Up
-- +goose StatementBegin

-- compensation_required is 21 characters; the original VARCHAR(20) columns
-- cannot persist it. Do not add a new CHECK here because older deployments
-- legitimately contain domain-specific values such as completed/processing.
ALTER TABLE topups       ALTER COLUMN status TYPE VARCHAR(32);
ALTER TABLE transfers    ALTER COLUMN status TYPE VARCHAR(32);
ALTER TABLE withdraws    ALTER COLUMN status TYPE VARCHAR(32);
ALTER TABLE transactions ALTER COLUMN status TYPE VARCHAR(32);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Keep the wider type on rollback: shrinking could make existing
-- compensation_required records unreadable or fail the rollback.

-- +goose StatementEnd
