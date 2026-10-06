package transfer_stats_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transferStatsStatusCache struct {
	store *cache.CacheStore
}

func NewTransferStatsStatusCache(store *cache.CacheStore) TransferStatsStatusCache {
	return &transferStatsStatusCache{store: store}
}

func (t *transferStatsStatusCache) GetCachedMonthTransferStatusSuccess(ctx context.Context, req *requests.MonthStatusTransfer) (*response.ApiResponseTransferMonthStatusSuccess, bool) {
	key := fmt.Sprintf(transferMonthTransferStatusSuccessKey, req.Month, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseTransferMonthStatusSuccess](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transferStatsStatusCache) SetCachedMonthTransferStatusSuccess(ctx context.Context, req *requests.MonthStatusTransfer, data *response.ApiResponseTransferMonthStatusSuccess) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transferMonthTransferStatusSuccessKey, req.Month, req.Year)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transferStatsStatusCache) GetCachedYearlyTransferStatusSuccess(ctx context.Context, year int) (*response.ApiResponseTransferYearStatusSuccess, bool) {
	key := fmt.Sprintf(transferYearTransferStatusSuccessKey, year)
	result, found := cache.GetFromCache[response.ApiResponseTransferYearStatusSuccess](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transferStatsStatusCache) SetCachedYearlyTransferStatusSuccess(ctx context.Context, year int, data *response.ApiResponseTransferYearStatusSuccess) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transferYearTransferStatusSuccessKey, year)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transferStatsStatusCache) GetCachedMonthTransferStatusFailed(ctx context.Context, req *requests.MonthStatusTransfer) (*response.ApiResponseTransferMonthStatusFailed, bool) {
	key := fmt.Sprintf(transferMonthTransferStatusFailedKey, req.Month, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseTransferMonthStatusFailed](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transferStatsStatusCache) SetCachedMonthTransferStatusFailed(ctx context.Context, req *requests.MonthStatusTransfer, data *response.ApiResponseTransferMonthStatusFailed) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transferMonthTransferStatusFailedKey, req.Month, req.Year)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transferStatsStatusCache) GetCachedYearlyTransferStatusFailed(ctx context.Context, year int) (*response.ApiResponseTransferYearStatusFailed, bool) {
	key := fmt.Sprintf(transferYearTransferStatusFailedKey, year)
	result, found := cache.GetFromCache[response.ApiResponseTransferYearStatusFailed](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transferStatsStatusCache) SetCachedYearlyTransferStatusFailed(ctx context.Context, year int, data *response.ApiResponseTransferYearStatusFailed) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transferYearTransferStatusFailedKey, year)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}
