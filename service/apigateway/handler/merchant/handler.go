package merchanthandler

import (
	merchant_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/merchant"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	pbstats "github.com/MamangRust/monolith-payment-gateway-pb/merchant/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	errors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/merchant"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type DepsMerchant struct {
	Client *grpc.ClientConn
	E      *echo.Echo
	Logger logger.LoggerInterface
	Cache  *cache.CacheStore

	ApiHandler errors.ApiHandler
}

func RegisterMerchantHandler(deps *DepsMerchant) {
	mapper := apimapper.NewMerchantResponseMapper()

	cache := merchant_cache.NewMerchantMencache(deps.Cache)

	handlers := []func(){
		setupMerchantQueryHandler(deps, mapper.QueryMapper(), cache),
		setupMerchantCommandHandler(deps, mapper.CommandMapper(), cache),
		setupMerchantStatsAmountHandler(deps, mapper.AmountStatsMapper(), cache),
		setupMerchantStatsMethodHandler(deps, mapper.MethodStatsMapper(), cache),
		setupMerchantStatsTotalAmountHandler(deps, mapper.TotalAmountStatsMapper(), cache),
		setupMerchantTransactionHandler(deps, mapper.TransactionMapper(), cache),
	}

	for _, h := range handlers {
		h()
	}
}

func setupMerchantQueryHandler(deps *DepsMerchant, mapper apimapper.MerchantQueryResponseMapper, cache merchant_cache.MerchantMencache) func() {
	return func() {
		NewMerchantQueryHandleApi(&merchantQueryHandleDeps{
			client:     pb.NewMerchantQueryServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			cache:      cache,
			apiHandler: deps.ApiHandler,
		})
	}
}

func setupMerchantCommandHandler(deps *DepsMerchant, mapper apimapper.MerchantCommandResponseMapper, cache merchant_cache.MerchantMencache) func() {
	return func() {
		NewMerchantCommandHandleApi(&merchantCommandHandleDeps{
			client:     pb.NewMerchantCommandServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			cache:      cache,
			apiHandler: deps.ApiHandler,
		})
	}
}

func setupMerchantStatsAmountHandler(deps *DepsMerchant, mapper apimapper.MerchantStatsAmountResponseMapper, cache merchant_cache.MerchantMencache) func() {
	return func() {
		NewMerchantStatsAmountHandleApi(&merchantStatsAmountHandleDeps{
			client:     pbstats.NewMerchantStatsAmountServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			cache:      cache,
			apiHandler: deps.ApiHandler,
		})
	}
}

func setupMerchantStatsMethodHandler(deps *DepsMerchant, mapper apimapper.MerchantStatsMethodResponseMapper, cache merchant_cache.MerchantMencache) func() {
	return func() {
		NewMerchantStatsMethodHandleApi(&merchantStatsMethodHandleDeps{
			client:     pbstats.NewMerchantStatsMethodServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			cache:      cache,
			apiHandler: deps.ApiHandler,
		})
	}
}

func setupMerchantStatsTotalAmountHandler(deps *DepsMerchant, mapper apimapper.MerchantStatsTotalAmountResponseMapper, cache merchant_cache.MerchantMencache) func() {
	return func() {
		NewMerchantStatsTotalAmountHandleApi(&merchantStatsTotalAmountHandleDeps{
			client:     pbstats.NewMerchantStatsTotalAmountServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			cache:      cache,
			apiHandler: deps.ApiHandler,
		})
	}
}

func setupMerchantTransactionHandler(deps *DepsMerchant, mapper apimapper.MerchantTransactionResponseMapper, cache merchant_cache.MerchantMencache) func() {
	return func() {
		NewMerchantTransactionHandleApi(&merchantTransactionHandleDeps{
			client:     pb.NewMerchantTransactionServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			cache:      cache,
			apiHandler: deps.ApiHandler,
		})
	}
}
