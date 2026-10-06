package transferstatsservice

import (
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-transfer/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-transfer/repository/stats"
)

type TransferStatsService interface {
	TransferStatsAmountService
	TransferStatsStatusService
}

type transferStatsService struct {
	TransferStatsAmountService
	TransferStatsStatusService
}

type DepsStats struct {
	Cache         mencache.TransferStatsCache
	Observability observability.TraceLoggerObservability
	Logger        logger.LoggerInterface
	Repository    repository.TransferStatsRepository
}

func NewTransferStatsService(deps *DepsStats) TransferStatsService {
	return &transferStatsService{
		TransferStatsAmountService: NewTransferStatsAmountService(&TransferStatsAmountDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		TransferStatsStatusService: NewTransferStatsStatusService(&TransferStatsStatusDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}

