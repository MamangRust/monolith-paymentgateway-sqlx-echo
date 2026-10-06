package mencache

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type transactionCachedResponseAll struct {
	Data         []*db.GetTransactionsRow `json:"data"`
	TotalRecords *int                     `json:"total_records"`
}

type transactionCachedResponseByCard struct {
	Data         []*db.GetTransactionsByCardNumberRow `json:"data"`
	TotalRecords *int                                 `json:"total_records"`
}

type transactionCachedResponseActive struct {
	Data         []*db.GetActiveTransactionsRow `json:"data"`
	TotalRecords *int                           `json:"total_records"`
}

type transactionCachedResponseTrashed struct {
	Data         []*db.GetTrashedTransactionsRow `json:"data"`
	TotalRecords *int                            `json:"total_records"`
}

type transactionQueryCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewTransactionQueryCache(store *sharedcachehelpers.CacheStore) TransactionQueryCache {
	return &transactionQueryCache{store: store}
}

func (t *transactionQueryCache) GetCachedTransactionsCache(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetTransactionsRow, *int, bool) {
	key := fmt.Sprintf(transactionAllCacheKey, req.Page, req.PageSize, req.Search)

	result, found := sharedcachehelpers.GetFromCache[transactionCachedResponseAll](ctx, t.store, key)

	if !found || result == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (t *transactionQueryCache) SetCachedTransactionsCache(ctx context.Context, req *requests.FindAllTransactions, data []*db.GetTransactionsRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}
	if data == nil {
		data = []*db.GetTransactionsRow{}
	}

	key := fmt.Sprintf(transactionAllCacheKey, req.Page, req.PageSize, req.Search)
	payload := &transactionCachedResponseAll{Data: data, TotalRecords: total}
	sharedcachehelpers.SetToCache(ctx, t.store, key, payload, ttlDefault)
}

func (t *transactionQueryCache) GetCachedTransactionByCardNumberCache(ctx context.Context, req *requests.FindAllTransactionCardNumber) ([]*db.GetTransactionsByCardNumberRow, *int, bool) {
	key := fmt.Sprintf(transactionByCardCacheKey, req.CardNumber, req.Page, req.PageSize, req.Search)

	result, found := sharedcachehelpers.GetFromCache[transactionCachedResponseByCard](ctx, t.store, key)

	if !found || result == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (t *transactionQueryCache) SetCachedTransactionByCardNumberCache(ctx context.Context, req *requests.FindAllTransactionCardNumber, data []*db.GetTransactionsByCardNumberRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}
	if data == nil {
		data = []*db.GetTransactionsByCardNumberRow{}
	}

	key := fmt.Sprintf(transactionByCardCacheKey, req.CardNumber, req.Page, req.PageSize, req.Search)
	payload := &transactionCachedResponseByCard{Data: data, TotalRecords: total}
	sharedcachehelpers.SetToCache(ctx, t.store, key, payload, ttlDefault)
}

func (t *transactionQueryCache) GetCachedTransactionActiveCache(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetActiveTransactionsRow, *int, bool) {
	key := fmt.Sprintf(transactionActiveCacheKey, req.Page, req.PageSize, req.Search)

	result, found := sharedcachehelpers.GetFromCache[transactionCachedResponseActive](ctx, t.store, key)

	if !found || result == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (t *transactionQueryCache) SetCachedTransactionActiveCache(ctx context.Context, req *requests.FindAllTransactions, data []*db.GetActiveTransactionsRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}
	if data == nil {
		data = []*db.GetActiveTransactionsRow{}
	}

	key := fmt.Sprintf(transactionActiveCacheKey, req.Page, req.PageSize, req.Search)
	payload := &transactionCachedResponseActive{Data: data, TotalRecords: total}
	sharedcachehelpers.SetToCache(ctx, t.store, key, payload, ttlDefault)
}

func (t *transactionQueryCache) GetCachedTransactionTrashedCache(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetTrashedTransactionsRow, *int, bool) {
	key := fmt.Sprintf(transactionTrashedCacheKey, req.Page, req.PageSize, req.Search)

	result, found := sharedcachehelpers.GetFromCache[transactionCachedResponseTrashed](ctx, t.store, key)

	if !found || result == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (t *transactionQueryCache) SetCachedTransactionTrashedCache(ctx context.Context, req *requests.FindAllTransactions, data []*db.GetTrashedTransactionsRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}
	if data == nil {
		data = []*db.GetTrashedTransactionsRow{}
	}

	key := fmt.Sprintf(transactionTrashedCacheKey, req.Page, req.PageSize, req.Search)
	payload := &transactionCachedResponseTrashed{Data: data, TotalRecords: total}
	sharedcachehelpers.SetToCache(ctx, t.store, key, payload, ttlDefault)
}

func (t *transactionQueryCache) GetCachedTransactionCache(ctx context.Context, transactionId int) (*db.GetTransactionByIDRow, bool) {
	key := fmt.Sprintf(transactionByIdCacheKey, transactionId)
	result, found := sharedcachehelpers.GetFromCache[*db.GetTransactionByIDRow](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (t *transactionQueryCache) SetCachedTransactionCache(ctx context.Context, data *db.GetTransactionByIDRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(transactionByIdCacheKey, data.TransactionID)
	sharedcachehelpers.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionQueryCache) GetCachedTransactionByMerchantIdCache(ctx context.Context, merchantId int) ([]*db.GetTransactionsByMerchantIDRow, bool) {
	key := fmt.Sprintf(transactionByMerchantIdCacheKey, merchantId)
	result, found := sharedcachehelpers.GetFromCache[[]*db.GetTransactionsByMerchantIDRow](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (t *transactionQueryCache) SetCachedTransactionByMerchantIdCache(ctx context.Context, merchantId int, data []*db.GetTransactionsByMerchantIDRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(transactionByMerchantIdCacheKey, merchantId)
	sharedcachehelpers.SetToCache(ctx, t.store, key, &data, ttlDefault)
}
