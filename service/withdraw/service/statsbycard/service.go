package withdrawstatsbycardservice

import (
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-withdraw/redis/statsbycard"
	repository "github.com/MamangRust/monolith-payment-gateway-withdraw/repository/statsbycard"
)

type WithdrawStatsByCardService interface {
	WithdrawStatsByCardAmountService
	WithdrawStatsByCardStatusService
}

type withdrawStatsByCardStatsByCardService struct {
	WithdrawStatsByCardAmountService
	WithdrawStatsByCardStatusService
}

type DepsStatsByCard struct {
	Cache mencache.WithdrawStatsByCardCache

	Logger logger.LoggerInterface

	Repository repository.WithdrawStatsByCardRepository

	Observability observability.TraceLoggerObservability
}

func NewWithdrawStatsByCardService(deps *DepsStatsByCard) WithdrawStatsByCardService {
	return &withdrawStatsByCardStatsByCardService{
		WithdrawStatsByCardAmountService: NewWithdrawStatsByCardAmountService(&WithdrawStatsByCardAmountDeps{
			Repository:    deps.Repository,
			Cache:         deps.Cache,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		WithdrawStatsByCardStatusService: NewWithdrawStatsByCardStatusService(&WithdrawStatsByCardStatusDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}

