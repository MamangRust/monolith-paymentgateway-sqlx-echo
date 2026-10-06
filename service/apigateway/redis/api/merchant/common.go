package merchant_cache

import "time"

const (
	merchantAllCacheKey = "merchant:all:page:%d:pageSize:%d:search:%s"

	merchantByIdCacheKey = "merchant:id:%d"

	merchantActiveCacheKey = "merchant:active:page:%d:pageSize:%d:search:%s"

	merchantTrashedCacheKey = "merchant:trashed:page:%d:pageSize:%d:search:%s"

	merchantByApiKeyCacheKey = "merchant:api_key:%s"

	merchantByUserIdCacheKey = "merchant:user_id:%d"

	merchantTransactionsCacheKey = "merchant:transaction:search:%s:page:%d:pageSize:%d"

	merchantTransactionApikeyCacheKey = "merchant:transaction:apikey:%s:search:%s:page:%d:pageSize:%d"

	merchantTransactionCacheKey = "merchant:transaction:merchant:%d:search:%s:page:%d:pageSize:%d"

	ttlDefault = 5 * time.Minute
)
