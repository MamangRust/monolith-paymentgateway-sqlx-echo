package merchantstatscache

import "time"

// Constants for merchant statistic cache keys.
//
// The keys are used to store the statistics data in the cache.
const (
	merchantMonthlyPaymentMethodCacheKey = "merchant:statistic:monthly:payment-method:year:%d"
	// merchantYearlyPaymentMethodCacheKey is the cache key for yearly
	// payment method statistics by merchant.
	merchantYearlyPaymentMethodCacheKey = "merchant:statistic:yearly:payment-method:year:%d"

	merchantMonthlyAmountCacheKey = "merchant:statistic:monthly:amount:year:%d"
	// merchantYearlyAmountCacheKey is the cache key for yearly amount
	// statistics by merchant.
	MerchantYearlyAmountCacheKey = "merchant:statistic:yearly:amount:year:%d"

	merchantMonthlyTotalAmountCacheKey = "merchant:statistic:monthly:total-amount:year:%d"
	// merchantYearlyTotalAmountCacheKey is the cache key for yearly total
	// amount statistics by merchant.
	merchantYearlyTotalAmountCacheKey = "merchant:statistic:yearly:total-amount:year:%d"

	ttlDefault = 5 * time.Minute
)
