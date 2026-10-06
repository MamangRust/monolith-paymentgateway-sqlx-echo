package service

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"

	mencache "github.com/MamangRust/monolith-payment-gateway-merchant/redis"
	"github.com/MamangRust/monolith-payment-gateway-merchant/repository"
	merchantstatsservice "github.com/MamangRust/monolith-payment-gateway-merchant/service/stats"
	merchantstatsbyapikeyservice "github.com/MamangRust/monolith-payment-gateway-merchant/service/statsbyapikey"
	merchantstatsbymerchantservice "github.com/MamangRust/monolith-payment-gateway-merchant/service/statsbymerchant"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
)

// Service exposes all merchant-related domain services.
type Service interface {
	MerchantQueryService() MerchantQueryService
	MerchantTransactionService() MerchantTransactionService
	MerchantCommandService() MerchantCommandService
	MerchantDocumentCommandService() MerchantDocumentCommandService
	MerchantDocumentQueryService() MerchantDocumentQueryService
	MerchantStatsService() merchantstatsservice.MerchantStatsService
	MerchantStatsByMerchantService() merchantstatsbymerchantservice.MerchantStatsByMerchantService
	MerchantStatsByApiKeyService() merchantstatsbyapikeyservice.MerchantStatsByApiKeyService
}

type service struct {
	merchantQuery           MerchantQueryService
	merchantTransaction     MerchantTransactionService
	merchantCommand         MerchantCommandService
	merchantDocumentCommand MerchantDocumentCommandService
	merchantDocumentQuery   MerchantDocumentQueryService
	merchantStats           merchantstatsservice.MerchantStatsService
	merchantStatsByMerchant merchantstatsbymerchantservice.MerchantStatsByMerchantService
	merchantStatsByApiKey   merchantstatsbyapikeyservice.MerchantStatsByApiKeyService
}

// Deps holds shared dependencies for merchant services.
type Deps struct {
	Repositories repository.Repositories
	Logger       logger.LoggerInterface
	Cache        *cache.CacheStore
	// Kafka is retained for backwards-compatible test/application wiring.
	// Domain email delivery is persisted through the SQLC outbox.
	Kafka *kafka.Kafka
}

// NewService wires and initializes all merchant services.
func NewService(deps *Deps) Service {
	observability, _ := observability.NewObservability("merchant-service", deps.Logger)
	cache := mencache.NewMencache(deps.Cache)

	return &service{
		merchantQuery:           newMerchantQueryService(deps, observability, cache),
		merchantTransaction:     newMerchantTransactionService(deps, observability, cache),
		merchantCommand:         newMerchantCommandService(deps, observability, cache),
		merchantDocumentCommand: newMerchantDocumentCommandService(deps, observability, cache),
		merchantDocumentQuery:   newMerchantDocumentQueryService(deps, observability, cache),

		merchantStats: merchantstatsservice.NewMerchantStatsService(&merchantstatsservice.DepsStats{
			Mencache:      cache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: observability,
		}),
		merchantStatsByMerchant: merchantstatsbymerchantservice.NewMerchantStatsByMerchantService(&merchantstatsbymerchantservice.DepsStatsByMerchant{
			Mencache:      cache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: observability,
		}),
		merchantStatsByApiKey: merchantstatsbyapikeyservice.NewMerchantStatsByApiKeyService(&merchantstatsbyapikeyservice.DepsStatsByApiKey{
			Mencache:      cache,
			Repository:    deps.Repositories,
			Logger:        deps.Logger,
			Observability: observability,
		}),
	}
}

func (s *service) MerchantQueryService() MerchantQueryService {
	return s.merchantQuery
}
func (s *service) MerchantTransactionService() MerchantTransactionService {
	return s.merchantTransaction
}
func (s *service) MerchantCommandService() MerchantCommandService {
	return s.merchantCommand
}
func (s *service) MerchantDocumentCommandService() MerchantDocumentCommandService {
	return s.merchantDocumentCommand
}
func (s *service) MerchantDocumentQueryService() MerchantDocumentQueryService {
	return s.merchantDocumentQuery
}
func (s *service) MerchantStatsService() merchantstatsservice.MerchantStatsService {
	return s.merchantStats
}
func (s *service) MerchantStatsByMerchantService() merchantstatsbymerchantservice.MerchantStatsByMerchantService {
	return s.merchantStatsByMerchant
}
func (s *service) MerchantStatsByApiKeyService() merchantstatsbyapikeyservice.MerchantStatsByApiKeyService {
	return s.merchantStatsByApiKey
}

func newMerchantQueryService(
	deps *Deps,
	observability observability.TraceLoggerObservability,
	cache mencache.Mencache,
) MerchantQueryService {
	return NewMerchantQueryService(&MerchantQueryDeps{
		Repository:    deps.Repositories,
		Cache:         cache,
		Logger:        deps.Logger,
		Observability: observability,
	})
}

func newMerchantDocumentQueryService(
	deps *Deps,
	observability observability.TraceLoggerObservability,
	cache mencache.Mencache,
) MerchantDocumentQueryService {
	return NewMerchantDocumentQueryService(&MerchantDocumentQueryDeps{
		Repository:    deps.Repositories,
		Cache:         cache,
		Logger:        deps.Logger,
		Observability: observability,
	})
}

func newMerchantTransactionService(
	deps *Deps,
	observability observability.TraceLoggerObservability,
	cache mencache.Mencache,
) MerchantTransactionService {
	return NewMerchantTransactionService(&MerchantTransactionDeps{
		Repository:    deps.Repositories,
		Cache:         cache,
		Logger:        deps.Logger,
		Observability: observability,
	})
}

func newMerchantCommandService(
	deps *Deps,
	observability observability.TraceLoggerObservability,
	cache mencache.Mencache,
) MerchantCommandService {
	return NewMerchantCommandService(&MerchantCommandServiceDeps{
		UserRepository:            deps.Repositories,
		MerchantQueryRepository:   deps.Repositories,
		MerchantCommandRepository: deps.Repositories,
		Logger:                    deps.Logger,
		Observability:             observability,
		Cache:                     cache,
	})
}

func newMerchantDocumentCommandService(
	deps *Deps,
	observability observability.TraceLoggerObservability,
	cache mencache.Mencache,
) MerchantDocumentCommandService {
	return NewMerchantDocumentCommandService(&MerchantDocumentCommandDeps{
		CommandRepository:       deps.Repositories,
		MerchantQueryRepository: deps.Repositories,
		UserRepository:          deps.Repositories,
		Logger:                  deps.Logger,
		Observability:           observability,
		Cache:                   cache,
	})
}
