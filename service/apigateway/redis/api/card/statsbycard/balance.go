package card_stats_bycard_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type cardStatsBalanceByCardCache struct {
	store *cache.CacheStore
}

func NewCardStatsBalanceByCardCache(store *cache.CacheStore) CardStatsBalanceByCardCache {
	return &cardStatsBalanceByCardCache{store: store}
}

func (c *cardStatsBalanceByCardCache) GetMonthlyBalanceByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseMonthlyBalance, bool) {
	key := fmt.Sprintf(cacheKeyMonthlyBalanceByCard, req.CardNumber, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseMonthlyBalance](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *cardStatsBalanceByCardCache) SetMonthlyBalanceByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseMonthlyBalance) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(cacheKeyMonthlyBalanceByCard, req.CardNumber, req.Year)
	cache.SetToCache(ctx, c.store, key, data, expirationCardStatistic)
}

func (c *cardStatsBalanceByCardCache) GetYearlyBalanceByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseYearlyBalance, bool) {
	key := fmt.Sprintf(cacheKeyYearlyBalanceByCard, req.CardNumber, req.Year)
	result, found := cache.GetFromCache[response.ApiResponseYearlyBalance](ctx, c.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (c *cardStatsBalanceByCardCache) SetYearlyBalanceByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseYearlyBalance) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(cacheKeyYearlyBalanceByCard, req.CardNumber, req.Year)
	cache.SetToCache(ctx, c.store, key, data, expirationCardStatistic)
}
