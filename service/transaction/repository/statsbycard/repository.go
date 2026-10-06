package transactionbycardrepository

import (
	"github.com/jmoiron/sqlx"
)

type TransactionStatsByCardRepository interface {
	TransactionStatsByCardAmountRepository
	TransactionStatsByCardMethodRepository
	TransactionStatsByCardStatusRepository
}

type repository struct {
	TransactionStatsByCardAmountRepository
	TransactionStatsByCardMethodRepository
	TransactionStatsByCardStatusRepository
}

func NewTransactionStatsRepository(db *sqlx.DB) TransactionStatsByCardRepository {

	return &repository{
		TransactionStatsByCardAmountRepository:  NewTransactionStatsByCardAmountRepository(db),
		TransactionStatsByCardMethodRepository: NewTransactionStatsByCardMethodRepository(db),
		TransactionStatsByCardStatusRepository:  NewTransactionStatsByCardStatusRepository(db),
	}
}

