package withdraw_stats_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type WithdrawStatsStatusCache interface {
	GetCachedMonthWithdrawStatusSuccessCache(ctx context.Context, req *requests.MonthStatusWithdraw) (*response.ApiResponseWithdrawMonthStatusSuccess, bool)
	SetCachedMonthWithdrawStatusSuccessCache(ctx context.Context, req *requests.MonthStatusWithdraw, data *response.ApiResponseWithdrawMonthStatusSuccess)

	GetCachedYearlyWithdrawStatusSuccessCache(ctx context.Context, year int) (*response.ApiResponseWithdrawYearStatusSuccess, bool)
	SetCachedYearlyWithdrawStatusSuccessCache(ctx context.Context, year int, data *response.ApiResponseWithdrawYearStatusSuccess)

	GetCachedMonthWithdrawStatusFailedCache(ctx context.Context, req *requests.MonthStatusWithdraw) (*response.ApiResponseWithdrawMonthStatusFailed, bool)
	SetCachedMonthWithdrawStatusFailedCache(ctx context.Context, req *requests.MonthStatusWithdraw, data *response.ApiResponseWithdrawMonthStatusFailed)

	GetCachedYearlyWithdrawStatusFailedCache(ctx context.Context, year int) (*response.ApiResponseWithdrawYearStatusFailed, bool)
	SetCachedYearlyWithdrawStatusFailedCache(ctx context.Context, year int, data *response.ApiResponseWithdrawYearStatusFailed)
}

type WithdrawStatsAmountCache interface {
	GetCachedMonthlyWithdraws(ctx context.Context, year int) (*response.ApiResponseWithdrawMonthAmount, bool)
	SetCachedMonthlyWithdraws(ctx context.Context, year int, data *response.ApiResponseWithdrawMonthAmount)

	GetCachedYearlyWithdraws(ctx context.Context, year int) (*response.ApiResponseWithdrawYearAmount, bool)
	SetCachedYearlyWithdraws(ctx context.Context, year int, data *response.ApiResponseWithdrawYearAmount)
}
