package merchantstatsmerchantrepository

import (
	"github.com/jmoiron/sqlx"
)

type MerchantStatsByMerchantRepository interface {
	MerchantStatsAmountByMerchantRepository
	MerchantStatsMethodByMerchantRepository
	MerchantStatsTotalAmountByMerchantRepository
}

type repository struct {
	MerchantStatsAmountByMerchantRepository
	MerchantStatsMethodByMerchantRepository
	MerchantStatsTotalAmountByMerchantRepository
}

func NewMerchantStatsByMerchantRepository(db *sqlx.DB) MerchantStatsByMerchantRepository {
	return &repository{
		MerchantStatsAmountByMerchantRepository:      NewMerchantStatsAmountByMerchantRepository(db),
		MerchantStatsMethodByMerchantRepository:      NewMerchantStatsMethodByMerchantRepository(db),
		MerchantStatsTotalAmountByMerchantRepository: NewMerchantStatsTotalAmountByMerchantRepository(db),
	}
}
