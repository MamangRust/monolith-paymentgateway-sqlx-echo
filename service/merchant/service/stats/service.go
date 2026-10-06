package merchantstatsservice

import (
	mencache "github.com/MamangRust/monolith-payment-gateway-merchant/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-merchant/repository/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
)

type MerchantStatsService interface {
	MerchantStatsAmountService
	MerchantStatsTotalAmountService
	MerchantStatsMethodService
}

type merchantStatsService struct {
	MerchantStatsAmountService
	MerchantStatsTotalAmountService
	MerchantStatsMethodService
}

type DepsStats struct {
	Mencache      mencache.MerchantStatsCache
	Logger        logger.LoggerInterface
	Repository    repository.MerchantStatsRepository
	Observability observability.TraceLoggerObservability
}

func NewMerchantStatsService(deps *DepsStats) MerchantStatsService {
	return &merchantStatsService{
		MerchantStatsAmountService: NewMerchantStatsAmountService(&MerchantStatsAmountDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		MerchantStatsTotalAmountService: NewMerchantStatsTotalAmountService(&MerchantStatsTotalAmountDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		MerchantStatsMethodService: NewMerchantStatsMethodService(&MerchantStatsMethodDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}

