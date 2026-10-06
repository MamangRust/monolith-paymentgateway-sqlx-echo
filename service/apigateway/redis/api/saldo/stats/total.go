package saldo_stats_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type saldoStatsTotalCache struct {
	store *cache.CacheStore
}

func NewSaldoStatsTotalCache(store *cache.CacheStore) SaldoStatsTotalCache {
	return &saldoStatsTotalCache{store: store}
}

func (c *saldoStatsTotalCache) GetMonthlyTotalSaldoBalanceCache(ctx context.Context, req *requests.MonthTotalSaldoBalance) (*response.ApiResponseMonthTotalSaldo, bool) {
	key := fmt.Sprintf(saldoMonthTotalBalanceCacheKey, req.Month, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseMonthTotalSaldo](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *saldoStatsTotalCache) SetMonthlyTotalSaldoCache(ctx context.Context, req *requests.MonthTotalSaldoBalance, data *response.ApiResponseMonthTotalSaldo) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(saldoMonthTotalBalanceCacheKey, req.Month, req.Year)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *saldoStatsTotalCache) GetYearTotalSaldoBalanceCache(ctx context.Context, year int) (*response.ApiResponseYearTotalSaldo, bool) {
	key := fmt.Sprintf(saldoYearTotalBalanceCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseYearTotalSaldo](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *saldoStatsTotalCache) SetYearTotalSaldoBalanceCache(ctx context.Context, year int, data *response.ApiResponseYearTotalSaldo) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(saldoYearTotalBalanceCacheKey, year)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}
