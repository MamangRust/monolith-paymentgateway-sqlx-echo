package topupstatsservice

import (
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-topup/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-topup/repository/stats"
)

type TopupStatsService interface {
	TopupStatsAmountService
	TopupStatsMethodService
	TopupStatsStatusService
}

type topupStatsService struct {
	TopupStatsAmountService
	TopupStatsMethodService
	TopupStatsStatusService
}

type DepsStats struct {
	Cache mencache.TopupStatsCache

	Logger        logger.LoggerInterface
	Repository    repository.TopupStatsRepository
	Observability observability.TraceLoggerObservability
}

func NewTopupStatsService(deps *DepsStats) TopupStatsService {
	return &topupStatsService{
		TopupStatsAmountService: NewTopupStatsAmountService(&topupStatsAmountDeps{
			Cache:         deps.Cache,
			Observability: deps.Observability,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
		}),
		TopupStatsMethodService: NewTopupStatsMethodService(&topupStatsMethodDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		TopupStatsStatusService: NewTopupStatsStatusService(&topupStatsStatusDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}
