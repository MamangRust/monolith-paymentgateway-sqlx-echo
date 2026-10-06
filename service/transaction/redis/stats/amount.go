package transactionstatscache

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type transactionStatsAmountCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewTransactionStatsAmountCache(store *sharedcachehelpers.CacheStore) TransactionStatsAmountCache {
	return &transactionStatsAmountCache{store: store}
}

func (t *transactionStatsAmountCache) GetMonthlyAmountsCache(ctx context.Context, year int) ([]*db.GetMonthlyAmountsRow, bool) {
	key := fmt.Sprintf(monthTransactionAmountCacheKey, year)
	result, found := sharedcachehelpers.GetFromCache[[]*db.GetMonthlyAmountsRow](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (t *transactionStatsAmountCache) SetMonthlyAmountsCache(ctx context.Context, year int, data []*db.GetMonthlyAmountsRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(monthTransactionAmountCacheKey, year)
	sharedcachehelpers.SetToCache(ctx, t.store, key, &data, ttlDefault)
}

func (t *transactionStatsAmountCache) GetYearlyAmountsCache(ctx context.Context, year int) ([]*db.GetYearlyAmountsRow, bool) {
	key := fmt.Sprintf(yearTransactionAmountCacheKey, year)

	result, found := sharedcachehelpers.GetFromCache[[]*db.GetYearlyAmountsRow](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return *result, true
}

func (t *transactionStatsAmountCache) SetYearlyAmountsCache(ctx context.Context, year int, data []*db.GetYearlyAmountsRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(yearTransactionAmountCacheKey, year)
	sharedcachehelpers.SetToCache(ctx, t.store, key, &data, ttlDefault)
}
