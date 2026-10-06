package transactionbycardrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyPaymentMethodsByCardNumber = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $2::timestamp), date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    ),
    payment_methods AS (
        SELECT DISTINCT
            payment_method
        FROM transactions
        WHERE
            deleted_at IS NULL
    )
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    pm.payment_method,
    COALESCE(COUNT(t.transaction_id), 0)::int AS total_transactions,
    COALESCE(SUM(t.amount), 0)::int AS total_amount
FROM
    months m
    CROSS JOIN payment_methods pm
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
    AND t.payment_method = pm.payment_method
    AND t.card_number = $1
    AND t.deleted_at IS NULL
GROUP BY
    m.month,
    pm.payment_method
ORDER BY m.month, pm.payment_method`

const getYearlyPaymentMethodsByCardNumber = `SELECT
    EXTRACT(
        YEAR
        FROM t.created_at
    ) AS year,
    t.payment_method,
    COUNT(t.transaction_id) AS total_transactions,
    SUM(t.amount) AS total_amount
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
    ),
    t.payment_method
ORDER BY year`

type transactionStatsByCardMethodRepository struct {
	db *sqlx.DB
}

func NewTransactionStatsByCardMethodRepository(db *sqlx.DB) TransactionStatsByCardMethodRepository {
	return &transactionStatsByCardMethodRepository{
		db: db,
	}
}

func (r *transactionStatsByCardMethodRepository) GetMonthlyPaymentMethodsByCardNumber(ctx context.Context, req *requests.MonthYearPaymentMethod) ([]*db.GetMonthlyPaymentMethodsByCardNumberRow, error) {
	year := req.Year
	cardNumber := req.CardNumber

	var rows []*db.GetMonthlyPaymentMethodsByCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getMonthlyPaymentMethodsByCardNumber,
		cardNumber,
		time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC),
	)

	if err != nil {
		return nil, transaction_errors.ErrGetMonthlyPaymentMethodsByCardFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *transactionStatsByCardMethodRepository) GetYearlyPaymentMethodsByCardNumber(ctx context.Context, req *requests.MonthYearPaymentMethod) ([]*db.GetYearlyPaymentMethodsByCardNumberRow, error) {
	year := req.Year
	cardNumber := req.CardNumber

	var rows []*db.GetYearlyPaymentMethodsByCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getYearlyPaymentMethodsByCardNumber,
		cardNumber,
		year,
	)

	if err != nil {
		return nil, transaction_errors.ErrGetYearlyPaymentMethodsByCardFailed.WithInternal(err)
	}

	return rows, nil
}
