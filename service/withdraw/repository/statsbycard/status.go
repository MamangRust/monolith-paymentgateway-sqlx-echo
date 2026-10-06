package withdrawstatsbycardrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	withdraw_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/withdraw_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthWithdrawStatusSuccessCardNumber = `WITH
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
            AND t.card_number = $1
            AND (
                (
                    t.withdraw_time >= $2::timestamp
                    AND t.withdraw_time <= $3::timestamp
                )
                OR (
                    t.withdraw_time >= $4::timestamp
                    AND t.withdraw_time <= $5::timestamp
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

const getYearlyWithdrawStatusSuccessCardNumber = `WITH
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
            AND t.card_number = $1
            AND (
                EXTRACT(
                    YEAR
                    FROM t.withdraw_time
                ) = $2::integer
                OR EXTRACT(
                    YEAR
                    FROM t.withdraw_time
                ) = $2::integer - 1
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
SELECT year, total_success, total_amount
FROM formatted_data
ORDER BY year DESC`

const getMonthWithdrawStatusFailedCardNumber = `WITH
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
            AND t.card_number = $1
            AND (
                (
                    t.withdraw_time >= $2::timestamp
                    AND t.withdraw_time <= $3::timestamp
                )
                OR (
                    t.withdraw_time >= $4::timestamp
                    AND t.withdraw_time <= $5::timestamp
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

const getYearlyWithdrawStatusFailedCardNumber = `WITH
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
            AND t.card_number = $1
            AND (
                EXTRACT(
                    YEAR
                    FROM t.withdraw_time
                ) = $2::integer
                OR EXTRACT(
                    YEAR
                    FROM t.withdraw_time
                ) = $2::integer - 1
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
SELECT year, total_failed, total_amount
FROM formatted_data
ORDER BY year DESC`

type withdrawStatsByCardStatusRepository struct {
	db *sqlx.DB
}

func NewWithdrawStatsByCardStatusRepository(db *sqlx.DB) WithdrawStatsByCardStatusRepository {
	return &withdrawStatsByCardStatusRepository{
		db: db,
	}
}

func (r *withdrawStatsByCardStatusRepository) GetMonthWithdrawStatusSuccessByCardNumber(ctx context.Context, req *requests.MonthStatusWithdrawCardNumber) ([]*db.GetMonthWithdrawStatusSuccessCardNumberRow, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	var rows []*db.GetMonthWithdrawStatusSuccessCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getMonthWithdrawStatusSuccessCardNumber,
		req.CardNumber,
		currentDate,
		lastDayCurrentMonth,
		prevDate,
		lastDayPrevMonth,
	)

	if err != nil {
		return nil, withdraw_errors.ErrGetMonthWithdrawStatusSuccessByCardFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *withdrawStatsByCardStatusRepository) GetYearlyWithdrawStatusSuccessByCardNumber(ctx context.Context, req *requests.YearStatusWithdrawCardNumber) ([]*db.GetYearlyWithdrawStatusSuccessCardNumberRow, error) {
	var rows []*db.GetYearlyWithdrawStatusSuccessCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getYearlyWithdrawStatusSuccessCardNumber,
		req.CardNumber,
		int32(req.Year),
	)

	if err != nil {
		return nil, withdraw_errors.ErrGetYearlyWithdrawStatusSuccessByCardFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *withdrawStatsByCardStatusRepository) GetMonthWithdrawStatusFailedByCardNumber(ctx context.Context, req *requests.MonthStatusWithdrawCardNumber) ([]*db.GetMonthWithdrawStatusFailedCardNumberRow, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	var rows []*db.GetMonthWithdrawStatusFailedCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getMonthWithdrawStatusFailedCardNumber,
		req.CardNumber,
		currentDate,
		lastDayCurrentMonth,
		prevDate,
		lastDayPrevMonth,
	)

	if err != nil {
		return nil, withdraw_errors.ErrGetMonthWithdrawStatusFailedByCardFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *withdrawStatsByCardStatusRepository) GetYearlyWithdrawStatusFailedByCardNumber(ctx context.Context, req *requests.YearStatusWithdrawCardNumber) ([]*db.GetYearlyWithdrawStatusFailedCardNumberRow, error) {
	var rows []*db.GetYearlyWithdrawStatusFailedCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getYearlyWithdrawStatusFailedCardNumber,
		req.CardNumber,
		int32(req.Year),
	)

	if err != nil {
		return nil, withdraw_errors.ErrGetYearlyWithdrawStatusFailedByCardFailed.WithInternal(err)
	}

	return rows, nil
}
