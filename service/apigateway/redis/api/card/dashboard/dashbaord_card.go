package card_dashboard_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type cardDashboardByCardNumberCache struct {
	store *cache.CacheStore
}

func NewCardDashboardByCardNumberCache(store *cache.CacheStore) CardDashboardByCardNumberCache {
	return &cardDashboardByCardNumberCache{store: store}
}

func (c *cardDashboardByCardNumberCache) GetDashboardCardCardNumberCache(ctx context.Context, cardNumber string) (*response.ApiResponseDashboardCardNumber, bool) {
	key := fmt.Sprintf(cacheKeyDashboardCardNumber, cardNumber)
	result, found := cache.GetFromCache[response.ApiResponseDashboardCardNumber](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (c *cardDashboardByCardNumberCache) SetDashboardCardCardNumberCache(ctx context.Context, cardNumber string, data *response.ApiResponseDashboardCardNumber) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(cacheKeyDashboardCardNumber, cardNumber)
	cache.SetToCache(ctx, c.store, key, data, ttlDashboardDefault)
}

func (c *cardDashboardByCardNumberCache) DeleteDashboardCardCardNumberCache(ctx context.Context, cardNumber string) {
	key := fmt.Sprintf(cacheKeyDashboardCardNumber, cardNumber)
	cache.DeleteFromCache(ctx, c.store, key)
}
