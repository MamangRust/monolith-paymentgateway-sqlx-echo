package repository

import (
	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pbsaldo "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	cardadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/card"
	saldoadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/saldo"
	withdrawstatsrepository "github.com/MamangRust/monolith-payment-gateway-withdraw/repository/stats"
	withdrawstatsbycardrepository "github.com/MamangRust/monolith-payment-gateway-withdraw/repository/statsbycard"
	"github.com/jmoiron/sqlx"
)

type GuardOptions struct {
	Card  []adapter.GuardOption
	Saldo []adapter.GuardOption
}

type Repositories interface {
	CardRepository
	SaldoRepository
	WithdrawQueryRepository
	WithdrawCommandRepository
	withdrawstatsrepository.WithdrawStatsRepository
	withdrawstatsbycardrepository.WithdrawStatsByCardRepository
}

type repositories struct {
	CardRepository
	SaldoRepository
	WithdrawQueryRepository
	WithdrawCommandRepository
	withdrawstatsrepository.WithdrawStatsRepository
	withdrawstatsbycardrepository.WithdrawStatsByCardRepository
}

func NewRepositories(
	db *sqlx.DB,
	cardQueryClient pbcard.CardQueryServiceClient,
	cardCommandClient pbcard.CardCommandServiceClient,
	saldoQueryClient pbsaldo.SaldoQueryServiceClient,
	saldoCommandClient pbsaldo.SaldoCommandServiceClient,
	dailyLimit int64,
	guards ...GuardOptions,
) Repositories {
	var g GuardOptions

	if len(guards) > 0 {
		g = guards[0]
	}

	return &repositories{
		CardRepository:                cardadapter.NewAdapter(cardQueryClient, cardCommandClient, g.Card...),
		SaldoRepository:               saldoadapter.NewAdapter(saldoQueryClient, saldoCommandClient, g.Saldo...),
		WithdrawQueryRepository:       NewWithdrawQueryRepository(db),
		WithdrawCommandRepository:     NewWithdrawCommandRepository(db, dailyLimit),
		WithdrawStatsRepository:       withdrawstatsrepository.NewWithdrawStatsRepository(db),
		WithdrawStatsByCardRepository: withdrawstatsbycardrepository.NewWithdrawStatsByCardRepository(db),
	}
}
