package cardstatsservice

import (
	mencache "github.com/MamangRust/monolith-payment-gateway-card/redis/stats"
	repositorystats "github.com/MamangRust/monolith-payment-gateway-card/repository/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
)

type CardStatsService interface {
	CardStatsBalanceService
	CardStatsTopupService
	CardStatsWithdrawService
	CardStatsTransferService
	CardStatsTransactionService
}

type cardStatsService struct {
	CardStatsBalanceService
	CardStatsTopupService
	CardStatsWithdrawService
	CardStatsTransferService
	CardStatsTransactionService
}

type DepsStats struct {
	Mencache      mencache.CardStatsCache
	Repositories  repositorystats.CardStatsRepository
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
}

func NewCardStatsService(deps *DepsStats) CardStatsService {
	return &cardStatsService{
		NewCardStatsBalanceService(&cardStatsBalanceServiceDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		NewCardStatsTopupService(&cardStatsTopupServiceDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		NewCardStatsWithdrawService(&cardStatsWithdrawServiceDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		NewCardStatsTransferService(&cardStatsTransferServiceDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		NewCardStatsTransactionService(&cardStatsTransactionServiceDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}
