package service

import (
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-transaction/redis"
	"github.com/MamangRust/monolith-payment-gateway-transaction/repository"
	transactionstatsservice "github.com/MamangRust/monolith-payment-gateway-transaction/service/stats"
	transactionstatsbycardservice "github.com/MamangRust/monolith-payment-gateway-transaction/service/statsbycard"
)

// Service is a struct that contains all the services
type service struct {
	TransactionQueryService
	TransactionCommandService
	transactionstatsservice.TransactionStatsService
	transactionstatsbycardservice.TransactionStatsByCardService
}

type Service interface {
	TransactionQueryService
	TransactionCommandService
	transactionstatsservice.TransactionStatsService
	transactionstatsbycardservice.TransactionStatsByCardService
}

type Deps struct {
	Kafka        *kafka.Kafka
	Repositories repository.Repositories
	Logger       logger.LoggerInterface
	Cache        *cache.CacheStore
}

func NewService(deps *Deps) Service {
	cache := mencache.NewMencache(deps.Cache)
	observability, _ := observability.NewObservability("transaction-service", deps.Logger)

	return &service{
		TransactionQueryService:   newTransactionQueryService(deps, observability, cache),
		TransactionCommandService: newTransactionCommandService(deps, observability, cache),
		TransactionStatsService: transactionstatsservice.NewTransactionStatsService(&transactionstatsservice.DepsStats{
			Cache:         cache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: observability,
		}),
		TransactionStatsByCardService: transactionstatsbycardservice.NewTransactionStatsByCardService(&transactionstatsbycardservice.DepsStats{
			Cache:         cache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: observability,
		}),
	}
}

func newTransactionQueryService(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.Mencache) TransactionQueryService {
	return NewTransactionQueryService(&TransactionQueryServiceDeps{
		Cache:                      cache,
		TransactionQueryRepository: deps.Repositories,
		Logger:                     deps.Logger,
		Observability:              observability,
	})
}

func newTransactionCommandService(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.Mencache) TransactionCommandService {
	return NewTransactionCommandService(&TransactionCommandServiceDeps{
		Kafka:                        deps.Kafka,
		Mencache:                     cache,
		MerchantRepository:           deps.Repositories,
		CardRepository:               deps.Repositories,
		SaldoRepository:              deps.Repositories,
		TransactionCommandRepository: deps.Repositories,
		TransactionQueryRepository:   deps.Repositories,
		Logger:                       deps.Logger,
		Observability:                observability,
	})
}
