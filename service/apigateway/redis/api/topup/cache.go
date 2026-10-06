package topup_cache

import (
	topup_stats_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/topup/stats"
	topup_stats_bycard_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/topup/statsbycard"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type TopupMencach interface {
	TopupQueryCache
	TopupCommandCache
	topup_stats_cache.TopupStatsCache
	topup_stats_bycard_cache.TopupStatsByCardCache
}

type mencache struct {
	TopupQueryCache
	TopupCommandCache
	topup_stats_cache.TopupStatsCache
	topup_stats_bycard_cache.TopupStatsByCardCache
}

func NewTopupMencache(cacheStore *cache.CacheStore) TopupMencach {

	return &mencache{
		TopupQueryCache:       NewTopupQueryCache(cacheStore),
		TopupCommandCache:     NewTopupCommandCache(cacheStore),
		TopupStatsCache:       topup_stats_cache.NewTopupStatsCache(cacheStore),
		TopupStatsByCardCache: topup_stats_bycard_cache.NewTopupStatsByCardCache(cacheStore),
	}
}
