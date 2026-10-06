package merchantstatscache

import sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"

type MerchantStatsCache interface {
	MerchantStatsAmountCache
	MerchantStatsMethodCache
	MerchantStatsTotalAmountCache
}

type merchantStatsCache struct {
	MerchantStatsAmountCache
	MerchantStatsMethodCache
	MerchantStatsTotalAmountCache
}

func NewMerchantStatsCache(store *sharedcachehelpers.CacheStore) MerchantStatsCache {
	return &merchantStatsCache{
		MerchantStatsAmountCache:      NewMerchantStatsAmountCache(store),
		MerchantStatsMethodCache:      NewMerchantStatsMethodCache(store),
		MerchantStatsTotalAmountCache: NewMerchantStatsTotalAmountCache(store),
	}
}
