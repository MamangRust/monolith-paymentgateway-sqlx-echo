package service

import (
	"fmt"

	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-topup/redis"
	"github.com/MamangRust/monolith-payment-gateway-topup/repository"
	topupstatsservice "github.com/MamangRust/monolith-payment-gateway-topup/service/stats"
	topupstatsbycardservice "github.com/MamangRust/monolith-payment-gateway-topup/service/statsbycard"
)

type Service interface {
	TopupQueryService
	TopupCommandService
	topupstatsservice.TopupStatsService
	topupstatsbycardservice.TopupStatsByCardService
}

type service struct {
	TopupQueryService
	TopupCommandService
	topupstatsservice.TopupStatsService
	topupstatsbycardservice.TopupStatsByCardService
}

type Deps struct {
	Kafka        *kafka.Kafka
	Cache        *cache.CacheStore
	Repositories repository.Repositories
	Logger       logger.LoggerInterface
}

func NewService(deps *Deps) Service {
	cache := mencache.NewMencache(deps.Cache)

	traceObservability, err := observability.NewObservability("topup-service", deps.Logger)
	if err != nil {
		// NewService cannot return an error without breaking the service
		// bootstrap contract. Fail fast instead of constructing services that
		// panic on their first request due to missing tracing dependencies.
		panic(fmt.Errorf("initialize topup observability: %w", err))
	}

	return &service{
		TopupQueryService:   newTopupQueryService(deps, traceObservability, cache),
		TopupCommandService: newTopupCommandService(deps, traceObservability, cache),
		TopupStatsService: topupstatsservice.NewTopupStatsService(&topupstatsservice.DepsStats{
			Cache:         cache,
			Logger:        deps.Logger,
			Repository:    deps.Repositories,
			Observability: traceObservability,
		}),
		TopupStatsByCardService: topupstatsbycardservice.NewTopupStatsByCardService(&topupstatsbycardservice.DepsStatsByCard{
			Cache:         cache,
			Logger:        deps.Logger,
			Repository:    deps.Repositories,
			Observability: traceObservability,
		}),
	}
}

func newTopupQueryService(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.TopupQueryCache) TopupQueryService {
	return NewTopupQueryService(&TopupQueryDeps{
		Cache:         cache,
		Repository:    deps.Repositories,
		Logger:        deps.Logger,
		Observability: observability,
	})
}

func newTopupCommandService(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.TopupCommandCache) TopupCommandService {
	return NewTopupCommandService(&TopupCommandDeps{
		Kafka:                  deps.Kafka,
		Cache:                  cache,
		CardRepository:         deps.Repositories,
		TopupQueryRepository:   deps.Repositories,
		TopupCommandRepository: deps.Repositories,
		SaldoRepository:        deps.Repositories,
		Logger:                 deps.Logger,
		Observability:          observability,
	})
}
