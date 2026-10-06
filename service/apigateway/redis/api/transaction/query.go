package transaction_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transactionQueryCache struct {
	store *cache.CacheStore
}

func NewTransactionQueryCache(store *cache.CacheStore) TransactionQueryCache {
	return &transactionQueryCache{store: store}
}

func (t *transactionQueryCache) GetCachedTransactionsCache(ctx context.Context, req *requests.FindAllTransactions) (*response.ApiResponsePaginationTransaction, bool) {
	key := fmt.Sprintf(transactionAllCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationTransaction](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionQueryCache) GetCachedTransactionByCardNumberCache(ctx context.Context, req *requests.FindAllTransactionCardNumber) (*response.ApiResponsePaginationTransaction, bool) {
	key := fmt.Sprintf(transactionByCardCacheKey, req.CardNumber, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationTransaction](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionQueryCache) GetCachedTransactionActiveCache(ctx context.Context, req *requests.FindAllTransactions) (*response.ApiResponsePaginationTransactionDeleteAt, bool) {
	key := fmt.Sprintf(transactionActiveCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationTransactionDeleteAt](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionQueryCache) GetCachedTransactionTrashedCache(ctx context.Context, req *requests.FindAllTransactions) (*response.ApiResponsePaginationTransactionDeleteAt, bool) {
	key := fmt.Sprintf(transactionTrashedCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationTransactionDeleteAt](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionQueryCache) GetCachedTransactionByMerchantIdCache(ctx context.Context, merchantId int) (*response.ApiResponseTransactions, bool) {
	key := fmt.Sprintf(transactionByMerchantIdCacheKey, merchantId)
	result, found := cache.GetFromCache[response.ApiResponseTransactions](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionQueryCache) GetCachedTransactionCache(ctx context.Context, transactionId int) (*response.ApiResponseTransaction, bool) {
	key := fmt.Sprintf(transactionByIdCacheKey, transactionId)
	result, found := cache.GetFromCache[response.ApiResponseTransaction](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionQueryCache) SetCachedTransactionsCache(ctx context.Context, req *requests.FindAllTransactions, data *response.ApiResponsePaginationTransaction) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactionAllCacheKey, req.Page, req.PageSize, req.Search)

	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionQueryCache) SetCachedTransactionByCardNumberCache(ctx context.Context, req *requests.FindAllTransactionCardNumber, data *response.ApiResponsePaginationTransaction) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactionByCardCacheKey, req.CardNumber, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionQueryCache) SetCachedTransactionActiveCache(ctx context.Context, req *requests.FindAllTransactions, data *response.ApiResponsePaginationTransactionDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactionActiveCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionQueryCache) SetCachedTransactionTrashedCache(ctx context.Context, req *requests.FindAllTransactions, data *response.ApiResponsePaginationTransactionDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactionTrashedCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionQueryCache) SetCachedTransactionByMerchantIdCache(ctx context.Context, merchantId int, data *response.ApiResponseTransactions) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactionByMerchantIdCacheKey, merchantId)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionQueryCache) SetCachedTransactionCache(ctx context.Context, data *response.ApiResponseTransaction) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(transactionByIdCacheKey, data.Data.ID)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}
