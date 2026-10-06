package topupstatsbycardrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyTopupAmountsByCardNumber = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $2::timestamp), date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(t.topup_amount), 0)::int AS total_amount
FROM
    months m
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
    AND t.card_number = $1
    AND t.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyTopupAmountsByCardNumber = `SELECT EXTRACT(
        YEAR
        FROM t.topup_time
    ) AS year, SUM(t.topup_amount) AS total_amount
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
    )
ORDER BY year`

type topupStatsByCardAmountRepository struct {
	db *sqlx.DB
}

func NewTopupStatsByCardAmountRepository(db *sqlx.DB) TopupStatsByCardAmountRepository {
	return &topupStatsByCardAmountRepository{
		db: db,
	}
}

func (r *topupStatsByCardAmountRepository) GetMonthlyTopupAmountsByCardNumber(ctx context.Context, req *requests.YearMonthMethod) ([]*db.GetMonthlyTopupAmountsByCardNumberRow, error) {
	year := req.Year
	cardNumber := req.CardNumber

	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var rows []*db.GetMonthlyTopupAmountsByCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getMonthlyTopupAmountsByCardNumber,
		cardNumber,
		yearStart,
	)

	if err != nil {
		return nil, topup_errors.ErrGetMonthlyTopupAmountsByCardFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *topupStatsByCardAmountRepository) GetYearlyTopupAmountsByCardNumber(ctx context.Context, req *requests.YearMonthMethod) ([]*db.GetYearlyTopupAmountsByCardNumberRow, error) {
	year := req.Year
	cardNumber := req.CardNumber

	var rows []*db.GetYearlyTopupAmountsByCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getYearlyTopupAmountsByCardNumber,
		cardNumber,
		year,
	)

	if err != nil {
		return nil, topup_errors.ErrGetYearlyTopupAmountsByCardFailed.WithInternal(err)
	}

	return rows, nil
}
