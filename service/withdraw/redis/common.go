package mencache

import "time"

const (
	withdrawAllCacheKey     = "withdraws:all:page:%d:pageSize:%d:search:%s"
	withdrawByCardCacheKey  = "withdraws:card_number:%s:page:%d:pageSize:%d:search:%s"
	withdrawByIdCacheKey    = "withdraws:id:%d"
	withdrawActiveCacheKey  = "withdraws:active:page:%d:pageSize:%d:search:%s"
	withdrawTrashedCacheKey = "withdraws:trashed:page:%d:pageSize:%d:search:%s"

	ttlDefault = 5 * time.Minute
)
