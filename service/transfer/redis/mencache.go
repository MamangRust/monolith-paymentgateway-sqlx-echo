package mencache

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	transferstatscache "github.com/MamangRust/monolith-payment-gateway-transfer/redis/stats"
	transferstatsbycardcache "github.com/MamangRust/monolith-payment-gateway-transfer/redis/statsbycard"
)

type Mencache interface {
	TransferQueryCache
	TransferCommandCache
	transferstatscache.TransferStatsCache
	transferstatsbycardcache.TransferStatsByCardCache
}

type mencache struct {
	TransferQueryCache
	TransferCommandCache
	transferstatscache.TransferStatsCache
	transferstatsbycardcache.TransferStatsByCardCache
}

func NewMencache(cacheStore *cache.CacheStore) Mencache {
	return &mencache{
		TransferQueryCache:       NewTransferQueryCache(cacheStore),
		TransferCommandCache:     NewTransferCommandCache(cacheStore),
		TransferStatsCache:       transferstatscache.NewTransferStatsCache(cacheStore),
		TransferStatsByCardCache: transferstatsbycardcache.NewTransferStatsByCardCache(cacheStore),
	}
}
