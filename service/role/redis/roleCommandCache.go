package mencache

import (
	"context"
	"fmt"

	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type roleCommandCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewRoleCommandCache(store *sharedcachehelpers.CacheStore) *roleCommandCache {
	return &roleCommandCache{store: store}
}

func (s *roleCommandCache) DeleteCachedRole(ctx context.Context, id int) {
	key := fmt.Sprintf(roleByIdCacheKey, id)

	sharedcachehelpers.DeleteFromCache(ctx, s.store, key)
}
