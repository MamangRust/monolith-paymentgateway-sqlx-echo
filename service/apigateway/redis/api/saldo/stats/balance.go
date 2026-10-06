package saldo_stats_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type saldoStatsBalanceCache struct {
	store *cache.CacheStore
}

func NewSaldoStatsBalanceCache(store *cache.CacheStore) SaldoStatsBalanceCache {
	return &saldoStatsBalanceCache{store: store}
}

func (c *saldoStatsBalanceCache) GetMonthlySaldoBalanceCache(ctx context.Context, year int) (*response.ApiResponseMonthSaldoBalances, bool) {
	key := fmt.Sprintf(saldoMonthBalanceCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseMonthSaldoBalances](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *saldoStatsBalanceCache) SetMonthlySaldoBalanceCache(ctx context.Context, year int, data *response.ApiResponseMonthSaldoBalances) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(saldoMonthBalanceCacheKey, year)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *saldoStatsBalanceCache) GetYearlySaldoBalanceCache(ctx context.Context, year int) (*response.ApiResponseYearSaldoBalances, bool) {
	key := fmt.Sprintf(saldoYearlyBalanceCacheKey, year)
	result, found := cache.GetFromCache[response.ApiResponseYearSaldoBalances](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *saldoStatsBalanceCache) SetYearlySaldoBalanceCache(ctx context.Context, year int, data *response.ApiResponseYearSaldoBalances) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(saldoYearlyBalanceCacheKey, year)
	cache.SetToCache(ctx, c.store, key, data, ttlDefault)
}
