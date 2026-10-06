package carddashboardmencache

import "time"

// Constants for cache keys and TTLs
const (
	cacheKeyDashboardDefault    = "dashboard:card"
	cacheKeyDashboardCardNumber = "dashboard:card:number:%s"
	ttlDashboardDefault         = 5 * time.Minute
)
