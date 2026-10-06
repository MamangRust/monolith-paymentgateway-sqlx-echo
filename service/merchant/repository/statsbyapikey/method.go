package merchantstatsapikeyrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyPaymentMethodByApikey = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    ),
    payment_methods AS (
        SELECT DISTINCT
            payment_method
        FROM transactions
        WHERE
            deleted_at IS NULL
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, pm.payment_method, COALESCE(SUM(t.amount), 0)::int AS total_amount
FROM
    months m
    CROSS JOIN payment_methods pm
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
    AND t.payment_method = pm.payment_method
    AND t.deleted_at IS NULL
    AND EXISTS (
        SELECT 1 FROM merchants mch
        WHERE mch.merchant_id = t.merchant_id
        AND mch.deleted_at IS NULL
        AND mch.api_key = $2
    )
GROUP BY
    m.month,
    pm.payment_method
ORDER BY m.month, pm.payment_method`

const getYearlyPaymentMethodByApikey = `WITH
    last_five_years AS (
        SELECT EXTRACT(
                YEAR
                FROM t.transaction_time
            ) AS year, t.payment_method, SUM(t.amount) AS total_amount
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
            ),
            t.payment_method
    )
SELECT year, payment_method, total_amount
FROM last_five_years
ORDER BY year`

type merchantStatsMethodByApiKeyRepository struct {
	db *sqlx.DB
}

func NewMerchantStatsMethodByApiKeyRepository(db *sqlx.DB) MerchantStatsMethodByApiKeyRepository {
	return &merchantStatsMethodByApiKeyRepository{
		db: db,
	}
}

func (r *merchantStatsMethodByApiKeyRepository) GetMonthlyPaymentMethodByApikey(ctx context.Context, req *requests.MonthYearPaymentMethodApiKey) ([]*db.GetMonthlyPaymentMethodByApikeyRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)
	var methods []*db.GetMonthlyPaymentMethodByApikeyRow
	err := r.db.SelectContext(ctx, &methods, getMonthlyPaymentMethodByApikey,
		yearStart,
		req.Apikey,
	)

	if err != nil {
		return nil, merchant_errors.ErrGetMonthlyPaymentMethodByApikeyFailed.WithInternal(err)
	}

	return methods, nil
}

func (r *merchantStatsMethodByApiKeyRepository) GetYearlyPaymentMethodByApikey(ctx context.Context, req *requests.MonthYearPaymentMethodApiKey) ([]*db.GetYearlyPaymentMethodByApikeyRow, error) {
	var methods []*db.GetYearlyPaymentMethodByApikeyRow
	err := r.db.SelectContext(ctx, &methods, getYearlyPaymentMethodByApikey,
		req.Apikey,
		req.Year,
	)

	if err != nil {
		return nil, merchant_errors.ErrGetYearlyPaymentMethodByApikeyFailed.WithInternal(err)
	}

	return methods, nil
}
