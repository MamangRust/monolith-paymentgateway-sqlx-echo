package withdrawstatsrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	withdraw_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/withdraw_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthWithdrawStatusSuccess = `WITH
    monthly_data AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.withdraw_time
            )::integer AS year,
            EXTRACT(
                MONTH
                FROM t.withdraw_time
            )::integer AS month,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.withdraw_amount), 0)::integer AS total_amount
        FROM withdraws t
        WHERE
            t.deleted_at IS NULL
            AND t.status = 'success'
            AND (
                (
                    t.withdraw_time >= $1::timestamp
                    AND t.withdraw_time <= $2::timestamp
                )
                OR (
                    t.withdraw_time >= $3::timestamp
                    AND t.withdraw_time <= $4::timestamp
                )
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.withdraw_time
            ),
            EXTRACT(
                MONTH
                FROM t.withdraw_time
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
                FROM $1::timestamp
            )::text AS year,
            TO_CHAR($1::timestamp, 'Mon') AS month,
            0 AS total_success,
            0 AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM monthly_data
                WHERE
                    year = EXTRACT(
                        YEAR
                        FROM $1::timestamp
                    )::integer
                    AND month = EXTRACT(
                        MONTH
                        FROM $1::timestamp
                    )::integer
            )
        UNION ALL
        SELECT
            EXTRACT(
                YEAR
                FROM $3::timestamp
            )::text AS year,
            TO_CHAR($3::timestamp, 'Mon') AS month,
            0 AS total_success,
            0 AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM monthly_data
                WHERE
                    year = EXTRACT(
                        YEAR
                        FROM $3::timestamp
                    )::integer
                    AND month = EXTRACT(
                        MONTH
                        FROM $3::timestamp
                    )::integer
            )
    )
SELECT year, month, total_success, total_amount
FROM formatted_data
ORDER BY year DESC, TO_DATE(month, 'Mon') DESC`

const getYearlyWithdrawStatusSuccess = `WITH
    yearly_data AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.withdraw_time
            )::integer AS year,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.withdraw_amount), 0)::integer AS total_amount
        FROM withdraws t
        WHERE
            t.deleted_at IS NULL
            AND t.status = 'success'
            AND (
                EXTRACT(
                    YEAR
                    FROM t.withdraw_time
                ) = $1::integer
                OR EXTRACT(
                    YEAR
                    FROM t.withdraw_time
                ) = $1::integer - 1
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.withdraw_time
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
            $1::text AS year,
            0::integer AS total_success,
            0::integer AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM yearly_data
                WHERE
                    year = $1::integer
            )
        UNION ALL
        SELECT ($1::integer - 1)::text AS year,
            0::integer AS total_success,
            0::integer AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM yearly_data
                WHERE
                    year = $1::integer - 1
            )
    )
SELECT year, total_success, total_amount
FROM formatted_data
ORDER BY year DESC`

const getMonthWithdrawStatusFailed = `WITH
    monthly_data AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.withdraw_time
            )::integer AS year,
            EXTRACT(
                MONTH
                FROM t.withdraw_time
            )::integer AS month,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.withdraw_amount), 0)::integer AS total_amount
        FROM withdraws t
        WHERE
            t.deleted_at IS NULL
            AND t.status = 'failed'
            AND (
                (
                    t.withdraw_time >= $1::timestamp
                    AND t.withdraw_time <= $2::timestamp
                )
                OR (
                    t.withdraw_time >= $3::timestamp
                    AND t.withdraw_time <= $4::timestamp
                )
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.withdraw_time
            ),
            EXTRACT(
                MONTH
                FROM t.withdraw_time
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
                FROM $1::timestamp
            )::text AS year,
            TO_CHAR($1::timestamp, 'Mon') AS month,
            0 AS total_failed,
            0 AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM monthly_data
                WHERE
                    year = EXTRACT(
                        YEAR
                        FROM $1::timestamp
                    )::integer
                    AND month = EXTRACT(
                        MONTH
                        FROM $1::timestamp
                    )::integer
            )
        UNION ALL
        SELECT
            EXTRACT(
                YEAR
                FROM $3::timestamp
            )::text AS year,
            TO_CHAR($3::timestamp, 'Mon') AS month,
            0 AS total_failed,
            0 AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM monthly_data
                WHERE
                    year = EXTRACT(
                        YEAR
                        FROM $3::timestamp
                    )::integer
                    AND month = EXTRACT(
                        MONTH
                        FROM $3::timestamp
                    )::integer
            )
    )
SELECT year, month, total_failed, total_amount
FROM formatted_data
ORDER BY year DESC, TO_DATE(month, 'Mon') DESC`

const getYearlyWithdrawStatusFailed = `WITH
    yearly_data AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.withdraw_time
            )::integer AS year,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.withdraw_amount), 0)::integer AS total_amount
        FROM withdraws t
        WHERE
            t.deleted_at IS NULL
            AND t.status = 'failed'
            AND (
                EXTRACT(
                    YEAR
                    FROM t.withdraw_time
                ) = $1::integer
                OR EXTRACT(
                    YEAR
                    FROM t.withdraw_time
                ) = $1::integer - 1
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.withdraw_time
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
            $1::text AS year,
            0::integer AS total_failed,
            0::integer AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM yearly_data
                WHERE
                    year = $1::integer
            )
        UNION ALL
        SELECT ($1::integer - 1)::text AS year,
            0::integer AS total_failed,
            0::integer AS total_amount
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM yearly_data
                WHERE
                    year = $1::integer - 1
            )
    )
SELECT year, total_failed, total_amount
FROM formatted_data
ORDER BY year DESC`

type withdrawStatsStatusRepository struct {
	db *sqlx.DB
}

func NewWithdrawStatsStatusRepository(db *sqlx.DB) WithdrawStatsStatusRepository {
	return &withdrawStatsStatusRepository{
		db: db,
	}
}

func (r *withdrawStatsStatusRepository) GetMonthWithdrawStatusSuccess(ctx context.Context, req *requests.MonthStatusWithdraw) ([]*db.GetMonthWithdrawStatusSuccessRow, error) {
	year := req.Year
	month := req.Month

	currentDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	var rows []*db.GetMonthWithdrawStatusSuccessRow
	err := r.db.SelectContext(ctx, &rows, getMonthWithdrawStatusSuccess,
		currentDate,
		lastDayCurrentMonth,
		prevDate,
		lastDayPrevMonth,
	)

	if err != nil {
		return nil, withdraw_errors.ErrGetMonthWithdrawStatusSuccessFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *withdrawStatsStatusRepository) GetYearlyWithdrawStatusSuccess(ctx context.Context, year int) ([]*db.GetYearlyWithdrawStatusSuccessRow, error) {
	var rows []*db.GetYearlyWithdrawStatusSuccessRow
	err := r.db.SelectContext(ctx, &rows, getYearlyWithdrawStatusSuccess, int32(year))

	if err != nil {
		return nil, withdraw_errors.ErrGetYearlyWithdrawStatusSuccessFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *withdrawStatsStatusRepository) GetMonthWithdrawStatusFailed(ctx context.Context, req *requests.MonthStatusWithdraw) ([]*db.GetMonthWithdrawStatusFailedRow, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	var rows []*db.GetMonthWithdrawStatusFailedRow
	err := r.db.SelectContext(ctx, &rows, getMonthWithdrawStatusFailed,
		currentDate,
		lastDayCurrentMonth,
		prevDate,
		lastDayPrevMonth,
	)

	if err != nil {
		return nil, withdraw_errors.ErrGetMonthWithdrawStatusFailedFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *withdrawStatsStatusRepository) GetYearlyWithdrawStatusFailed(ctx context.Context, year int) ([]*db.GetYearlyWithdrawStatusFailedRow, error) {
	var rows []*db.GetYearlyWithdrawStatusFailedRow
	err := r.db.SelectContext(ctx, &rows, getYearlyWithdrawStatusFailed, int32(year))

	if err != nil {
		return nil, withdraw_errors.ErrGetYearlyWithdrawStatusFailedFailed.WithInternal(err)
	}

	return rows, nil
}
