package transactionstatsrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyPaymentMethods = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
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
    AND t.deleted_at IS NULL
GROUP BY
    m.month,
    pm.payment_method
ORDER BY m.month, pm.payment_method`

const getYearlyPaymentMethods = `SELECT
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
    AND EXTRACT(
        YEAR
        FROM t.created_at
    ) >= $1 - 4
    AND EXTRACT(
        YEAR
        FROM t.created_at
    ) <= $1
GROUP BY
    EXTRACT(
        YEAR
        FROM t.created_at
    ),
    t.payment_method
ORDER BY year`

type transactionStatsMethodRepository struct {
	db *sqlx.DB
}

func NewTransactionStatsMethodRepository(db *sqlx.DB) TransactionStatsMethodRepository {
	return &transactionStatsMethodRepository{
		db: db,
	}
}

func (r *transactionStatsMethodRepository) GetMonthlyPaymentMethods(ctx context.Context, year int) ([]*db.GetMonthlyPaymentMethodsRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var rows []*db.GetMonthlyPaymentMethodsRow
	err := r.db.SelectContext(ctx, &rows, getMonthlyPaymentMethods, yearStart)

	if err != nil {
		return nil, transaction_errors.ErrGetMonthlyPaymentMethodsFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *transactionStatsMethodRepository) GetYearlyPaymentMethods(ctx context.Context, year int) ([]*db.GetYearlyPaymentMethodsRow, error) {
	var rows []*db.GetYearlyPaymentMethodsRow
	err := r.db.SelectContext(ctx, &rows, getYearlyPaymentMethods, year)

	if err != nil {
		return nil, transaction_errors.ErrGetYearlyPaymentMethodsFailed.WithInternal(err)
	}

	return rows, nil
}
