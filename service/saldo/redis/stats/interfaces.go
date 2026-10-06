package saldostatscache

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type SaldoStatsTotalCache interface {
	GetMonthlyTotalSaldoBalanceCache(ctx context.Context, req *requests.MonthTotalSaldoBalance) ([]*db.GetMonthlyTotalSaldoBalanceRow, bool)
	SetMonthlyTotalSaldoCache(ctx context.Context, req *requests.MonthTotalSaldoBalance, data []*db.GetMonthlyTotalSaldoBalanceRow)

	GetYearTotalSaldoBalanceCache(ctx context.Context, year int) ([]*db.GetYearlyTotalSaldoBalancesRow, bool)
	SetYearTotalSaldoBalanceCache(ctx context.Context, year int, data []*db.GetYearlyTotalSaldoBalancesRow)
}

type SaldoStatsBalanceCache interface {
	GetMonthlySaldoBalanceCache(ctx context.Context, year int) ([]*db.GetMonthlySaldoBalancesRow, bool)
	SetMonthlySaldoBalanceCache(ctx context.Context, year int, data []*db.GetMonthlySaldoBalancesRow)

	GetYearlySaldoBalanceCache(ctx context.Context, year int) ([]*db.GetYearlySaldoBalancesRow, bool)
	SetYearlySaldoBalanceCache(ctx context.Context, year int, data []*db.GetYearlySaldoBalancesRow)
}
