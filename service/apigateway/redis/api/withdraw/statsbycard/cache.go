package withdraw_stats_bycard_cache

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type WithdrawStatsByCardCache interface {
	WithdrawStatsByCardAmountCache
	WithdrawStatsByCardStatusCache
}

type withdrawStatsByCardCache struct {
	WithdrawStatsByCardAmountCache
	WithdrawStatsByCardStatusCache
}

func NewWithdrawStatsByCardCache(store *cache.CacheStore) WithdrawStatsByCardCache {
	return &withdrawStatsByCardCache{
		WithdrawStatsByCardAmountCache: NewWithdrawStatsByCardAmountCache(store),
		WithdrawStatsByCardStatusCache: NewWithdrawStatsByCardStatusCache(store),
	}
}
