package withdrawstatsservice

import (
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-withdraw/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-withdraw/repository/stats"
)

type WithdrawStatsService interface {
	WithdrawStatsAmountService
	WithdrawStatsStatusService
}

type withdrawStatsStatsService struct {
	WithdrawStatsAmountService
	WithdrawStatsStatusService
}

type DepsStats struct {
	Cache mencache.WithdrawStatsCache

	Logger logger.LoggerInterface

	Repository    repository.WithdrawStatsRepository
	Observability observability.TraceLoggerObservability
}

func NewWithdrawStatsService(deps *DepsStats) WithdrawStatsService {
	return &withdrawStatsStatsService{
		WithdrawStatsAmountService: NewWithdrawStatsAmountService(&WithdrawStatsAmountDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		WithdrawStatsStatusService: NewWithdrawStatsStatusService(&WithdrawStatsStatusDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}

