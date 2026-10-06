package repositorystatsbycard

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyTransferAmountBySender = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $2::timestamp), date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(t.transfer_amount), 0)::int AS total_sent_amount
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
    AND t.deleted_at IS NULL
    AND t.transfer_from = $1
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyTransferAmountBySender = `WITH
    last_five_years AS (
        SELECT EXTRACT(
                YEAR
                FROM t.transfer_time
            ) AS year, SUM(t.transfer_amount) AS total_sent_amount
        FROM transfers t
        WHERE
            t.deleted_at IS NULL
            AND t.transfer_from = $1
            AND EXTRACT(
                YEAR
                FROM t.transfer_time
            ) >= $2 - 4
            AND EXTRACT(
                YEAR
                FROM t.transfer_time
            ) <= $2
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.transfer_time
            )
    )
SELECT year, total_sent_amount
FROM last_five_years
ORDER BY year`

const getMonthlyTransferAmountByReceiver = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $2::timestamp), date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(t.transfer_amount), 0)::int AS total_received_amount
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
    AND t.deleted_at IS NULL
    AND t.transfer_to = $1
GROUP BY
    m.month
ORDER BY m.month`

const getYearlyTransferAmountByReceiver = `WITH
    last_five_years AS (
        SELECT EXTRACT(
                YEAR
                FROM t.transfer_time
            ) AS year, SUM(t.transfer_amount) AS total_received_amount
        FROM transfers t
        WHERE
            t.deleted_at IS NULL
            AND t.transfer_to = $1
            AND EXTRACT(
                YEAR
                FROM t.transfer_time
            ) >= $2 - 4
            AND EXTRACT(
                YEAR
                FROM t.transfer_time
            ) <= $2
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.transfer_time
            )
    )
SELECT year, total_received_amount
FROM last_five_years
ORDER BY year`

type cardStatsTransferByCardRepository struct {
	db *sqlx.DB
}

func NewCardStatsTransferByCardRepository(db *sqlx.DB) CardStatsTransferByCardRepository {
	return &cardStatsTransferByCardRepository{
		db: db,
	}
}

func (r *cardStatsTransferByCardRepository) GetMonthlyTransferAmountBySender(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetMonthlyTransferAmountBySenderRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	var res []*db.GetMonthlyTransferAmountBySenderRow
	err := r.db.SelectContext(ctx, &res, getMonthlyTransferAmountBySender,
		req.CardNumber,
		yearStart,
	)

	if err != nil {
		return nil, card_errors.ErrGetMonthlyTransferAmountBySenderFailed
	}

	return res, nil
}

func (r *cardStatsTransferByCardRepository) GetYearlyTransferAmountBySender(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetYearlyTransferAmountBySenderRow, error) {
	var res []*db.GetYearlyTransferAmountBySenderRow
	err := r.db.SelectContext(ctx, &res, getYearlyTransferAmountBySender,
		req.CardNumber,
		int32(req.Year),
	)

	if err != nil {
		return nil, card_errors.ErrGetYearlyTransferAmountBySenderFailed
	}

	return res, nil
}

func (r *cardStatsTransferByCardRepository) GetMonthlyTransferAmountByReceiver(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetMonthlyTransferAmountByReceiverRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	var res []*db.GetMonthlyTransferAmountByReceiverRow
	err := r.db.SelectContext(ctx, &res, getMonthlyTransferAmountByReceiver,
		req.CardNumber,
		yearStart,
	)

	if err != nil {
		return nil, card_errors.ErrGetMonthlyTransferAmountByReceiverFailed.WithInternal(err)
	}

	return res, nil
}

func (r *cardStatsTransferByCardRepository) GetYearlyTransferAmountByReceiver(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetYearlyTransferAmountByReceiverRow, error) {
	var res []*db.GetYearlyTransferAmountByReceiverRow
	err := r.db.SelectContext(ctx, &res, getYearlyTransferAmountByReceiver,
		req.CardNumber,
		int32(req.Year),
	)

	if err != nil {
		return nil, card_errors.ErrGetYearlyTransferAmountByReceiverFailed.WithInternal(err)
	}

	return res, nil
}
