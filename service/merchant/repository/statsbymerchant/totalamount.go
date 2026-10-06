package merchantstatsmerchantrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyTotalAmountByMerchant = `WITH
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
            AND t.merchant_id = $2::integer
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
SELECT *
FROM formatted_data
ORDER BY year ASC, TO_DATE(month, 'Mon') ASC`

const getYearlyTotalAmountByMerchant = `WITH
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
            AND EXTRACT(
                YEAR
                FROM t.transaction_time
            ) >= $1::integer - 4
            AND EXTRACT(
                YEAR
                FROM t.transaction_time
            ) <= $1::integer
            AND t.merchant_id = $2::integer
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.transaction_time
            )
    ),
    formatted_data AS (
        SELECT year::text, total_amount
        FROM yearly_data
        UNION ALL
        SELECT y::text AS year, 0::integer AS total_amount
        FROM generate_series($1::integer - 4, $1::integer) AS y
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM yearly_data yd
                WHERE
                    yd.year = y
            )
    )
SELECT *
FROM formatted_data
ORDER BY year DESC`

type merchantStatsTotalAmountByMerchantRepository struct {
	db *sqlx.DB
}

func NewMerchantStatsTotalAmountByMerchantRepository(db *sqlx.DB) MerchantStatsTotalAmountByMerchantRepository {
	return &merchantStatsTotalAmountByMerchantRepository{
		db: db,
	}
}

func (r *merchantStatsTotalAmountByMerchantRepository) GetMonthlyTotalAmountByMerchants(ctx context.Context, req *requests.MonthYearTotalAmountMerchant) ([]*db.GetMonthlyTotalAmountByMerchantRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)
	var amounts []*db.GetMonthlyTotalAmountByMerchantRow
	err := r.db.SelectContext(ctx, &amounts, getMonthlyTotalAmountByMerchant,
		yearStart,
		int32(req.MerchantID),
	)

	if err != nil {
		return nil, merchant_errors.ErrGetMonthlyTotalAmountByMerchantsFailed.WithInternal(err)
	}

	return amounts, nil
}

func (r *merchantStatsTotalAmountByMerchantRepository) GetYearlyTotalAmountByMerchants(ctx context.Context, req *requests.MonthYearTotalAmountMerchant) ([]*db.GetYearlyTotalAmountByMerchantRow, error) {
	var amounts []*db.GetYearlyTotalAmountByMerchantRow
	err := r.db.SelectContext(ctx, &amounts, getYearlyTotalAmountByMerchant,
		int32(req.Year),
		int32(req.MerchantID),
	)

	if err != nil {
		return nil, merchant_errors.ErrGetYearlyTotalAmountByMerchantsFailed.WithInternal(err)
	}

	return amounts, nil
}
