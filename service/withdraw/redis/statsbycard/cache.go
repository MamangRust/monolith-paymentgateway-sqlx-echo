package withdrawstatsbycardcache

import sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"

type WithdrawStatsByCardCache interface {
	WithdrawStatsByCardAmountCache
	WithdrawStatsByCardStatusCache
}

type withdrawStatsByCardCache struct {
	WithdrawStatsByCardAmountCache
	WithdrawStatsByCardStatusCache
}

func NewWithdrawStatsByCardCache(store *sharedcachehelpers.CacheStore) WithdrawStatsByCardCache {
	return &withdrawStatsByCardCache{
		WithdrawStatsByCardAmountCache: NewWithdrawStatsAmountCache(store),
		WithdrawStatsByCardStatusCache: NewWithdrawStatsStatusCache(store),
	}
}
