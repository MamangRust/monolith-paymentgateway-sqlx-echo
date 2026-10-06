package saldostatsrepository

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type SaldoStatsBalanceRepository interface {
	GetMonthlySaldoBalances(ctx context.Context, year int) ([]*db.GetMonthlySaldoBalancesRow, error)
	GetYearlySaldoBalances(ctx context.Context, year int) ([]*db.GetYearlySaldoBalancesRow, error)
}

type SaldoStatsTotalSaldoRepository interface {
	GetMonthlyTotalSaldoBalance(ctx context.Context, req *requests.MonthTotalSaldoBalance) ([]*db.GetMonthlyTotalSaldoBalanceRow, error)
	GetYearTotalSaldoBalance(ctx context.Context, year int) ([]*db.GetYearlyTotalSaldoBalancesRow, error)
}
