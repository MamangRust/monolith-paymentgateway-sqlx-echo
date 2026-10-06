package transferstatsrepository

import (
	"github.com/jmoiron/sqlx"
)

type TransferStatsRepository interface {
	TransferStatsAmountRepository
	TransferStatsStatusRepository
}

type repositories struct {
	TransferStatsAmountRepository
	TransferStatsStatusRepository
}

func NewTransferStatsRepository(db *sqlx.DB) TransferStatsRepository {

	return &repositories{
		TransferStatsAmountRepository: NewTransferStatsAmountRepository(db),
		TransferStatsStatusRepository: NewTransferStatsStatusRepository(db),
	}
}
