package repository

import (
	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
	repositorydashboard "github.com/MamangRust/monolith-payment-gateway-card/repository/dashboard"
	repositorystats "github.com/MamangRust/monolith-payment-gateway-card/repository/stats"
	repositorystatsbycard "github.com/MamangRust/monolith-payment-gateway-card/repository/statsbycard"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	useradapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/user"
	"github.com/jmoiron/sqlx"
)

type GuardOptions struct {
	User []adapter.GuardOption
}

type Repositories struct {
	CardCommand         CardCommandRepository
	CardQuery           CardQueryRepository
	CardDashboard       repositorydashboard.CardDashboardRepository
	CardStatistic       repositorystats.CardStatsRepository
	CardStatisticByCard repositorystatsbycard.CardStatsByCardRepository
	User                UserRepository
	CardAuthTransaction CardAuthTransactionRepository
	CardPayment         CardPaymentRepository
	CardReward          CardRewardRepository
	BillingCycle        BillingCycleRepository
}

func NewRepositories(db *sqlx.DB, userQueryClient pbuser.UserQueryServiceClient, guards ...GuardOptions) *Repositories {
	var g GuardOptions

	if len(guards) > 0 {
		g = guards[0]
	}

	return &Repositories{
		CardQuery:           NewCardQueryRepository(db),
		CardCommand:         NewCardCommandRepository(db),
		CardDashboard:       repositorydashboard.NewCardDashboardRepository(db),
		CardStatistic:       repositorystats.NewCardStatsRepository(db),
		CardStatisticByCard: repositorystatsbycard.NewCardStatsByCardRepository(db),
		User:                useradapter.NewQueryAdapter(userQueryClient, g.User...),
		CardAuthTransaction: NewCardAuthTransactionRepository(db),
		CardPayment:         NewCardPaymentRepository(db),
		CardReward:          NewCardRewardRepository(db),
		BillingCycle:        NewBillingCycleRepository(db),
	}
}
