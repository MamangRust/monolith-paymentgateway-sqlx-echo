package transferstatsbycardrepository

import (
	"github.com/jmoiron/sqlx"
)

type TransferStatsByCardRepository interface {
	TransferStatsByCardAmountSenderRepository
	TransferStatsByCardAmountReceiverRepository
	TransferStatsByCardStatusRepository
}

type repositories struct {
	TransferStatsByCardAmountSenderRepository
	TransferStatsByCardAmountReceiverRepository
	TransferStatsByCardStatusRepository
}

func NewTransferStatsByCardRepository(db *sqlx.DB) TransferStatsByCardRepository {

	return &repositories{
		TransferStatsByCardAmountSenderRepository:   NewTransferStatsAmountSenderRepository(db),
		TransferStatsByCardAmountReceiverRepository: NewTransferStatsAmountReceiverRepository(db),
		TransferStatsByCardStatusRepository:         NewTransferStatsByCardStatusRepository(db),
	}
}
