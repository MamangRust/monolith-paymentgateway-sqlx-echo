package transferstatsbycardservice

import (
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-transfer/redis/statsbycard"
	repository "github.com/MamangRust/monolith-payment-gateway-transfer/repository/statsbycard"
)

type TransferStatsByCardService interface {
	TransferStatsByCardAmountService
	TransferStatsByCardStatusService
}

type transferStatsByCardService struct {
	TransferStatsByCardAmountService
	TransferStatsByCardStatusService
}

type DepsStats struct {
	Cache         mencache.TransferStatsByCardCache
	Logger        logger.LoggerInterface
	Repository    repository.TransferStatsByCardRepository
	Observability observability.TraceLoggerObservability
}

func NewTransferStatsByCardService(deps *DepsStats) TransferStatsByCardService {
	return &transferStatsByCardService{
		TransferStatsByCardAmountService: NewTransferStatsByCardAmountService(&TransferStatsByCardAmountDeps{
			Cache:         deps.Cache,
			Sender:        deps.Repository,
			Receiver:      deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		TransferStatsByCardStatusService: NewTransferStatsByCardStatusService(&TransferStatsByCardStatusDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}

