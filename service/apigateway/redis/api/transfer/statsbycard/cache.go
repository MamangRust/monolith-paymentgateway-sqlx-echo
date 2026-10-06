package transfer_stats_bycard_cache

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type TransferStatsByCardCache interface {
	TransferStatsByCardAmountCache
	TransferStatsByCardStatusCache
}

type transferStatsByCardCache struct {
	TransferStatsByCardAmountCache
	TransferStatsByCardStatusCache
}

func NewTransferStatsByCardCache(store *cache.CacheStore) TransferStatsByCardCache {
	return &transferStatsByCardCache{
		TransferStatsByCardAmountCache: NewTransferStatsByCardAmountCache(store),
		TransferStatsByCardStatusCache: NewTransferStatsByCardStatusCache(store),
	}
}
