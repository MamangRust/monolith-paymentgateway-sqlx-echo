package card_stats_bycard_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type cardStatsTransactionByCardCache struct {
	store *cache.CacheStore
}

func NewCardStatsTransactionByCardCache(store *cache.CacheStore) CardStatsTransactionByCardCache {
	return &cardStatsTransactionByCardCache{store: store}
}

func (c *cardStatsTransactionByCardCache) GetMonthlyTransactionByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseMonthlyAmount, bool) {
	key := fmt.Sprintf(cacheKeyMonthlyTxnByCard, req.CardNumber, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseMonthlyAmount](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *cardStatsTransactionByCardCache) SetMonthlyTransactionByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseMonthlyAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(cacheKeyMonthlyTxnByCard, req.CardNumber, req.Year)
	cache.SetToCache(ctx, c.store, key, data, expirationCardStatistic)
}

func (c *cardStatsTransactionByCardCache) GetYearlyTransactionByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseYearlyAmount, bool) {
	key := fmt.Sprintf(cacheKeyYearlyTxnByCard, req.CardNumber, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseYearlyAmount](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *cardStatsTransactionByCardCache) SetYearlyTransactionByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseYearlyAmount) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(cacheKeyYearlyTxnByCard, req.CardNumber, req.Year)
	cache.SetToCache(ctx, c.store, key, data, expirationCardStatistic)
}
