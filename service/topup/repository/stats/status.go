package topupstatsrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthTopupStatusSuccess = `WITH
    monthly_data AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.topup_time
            )::integer AS year,
            EXTRACT(
                MONTH
                FROM t.topup_time
            )::integer AS month,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.topup_amount), 0)::integer AS total_amount
        FROM topups t
        WHERE
            t.deleted_at IS NULL
            AND t.status = 'success'
            AND (
                (
                    t.topup_time >= $1::timestamp
                    AND t.topup_time <= $2::timestamp
                )
                OR (
                    t.topup_time >= $3::timestamp
                    AND t.topup_time <= $4::timestamp
                )
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.topup_time
            ),
            EXTRACT(
                MONTH
                FROM t.topup_time
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

const getYearlyTopupStatusSuccess = `WITH
    yearly_data AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.topup_time
            )::integer AS year,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.topup_amount), 0)::integer AS total_amount
        FROM topups t
        WHERE
            t.deleted_at IS NULL
            AND t.status = 'success'
            AND (
                EXTRACT(
                    YEAR
                    FROM t.topup_time
                ) = $1::integer
                OR EXTRACT(
                    YEAR
                    FROM t.topup_time
                ) = $1::integer - 1
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.topup_time
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

const getMonthTopupStatusFailed = `WITH
    monthly_data AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.topup_time
            )::integer AS year,
            EXTRACT(
                MONTH
                FROM t.topup_time
            )::integer AS month,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.topup_amount), 0)::integer AS total_amount
        FROM topups t
        WHERE
            t.deleted_at IS NULL
            AND t.status = 'failed'
            AND (
                (
                    t.topup_time >= $1::timestamp
                    AND t.topup_time <= $2::timestamp
                )
                OR (
                    t.topup_time >= $3::timestamp
                    AND t.topup_time <= $4::timestamp
                )
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.topup_time
            ),
            EXTRACT(
                MONTH
                FROM t.topup_time
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

const getYearlyTopupStatusFailed = `WITH
    yearly_data AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.topup_time
            )::integer AS year,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.topup_amount), 0)::integer AS total_amount
        FROM topups t
        WHERE
            t.deleted_at IS NULL
            AND t.status = 'failed'
            AND (
                EXTRACT(
                    YEAR
                    FROM t.topup_time
                ) = $1::integer
                OR EXTRACT(
                    YEAR
                    FROM t.topup_time
                ) = $1::integer - 1
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.topup_time
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

type topupStatsStatusRepository struct {
	db *sqlx.DB
}

func NewTopupStatsStatusRepository(db *sqlx.DB) TopupStatsStatusRepository {
	return &topupStatsStatusRepository{
		db: db,
	}
}

func (r *topupStatsStatusRepository) GetMonthTopupStatusSuccess(ctx context.Context, req *requests.MonthTopupStatus) ([]*db.GetMonthTopupStatusSuccessRow, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	var rows []*db.GetMonthTopupStatusSuccessRow
	err := r.db.SelectContext(ctx, &rows, getMonthTopupStatusSuccess,
		currentDate,
		lastDayCurrentMonth,
		prevDate,
		lastDayPrevMonth,
	)

	if err != nil {
		return nil, topup_errors.ErrGetMonthTopupStatusSuccessFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *topupStatsStatusRepository) GetYearlyTopupStatusSuccess(ctx context.Context, year int) ([]*db.GetYearlyTopupStatusSuccessRow, error) {
	var rows []*db.GetYearlyTopupStatusSuccessRow
	err := r.db.SelectContext(ctx, &rows, getYearlyTopupStatusSuccess, int32(year))

	if err != nil {
		return nil, topup_errors.ErrGetYearlyTopupStatusSuccessFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *topupStatsStatusRepository) GetMonthTopupStatusFailed(ctx context.Context, req *requests.MonthTopupStatus) ([]*db.GetMonthTopupStatusFailedRow, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	var rows []*db.GetMonthTopupStatusFailedRow
	err := r.db.SelectContext(ctx, &rows, getMonthTopupStatusFailed,
		currentDate,
		lastDayCurrentMonth,
		prevDate,
		lastDayPrevMonth,
	)

	if err != nil {
		return nil, topup_errors.ErrGetMonthTopupStatusFailedFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *topupStatsStatusRepository) GetYearlyTopupStatusFailed(ctx context.Context, year int) ([]*db.GetYearlyTopupStatusFailedRow, error) {
	var rows []*db.GetYearlyTopupStatusFailedRow
	err := r.db.SelectContext(ctx, &rows, getYearlyTopupStatusFailed, int32(year))

	if err != nil {
		return nil, topup_errors.ErrGetYearlyTopupStatusFailedFailed.WithInternal(err)
	}

	return rows, nil
}
