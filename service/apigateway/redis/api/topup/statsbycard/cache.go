package topup_stats_bycard_cache

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

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

func NewTopupStatsByCardCache(store *cache.CacheStore) TopupStatsByCardCache {
	return &topupStatsByCardCache{
		TopupStatsAmountByCardCache: NewTopupStatsAmountByCardCache(store),
		TopupStatsMethodByCardCache: NewTopupStatsMethodByCardCache(store),
		TopupStatsStatusByCardCache: NewTopupStatsStatusByCardCache(store),
	}
}
