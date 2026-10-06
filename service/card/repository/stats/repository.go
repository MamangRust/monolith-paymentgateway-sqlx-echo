package repositorystats

import (
	"github.com/jmoiron/sqlx"
)

type CardStatsRepository interface {
	CardStatsBalanceRepository
	CardStatsTopupRepository
	CardStatsTransactionRepository
	CardStatsTransferRepository
	CardStatsWithdrawRepository
}

type repository struct {
	CardStatsBalanceRepository
	CardStatsTopupRepository
	CardStatsTransactionRepository
	CardStatsTransferRepository
	CardStatsWithdrawRepository
}

func NewCardStatsRepository(db *sqlx.DB) CardStatsRepository {

	return &repository{
		CardStatsBalanceRepository:     NewCardStatsBalanceRepository(db),
		CardStatsTopupRepository:       NewCardStatsTopupRepository(db),
		CardStatsTransactionRepository: NewCardStatsTransactionRepository(db),
		CardStatsTransferRepository:    NewCardStatsTransferRepository(db),
		CardStatsWithdrawRepository:    NewCardStatsWithdrawRepository(db),
	}
}
