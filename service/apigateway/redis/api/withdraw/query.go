package withdraw_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type withdrawQueryCache struct {
	store *cache.CacheStore
}

func NewWithdrawQueryCache(store *cache.CacheStore) WithdrawQueryCache {
	return &withdrawQueryCache{store: store}
}

func (w *withdrawQueryCache) GetCachedWithdrawsCache(ctx context.Context, req *requests.FindAllWithdraws) (*response.ApiResponsePaginationWithdraw, bool) {
	key := fmt.Sprintf(withdrawAllCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationWithdraw](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawQueryCache) GetCachedWithdrawByCardCache(ctx context.Context, req *requests.FindAllWithdrawCardNumber) (*response.ApiResponsePaginationWithdraw, bool) {
	key := fmt.Sprintf(withdrawByCardCacheKey, req.CardNumber, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationWithdraw](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawQueryCache) GetCachedWithdrawActiveCache(ctx context.Context, req *requests.FindAllWithdraws) (*response.ApiResponsePaginationWithdrawDeleteAt, bool) {
	key := fmt.Sprintf(withdrawActiveCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationWithdrawDeleteAt](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawQueryCache) GetCachedWithdrawTrashedCache(ctx context.Context, req *requests.FindAllWithdraws) (*response.ApiResponsePaginationWithdrawDeleteAt, bool) {
	key := fmt.Sprintf(withdrawTrashedCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationWithdrawDeleteAt](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawQueryCache) GetCachedWithdrawCache(ctx context.Context, id int) (*response.ApiResponseWithdraw, bool) {
	key := fmt.Sprintf(withdrawByIdCacheKey, id)
	result, found := cache.GetFromCache[response.ApiResponseWithdraw](ctx, w.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (w *withdrawQueryCache) SetCachedWithdrawsCache(ctx context.Context, req *requests.FindAllWithdraws, data *response.ApiResponsePaginationWithdraw) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(withdrawAllCacheKey, req.Page, req.PageSize, req.Search)

	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}

func (w *withdrawQueryCache) SetCachedWithdrawByCardCache(ctx context.Context, req *requests.FindAllWithdrawCardNumber, data *response.ApiResponsePaginationWithdraw) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(withdrawByCardCacheKey, req.CardNumber, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}

func (w *withdrawQueryCache) SetCachedWithdrawActiveCache(ctx context.Context, req *requests.FindAllWithdraws, data *response.ApiResponsePaginationWithdrawDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(withdrawActiveCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}

func (w *withdrawQueryCache) SetCachedWithdrawTrashedCache(ctx context.Context, req *requests.FindAllWithdraws, data *response.ApiResponsePaginationWithdrawDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(withdrawTrashedCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}

func (w *withdrawQueryCache) SetCachedWithdrawCache(ctx context.Context, data *response.ApiResponseWithdraw) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(withdrawByIdCacheKey, data.Data.ID)
	cache.SetToCache(ctx, w.store, key, data, ttlDefault)
}
