package mencache

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type merchantCachedResponseAll struct {
	Data         []*db.GetMerchantsRow `json:"data"`
	TotalRecords *int                  `json:"total_records"`
}

type merchantCachedResponseActive struct {
	Data         []*db.GetActiveMerchantsRow `json:"data"`
	TotalRecords *int                        `json:"total_records"`
}

type merchantCachedResponseTrashed struct {
	Data         []*db.GetTrashedMerchantsRow `json:"data"`
	TotalRecords *int                         `json:"total_records"`
}

type merchantQueryCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewMerchantQueryCache(store *sharedcachehelpers.CacheStore) MerchantQueryCache {
	return &merchantQueryCache{store: store}
}

func (m *merchantQueryCache) GetCachedMerchants(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetMerchantsRow, *int, bool) {
	key := fmt.Sprintf(merchantAllCacheKey, req.Page, req.PageSize, req.Search)

	result, found := sharedcachehelpers.GetFromCache[merchantCachedResponseAll](ctx, m.store, key)

	if !found || result == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (m *merchantQueryCache) SetCachedMerchants(ctx context.Context, req *requests.FindAllMerchants, data []*db.GetMerchantsRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}

	if data == nil {
		data = []*db.GetMerchantsRow{}
	}

	key := fmt.Sprintf(merchantAllCacheKey, req.Page, req.PageSize, req.Search)

	payload := &merchantCachedResponseAll{Data: data, TotalRecords: total}
	sharedcachehelpers.SetToCache(ctx, m.store, key, payload, ttlDefault)
}

func (m *merchantQueryCache) GetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetActiveMerchantsRow, *int, bool) {
	key := fmt.Sprintf(merchantActiveCacheKey, req.Page, req.PageSize, req.Search)

	result, found := sharedcachehelpers.GetFromCache[merchantCachedResponseActive](ctx, m.store, key)

	if !found || result == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (m *merchantQueryCache) SetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchants, data []*db.GetActiveMerchantsRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}

	if data == nil {
		data = []*db.GetActiveMerchantsRow{}
	}

	key := fmt.Sprintf(merchantActiveCacheKey, req.Page, req.PageSize, req.Search)

	payload := &merchantCachedResponseActive{Data: data, TotalRecords: total}
	sharedcachehelpers.SetToCache(ctx, m.store, key, payload, ttlDefault)
}

func (m *merchantQueryCache) GetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetTrashedMerchantsRow, *int, bool) {
	key := fmt.Sprintf(merchantTrashedCacheKey, req.Page, req.PageSize, req.Search)

	result, found := sharedcachehelpers.GetFromCache[merchantCachedResponseTrashed](ctx, m.store, key)

	if !found || result == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (m *merchantQueryCache) SetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchants, data []*db.GetTrashedMerchantsRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}

	if data == nil {
		data = []*db.GetTrashedMerchantsRow{}
	}

	key := fmt.Sprintf(merchantTrashedCacheKey, req.Page, req.PageSize, req.Search)

	payload := &merchantCachedResponseTrashed{Data: data, TotalRecords: total}
	sharedcachehelpers.SetToCache(ctx, m.store, key, payload, ttlDefault)
}

func (m *merchantQueryCache) GetCachedMerchant(ctx context.Context, id int) (*db.GetMerchantByIDRow, bool) {
	key := fmt.Sprintf(merchantByIdCacheKey, id)

	result, found := sharedcachehelpers.GetFromCache[*db.GetMerchantByIDRow](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (m *merchantQueryCache) SetCachedMerchant(ctx context.Context, data *db.GetMerchantByIDRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(merchantByIdCacheKey, data.MerchantID)

	sharedcachehelpers.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantQueryCache) GetCachedMerchantsByUserId(ctx context.Context, userId int) ([]*db.GetMerchantsByUserIDRow, bool) {
	key := fmt.Sprintf(merchantByUserIdCacheKey, userId)

	result, found := sharedcachehelpers.GetFromCache[[]*db.GetMerchantsByUserIDRow](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (m *merchantQueryCache) SetCachedMerchantsByUserId(ctx context.Context, userId int, data []*db.GetMerchantsByUserIDRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(merchantByUserIdCacheKey, userId)

	sharedcachehelpers.SetToCache(ctx, m.store, key, &data, ttlDefault)
}

func (m *merchantQueryCache) GetCachedMerchantByApiKey(ctx context.Context, apiKey string) (*db.GetMerchantByApiKeyRow, bool) {
	key := fmt.Sprintf(merchantByApiKeyCacheKey, apiKey)

	result, found := sharedcachehelpers.GetFromCache[*db.GetMerchantByApiKeyRow](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (m *merchantQueryCache) SetCachedMerchantByApiKey(ctx context.Context, apiKey string, data *db.GetMerchantByApiKeyRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(merchantByApiKeyCacheKey, apiKey)

	sharedcachehelpers.SetToCache(ctx, m.store, key, data, ttlDefault)
}
