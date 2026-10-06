package repository

import (
	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pbsaldo "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	cardadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/card"
	saldoadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/saldo"
	transferstatsrepository "github.com/MamangRust/monolith-payment-gateway-transfer/repository/stats"
	transferstatsbycardrepository "github.com/MamangRust/monolith-payment-gateway-transfer/repository/statsbycard"
	"github.com/jmoiron/sqlx"
)

type GuardOptions struct {
	Card  []adapter.GuardOption
	Saldo []adapter.GuardOption
}

type Repositories interface {
	SaldoRepository
	TransferQueryRepository
	TransferCommandRepository
	CardRepository
	transferstatsrepository.TransferStatsRepository
	transferstatsbycardrepository.TransferStatsByCardRepository
}

type repositories struct {
	SaldoRepository
	TransferQueryRepository
	TransferCommandRepository
	CardRepository
	transferstatsrepository.TransferStatsRepository
	transferstatsbycardrepository.TransferStatsByCardRepository
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
		SaldoRepository:               saldoadapter.NewAdapter(saldoQueryClient, saldoCommandClient, g.Saldo...),
		TransferQueryRepository:       NewTransferQueryRepository(db),
		TransferCommandRepository:     NewTransferCommandRepository(db),
		TransferStatsRepository:       transferstatsrepository.NewTransferStatsRepository(db),
		TransferStatsByCardRepository: transferstatsbycardrepository.NewTransferStatsByCardRepository(db),
		CardRepository:                cardadapter.NewAdapter(cardQueryClient, cardCommandClient, g.Card...),
	}
}
