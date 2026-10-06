package repositorystats

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyWithdrawAmount = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(w.withdraw_amount), 0)::int AS total_withdraw_amount
FROM
    months m
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
    LEFT JOIN cards c ON w.card_number = c.card_number
    AND c.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyWithdrawAmount = `WITH
    last_five_years AS (
        SELECT EXTRACT(
                YEAR
                FROM w.withdraw_time
            ) AS year, SUM(w.withdraw_amount) AS total_withdraw_amount
        FROM withdraws w
            JOIN cards c ON w.card_number = c.card_number
        WHERE
            w.deleted_at IS NULL
            AND c.deleted_at IS NULL
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
    )
SELECT year, total_withdraw_amount
FROM last_five_years
ORDER BY year`

type cardStatsWithdrawRepository struct {
	db *sqlx.DB
}

func NewCardStatsWithdrawRepository(db *sqlx.DB) CardStatsWithdrawRepository {
	return &cardStatsWithdrawRepository{
		db: db,
	}
}

func (r *cardStatsWithdrawRepository) GetMonthlyWithdrawAmount(ctx context.Context, year int) ([]*db.GetMonthlyWithdrawAmountRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var res []*db.GetMonthlyWithdrawAmountRow
	err := r.db.SelectContext(ctx, &res, getMonthlyWithdrawAmount, yearStart)

	if err != nil {
		return nil, card_errors.ErrGetMonthlyWithdrawAmountFailed
	}

	return res, nil
}

func (r *cardStatsWithdrawRepository) GetYearlyWithdrawAmount(ctx context.Context, year int) ([]*db.GetYearlyWithdrawAmountRow, error) {
	var res []*db.GetYearlyWithdrawAmountRow
	err := r.db.SelectContext(ctx, &res, getYearlyWithdrawAmount, int32(year))

	if err != nil {
		return nil, card_errors.ErrGetYearlyWithdrawAmountFailed
	}

	return res, nil
}
