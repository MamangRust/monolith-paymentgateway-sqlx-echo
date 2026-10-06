package carddashboardmencache

import (
	"context"
	"fmt"

	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type cardDashboardByCardNumberCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewCardDashboardByCardNumberCache(store *sharedcachehelpers.CacheStore) CardDashboardByCardNumberCache {
	return &cardDashboardByCardNumberCache{store: store}
}

func (c *cardDashboardByCardNumberCache) GetDashboardCardCardNumberCache(ctx context.Context, cardNumber string) (*response.DashboardCardCardNumber, bool) {
	key := fmt.Sprintf(cacheKeyDashboardCardNumber, cardNumber)
	result, found := sharedcachehelpers.GetFromCache[response.DashboardCardCardNumber](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (c *cardDashboardByCardNumberCache) SetDashboardCardCardNumberCache(ctx context.Context, cardNumber string, data *response.DashboardCardCardNumber) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(cacheKeyDashboardCardNumber, cardNumber)
	sharedcachehelpers.SetToCache(ctx, c.store, key, data, ttlDashboardDefault)
}

func (c *cardDashboardByCardNumberCache) DeleteDashboardCardCardNumberCache(ctx context.Context, cardNumber string) {
	key := fmt.Sprintf(cacheKeyDashboardCardNumber, cardNumber)
	sharedcachehelpers.DeleteFromCache(ctx, c.store, key)
}
