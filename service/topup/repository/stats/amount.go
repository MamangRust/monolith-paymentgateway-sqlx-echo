package topupstatsrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyTopupAmounts = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(t.topup_amount), 0)::int AS total_amount
FROM months m
    LEFT JOIN topups t ON EXTRACT(
        MONTH
        FROM t.topup_time
    ) = EXTRACT(
        MONTH
        FROM m.month
    )
    AND EXTRACT(
        YEAR
        FROM t.topup_time
    ) = EXTRACT(
        YEAR
        FROM m.month
    )
    AND t.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyTopupAmounts = `SELECT EXTRACT(
        YEAR
        FROM t.topup_time
    ) AS year, SUM(t.topup_amount) AS total_amount
FROM topups t
WHERE
    t.deleted_at IS NULL
    AND EXTRACT(
        YEAR
        FROM t.topup_time
    ) >= $1 - 4
    AND EXTRACT(
        YEAR
        FROM t.topup_time
    ) <= $1
GROUP BY
    EXTRACT(
        YEAR
        FROM t.topup_time
    )
ORDER BY year`

type topupStatsAmountRepository struct {
	db *sqlx.DB
}

func NewTopupStatsAmountRepository(db *sqlx.DB) TopupStatsAmountRepository {
	return &topupStatsAmountRepository{
		db: db,
	}
}

func (r *topupStatsAmountRepository) GetMonthlyTopupAmounts(ctx context.Context, year int) ([]*db.GetMonthlyTopupAmountsRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var rows []*db.GetMonthlyTopupAmountsRow
	err := r.db.SelectContext(ctx, &rows, getMonthlyTopupAmounts, yearStart)

	if err != nil {
		return nil, topup_errors.ErrGetMonthlyTopupAmountsFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *topupStatsAmountRepository) GetYearlyTopupAmounts(ctx context.Context, year int) ([]*db.GetYearlyTopupAmountsRow, error) {
	var rows []*db.GetYearlyTopupAmountsRow
	err := r.db.SelectContext(ctx, &rows, getYearlyTopupAmounts, year)

	if err != nil {
		return nil, topup_errors.ErrGetYearlyTopupAmountsFailed.WithInternal(err)
	}

	return rows, nil
}
