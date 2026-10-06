package merchant_stats_bymerchant_cache

import "time"

const (
	merchantMonthlyPaymentMethodByMerchantCacheKey = "merchant:statistic:monthly:payment-method:merchant-id:%d:year:%d"

	merchantYearlyPaymentMethodByMerchantCacheKey = "merchant:statistic:yearly:payment-method:merchant-id:%d:year:%d"

	merchantMonthlyAmountByMerchantCacheKey = "merchant:statistic:monthly:amount:merchant-id:%d:year:%d"

	merchantYearlyAmountByMerchantCacheKey = "merchant:statistic:yearly:amount:merchant-id:%d:year:%d"

	merchantMonthlyTotalAmountByMerchantCacheKey = "merchant:statistic:monthly:total-amount:merchant-id:%d:year:%d"

	merchantYearlyTotalAmountByMerchantCacheKey = "merchant:statistic:yearly:total-amount:merchant-id:%d:year:%d"

	ttlDefault = 5 * time.Minute
)
