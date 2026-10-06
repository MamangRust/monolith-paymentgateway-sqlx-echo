package transfer_cache

import (
	transfer_stats_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/transfer/stats"
	transfer_stats_bycard_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/transfer/statsbycard"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type TransferMencache interface {
	TransferQueryCache
	TransferCommandCache
	transfer_stats_cache.TransferStatsCache
	transfer_stats_bycard_cache.TransferStatsByCardCache
}

type transfermencache struct {
	TransferQueryCache
	TransferCommandCache
	transfer_stats_cache.TransferStatsCache
	transfer_stats_bycard_cache.TransferStatsByCardCache
}

func NewTransferMencache(cacheStore *cache.CacheStore) TransferMencache {
	return &transfermencache{
		TransferQueryCache:       NewTransferQueryCache(cacheStore),
		TransferCommandCache:     NewTransferCommandCache(cacheStore),
		TransferStatsCache:       transfer_stats_cache.NewTransferStatsCache(cacheStore),
		TransferStatsByCardCache: transfer_stats_bycard_cache.NewTransferStatsByCardCache(cacheStore),
	}
}
