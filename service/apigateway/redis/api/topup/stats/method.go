package topup_stats_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type topupStatsMethodCache struct {
	store *cache.CacheStore
}

func NewTopupStatsMethodCache(store *cache.CacheStore) TopupStatsMethodCache {
	return &topupStatsMethodCache{store: store}
}

func (c *topupStatsMethodCache) GetMonthlyTopupMethodsCache(ctx context.Context, year int) (*response.ApiResponseTopupMonthMethod, bool) {
	key := fmt.Sprintf(monthTopupMethodCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseTopupMonthMethod](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *topupStatsMethodCache) SetMonthlyTopupMethodsCache(ctx context.Context, year int, data *response.ApiResponseTopupMonthMethod) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(monthTopupMethodCacheKey, year)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *topupStatsMethodCache) GetYearlyTopupMethodsCache(ctx context.Context, year int) (*response.ApiResponseTopupYearMethod, bool) {
	key := fmt.Sprintf(yearTopupMethodCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseTopupYearMethod](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *topupStatsMethodCache) SetYearlyTopupMethodsCache(ctx context.Context, year int, data *response.ApiResponseTopupYearMethod) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearTopupMethodCacheKey, year)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}
