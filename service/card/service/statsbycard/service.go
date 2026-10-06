package cardstatsbycard

import (
	mencache "github.com/MamangRust/monolith-payment-gateway-card/redis/statsbycard"
	repositorystats "github.com/MamangRust/monolith-payment-gateway-card/repository/statsbycard"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
)

type CardStatsByCardService interface {
	CardStatsBalanceByCardService
	CardStatsTopupByCardService
	CardStatsWithdrawByCardService
	CardStatsTransferByCardService
	CardStatsTransactionByCardService
}

type cardStatsByCardService struct {
	CardStatsBalanceByCardService
	CardStatsTopupByCardService
	CardStatsWithdrawByCardService
	CardStatsTransferByCardService
	CardStatsTransactionByCardService
}

type DepsStatsByCard struct {
	Mencache      mencache.CardStatsByCardCache
	Repositories  repositorystats.CardStatsByCardRepository
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
}

func NewCardStatsByCardService(deps *DepsStatsByCard) CardStatsByCardService {
	return &cardStatsByCardService{
		NewCardStatsBalanceByCardService(&cardStatsBalanceByCardServiceDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		NewCardStatsTopupByCardService(&cardStatsTopupByCardServiceDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		NewCardStatsWithdrawByCardService(&cardStatsWithdrawByCardServiceDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		NewCardStatsTransferByCardService(&cardStatsTransferByCardServiceDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		NewCardStatsTransactionByCardService(&cardStatsTransactionByCardServiceDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}
