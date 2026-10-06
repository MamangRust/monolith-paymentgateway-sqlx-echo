package merchant_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type merchantCommandCache struct {
	store *cache.CacheStore
}

func NewMerchantCommandCache(store *cache.CacheStore) MerchantCommandCache {
	return &merchantCommandCache{store: store}
}

func (s *merchantCommandCache) DeleteCachedMerchant(ctx context.Context, id int) {
	key := fmt.Sprintf(merchantByIdCacheKey, id)

	cache.DeleteFromCache(ctx, s.store, key)
}
