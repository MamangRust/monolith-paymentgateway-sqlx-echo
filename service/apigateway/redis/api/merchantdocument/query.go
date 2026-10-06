package merchantdocument_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type merchantDocumentQueryCache struct {
	store *cache.CacheStore
}

func NewMerchantDocumentQueryCache(store *cache.CacheStore) MerchantDocumentQueryCache {
	return &merchantDocumentQueryCache{store: store}
}

func (m *merchantDocumentQueryCache) GetCachedMerchants(ctx context.Context, req *requests.FindAllMerchantDocuments) (*response.ApiResponsePaginationMerchantDocument, bool) {
	key := fmt.Sprintf(merchantAllCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationMerchantDocument](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantDocumentQueryCache) GetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchantDocuments) (*response.ApiResponsePaginationMerchantDocumentDeleteAt, bool) {
	key := fmt.Sprintf(merchantActiveCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationMerchantDocumentDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantDocumentQueryCache) GetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchantDocuments) (*response.ApiResponsePaginationMerchantDocumentDeleteAt, bool) {
	key := fmt.Sprintf(merchantTrashedCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationMerchantDocumentDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantDocumentQueryCache) GetCachedMerchant(ctx context.Context, id int) (*response.ApiResponseMerchantDocument, bool) {
	key := fmt.Sprintf(merchantByIdCacheKey, id)
	result, found := cache.GetFromCache[response.ApiResponseMerchantDocument](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantDocumentQueryCache) SetCachedMerchants(ctx context.Context, req *requests.FindAllMerchantDocuments, data *response.ApiResponsePaginationMerchantDocument) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantAllCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantDocumentQueryCache) SetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchantDocuments, data *response.ApiResponsePaginationMerchantDocumentDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantActiveCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantDocumentQueryCache) SetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchantDocuments, data *response.ApiResponsePaginationMerchantDocumentDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantTrashedCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantDocumentQueryCache) SetCachedMerchant(ctx context.Context, data *response.ApiResponseMerchantDocument) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantByIdCacheKey, data.Data.ID)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}
