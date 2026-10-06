package card_dashboard_cache

import "github.com/MamangRust/monolith-payment-gateway-shared/cache"

type CardDashboardCache interface {
	CardDashboardTotalCache
	CardDashboardByCardNumberCache
}

type cardDashboardCaches struct {
	CardDashboardTotalCache
	CardDashboardByCardNumberCache
}

func NewMencacheDashboard(store *cache.CacheStore) CardDashboardCache {
	return &cardDashboardCaches{
		CardDashboardTotalCache:        NewCardDashboardCache(store),
		CardDashboardByCardNumberCache: NewCardDashboardByCardNumberCache(store),
	}
}
