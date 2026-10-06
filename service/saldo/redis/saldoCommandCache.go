package mencache

import (
	"context"
	"fmt"

	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"go.uber.org/zap"
)

// saldoCommandCache is a struct that represents the cache store
type saldoCommandCache struct {
	store *sharedcachehelpers.CacheStore
}

// NewSaldoCommandCache creates a new instance of saldoCommandCache with the provided CacheStore.
// It returns a pointer to the newly created saldoCommandCache.
func NewSaldoCommandCache(store *sharedcachehelpers.CacheStore) SaldoCommandCache {
	return &saldoCommandCache{store: store}
}

// DeleteSaldoCache removes the cache entry associated with the specified saldo ID.
// It formats the cache key using the saldo ID and deletes the entry from the cache store.
func (s *saldoCommandCache) DeleteSaldoCache(ctx context.Context, saldo_id int) {
	key := fmt.Sprintf(saldoByIdCacheKey, saldo_id)
	sharedcachehelpers.DeleteFromCache(ctx, s.store, key)
}

// InvalidateSaldoCache removes saldo detail, list, and card-number cache entries
// after a mutation so upserts cannot serve stale balances.
func (s *saldoCommandCache) InvalidateSaldoCache(ctx context.Context) {
	if _, err := s.store.InvalidateCache(ctx, "saldo:*"); err != nil {
		s.store.Logger.Error("Failed to invalidate saldo cache", zap.Error(err))
	}
}
