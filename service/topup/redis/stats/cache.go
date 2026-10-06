package topupstatscache

import sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"

type TopupStatsCache interface {
	TopupStatsAmountCache
	TopupStatsMethodCache
	TopupStatsStatusCache
}

type topupStatsCache struct {
	TopupStatsAmountCache
	TopupStatsMethodCache
	TopupStatsStatusCache
}

func NewTopupStatsCache(store *sharedcachehelpers.CacheStore) TopupStatsCache {
	return &topupStatsCache{
		TopupStatsAmountCache: NewTopupStatsAmountCache(store),
		TopupStatsMethodCache: NewTopupStatsMethodCache(store),
		TopupStatsStatusCache: NewTopupStatsStatusCache(store),
	}
}
