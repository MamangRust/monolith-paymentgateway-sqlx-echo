package withdrawstatscache

import "time"

// Constants for cache keys
const (
	montWithdrawStatusSuccessKey = "withdraws:mont:status:success:month%d:year:%d"
	yearWithdrawStatusSuccessKey = "withdraws:year:status:success:year:%d"

	montWithdrawStatusFailedKey = "withdraws:mont:status:failed:month:%d:year:%d"
	yearWithdrawStatusFailedKey = "withdraws:year:status:failed:year:%d"

	montWithdrawAmountKey = "withdraws:mont:amount:year:%d"
	yearWithdrawAmountKey = "withdraws:year:amount:year:%d"

	ttlDefault = 5 * time.Minute
)
