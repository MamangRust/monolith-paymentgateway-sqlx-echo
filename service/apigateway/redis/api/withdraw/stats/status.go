package withdraw_stats_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type withdrawStatsStatusCache struct {
	store *cache.CacheStore
}

func NewWithdrawStatsStatusCache(store *cache.CacheStore) WithdrawStatsStatusCache {
	return &withdrawStatsStatusCache{store: store}
}

func (w *withdrawStatsStatusCache) GetCachedMonthWithdrawStatusSuccessCache(ctx context.Context, req *requests.MonthStatusWithdraw) (*response.ApiResponseWithdrawMonthStatusSuccess, bool) {
	key := fmt.Sprintf(montWithdrawStatusSuccessKey, req.Month, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseWithdrawMonthStatusSuccess](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawStatsStatusCache) SetCachedMonthWithdrawStatusSuccessCache(ctx context.Context, req *requests.MonthStatusWithdraw, data *response.ApiResponseWithdrawMonthStatusSuccess) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(montWithdrawStatusSuccessKey, req.Month, req.Year)
	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}

func (w *withdrawStatsStatusCache) GetCachedYearlyWithdrawStatusSuccessCache(ctx context.Context, year int) (*response.ApiResponseWithdrawYearStatusSuccess, bool) {
	key := fmt.Sprintf(yearWithdrawStatusSuccessKey, year)
	result, found := cache.GetFromCache[response.ApiResponseWithdrawYearStatusSuccess](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawStatsStatusCache) SetCachedYearlyWithdrawStatusSuccessCache(ctx context.Context, year int, data *response.ApiResponseWithdrawYearStatusSuccess) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearWithdrawStatusSuccessKey, year)
	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}

func (w *withdrawStatsStatusCache) GetCachedMonthWithdrawStatusFailedCache(ctx context.Context, req *requests.MonthStatusWithdraw) (*response.ApiResponseWithdrawMonthStatusFailed, bool) {
	key := fmt.Sprintf(montWithdrawStatusFailedKey, req.Month, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseWithdrawMonthStatusFailed](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawStatsStatusCache) SetCachedMonthWithdrawStatusFailedCache(ctx context.Context, req *requests.MonthStatusWithdraw, data *response.ApiResponseWithdrawMonthStatusFailed) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(montWithdrawStatusFailedKey, req.Month, req.Year)
	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}

func (w *withdrawStatsStatusCache) GetCachedYearlyWithdrawStatusFailedCache(ctx context.Context, year int) (*response.ApiResponseWithdrawYearStatusFailed, bool) {
	key := fmt.Sprintf(yearWithdrawStatusFailedKey, year)
	result, found := cache.GetFromCache[response.ApiResponseWithdrawYearStatusFailed](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawStatsStatusCache) SetCachedYearlyWithdrawStatusFailedCache(ctx context.Context, year int, data *response.ApiResponseWithdrawYearStatusFailed) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearWithdrawStatusFailedKey, year)
	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}
