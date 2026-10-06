package saldostatsrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	saldo_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/saldo_errors/repository"
	"github.com/jmoiron/sqlx"
)

type saldoStatsTotalBalanceRepository struct {
	db *sqlx.DB
}

func NewSaldoStatsTotalBalanceRepository(db *sqlx.DB) SaldoStatsTotalSaldoRepository {
	return &saldoStatsTotalBalanceRepository{
		db: db,
	}
}

const getMonthlyTotalSaldoBalance = `WITH
    monthly_data AS (
        SELECT EXTRACT(
                YEAR
                FROM s.created_at
            )::integer AS year, EXTRACT(
                MONTH
                FROM s.created_at
            )::integer AS month, COALESCE(SUM(s.total_balance), 0) AS total_balance
        FROM saldos s
        WHERE
            s.deleted_at IS NULL
            AND (
                (
                    s.created_at >= $1::timestamp
                    AND s.created_at <= $2::timestamp
                )
                OR (
                    s.created_at >= $3::timestamp
                    AND s.created_at <= $4::timestamp
                )
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM s.created_at
            ),
            EXTRACT(
                MONTH
                FROM s.created_at
            )
    ),
    formatted_data AS (
        SELECT year::text, TO_CHAR(
                TO_DATE(month::text, 'MM'), 'Mon'
            ) AS month, total_balance::integer
        FROM monthly_data
        UNION ALL
        SELECT EXTRACT(
                YEAR
                FROM $1::timestamp
            )::text AS year, TO_CHAR($1::timestamp, 'Mon') AS month, 0::integer AS total_balance
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
        SELECT EXTRACT(
                YEAR
                FROM $3::timestamp
            )::text AS year, TO_CHAR($3::timestamp, 'Mon') AS month, 0::integer AS total_balance
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
SELECT year, month, total_balance
FROM formatted_data
ORDER BY year DESC, TO_DATE(month, 'Mon') DESC`

const getYearlyTotalSaldoBalances = `WITH
    yearly_data AS (
        SELECT EXTRACT(
                YEAR
                FROM s.created_at
            )::integer AS year, COALESCE(SUM(s.total_balance), 0)::integer AS total_balance
        FROM saldos s
        WHERE
            s.deleted_at IS NULL
            AND (
                EXTRACT(
                    YEAR
                    FROM s.created_at
                ) = $1::integer
                OR EXTRACT(
                    YEAR
                    FROM s.created_at
                ) = $1::integer - 1
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM s.created_at
            )
    ),
    formatted_data AS (
        SELECT year::text, total_balance::integer
        FROM yearly_data
        UNION ALL
        SELECT $1::text AS year, 0::integer AS total_balance
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM yearly_data
                WHERE
                    year = $1::integer
            )
        UNION ALL
        SELECT ($1::integer - 1)::text AS year, 0::integer AS total_balance
        WHERE
            NOT EXISTS (
                SELECT 1
                FROM yearly_data
                WHERE
                    year = $1::integer - 1
            )
    )
SELECT year, total_balance
FROM formatted_data
ORDER BY year DESC`

func (r *saldoStatsTotalBalanceRepository) GetMonthlyTotalSaldoBalance(ctx context.Context, req *requests.MonthTotalSaldoBalance) ([]*db.GetMonthlyTotalSaldoBalanceRow, error) {
	year := req.Year
	month := req.Month

	currentDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	var rows []*db.GetMonthlyTotalSaldoBalanceRow
	err := r.db.SelectContext(ctx, &rows, getMonthlyTotalSaldoBalance,
		currentDate,
		lastDayCurrentMonth,
		prevDate,
		lastDayPrevMonth,
	)

	if err != nil {
		return nil, saldo_errors.ErrGetMonthlyTotalSaldoBalanceFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *saldoStatsTotalBalanceRepository) GetYearTotalSaldoBalance(ctx context.Context, year int) ([]*db.GetYearlyTotalSaldoBalancesRow, error) {
	var rows []*db.GetYearlyTotalSaldoBalancesRow
	err := r.db.SelectContext(ctx, &rows, getYearlyTotalSaldoBalances, int32(year))

	if err != nil {
		return nil, saldo_errors.ErrGetYearTotalSaldoBalanceFailed.WithInternal(err)
	}

	return rows, nil
}
