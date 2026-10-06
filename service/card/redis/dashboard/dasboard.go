package carddashboardmencache

import (
	"context"

	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type cardDashboardCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewCardDashboardCache(store *sharedcachehelpers.CacheStore) CardDashboardTotalCache {
	return &cardDashboardCache{store: store}
}

func (c *cardDashboardCache) GetDashboardCardCache(ctx context.Context) (*response.DashboardCard, bool) {
	result, found := sharedcachehelpers.GetFromCache[*response.DashboardCard](ctx, c.store, cacheKeyDashboardDefault)

	if !found || result == nil {
		return nil, false
	}

	return *result, true

}

func (c *cardDashboardCache) SetDashboardCardCache(ctx context.Context, data *response.DashboardCard) {
	if data == nil {
		return
	}

	sharedcachehelpers.SetToCache(ctx, c.store, cacheKeyDashboardDefault, data, ttlDashboardDefault)
}

func (c *cardDashboardCache) DeleteDashboardCardCache(ctx context.Context) {
	sharedcachehelpers.DeleteFromCache(ctx, c.store, cacheKeyDashboardDefault)
}
