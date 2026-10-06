package repository

import (
	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pbmerchant "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	pbsaldo "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	cardadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/card"
	merchantadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/merchant"
	saldoadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/saldo"
	transactionstatsrepository "github.com/MamangRust/monolith-payment-gateway-transaction/repository/stats"
	transactionbycardrepository "github.com/MamangRust/monolith-payment-gateway-transaction/repository/statsbycard"
	"github.com/jmoiron/sqlx"
)

type GuardOptions struct {
	Card     []adapter.GuardOption
	Saldo    []adapter.GuardOption
	Merchant []adapter.GuardOption
}

type Repositories interface {
	SaldoRepository
	MerchantRepository
	CardRepository
	TransactionQueryRepository
	TransactionCommandRepository
	transactionstatsrepository.TransactionStatsRepository
	transactionbycardrepository.TransactionStatsByCardRepository
}

type repositories struct {
	SaldoRepository
	MerchantRepository
	CardRepository
	TransactionQueryRepository
	TransactionCommandRepository
	transactionstatsrepository.TransactionStatsRepository
	transactionbycardrepository.TransactionStatsByCardRepository
}

func NewRepositories(
	db *sqlx.DB,
	cardQueryClient pbcard.CardQueryServiceClient,
	cardCommandClient pbcard.CardCommandServiceClient,
	saldoQueryClient pbsaldo.SaldoQueryServiceClient,
	saldoCommandClient pbsaldo.SaldoCommandServiceClient,
	merchantQueryClient pbmerchant.MerchantQueryServiceClient,
	guards ...GuardOptions,
) Repositories {
	var g GuardOptions

	if len(guards) > 0 {
		g = guards[0]
	}

	return &repositories{
		SaldoRepository:                  saldoadapter.NewAdapter(saldoQueryClient, saldoCommandClient, g.Saldo...),
		MerchantRepository:               merchantadapter.NewAdapter(merchantQueryClient, g.Merchant...),
		CardRepository:                   cardadapter.NewAdapter(cardQueryClient, cardCommandClient, g.Card...),
		TransactionQueryRepository:       NewTransactionQueryRepository(db),
		TransactionCommandRepository:     NewTransactionCommandRepository(db),
		TransactionStatsRepository:       transactionstatsrepository.NewTransactionStatsRepository(db),
		TransactionStatsByCardRepository: transactionbycardrepository.NewTransactionStatsRepository(db),
	}
}
