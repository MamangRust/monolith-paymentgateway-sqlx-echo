package topupstatsbycardservice

import (
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-topup/redis/statsbycard"
	repository "github.com/MamangRust/monolith-payment-gateway-topup/repository/statsbycard"
)

type TopupStatsByCardService interface {
	TopupStatsByCardAmountService
	TopupStatsByCardMethodService
	TopupStatsByCardStatusService
}

type topupStatsByCardService struct {
	TopupStatsByCardAmountService
	TopupStatsByCardMethodService
	TopupStatsByCardStatusService
}

type DepsStatsByCard struct {
	Cache mencache.TopupStatsByCardCache

	Logger        logger.LoggerInterface
	Repository    repository.TopupStatsByCardRepository
	Observability observability.TraceLoggerObservability
}

func NewTopupStatsByCardService(deps *DepsStatsByCard) TopupStatsByCardService {
	return &topupStatsByCardService{
		TopupStatsByCardAmountService: NewTopupStatsByCardAmountService(&topupStatsByCardAmountDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		TopupStatsByCardMethodService: NewTopupStatsByCardMethodService(&topupStatsByCardMethodDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		TopupStatsByCardStatusService: NewTopupStatsByCardStatusService(&topupStatsByCardStatusDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}
