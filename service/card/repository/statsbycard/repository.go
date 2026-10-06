package repositorystatsbycard

import (
	"github.com/jmoiron/sqlx"
)

type CardStatsByCardRepository interface {
	CardStatsBalanceByCardRepository
	CardStatsTopupByCardRepository
	CardStatsTransactionByCardRepository
	CardStatsTransferByCardRepository
	CardStatsWithdrawByCardRepository
}

type repository struct {
	CardStatsBalanceByCardRepository
	CardStatsTopupByCardRepository
	CardStatsTransactionByCardRepository
	CardStatsTransferByCardRepository
	CardStatsWithdrawByCardRepository
}

func NewCardStatsByCardRepository(db *sqlx.DB) CardStatsByCardRepository {
	return &repository{
		CardStatsBalanceByCardRepository:     NewCardStatsBalanceByCardRepository(db),
		CardStatsTopupByCardRepository:       NewCardStatsTopupByCardRepository(db),
		CardStatsTransactionByCardRepository: NewCardStatsTransactionByCardRepository(db),
		CardStatsTransferByCardRepository:    NewCardStatsTransferByCardRepository(db),
		CardStatsWithdrawByCardRepository:    NewCardStatsWithdrawByCardRepository(db),
	}
}
