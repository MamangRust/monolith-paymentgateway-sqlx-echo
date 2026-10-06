package transactionstatsrepository

import (
	"github.com/jmoiron/sqlx"
)

type TransactionStatsRepository interface {
	TransactionStatsAmountRepository
	TransactionStatsMethodRepository
	TransactionStatsStatusRepository
}

type repository struct {
	TransactionStatsAmountRepository
	TransactionStatsMethodRepository
	TransactionStatsStatusRepository
}

func NewTransactionStatsRepository(db *sqlx.DB) TransactionStatsRepository {

	return &repository{
		TransactionStatsAmountRepository: NewTransactionStatsAmountRepository(db),
		TransactionStatsMethodRepository: NewTransactionStatsMethodRepository(db),
		TransactionStatsStatusRepository: NewTransactionStatsStatusRepository(db),
	}
}
