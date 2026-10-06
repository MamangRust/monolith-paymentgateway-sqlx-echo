package mencache

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	topupstatscache "github.com/MamangRust/monolith-payment-gateway-topup/redis/stats"
	topupstatsbycardcache "github.com/MamangRust/monolith-payment-gateway-topup/redis/statsbycard"
)

type Mencache interface {
	TopupQueryCache
	TopupCommandCache
	topupstatscache.TopupStatsCache
	topupstatsbycardcache.TopupStatsByCardCache
}

type mencache struct {
	TopupQueryCache
	TopupCommandCache
	topupstatscache.TopupStatsCache
	topupstatsbycardcache.TopupStatsByCardCache
}

// NewMencache creates a new Mencache instance using the given dependencies.
// It creates a new cache store using the given context, Redis client, and logger,
// and returns a Mencache struct with initialized caches for topup query, topup command,
// topup statistic, and topup statistic by card.
func NewMencache(cacheStore *cache.CacheStore) Mencache {

	return &mencache{
		TopupQueryCache:       NewTopupQueryCache(cacheStore),
		TopupCommandCache:     NewTopupCommandCache(cacheStore),
		TopupStatsCache:       topupstatscache.NewTopupStatsCache(cacheStore),
		TopupStatsByCardCache: topupstatsbycardcache.NewTopupStatsByCardCache(cacheStore),
	}
}
