package transaction_stats_bycard_cache

import "time"

const (
	monthTransactionStatusSuccessByCardCacheKey = "transaction:bycard:month:status:success:card:%s:month:%d:year:%d"
	yearTransactionStatusSuccessByCardCacheKey  = "transaction:bycard:year:status:success:card:%s:year:%d"
	monthTransactionStatusFailedByCardCacheKey  = "transaction:bycard:month:status:failed:card:%s:month:%d:year:%d"
	yearTransactionStatusFailedByCardCacheKey   = "transaction:bycard:year:status:failed:card:%s:year:%d"

	monthTransactionAmountByCardCacheKey = "transaction:bycard:month:amount:card:%s:year:%d"
	yearTransactionAmountByCardCacheKey  = "transaction:bycard:year:amount:card:%s:year:%d"

	monthTransactionMethodByCardCacheKey = "transaction:bycard:month:method:card:%s:year:%d"
	yearTransactionMethodByCardCacheKey  = "transaction:bycard:year:method:card:%s:year:%d"

	ttlDefault = 5 * time.Minute
)
