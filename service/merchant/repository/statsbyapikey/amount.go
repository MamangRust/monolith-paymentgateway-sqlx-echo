package merchantstatsapikeyrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyAmountByApikey = `WITH
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
    AND EXISTS (
        SELECT 1 FROM merchants mch
        WHERE mch.merchant_id = t.merchant_id
        AND mch.deleted_at IS NULL
        AND mch.api_key = $2
    )
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyAmountByApikey = `WITH
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
            AND m.api_key = $1
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

type merchantStatsAmountByApiKeyRepository struct {
	db *sqlx.DB
}

func NewMerchantStatsAmountByApiKeyRepository(db *sqlx.DB) MerchantStatsAmountByApiKeyRepository {
	return &merchantStatsAmountByApiKeyRepository{
		db: db,
	}
}

func (r *merchantStatsAmountByApiKeyRepository) GetMonthlyAmountByApikey(ctx context.Context, req *requests.MonthYearAmountApiKey) ([]*db.GetMonthlyAmountByApikeyRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)
	var amounts []*db.GetMonthlyAmountByApikeyRow
	err := r.db.SelectContext(ctx, &amounts, getMonthlyAmountByApikey,
		yearStart,
		req.Apikey,
	)

	if err != nil {
		return nil, merchant_errors.ErrGetMonthlyAmountByApikeyFailed.WithInternal(err)
	}

	return amounts, nil
}

func (r *merchantStatsAmountByApiKeyRepository) GetYearlyAmountByApikey(ctx context.Context, req *requests.MonthYearAmountApiKey) ([]*db.GetYearlyAmountByApikeyRow, error) {
	var amounts []*db.GetYearlyAmountByApikeyRow
	err := r.db.SelectContext(ctx, &amounts, getYearlyAmountByApikey,
		req.Apikey,
		req.Year,
	)

	if err != nil {
		return nil, merchant_errors.ErrGetYearlyAmountByApikeyFailed.WithInternal(err)
	}

	return amounts, nil
}
