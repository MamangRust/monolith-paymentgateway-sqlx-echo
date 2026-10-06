package cardpaymentmencache

import "time"

const (
    ttlDefault = 5 * time.Minute

    paymentHistoryCacheKey = "card:payment:history:%s:page:%d:pageSize:%d"
    paymentCountCacheKey   = "card:payment:count:%s"
)
