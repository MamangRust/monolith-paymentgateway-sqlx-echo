package transactionstatsbycarcache

import sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"

type TransactionStatsByCardCache interface {
	TransactionStatsByCardAmountCache
	TransactionStatsByCardStatusCache
	TransactionStatsByCardMethodCache
}

type transactionStatsByCardCache struct {
	TransactionStatsByCardAmountCache
	TransactionStatsByCardStatusCache
	TransactionStatsByCardMethodCache
}

func NewTransactionStatsByCardCache(store *sharedcachehelpers.CacheStore) TransactionStatsByCardCache {
	return &transactionStatsByCardCache{
		TransactionStatsByCardAmountCache: NewTransactionStatsByCardAmountCache(store),
		TransactionStatsByCardStatusCache: NewTransactionStatsByCardStatusCache(store),
		TransactionStatsByCardMethodCache: NewTransactionStatsByCardMethodCache(store),
	}
}
