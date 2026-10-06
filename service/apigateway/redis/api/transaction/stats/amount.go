package transaction_stats_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transactionStatsAmountCache struct {
	store *cache.CacheStore
}

func NewTransactionStatsAmountCache(store *cache.CacheStore) TransactionStatsAmountCache {
	return &transactionStatsAmountCache{store: store}
}

func (t *transactionStatsAmountCache) GetMonthlyAmountsCache(ctx context.Context, year int) (*response.ApiResponseTransactionMonthAmount, bool) {
	key := fmt.Sprintf(monthTransactionAmountCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseTransactionMonthAmount](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionStatsAmountCache) SetMonthlyAmountsCache(ctx context.Context, year int, data *response.ApiResponseTransactionMonthAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(monthTransactionAmountCacheKey, year)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionStatsAmountCache) GetYearlyAmountsCache(ctx context.Context, year int) (*response.ApiResponseTransactionYearAmount, bool) {
	key := fmt.Sprintf(yearTransactionAmountCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseTransactionYearAmount](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionStatsAmountCache) SetYearlyAmountsCache(ctx context.Context, year int, data *response.ApiResponseTransactionYearAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearTransactionAmountCacheKey, year)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}
