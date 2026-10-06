package mencache

import (
	cardauthmencache "github.com/MamangRust/monolith-payment-gateway-card/redis/auth"
	cardbillingmencache "github.com/MamangRust/monolith-payment-gateway-card/redis/billing"
	carddashboardmencache "github.com/MamangRust/monolith-payment-gateway-card/redis/dashboard"
	cardpaymentmencache "github.com/MamangRust/monolith-payment-gateway-card/redis/payment"
	cardrewardmencache "github.com/MamangRust/monolith-payment-gateway-card/redis/reward"
	cardstatsmencache "github.com/MamangRust/monolith-payment-gateway-card/redis/stats"
	cardstatsbycardmencache "github.com/MamangRust/monolith-payment-gateway-card/redis/statsbycard"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/redis/go-redis/v9"

	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
)

type Mencache interface {
	CardQueryCache
	CardCommandCache
	cardstatsmencache.CardStatsCache
	cardstatsbycardmencache.CardStatsByCardCache
	carddashboardmencache.CardDashboardCache
	cardauthmencache.CardAuthCache
	cardpaymentmencache.CardPaymentCache
	cardrewardmencache.CardRewardCache
	cardbillingmencache.CardBillingCache
}

// Mencache is a struct that represents the cache store
type mencache struct {
	CardQueryCache
	CardCommandCache
	cardstatsmencache.CardStatsCache
	cardstatsbycardmencache.CardStatsByCardCache
	carddashboardmencache.CardDashboardCache
	cardauthmencache.CardAuthCache
	cardpaymentmencache.CardPaymentCache
	cardrewardmencache.CardRewardCache
	cardbillingmencache.CardBillingCache
}

// Deps is a struct that represents the dependencies needed to create a Mencache
type Deps struct {
	Redis   *redis.Client
	Logger  logger.LoggerInterface
	Metrics observability.CacheMetricsInterface
}

func NewMencache(cacheStore *cache.CacheStore) Mencache {
	return &mencache{
		CardCommandCache:     NewCardCommandCache(cacheStore),
		CardQueryCache:       NewCardQueryCache(cacheStore),
		CardStatsCache:       cardstatsmencache.NewMencacheStats(cacheStore),
		CardStatsByCardCache: cardstatsbycardmencache.NewMencacheStatsByCard(cacheStore),
		CardDashboardCache:   carddashboardmencache.NewMencacheDashboard(cacheStore),
		CardAuthCache:        cardauthmencache.NewCardAuthCache(cacheStore),
		CardPaymentCache:     cardpaymentmencache.NewCardPaymentCache(cacheStore),
		CardRewardCache:      cardrewardmencache.NewCardRewardCache(cacheStore),
		CardBillingCache:     cardbillingmencache.NewCardBillingCache(cacheStore),
	}
}
