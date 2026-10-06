package merchant_stats_cache

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

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

func NewMerchantStatsCache(store *cache.CacheStore) MerchantStatsCache {
	return &merchantStatsCache{
		MerchantStatsAmountCache:      NewMerchantStatsAmountCache(store),
		MerchantStatsMethodCache:      NewMerchantStatsMethodCache(store),
		MerchantStatsTotalAmountCache: NewMerchantStatsTotalAmountCache(store),
	}
}
