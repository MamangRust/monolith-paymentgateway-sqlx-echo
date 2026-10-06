package withdraw_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type withdrawCommandCache struct {
	store *cache.CacheStore
}

func NewWithdrawCommandCache(store *cache.CacheStore) WithdrawCommandCache {
	return &withdrawCommandCache{store: store}
}

func (wc *withdrawCommandCache) DeleteCachedWithdrawCache(ctx context.Context, id int) {
	key := fmt.Sprintf(withdrawByIdCacheKey, id)
	cache.DeleteFromCache(ctx, wc.store, key)
}
