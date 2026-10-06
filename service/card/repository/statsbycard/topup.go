package repositorystatsbycard

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyTopupAmountByCardNumber = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $2::timestamp), date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
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
    AND t.card_number = $1
    LEFT JOIN cards c ON t.card_number = c.card_number
    AND c.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyTopupAmountByCardNumber = `WITH
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
            AND t.card_number = $1
            AND EXTRACT(
                YEAR
                FROM t.topup_time
            ) >= $2 - 4
            AND EXTRACT(
                YEAR
                FROM t.topup_time
            ) <= $2
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.topup_time
            )
    )
SELECT year, total_topup_amount
FROM last_five_years
ORDER BY year`

type cardStatsTopupByCardRepository struct {
	db *sqlx.DB
}

func NewCardStatsTopupByCardRepository(db *sqlx.DB) CardStatsTopupByCardRepository {
	return &cardStatsTopupByCardRepository{
		db: db,
	}
}

func (r *cardStatsTopupByCardRepository) GetMonthlyTopupAmountByCardNumber(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetMonthlyTopupAmountByCardNumberRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	var res []*db.GetMonthlyTopupAmountByCardNumberRow
	err := r.db.SelectContext(ctx, &res, getMonthlyTopupAmountByCardNumber,
		req.CardNumber,
		yearStart,
	)

	if err != nil {
		return nil, card_errors.ErrGetMonthlyTopupAmountByCardFailed.WithInternal(err)
	}

	return res, nil
}

func (r *cardStatsTopupByCardRepository) GetYearlyTopupAmountByCardNumber(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetYearlyTopupAmountByCardNumberRow, error) {
	var res []*db.GetYearlyTopupAmountByCardNumberRow
	err := r.db.SelectContext(ctx, &res, getYearlyTopupAmountByCardNumber,
		req.CardNumber,
		int32(req.Year),
	)

	if err != nil {
		return nil, card_errors.ErrGetYearlyTopupAmountByCardFailed.WithInternal(err)
	}

	return res, nil
}
