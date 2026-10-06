package cardbillingmencache

import "time"

const (
    ttlDefault = 5 * time.Minute

    billingByCardCacheKey = "card:billing:card_number:%s"
)
