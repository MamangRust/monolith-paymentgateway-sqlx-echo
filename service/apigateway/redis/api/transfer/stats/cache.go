package transfer_stats_cache

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type TransferStatsCache interface {
	TransferStatsAmountCache
	TransferStatsStatusCache
}

type transferStatsCache struct {
	TransferStatsAmountCache
	TransferStatsStatusCache
}

func NewTransferStatsCache(store *cache.CacheStore) TransferStatsCache {
	return &transferStatsCache{
		TransferStatsAmountCache: NewTransferStatsAmountCache(store),
		TransferStatsStatusCache: NewTransferStatsStatusCache(store),
	}
}
