package withdraw_stats_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type withdrawStatsAmountCache struct {
	store *cache.CacheStore
}

func NewWithdrawStatsAmountCache(store *cache.CacheStore) WithdrawStatsAmountCache {
	return &withdrawStatsAmountCache{store: store}
}

func (w *withdrawStatsAmountCache) GetCachedMonthlyWithdraws(ctx context.Context, year int) (*response.ApiResponseWithdrawMonthAmount, bool) {
	key := fmt.Sprintf(montWithdrawAmountKey, year)
	result, found := cache.GetFromCache[response.ApiResponseWithdrawMonthAmount](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawStatsAmountCache) SetCachedMonthlyWithdraws(ctx context.Context, year int, data *response.ApiResponseWithdrawMonthAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(montWithdrawAmountKey, year)
	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}

func (w *withdrawStatsAmountCache) GetCachedYearlyWithdraws(ctx context.Context, year int) (*response.ApiResponseWithdrawYearAmount, bool) {
	key := fmt.Sprintf(yearWithdrawAmountKey, year)
	result, found := cache.GetFromCache[response.ApiResponseWithdrawYearAmount](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawStatsAmountCache) SetCachedYearlyWithdraws(ctx context.Context, year int, data *response.ApiResponseWithdrawYearAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearWithdrawAmountKey, year)
	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}
