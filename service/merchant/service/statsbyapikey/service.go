package merchantstatsbyapikeyservice

import (
	mencache "github.com/MamangRust/monolith-payment-gateway-merchant/redis/statsbyapikey"
	repository "github.com/MamangRust/monolith-payment-gateway-merchant/repository/statsbyapikey"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
)

type MerchantStatsByApiKeyService interface {
	MerchantStatsByApiKeyAmountService
	MerchantStatsByApiKeyTotalAmountService
	MerchantStatsByApiKeyMethodService
}

type merchantStatsByApiKeyService struct {
	MerchantStatsByApiKeyAmountService
	MerchantStatsByApiKeyTotalAmountService
	MerchantStatsByApiKeyMethodService
}

type DepsStatsByApiKey struct {
	Mencache      mencache.MerchantStatsByApiKeyCache
	Logger        logger.LoggerInterface
	Repository    repository.MerchantStatsByApiKeyRepository
	Observability observability.TraceLoggerObservability
}

func NewMerchantStatsByApiKeyService(deps *DepsStatsByApiKey) MerchantStatsByApiKeyService {
	return &merchantStatsByApiKeyService{
		NewMerchantStatsAmountByApiKeyService(&MerchantStatsAmountByApiKeyDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		NewMerchantStatsTotalAmountByApiKeyService(&MerchantStatsTotalAmountByApiKeyDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
		NewMerchantStatsMethodByApiKeyService(&MerchantStatsMethodByApiKeyDeps{
			Cache:         deps.Mencache,
			Repository:    deps.Repository,
			Logger:        deps.Logger,
			Observability: deps.Observability,
		}),
	}
}

