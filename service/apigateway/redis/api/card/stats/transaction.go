package card_stats_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type cardStatsTransactionCache struct {
	store *cache.CacheStore
}

func NewCardStatsTransactionCache(store *cache.CacheStore) CardStatsTransactionCache {
	return &cardStatsTransactionCache{store: store}
}

func (c *cardStatsTransactionCache) GetMonthlyTransactionCache(ctx context.Context, year int) (*response.ApiResponseMonthlyAmount, bool) {
	key := fmt.Sprintf(cacheKeyMonthlyTransactionAmount, year)
	result, found := cache.GetFromCache[response.ApiResponseMonthlyAmount](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *cardStatsTransactionCache) SetMonthlyTransactionCache(ctx context.Context, year int, data *response.ApiResponseMonthlyAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(cacheKeyMonthlyTransactionAmount, year)
	cache.SetToCache(ctx, c.store, key, data, ttlStatistic)
}

func (c *cardStatsTransactionCache) GetYearlyTransactionCache(ctx context.Context, year int) (*response.ApiResponseYearlyAmount, bool) {
	key := fmt.Sprintf(cacheKeyYearlyTransactionAmount, year)
	result, found := cache.GetFromCache[response.ApiResponseYearlyAmount](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *cardStatsTransactionCache) SetYearlyTransactionCache(ctx context.Context, year int, data *response.ApiResponseYearlyAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(cacheKeyYearlyTransactionAmount, year)
	cache.SetToCache(ctx, c.store, key, data, ttlStatistic)
}
