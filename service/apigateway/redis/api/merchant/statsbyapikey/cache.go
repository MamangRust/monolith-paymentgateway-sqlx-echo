package merchant_stats_byapikey_cache

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type MerchantStatsByApiKeyCache interface {
	MerchantStatsAmountByApiKeyCache
	MerchantStatsMethodByApiKeyCache
	MerchantStatsTotalAmountByApiKeyCache
}

type merchantStatsByApiKeyCache struct {
	MerchantStatsAmountByApiKeyCache
	MerchantStatsMethodByApiKeyCache
	MerchantStatsTotalAmountByApiKeyCache
}

func NewMerchantStatsByApiKeyCache(store *cache.CacheStore) MerchantStatsByApiKeyCache {
	return &merchantStatsByApiKeyCache{
		MerchantStatsAmountByApiKeyCache:      NewMerchantStatsAmountByApiKeyCache(store),
		MerchantStatsMethodByApiKeyCache:      NewMerchantStatsMethodByApiKeyCache(store),
		MerchantStatsTotalAmountByApiKeyCache: NewMerchantStatsTotalAmountByApiKeyCache(store),
	}
}
