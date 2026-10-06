package transactionstatsrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyAmounts = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
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
    AND t.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyAmounts = `SELECT EXTRACT(
        YEAR
        FROM t.created_at
    ) AS year, SUM(t.amount) AS total_amount
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
    )
ORDER BY year`

type transactionStatsAmountRepository struct {
	db *sqlx.DB
}

func NewTransactionStatsAmountRepository(db *sqlx.DB) TransactionStatsAmountRepository {
	return &transactionStatsAmountRepository{
		db: db,
	}
}

func (r *transactionStatsAmountRepository) GetMonthlyAmounts(ctx context.Context, year int) ([]*db.GetMonthlyAmountsRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var rows []*db.GetMonthlyAmountsRow
	err := r.db.SelectContext(ctx, &rows, getMonthlyAmounts, yearStart)

	if err != nil {
		return nil, transaction_errors.ErrGetMonthlyAmountsFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *transactionStatsAmountRepository) GetYearlyAmounts(ctx context.Context, year int) ([]*db.GetYearlyAmountsRow, error) {
	var rows []*db.GetYearlyAmountsRow
	err := r.db.SelectContext(ctx, &rows, getYearlyAmounts, year)

	if err != nil {
		return nil, transaction_errors.ErrGetYearlyAmountsFailed.WithInternal(err)
	}

	return rows, nil
}
