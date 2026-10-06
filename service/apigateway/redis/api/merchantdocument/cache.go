package merchantdocument_cache

import "github.com/MamangRust/monolith-payment-gateway-shared/cache"

type mencache struct {
	MerchantDocumentQueryCache
	MerchantDocumentCommandCache
}

type MerchantDocumentMencache interface {
	MerchantDocumentQueryCache
	MerchantDocumentCommandCache
}

func NewMerchantDocumentMencache(cacheStore *cache.CacheStore) MerchantDocumentMencache {
	return &mencache{
		MerchantDocumentQueryCache:   NewMerchantDocumentQueryCache(cacheStore),
		MerchantDocumentCommandCache: NewMerchantDocumentCommandCache(cacheStore),
	}
}
