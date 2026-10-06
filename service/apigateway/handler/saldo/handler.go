package saldohandler

import (
	saldo_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/saldo"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	pbstats "github.com/MamangRust/monolith-payment-gateway-pb/saldo/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/saldo"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type DepsSaldo struct {
	Client *grpc.ClientConn

	E *echo.Echo

	Logger logger.LoggerInterface

	Cache *cache.CacheStore

	ApiHandler errors.ApiHandler
}

func RegisterSaldoHandler(deps *DepsSaldo) {
	mapper := apimapper.NewSaldoResponseMapper()

	cache := saldo_cache.NewSaldoMencache(deps.Cache)

	handlers := []func(){
		setupSaldoQueryHandler(deps, mapper.QueryMapper(), cache),
		setupSaldoCommandHandler(deps, mapper.CommandMapper(), cache),
		setupSaldoStatsBalanceHandler(deps, mapper.BalanceStatsMapper(), cache),
		setupStatsSaldoTotalBalanceHandler(deps, mapper.TotalStatsMapper(), cache),
	}

	for _, h := range handlers {
		h()
	}
}

func setupSaldoQueryHandler(deps *DepsSaldo, mapper apimapper.SaldoQueryResponseMapper, cache saldo_cache.SaldoMencache) func() {
	return func() {
		NewSaldoQueryHandleApi(
			&saldoQueryHandleDeps{
				client:     pb.NewSaldoQueryServiceClient(deps.Client),
				router:     deps.E,
				logger:     deps.Logger,
				mapper:     mapper,
				cache:      cache,
				apiHandler: deps.ApiHandler,
			},
		)
	}
}

func setupSaldoCommandHandler(deps *DepsSaldo, mapper apimapper.SaldoCommandResponseMapper, cache saldo_cache.SaldoMencache) func() {
	return func() {
		NewSaldoCommandHandleApi(
			&saldoCommandHandleDeps{
				client:     pb.NewSaldoCommandServiceClient(deps.Client),
				router:     deps.E,
				logger:     deps.Logger,
				mapper:     mapper,
				apiHandler: deps.ApiHandler,
				cache:      cache,
			},
		)
	}
}

func setupSaldoStatsBalanceHandler(deps *DepsSaldo, mapper apimapper.SaldoStatsBalanceResponseMapper, cache saldo_cache.SaldoMencache) func() {
	return func() {
		NewSaldoStatsBalanceHandleApi(
			&saldoStatsBalanceHandleDeps{
				client:     pbstats.NewSaldoStatsBalanceServiceClient(deps.Client),
				router:     deps.E,
				logger:     deps.Logger,
				mapper:     mapper,
				apiHandler: deps.ApiHandler,
				cache:      cache,
			},
		)
	}
}

func setupStatsSaldoTotalBalanceHandler(deps *DepsSaldo, mapper apimapper.SaldoStatsTotalResponseMapper, cache saldo_cache.SaldoMencache) func() {
	return func() {
		NewSaldoTotalBalanceHandleApi(
			&saldoTotalBalanceHandleDeps{
				client:     pbstats.NewSaldoStatsTotalBalanceClient(deps.Client),
				router:     deps.E,
				logger:     deps.Logger,
				mapper:     mapper,
				cache:      cache,
				apiHandler: deps.ApiHandler,
			},
		)
	}
}
