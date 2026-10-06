package topupstatscache

import "time"

const (
	monthTopupStatusSuccessCacheKey = "topup:month:status:success:month:%d:year:%d"
	yearTopupStatusSuccessCacheKey  = "topup:year:status:success:year:%d"
	monthTopupStatusFailedCacheKey  = "topup:month:status:failed:month:%d:year:%d"
	yearTopupStatusFailedCacheKey   = "topup:year:status:failed:year:%d"

	monthTopupAmountCacheKey = "topup:month:amount:year:%d"
	yearTopupAmountCacheKey  = "topup:year:amount:year:%d"

	monthTopupMethodCacheKey = "topup:month:method:year:%d"
	yearTopupMethodCacheKey  = "topup:year:method:year:%d"

	ttlDefault = 5 * time.Minute
)
