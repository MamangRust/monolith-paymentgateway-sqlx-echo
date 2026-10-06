package withdrawstatsrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	withdraw_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/withdraw_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyWithdraws = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(w.withdraw_amount), 0)::int AS total_withdraw_amount
FROM months m
    LEFT JOIN withdraws w ON EXTRACT(
        MONTH
        FROM w.withdraw_time
    ) = EXTRACT(
        MONTH
        FROM m.month
    )
    AND EXTRACT(
        YEAR
        FROM w.withdraw_time
    ) = EXTRACT(
        YEAR
        FROM m.month
    )
    AND w.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyWithdraws = `SELECT EXTRACT(
        YEAR
        FROM w.withdraw_time
    ) AS year, SUM(w.withdraw_amount) AS total_withdraw_amount
FROM withdraws w
WHERE
    w.deleted_at IS NULL
    AND EXTRACT(
        YEAR
        FROM w.withdraw_time
    ) >= $1 - 4
    AND EXTRACT(
        YEAR
        FROM w.withdraw_time
    ) <= $1
GROUP BY
    EXTRACT(
        YEAR
        FROM w.withdraw_time
    )
ORDER BY year`

type withdrawStatsAmountRepository struct {
	db *sqlx.DB
}

func NewWithdrawStatsAmountRepository(db *sqlx.DB) WithdrawStatsAmountRepository {
	return &withdrawStatsAmountRepository{
		db: db,
	}
}

func (r *withdrawStatsAmountRepository) GetMonthlyWithdraws(ctx context.Context, year int) ([]*db.GetMonthlyWithdrawsRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var rows []*db.GetMonthlyWithdrawsRow
	err := r.db.SelectContext(ctx, &rows, getMonthlyWithdraws, yearStart)

	if err != nil {
		return nil, withdraw_errors.ErrGetMonthlyWithdrawsFailed.WithInternal(err)
	}

	return rows, nil

}

func (r *withdrawStatsAmountRepository) GetYearlyWithdraws(ctx context.Context, year int) ([]*db.GetYearlyWithdrawsRow, error) {
	var rows []*db.GetYearlyWithdrawsRow
	err := r.db.SelectContext(ctx, &rows, getYearlyWithdraws, year)

	if err != nil {
		return nil, withdraw_errors.ErrGetYearlyWithdrawsFailed.WithInternal(err)
	}

	return rows, nil

}
