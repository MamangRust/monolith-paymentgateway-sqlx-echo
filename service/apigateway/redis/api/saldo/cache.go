package saldo_cache

import (
	saldo_stats_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/saldo/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type SaldoMencache interface {
	SaldoQueryCache
	SaldoCommandCache
	saldo_stats_cache.SaldoStatsCache
}

type saldomencache struct {
	SaldoQueryCache
	SaldoCommandCache
	saldo_stats_cache.SaldoStatsCache
}

func NewSaldoMencache(cacheStore *cache.CacheStore) SaldoMencache {
	return &saldomencache{
		SaldoQueryCache:   NewSaldoQueryCache(cacheStore),
		SaldoCommandCache: NewSaldoCommandCache(cacheStore),
		SaldoStatsCache:   saldo_stats_cache.NewSaldoStatsCache(cacheStore),
	}
}
