package repositorydashboard

import (
	"context"

	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getTotalTransactionAmount = `SELECT SUM(t.amount) AS total_transaction_amount
FROM transactions t
    JOIN cards c ON t.card_number = c.card_number
WHERE
    t.deleted_at IS NULL
    AND c.deleted_at IS NULL`

const getTotalTransactionAmountByCardNumber = `SELECT SUM(t.amount) AS total_transaction_amount
FROM transactions t
    JOIN cards c ON t.card_number = c.card_number
WHERE
    t.deleted_at IS NULL
    AND c.deleted_at IS NULL
    AND c.card_number = $1`

type cardDashboardTransactionRepository struct {
	db *sqlx.DB
}

func NewCardDashboardTransactionRepository(db *sqlx.DB) CardDashboardTransactionRepository {
	return &cardDashboardTransactionRepository{
		db: db,
	}
}

func (r *cardDashboardTransactionRepository) GetTotalTransactionAmount(ctx context.Context) (*int64, error) {
	var res int64
	err := r.db.GetContext(ctx, &res, getTotalTransactionAmount)

	if err != nil {
		return nil, card_errors.ErrGetTotalTransactionsFailed.WithInternal(err)
	}

	return &res, nil
}

func (r *cardDashboardTransactionRepository) GetTotalTransactionAmountByCardNumber(ctx context.Context, cardNumber string) (*int64, error) {
	var res int64
	err := r.db.GetContext(ctx, &res, getTotalTransactionAmountByCardNumber, cardNumber)

	if err != nil {
		return nil, card_errors.ErrGetTotalTransactionAmountByCardFailed.WithInternal(err)
	}

	return &res, nil
}
