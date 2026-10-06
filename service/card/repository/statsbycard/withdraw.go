package repositorystatsbycard

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyWithdrawAmountByCardNumber = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $2::timestamp), date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
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
    AND w.card_number = $1
    LEFT JOIN cards c ON w.card_number = c.card_number
    AND c.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyWithdrawAmountByCardNumber = `WITH
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
            AND w.card_number = $1
            AND EXTRACT(
                YEAR
                FROM w.withdraw_time
            ) >= $2 - 4
            AND EXTRACT(
                YEAR
                FROM w.withdraw_time
            ) <= $2
        GROUP BY
            EXTRACT(
                YEAR
                FROM w.withdraw_time
            )
    )
SELECT year, total_withdraw_amount
FROM last_five_years
ORDER BY year`

type cardStatsWithdrawByCardRepository struct {
	db *sqlx.DB
}

func NewCardStatsWithdrawByCardRepository(db *sqlx.DB) CardStatsWithdrawByCardRepository {
	return &cardStatsWithdrawByCardRepository{
		db: db,
	}
}

func (r *cardStatsWithdrawByCardRepository) GetMonthlyWithdrawAmountByCardNumber(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetMonthlyWithdrawAmountByCardNumberRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	var res []*db.GetMonthlyWithdrawAmountByCardNumberRow
	err := r.db.SelectContext(ctx, &res, getMonthlyWithdrawAmountByCardNumber,
		req.CardNumber,
		yearStart,
	)

	if err != nil {
		return nil, card_errors.ErrGetMonthlyWithdrawAmountByCardFailed.WithInternal(err)
	}

	return res, nil
}

func (r *cardStatsWithdrawByCardRepository) GetYearlyWithdrawAmountByCardNumber(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetYearlyWithdrawAmountByCardNumberRow, error) {
	var res []*db.GetYearlyWithdrawAmountByCardNumberRow
	err := r.db.SelectContext(ctx, &res, getYearlyWithdrawAmountByCardNumber,
		req.CardNumber,
		int32(req.Year),
	)

	if err != nil {
		return nil, card_errors.ErrGetYearlyWithdrawAmountByCardFailed.WithInternal(err)
	}

	return res, nil
}
