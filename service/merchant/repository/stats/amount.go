package merchantstatsrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyAmountMerchant = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(t.amount), 0)::int AS total_amount
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
    LEFT JOIN merchants mch ON t.merchant_id = mch.merchant_id
    AND mch.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyAmountMerchant = `WITH
    last_five_years AS (
        SELECT EXTRACT(
                YEAR
                FROM t.transaction_time
            ) AS year, SUM(t.amount) AS total_amount
        FROM transactions t
            JOIN merchants m ON t.merchant_id = m.merchant_id
        WHERE
            t.deleted_at IS NULL
            AND m.deleted_at IS NULL
            AND EXTRACT(
                YEAR
                FROM t.transaction_time
            ) >= $1 - 4
            AND EXTRACT(
                YEAR
                FROM t.transaction_time
            ) <= $1
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.transaction_time
            )
    )
SELECT year, total_amount
FROM last_five_years
ORDER BY year`

type merchantStatsAmountRepository struct {
	db *sqlx.DB
}

func NewMerchantStatsAmountRepository(db *sqlx.DB) MerchantStatsAmountRepository {
	return &merchantStatsAmountRepository{
		db: db,
	}
}

func (r *merchantStatsAmountRepository) GetMonthlyAmountMerchant(ctx context.Context, year int) ([]*db.GetMonthlyAmountMerchantRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var amounts []*db.GetMonthlyAmountMerchantRow
	err := r.db.SelectContext(ctx, &amounts, getMonthlyAmountMerchant, yearStart)

	if err != nil {
		return nil, merchant_errors.ErrGetMonthlyAmountMerchantFailed.WithInternal(err)
	}

	return amounts, nil
}

func (r *merchantStatsAmountRepository) GetYearlyAmountMerchant(ctx context.Context, year int) ([]*db.GetYearlyAmountMerchantRow, error) {
	var amounts []*db.GetYearlyAmountMerchantRow
	err := r.db.SelectContext(ctx, &amounts, getYearlyAmountMerchant, year)

	if err != nil {
		return nil, merchant_errors.ErrGetYearlyAmountMerchantFailed.WithInternal(err)
	}

	return amounts, nil
}
