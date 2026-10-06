package cardrewardmencache

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type cardRewardCache struct {
	store *sharedcachehelpers.CacheStore
}

func NewCardRewardCache(store *sharedcachehelpers.CacheStore) CardRewardCache {
	return &cardRewardCache{store: store}
}

func (c *cardRewardCache) GetBalance(ctx context.Context, cardNumber string) (int64, bool) {
	key := fmt.Sprintf(rewardBalanceCacheKey, cardNumber)
	result, found := sharedcachehelpers.GetFromCache[int64](ctx, c.store, key)
	if !found || result == nil {
		return 0, false
	}
	return *result, true
}

func (c *cardRewardCache) SetBalance(ctx context.Context, cardNumber string, balance int64) {
	key := fmt.Sprintf(rewardBalanceCacheKey, cardNumber)
	sharedcachehelpers.SetToCache(ctx, c.store, key, &balance, ttlDefault)
}

func (c *cardRewardCache) DeleteBalance(ctx context.Context, cardNumber string) {
	key := fmt.Sprintf(rewardBalanceCacheKey, cardNumber)
	sharedcachehelpers.DeleteFromCache(ctx, c.store, key)
}

func (c *cardRewardCache) GetHistory(ctx context.Context, cardNumber string) ([]*db.CardReward, bool) {
	key := fmt.Sprintf(rewardHistoryCacheKey, cardNumber)
	result, found := sharedcachehelpers.GetFromCache[[]*db.CardReward](ctx, c.store, key)
	if !found || result == nil {
		return nil, false
	}
	return *result, true
}

func (c *cardRewardCache) SetHistory(ctx context.Context, cardNumber string, data []*db.CardReward) {
	key := fmt.Sprintf(rewardHistoryCacheKey, cardNumber)
	sharedcachehelpers.SetToCache(ctx, c.store, key, &data, ttlDefault)
}

func (c *cardRewardCache) DeleteHistory(ctx context.Context, cardNumber string) {
	key := fmt.Sprintf(rewardHistoryCacheKey, cardNumber)
	sharedcachehelpers.DeleteFromCache(ctx, c.store, key)
}
