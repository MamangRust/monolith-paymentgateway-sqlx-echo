package cardrewardmencache

import "time"

const (
    ttlDefault = 5 * time.Minute

    rewardBalanceCacheKey = "card:reward:balance:%s"
    rewardHistoryCacheKey = "card:reward:history:%s"
)
