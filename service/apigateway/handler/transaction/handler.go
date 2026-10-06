package transactionhandler

import (
	mencache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis"
	transaction_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/transaction"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction"
	pbstats "github.com/MamangRust/monolith-payment-gateway-pb/transaction/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/transaction"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type DepsTransaction struct {
	Client *grpc.ClientConn

	E *echo.Echo

	Kafka *kafka.Kafka

	Logger logger.LoggerInterface

	CacheApiGateway mencache.CacheApiGateway

	Cache *cache.CacheStore

	ApiHandler errors.ApiHandler
}

func RegisterTransactionHandler(deps *DepsTransaction) {
	mapper := apimapper.NewTransactionResponseMapper()

	cache := transaction_cache.NewTransactionMencache(deps.Cache)

	handlers := []func(){
		setupTransactionQueryHandler(deps, mapper.QueryMapper(), cache),
		setupTransactionCommandHandler(deps, deps.CacheApiGateway, mapper.CommandMapper(), cache),
		setupTransactionStatsAmountHandler(deps, mapper.AmountStatsMapper(), cache),
		setupTransactionStatsStatusHandler(deps, mapper.StatusStatsMapper(), cache),
		setupTransactionStatsMethodHandler(deps, mapper.MethodStatsMapper(), cache),
	}

	for _, h := range handlers {
		h()
	}
}

func setupTransactionQueryHandler(deps *DepsTransaction, mapper apimapper.TransactionQueryResponseMapper, cache transaction_cache.TransactionMencache) func() {
	return func() {
		NewTransactionQueryHandleApi(&transactionQueryHandleDeps{
			client:     pb.NewTransactionQueryServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			cache:      cache,
			apiHandler: deps.ApiHandler,
		})
	}
}

func setupTransactionCommandHandler(deps *DepsTransaction, cache mencache.MerchantCache, mapper apimapper.TransactionCommandResponseMapper, cache_ transaction_cache.TransactionMencache) func() {
	return func() {
		NewTransactionCommandHandleApi(&transactionCommandHandleDeps{
			kafka:      deps.Kafka,
			client:     pb.NewTransactionCommandServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			cache:             cache,
			cache_transaction: cache_,
			apiHandler: deps.ApiHandler,
		})
	}
}

func setupTransactionStatsAmountHandler(deps *DepsTransaction, mapper apimapper.TransactionStatsAmountResponseMapper, cache transaction_cache.TransactionMencache) func() {
	return func() {
		NewTransactionStatsAmountHandleApi(&transactionStatsAmountHandleDeps{
			client:     pbstats.NewTransactionStatsAmountServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			cache:      cache,
			apiHandler: deps.ApiHandler,
		})
	}
}

func setupTransactionStatsMethodHandler(deps *DepsTransaction, mapper apimapper.TransactionStatsMethodResponseMapper, cache transaction_cache.TransactionMencache) func() {
	return func() {
		NewTransactionStatsMethodHandleApi(&transactionStatsMethodHandleDeps{
			client:     pbstats.NewTransactionStatsMethodServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			cache:      cache,
			apiHandler: deps.ApiHandler,
		})
	}
}

func setupTransactionStatsStatusHandler(deps *DepsTransaction, mapper apimapper.TransactionStatsStatusResponseMapper, cache transaction_cache.TransactionMencache) func() {
	return func() {
		NewTransactionStatsStatusHandleApi(&transactionStatsStatusHandleDeps{
			client:     pbstats.NewTransactionStatsStatusServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			apiHandler: deps.ApiHandler,
			cache:      cache,
		})
	}
}
