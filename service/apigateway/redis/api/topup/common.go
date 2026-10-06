package topup_cache

import "time"

const (
	topupAllCacheKey     = "topup:all:page:%d:pageSize:%d:search:%s"
	topupByCardCacheKey  = "topup:card_number:%s:page:%d:pageSize:%d:search:%s"
	topupByIdCacheKey    = "topup:id:%d"
	topupActiveCacheKey  = "topup:active:page:%d:pageSize:%d:search:%s"
	topupTrashedCacheKey = "topup:trashed:page:%d:pageSize:%d:search:%s"

	ttlDefault = 5 * time.Minute
)
