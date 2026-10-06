package saldostatsrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	saldo_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/saldo_errors/repository"
	"github.com/jmoiron/sqlx"
)

type saldoStatsBalanceRepository struct {
	db *sqlx.DB
}

func NewSaldoStatsBalanceRepository(db *sqlx.DB) SaldoStatsBalanceRepository {
	return &saldoStatsBalanceRepository{
		db: db,
	}
}

const getMonthlySaldoBalances = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(s.total_balance), 0)::int AS total_balance
FROM months m
    LEFT JOIN saldos s ON EXTRACT(
        MONTH
        FROM s.created_at
    ) = EXTRACT(
        MONTH
        FROM m.month
    )
    AND EXTRACT(
        YEAR
        FROM s.created_at
    ) = EXTRACT(
        YEAR
        FROM m.month
    )
    AND s.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlySaldoBalances = `WITH
    years AS (
        SELECT generate_series($1::int - 4, $1::int) AS year
    ),
    yearly_data AS (
        SELECT EXTRACT(
                YEAR
                FROM s.created_at
            )::int AS year, SUM(s.total_balance) AS total_balance
        FROM saldos s
        WHERE
            s.deleted_at IS NULL
            AND EXTRACT(
                YEAR
                FROM s.created_at
            ) >= $1::int - 4
            AND EXTRACT(
                YEAR
                FROM s.created_at
            ) <= $1::int
        GROUP BY
            EXTRACT(
                YEAR
                FROM s.created_at
            )
    )
SELECT y.year::text, COALESCE(yd.total_balance, 0)::bigint AS total_balance
FROM years y
    LEFT JOIN yearly_data yd ON y.year = yd.year
ORDER BY y.year`

func (r *saldoStatsBalanceRepository) GetMonthlySaldoBalances(ctx context.Context, year int) ([]*db.GetMonthlySaldoBalancesRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var rows []*db.GetMonthlySaldoBalancesRow
	err := r.db.SelectContext(ctx, &rows, getMonthlySaldoBalances, yearStart)

	if err != nil {
		return nil, saldo_errors.ErrGetMonthlySaldoBalancesFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *saldoStatsBalanceRepository) GetYearlySaldoBalances(ctx context.Context, year int) ([]*db.GetYearlySaldoBalancesRow, error) {
	var rows []*db.GetYearlySaldoBalancesRow
	err := r.db.SelectContext(ctx, &rows, getYearlySaldoBalances, int32(year))

	if err != nil {
		return nil, saldo_errors.ErrGetYearlySaldoBalancesFailed.WithInternal(err)
	}

	return rows, nil
}
