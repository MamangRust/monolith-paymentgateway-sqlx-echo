package mencache

import (
	saldostatscache "github.com/MamangRust/monolith-payment-gateway-saldo/redis/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type Mencache interface {
	SaldoQueryCache
	SaldoCommandCache
	saldostatscache.SaldoStatsCache
}

type mencache struct {
	SaldoQueryCache
	SaldoCommandCache
	saldostatscache.SaldoStatsCache
}

func NewMencache(cacheStore *cache.CacheStore) Mencache {
	return &mencache{
		SaldoQueryCache:   NewSaldoQueryCache(cacheStore),
		SaldoCommandCache: NewSaldoCommandCache(cacheStore),
		SaldoStatsCache:   saldostatscache.NewSaldoStatsCache(cacheStore),
	}
}
