package card_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type cardQueryCache struct {
	store *cache.CacheStore
}

func NewCardQueryCache(store *cache.CacheStore) CardQueryCache {
	return &cardQueryCache{store: store}
}

func (c *cardQueryCache) GetByIdCache(ctx context.Context, cardID int) (*response.ApiResponseCard, bool) {
	key := fmt.Sprintf(cardByIdCacheKey, cardID)
	result, found := cache.GetFromCache[*response.ApiResponseCard](ctx, c.store, key)
	if !found || result == nil {
		return nil, false
	}
	return *result, true
}

func (c *cardQueryCache) GetFindAllCache(ctx context.Context, req *requests.FindAllCards) (*response.ApiResponsePaginationCard, bool) {
	key := fmt.Sprintf(cardAllCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationCard](ctx, c.store, key)
	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *cardQueryCache) GetByActiveCache(ctx context.Context, req *requests.FindAllCards) (*response.ApiResponsePaginationCardDeleteAt, bool) {
	key := fmt.Sprintf(cardActiveCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationCardDeleteAt](ctx, c.store, key)
	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *cardQueryCache) GetByTrashedCache(ctx context.Context, req *requests.FindAllCards) (*response.ApiResponsePaginationCardDeleteAt, bool) {
	key := fmt.Sprintf(cardTrashedCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationCardDeleteAt](ctx, c.store, key)
	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *cardQueryCache) GetByUserIDCache(ctx context.Context, userID int) (*response.ApiResponseCard, bool) {
	key := fmt.Sprintf(cardByUserIdCacheKey, userID)
	result, found := cache.GetFromCache[*response.ApiResponseCard](ctx, c.store, key)
	if !found || result == nil {
		return nil, false
	}
	return *result, true
}

func (c *cardQueryCache) GetByCardNumberCache(
	ctx context.Context,
	cardNumber string,
) (*response.ApiResponseCard, bool) {

	key := fmt.Sprintf(cardByCardNumCacheKey, cardNumber)

	result, found := cache.GetFromCache[response.ApiResponseCard](ctx, c.store, key)
	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (c *cardQueryCache) SetByIdCache(ctx context.Context, cardID int, data *response.ApiResponseCard) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(cardByIdCacheKey, cardID)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *cardQueryCache) SetByCardNumberCache(
	ctx context.Context,
	cardNumber string,
	data *response.ApiResponseCard,
) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(cardByCardNumCacheKey, cardNumber)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *cardQueryCache) SetFindAllCache(ctx context.Context, req *requests.FindAllCards, data *response.ApiResponsePaginationCard) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(cardAllCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *cardQueryCache) SetByActiveCache(ctx context.Context, req *requests.FindAllCards, data *response.ApiResponsePaginationCardDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(cardActiveCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *cardQueryCache) SetByTrashedCache(ctx context.Context, req *requests.FindAllCards, data *response.ApiResponsePaginationCardDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(cardTrashedCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *cardQueryCache) SetByUserIDCache(ctx context.Context, userID int, data *response.ApiResponseCard) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(cardByUserIdCacheKey, userID)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *cardQueryCache) DeleteByIdCache(ctx context.Context, cardID int) {
	key := fmt.Sprintf(cardByIdCacheKey, cardID)
	cache.DeleteFromCache(ctx, c.store, key)
}

func (c *cardQueryCache) DeleteByUserIDCache(ctx context.Context, userID int) {
	key := fmt.Sprintf(cardByUserIdCacheKey, userID)
	cache.DeleteFromCache(ctx, c.store, key)
}

func (c *cardQueryCache) DeleteByCardNumberCache(ctx context.Context, cardNumber string) {
	key := fmt.Sprintf(cardByCardNumCacheKey, cardNumber)
	cache.DeleteFromCache(ctx, c.store, key)
}
