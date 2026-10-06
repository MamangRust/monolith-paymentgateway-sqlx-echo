package transferstatsbycardrepository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	transfer_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transfer_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyTransferAmountsBySenderCardNumber = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $2::timestamp), date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(t.transfer_amount), 0)::int AS total_transfer_amount
FROM
    months m
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
    AND t.transfer_from = $1
    AND t.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyTransferAmountsBySenderCardNumber = `SELECT EXTRACT(
        YEAR
        FROM t.created_at
    ) AS year, SUM(t.transfer_amount) AS total_transfer_amount
FROM transfers t
WHERE
    t.deleted_at IS NULL
    AND t.transfer_from = $1
    AND EXTRACT(
        YEAR
        FROM t.created_at
    ) >= $2 - 4
    AND EXTRACT(
        YEAR
        FROM t.created_at
    ) <= $2
GROUP BY
    EXTRACT(
        YEAR
        FROM t.created_at
    )
ORDER BY year`

type transferStatsAmountSenderRepository struct {
	db *sqlx.DB
}

func NewTransferStatsAmountSenderRepository(db *sqlx.DB) TransferStatsByCardAmountSenderRepository {
	return &transferStatsAmountSenderRepository{
		db: db,
	}
}

func (r *transferStatsAmountSenderRepository) GetMonthlyTransferAmountsBySenderCardNumber(ctx context.Context, req *requests.MonthYearCardNumber) ([]*db.GetMonthlyTransferAmountsBySenderCardNumberRow, error) {
	var rows []*db.GetMonthlyTransferAmountsBySenderCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getMonthlyTransferAmountsBySenderCardNumber,
		req.CardNumber,
		time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC),
	)

	if err != nil {
		return nil, transfer_errors.ErrGetMonthlyTransferAmountsBySenderCardFailed.WithInternal(err)
	}

	return rows, nil
}

func (r *transferStatsAmountSenderRepository) GetYearlyTransferAmountsBySenderCardNumber(ctx context.Context, req *requests.MonthYearCardNumber) ([]*db.GetYearlyTransferAmountsBySenderCardNumberRow, error) {
	var rows []*db.GetYearlyTransferAmountsBySenderCardNumberRow
	err := r.db.SelectContext(ctx, &rows, getYearlyTransferAmountsBySenderCardNumber,
		req.CardNumber,
		req.Year,
	)

	if err != nil {
		return nil, transfer_errors.ErrGetYearlyTransferAmountsBySenderCardFailed.WithInternal(err)
	}

	return rows, nil
}
