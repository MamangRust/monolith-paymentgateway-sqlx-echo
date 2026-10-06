package repositorystats

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyBalances = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(s.total_balance), 0)::int AS total_balance
FROM
    months m
    LEFT JOIN saldos s ON EXTRACT(
        MONTH
        FROM s.created_at
    ) = EXTRACT(
        MONTH
        FROM m.month
    )
    AND EXTRACT(
        YEAR
        FROM s.created_at
    ) = EXTRACT(
        YEAR
        FROM m.month
    )
    LEFT JOIN cards c ON s.card_number = c.card_number
    AND c.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyBalances = `WITH
    last_five_years AS (
        SELECT EXTRACT(
                YEAR
                FROM s.created_at
            ) AS year, SUM(s.total_balance) AS total_balance
        FROM saldos s
            JOIN cards c ON s.card_number = c.card_number
        WHERE
            c.deleted_at IS NULL
            AND EXTRACT(
                YEAR
                FROM s.created_at
            ) >= $1 - 4
            AND EXTRACT(
                YEAR
                FROM s.created_at
            ) <= $1
        GROUP BY
            EXTRACT(
                YEAR
                FROM s.created_at
            )
    )
SELECT year, total_balance
FROM last_five_years
ORDER BY year`

type cardStatsBalanceRepository struct {
	db *sqlx.DB
}

func NewCardStatsBalanceRepository(db *sqlx.DB) CardStatsBalanceRepository {
	return &cardStatsBalanceRepository{
		db: db,
	}
}

func (r *cardStatsBalanceRepository) GetMonthlyBalance(ctx context.Context, year int) ([]*db.GetMonthlyBalancesRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var res []*db.GetMonthlyBalancesRow
	err := r.db.SelectContext(ctx, &res, getMonthlyBalances, yearStart)

	if err != nil {
		return nil, card_errors.ErrGetMonthlyBalanceFailed.WithInternal(err)
	}

	return res, nil
}

func (r *cardStatsBalanceRepository) GetYearlyBalance(ctx context.Context, year int) ([]*db.GetYearlyBalancesRow, error) {
	var res []*db.GetYearlyBalancesRow
	err := r.db.SelectContext(ctx, &res, getYearlyBalances, year)

	if err != nil {
		return nil, card_errors.ErrGetYearlyBalanceFailed.WithInternal(err)
	}

	return res, nil
}
