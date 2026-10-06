package withdrawstatsbycardrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	withdraw_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/withdraw_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyWithdrawsByCardNumber = `WITH
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
    AND w.card_number = $1
    AND w.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyWithdrawsByCardNumber = `SELECT EXTRACT(
        YEAR
        FROM w.created_at
    ) AS year, SUM(w.withdraw_amount) AS total_withdraw_amount
FROM withdraws w
WHERE
    w.deleted_at IS NULL
    AND w.card_number = $1
    AND EXTRACT(
        YEAR
        FROM w.created_at
    ) >= $2 - 4
    AND EXTRACT(
        YEAR
        FROM w.created_at
    ) <= $2
GROUP BY
    EXTRACT(
        YEAR
        FROM w.created_at
    )
ORDER BY year`

type withdrawStatsByCardAmountRepository struct {
	db *sqlx.DB
}

func NewWithdrawStatsByCardAmountRepository(db *sqlx.DB) WithdrawStatsByCardAmountRepository {
	return &withdrawStatsByCardAmountRepository{
		db: db,
	}
}

func (r *withdrawStatsByCardAmountRepository) GetMonthlyWithdrawsByCardNumber(ctx context.Context, req *requests.YearMonthCardNumber) ([]*db.GetMonthlyWithdrawsByCardNumberRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	var rows []*db.GetMonthlyWithdrawsByCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getMonthlyWithdrawsByCardNumber,
		req.CardNumber,
		yearStart,
	)

	if err != nil {
		return nil, withdraw_errors.ErrGetMonthlyWithdrawsByCardFailed.WithInternal(err)
	}

	return rows, nil

}

func (r *withdrawStatsByCardAmountRepository) GetYearlyWithdrawsByCardNumber(ctx context.Context, req *requests.YearMonthCardNumber) ([]*db.GetYearlyWithdrawsByCardNumberRow, error) {
	var rows []*db.GetYearlyWithdrawsByCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getYearlyWithdrawsByCardNumber,
		req.CardNumber,
		req.Year,
	)

	if err != nil {
		return nil, withdraw_errors.ErrGetYearlyWithdrawsByCardFailed.WithInternal(err)
	}

	return rows, nil
}
