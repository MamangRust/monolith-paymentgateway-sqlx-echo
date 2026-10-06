package mencache

import (
	"context"
	"fmt"

	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

// withdrawCommandCache is a struct that represents the withdraw command cache
type withdrawCommandCache struct {
	store *sharedcachehelpers.CacheStore
}

// NewWithdrawCommandCache creates a new instance of withdrawCommandCache using the provided cache store
// and returns a pointer to the newly initialized withdrawCommandCache
func NewWithdrawCommandCache(store *sharedcachehelpers.CacheStore) WithdrawCommandCache {
	return &withdrawCommandCache{store: store}
}

// DeleteCachedWithdrawCache deletes a cached withdraw entry by ID.
//
// Parameters:
//   - ctx: The context for timeout and cancellation.
//   - id: The ID of the withdraw cache to delete.
func (wc *withdrawCommandCache) DeleteCachedWithdrawCache(ctx context.Context, id int) {
	key := fmt.Sprintf(withdrawByIdCacheKey, id)
	sharedcachehelpers.DeleteFromCache(ctx, wc.store, key)
}
