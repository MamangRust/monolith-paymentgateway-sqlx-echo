package repositorystats

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyTopupAmount = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(t.topup_amount), 0)::int AS total_topup_amount
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
    AND t.deleted_at IS NULL
    LEFT JOIN cards c ON t.card_number = c.card_number
    AND c.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyTopupAmount = `WITH
    last_five_years AS (
        SELECT EXTRACT(
                YEAR
                FROM t.topup_time
            ) AS year, SUM(t.topup_amount) AS total_topup_amount
        FROM topups t
            JOIN cards c ON t.card_number = c.card_number
        WHERE
            t.deleted_at IS NULL
            AND c.deleted_at IS NULL
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
    )
SELECT year, total_topup_amount
FROM last_five_years
ORDER BY year`

type cardStatsTopupRepository struct {
	db *sqlx.DB
}

func NewCardStatsTopupRepository(db *sqlx.DB) CardStatsTopupRepository {
	return &cardStatsTopupRepository{
		db: db,
	}
}

func (r *cardStatsTopupRepository) GetMonthlyTopupAmount(ctx context.Context, year int) ([]*db.GetMonthlyTopupAmountRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var res []*db.GetMonthlyTopupAmountRow
	err := r.db.SelectContext(ctx, &res, getMonthlyTopupAmount, yearStart)

	if err != nil {
		return nil, card_errors.ErrGetMonthlyTopupAmountFailed.WithInternal(err)
	}

	return res, nil
}

func (r *cardStatsTopupRepository) GetYearlyTopupAmount(ctx context.Context, year int) ([]*db.GetYearlyTopupAmountRow, error) {
	var res []*db.GetYearlyTopupAmountRow
	err := r.db.SelectContext(ctx, &res, getYearlyTopupAmount, year)

	if err != nil {
		return nil, card_errors.ErrGetYearlyTopupAmountFailed.WithInternal(err)
	}

	return res, nil
}
