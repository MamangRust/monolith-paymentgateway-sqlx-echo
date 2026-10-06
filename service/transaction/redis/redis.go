package mencache

import (
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	transactionstatscache "github.com/MamangRust/monolith-payment-gateway-transaction/redis/stats"
	transactionstatsbycarcache "github.com/MamangRust/monolith-payment-gateway-transaction/redis/statsbycard"
	"github.com/redis/go-redis/v9"
)

type Mencache interface {
	TransactionQueryCache
	TransactionCommandCache
	transactionstatscache.TransactionStatsCache
	transactionstatsbycarcache.TransactionStatsByCardCache
}

// Mencache represents a cache store for transaction queries, commands, statistics, and statistics by card.
type mencache struct {
	TransactionQueryCache
	TransactionCommandCache
	transactionstatscache.TransactionStatsCache
	transactionstatsbycarcache.TransactionStatsByCardCache
}

// Deps represents the dependencies required to initialize a Mencache.
type Deps struct {
	Redis   *redis.Client
	Logger  logger.LoggerInterface
	Metrics observability.CacheMetricsInterface
}

func NewMencache(cacheStore *cache.CacheStore) Mencache {
	return &mencache{
		TransactionQueryCache:       NewTransactionQueryCache(cacheStore),
		TransactionCommandCache:     NewTransactionCommandCache(cacheStore),
		TransactionStatsCache:       transactionstatscache.NewTransactionStatsCache(cacheStore),
		TransactionStatsByCardCache: transactionstatsbycarcache.NewTransactionStatsByCardCache(cacheStore),
	}
}
