package merchant_stats_byapikey_cache

import "time"

const (
	merchantMonthlyPaymentMethodByApikeyCacheKey = "merchant:statistic:monthly:payment-method:apikey:%s:year:%d"

	merchantYearlyPaymentMethodByApikeyCacheKey = "merchant:statistic:yearly:payment-method:apikey:%s:year:%d"

	merchantMonthlyAmountByApikeyCacheKey = "merchant:statistic:monthly:amount:apikey:%s:year:%d"

	merchantYearlyAmountByApikeyCacheKey = "merchant:statistic:yearly:amount:apikey:%s:year:%d"

	merchantMonthlyTotalAmountByApikeyCacheKey = "merchant:statistic:monthly:total-amount:apikey:%s:year:%d"

	merchantYearlyTotalAmountByApikeyCacheKey = "merchant:statistic:yearly:total-amount:apikey:%s:year:%d"

	ttlDefault = 5 * time.Minute
)
