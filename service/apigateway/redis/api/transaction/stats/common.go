package transaction_stats_cache

import "time"

const (
	monthTransactionStatusSuccessCacheKey = "transaction:month:status:success:month:%d:year:%d"
	yearTransactionStatusSuccessCacheKey  = "transaction:year:status:success:year:%d"
	monthTransactionStatusFailedCacheKey  = "transaction:month:status:failed:month:%d:year:%d"
	yearTransactionStatusFailedCacheKey   = "transaction:year:status:failed:year:%d"

	monthTransactionAmountCacheKey = "transaction:month:amount:year:%d"
	yearTransactionAmountCacheKey  = "transaction:year:amount:year:%d"

	monthTransactionMethodCacheKey = "transaction:month:method:year:%d"
	yearTransactionMethodCacheKey  = "transaction:year:method:year:%d"

	ttlDefault = 5 * time.Minute
)
