package mencache

import (
	"context"
	"fmt"

	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type roleCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewRoleCache(store *sharedcachehelpers.CacheStore) RoleCache {
	return &roleCache{store: store}
}

func (c *roleCache) GetRoleCache(ctx context.Context, userID string) ([]string, bool) {
	key := fmt.Sprintf(cacheRoleKey, userID)

	result, found := sharedcachehelpers.GetFromCache[[]string](ctx, c.store, key)
	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (c *roleCache) SetRoleCache(ctx context.Context, userID string, roles []string) {
	if userID == "" || len(roles) == 0 {
		return
	}

	key := fmt.Sprintf(cacheRoleKey, userID)

	sharedcachehelpers.SetToCache(ctx, c.store, key, &roles, ttlDefault)
}
