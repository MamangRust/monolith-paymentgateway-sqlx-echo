package topupstatsbycardrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthTopupStatusSuccessCardNumber = `WITH
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
            AND t.card_number = $1
            AND (
                (
                    t.topup_time >= $2::timestamp
                    AND t.topup_time <= $3::timestamp
                )
                OR (
                    t.topup_time >= $4::timestamp
                    AND t.topup_time <= $5::timestamp
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

const getYearlyTopupStatusSuccessCardNumber = `WITH
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
            AND t.card_number = $1
            AND (
                EXTRACT(
                    YEAR
                    FROM t.topup_time
                ) = $2::integer
                OR EXTRACT(
                    YEAR
                    FROM t.topup_time
                ) = $2::integer - 1
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

const getMonthTopupStatusFailedCardNumber = `WITH
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
            AND t.card_number = $1
            AND (
                (
                    t.topup_time >= $2::timestamp
                    AND t.topup_time <= $3::timestamp
                )
                OR (
                    t.topup_time >= $4::timestamp
                    AND t.topup_time <= $5::timestamp
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
SELECT year, month, total_failed, total_amount
FROM formatted_data
ORDER BY year DESC, TO_DATE(month, 'Mon') DESC`

const getYearlyTopupStatusFailedCardNumber = `WITH
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
            AND t.card_number = $1
            AND (
                EXTRACT(
                    YEAR
                    FROM t.topup_time
                ) = $2::integer
                OR EXTRACT(
                    YEAR
                    FROM t.topup_time
                ) = $2::integer - 1
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

type topupStatsByCardStatusRepository struct {
	db *sqlx.DB
}

func NewTopupStatsByCardStatusRepository(db *sqlx.DB) TopupStatsByCardStatusRepository {
	return &topupStatsByCardStatusRepository{
		db: db,
	}
}

func (r *topupStatsByCardStatusRepository) GetMonthTopupStatusSuccessByCardNumber(ctx context.Context, req *requests.MonthTopupStatusCardNumber) ([]*db.GetMonthTopupStatusSuccessCardNumberRow, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	var rows []*db.GetMonthTopupStatusSuccessCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getMonthTopupStatusSuccessCardNumber,
		req.CardNumber,
		currentDate,
		lastDayCurrentMonth,
		prevDate,
		lastDayPrevMonth,
	)

	if err != nil {
		return nil, topup_errors.ErrGetMonthTopupStatusSuccessByCardFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *topupStatsByCardStatusRepository) GetYearlyTopupStatusSuccessByCardNumber(ctx context.Context, req *requests.YearTopupStatusCardNumber) ([]*db.GetYearlyTopupStatusSuccessCardNumberRow, error) {
	var rows []*db.GetYearlyTopupStatusSuccessCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getYearlyTopupStatusSuccessCardNumber,
		req.CardNumber,
		int32(req.Year),
	)

	if err != nil {
		return nil, topup_errors.ErrGetYearlyTopupStatusSuccessByCardFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *topupStatsByCardStatusRepository) GetMonthTopupStatusFailedByCardNumber(ctx context.Context, req *requests.MonthTopupStatusCardNumber) ([]*db.GetMonthTopupStatusFailedCardNumberRow, error) {
	cardNumber := req.CardNumber
	year := req.Year
	month := req.Month

	currentDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	var rows []*db.GetMonthTopupStatusFailedCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getMonthTopupStatusFailedCardNumber,
		cardNumber,
		currentDate,
		lastDayCurrentMonth,
		prevDate,
		lastDayPrevMonth,
	)

	if err != nil {
		return nil, topup_errors.ErrGetMonthTopupStatusFailedByCardFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *topupStatsByCardStatusRepository) GetYearlyTopupStatusFailedByCardNumber(ctx context.Context, req *requests.YearTopupStatusCardNumber) ([]*db.GetYearlyTopupStatusFailedCardNumberRow, error) {
	var rows []*db.GetYearlyTopupStatusFailedCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getYearlyTopupStatusFailedCardNumber,
		req.CardNumber,
		int32(req.Year),
	)

	if err != nil {
		return nil, topup_errors.ErrGetYearlyTopupStatusSuccessByCardFailed.WithInternal(err)
	}

	return rows, nil
}
