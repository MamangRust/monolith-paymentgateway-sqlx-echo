package cardstatsmencache

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type cardStatsBalanceCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewCardStatsBalanceCache(store *sharedcachehelpers.CacheStore) CardStatsBalanceCache {
	return &cardStatsBalanceCache{store: store}
}

func (c *cardStatsBalanceCache) GetMonthlyBalanceCache(ctx context.Context, year int) ([]*db.GetMonthlyBalancesRow, bool) {
	key := fmt.Sprintf(cacheKeyMonthlyBalance, year)

	result, found := sharedcachehelpers.GetFromCache[[]*db.GetMonthlyBalancesRow](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (c *cardStatsBalanceCache) SetMonthlyBalanceCache(ctx context.Context, year int, data []*db.GetMonthlyBalancesRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(cacheKeyMonthlyBalance, year)
	sharedcachehelpers.SetToCache(ctx, c.store, key, &data, ttlStatistic)
}

func (c *cardStatsBalanceCache) GetYearlyBalanceCache(ctx context.Context, year int) ([]*db.GetYearlyBalancesRow, bool) {
	key := fmt.Sprintf(cacheKeyYearlyBalance, year)

	result, found := sharedcachehelpers.GetFromCache[[]*db.GetYearlyBalancesRow](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (c *cardStatsBalanceCache) SetYearlyBalanceCache(ctx context.Context, year int, data []*db.GetYearlyBalancesRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(cacheKeyYearlyBalance, year)
	sharedcachehelpers.SetToCache(ctx, c.store, key, &data, ttlStatistic)
}
