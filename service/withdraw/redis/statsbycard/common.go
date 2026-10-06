package withdrawstatsbycardcache

import "time"

// Constants for cache keys
const (
	monthWithdrawStatusSuccessByCardKey = "withdraws:month:status:success:card_number:%s:month:%d:year:%d"
	yearWithdrawStatusSuccessByCardKey  = "withdraws:year:status:success:card_number:%s:year:%d"

	monthWithdrawStatusFailedByCardKey = "withdraws:month:status:failed:card_number:%s:month:%d:year:%d"
	yearWithdrawStatusFailedByCardKey  = "withdraws:year:status:failed:card_number:%s:year:%d"

	monthWithdrawAmountByCardKey = "withdraws:month:amount:card_number:%s:year:%d"
	yearWithdrawAmountByCardKey  = "withdraws:year:amount:card_number:%s:year:%d"

	ttlDefault = 5 * time.Minute
)
