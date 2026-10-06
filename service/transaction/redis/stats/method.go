package transactionstatscache

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type transactionStatsMethodCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewTransactionStatsMethodCache(store *sharedcachehelpers.CacheStore) TransactionStatsMethodCache {
	return &transactionStatsMethodCache{store: store}
}

func (t *transactionStatsMethodCache) GetMonthlyPaymentMethodsCache(ctx context.Context, year int) ([]*db.GetMonthlyPaymentMethodsRow, bool) {
	key := fmt.Sprintf(monthTransactionMethodCacheKey, year)
	result, found := sharedcachehelpers.GetFromCache[[]*db.GetMonthlyPaymentMethodsRow](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (t *transactionStatsMethodCache) SetMonthlyPaymentMethodsCache(ctx context.Context, year int, data []*db.GetMonthlyPaymentMethodsRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(monthTransactionMethodCacheKey, year)
	sharedcachehelpers.SetToCache(ctx, t.store, key, &data, ttlDefault)
}

func (t *transactionStatsMethodCache) GetYearlyPaymentMethodsCache(ctx context.Context, year int) ([]*db.GetYearlyPaymentMethodsRow, bool) {
	key := fmt.Sprintf(yearTransactionMethodCacheKey, year)
	result, found := sharedcachehelpers.GetFromCache[[]*db.GetYearlyPaymentMethodsRow](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return *result, true
}

func (t *transactionStatsMethodCache) SetYearlyPaymentMethodsCache(ctx context.Context, year int, data []*db.GetYearlyPaymentMethodsRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(yearTransactionMethodCacheKey, year)
	sharedcachehelpers.SetToCache(ctx, t.store, key, &data, ttlDefault)
}
