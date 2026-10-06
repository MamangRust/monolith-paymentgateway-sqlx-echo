package topup_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type topupQueryCache struct {
	store *cache.CacheStore
}

func NewTopupQueryCache(store *cache.CacheStore) TopupQueryCache {
	return &topupQueryCache{store: store}
}

func (c *topupQueryCache) GetCachedTopupsCache(ctx context.Context, req *requests.FindAllTopups) (*response.ApiResponsePaginationTopup, bool) {
	key := fmt.Sprintf(topupAllCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationTopup](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *topupQueryCache) GetCacheTopupByCardCache(ctx context.Context, req *requests.FindAllTopupsByCardNumber) (*response.ApiResponsePaginationTopup, bool) {
	key := fmt.Sprintf(topupByCardCacheKey, req.CardNumber, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationTopup](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *topupQueryCache) GetCachedTopupActiveCache(ctx context.Context, req *requests.FindAllTopups) (*response.ApiResponsePaginationTopupDeleteAt, bool) {
	key := fmt.Sprintf(topupActiveCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationTopupDeleteAt](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *topupQueryCache) GetCachedTopupTrashedCache(ctx context.Context, req *requests.FindAllTopups) (*response.ApiResponsePaginationTopupDeleteAt, bool) {
	key := fmt.Sprintf(topupTrashedCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationTopupDeleteAt](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *topupQueryCache) GetCachedTopupCache(ctx context.Context, id int) (*response.ApiResponseTopup, bool) {
	key := fmt.Sprintf(topupByIdCacheKey, id)
	result, found := cache.GetFromCache[response.ApiResponseTopup](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *topupQueryCache) SetCachedTopupsCache(ctx context.Context, req *requests.FindAllTopups, data *response.ApiResponsePaginationTopup) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(topupAllCacheKey, req.Page, req.PageSize, req.Search)

	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *topupQueryCache) SetCacheTopupByCardCache(ctx context.Context, req *requests.FindAllTopupsByCardNumber, data *response.ApiResponsePaginationTopup) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(topupByCardCacheKey, req.CardNumber, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *topupQueryCache) SetCachedTopupActiveCache(ctx context.Context, req *requests.FindAllTopups, data *response.ApiResponsePaginationTopupDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(topupActiveCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *topupQueryCache) SetCachedTopupTrashedCache(ctx context.Context, req *requests.FindAllTopups, data *response.ApiResponsePaginationTopupDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(topupTrashedCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *topupQueryCache) SetCachedTopupCache(ctx context.Context, data *response.ApiResponseTopup) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(topupByIdCacheKey, data.Data.ID)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}
