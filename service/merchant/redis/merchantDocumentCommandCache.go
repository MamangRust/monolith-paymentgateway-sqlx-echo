package mencache

import (
	"context"
	"fmt"

	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

// merchantDocumentCommandCache is a struct that represents the cache store
type merchantDocumentCommandCache struct {
	store *sharedcachehelpers.CacheStore
}

// NewMerchantDocumentCommandCache creates a new instance of merchantDocumentCommandCache
func NewMerchantDocumentCommandCache(store *sharedcachehelpers.CacheStore) MerchantDocumentCommandCache {
	return &merchantDocumentCommandCache{store: store}
}

// DeleteCachedMerchantDocuments deletes the cache entry associated with the specified merchant ID
func (s *merchantDocumentCommandCache) DeleteCachedMerchantDocuments(ctx context.Context, id int) {
	key := fmt.Sprintf(merchantDocumentByIdCacheKey, id)
	sharedcachehelpers.DeleteFromCache(ctx, s.store, key)
}
