package transaction_stats_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type TransactionStatsAmountCache interface {
	GetMonthlyAmountsCache(ctx context.Context, year int) (*response.ApiResponseTransactionMonthAmount, bool)
	SetMonthlyAmountsCache(ctx context.Context, year int, data *response.ApiResponseTransactionMonthAmount)

	GetYearlyAmountsCache(ctx context.Context, year int) (*response.ApiResponseTransactionYearAmount, bool)
	SetYearlyAmountsCache(ctx context.Context, year int, data *response.ApiResponseTransactionYearAmount)
}

type TransactionStatsMethodCache interface {
	GetMonthlyPaymentMethodsCache(ctx context.Context, year int) (*response.ApiResponseTransactionMonthMethod, bool)
	SetMonthlyPaymentMethodsCache(ctx context.Context, year int, data *response.ApiResponseTransactionMonthMethod)

	GetYearlyPaymentMethodsCache(ctx context.Context, year int) (*response.ApiResponseTransactionYearMethod, bool)
	SetYearlyPaymentMethodsCache(ctx context.Context, year int, data *response.ApiResponseTransactionYearMethod)
}

type TransactionStatsStatusCache interface {
	GetMonthTransactionStatusSuccessCache(ctx context.Context, req *requests.MonthStatusTransaction) (*response.ApiResponseTransactionMonthStatusSuccess, bool)
	SetMonthTransactionStatusSuccessCache(ctx context.Context, req *requests.MonthStatusTransaction, data *response.ApiResponseTransactionMonthStatusSuccess)

	GetYearTransactionStatusSuccessCache(ctx context.Context, year int) (*response.ApiResponseTransactionYearStatusSuccess, bool)
	SetYearTransactionStatusSuccessCache(ctx context.Context, year int, data *response.ApiResponseTransactionYearStatusSuccess)

	GetMonthTransactionStatusFailedCache(ctx context.Context, req *requests.MonthStatusTransaction) (*response.ApiResponseTransactionMonthStatusFailed, bool)
	SetMonthTransactionStatusFailedCache(ctx context.Context, req *requests.MonthStatusTransaction, data *response.ApiResponseTransactionMonthStatusFailed)

	GetYearTransactionStatusFailedCache(ctx context.Context, year int) (*response.ApiResponseTransactionYearStatusFailed, bool)
	SetYearTransactionStatusFailedCache(ctx context.Context, year int, data *response.ApiResponseTransactionYearStatusFailed)
}
