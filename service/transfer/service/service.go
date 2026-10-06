package service

import (
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"

	mencache "github.com/MamangRust/monolith-payment-gateway-transfer/redis"
	"github.com/MamangRust/monolith-payment-gateway-transfer/repository"
	transferstatsservice "github.com/MamangRust/monolith-payment-gateway-transfer/service/stats"
	transferstatsbycardservice "github.com/MamangRust/monolith-payment-gateway-transfer/service/statsbycard"
)

type Service interface {
	TransferQueryService
	TransferCommandService
	transferstatsservice.TransferStatsService
	transferstatsbycardservice.TransferStatsByCardService
}

type service struct {
	TransferQueryService
	TransferCommandService
	transferstatsservice.TransferStatsService
	transferstatsbycardservice.TransferStatsByCardService
}

type Deps struct {
	Kafka        *kafka.Kafka
	Cache        *cache.CacheStore
	Repositories repository.Repositories
	Logger       logger.LoggerInterface
}

func NewService(deps *Deps) Service {
	cache := mencache.NewMencache(deps.Cache)

	observability, _ := observability.NewObservability("transfer-service", deps.Logger)

	return &service{
		TransferQueryService:   newTransferQueryService(deps, observability, cache),
		TransferCommandService: newTransferCommandService(deps, observability, cache),
		TransferStatsService: transferstatsservice.NewTransferStatsService(&transferstatsservice.DepsStats{
			Cache:         cache,
			Logger:        deps.Logger,
			Repository:    deps.Repositories,
			Observability: observability,
		}),
		TransferStatsByCardService: transferstatsbycardservice.NewTransferStatsByCardService(&transferstatsbycardservice.DepsStats{
			Observability: observability,
			Logger:        deps.Logger,
			Repository:    deps.Repositories,
			Cache:         cache,
		}),
	}
}

func newTransferQueryService(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.Mencache) TransferQueryService {
	return NewTransferQueryService(
		&TransferQueryDeps{
			Cache:         cache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: observability,
		},
	)
}

func newTransferCommandService(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.Mencache) TransferCommandService {
	return NewTransferCommandService(
		&TransferCommandDeps{
			Kafka:                     deps.Kafka,
			Cache:                     cache,
			CardRepository:            deps.Repositories,
			SaldoRepository:           deps.Repositories,
			TransferQueryRepository:   deps.Repositories,
			TransferCommandRepository: deps.Repositories,
			Logger:                    deps.Logger,
			Observability:             observability,
		},
	)
}
