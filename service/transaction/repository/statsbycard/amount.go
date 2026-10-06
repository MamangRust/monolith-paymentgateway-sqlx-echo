package transactionbycardrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyAmountsByCardNumber = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $2::timestamp), date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(t.amount), 0)::int AS total_amount
FROM
    months m
    LEFT JOIN transactions t ON EXTRACT(
        MONTH
        FROM t.transaction_time
    ) = EXTRACT(
        MONTH
        FROM m.month
    )
    AND EXTRACT(
        YEAR
        FROM t.transaction_time
    ) = EXTRACT(
        YEAR
        FROM m.month
    )
    AND t.card_number = $1
    AND t.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyAmountsByCardNumber = `SELECT EXTRACT(
        YEAR
        FROM t.created_at
    ) AS year, SUM(t.amount) AS total_amount
FROM transactions t
WHERE
    t.deleted_at IS NULL
    AND t.card_number = $1
    AND EXTRACT(
        YEAR
        FROM t.created_at
    ) >= $2 - 4
    AND EXTRACT(
        YEAR
        FROM t.created_at
    ) <= $2
GROUP BY
    EXTRACT(
        YEAR
        FROM t.created_at
    )
ORDER BY year`

type transactionStatsByCardAmountRepository struct {
	db *sqlx.DB
}

func NewTransactionStatsByCardAmountRepository(db *sqlx.DB) TransactionStatsByCardAmountRepository {
	return &transactionStatsByCardAmountRepository{
		db: db,
	}
}

func (r *transactionStatsByCardAmountRepository) GetMonthlyAmountsByCardNumber(ctx context.Context, req *requests.MonthYearPaymentMethod) ([]*db.GetMonthlyAmountsByCardNumberRow, error) {
	cardNumber := req.CardNumber
	year := req.Year

	var rows []*db.GetMonthlyAmountsByCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getMonthlyAmountsByCardNumber,
		cardNumber,
		time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC),
	)
	if err != nil {
		return nil, transaction_errors.ErrGetMonthlyAmountsByCardFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *transactionStatsByCardAmountRepository) GetYearlyAmountsByCardNumber(ctx context.Context, req *requests.MonthYearPaymentMethod) ([]*db.GetYearlyAmountsByCardNumberRow, error) {
	cardNumber := req.CardNumber
	year := req.Year

	var rows []*db.GetYearlyAmountsByCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getYearlyAmountsByCardNumber,
		cardNumber,
		year,
	)
	if err != nil {
		return nil, transaction_errors.ErrGetYearlyAmountsByCardFailed.WithInternal(err)
	}

	return rows, nil
}
