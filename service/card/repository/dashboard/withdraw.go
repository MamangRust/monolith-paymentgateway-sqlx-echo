package repositorydashboard

import (
	"context"

	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getTotalWithdrawAmount = `SELECT SUM(s.withdraw_amount) AS total_withdraw_amount
FROM withdraws s
    JOIN cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL
    AND c.deleted_at IS NULL`

const getTotalWithdrawAmountByCardNumber = `SELECT SUM(w.withdraw_amount) AS total_withdraw_amount
FROM withdraws w
    JOIN cards c ON w.card_number = c.card_number
WHERE
    w.deleted_at IS NULL
    AND c.deleted_at IS NULL
    AND c.card_number = $1`

type cardDashboardWithdrawRepository struct {
	db *sqlx.DB
}

func NewCardDashboardWithdrawRepository(db *sqlx.DB) CardDashboardWithdrawRepository {
	return &cardDashboardWithdrawRepository{
		db: db,
	}
}

func (r *cardDashboardWithdrawRepository) GetTotalWithdrawAmount(ctx context.Context) (*int64, error) {
	var res int64
	err := r.db.GetContext(ctx, &res, getTotalWithdrawAmount)

	if err != nil {
		return nil, card_errors.ErrGetTotalWithdrawsFailed.WithInternal(err)
	}

	return &res, nil
}

func (r *cardDashboardWithdrawRepository) GetTotalWithdrawAmountByCardNumber(ctx context.Context, cardNumber string) (*int64, error) {
	var res int64
	err := r.db.GetContext(ctx, &res, getTotalWithdrawAmountByCardNumber, cardNumber)

	if err != nil {
		return nil, card_errors.ErrGetTotalWithdrawAmountByCardFailed.WithInternal(err)
	}

	return &res, nil
}
