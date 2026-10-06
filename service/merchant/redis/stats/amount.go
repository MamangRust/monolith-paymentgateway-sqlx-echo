package merchantstatscache

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type merchantStatsAmountCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewMerchantStatsAmountCache(store *sharedcachehelpers.CacheStore) MerchantStatsAmountCache {
	return &merchantStatsAmountCache{store: store}
}

func (s *merchantStatsAmountCache) GetMonthlyAmountMerchantCache(ctx context.Context, year int) ([]*db.GetMonthlyAmountMerchantRow, bool) {
	key := fmt.Sprintf(merchantMonthlyAmountCacheKey, year)

	result, found := sharedcachehelpers.GetFromCache[[]*db.GetMonthlyAmountMerchantRow](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (s *merchantStatsAmountCache) SetMonthlyAmountMerchantCache(ctx context.Context, year int, data []*db.GetMonthlyAmountMerchantRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(merchantMonthlyAmountCacheKey, year)

	sharedcachehelpers.SetToCache(ctx, s.store, key, &data, ttlDefault)
}

func (s *merchantStatsAmountCache) GetYearlyAmountMerchantCache(ctx context.Context, year int) ([]*db.GetYearlyAmountMerchantRow, bool) {
	key := fmt.Sprintf(MerchantYearlyAmountCacheKey, year)

	result, found := sharedcachehelpers.GetFromCache[[]*db.GetYearlyAmountMerchantRow](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (s *merchantStatsAmountCache) SetYearlyAmountMerchantCache(ctx context.Context, year int, data []*db.GetYearlyAmountMerchantRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(MerchantYearlyAmountCacheKey, year)

	sharedcachehelpers.SetToCache(ctx, s.store, key, &data, ttlDefault)
}
