package transferstatsbycardcache

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type transferStatsByCardAmountCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewTransferStatsByCardAmountCache(store *sharedcachehelpers.CacheStore) TransferStatsByCardAmountCache {
	return &transferStatsByCardAmountCache{store: store}
}

func (t *transferStatsByCardAmountCache) GetMonthlyTransferAmountsBySenderCard(ctx context.Context, req *requests.MonthYearCardNumber) ([]*db.GetMonthlyTransferAmountsBySenderCardNumberRow, bool) {
	key := fmt.Sprintf(transferMonthTransferAmountBySenderCardKey, req.CardNumber, req.Year)
	result, found := sharedcachehelpers.GetFromCache[[]*db.GetMonthlyTransferAmountsBySenderCardNumberRow](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (t *transferStatsByCardAmountCache) SetMonthlyTransferAmountsBySenderCard(ctx context.Context, req *requests.MonthYearCardNumber, data []*db.GetMonthlyTransferAmountsBySenderCardNumberRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(transferMonthTransferAmountBySenderCardKey, req.CardNumber, req.Year)
	sharedcachehelpers.SetToCache(ctx, t.store, key, &data, ttlDefault)
}

func (t *transferStatsByCardAmountCache) GetMonthlyTransferAmountsByReceiverCard(ctx context.Context, req *requests.MonthYearCardNumber) ([]*db.GetMonthlyTransferAmountsByReceiverCardNumberRow, bool) {
	key := fmt.Sprintf(transferMonthTransferAmountByReceiverCardKey, req.CardNumber, req.Year)
	result, found := sharedcachehelpers.GetFromCache[[]*db.GetMonthlyTransferAmountsByReceiverCardNumberRow](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (t *transferStatsByCardAmountCache) SetMonthlyTransferAmountsByReceiverCard(ctx context.Context, req *requests.MonthYearCardNumber, data []*db.GetMonthlyTransferAmountsByReceiverCardNumberRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(transferMonthTransferAmountByReceiverCardKey, req.CardNumber, req.Year)
	sharedcachehelpers.SetToCache(ctx, t.store, key, &data, ttlDefault)
}

func (t *transferStatsByCardAmountCache) GetYearlyTransferAmountsBySenderCard(ctx context.Context, req *requests.MonthYearCardNumber) ([]*db.GetYearlyTransferAmountsBySenderCardNumberRow, bool) {
	key := fmt.Sprintf(transferYearTransferAmountBySenderCardKey, req.CardNumber, req.Year)
	result, found := sharedcachehelpers.GetFromCache[[]*db.GetYearlyTransferAmountsBySenderCardNumberRow](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (t *transferStatsByCardAmountCache) SetYearlyTransferAmountsBySenderCard(ctx context.Context, req *requests.MonthYearCardNumber, data []*db.GetYearlyTransferAmountsBySenderCardNumberRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(transferYearTransferAmountBySenderCardKey, req.CardNumber, req.Year)
	sharedcachehelpers.SetToCache(ctx, t.store, key, &data, ttlDefault)
}

func (t *transferStatsByCardAmountCache) GetYearlyTransferAmountsByReceiverCard(ctx context.Context, req *requests.MonthYearCardNumber) ([]*db.GetYearlyTransferAmountsByReceiverCardNumberRow, bool) {
	key := fmt.Sprintf(transferYearTransferAmountByReceiverCardKey, req.CardNumber, req.Year)
	result, found := sharedcachehelpers.GetFromCache[[]*db.GetYearlyTransferAmountsByReceiverCardNumberRow](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}
	return *result, true
}

func (t *transferStatsByCardAmountCache) SetYearlyTransferAmountsByReceiverCard(ctx context.Context, req *requests.MonthYearCardNumber, data []*db.GetYearlyTransferAmountsByReceiverCardNumberRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(transferYearTransferAmountByReceiverCardKey, req.CardNumber, req.Year)
	sharedcachehelpers.SetToCache(ctx, t.store, key, &data, ttlDefault)
}
