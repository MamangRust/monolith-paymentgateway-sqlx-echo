package mencache

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type transferCacheResponseAll struct {
	Data         []*db.GetTransfersRow `json:"data"`
	TotalRecords *int                  `json:"total_records"`
}

type transferCacheResponseActive struct {
	Data         []*db.GetActiveTransfersRow `json:"data"`
	TotalRecords *int                        `json:"total_records"`
}

type transferCacheResponseTrashed struct {
	Data         []*db.GetTrashedTransfersRow `json:"data"`
	TotalRecords *int                         `json:"total_records"`
}

type transferQueryCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewTransferQueryCache(store *sharedcachehelpers.CacheStore) TransferQueryCache {
	return &transferQueryCache{store: store}
}

func (c *transferQueryCache) GetCachedTransfersCache(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetTransfersRow, *int, bool) {
	key := fmt.Sprintf(transferAllCacheKey, req.Page, req.PageSize, req.Search)

	result, found := sharedcachehelpers.GetFromCache[transferCacheResponseAll](ctx, c.store, key)

	if !found || result == nil {
		return nil, nil, false
	}
	return result.Data, result.TotalRecords, true
}

func (c *transferQueryCache) SetCachedTransfersCache(ctx context.Context, req *requests.FindAllTransfers, data []*db.GetTransfersRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}
	if data == nil {
		data = []*db.GetTransfersRow{}
	}

	key := fmt.Sprintf(transferAllCacheKey, req.Page, req.PageSize, req.Search)
	payload := &transferCacheResponseAll{Data: data, TotalRecords: total}
	sharedcachehelpers.SetToCache(ctx, c.store, key, payload, ttlDefault)
}

func (c *transferQueryCache) GetCachedTransferActiveCache(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetActiveTransfersRow, *int, bool) {
	key := fmt.Sprintf(transferActiveCacheKey, req.Page, req.PageSize, req.Search)

	result, found := sharedcachehelpers.GetFromCache[transferCacheResponseActive](ctx, c.store, key)

	if !found || result == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (c *transferQueryCache) SetCachedTransferActiveCache(ctx context.Context, req *requests.FindAllTransfers, data []*db.GetActiveTransfersRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}
	if data == nil {
		data = []*db.GetActiveTransfersRow{}
	}

	key := fmt.Sprintf(transferActiveCacheKey, req.Page, req.PageSize, req.Search)
	payload := &transferCacheResponseActive{Data: data, TotalRecords: total}
	sharedcachehelpers.SetToCache(ctx, c.store, key, payload, ttlDefault)
}

func (c *transferQueryCache) GetCachedTransferTrashedCache(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetTrashedTransfersRow, *int, bool) {
	key := fmt.Sprintf(transferTrashedCacheKey, req.Page, req.PageSize, req.Search)

	result, found := sharedcachehelpers.GetFromCache[transferCacheResponseTrashed](ctx, c.store, key)

	if !found || result == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (c *transferQueryCache) SetCachedTransferTrashedCache(ctx context.Context, req *requests.FindAllTransfers, data []*db.GetTrashedTransfersRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}
	if data == nil {
		data = []*db.GetTrashedTransfersRow{}
	}

	key := fmt.Sprintf(transferTrashedCacheKey, req.Page, req.PageSize, req.Search)
	payload := &transferCacheResponseTrashed{Data: data, TotalRecords: total}
	sharedcachehelpers.SetToCache(ctx, c.store, key, payload, ttlDefault)
}

func (c *transferQueryCache) GetCachedTransferCache(ctx context.Context, id int) (*db.GetTransferByIDRow, bool) {
	key := fmt.Sprintf(transferByIdCacheKey, id)
	result, found := sharedcachehelpers.GetFromCache[*db.GetTransferByIDRow](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (c *transferQueryCache) SetCachedTransferCache(ctx context.Context, data *db.GetTransferByIDRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(transferByIdCacheKey, data.TransferID)
	sharedcachehelpers.SetToCache(ctx, c.store, key, data, ttlDefault)
}

func (c *transferQueryCache) GetCachedTransferByFrom(ctx context.Context, from string) ([]*db.GetTransfersBySourceCardRow, bool) {
	key := fmt.Sprintf(transferByFromCacheKey, from)
	result, found := sharedcachehelpers.GetFromCache[[]*db.GetTransfersBySourceCardRow](ctx, c.store, key)
	if !found || result == nil {
		return nil, false
	}
	return *result, true
}

func (c *transferQueryCache) SetCachedTransferByFrom(ctx context.Context, from string, data []*db.GetTransfersBySourceCardRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(transferByFromCacheKey, from)
	sharedcachehelpers.SetToCache(ctx, c.store, key, &data, ttlDefault)
}

func (c *transferQueryCache) GetCachedTransferByTo(ctx context.Context, to string) ([]*db.GetTransfersByDestinationCardRow, bool) {
	key := fmt.Sprintf(transferByToCacheKey, to)

	result, found := sharedcachehelpers.GetFromCache[[]*db.GetTransfersByDestinationCardRow](ctx, c.store, key)
	if !found || result == nil {
		return nil, false
	}
	return *result, true
}

func (c *transferQueryCache) SetCachedTransferByTo(ctx context.Context, to string, data []*db.GetTransfersByDestinationCardRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(transferByToCacheKey, to)
	sharedcachehelpers.SetToCache(ctx, c.store, key, &data, ttlDefault)
}
