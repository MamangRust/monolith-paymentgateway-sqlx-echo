package mencache

import (
	"context"
	"fmt"

	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

// transferCommandCache represents a cache store for transfer commands.
type transferCommandCache struct {
	store *sharedcachehelpers.CacheStore
}

// NewTransferCommandCache creates a new instance of transferCommandCache using the provided cache store.
func NewTransferCommandCache(store *sharedcachehelpers.CacheStore) TransferCommandCache {
	return &transferCommandCache{store: store}
}

// DeleteTransferCache removes a cached transfer by its ID.
//
// Parameters:
//   - ctx: The context for timeout and cancellation.
//   - id: The transfer ID.
func (t *transferCommandCache) DeleteTransferCache(ctx context.Context, id int) {
	sharedcachehelpers.DeleteFromCache(ctx, t.store, fmt.Sprintf(transferByIdCacheKey, id))
}
