package repositorydashboard

import (
	"context"

	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getTotalBalance = `SELECT SUM(s.total_balance) AS total_balance
FROM saldos s
    JOIN cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL
    AND c.deleted_at IS NULL`

const getTotalBalanceByCardNumber = `SELECT SUM(s.total_balance) AS total_balance
FROM saldos s
    JOIN cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL
    AND c.deleted_at IS NULL
    AND c.card_number = $1`

type cardDashboardBalanceRepository struct {
	db *sqlx.DB
}

func NewCardDashboardBalanceRepository(db *sqlx.DB) CardDashboardBalanceRepository {
	return &cardDashboardBalanceRepository{
		db: db,
	}
}

func (r *cardDashboardBalanceRepository) GetTotalBalances(ctx context.Context) (*int64, error) {
	var res int64
	err := r.db.GetContext(ctx, &res, getTotalBalance)

	if err != nil {
		return nil, card_errors.ErrGetTotalBalancesFailed.WithInternal(err)
	}

	return &res, nil
}

func (r *cardDashboardBalanceRepository) GetTotalBalanceByCardNumber(ctx context.Context, cardNumber string) (*int64, error) {
	var res int64
	err := r.db.GetContext(ctx, &res, getTotalBalanceByCardNumber, cardNumber)

	if err != nil {
		return nil, card_errors.ErrGetTotalBalanceByCardFailed.WithInternal(err)
	}

	return &res, nil
}
