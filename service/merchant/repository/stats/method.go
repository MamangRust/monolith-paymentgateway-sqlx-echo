package merchantstatsrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyPaymentMethodsMerchant = `WITH
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
    LEFT JOIN merchants mch ON t.merchant_id = mch.merchant_id
    AND mch.deleted_at IS NULL
GROUP BY
    m.month,
    pm.payment_method
ORDER BY m.month, pm.payment_method`

const getYearlyPaymentMethodMerchant = `WITH
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
            ),
            t.payment_method
    )
SELECT year, payment_method, total_amount
FROM last_five_years
ORDER BY year`

type merchantStatsMethodRepository struct {
	db *sqlx.DB
}

func NewMerchantStatsMethodRepository(db *sqlx.DB) MerchantStatsMethodRepository {
	return &merchantStatsMethodRepository{
		db: db,
	}
}

func (r *merchantStatsMethodRepository) GetMonthlyPaymentMethodsMerchant(ctx context.Context, year int) ([]*db.GetMonthlyPaymentMethodsMerchantRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var methods []*db.GetMonthlyPaymentMethodsMerchantRow
	err := r.db.SelectContext(ctx, &methods, getMonthlyPaymentMethodsMerchant, yearStart)

	if err != nil {
		return nil, merchant_errors.ErrGetMonthlyPaymentMethodsMerchantFailed.WithInternal(err)
	}

	return methods, nil
}

func (r *merchantStatsMethodRepository) GetYearlyPaymentMethodMerchant(ctx context.Context, year int) ([]*db.GetYearlyPaymentMethodMerchantRow, error) {
	var methods []*db.GetYearlyPaymentMethodMerchantRow
	err := r.db.SelectContext(ctx, &methods, getYearlyPaymentMethodMerchant, year)

	if err != nil {
		return nil, merchant_errors.ErrGetYearlyPaymentMethodMerchantFailed.WithInternal(err)
	}

	return methods, nil
}
