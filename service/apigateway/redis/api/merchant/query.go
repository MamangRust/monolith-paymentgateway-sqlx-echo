package merchant_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type merchantQueryCache struct {
	store *cache.CacheStore
}

func NewMerchantQueryCache(store *cache.CacheStore) MerchantQueryCache {
	return &merchantQueryCache{store: store}
}

func (m *merchantQueryCache) GetCachedMerchants(ctx context.Context, req *requests.FindAllMerchants) (*response.ApiResponsePaginationMerchant, bool) {
	key := fmt.Sprintf(merchantAllCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationMerchant](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantQueryCache) GetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchants) (*response.ApiResponsePaginationMerchantDeleteAt, bool) {
	key := fmt.Sprintf(merchantActiveCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationMerchantDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantQueryCache) GetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchants) (*response.ApiResponsePaginationMerchantDeleteAt, bool) {
	key := fmt.Sprintf(merchantTrashedCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationMerchantDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantQueryCache) GetCachedMerchant(ctx context.Context, id int) (*response.ApiResponseMerchant, bool) {
	key := fmt.Sprintf(merchantByIdCacheKey, id)
	result, found := cache.GetFromCache[response.ApiResponseMerchant](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantQueryCache) GetCachedMerchantsByUserId(ctx context.Context, userId int) (*response.ApiResponsesMerchant, bool) {
	key := fmt.Sprintf(merchantByUserIdCacheKey, userId)
	result, found := cache.GetFromCache[response.ApiResponsesMerchant](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantQueryCache) GetCachedMerchantByApiKey(ctx context.Context, apiKey string) (*response.ApiResponseMerchant, bool) {
	key := fmt.Sprintf(merchantByApiKeyCacheKey, apiKey)
	result, found := cache.GetFromCache[response.ApiResponseMerchant](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantQueryCache) SetCachedMerchants(ctx context.Context, req *requests.FindAllMerchants, data *response.ApiResponsePaginationMerchant) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantAllCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantQueryCache) SetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchants, data *response.ApiResponsePaginationMerchantDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantActiveCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantQueryCache) SetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchants, data *response.ApiResponsePaginationMerchantDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantTrashedCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantQueryCache) SetCachedMerchant(ctx context.Context, data *response.ApiResponseMerchant) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantByIdCacheKey, data.Data.ID)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantQueryCache) SetCachedMerchantsByUserId(ctx context.Context, userId int, data *response.ApiResponsesMerchant) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantByUserIdCacheKey, userId)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantQueryCache) SetCachedMerchantByApiKey(ctx context.Context, apiKey string, data *response.ApiResponseMerchant) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantByApiKeyCacheKey, apiKey)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}
