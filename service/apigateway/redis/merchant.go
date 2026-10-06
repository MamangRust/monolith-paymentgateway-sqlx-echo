package mencache

import (
	"context"
	"fmt"

	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type merchantCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewMerchantCache(store *sharedcachehelpers.CacheStore) MerchantCache {
	return &merchantCache{store: store}
}

func (c *merchantCache) GetMerchantCache(ctx context.Context, apiKey string) (string, bool) {
	key := fmt.Sprintf(cacheMerchantKey, apiKey)

	result, found := sharedcachehelpers.GetFromCache[string](ctx, c.store, key)
	if !found || result == nil {
		return "", false
	}

	return *result, true
}

func (c *merchantCache) SetMerchantCache(ctx context.Context, merchantID string, apiKey string) {
	if merchantID == "" || apiKey == "" {
		return
	}

	key := fmt.Sprintf(cacheMerchantKey, apiKey)

	sharedcachehelpers.SetToCache(ctx, c.store, key, &merchantID, ttlDefault)
}
