package mencache

import "time"

const (
	ttlDefault = 5 * time.Minute

	cacheMerchantKey = "merchant_api_key:%s"
	cacheRoleKey     = "user_roles:%s"
)
