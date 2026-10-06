-- +goose Up
-- +goose StatementBegin

-- Backfill durable statement events for billing cycles created before the
-- billing-cycle CTE began enqueueing them. The unique event key makes this
-- safe to run repeatedly.
INSERT INTO outbox_events (event_key, topic, message_key, payload)
SELECT
    'card-statement:' || b.billing_id::text,
    'card.statement.generated',
    b.billing_id::text,
    jsonb_build_object(
        'billing_id', b.billing_id,
        'card_number', b.card_number,
        'cycle_start', b.cycle_start,
        'cycle_end', b.cycle_end,
        'amount_due', b.amount_due
    )
FROM billing_cycles b
WHERE NOT EXISTS (
    SELECT 1
    FROM outbox_events e
    WHERE e.event_key = 'card-statement:' || b.billing_id::text
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Keep already-created outbox records on rollback; deleting them could lose
-- audit events that are still pending delivery.

-- +goose StatementEnd
