package merchantstatsmerchantrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyAmountByMerchants = `WITH
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
    AND t.merchant_id = $2
    LEFT JOIN merchants mch ON t.merchant_id = mch.merchant_id
    AND mch.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyAmountByMerchants = `WITH
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
            AND t.merchant_id = $1
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
SELECT year, total_amount
FROM last_five_years
ORDER BY year`

type merchantStatsAmountByMerchantRepository struct {
	db *sqlx.DB
}

func NewMerchantStatsAmountByMerchantRepository(db *sqlx.DB) MerchantStatsAmountByMerchantRepository {
	return &merchantStatsAmountByMerchantRepository{
		db: db,
	}
}

func (r *merchantStatsAmountByMerchantRepository) GetMonthlyAmountByMerchants(ctx context.Context, req *requests.MonthYearAmountMerchant) ([]*db.GetMonthlyAmountByMerchantsRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)
	var amounts []*db.GetMonthlyAmountByMerchantsRow
	err := r.db.SelectContext(ctx, &amounts, getMonthlyAmountByMerchants,
		yearStart,
		int32(req.MerchantID),
	)

	if err != nil {
		return nil, merchant_errors.ErrGetMonthlyAmountByMerchantsFailed.WithInternal(err)
	}

	return amounts, nil
}

func (r *merchantStatsAmountByMerchantRepository) GetYearlyAmountByMerchants(ctx context.Context, req *requests.MonthYearAmountMerchant) ([]*db.GetYearlyAmountByMerchantsRow, error) {
	var amounts []*db.GetYearlyAmountByMerchantsRow
	err := r.db.SelectContext(ctx, &amounts, getYearlyAmountByMerchants,
		int32(req.MerchantID),
		req.Year,
	)

	if err != nil {
		return nil, merchant_errors.ErrGetYearlyAmountByMerchantsFailed.WithInternal(err)
	}

	return amounts, nil
}
