package merchantstatsrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyTotalAmountMerchant = `WITH
    monthly_data AS (
        SELECT EXTRACT(
                YEAR
                FROM t.transaction_time
            )::integer AS year, EXTRACT(
                MONTH
                FROM t.transaction_time
            )::integer AS month, COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
            INNER JOIN merchants m ON t.merchant_id = m.merchant_id
        WHERE
            t.deleted_at IS NULL
            AND m.deleted_at IS NULL
            AND EXTRACT(
                YEAR
                FROM t.transaction_time
            ) = EXTRACT(
                YEAR
                FROM $1::timestamp
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.transaction_time
            ),
            EXTRACT(
                MONTH
                FROM t.transaction_time
            )
    ),
    formatted_data AS (
        SELECT year::text, TO_CHAR(
                TO_DATE(month::text, 'MM'), 'Mon'
            ) AS month, total_amount
        FROM monthly_data
        UNION ALL
        SELECT EXTRACT(
                YEAR
                FROM gs.month
            )::text AS year, TO_CHAR(gs.month, 'Mon') AS month, 0::integer AS total_amount
        FROM generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '11 month', interval '1 month'
            ) AS gs (month)
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM monthly_data md
                WHERE
                    md.year = EXTRACT(
                        YEAR
                        FROM gs.month
                    )::integer
                    AND md.month = EXTRACT(
                        MONTH
                        FROM gs.month
                    )::integer
            )
    )
SELECT year, month, total_amount
FROM formatted_data
ORDER BY year ASC, TO_DATE(month, 'Mon') ASC`

const getYearlyTotalAmountMerchant = `WITH
    yearly_data AS (
        SELECT EXTRACT(
                YEAR
                FROM t.transaction_time
            )::integer AS year, COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
            INNER JOIN merchants m ON t.merchant_id = m.merchant_id
        WHERE
            t.deleted_at IS NULL
            AND m.deleted_at IS NULL
            AND (
                EXTRACT(
                    YEAR
                    FROM t.transaction_time
                ) = $1::integer
                OR EXTRACT(
                    YEAR
                    FROM t.transaction_time
                ) = $1::integer - 1
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.transaction_time
            )
    ),
    formatted_data AS (
        SELECT year::text, total_amount::integer
        FROM yearly_data
        UNION ALL
        SELECT $1::text AS year, 0::integer AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM yearly_data
                WHERE
                    year = $1::integer
            )
        UNION ALL
        SELECT ($1::integer - 1)::text AS year, 0::integer AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM yearly_data
                WHERE
                    year = $1::integer - 1
            )
    )
SELECT *
FROM formatted_data
ORDER BY year DESC`

type merchantStatsTotalAmountRepository struct {
	db *sqlx.DB
}

func NewMerchantStatsTotalAmountRepository(db *sqlx.DB) MerchantStatsTotalAmountRepository {
	return &merchantStatsTotalAmountRepository{
		db: db,
	}
}

func (r *merchantStatsTotalAmountRepository) GetMonthlyTotalAmountMerchant(ctx context.Context, year int) ([]*db.GetMonthlyTotalAmountMerchantRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	var amounts []*db.GetMonthlyTotalAmountMerchantRow
	err := r.db.SelectContext(ctx, &amounts, getMonthlyTotalAmountMerchant, yearStart)

	if err != nil {
		return nil, merchant_errors.ErrGetMonthlyTotalAmountMerchantFailed.WithInternal(err)
	}

	return amounts, nil
}

func (r *merchantStatsTotalAmountRepository) GetYearlyTotalAmountMerchant(ctx context.Context, year int) ([]*db.GetYearlyTotalAmountMerchantRow, error) {
	var amounts []*db.GetYearlyTotalAmountMerchantRow
	err := r.db.SelectContext(ctx, &amounts, getYearlyTotalAmountMerchant, int32(year))

	if err != nil {
		return nil, merchant_errors.ErrGetYearlyTotalAmountMerchantFailed.WithInternal(err)
	}

	return amounts, nil
}
