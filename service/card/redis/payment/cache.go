package cardpaymentmencache

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type cardPaymentCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewCardPaymentCache(store *sharedcachehelpers.CacheStore) CardPaymentCache {
	return &cardPaymentCache{store: store}
}

func (c *cardPaymentCache) GetPaymentHistory(ctx context.Context, cardNumber string, page, pageSize int) ([]*db.GetCardPaymentsByCardNumberRow, bool) {
	key := fmt.Sprintf(paymentHistoryCacheKey, cardNumber, page, pageSize)
	result, found := sharedcachehelpers.GetFromCache[[]*db.GetCardPaymentsByCardNumberRow](ctx, c.store, key)
	if !found || result == nil {
		return nil, false
	}
	return *result, true
}

func (c *cardPaymentCache) SetPaymentHistory(ctx context.Context, cardNumber string, page, pageSize int, data []*db.GetCardPaymentsByCardNumberRow) {
	key := fmt.Sprintf(paymentHistoryCacheKey, cardNumber, page, pageSize)
	sharedcachehelpers.SetToCache(ctx, c.store, key, &data, ttlDefault)
}

func (c *cardPaymentCache) DeletePaymentHistory(ctx context.Context, cardNumber string) {
	// Pattern delete: hapus semua key yang mengandung cardNumber
	// Untuk kesederhanaan, kita hapus key spesifik. Di Redis bisa pakai wildcard scan.
	keys := []string{
		fmt.Sprintf(paymentHistoryCacheKey, cardNumber, 0, 0),
		fmt.Sprintf(paymentCountCacheKey, cardNumber),
	}
	for _, key := range keys {
		sharedcachehelpers.DeleteFromCache(ctx, c.store, key)
	}
}

func (c *cardPaymentCache) GetPaymentCount(ctx context.Context, cardNumber string) (int, bool) {
	key := fmt.Sprintf(paymentCountCacheKey, cardNumber)
	result, found := sharedcachehelpers.GetFromCache[int](ctx, c.store, key)
	if !found || result == nil {
		return 0, false
	}
	return *result, true
}

func (c *cardPaymentCache) SetPaymentCount(ctx context.Context, cardNumber string, count int) {
	key := fmt.Sprintf(paymentCountCacheKey, cardNumber)
	sharedcachehelpers.SetToCache(ctx, c.store, key, &count, ttlDefault)
}

func (c *cardPaymentCache) DeletePaymentCount(ctx context.Context, cardNumber string) {
	key := fmt.Sprintf(paymentCountCacheKey, cardNumber)
	sharedcachehelpers.DeleteFromCache(ctx, c.store, key)
}
