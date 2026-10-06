package service

import (
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-withdraw/redis"
	"github.com/MamangRust/monolith-payment-gateway-withdraw/repository"
	withdrawstatsservice "github.com/MamangRust/monolith-payment-gateway-withdraw/service/stats"
	withdrawstatsbycardservice "github.com/MamangRust/monolith-payment-gateway-withdraw/service/statsbycard"
)

type Service interface {
	WithdrawQueryService
	WithdrawCommandService
	withdrawstatsservice.WithdrawStatsService
	withdrawstatsbycardservice.WithdrawStatsByCardService
}

type service struct {
	WithdrawQueryService
	WithdrawCommandService
	withdrawstatsservice.WithdrawStatsService
	withdrawstatsbycardservice.WithdrawStatsByCardService
}

type Deps struct {
	Kafka        *kafka.Kafka
	Repositories repository.Repositories
	Logger       logger.LoggerInterface
	Cache        *cache.CacheStore
}

func NewService(deps *Deps) Service {

	cache := mencache.NewMencache(deps.Cache)

	observability, _ := observability.NewObservability("withdraw-service", deps.Logger)

	return &service{
		WithdrawQueryService:   newWithdrawQueryService(deps, observability, cache),
		WithdrawCommandService: newWithdrawCommandService(deps, observability, cache),
		WithdrawStatsService: withdrawstatsservice.NewWithdrawStatsService(&withdrawstatsservice.DepsStats{
			Cache:         cache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: observability,
		}),
		WithdrawStatsByCardService: withdrawstatsbycardservice.NewWithdrawStatsByCardService(&withdrawstatsbycardservice.DepsStatsByCard{
			Cache:         cache,
			Logger:        deps.Logger,
			Repository:    deps.Repositories,
			Observability: observability,
		}),
	}
}

func newWithdrawQueryService(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.Mencache) WithdrawQueryService {
	return NewWithdrawQueryService(
		&WithdrawQueryServiceDeps{
			Cache:         cache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: observability,
		},
	)
}

func newWithdrawCommandService(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.Mencache) WithdrawCommandService {
	return NewWithdrawCommandService(
		&WithdrawCommandServiceDeps{
			Cache:             cache,
			Kafka:             deps.Kafka,
			CardRepository:    deps.Repositories,
			SaldoRepository:   deps.Repositories,
			CommandRepository: deps.Repositories,
			QueryRepository:   deps.Repositories,
			Logger:            deps.Logger,
			Observability:     observability,
		},
	)
}
