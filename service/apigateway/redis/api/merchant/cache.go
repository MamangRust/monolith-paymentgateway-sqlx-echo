package merchant_cache

import (
	merchant_stats_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/merchant/stats"
	merchant_stats_byapikey_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/merchant/statsbyapikey"
	merchant_stats_bymerchant_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/merchant/statsbymerchant"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type mencache struct {
	MerchantQueryCache
	MerchantCommandCache
	MerchantTransactionCache
	merchant_stats_cache.MerchantStatsCache
	merchant_stats_byapikey_cache.MerchantStatsByApiKeyCache
	merchant_stats_bymerchant_cache.MerchantStatsByMerchantCache
}

type MerchantMencache interface {
	MerchantQueryCache
	MerchantCommandCache
	MerchantTransactionCache
	merchant_stats_cache.MerchantStatsCache
	merchant_stats_byapikey_cache.MerchantStatsByApiKeyCache
	merchant_stats_bymerchant_cache.MerchantStatsByMerchantCache
}

func NewMerchantMencache(cacheStore *cache.CacheStore) MerchantMencache {

	return &mencache{
		MerchantQueryCache:   NewMerchantQueryCache(cacheStore),
		MerchantCommandCache: NewMerchantCommandCache(cacheStore),

		MerchantTransactionCache:     NewMerchantTransactionCache(cacheStore),
		MerchantStatsCache:           merchant_stats_cache.NewMerchantStatsCache(cacheStore),
		MerchantStatsByApiKeyCache:   merchant_stats_byapikey_cache.NewMerchantStatsByApiKeyCache(cacheStore),
		MerchantStatsByMerchantCache: merchant_stats_bymerchant_cache.NewMerchantStatsByMerchantCache(cacheStore),
	}
}
