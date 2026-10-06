package topup_stats_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type topupStatsAmountCache struct {
	store *cache.CacheStore
}

func NewTopupStatsAmountCache(store *cache.CacheStore) TopupStatsAmountCache {
	return &topupStatsAmountCache{store: store}
}

func (c *topupStatsAmountCache) GetMonthlyTopupAmountsCache(ctx context.Context, year int) (*response.ApiResponseTopupMonthAmount, bool) {
	key := fmt.Sprintf(monthTopupAmountCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseTopupMonthAmount](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *topupStatsAmountCache) SetMonthlyTopupAmountsCache(ctx context.Context, year int, data *response.ApiResponseTopupMonthAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(monthTopupAmountCacheKey, year)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *topupStatsAmountCache) GetYearlyTopupAmountsCache(ctx context.Context, year int) (*response.ApiResponseTopupYearAmount, bool) {
	key := fmt.Sprintf(yearTopupAmountCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseTopupYearAmount](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *topupStatsAmountCache) SetYearlyTopupAmountsCache(ctx context.Context, year int, data *response.ApiResponseTopupYearAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearTopupAmountCacheKey, year)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}
