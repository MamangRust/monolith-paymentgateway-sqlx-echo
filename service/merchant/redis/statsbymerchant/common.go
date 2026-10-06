package merchantstatsbymerchant

import "time"

// Cache keys for merchant statistics by merchant ID and year.
const (
	// Key for monthly payment method statistics by merchant.
	merchantMonthlyPaymentMethodByMerchantCacheKey = "merchant:statistic:monthly:payment-method:merchant-id:%d:year:%d"

	// Key for yearly payment method statistics by merchant.
	merchantYearlyPaymentMethodByMerchantCacheKey = "merchant:statistic:yearly:payment-method:merchant-id:%d:year:%d"

	// Key for monthly amount statistics by merchant.
	merchantMonthlyAmountByMerchantCacheKey = "merchant:statistic:monthly:amount:merchant-id:%d:year:%d"

	// Key for yearly amount statistics by merchant.
	merchantYearlyAmountByMerchantCacheKey = "merchant:statistic:yearly:amount:merchant-id:%d:year:%d"

	// Key for monthly total amount statistics by merchant.
	merchantMonthlyTotalAmountByMerchantCacheKey = "merchant:statistic:monthly:total-amount:merchant-id:%d:year:%d"

	// Key for yearly total amount statistics by merchant.
	merchantYearlyTotalAmountByMerchantCacheKey = "merchant:statistic:yearly:total-amount:merchant-id:%d:year:%d"

	ttlDefault = 5 * time.Minute
)
