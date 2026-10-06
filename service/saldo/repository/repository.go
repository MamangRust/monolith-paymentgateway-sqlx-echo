package repository

import (
	"github.com/jmoiron/sqlx"

	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	cardadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/card"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	saldostatsrepository "github.com/MamangRust/monolith-payment-gateway-saldo/repository/stats"
)

type GuardOptions struct {
	Card []adapter.GuardOption
}

type Repositories interface {
	SaldoQueryRepository
	SaldoCommandRepository
	saldostatsrepository.SaldoStatsRepository
	CardRepository
}

type repositories struct {
	SaldoQueryRepository
	SaldoCommandRepository
	saldostatsrepository.SaldoStatsRepository
	CardRepository
}

func NewRepositories(
	db *sqlx.DB,
	cardQueryClient pbcard.CardQueryServiceClient,
	cardCommandClient pbcard.CardCommandServiceClient,
	guards ...GuardOptions,
) Repositories {
	var g GuardOptions

	if len(guards) > 0 {
		g = guards[0]
	}

	return &repositories{
		SaldoQueryRepository:   NewSaldoQueryRepository(db),
		SaldoCommandRepository: NewSaldoCommandRepository(db),
		SaldoStatsRepository:   saldostatsrepository.NewSaldoStatsRepository(db),
		CardRepository:         cardadapter.NewAdapter(cardQueryClient, cardCommandClient, g.Card...),
	}
}
