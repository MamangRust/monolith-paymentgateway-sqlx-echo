package repositorydashboard

import (
	"context"

	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getTotalTransferAmount = `SELECT SUM(transfer_amount) AS total_transfer_amount
FROM transfers
WHERE
    deleted_at IS NULL`

const getTotalTransferAmountBySender = `SELECT SUM(transfer_amount) AS total_transfer_amount
FROM transfers
WHERE
    transfer_from = $1
    AND deleted_at IS NULL`

const getTotalTransferAmountByReceiver = `SELECT SUM(transfer_amount) AS total_transfer_amount
FROM transfers
WHERE
    transfer_to = $1
    AND deleted_at IS NULL`

type cardDashboardTransferRepository struct {
	db *sqlx.DB
}

func NewCardDashboardTransferRepository(db *sqlx.DB) CardDashboardTransferRepository {
	return &cardDashboardTransferRepository{
		db: db,
	}
}

func (r *cardDashboardTransferRepository) GetTotalTransferAmount(ctx context.Context) (*int64, error) {
	var res int64
	err := r.db.GetContext(ctx, &res, getTotalTransferAmount)

	if err != nil {
		return nil, card_errors.ErrGetTotalTransfersFailed.WithInternal(err)
	}

	return &res, nil
}

func (r *cardDashboardTransferRepository) GetTotalTransferAmountBySender(ctx context.Context, senderCardNumber string) (*int64, error) {
	var res int64
	err := r.db.GetContext(ctx, &res, getTotalTransferAmountBySender, senderCardNumber)

	if err != nil {
		return nil, card_errors.ErrGetTotalTransferSenderByCardFailed.WithInternal(err)
	}

	return &res, nil
}

func (r *cardDashboardTransferRepository) GetTotalTransferAmountByReceiver(ctx context.Context, receiverCardNumber string) (*int64, error) {
	var res int64
	err := r.db.GetContext(ctx, &res, getTotalTransferAmountByReceiver, receiverCardNumber)

	if err != nil {
		return nil, card_errors.ErrGetTotalTransferSenderByCardFailed.WithInternal(err)
	}

	return &res, nil
}
