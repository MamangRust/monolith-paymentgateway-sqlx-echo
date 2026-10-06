package topupstatsrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyTopupMethods = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    ),
    topup_methods AS (
        SELECT DISTINCT
            topup_method
        FROM topups
        WHERE
            deleted_at IS NULL
    )
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    tm.topup_method,
    COALESCE(COUNT(t.topup_id), 0)::int AS total_topups,
    COALESCE(SUM(t.topup_amount), 0)::int AS total_amount
FROM
    months m
    CROSS JOIN topup_methods tm
    LEFT JOIN topups t ON EXTRACT(
        MONTH
        FROM t.topup_time
    ) = EXTRACT(
        MONTH
        FROM m.month
    )
    AND EXTRACT(
        YEAR
        FROM t.created_at
    ) = EXTRACT(
        YEAR
        FROM m.month
    )
    AND t.topup_method = tm.topup_method
    AND t.deleted_at IS NULL
GROUP BY
    m.month,
    tm.topup_method
ORDER BY m.month, tm.topup_method`

const getYearlyTopupMethods = `SELECT
    EXTRACT(
        YEAR
        FROM t.created_at
    ) AS year,
    t.topup_method,
    COUNT(t.topup_id) AS total_topups,
    SUM(t.topup_amount) AS total_amount
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
        FROM t.created_at
    ),
    t.topup_method
ORDER BY year`

type topupStatsMethodRepository struct {
	db *sqlx.DB
}

func NewTopupStatsMethodRepository(db *sqlx.DB) TopupStatsMethodRepository {
	return &topupStatsMethodRepository{
		db: db,
	}
}

func (r *topupStatsMethodRepository) GetMonthlyTopupMethods(ctx context.Context, year int) ([]*db.GetMonthlyTopupMethodsRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var rows []*db.GetMonthlyTopupMethodsRow
	err := r.db.SelectContext(ctx, &rows, getMonthlyTopupMethods, yearStart)

	if err != nil {
		return nil, topup_errors.ErrGetMonthlyTopupMethodsFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *topupStatsMethodRepository) GetYearlyTopupMethods(ctx context.Context, year int) ([]*db.GetYearlyTopupMethodsRow, error) {
	var rows []*db.GetYearlyTopupMethodsRow
	err := r.db.SelectContext(ctx, &rows, getYearlyTopupMethods, year)

	if err != nil {
		return nil, topup_errors.ErrGetYearlyTopupMethodsFailed.WithInternal(err)
	}

	return rows, nil
}
