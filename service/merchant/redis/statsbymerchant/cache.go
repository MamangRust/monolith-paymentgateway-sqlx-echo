package merchantstatsbymerchant

import sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"

type MerchantStatsByMerchantCache interface {
	MerchantStatsAmountByMerchantCache
	MerchantStatsMethodByMerchantCache
	MerchantStatsTotalAmountByMerchantCache
}

type merchantStatsByMerchantCache struct {
	MerchantStatsAmountByMerchantCache
	MerchantStatsMethodByMerchantCache
	MerchantStatsTotalAmountByMerchantCache
}

func NewMerchantStatsByMerchantCache(store *sharedcachehelpers.CacheStore) MerchantStatsByMerchantCache {
	return &merchantStatsByMerchantCache{
		MerchantStatsAmountByMerchantCache:      NewMerchantStatsAmountByMerchantCache(store),
		MerchantStatsMethodByMerchantCache:      NewMerchantStatsMethodByMerchantCache(store),
		MerchantStatsTotalAmountByMerchantCache: NewMerchantStatsTotalAmountByMerchantCache(store),
	}
}
