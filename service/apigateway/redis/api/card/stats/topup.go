package card_stats_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type cardStatsTopupCache struct {
	store *cache.CacheStore
}

func NewCardStatsTopupCache(store *cache.CacheStore) CardStatsTopupCache {
	return &cardStatsTopupCache{store: store}
}

func (c *cardStatsTopupCache) GetMonthlyTopupCache(ctx context.Context, year int) (*response.ApiResponseMonthlyAmount, bool) {
	key := fmt.Sprintf(cacheKeyMonthlyTopupAmount, year)
	result, found := cache.GetFromCache[response.ApiResponseMonthlyAmount](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *cardStatsTopupCache) SetMonthlyTopupCache(ctx context.Context, year int, data *response.ApiResponseMonthlyAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(cacheKeyMonthlyTopupAmount, year)
	cache.SetToCache(ctx, c.store, key, data, ttlStatistic)
}

func (c *cardStatsTopupCache) GetYearlyTopupCache(ctx context.Context, year int) (*response.ApiResponseYearlyAmount, bool) {
	key := fmt.Sprintf(cacheKeyYearlyTopupAmount, year)
	result, found := cache.GetFromCache[response.ApiResponseYearlyAmount](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *cardStatsTopupCache) SetYearlyTopupCache(ctx context.Context, year int, data *response.ApiResponseYearlyAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(cacheKeyYearlyTopupAmount, year)
	cache.SetToCache(ctx, c.store, key, data, ttlStatistic)
}
