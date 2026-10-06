package saldostatsrepository

import (
	"github.com/jmoiron/sqlx"
)

type SaldoStatsRepository interface {
	SaldoStatsBalanceRepository
	SaldoStatsTotalSaldoRepository
}

type repository struct {
	SaldoStatsBalanceRepository
	SaldoStatsTotalSaldoRepository
}

func NewSaldoStatsRepository(db *sqlx.DB) SaldoStatsRepository {
	return &repository{
		SaldoStatsBalanceRepository:    NewSaldoStatsBalanceRepository(db),
		SaldoStatsTotalSaldoRepository: NewSaldoStatsTotalBalanceRepository(db),
	}
}
