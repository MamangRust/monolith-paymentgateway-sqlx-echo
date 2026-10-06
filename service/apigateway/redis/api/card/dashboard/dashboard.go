package card_dashboard_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type cardDashboardCache struct {
	store *cache.CacheStore
}

func NewCardDashboardCache(store *cache.CacheStore) CardDashboardTotalCache {
	return &cardDashboardCache{store: store}
}

func (c *cardDashboardCache) GetDashboardCardCache(ctx context.Context) (*response.ApiResponseDashboardCard, bool) {
	result, found := cache.GetFromCache[response.ApiResponseDashboardCard](ctx, c.store, cacheKeyDashboardDefault)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (c *cardDashboardCache) SetDashboardCardCache(ctx context.Context, data *response.ApiResponseDashboardCard) {
	if data == nil {
		return
	}

	cache.SetToCache(ctx, c.store, cacheKeyDashboardDefault, data, ttlDashboardDefault)
}

func (c *cardDashboardCache) DeleteDashboardCardCache(ctx context.Context) {
	cache.DeleteFromCache(ctx, c.store, cacheKeyDashboardDefault)
}
