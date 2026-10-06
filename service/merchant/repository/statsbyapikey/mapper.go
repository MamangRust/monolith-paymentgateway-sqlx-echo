package merchantstatsapikeyrepository

import (
	"github.com/jmoiron/sqlx"
)

type MerchantStatsByApiKeyRepository interface {
	MerchantStatsAmountByApiKeyRepository
	MerchantStatsMethodByApiKeyRepository
	MerchantStatsTotalAmountByApiKeyRepository
}

type repository struct {
	MerchantStatsAmountByApiKeyRepository
	MerchantStatsMethodByApiKeyRepository
	MerchantStatsTotalAmountByApiKeyRepository
}

func NewMerchantStatsByApiKeyRepository(db *sqlx.DB) MerchantStatsByApiKeyRepository {

	return &repository{
		MerchantStatsAmountByApiKeyRepository:      NewMerchantStatsAmountByApiKeyRepository(db),
		MerchantStatsMethodByApiKeyRepository:      NewMerchantStatsMethodByApiKeyRepository(db),
		MerchantStatsTotalAmountByApiKeyRepository: NewMerchantStatsTotalAmountByApiKeyRepository(db),
	}
}
