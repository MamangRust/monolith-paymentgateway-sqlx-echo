package transferstatscache

import "time"

// Constants for cache keys
const (
	transferMonthTransferStatusSuccessKey = "transfer:month:transfer_status:success:month:%d:year:%d"
	transferMonthTransferStatusFailedKey  = "transfer:month:transfer_status:failed:month:%d:year:%d"

	transferYearTransferStatusSuccessKey = "transfer:year:transfer_status:success:year:%d"
	transferYearTransferStatusFailedKey  = "transfer:year:transfer_status:failed:year:%d"

	transferMonthTransferAmountKey = "transfer:month:transfer_amount:year:%d"
	transferYearTransferAmountKey  = "transfer:year:transfer_amount:year:%d"

	ttlDefault = 5 * time.Minute
)
