package topup_stats_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type topupStatsStatusCache struct {
	store *cache.CacheStore
}

func NewTopupStatsStatusCache(store *cache.CacheStore) TopupStatsStatusCache {
	return &topupStatsStatusCache{store: store}
}

func (c *topupStatsStatusCache) GetMonthTopupStatusSuccessCache(ctx context.Context, req *requests.MonthTopupStatus) (*response.ApiResponseTopupMonthStatusSuccess, bool) {
	key := fmt.Sprintf(monthTopupStatusSuccessCacheKey, req.Month, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseTopupMonthStatusSuccess](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *topupStatsStatusCache) SetMonthTopupStatusSuccessCache(ctx context.Context, req *requests.MonthTopupStatus, data *response.ApiResponseTopupMonthStatusSuccess) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(monthTopupStatusSuccessCacheKey, req.Month, req.Year)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *topupStatsStatusCache) GetYearlyTopupStatusSuccessCache(ctx context.Context, year int) (*response.ApiResponseTopupYearStatusSuccess, bool) {
	key := fmt.Sprintf(yearTopupStatusSuccessCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseTopupYearStatusSuccess](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *topupStatsStatusCache) SetYearlyTopupStatusSuccessCache(ctx context.Context, year int, data *response.ApiResponseTopupYearStatusSuccess) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearTopupStatusSuccessCacheKey, year)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *topupStatsStatusCache) GetMonthTopupStatusFailedCache(ctx context.Context, req *requests.MonthTopupStatus) (*response.ApiResponseTopupMonthStatusFailed, bool) {
	key := fmt.Sprintf(monthTopupStatusFailedCacheKey, req.Month, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseTopupMonthStatusFailed](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *topupStatsStatusCache) SetMonthTopupStatusFailedCache(ctx context.Context, req *requests.MonthTopupStatus, data *response.ApiResponseTopupMonthStatusFailed) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(monthTopupStatusFailedCacheKey, req.Month, req.Year)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *topupStatsStatusCache) GetYearlyTopupStatusFailedCache(ctx context.Context, year int) (*response.ApiResponseTopupYearStatusFailed, bool) {
	key := fmt.Sprintf(yearTopupStatusFailedCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseTopupYearStatusFailed](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *topupStatsStatusCache) SetYearlyTopupStatusFailedCache(ctx context.Context, year int, data *response.ApiResponseTopupYearStatusFailed) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearTopupStatusFailedCacheKey, year)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}
