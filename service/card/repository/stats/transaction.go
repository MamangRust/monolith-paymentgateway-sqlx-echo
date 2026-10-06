package repositorystats

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyTransactionAmount = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(t.amount), 0)::int AS total_transaction_amount
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
    LEFT JOIN cards c ON t.card_number = c.card_number
    AND c.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyTransactionAmount = `WITH
    last_five_years AS (
        SELECT EXTRACT(
                YEAR
                FROM t.transaction_time
            ) AS year, SUM(t.amount) AS total_transaction_amount
        FROM transactions t
            JOIN cards c ON t.card_number = c.card_number
        WHERE
            t.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND EXTRACT(
                YEAR
                FROM t.transaction_time
            ) >= $1 - 4
            AND EXTRACT(
                YEAR
                FROM t.transaction_time
            ) <= $1
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.transaction_time
            )
    )
SELECT year, total_transaction_amount
FROM last_five_years
ORDER BY year`

type cardStatsTransactionRepository struct {
	db *sqlx.DB
}

func NewCardStatsTransactionRepository(db *sqlx.DB) CardStatsTransactionRepository {
	return &cardStatsTransactionRepository{
		db: db,
	}
}

func (r *cardStatsTransactionRepository) GetMonthlyTransactionAmount(ctx context.Context, year int) ([]*db.GetMonthlyTransactionAmountRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var res []*db.GetMonthlyTransactionAmountRow
	err := r.db.SelectContext(ctx, &res, getMonthlyTransactionAmount, yearStart)

	if err != nil {
		return nil, card_errors.ErrGetMonthlyTransactionAmountFailed.WithInternal(err)
	}

	return res, nil
}

func (r *cardStatsTransactionRepository) GetYearlyTransactionAmount(ctx context.Context, year int) ([]*db.GetYearlyTransactionAmountRow, error) {
	var res []*db.GetYearlyTransactionAmountRow
	err := r.db.SelectContext(ctx, &res, getYearlyTransactionAmount, year)

	if err != nil {
		return nil, card_errors.ErrGetYearlyTransactionAmountFailed.WithInternal(err)
	}

	return res, nil
}
