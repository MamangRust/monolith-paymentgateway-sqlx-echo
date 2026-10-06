package transferstatsbycardrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	transfer_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transfer_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthTransferStatusSuccessCardNumber = `WITH
    monthly_data AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.transfer_time
            )::integer AS year,
            EXTRACT(
                MONTH
                FROM t.transfer_time
            )::integer AS month,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.transfer_amount), 0)::integer AS total_amount
        FROM transfers t
        WHERE
            t.deleted_at IS NULL
            AND t.status = 'success'
            AND (
                t.transfer_from = $1
                OR t.transfer_to = $1
            )
            AND (
                (
                    t.transfer_time >= $2::timestamp
                    AND t.transfer_time <= $3::timestamp
                )
                OR (
                    t.transfer_time >= $4::timestamp
                    AND t.transfer_time <= $5::timestamp
                )
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.transfer_time
            ),
            EXTRACT(
                MONTH
                FROM t.transfer_time
            )
    ),
    formatted_data AS (
        SELECT
            year::text,
            TO_CHAR(
                TO_DATE(month::text, 'MM'),
                'Mon'
            ) AS month,
            total_success,
            total_amount
        FROM monthly_data
        UNION ALL
        SELECT
            EXTRACT(
                YEAR
                FROM $2::timestamp
            )::text AS year,
            TO_CHAR($2::timestamp, 'Mon') AS month,
            0 AS total_success,
            0 AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM monthly_data
                WHERE
                    year = EXTRACT(
                        YEAR
                        FROM $2::timestamp
                    )::integer
                    AND month = EXTRACT(
                        MONTH
                        FROM $2::timestamp
                    )::integer
            )
        UNION ALL
        SELECT
            EXTRACT(
                YEAR
                FROM $4::timestamp
            )::text AS year,
            TO_CHAR($4::timestamp, 'Mon') AS month,
            0 AS total_success,
            0 AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM monthly_data
                WHERE
                    year = EXTRACT(
                        YEAR
                        FROM $4::timestamp
                    )::integer
                    AND month = EXTRACT(
                        MONTH
                        FROM $4::timestamp
                    )::integer
            )
    )
SELECT *
FROM formatted_data
ORDER BY year DESC, TO_DATE(month, 'Mon') DESC`

const getYearlyTransferStatusSuccessCardNumber = `WITH
    yearly_data AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.transfer_time
            )::integer AS year,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.transfer_amount), 0)::integer AS total_amount
        FROM transfers t
        WHERE
            t.deleted_at IS NULL
            AND t.status = 'success'
            AND (
                t.transfer_from = $1
                OR t.transfer_to = $1
            )
            AND (
                EXTRACT(
                    YEAR
                    FROM t.transfer_time
                ) = $2::integer
                OR EXTRACT(
                    YEAR
                    FROM t.transfer_time
                ) = $2::integer - 1
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.transfer_time
            )
    ),
    formatted_data AS (
        SELECT
            year::text,
            total_success::integer,
            total_amount::integer
        FROM yearly_data
        UNION ALL
        SELECT
            $2::text AS year,
            0::integer AS total_success,
            0::integer AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM yearly_data
                WHERE
                    year = $2::integer
            )
        UNION ALL
        SELECT ($2::integer - 1)::text AS year,
            0::integer AS total_success,
            0::integer AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM yearly_data
                WHERE
                    year = $2::integer - 1
            )
    )
SELECT *
FROM formatted_data
ORDER BY year DESC`

const getMonthTransferStatusFailedCardNumber = `WITH
    monthly_data AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.transfer_time
            )::integer AS year,
            EXTRACT(
                MONTH
                FROM t.transfer_time
            )::integer AS month,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.transfer_amount), 0)::integer AS total_amount
        FROM transfers t
        WHERE
            t.deleted_at IS NULL
            AND t.status = 'failed'
            AND (
                t.transfer_from = $1
                OR t.transfer_to = $1
            )
            AND (
                (
                    t.transfer_time >= $2::timestamp
                    AND t.transfer_time <= $3::timestamp
                )
                OR (
                    t.transfer_time >= $4::timestamp
                    AND t.transfer_time <= $5::timestamp
                )
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.transfer_time
            ),
            EXTRACT(
                MONTH
                FROM t.transfer_time
            )
    ),
    formatted_data AS (
        SELECT
            year::text,
            TO_CHAR(
                TO_DATE(month::text, 'MM'),
                'Mon'
            ) AS month,
            total_failed,
            total_amount
        FROM monthly_data
        UNION ALL
        SELECT
            EXTRACT(
                YEAR
                FROM $2::timestamp
            )::text AS year,
            TO_CHAR($2::timestamp, 'Mon') AS month,
            0 AS total_failed,
            0 AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM monthly_data
                WHERE
                    year = EXTRACT(
                        YEAR
                        FROM $2::timestamp
                    )::integer
                    AND month = EXTRACT(
                        MONTH
                        FROM $2::timestamp
                    )::integer
            )
        UNION ALL
        SELECT
            EXTRACT(
                YEAR
                FROM $4::timestamp
            )::text AS year,
            TO_CHAR($4::timestamp, 'Mon') AS month,
            0 AS total_failed,
            0 AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM monthly_data
                WHERE
                    year = EXTRACT(
                        YEAR
                        FROM $4::timestamp
                    )::integer
                    AND month = EXTRACT(
                        MONTH
                        FROM $4::timestamp
                    )::integer
            )
    )
SELECT *
FROM formatted_data
ORDER BY year DESC, TO_DATE(month, 'Mon') DESC`

const getYearlyTransferStatusFailedCardNumber = `WITH
    yearly_data AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.transfer_time
            )::integer AS year,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.transfer_amount), 0)::integer AS total_amount
        FROM transfers t
        WHERE
            t.deleted_at IS NULL
            AND t.status = 'failed'
            AND (
                t.transfer_from = $1
                OR t.transfer_to = $1
            )
            AND (
                EXTRACT(
                    YEAR
                    FROM t.transfer_time
                ) = $2::integer
                OR EXTRACT(
                    YEAR
                    FROM t.transfer_time
                ) = $2::integer - 1
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.transfer_time
            )
    ),
    formatted_data AS (
        SELECT
            year::text,
            total_failed::integer,
            total_amount::integer
        FROM yearly_data
        UNION ALL
        SELECT
            $2::text AS year,
            0::integer AS total_failed,
            0::integer AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM yearly_data
                WHERE
                    year = $2::integer
            )
        UNION ALL
        SELECT ($2::integer - 1)::text AS year,
            0::integer AS total_failed,
            0::integer AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM yearly_data
                WHERE
                    year = $2::integer - 1
            )
    )
SELECT *
FROM formatted_data
ORDER BY year DESC`

type transferStatsStatusByCardRepository struct {
	db *sqlx.DB
}

func NewTransferStatsByCardStatusRepository(db *sqlx.DB) TransferStatsByCardStatusRepository {
	return &transferStatsStatusByCardRepository{
		db: db,
	}
}

func (r *transferStatsStatusByCardRepository) GetMonthTransferStatusSuccessByCardNumber(ctx context.Context, req *requests.MonthStatusTransferCardNumber) ([]*db.GetMonthTransferStatusSuccessCardNumberRow, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	var rows []*db.GetMonthTransferStatusSuccessCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getMonthTransferStatusSuccessCardNumber,
		req.CardNumber,
		currentDate,
		lastDayCurrentMonth,
		prevDate,
		lastDayPrevMonth,
	)

	if err != nil {
		return nil, transfer_errors.ErrGetMonthTransferStatusSuccessByCardFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *transferStatsStatusByCardRepository) GetYearlyTransferStatusSuccessByCardNumber(ctx context.Context, req *requests.YearStatusTransferCardNumber) ([]*db.GetYearlyTransferStatusSuccessCardNumberRow, error) {
	var rows []*db.GetYearlyTransferStatusSuccessCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getYearlyTransferStatusSuccessCardNumber,
		req.CardNumber,
		int32(req.Year),
	)

	if err != nil {
		return nil, transfer_errors.ErrGetYearlyTransferStatusSuccessByCardFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *transferStatsStatusByCardRepository) GetMonthTransferStatusFailedByCardNumber(ctx context.Context, req *requests.MonthStatusTransferCardNumber) ([]*db.GetMonthTransferStatusFailedCardNumberRow, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	var rows []*db.GetMonthTransferStatusFailedCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getMonthTransferStatusFailedCardNumber,
		req.CardNumber,
		currentDate,
		lastDayCurrentMonth,
		prevDate,
		lastDayPrevMonth,
	)

	if err != nil {
		return nil, transfer_errors.ErrGetMonthTransferStatusFailedByCardFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *transferStatsStatusByCardRepository) GetYearlyTransferStatusFailedByCardNumber(ctx context.Context, req *requests.YearStatusTransferCardNumber) ([]*db.GetYearlyTransferStatusFailedCardNumberRow, error) {
	var rows []*db.GetYearlyTransferStatusFailedCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getYearlyTransferStatusFailedCardNumber,
		req.CardNumber,
		int32(req.Year),
	)

	if err != nil {
		return nil, transfer_errors.ErrGetYearlyTransferStatusFailedByCardFailed.WithInternal(err)
	}

	return rows, nil
}
