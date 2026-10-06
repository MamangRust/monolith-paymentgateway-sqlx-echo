package mencache

import (
	"context"
	"fmt"

	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

// userCommandCache is a struct that represents the user command cache.
type userCommandCache struct {
	store *sharedcachehelpers.CacheStore
}

// NewUserCommandCache creates a new instance of userCommandCache using the provided cache store.
// It returns a pointer to the newly initialized userCommandCache.
func NewUserCommandCache(store *sharedcachehelpers.CacheStore) UserCommandCache {
	return &userCommandCache{store: store}
}

// DeleteUserCache removes the cache entry for a user associated with the given ID.
// It constructs the cache key using the user ID and deletes the entry from the cache store.
func (u *userCommandCache) DeleteUserCache(ctx context.Context, id int) {
	key := fmt.Sprintf(userByIdCacheKey, id)

	sharedcachehelpers.DeleteFromCache(ctx, u.store, key)
}

func (u *userCommandCache) DeleteUserListCache(ctx context.Context) {
	u.store.InvalidateCache(ctx, "user:all:*")
	u.store.InvalidateCache(ctx, "user:active:*")
	u.store.InvalidateCache(ctx, "user:trashed:*")
}
