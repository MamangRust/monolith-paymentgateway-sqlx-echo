package mencache

import (
	"context"
	"fmt"

	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

// topupCommandCache is a struct that represents a cache for topups.
type topupCommandCache struct {
	store *sharedcachehelpers.CacheStore
}

// NewTopupCommandCache creates a new instance of topupCommandCache with the provided CacheStore.
// It returns a pointer to the newly created topupCommandCache.
func NewTopupCommandCache(store *sharedcachehelpers.CacheStore) TopupCommandCache {
	return &topupCommandCache{store: store}
}

// DeleteCachedTopupCache removes the cache entry associated with the specified topup ID.
// It formats the cache key using the topup ID and deletes the entry from the cache store.
func (c *topupCommandCache) DeleteCachedTopupCache(ctx context.Context, id int) {
	key := fmt.Sprintf(topupByIdCacheKey, id)
	sharedcachehelpers.DeleteFromCache(ctx, c.store, key)
}
