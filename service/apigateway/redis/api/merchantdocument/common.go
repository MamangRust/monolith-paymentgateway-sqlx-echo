package merchantdocument_cache

import "time"

const (
	merchantAllCacheKey = "merchant_document:all:page:%d:pageSize:%d:search:%s"

	merchantByIdCacheKey = "merchant_document:id:%d"

	merchantActiveCacheKey = "merchant_document:active:page:%d:pageSize:%d:search:%s"

	merchantTrashedCacheKey = "merchant_document:trashed:page:%d:pageSize:%d:search:%s"

	ttlDefault = 5 * time.Minute
)
