package topupstatsrepository

import (
	"github.com/jmoiron/sqlx"
)

type TopupStatsRepository interface {
	TopupStatsAmountRepository
	TopupStatsStatusRepository
	TopupStatsMethodRepository
}

type repository struct {
	TopupStatsAmountRepository
	TopupStatsStatusRepository
	TopupStatsMethodRepository
}

func NewTopupStatsRepository(db *sqlx.DB) TopupStatsRepository {
	return &repository{
		TopupStatsAmountRepository: NewTopupStatsAmountRepository(db),
		TopupStatsStatusRepository: NewTopupStatsStatusRepository(db),
		TopupStatsMethodRepository: NewTopupStatsMethodRepository(db),
	}
}
