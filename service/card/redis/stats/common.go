package cardstatsmencache

import "time"

// cardStatisticCache is a struct that represents the cache store
const (
	cacheKeyMonthlyBalance = "stat:monthly:balance:%d"
	cacheKeyYearlyBalance  = "stat:yearly:balance:%d"

	cacheKeyMonthlyTopupAmount = "stat:monthly:topup:%d"
	cacheKeyYearlyTopupAmount  = "stat:yearly:topup:%d"

	cacheKeyMonthlyWithdrawAmount = "stat:monthly:withdraw:%d"
	cacheKeyYearlyWithdrawAmount  = "stat:yearly:withdraw:%d"

	cacheKeyMonthlyTransactionAmount = "stat:monthly:transaction:%d"
	cacheKeyYearlyTransactionAmount  = "stat:yearly:transaction:%d"

	cacheKeyMonthlyTransferSender = "stat:monthly:transfer:sender:%d"
	cacheKeyYearlyTransferSender  = "stat:yearly:transfer:sender:%d"

	cacheKeyMonthlyTransferReceiver = "stat:monthly:transfer:receiver:%d"
	cacheKeyYearlyTransferReceiver  = "stat:yearly:transfer:receiver:%d"

	ttlStatistic = 10 * time.Minute
)
