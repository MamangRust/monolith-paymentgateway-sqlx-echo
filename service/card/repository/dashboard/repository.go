package repositorydashboard

import (
	"github.com/jmoiron/sqlx"
)

type CardDashboardRepository interface {
	CardDashboardBalanceRepository
	CardDashboardTopupRepository
	CardDashboardTransactionRepository
	CardDashboardTransferRepository
	CardDashboardWithdrawRepository
}

type repository struct {
	CardDashboardBalanceRepository
	CardDashboardTopupRepository
	CardDashboardTransactionRepository
	CardDashboardTransferRepository
	CardDashboardWithdrawRepository
}

func NewCardDashboardRepository(db *sqlx.DB) CardDashboardRepository {
	return &repository{
		CardDashboardBalanceRepository:     NewCardDashboardBalanceRepository(db),
		CardDashboardTopupRepository:       NewCardDashboardTopupRepository(db),
		CardDashboardTransactionRepository: NewCardDashboardTransactionRepository(db),
		CardDashboardTransferRepository:    NewCardDashboardTransferRepository(db),
		CardDashboardWithdrawRepository:    NewCardDashboardWithdrawRepository(db),
	}
}
