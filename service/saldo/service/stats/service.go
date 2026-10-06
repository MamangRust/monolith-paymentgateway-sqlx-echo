package saldostatsservice

import (
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	mencache "github.com/MamangRust/monolith-payment-gateway-saldo/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-saldo/repository/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
)

type SaldoStatsService interface {
	SaldoStatsBalanceService
	SaldoStatsTotalBalanceService
}

type saldoStatsService struct {
	SaldoStatsBalanceService
	SaldoStatsTotalBalanceService
}

type DepsStats struct {
	Mencache mencache.SaldoStatsCache

	Logger logger.LoggerInterface

	Repository repository.SaldoStatsRepository

	Observability observability.TraceLoggerObservability
}

func NewSaldoStatsService(deps *DepsStats) SaldoStatsService {
	return &saldoStatsService{
		NewSaldoStatsBalanceService(&saldoStatsBalanceDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		NewSaldoStatsTotalBalanceService(&saldoStatsTotalBalanceDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}
