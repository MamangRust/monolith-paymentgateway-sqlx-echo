package merchant_stats_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type merchantStatsAmountCache struct {
	store *cache.CacheStore
}

func NewMerchantStatsAmountCache(store *cache.CacheStore) MerchantStatsAmountCache {
	return &merchantStatsAmountCache{store: store}
}

func (s *merchantStatsAmountCache) GetMonthlyAmountMerchantCache(ctx context.Context, year int) (*response.ApiResponseMerchantMonthlyAmount, bool) {
	key := fmt.Sprintf(merchantMonthlyAmountCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseMerchantMonthlyAmount](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (s *merchantStatsAmountCache) SetMonthlyAmountMerchantCache(ctx context.Context, year int, data *response.ApiResponseMerchantMonthlyAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantMonthlyAmountCacheKey, year)
	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *merchantStatsAmountCache) GetYearlyAmountMerchantCache(ctx context.Context, year int) (*response.ApiResponseMerchantYearlyAmount, bool) {
	key := fmt.Sprintf(MerchantYearlyAmountCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseMerchantYearlyAmount](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (s *merchantStatsAmountCache) SetYearlyAmountMerchantCache(ctx context.Context, year int, data *response.ApiResponseMerchantYearlyAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(MerchantYearlyAmountCacheKey, year)
	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}
