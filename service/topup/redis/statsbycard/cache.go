package topupstatsbycardcache

import sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"

type TopupStatsByCardCache interface {
	TopupStatsAmountByCardCache
	TopupStatsMethodByCardCache
	TopupStatsStatusByCardCache
}

type topupStatsByCardCache struct {
	TopupStatsAmountByCardCache
	TopupStatsMethodByCardCache
	TopupStatsStatusByCardCache
}

func NewTopupStatsByCardCache(store *sharedcachehelpers.CacheStore) TopupStatsByCardCache {
	return &topupStatsByCardCache{
		TopupStatsAmountByCardCache: NewTopupStatsAmountByCardCache(store),
		TopupStatsMethodByCardCache: NewTopupStatsMethodByCardCache(store),
		TopupStatsStatusByCardCache: NewTopupStatsStatusByCardCache(store),
	}
}
