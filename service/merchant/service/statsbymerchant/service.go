package merchantstatsbymerchantservice

import (
	mencache "github.com/MamangRust/monolith-payment-gateway-merchant/redis/statsbymerchant"
	repository "github.com/MamangRust/monolith-payment-gateway-merchant/repository/statsbymerchant"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
)

type MerchantStatsByMerchantService interface {
	MerchantStatsByMerchantAmountService
	MerchantStatsByMerchantTotalAmountService
	MerchantStatsByMerchantMethodService
}

type merchantStatsByMerchantService struct {
	MerchantStatsByMerchantAmountService
	MerchantStatsByMerchantTotalAmountService
	MerchantStatsByMerchantMethodService
}

type DepsStatsByMerchant struct {
	Mencache      mencache.MerchantStatsByMerchantCache
	Observability observability.TraceLoggerObservability
	Logger        logger.LoggerInterface
	Repository    repository.MerchantStatsByMerchantRepository
}

func NewMerchantStatsByMerchantService(deps *DepsStatsByMerchant) MerchantStatsByMerchantService {
	return &merchantStatsByMerchantService{
		NewMerchantStatsByMerchantAmountService(&MerchantStatsByMerchantAmountServiceDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		NewMerchantStatsByMerchantTotalAmountService(&MerchantStatsByMerchantTotalAmountDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		NewMerchantStatsByMerchantMethodService(&MerchantStatsByMerchantMethodDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}

