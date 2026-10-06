package transaction_stats_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transactionStatsStatusCache struct {
	store *cache.CacheStore
}

func NewTransactionStatsStatusCache(store *cache.CacheStore) TransactionStatsStatusCache {
	return &transactionStatsStatusCache{store: store}
}

func (t *transactionStatsStatusCache) GetMonthTransactionStatusSuccessCache(ctx context.Context, req *requests.MonthStatusTransaction) (*response.ApiResponseTransactionMonthStatusSuccess, bool) {
	key := fmt.Sprintf(monthTransactionStatusSuccessCacheKey, req.Month, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseTransactionMonthStatusSuccess](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionStatsStatusCache) SetMonthTransactionStatusSuccessCache(ctx context.Context, req *requests.MonthStatusTransaction, data *response.ApiResponseTransactionMonthStatusSuccess) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(monthTransactionStatusSuccessCacheKey, req.Month, req.Year)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionStatsStatusCache) GetYearTransactionStatusSuccessCache(ctx context.Context, year int) (*response.ApiResponseTransactionYearStatusSuccess, bool) {
	key := fmt.Sprintf(yearTransactionStatusSuccessCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseTransactionYearStatusSuccess](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionStatsStatusCache) SetYearTransactionStatusSuccessCache(ctx context.Context, year int, data *response.ApiResponseTransactionYearStatusSuccess) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearTransactionStatusSuccessCacheKey, year)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionStatsStatusCache) GetMonthTransactionStatusFailedCache(ctx context.Context, req *requests.MonthStatusTransaction) (*response.ApiResponseTransactionMonthStatusFailed, bool) {
	key := fmt.Sprintf(monthTransactionStatusFailedCacheKey, req.Month, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseTransactionMonthStatusFailed](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionStatsStatusCache) SetMonthTransactionStatusFailedCache(ctx context.Context, req *requests.MonthStatusTransaction, data *response.ApiResponseTransactionMonthStatusFailed) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(monthTransactionStatusFailedCacheKey, req.Month, req.Year)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionStatsStatusCache) GetYearTransactionStatusFailedCache(ctx context.Context, year int) (*response.ApiResponseTransactionYearStatusFailed, bool) {
	key := fmt.Sprintf(yearTransactionStatusFailedCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseTransactionYearStatusFailed](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionStatsStatusCache) SetYearTransactionStatusFailedCache(ctx context.Context, year int, data *response.ApiResponseTransactionYearStatusFailed) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearTransactionStatusFailedCacheKey, year)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}
