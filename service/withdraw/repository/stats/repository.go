package withdrawstatsrepository

import (
	"github.com/jmoiron/sqlx"
)

type WithdrawStatsRepository interface {
	WithdrawStatsAmountRepository
	WithdrawStatsStatusRepository
}

type repositories struct {
	WithdrawStatsAmountRepository
	WithdrawStatsStatusRepository
}

func NewWithdrawStatsRepository(db *sqlx.DB) WithdrawStatsRepository {

	return &repositories{
		WithdrawStatsAmountRepository: NewWithdrawStatsAmountRepository(db),
		WithdrawStatsStatusRepository: NewWithdrawStatsStatusRepository(db),
	}
}
