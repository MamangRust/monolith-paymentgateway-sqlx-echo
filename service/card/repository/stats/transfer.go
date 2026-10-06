package repositorystats

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMonthlyTransferAmountSender = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(t.transfer_amount), 0)::int AS total_sent_amount
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

const getYearlyTransferAmountSender = `WITH
    last_five_years AS (
        SELECT EXTRACT(
                YEAR
                FROM t.transfer_time
            ) AS year, SUM(t.transfer_amount) AS total_sent_amount
        FROM transfers t
        WHERE
            t.deleted_at IS NULL
            AND EXTRACT(
                YEAR
                FROM t.transfer_time
            ) >= $1 - 4
            AND EXTRACT(
                YEAR
                FROM t.transfer_time
            ) <= $1
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.transfer_time
            )
    )
SELECT year, total_sent_amount
FROM last_five_years
ORDER BY year`

const getMonthlyTransferAmountReceiver = `WITH
    months AS (
        SELECT generate_series(
                date_trunc('year', $1::timestamp), date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day', interval '1 month'
            ) AS month
    )
SELECT TO_CHAR(m.month, 'Mon') AS month, COALESCE(SUM(t.transfer_amount), 0)::int AS total_received_amount
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

const getYearlyTransferAmountReceiver = `WITH
    last_five_years AS (
        SELECT EXTRACT(
                YEAR
                FROM t.transfer_time
            ) AS year, SUM(t.transfer_amount) AS total_received_amount
        FROM transfers t
        WHERE
            t.deleted_at IS NULL
            AND EXTRACT(
                YEAR
                FROM t.transfer_time
            ) >= $1 - 4
            AND EXTRACT(
                YEAR
                FROM t.transfer_time
            ) <= $1
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.transfer_time
            )
    )
SELECT year, total_received_amount
FROM last_five_years
ORDER BY year`

type cardStaticsTransferRepository struct {
	db *sqlx.DB
}

func NewCardStatsTransferRepository(db *sqlx.DB) CardStatsTransferRepository {
	return &cardStaticsTransferRepository{
		db: db,
	}
}

func (r *cardStaticsTransferRepository) GetMonthlyTransferAmountSender(ctx context.Context, year int) ([]*db.GetMonthlyTransferAmountSenderRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var res []*db.GetMonthlyTransferAmountSenderRow
	err := r.db.SelectContext(ctx, &res, getMonthlyTransferAmountSender, yearStart)

	if err != nil {
		return nil, card_errors.ErrGetMonthlyTransferAmountSenderFailed
	}

	return res, nil
}

func (r *cardStaticsTransferRepository) GetYearlyTransferAmountSender(ctx context.Context, year int) ([]*db.GetYearlyTransferAmountSenderRow, error) {
	var res []*db.GetYearlyTransferAmountSenderRow
	err := r.db.SelectContext(ctx, &res, getYearlyTransferAmountSender, int32(year))

	if err != nil {
		return nil, card_errors.ErrGetYearlyTransferAmountSenderFailed
	}

	return res, nil
}

func (r *cardStaticsTransferRepository) GetMonthlyTransferAmountReceiver(ctx context.Context, year int) ([]*db.GetMonthlyTransferAmountReceiverRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	var res []*db.GetMonthlyTransferAmountReceiverRow
	err := r.db.SelectContext(ctx, &res, getMonthlyTransferAmountReceiver, yearStart)

	if err != nil {
		return nil, card_errors.ErrGetMonthlyTransferAmountReceiverFailed.WithInternal(err)
	}

	return res, nil
}

func (r *cardStaticsTransferRepository) GetYearlyTransferAmountReceiver(ctx context.Context, year int) ([]*db.GetYearlyTransferAmountReceiverRow, error) {
	var res []*db.GetYearlyTransferAmountReceiverRow
	err := r.db.SelectContext(ctx, &res, getYearlyTransferAmountReceiver, int32(year))

	if err != nil {
		return nil, card_errors.ErrGetYearlyTransferAmountReceiverFailed.WithInternal(err)
	}

	return res, nil
}
