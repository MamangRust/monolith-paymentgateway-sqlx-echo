package mencache

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	withdrawstatscache "github.com/MamangRust/monolith-payment-gateway-withdraw/redis/stats"
	withdrawstatsbycardcache "github.com/MamangRust/monolith-payment-gateway-withdraw/redis/statsbycard"
)

// Mencache represents a cache store for withdraw operations.
type Mencache interface {
	WithdrawQueryCache
	WithdrawCommandCache
	withdrawstatscache.WithdrawStatsCache
	withdrawstatsbycardcache.WithdrawStatsByCardCache
}

type mencache struct {
	WithdrawQueryCache
	WithdrawCommandCache
	withdrawstatscache.WithdrawStatsCache
	withdrawstatsbycardcache.WithdrawStatsByCardCache
}

func NewMencache(cacheStore *cache.CacheStore) Mencache {
	return &mencache{
		WithdrawQueryCache:       NewWithdrawQueryCache(cacheStore),
		WithdrawCommandCache:     NewWithdrawCommandCache(cacheStore),
		WithdrawStatsCache:       withdrawstatscache.NewWithdrawStatsCache(cacheStore),
		WithdrawStatsByCardCache: withdrawstatsbycardcache.NewWithdrawStatsByCardCache(cacheStore),
	}
}
