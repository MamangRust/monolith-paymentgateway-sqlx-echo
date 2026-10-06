package mencache

import (
	"context"
	"fmt"

	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"go.uber.org/zap"
)

// merchantCommandCache is a struct that represents the cache store
type merchantCommandCache struct {
	store *sharedcachehelpers.CacheStore
}

// NewMerchantCommandCache returns a new instance of merchantCommandCache
func NewMerchantCommandCache(store *sharedcachehelpers.CacheStore) MerchantCommandCache {
	return &merchantCommandCache{store: store}
}

// DeleteCachedMerchant removes the cache entry associated with the specified merchant ID.
// It formats the cache key using the merchant ID and deletes the entry from the cache store.
func (s *merchantCommandCache) DeleteCachedMerchant(ctx context.Context, id int) {
	if s == nil || s.store == nil || s.store.Redis == nil {
		return
	}

	key := fmt.Sprintf(merchantByIdCacheKey, id)
	sharedcachehelpers.DeleteFromCache(ctx, s.store, key)
}

func (s *merchantCommandCache) InvalidateMerchantListCaches(ctx context.Context) {
	if s == nil || s.store == nil || s.store.Redis == nil {
		return
	}

	for _, pattern := range []string{
		"merchant:all:*",
		"merchant:active:*",
		"merchant:trashed:*",
		"merchant:user_id:*",
		"merchant:api_key:*",
		"merchant:statistic:*",
		"merchant:transaction:*",
	} {
		if _, err := s.store.InvalidateCache(ctx, pattern); err != nil {
			s.store.Logger.Error("failed to invalidate merchant cache", zap.Error(err), zap.String("pattern", pattern))
		}
	}
}

func (s *merchantCommandCache) DeleteCachedMerchantByUserID(ctx context.Context, userID int) {
	if s == nil || s.store == nil || s.store.Redis == nil {
		return
	}
	sharedcachehelpers.DeleteFromCache(ctx, s.store, fmt.Sprintf(merchantByUserIdCacheKey, userID))
}

func (s *merchantCommandCache) DeleteCachedMerchantByAPIKey(ctx context.Context, apiKey string) {
	if s == nil || s.store == nil || s.store.Redis == nil {
		return
	}
	sharedcachehelpers.DeleteFromCache(ctx, s.store, fmt.Sprintf(merchantByApiKeyCacheKey, apiKey))
}
