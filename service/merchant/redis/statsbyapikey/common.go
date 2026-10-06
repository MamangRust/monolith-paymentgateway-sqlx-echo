package merchantstatsapikey

import "time"

const (
	// Cache key for monthly payment method statistics by API key
	merchantMonthlyPaymentMethodByApikeyCacheKey = "merchant:statistic:monthly:payment-method:apikey:%s:year:%d"

	// Cache key for yearly payment method statistics by API key
	merchantYearlyPaymentMethodByApikeyCacheKey = "merchant:statistic:yearly:payment-method:apikey:%s:year:%d"

	// Cache key for monthly amount statistics by API key
	merchantMonthlyAmountByApikeyCacheKey = "merchant:statistic:monthly:amount:apikey:%s:year:%d"

	// Cache key for yearly amount statistics by API key
	merchantYearlyAmountByApikeyCacheKey = "merchant:statistic:yearly:amount:apikey:%s:year:%d"

	// Cache key for monthly total amount statistics by API key
	merchantMonthlyTotalAmountByApikeyCacheKey = "merchant:statistic:monthly:total-amount:apikey:%s:year:%d"

	// Cache key for yearly total amount statistics by API key
	merchantYearlyTotalAmountByApikeyCacheKey = "merchant:statistic:yearly:total-amount:apikey:%s:year:%d"

	ttlDefault = 5 * time.Minute
)
