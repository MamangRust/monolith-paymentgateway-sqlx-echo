package withdraw_stats_bycard_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type withdrawStatsByCardStatusCache struct {
	store *cache.CacheStore
}

func NewWithdrawStatsByCardStatusCache(store *cache.CacheStore) WithdrawStatsByCardStatusCache {
	return &withdrawStatsByCardStatusCache{store: store}
}

func (w *withdrawStatsByCardStatusCache) GetCachedMonthWithdrawStatusSuccessByCardNumber(ctx context.Context, req *requests.MonthStatusWithdrawCardNumber) (*response.ApiResponseWithdrawMonthStatusSuccess, bool) {
	key := fmt.Sprintf(monthWithdrawStatusSuccessByCardKey, req.CardNumber, req.Month, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseWithdrawMonthStatusSuccess](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawStatsByCardStatusCache) SetCachedMonthWithdrawStatusSuccessByCardNumber(ctx context.Context, req *requests.MonthStatusWithdrawCardNumber, data *response.ApiResponseWithdrawMonthStatusSuccess) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(monthWithdrawStatusSuccessByCardKey, req.CardNumber, req.Month, req.Year)
	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}

func (w *withdrawStatsByCardStatusCache) GetCachedYearlyWithdrawStatusSuccessByCardNumber(ctx context.Context, req *requests.YearStatusWithdrawCardNumber) (*response.ApiResponseWithdrawYearStatusSuccess, bool) {
	key := fmt.Sprintf(yearWithdrawStatusSuccessByCardKey, req.CardNumber, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseWithdrawYearStatusSuccess](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawStatsByCardStatusCache) SetCachedYearlyWithdrawStatusSuccessByCardNumber(ctx context.Context, req *requests.YearStatusWithdrawCardNumber, data *response.ApiResponseWithdrawYearStatusSuccess) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearWithdrawStatusSuccessByCardKey, req.CardNumber, req.Year)
	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}

func (w *withdrawStatsByCardStatusCache) GetCachedMonthWithdrawStatusFailedByCardNumber(ctx context.Context, req *requests.MonthStatusWithdrawCardNumber) (*response.ApiResponseWithdrawMonthStatusFailed, bool) {
	key := fmt.Sprintf(monthWithdrawStatusFailedByCardKey, req.CardNumber, req.Month, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseWithdrawMonthStatusFailed](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawStatsByCardStatusCache) SetCachedMonthWithdrawStatusFailedByCardNumber(ctx context.Context, req *requests.MonthStatusWithdrawCardNumber, data *response.ApiResponseWithdrawMonthStatusFailed) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(monthWithdrawStatusFailedByCardKey, req.CardNumber, req.Month, req.Year)
	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}

func (w *withdrawStatsByCardStatusCache) GetCachedYearlyWithdrawStatusFailedByCardNumber(ctx context.Context, req *requests.YearStatusWithdrawCardNumber) (*response.ApiResponseWithdrawYearStatusFailed, bool) {
	key := fmt.Sprintf(yearWithdrawStatusFailedByCardKey, req.CardNumber, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseWithdrawYearStatusFailed](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawStatsByCardStatusCache) SetCachedYearlyWithdrawStatusFailedByCardNumber(ctx context.Context, req *requests.YearStatusWithdrawCardNumber, data *response.ApiResponseWithdrawYearStatusFailed) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearWithdrawStatusFailedByCardKey, req.CardNumber, req.Year)
	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}
