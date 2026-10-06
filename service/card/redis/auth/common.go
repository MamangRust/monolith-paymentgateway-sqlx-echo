package cardauthmencache

import "time"

const (
	ttlDefault = 5 * time.Minute

	authTxnByTxnIDCacheKey     = "card:auth:txn_id:%s"
	authTxnByCardNumCacheKey   = "card:auth:card_number:%s:page:%d:pageSize:%d"
)
