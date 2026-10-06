package saldo_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type saldoQueryCache struct {
	store *cache.CacheStore
}

func NewSaldoQueryCache(store *cache.CacheStore) SaldoQueryCache {
	return &saldoQueryCache{store: store}
}

func (s *saldoQueryCache) GetCachedSaldos(ctx context.Context, req *requests.FindAllSaldos) (*response.ApiResponsePaginationSaldo, bool) {
	key := fmt.Sprintf(saldoAllCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationSaldo](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (s *saldoQueryCache) GetCachedSaldoByActive(ctx context.Context, req *requests.FindAllSaldos) (*response.ApiResponsePaginationSaldoDeleteAt, bool) {
	key := fmt.Sprintf(saldoActiveCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationSaldoDeleteAt](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (s *saldoQueryCache) GetCachedSaldoByTrashed(ctx context.Context, req *requests.FindAllSaldos) (*response.ApiResponsePaginationSaldoDeleteAt, bool) {
	key := fmt.Sprintf(saldoTrashedCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[response.ApiResponsePaginationSaldoDeleteAt](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (s *saldoQueryCache) GetCachedSaldoById(ctx context.Context, saldo_id int) (*response.ApiResponseSaldo, bool) {
	key := fmt.Sprintf(saldoByIdCacheKey, saldo_id)
	result, found := cache.GetFromCache[response.ApiResponseSaldo](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (s *saldoQueryCache) GetCachedSaldoByCardNumber(ctx context.Context, card_number string) (*response.ApiResponseSaldo, bool) {
	key := fmt.Sprintf(saldoByCardNumberKey, card_number)
	result, found := cache.GetFromCache[response.ApiResponseSaldo](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (s *saldoQueryCache) SetCachedSaldos(ctx context.Context, req *requests.FindAllSaldos, data *response.ApiResponsePaginationSaldo) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(saldoAllCacheKey, req.Page, req.PageSize, req.Search)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *saldoQueryCache) SetCachedSaldoByActive(ctx context.Context, req *requests.FindAllSaldos, data *response.ApiResponsePaginationSaldoDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(saldoActiveCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *saldoQueryCache) SetCachedSaldoByTrashed(ctx context.Context, req *requests.FindAllSaldos, data *response.ApiResponsePaginationSaldoDeleteAt) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(saldoTrashedCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *saldoQueryCache) SetCachedSaldoById(ctx context.Context, saldo_id int, data *response.ApiResponseSaldo) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(saldoByIdCacheKey, saldo_id)
	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *saldoQueryCache) SetCachedSaldoByCardNumber(ctx context.Context, card_number string, data *response.ApiResponseSaldo) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(saldoByCardNumberKey, card_number)
	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}
