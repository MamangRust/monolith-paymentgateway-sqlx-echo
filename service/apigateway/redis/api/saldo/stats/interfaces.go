package saldo_stats_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type SaldoStatsTotalCache interface {
	GetMonthlyTotalSaldoBalanceCache(ctx context.Context, req *requests.MonthTotalSaldoBalance) (*response.ApiResponseMonthTotalSaldo, bool)
	SetMonthlyTotalSaldoCache(ctx context.Context, req *requests.MonthTotalSaldoBalance, data *response.ApiResponseMonthTotalSaldo)

	GetYearTotalSaldoBalanceCache(ctx context.Context, year int) (*response.ApiResponseYearTotalSaldo, bool)
	SetYearTotalSaldoBalanceCache(ctx context.Context, year int, data *response.ApiResponseYearTotalSaldo)
}

type SaldoStatsBalanceCache interface {
	GetMonthlySaldoBalanceCache(ctx context.Context, year int) (*response.ApiResponseMonthSaldoBalances, bool)
	SetMonthlySaldoBalanceCache(ctx context.Context, year int, data *response.ApiResponseMonthSaldoBalances)

	GetYearlySaldoBalanceCache(ctx context.Context, year int) (*response.ApiResponseYearSaldoBalances, bool)
	SetYearlySaldoBalanceCache(ctx context.Context, year int, data *response.ApiResponseYearSaldoBalances)
}
