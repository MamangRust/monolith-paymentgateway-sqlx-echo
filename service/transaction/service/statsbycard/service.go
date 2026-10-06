package transactionstatsbycardservice

import (
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-transaction/redis/statsbycard"
	repository "github.com/MamangRust/monolith-payment-gateway-transaction/repository/statsbycard"
)

type TransactionStatsByCardService interface {
	TransactionStatsByCardAmountService
	TransactionStatsByCardStatusService
	TransactionStatsByCardMethodService
}

type transactionStatsByCardService struct {
	TransactionStatsByCardAmountService
	TransactionStatsByCardStatusService
	TransactionStatsByCardMethodService
}

type DepsStats struct {
	Cache mencache.TransactionStatsByCardCache

	Repository repository.TransactionStatsByCardRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

func NewTransactionStatsByCardService(deps *DepsStats) TransactionStatsByCardService {
	return &transactionStatsByCardService{
		TransactionStatsByCardAmountService: NewTransactionStatsByCardAmountService(&transactionStatsByCardAmountServiceDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		TransactionStatsByCardStatusService: NewTransactionStatsByCardStatusService(&transactionStatsByCardStatusServiceDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		TransactionStatsByCardMethodService: NewTransactionStatsByCardMethodService(&transactionStatsByCardMethodServiceDeps{
			Cache:         deps.Cache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}
