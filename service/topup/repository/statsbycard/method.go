package topupstatsbycardrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyTopupMethodsByCardNumber = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $2::timestamp), date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
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
        FROM t.topup_time
    ) = EXTRACT(
        YEAR
        FROM m.month
    )
    AND t.topup_method = tm.topup_method
    AND t.card_number = $1
    AND t.deleted_at IS NULL
GROUP BY
    m.month,
    tm.topup_method
ORDER BY m.month, tm.topup_method`

const getYearlyTopupMethodsByCardNumber = `SELECT
    EXTRACT(
        YEAR
        FROM t.topup_time
    ) AS year,
    t.topup_method,
    COUNT(t.topup_id) AS total_topups,
    SUM(t.topup_amount) AS total_amount
FROM topups t
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
        FROM t.topup_time
    ),
    t.topup_method
ORDER BY year`

type topupStatsByCardMethodRepository struct {
	db *sqlx.DB
}

func NewTopupStatsByCardMethodRepository(db *sqlx.DB) TopupStatsByCardMethodRepository {
	return &topupStatsByCardMethodRepository{
		db: db,
	}
}

func (r *topupStatsByCardMethodRepository) GetMonthlyTopupMethodsByCardNumber(ctx context.Context, req *requests.YearMonthMethod) ([]*db.GetMonthlyTopupMethodsByCardNumberRow, error) {
	year := req.Year
	cardNumber := req.CardNumber

	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var rows []*db.GetMonthlyTopupMethodsByCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getMonthlyTopupMethodsByCardNumber,
		cardNumber,
		yearStart,
	)

	if err != nil {
		return nil, topup_errors.ErrGetMonthlyTopupMethodsByCardFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *topupStatsByCardMethodRepository) GetYearlyTopupMethodsByCardNumber(ctx context.Context, req *requests.YearMonthMethod) ([]*db.GetYearlyTopupMethodsByCardNumberRow, error) {
	year := req.Year
	cardNumber := req.CardNumber

	var rows []*db.GetYearlyTopupMethodsByCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getYearlyTopupMethodsByCardNumber,
		cardNumber,
		year,
	)

	if err != nil {
		return nil, topup_errors.ErrGetYearlyTopupMethodsByCardFailed.WithInternal(err)
	}

	return rows, nil
}
