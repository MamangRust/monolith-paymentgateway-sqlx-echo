package topup_stats_bycard_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type TopupStatsStatusByCardCache interface {
	GetMonthTopupStatusSuccessByCardNumberCache(ctx context.Context, req *requests.MonthTopupStatusCardNumber) (*response.ApiResponseTopupMonthStatusSuccess, bool)
	SetMonthTopupStatusSuccessByCardNumberCache(ctx context.Context, req *requests.MonthTopupStatusCardNumber, data *response.ApiResponseTopupMonthStatusSuccess)

	GetYearlyTopupStatusSuccessByCardNumberCache(ctx context.Context, req *requests.YearTopupStatusCardNumber) (*response.ApiResponseTopupYearStatusSuccess, bool)
	SetYearlyTopupStatusSuccessByCardNumberCache(ctx context.Context, req *requests.YearTopupStatusCardNumber, data *response.ApiResponseTopupYearStatusSuccess)

	GetMonthTopupStatusFailedByCardNumberCache(ctx context.Context, req *requests.MonthTopupStatusCardNumber) (*response.ApiResponseTopupMonthStatusFailed, bool)
	SetMonthTopupStatusFailedByCardNumberCache(ctx context.Context, req *requests.MonthTopupStatusCardNumber, data *response.ApiResponseTopupMonthStatusFailed)

	GetYearlyTopupStatusFailedByCardNumberCache(ctx context.Context, req *requests.YearTopupStatusCardNumber) (*response.ApiResponseTopupYearStatusFailed, bool)
	SetYearlyTopupStatusFailedByCardNumberCache(ctx context.Context, req *requests.YearTopupStatusCardNumber, data *response.ApiResponseTopupYearStatusFailed)
}

type TopupStatsMethodByCardCache interface {
	GetMonthlyTopupMethodsByCardNumberCache(ctx context.Context, req *requests.YearMonthMethod) (*response.ApiResponseTopupMonthMethod, bool)
	SetMonthlyTopupMethodsByCardNumberCache(ctx context.Context, req *requests.YearMonthMethod, data *response.ApiResponseTopupMonthMethod)

	GetYearlyTopupMethodsByCardNumberCache(ctx context.Context, req *requests.YearMonthMethod) (*response.ApiResponseTopupYearMethod, bool)
	SetYearlyTopupMethodsByCardNumberCache(ctx context.Context, req *requests.YearMonthMethod, data *response.ApiResponseTopupYearMethod)
}

type TopupStatsAmountByCardCache interface {
	GetMonthlyTopupAmountsByCardNumberCache(ctx context.Context, req *requests.YearMonthMethod) (*response.ApiResponseTopupMonthAmount, bool)
	SetMonthlyTopupAmountsByCardNumberCache(ctx context.Context, req *requests.YearMonthMethod, data *response.ApiResponseTopupMonthAmount)

	GetYearlyTopupAmountsByCardNumberCache(ctx context.Context, req *requests.YearMonthMethod) (*response.ApiResponseTopupYearAmount, bool)
	SetYearlyTopupAmountsByCardNumberCache(ctx context.Context, req *requests.YearMonthMethod, data *response.ApiResponseTopupYearAmount)
}
