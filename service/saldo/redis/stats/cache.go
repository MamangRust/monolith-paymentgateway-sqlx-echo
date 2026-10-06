package saldostatscache

import sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"

type SaldoStatsCache interface {
	SaldoStatsBalanceCache
	SaldoStatsTotalCache
}

type saldoStatsCache struct {
	SaldoStatsBalanceCache
	SaldoStatsTotalCache
}

func NewSaldoStatsCache(store *sharedcachehelpers.CacheStore) SaldoStatsCache {
	return &saldoStatsCache{
		SaldoStatsBalanceCache: NewSaldoStatsBalanceCache(store),
		SaldoStatsTotalCache:   NewSaldoStatsTotalCache(store),
	}
}
