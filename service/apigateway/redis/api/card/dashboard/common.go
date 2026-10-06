package card_dashboard_cache

import "time"

const (
	cacheKeyDashboardDefault    = "dashboard:card"
	cacheKeyDashboardCardNumber = "dashboard:card:number:%s"
	ttlDashboardDefault         = 5 * time.Minute
)
