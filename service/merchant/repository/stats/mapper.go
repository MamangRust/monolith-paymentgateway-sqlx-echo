package merchantstatsrepository

import (
	"github.com/jmoiron/sqlx"
)

type MerchantStatsRepository interface {
	MerchantStatsAmountRepository
	MerchantStatsMethodRepository
	MerchantStatsTotalAmountRepository
}

type repository struct {
	MerchantStatsAmountRepository
	MerchantStatsMethodRepository
	MerchantStatsTotalAmountRepository
}

func NewMerchantStatsRepository(db *sqlx.DB) MerchantStatsRepository {
	return &repository{
		MerchantStatsAmountRepository:      NewMerchantStatsAmountRepository(db),
		MerchantStatsMethodRepository:      NewMerchantStatsMethodRepository(db),
		MerchantStatsTotalAmountRepository: NewMerchantStatsTotalAmountRepository(db),
	}
}
