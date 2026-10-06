package topupstatsbycardcache

import "time"

const (
	monthTopupStatusSuccessByCardCacheKey = "topup:month:status:success:card_number:%s:month:%d:year:%d"
	yearTopupStatusSuccessByCardCacheKey  = "topup:year:status:success:card_number:%s:year:%d"
	monthTopupStatusFailedByCardCacheKey  = "topup:month:status:failed:card_number:%s:month:%d:year:%d"
	yearTopupStatusFailedByCardCacheKey   = "topup:year:status:failed:card_number:%s:year:%d"

	monthTopupAmountByCardCacheKey = "topup:month:amount:card_number:%s:year:%d"
	yearTopupAmountByCardCacheKey  = "topup:year:amount:card_number:%s:year:%d"

	monthTopupMethodByCardCacheKey = "topup:month:method:card_number:%s:year:%d"
	yearTopupMethodByCardCacheKey  = "topup:year:method:card_number:%s:year:%d"

	ttlDefault = 5 * time.Minute
)
