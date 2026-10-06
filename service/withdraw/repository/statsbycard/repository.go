package withdrawstatsbycardrepository

import (
	"github.com/jmoiron/sqlx"
)

type WithdrawStatsByCardRepository interface {
	WithdrawStatsByCardAmountRepository
	WithdrawStatsByCardStatusRepository
}

type repositories struct {
	WithdrawStatsByCardAmountRepository
	WithdrawStatsByCardStatusRepository
}

func NewWithdrawStatsByCardRepository(db *sqlx.DB) WithdrawStatsByCardRepository {

	return &repositories{
		WithdrawStatsByCardAmountRepository: NewWithdrawStatsByCardAmountRepository(db),
		WithdrawStatsByCardStatusRepository: NewWithdrawStatsByCardStatusRepository(db),
	}
}
