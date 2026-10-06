package withdraw_stats_cache

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type WithdrawStatsCache interface {
	WithdrawStatsAmountCache
	WithdrawStatsStatusCache
}

type withdrawStatsCache struct {
	WithdrawStatsAmountCache
	WithdrawStatsStatusCache
}

func NewWithdrawStatsCache(store *cache.CacheStore) WithdrawStatsCache {
	return &withdrawStatsCache{
		WithdrawStatsAmountCache: NewWithdrawStatsAmountCache(store),
		WithdrawStatsStatusCache: NewWithdrawStatsStatusCache(store),
	}
}
