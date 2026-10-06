package repository

import (
	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pbsaldo "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	cardadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/card"
	saldoadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/saldo"
	topupstatsrepository "github.com/MamangRust/monolith-payment-gateway-topup/repository/stats"
	topupstatsbycardrepository "github.com/MamangRust/monolith-payment-gateway-topup/repository/statsbycard"
	"github.com/jmoiron/sqlx"
)

type GuardOptions struct {
	Card  []adapter.GuardOption
	Saldo []adapter.GuardOption
}

type Repositories interface {
	TopupQueryRepository
	TopupCommandRepository
	CardRepository
	SaldoRepository
	topupstatsrepository.TopupStatsRepository
	topupstatsbycardrepository.TopupStatsByCardRepository
}

type repositories struct {
	TopupQueryRepository
	TopupCommandRepository
	CardRepository
	SaldoRepository
	topupstatsrepository.TopupStatsRepository
	topupstatsbycardrepository.TopupStatsByCardRepository
}

func NewRepositories(
	db *sqlx.DB,
	cardQueryClient pbcard.CardQueryServiceClient,
	cardCommandClient pbcard.CardCommandServiceClient,
	saldoQueryClient pbsaldo.SaldoQueryServiceClient,
	saldoCommandClient pbsaldo.SaldoCommandServiceClient,
	guards ...GuardOptions,
) Repositories {
	var g GuardOptions

	if len(guards) > 0 {
		g = guards[0]
	}

	return &repositories{
		TopupQueryRepository:       NewTopupQueryRepository(db),
		TopupCommandRepository:     NewTopupCommandRepository(db),
		TopupStatsRepository:       topupstatsrepository.NewTopupStatsRepository(db),
		TopupStatsByCardRepository: topupstatsbycardrepository.NewTopupStatsByCardRepository(db),
		CardRepository:             cardadapter.NewAdapter(cardQueryClient, cardCommandClient, g.Card...),
		SaldoRepository:            saldoadapter.NewAdapter(saldoQueryClient, saldoCommandClient, g.Saldo...),
	}
}
