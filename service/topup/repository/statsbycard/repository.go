package topupstatsbycardrepository

import (
	"github.com/jmoiron/sqlx"
)

type TopupStatsByCardRepository interface {
	TopupStatsByCardAmountRepository
	TopupStatsByCardMethodRepository
	TopupStatsByCardStatusRepository
}

type repository struct {
	TopupStatsByCardAmountRepository
	TopupStatsByCardMethodRepository
	TopupStatsByCardStatusRepository
}

func NewTopupStatsByCardRepository(db *sqlx.DB) TopupStatsByCardRepository {

	return &repository{
		TopupStatsByCardAmountRepository: NewTopupStatsByCardAmountRepository(db),
		TopupStatsByCardMethodRepository: NewTopupStatsByCardMethodRepository(db),
		TopupStatsByCardStatusRepository: NewTopupStatsByCardStatusRepository(db),
	}
}
