package repository

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getBillingCycleByID = `SELECT
    billing_id,
    card_number,
    cycle_start,
    cycle_end,
    amount_due,
    due_date,
    status,
    created_at,
    updated_at
FROM billing_cycles
WHERE billing_id = $1`

const getBillingCyclesByCardNumber = `SELECT
    billing_id,
    card_number,
    cycle_start,
    cycle_end,
    amount_due,
    due_date,
    status,
    created_at,
    updated_at
FROM billing_cycles
WHERE
    card_number = $1
ORDER BY cycle_start DESC`

type billingCycleRepository struct {
	db *sqlx.DB
}

func NewBillingCycleRepository(db *sqlx.DB) BillingCycleRepository {
	return &billingCycleRepository{
		db: db,
	}
}

func (r *billingCycleRepository) GetBillingCycleByID(ctx context.Context, billingID int) (*db.BillingCycle, error) {
	var res db.BillingCycle
	err := r.db.GetContext(ctx, &res, getBillingCycleByID, int32(billingID))
	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Billing cycle", "get billing cycle")
	}
	return &res, nil
}

func (r *billingCycleRepository) GetBillingCyclesByCardNumber(ctx context.Context, cardNumber string) ([]*db.BillingCycle, error) {
	var res []*db.BillingCycle
	err := r.db.SelectContext(ctx, &res, getBillingCyclesByCardNumber, cardNumber)
	if err != nil {
		return nil, card_errors.ErrGetBillingCyclesFailed.WithInternal(err)
	}

	return res, nil
}
