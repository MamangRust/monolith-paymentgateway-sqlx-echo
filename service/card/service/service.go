package service

import (
	mencache "github.com/MamangRust/monolith-payment-gateway-card/redis"
	"github.com/MamangRust/monolith-payment-gateway-card/repository"
	cardstatsservice "github.com/MamangRust/monolith-payment-gateway-card/service/stats"
	cardstatsbycard "github.com/MamangRust/monolith-payment-gateway-card/service/statsbycard"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
)

type Service interface {
	CardQueryService
	CardDashboardService
	cardstatsservice.CardStatsService
	cardstatsbycard.CardStatsByCardService
	CardCommandService
	CardAuthorizationService
	CardPaymentService
	CardRewardService
	BillingEngineService
}

type service struct {
	CardQueryService
	CardDashboardService
	cardstatsservice.CardStatsService
	cardstatsbycard.CardStatsByCardService
	CardCommandService
	CardAuthorizationService
	CardPaymentService
	CardRewardService
	BillingEngineService
}

type Deps struct {
	Cache        *cache.CacheStore
	Repositories *repository.Repositories
	Logger       logger.LoggerInterface
	Kafka        *kafka.Kafka
}

func NewService(deps *Deps) Service {
	observability, _ := observability.NewObservability("card-server", deps.Logger)

	cache := mencache.NewMencache(deps.Cache)

	return &service{
		CardQueryService:     newCardQuery(deps, observability, cache),
		CardCommandService:   newCardCommand(deps, observability, cache),
		CardDashboardService: newCardDashboard(deps, observability, cache),
		CardStatsService: cardstatsservice.NewCardStatsService(&cardstatsservice.DepsStats{
			Mencache:      cache,
			Repositories:  deps.Repositories.CardStatistic,
			Logger:        deps.Logger,
			Observability: observability,
		}),
		CardStatsByCardService: cardstatsbycard.NewCardStatsByCardService(&cardstatsbycard.DepsStatsByCard{
			Mencache:      cache,
			Repositories:  deps.Repositories.CardStatisticByCard,
			Logger:        deps.Logger,
			Observability: observability,
		}),
		CardAuthorizationService: newCardAuthorization(deps, observability, cache),
		CardPaymentService:       newCardPayment(deps, observability, cache),
		CardRewardService:        newCardReward(deps, observability, cache),
		BillingEngineService:     newBillingEngine(deps, observability, cache),
	}
}

// newCardQuery initializes a new instance of the CardQueryService.
// It takes a pointer to Deps and a mapper for CardResponse.
// It returns a pointer to CardQueryService.
func newCardQuery(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.Mencache) CardQueryService {
	return NewCardQueryService(&CardQueryServiceDeps{
		Cache:               cache,
		CardQueryRepository: deps.Repositories.CardQuery,
		Logger:              deps.Logger,
		Observability:       observability,
	})
}

// newCardDashboard initializes a new instance of the CardDashboardService.
// It takes a pointer to Deps and a mapper for CardResponse.
// It returns a pointer to CardDashboardService.
func newCardDashboard(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.Mencache) CardDashboardService {
	return NewCardDashboardService(&cardDashboardDeps{
		Cache:                   cache,
		CardDashboardRepository: deps.Repositories.CardDashboard,
		Logger:                  deps.Logger,
		Observability:           observability,
	})
}

// newCardCommand initializes a new instance of the CardCommandService.
// It takes a pointer to Deps and a mapper for CardResponse.
// It returns a pointer to CardCommandService.
func newCardCommand(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.Mencache) CardCommandService {
	return NewCardCommandService(&CardCommandServiceDeps{
		Cache:                 cache,
		UserRepository:        deps.Repositories.User,
		CardCommandRepository: deps.Repositories.CardCommand,
		Logger:                deps.Logger,
		Observability:         observability,
	})
}

func newCardAuthorization(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.Mencache) CardAuthorizationService {
	return NewCardAuthorizationService(&CardAuthorizationServiceDeps{
		Cache:                  cache,
		AuthCache:              cache,
		CardCommandRepository:  deps.Repositories.CardCommand,
		CardAuthRepository:     deps.Repositories.CardAuthTransaction,
		CardPaymentRepository:  deps.Repositories.CardPayment,
		CardRewardRepository:   deps.Repositories.CardReward,
		BillingCycleRepository: deps.Repositories.BillingCycle,
		CardQueryRepository:    deps.Repositories.CardQuery,
		Logger:                 deps.Logger,
		Observability:          observability,
	})
}

func newCardPayment(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.Mencache) CardPaymentService {
	return NewCardPaymentService(&CardPaymentServiceDeps{
		CardPaymentRepository:  deps.Repositories.CardPayment,
		BillingCycleRepository: deps.Repositories.BillingCycle,
		CardCommandRepository:  deps.Repositories.CardCommand,
		Logger:                 deps.Logger,
		Observability:          observability,
	})
}

func newCardReward(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.Mencache) CardRewardService {
	return NewCardRewardService(&CardRewardServiceDeps{
		CardRewardRepository:  deps.Repositories.CardReward,
		CardCommandRepository: deps.Repositories.CardCommand,
		Logger:                deps.Logger,
		Observability:         observability,
	})
}

func newBillingEngine(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.Mencache) BillingEngineService {
	return NewBillingEngineService(&BillingEngineServiceDeps{
		CardCommandRepository:  deps.Repositories.CardCommand,
		BillingCycleRepository: deps.Repositories.BillingCycle,
		CardQueryRepository:    deps.Repositories.CardQuery,
		Kafka:                  deps.Kafka,
		Logger:                 deps.Logger,
		Observability:          observability,
	})
}
