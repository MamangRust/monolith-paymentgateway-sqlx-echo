package merchantdocument_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type merchantDocumentCommandCache struct {
	store *cache.CacheStore
}

func NewMerchantDocumentCommandCache(store *cache.CacheStore) MerchantDocumentCommandCache {
	return &merchantDocumentCommandCache{store: store}
}

func (s *merchantDocumentCommandCache) DeleteCachedMerchantDocument(ctx context.Context, id int) {
	key := fmt.Sprintf(merchantByIdCacheKey, id)

	cache.DeleteFromCache(ctx, s.store, key)
}
