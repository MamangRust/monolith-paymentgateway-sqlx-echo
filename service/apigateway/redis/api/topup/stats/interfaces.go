package topup_stats_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type TopupStatsStatusCache interface {
	GetMonthTopupStatusSuccessCache(ctx context.Context, req *requests.MonthTopupStatus) (*response.ApiResponseTopupMonthStatusSuccess, bool)
	SetMonthTopupStatusSuccessCache(ctx context.Context, req *requests.MonthTopupStatus, data *response.ApiResponseTopupMonthStatusSuccess)

	GetYearlyTopupStatusSuccessCache(ctx context.Context, year int) (*response.ApiResponseTopupYearStatusSuccess, bool)
	SetYearlyTopupStatusSuccessCache(ctx context.Context, year int, data *response.ApiResponseTopupYearStatusSuccess)

	GetMonthTopupStatusFailedCache(ctx context.Context, req *requests.MonthTopupStatus) (*response.ApiResponseTopupMonthStatusFailed, bool)
	SetMonthTopupStatusFailedCache(ctx context.Context, req *requests.MonthTopupStatus, data *response.ApiResponseTopupMonthStatusFailed)

	GetYearlyTopupStatusFailedCache(ctx context.Context, year int) (*response.ApiResponseTopupYearStatusFailed, bool)
	SetYearlyTopupStatusFailedCache(ctx context.Context, year int, data *response.ApiResponseTopupYearStatusFailed)
}

type TopupStatsMethodCache interface {
	GetMonthlyTopupMethodsCache(ctx context.Context, year int) (*response.ApiResponseTopupMonthMethod, bool)
	SetMonthlyTopupMethodsCache(ctx context.Context, year int, data *response.ApiResponseTopupMonthMethod)

	GetYearlyTopupMethodsCache(ctx context.Context, year int) (*response.ApiResponseTopupYearMethod, bool)
	SetYearlyTopupMethodsCache(ctx context.Context, year int, data *response.ApiResponseTopupYearMethod)
}

type TopupStatsAmountCache interface {
	GetMonthlyTopupAmountsCache(ctx context.Context, year int) (*response.ApiResponseTopupMonthAmount, bool)
	SetMonthlyTopupAmountsCache(ctx context.Context, year int, data *response.ApiResponseTopupMonthAmount)

	GetYearlyTopupAmountsCache(ctx context.Context, year int) (*response.ApiResponseTopupYearAmount, bool)
	SetYearlyTopupAmountsCache(ctx context.Context, year int, data *response.ApiResponseTopupYearAmount)
}
