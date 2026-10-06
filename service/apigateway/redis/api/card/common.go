package card_cache

import "time"

const (
	ttlDefault = 5 * time.Minute

	cardAllCacheKey       = "card:all:page:%d:pageSize:%d:search:%s"
	cardByIdCacheKey      = "card:id:%d"
	cardActiveCacheKey    = "card:active:page:%d:pageSize:%d:search:%s"
	cardTrashedCacheKey   = "card:trashed:page:%d:pageSize:%d:search:%s"
	cardByUserIdCacheKey  = "card:user_id:%d"
	cardByCardNumCacheKey = "card:card_number:%s"
)
