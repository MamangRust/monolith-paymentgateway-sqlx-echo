package transferhandler

import (
	transfer_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/transfer"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transfer"
	pbstats "github.com/MamangRust/monolith-payment-gateway-pb/transfer/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/transfer"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type DepsTransfer struct {
	Client *grpc.ClientConn

	E *echo.Echo

	Logger logger.LoggerInterface

	Cache *cache.CacheStore

	ApiHandler errors.ApiHandler
}

func RegisterTransferHandler(deps *DepsTransfer) {
	mapper := apimapper.NewTransferResponseMapper()
	cache := transfer_cache.NewTransferMencache(deps.Cache)

	handlers := []func(){
		setupTransferQueryHandler(deps, mapper.QueryMapper(), cache),
		setupTransferCommandHandler(deps, mapper.CommandMapper(), cache),
		setupTransferStatsAmountHandler(deps, mapper.AmountStatsMapper(), cache),
		setupTransferStatsStatusHandler(deps, mapper.StatusStatsMapper(), cache),
	}

	for _, h := range handlers {
		h()
	}
}

func setupTransferQueryHandler(deps *DepsTransfer, mapper apimapper.TransferQueryResponseMapper, cache transfer_cache.TransferMencache) func() {
	return func() {
		NewTransferQueryHandleApi(&transferQueryHandleDeps{
			client:     pb.NewTransferQueryServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			cache:      cache,
			apiHandler: deps.ApiHandler,
		})
	}
}

func setupTransferCommandHandler(deps *DepsTransfer, mapper apimapper.TransferCommandResponseMapper, cache transfer_cache.TransferMencache) func() {
	return func() {
		NewTransferCommandHandleApi(&transferCommandHandleDeps{
			client:     pb.NewTransferCommandServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			cache:      cache,
			apiHandler: deps.ApiHandler,
		})
	}
}

func setupTransferStatsAmountHandler(deps *DepsTransfer, mapper apimapper.TransferStatsAmountResponseMapper, cache transfer_cache.TransferMencache) func() {
	return func() {
		NewTransferStatsAmountHandleApi(&transferStatsAmountHandleDeps{
			client:     pbstats.NewTransferStatsAmountServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			cache:      cache,
			apiHandler: deps.ApiHandler,
		})
	}
}

func setupTransferStatsStatusHandler(deps *DepsTransfer, mapper apimapper.TransferStatsStatusResponseMapper, cache transfer_cache.TransferMencache) func() {
	return func() {
		NewTransferStatsStatusHandleApi(&transferStatsStatusHandleDeps{
			client:     pbstats.NewTransferStatsStatusServiceClient(deps.Client),
			router:     deps.E,
			logger:     deps.Logger,
			mapper:     mapper,
			cache:      cache,
			apiHandler: deps.ApiHandler,
		})
	}
}
