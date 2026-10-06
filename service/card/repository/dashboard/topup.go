package repositorydashboard

import (
	"context"

	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getTotalTopupAmount = `SELECT SUM(t.topup_amount) AS total_topup_amount
FROM topups t
    JOIN cards c ON t.card_number = c.card_number
WHERE
    t.deleted_at IS NULL
    AND c.deleted_at IS NULL`

const getTotalTopupAmountByCardNumber = `SELECT SUM(t.topup_amount) AS total_topup_amount
FROM topups t
    JOIN cards c ON t.card_number = c.card_number
WHERE
    t.deleted_at IS NULL
    AND c.deleted_at IS NULL
    AND c.card_number = $1`

type cardDashboardTopupRepository struct {
	db *sqlx.DB
}

func NewCardDashboardTopupRepository(db *sqlx.DB) CardDashboardTopupRepository {
	return &cardDashboardTopupRepository{
		db: db,
	}
}

func (r *cardDashboardTopupRepository) GetTotalTopAmount(ctx context.Context) (*int64, error) {
	var res int64
	err := r.db.GetContext(ctx, &res, getTotalTopupAmount)

	if err != nil {
		return nil, card_errors.ErrGetTotalTopupsFailed.WithInternal(err)
	}

	return &res, nil
}

func (r *cardDashboardTopupRepository) GetTotalTopupAmountByCardNumber(ctx context.Context, cardNumber string) (*int64, error) {
	var res int64
	err := r.db.GetContext(ctx, &res, getTotalTopupAmountByCardNumber, cardNumber)

	if err != nil {
		return nil, card_errors.ErrGetTotalTopupAmountByCardFailed.WithInternal(err)
	}

	return &res, nil
}
