package merchant_stats_byapikey_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type merchantStatsAmountByApiKeyCache struct {
	store *cache.CacheStore
}

func NewMerchantStatsAmountByApiKeyCache(store *cache.CacheStore) MerchantStatsAmountByApiKeyCache {
	return &merchantStatsAmountByApiKeyCache{store: store}
}

func (m *merchantStatsAmountByApiKeyCache) GetMonthlyAmountByApikeysCache(ctx context.Context, req *requests.MonthYearAmountApiKey) (*response.ApiResponseMerchantMonthlyAmount, bool) {
	key := fmt.Sprintf(merchantMonthlyAmountByApikeyCacheKey, req.Apikey, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseMerchantMonthlyAmount](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantStatsAmountByApiKeyCache) SetMonthlyAmountByApikeysCache(ctx context.Context, req *requests.MonthYearAmountApiKey, data *response.ApiResponseMerchantMonthlyAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantMonthlyAmountByApikeyCacheKey, req.Apikey, req.Year)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantStatsAmountByApiKeyCache) GetYearlyAmountByApikeysCache(ctx context.Context, req *requests.MonthYearAmountApiKey) (*response.ApiResponseMerchantYearlyAmount, bool) {
	key := fmt.Sprintf(merchantYearlyAmountByApikeyCacheKey, req.Apikey, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseMerchantYearlyAmount](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantStatsAmountByApiKeyCache) SetYearlyAmountByApikeysCache(ctx context.Context, req *requests.MonthYearAmountApiKey, data *response.ApiResponseMerchantYearlyAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantYearlyAmountByApikeyCacheKey, req.Apikey, req.Year)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}
