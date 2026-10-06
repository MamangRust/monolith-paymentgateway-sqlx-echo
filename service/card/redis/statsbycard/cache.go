package cardstatsbycardmencache

import sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"

type CardStatsByCardCache interface {
	CardStatsBalanceByCardCache
	CardStatsTopupByCardCache
	CardStatsTransactionByCardCache
	CardStatsTransferByCardCache
	CardStatsWithdrawByCardCache
}

type MencacheStatsByCard struct {
	CardStatsBalanceByCardCache
	CardStatsTopupByCardCache
	CardStatsTransactionByCardCache
	CardStatsTransferByCardCache
	CardStatsWithdrawByCardCache
}

func NewMencacheStatsByCard(store *sharedcachehelpers.CacheStore) CardStatsByCardCache {
	return &MencacheStatsByCard{
		CardStatsBalanceByCardCache:     NewCardStatsBalanceByCardCache(store),
		CardStatsTopupByCardCache:       NewCardStatsTopupByCardCache(store),
		CardStatsTransactionByCardCache: NewCardStatsTransactionByCardCache(store),
		CardStatsTransferByCardCache:    NewCardStatsTransferByCardCache(store),
		CardStatsWithdrawByCardCache:    NewCardStatsWithdrawByCardCache(store),
	}
}
