package mencache

import (
	"context"
	"fmt"

	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type cardCommandCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewCardCommandCache(store *sharedcachehelpers.CacheStore) CardCommandCache {
	return &cardCommandCache{store: store}
}

func (c *cardCommandCache) DeleteCardCommandCache(ctx context.Context, id int) {
	key := fmt.Sprintf(cardByIdCacheKey, id)
	sharedcachehelpers.DeleteFromCache(ctx, c.store, key)
}

func (c *cardCommandCache) DeleteCardCache(ctx context.Context, id, userID int, cardNumber string) {
	c.DeleteCardCommandCache(ctx, id)

	userKey := fmt.Sprintf(cardByUserIdCacheKey, userID)
	sharedcachehelpers.DeleteFromCache(ctx, c.store, userKey)

	cardNumberKey := fmt.Sprintf(cardByCardNumCacheKey, cardNumber)
	sharedcachehelpers.DeleteFromCache(ctx, c.store, cardNumberKey)

	userCardKey := fmt.Sprintf(cardUserByCardNumCacheKey, cardNumber)
	sharedcachehelpers.DeleteFromCache(ctx, c.store, userCardKey)
}
