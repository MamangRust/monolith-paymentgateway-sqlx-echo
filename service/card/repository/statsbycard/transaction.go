package repositorystatsbycard

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyTransactionAmountByCardNumber = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $2::timestamp), date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(t.amount), 0)::int AS total_transaction_amount
FROM
    months m
    LEFT JOIN transactions t ON EXTRACT(
        MONTH
        FROM t.transaction_time
    ) = EXTRACT(
        MONTH
        FROM m.month
    )
    AND EXTRACT(
        YEAR
        FROM t.transaction_time
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

const getYearlyTransactionAmountByCardNumber = `WITH
    last_five_years AS (
        SELECT EXTRACT(
                YEAR
                FROM t.transaction_time
            ) AS year, SUM(t.amount) AS total_transaction_amount
        FROM transactions t
            JOIN cards c ON t.card_number = c.card_number
        WHERE
            t.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.card_number = $1
            AND EXTRACT(
                YEAR
                FROM t.transaction_time
            ) >= $2 - 4
            AND EXTRACT(
                YEAR
                FROM t.transaction_time
            ) <= $2
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.transaction_time
            )
    )
SELECT year, total_transaction_amount
FROM last_five_years
ORDER BY year`

type cardStatsTransactionByCardRepository struct {
	db *sqlx.DB
}

func NewCardStatsTransactionByCardRepository(db *sqlx.DB) CardStatsTransactionByCardRepository {
	return &cardStatsTransactionByCardRepository{
		db: db,
	}
}

func (r *cardStatsTransactionByCardRepository) GetMonthlyTransactionAmountByCardNumber(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetMonthlyTransactionAmountByCardNumberRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	var res []*db.GetMonthlyTransactionAmountByCardNumberRow
	err := r.db.SelectContext(ctx, &res, getMonthlyTransactionAmountByCardNumber,
		req.CardNumber,
		yearStart,
	)

	if err != nil {
		return nil, card_errors.ErrGetMonthlyTransactionAmountByCardFailed.WithInternal(err)
	}

	return res, nil
}

func (r *cardStatsTransactionByCardRepository) GetYearlyTransactionAmountByCardNumber(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetYearlyTransactionAmountByCardNumberRow, error) {
	var res []*db.GetYearlyTransactionAmountByCardNumberRow
	err := r.db.SelectContext(ctx, &res, getYearlyTransactionAmountByCardNumber,
		req.CardNumber,
		int32(req.Year),
	)

	if err != nil {
		return nil, card_errors.ErrGetYearlyTransactionAmountByCardFailed.WithInternal(err)
	}

	return res, nil
}
