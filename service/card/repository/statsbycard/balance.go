package repositorystatsbycard

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyBalancesByCardNumber = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(s.total_balance), 0)::int AS total_balance
FROM
    months m
    LEFT JOIN saldos s ON EXTRACT(
        MONTH
        FROM s.created_at
    ) = EXTRACT(
        MONTH
        FROM m.month
    )
    AND EXTRACT(
        YEAR
        FROM s.created_at
    ) = EXTRACT(
        YEAR
        FROM m.month
    )
    AND s.card_number = $2
    LEFT JOIN cards c ON s.card_number = c.card_number
    AND c.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyBalancesByCardNumber = `WITH
    last_five_years AS (
        SELECT EXTRACT(
                YEAR
                FROM s.created_at
            ) AS year, SUM(s.total_balance) AS total_balance
        FROM saldos s
            JOIN cards c ON s.card_number = c.card_number
        WHERE
            c.deleted_at IS NULL
            AND EXTRACT(
                YEAR
                FROM s.created_at
            ) >= $1 - 4
            AND EXTRACT(
                YEAR
                FROM s.created_at
            ) <= $1
            AND c.card_number = $2
        GROUP BY
            EXTRACT(
                YEAR
                FROM s.created_at
            )
    )
SELECT year, total_balance
FROM last_five_years
ORDER BY year`

type cardStatsBalanceByCardRepository struct {
	db *sqlx.DB
}

func NewCardStatsBalanceByCardRepository(db *sqlx.DB) CardStatsBalanceByCardRepository {
	return &cardStatsBalanceByCardRepository{
		db: db,
	}
}

func (r *cardStatsBalanceByCardRepository) GetMonthlyBalancesByCardNumber(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetMonthlyBalancesByCardNumberRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	var res []*db.GetMonthlyBalancesByCardNumberRow
	err := r.db.SelectContext(ctx, &res, getMonthlyBalancesByCardNumber,
		yearStart,
		req.CardNumber,
	)

	if err != nil {
		return nil, card_errors.ErrGetMonthlyBalanceByCardFailed.WithInternal(err)
	}

	return res, nil
}

func (r *cardStatsBalanceByCardRepository) GetYearlyBalanceByCardNumber(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetYearlyBalancesByCardNumberRow, error) {
	var res []*db.GetYearlyBalancesByCardNumberRow
	err := r.db.SelectContext(ctx, &res, getYearlyBalancesByCardNumber,
		req.Year,
		req.CardNumber,
	)

	if err != nil {
		return nil, card_errors.ErrGetYearlyBalanceByCardFailed.WithInternal(err)
	}

	return res, nil
}
