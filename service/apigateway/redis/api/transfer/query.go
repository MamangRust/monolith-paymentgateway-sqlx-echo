package transfer_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transferQueryCache struct {
	store *cache.CacheStore
}

func NewTransferQueryCache(store *cache.CacheStore) TransferQueryCache {
	return &transferQueryCache{store: store}
}

func (c *transferQueryCache) GetCachedTransfersCache(ctx context.Context, req *requests.FindAllTransfers) (*response.ApiResponsePaginationTransfer, bool) {
	key := fmt.Sprintf(transferAllCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[response.ApiResponsePaginationTransfer](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *transferQueryCache) GetCachedTransferActiveCache(ctx context.Context, req *requests.FindAllTransfers) (*response.ApiResponsePaginationTransferDeleteAt, bool) {
	key := fmt.Sprintf(transferActiveCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[response.ApiResponsePaginationTransferDeleteAt](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (c *transferQueryCache) GetCachedTransferTrashedCache(ctx context.Context, req *requests.FindAllTransfers) (*response.ApiResponsePaginationTransferDeleteAt, bool) {
	key := fmt.Sprintf(transferTrashedCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[response.ApiResponsePaginationTransferDeleteAt](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (c *transferQueryCache) GetCachedTransferCache(ctx context.Context, id int) (*response.ApiResponseTransfer, bool) {
	key := fmt.Sprintf(transferByIdCacheKey, id)
	result, found := cache.GetFromCache[response.ApiResponseTransfer](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (c *transferQueryCache) GetCachedTransferByFrom(ctx context.Context, from string) (*response.ApiResponseTransfers, bool) {
	key := fmt.Sprintf(transferByFromCacheKey, from)
	result, found := cache.GetFromCache[response.ApiResponseTransfers](ctx, c.store, key)
	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *transferQueryCache) GetCachedTransferByTo(ctx context.Context, to string) (*response.ApiResponseTransfers, bool) {
	key := fmt.Sprintf(transferByToCacheKey, to)
	result, found := cache.GetFromCache[response.ApiResponseTransfers](ctx, c.store, key)
	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *transferQueryCache) SetCachedTransfersCache(ctx context.Context, req *requests.FindAllTransfers, data *response.ApiResponsePaginationTransfer) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transferAllCacheKey, req.Page, req.PageSize, req.Search)

	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *transferQueryCache) SetCachedTransferActiveCache(ctx context.Context, req *requests.FindAllTransfers, data *response.ApiResponsePaginationTransferDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transferActiveCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *transferQueryCache) SetCachedTransferTrashedCache(ctx context.Context, req *requests.FindAllTransfers, data *response.ApiResponsePaginationTransferDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transferTrashedCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *transferQueryCache) SetCachedTransferCache(ctx context.Context, data *response.ApiResponseTransfer) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(transferByIdCacheKey, data.Data.ID)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *transferQueryCache) SetCachedTransferByFrom(ctx context.Context, from string, data *response.ApiResponseTransfers) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transferByFromCacheKey, from)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *transferQueryCache) SetCachedTransferByTo(ctx context.Context, to string, data *response.ApiResponseTransfers) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transferByToCacheKey, to)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}
