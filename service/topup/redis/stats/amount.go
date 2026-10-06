package topupstatscache

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type topupStatsAmountCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewTopupStatsAmountCache(store *sharedcachehelpers.CacheStore) TopupStatsAmountCache {
	return &topupStatsAmountCache{store: store}
}

func (c *topupStatsAmountCache) GetMonthlyTopupAmountsCache(ctx context.Context, year int) ([]*db.GetMonthlyTopupAmountsRow, bool) {
	key := fmt.Sprintf(monthTopupAmountCacheKey, year)

	result, found := sharedcachehelpers.GetFromCache[[]*db.GetMonthlyTopupAmountsRow](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (c *topupStatsAmountCache) SetMonthlyTopupAmountsCache(ctx context.Context, year int, data []*db.GetMonthlyTopupAmountsRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(monthTopupAmountCacheKey, year)
	sharedcachehelpers.SetToCache(ctx, c.store, key, &data, ttlDefault)
}

func (c *topupStatsAmountCache) GetYearlyTopupAmountsCache(ctx context.Context, year int) ([]*db.GetYearlyTopupAmountsRow, bool) {
	key := fmt.Sprintf(yearTopupAmountCacheKey, year)

	result, found := sharedcachehelpers.GetFromCache[[]*db.GetYearlyTopupAmountsRow](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (c *topupStatsAmountCache) SetYearlyTopupAmountsCache(ctx context.Context, year int, data []*db.GetYearlyTopupAmountsRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(yearTopupAmountCacheKey, year)
	sharedcachehelpers.SetToCache(ctx, c.store, key, &data, ttlDefault)
}
