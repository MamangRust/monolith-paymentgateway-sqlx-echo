package transferstatscache

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type transferStatsAmountCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewTransferStatsAmountCache(store *sharedcachehelpers.CacheStore) TransferStatsAmountCache {
	return &transferStatsAmountCache{store: store}
}

func (t *transferStatsAmountCache) GetCachedMonthTransferAmounts(ctx context.Context, year int) ([]*db.GetMonthlyTransferAmountsRow, bool) {
	key := fmt.Sprintf(transferMonthTransferAmountKey, year)
	result, found := sharedcachehelpers.GetFromCache[[]*db.GetMonthlyTransferAmountsRow](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (t *transferStatsAmountCache) SetCachedMonthTransferAmounts(ctx context.Context, year int, data []*db.GetMonthlyTransferAmountsRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(transferMonthTransferAmountKey, year)
	sharedcachehelpers.SetToCache(ctx, t.store, key, &data, ttlDefault)
}

func (t *transferStatsAmountCache) GetCachedYearlyTransferAmounts(ctx context.Context, year int) ([]*db.GetYearlyTransferAmountsRow, bool) {
	key := fmt.Sprintf(transferYearTransferAmountKey, year)
	result, found := sharedcachehelpers.GetFromCache[[]*db.GetYearlyTransferAmountsRow](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (t *transferStatsAmountCache) SetCachedYearlyTransferAmounts(ctx context.Context, year int, data []*db.GetYearlyTransferAmountsRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(transferYearTransferAmountKey, year)
	sharedcachehelpers.SetToCache(ctx, t.store, key, &data, ttlDefault)
}
