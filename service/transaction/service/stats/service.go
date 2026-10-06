package transactionstatsservice

import (
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-transaction/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-transaction/repository/stats"
)

type TransactionStatsService interface {
	TransactionStatsAmountService
	TransactionStatsStatusService
	TransactionStatsMethodService
}

type transactionStatsService struct {
	TransactionStatsAmountService
	TransactionStatsStatusService
	TransactionStatsMethodService
}

type DepsStats struct {
	Cache mencache.TransactionStatsCache

	Repository repository.TransactionStatsRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

func NewTransactionStatsService(deps *DepsStats) TransactionStatsService {
	return &transactionStatsService{
		TransactionStatsAmountService: NewTransactionStatsAmountService(&transactionStatsAmountServiceDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		TransactionStatsStatusService: NewTransactionStatsStatusService(&transactionStatsStatusServiceDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		TransactionStatsMethodService: NewTransactionStatsMethodService(&transactionStatsMethodServiceDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}
