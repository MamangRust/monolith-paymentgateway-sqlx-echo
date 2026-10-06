package withdraw_cache

import (
	withdraw_stats_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/withdraw/stats"
	withdraw_stats_bycard_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/withdraw/statsbycard"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type WithdrawMencache interface {
	WithdrawQueryCache
	WithdrawCommandCache
	withdraw_stats_cache.WithdrawStatsCache
	withdraw_stats_bycard_cache.WithdrawStatsByCardCache
}

type withdrawmencache struct {
	WithdrawQueryCache
	WithdrawCommandCache
	withdraw_stats_cache.WithdrawStatsCache
	withdraw_stats_bycard_cache.WithdrawStatsByCardCache
}

func NewWithdrawMencache(cacheStore *cache.CacheStore) WithdrawMencache {
	return &withdrawmencache{
		WithdrawQueryCache:       NewWithdrawQueryCache(cacheStore),
		WithdrawCommandCache:     NewWithdrawCommandCache(cacheStore),
		WithdrawStatsCache:       withdraw_stats_cache.NewWithdrawStatsCache(cacheStore),
		WithdrawStatsByCardCache: withdraw_stats_bycard_cache.NewWithdrawStatsByCardCache(cacheStore),
	}
}
