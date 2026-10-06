package transferstatsrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	transfer_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transfer_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyTransferAmounts = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(t.transfer_amount), 0)::int AS total_transfer_amount
FROM months m
    LEFT JOIN transfers t ON EXTRACT(
        MONTH
        FROM t.transfer_time
    ) = EXTRACT(
        MONTH
        FROM m.month
    )
    AND EXTRACT(
        YEAR
        FROM t.transfer_time
    ) = EXTRACT(
        YEAR
        FROM m.month
    )
    AND t.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyTransferAmounts = `SELECT EXTRACT(
        YEAR
        FROM t.created_at
    ) AS year, SUM(t.transfer_amount) AS total_transfer_amount
FROM transfers t
WHERE
    t.deleted_at IS NULL
    AND EXTRACT(
        YEAR
        FROM t.created_at
    ) >= $1 - 4
    AND EXTRACT(
        YEAR
        FROM t.created_at
    ) <= $1
GROUP BY
    EXTRACT(
        YEAR
        FROM t.created_at
    )
ORDER BY year`

type transferStatsAmountRepository struct {
	db *sqlx.DB
}

func NewTransferStatsAmountRepository(db *sqlx.DB) TransferStatsAmountRepository {
	return &transferStatsAmountRepository{
		db: db,
	}
}

func (r *transferStatsAmountRepository) GetMonthlyTransferAmounts(ctx context.Context, year int) ([]*db.GetMonthlyTransferAmountsRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var rows []*db.GetMonthlyTransferAmountsRow
	err := r.db.SelectContext(ctx, &rows, getMonthlyTransferAmounts, yearStart)

	if err != nil {
		return nil, transfer_errors.ErrGetMonthlyTransferAmountsFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *transferStatsAmountRepository) GetYearlyTransferAmounts(ctx context.Context, year int) ([]*db.GetYearlyTransferAmountsRow, error) {
	var rows []*db.GetYearlyTransferAmountsRow
	err := r.db.SelectContext(ctx, &rows, getYearlyTransferAmounts, year)

	if err != nil {
		return nil, transfer_errors.ErrGetYearlyTransferAmountsFailed.WithInternal(err)
	}
	return rows, nil
}
