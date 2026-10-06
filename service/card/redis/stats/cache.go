package cardstatsmencache

import (
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type CardStatsCache interface {
	CardStatsBalanceCache
	CardStatsTopupCache
	CardStatsTransactionCache
	CardStatsTransferCache
	CardStatsWithdrawCache
}

type mencacheStats struct {
	CardStatsBalanceCache
	CardStatsTopupCache
	CardStatsTransactionCache
	CardStatsTransferCache
	CardStatsWithdrawCache
}

func NewMencacheStats(store *sharedcachehelpers.CacheStore) CardStatsCache {
	return &mencacheStats{
		CardStatsBalanceCache:     NewCardStatsBalanceCache(store),
		CardStatsTopupCache:       NewCardStatsTopupCache(store),
		CardStatsTransactionCache: NewCardStatsTransactionCache(store),
		CardStatsTransferCache:    NewCardStatsTransferCache(store),
		CardStatsWithdrawCache:    NewCardStatsWithdrawCache(store),
	}
}
