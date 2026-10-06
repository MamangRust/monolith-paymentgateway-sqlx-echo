package merchant_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type merchantTransactionCache struct {
	store *cache.CacheStore
}

func NewMerchantTransactionCache(store *cache.CacheStore) MerchantTransactionCache {
	return &merchantTransactionCache{store: store}
}

func (m *merchantTransactionCache) GetCacheAllMerchantTransactions(ctx context.Context, req *requests.FindAllMerchantTransactions) (*response.ApiResponsePaginationMerchantTransaction, bool) {
	key := fmt.Sprintf(merchantTransactionsCacheKey, req.Search, req.Page, req.PageSize)
	result, found := cache.GetFromCache[response.ApiResponsePaginationMerchantTransaction](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantTransactionCache) GetCacheMerchantTransactions(ctx context.Context, req *requests.FindAllMerchantTransactionsById) (*response.ApiResponsePaginationMerchantTransaction, bool) {
	key := fmt.Sprintf(merchantTransactionCacheKey, req.MerchantID, req.Search, req.Page, req.PageSize)
	result, found := cache.GetFromCache[response.ApiResponsePaginationMerchantTransaction](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantTransactionCache) GetCacheMerchantTransactionApikey(ctx context.Context, req *requests.FindAllMerchantTransactionsByApiKey) (*response.ApiResponsePaginationMerchantTransaction, bool) {
	key := fmt.Sprintf(merchantTransactionApikeyCacheKey, req.ApiKey, req.Search, req.Page, req.PageSize)
	result, found := cache.GetFromCache[response.ApiResponsePaginationMerchantTransaction](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (m *merchantTransactionCache) SetCacheAllMerchantTransactions(ctx context.Context, req *requests.FindAllMerchantTransactions, data *response.ApiResponsePaginationMerchantTransaction) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantTransactionsCacheKey, req.Search, req.Page, req.PageSize)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantTransactionCache) SetCacheMerchantTransactions(ctx context.Context, req *requests.FindAllMerchantTransactionsById, data *response.ApiResponsePaginationMerchantTransaction) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantTransactionCacheKey, req.MerchantID, req.Search, req.Page, req.PageSize)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantTransactionCache) SetCacheMerchantTransactionApikey(ctx context.Context, req *requests.FindAllMerchantTransactionsByApiKey, data *response.ApiResponsePaginationMerchantTransaction) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(merchantTransactionApikeyCacheKey, req.ApiKey, req.Search, req.Page, req.PageSize)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}
