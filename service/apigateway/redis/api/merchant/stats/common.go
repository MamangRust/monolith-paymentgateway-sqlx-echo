package merchant_stats_cache

import "time"

const (
	merchantMonthlyPaymentMethodCacheKey = "merchant:statistic:monthly:payment-method:year:%d"
	merchantYearlyPaymentMethodCacheKey  = "merchant:statistic:yearly:payment-method:year:%d"

	merchantMonthlyAmountCacheKey = "merchant:statistic:monthly:amount:year:%d"
	MerchantYearlyAmountCacheKey  = "merchant:statistic:yearly:amount:year:%d"

	merchantMonthlyTotalAmountCacheKey = "merchant:statistic:monthly:total-amount:year:%d"
	merchantYearlyTotalAmountCacheKey  = "merchant:statistic:yearly:total-amount:year:%d"

	ttlDefault = 5 * time.Minute
)
