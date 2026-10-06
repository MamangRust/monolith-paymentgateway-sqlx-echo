package mencache

import (
	merchantstatscache "github.com/MamangRust/monolith-payment-gateway-merchant/redis/stats"
	merchantstatsapikey "github.com/MamangRust/monolith-payment-gateway-merchant/redis/statsbyapikey"
	merchantstatsbymerchant "github.com/MamangRust/monolith-payment-gateway-merchant/redis/statsbymerchant"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type mencache struct {
	MerchantQueryCache
	MerchantCommandCache
	MerchantDocumentQueryCache
	MerchantDocumentCommandCache
	MerchantTransactionCache
	merchantstatscache.MerchantStatsCache
	merchantstatsapikey.MerchantStatsByApiKeyCache
	merchantstatsbymerchant.MerchantStatsByMerchantCache
}

type Mencache interface {
	MerchantQueryCache
	MerchantCommandCache
	MerchantDocumentQueryCache
	MerchantDocumentCommandCache
	MerchantTransactionCache
	merchantstatscache.MerchantStatsCache
	merchantstatsapikey.MerchantStatsByApiKeyCache
	merchantstatsbymerchant.MerchantStatsByMerchantCache
}

func NewMencache(cacheStore *cache.CacheStore) Mencache {
	return &mencache{
		MerchantQueryCache:           NewMerchantQueryCache(cacheStore),
		MerchantCommandCache:         NewMerchantCommandCache(cacheStore),
		MerchantDocumentQueryCache:   NewMerchantDocumentQueryCache(cacheStore),
		MerchantDocumentCommandCache: NewMerchantDocumentCommandCache(cacheStore),
		MerchantTransactionCache:     NewMerchantTransactionCache(cacheStore),
		MerchantStatsCache:           merchantstatscache.NewMerchantStatsCache(cacheStore),
		MerchantStatsByApiKeyCache:   merchantstatsapikey.NewMerchantStatsByApiKeyCache(cacheStore),
		MerchantStatsByMerchantCache: merchantstatsbymerchant.NewMerchantStatsByMerchantCache(cacheStore),
	}
}
