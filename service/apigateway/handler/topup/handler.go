package topuphandler

import (
	topup_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/topup"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup"
	pbstats "github.com/MamangRust/monolith-payment-gateway-pb/topup/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/topup"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type DepsTopup struct {
	Client *grpc.ClientConn

	E *echo.Echo

	Logger logger.LoggerInterface

	Cache *cache.CacheStore

	ApiHandler errors.ApiHandler
}

func RegisterTopupHandler(deps *DepsTopup) {
	mapper := apimapper.NewTopupResponseMapper()

	cache := topup_cache.NewTopupMencache(deps.Cache)

	handlers := []func(){
		setupTopupQueryHandler(deps, mapper.QueryMapper(), cache),
		setupTopupCommandHandler(deps, mapper.CommandMapper(), cache),
		setupTopupStatsMethodHandler(deps, mapper.MethodStatsMapper(), cache),
		setupTopupStatsStatusHandler(deps, mapper.StatusStatsMapper(), cache),
		setupTopupStatsAmountHandler(deps, mapper.AmountStatsMapper(), cache),
	}

	for _, h := range handlers {
		h()
	}
}

func setupTopupQueryHandler(deps *DepsTopup, mapper apimapper.TopupQueryResponseMapper, cache topup_cache.TopupMencach) func() {
	return func() {
		NewTopupQueryHandleApi(
			&topupQueryHandleDeps{
				client:     pb.NewTopupQueryServiceClient(deps.Client),
				router:     deps.E,
				logger:     deps.Logger,
				mapper:     mapper,
				cache:      cache,
				apiHandler: deps.ApiHandler,
			},
		)
	}
}

func setupTopupCommandHandler(deps *DepsTopup, mapper apimapper.TopupCommandResponseMapper, cache topup_cache.TopupMencach) func() {
	return func() {
		NewTopupCommandHandleApi(
			&topupCommandHandleDeps{
				client:     pb.NewTopupCommandServiceClient(deps.Client),
				router:     deps.E,
				logger:     deps.Logger,
				mapper:     mapper,
				apiHandler: deps.ApiHandler,
				cache:      cache,
			},
		)
	}
}

func setupTopupStatsMethodHandler(deps *DepsTopup, mapper apimapper.TopupStatsMethodResponseMapper, cache topup_cache.TopupMencach) func() {
	return func() {
		NewTopupStatsMethodHandleApi(
			&topupStatsMethodHandleDeps{
				client:     pbstats.NewTopupStatsMethodServiceClient(deps.Client),
				router:     deps.E,
				logger:     deps.Logger,
				mapper:     mapper,
				cache:      cache,
				apiHandler: deps.ApiHandler,
			},
		)
	}
}

func setupTopupStatsAmountHandler(deps *DepsTopup, mapper apimapper.TopupStatsAmountResponseMapper, cache topup_cache.TopupMencach) func() {
	return func() {
		NewTopupStatsAmountHandleApi(
			&topupStatsAmountHandleDeps{
				client:     pbstats.NewTopupStatsAmountServiceClient(deps.Client),
				router:     deps.E,
				logger:     deps.Logger,
				mapper:     mapper,
				cache:      cache,
				apiHandler: deps.ApiHandler,
			},
		)
	}
}

func setupTopupStatsStatusHandler(deps *DepsTopup, mapper apimapper.TopupStatsStatusResponseMapper, cache topup_cache.TopupMencach) func() {
	return func() {
		NewTopupStatsStatusHandleApi(
			&topupStatsStatusHandleDeps{
				client:     pbstats.NewTopupStatsStatusServiceClient(deps.Client),
				router:     deps.E,
				logger:     deps.Logger,
				mapper:     mapper,
				apiHandler: deps.ApiHandler,
				cache:      cache,
			},
		)
	}
}
