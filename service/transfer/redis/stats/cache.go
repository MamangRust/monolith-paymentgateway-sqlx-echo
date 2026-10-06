package transferstatscache

import (
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type TransferStatsCache interface {
	TransferStatsAmountCache
	TransferStatsStatusCache
}

type transferStatsCache struct {
	TransferStatsAmountCache
	TransferStatsStatusCache
}

func NewTransferStatsCache(store *sharedcachehelpers.CacheStore) TransferStatsCache {
	return &transferStatsCache{
		TransferStatsAmountCache: NewTransferStatsAmountCache(store),
		TransferStatsStatusCache: NewTransferStatsStatusCache(store),
	}
}
