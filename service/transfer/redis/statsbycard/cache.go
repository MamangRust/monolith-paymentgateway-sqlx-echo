package transferstatsbycardcache

import (
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type TransferStatsByCardCache interface {
	TransferStatsByCardAmountCache
	TransferStatsByCardStatusCache
}

type transferStatsByCardCache struct {
	TransferStatsByCardAmountCache
	TransferStatsByCardStatusCache
}

func NewTransferStatsByCardCache(store *sharedcachehelpers.CacheStore) TransferStatsByCardCache {
	return &transferStatsByCardCache{
		TransferStatsByCardAmountCache: NewTransferStatsByCardAmountCache(store),
		TransferStatsByCardStatusCache: NewTransferStatsByCardStatusCache(store),
	}
}
