package saldostatscache

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type saldoStatsBalanceCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewSaldoStatsBalanceCache(store *sharedcachehelpers.CacheStore) SaldoStatsBalanceCache {
	return &saldoStatsBalanceCache{store: store}
}

func (c *saldoStatsBalanceCache) GetMonthlySaldoBalanceCache(ctx context.Context, year int) ([]*db.GetMonthlySaldoBalancesRow, bool) {
	key := fmt.Sprintf(saldoMonthBalanceCacheKey, year)
	result, found := sharedcachehelpers.GetFromCache[[]*db.GetMonthlySaldoBalancesRow](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (c *saldoStatsBalanceCache) SetMonthlySaldoBalanceCache(ctx context.Context, year int, data []*db.GetMonthlySaldoBalancesRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(saldoMonthBalanceCacheKey, year)
	sharedcachehelpers.SetToCache(ctx, c.store, key, &data, ttlDefault)
}

func (c *saldoStatsBalanceCache) GetYearlySaldoBalanceCache(ctx context.Context, year int) ([]*db.GetYearlySaldoBalancesRow, bool) {
	key := fmt.Sprintf(saldoYearlyBalanceCacheKey, year)
	result, found := sharedcachehelpers.GetFromCache[[]*db.GetYearlySaldoBalancesRow](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (c *saldoStatsBalanceCache) SetYearlySaldoBalanceCache(ctx context.Context, year int, data []*db.GetYearlySaldoBalancesRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(saldoYearlyBalanceCacheKey, year)
	sharedcachehelpers.SetToCache(ctx, c.store, key, &data, ttlDefault)
}
