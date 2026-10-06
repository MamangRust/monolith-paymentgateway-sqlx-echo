package mencache

import "time"

const (
	transferAllCacheKey     = "transfer:all:page:%d:pageSize:%d:search:%s"
	transferByIdCacheKey    = "transfer:id:%d"
	transferActiveCacheKey  = "transfer:active:page:%d:pageSize:%d:search:%s"
	transferTrashedCacheKey = "transfer:trashed:page:%d:pageSize:%d:search:%s"

	transferByFromCacheKey = "transfer:from_card_number:%s:"
	transferByToCacheKey   = "transfer:to_card_number:%s"

	ttlDefault = 5 * time.Minute
)
