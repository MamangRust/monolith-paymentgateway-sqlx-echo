package transaction_stats_bycard_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type TransactionStatsByCardAmountCache interface {
	GetMonthlyAmountsByCardCache(ctx context.Context, req *requests.MonthYearPaymentMethod) (*response.ApiResponseTransactionMonthAmount, bool)
	SetMonthlyAmountsByCardCache(ctx context.Context, req *requests.MonthYearPaymentMethod, data *response.ApiResponseTransactionMonthAmount)

	GetYearlyAmountsByCardCache(ctx context.Context, req *requests.MonthYearPaymentMethod) (*response.ApiResponseTransactionYearAmount, bool)
	SetYearlyAmountsByCardCache(ctx context.Context, req *requests.MonthYearPaymentMethod, data *response.ApiResponseTransactionYearAmount)
}

type TransactionStatsByCardMethodCache interface {
	GetMonthlyPaymentMethodsByCardCache(ctx context.Context, req *requests.MonthYearPaymentMethod) (*response.ApiResponseTransactionMonthMethod, bool)
	SetMonthlyPaymentMethodsByCardCache(ctx context.Context, req *requests.MonthYearPaymentMethod, data *response.ApiResponseTransactionMonthMethod)

	GetYearlyPaymentMethodsByCardCache(ctx context.Context, req *requests.MonthYearPaymentMethod) (*response.ApiResponseTransactionYearMethod, bool)
	SetYearlyPaymentMethodsByCardCache(ctx context.Context, req *requests.MonthYearPaymentMethod, data *response.ApiResponseTransactionYearMethod)
}

type TransactionStatsByCardStatusCache interface {
	GetMonthTransactionStatusSuccessByCardCache(ctx context.Context, req *requests.MonthStatusTransactionCardNumber) (*response.ApiResponseTransactionMonthStatusSuccess, bool)
	SetMonthTransactionStatusSuccessByCardCache(ctx context.Context, req *requests.MonthStatusTransactionCardNumber, data *response.ApiResponseTransactionMonthStatusSuccess)

	GetYearTransactionStatusSuccessByCardCache(ctx context.Context, req *requests.YearStatusTransactionCardNumber) (*response.ApiResponseTransactionYearStatusSuccess, bool)
	SetYearTransactionStatusSuccessByCardCache(ctx context.Context, req *requests.YearStatusTransactionCardNumber, data *response.ApiResponseTransactionYearStatusSuccess)

	GetMonthTransactionStatusFailedByCardCache(ctx context.Context, req *requests.MonthStatusTransactionCardNumber) (*response.ApiResponseTransactionMonthStatusFailed, bool)
	SetMonthTransactionStatusFailedByCardCache(ctx context.Context, req *requests.MonthStatusTransactionCardNumber, data *response.ApiResponseTransactionMonthStatusFailed)

	GetYearTransactionStatusFailedByCardCache(ctx context.Context, req *requests.YearStatusTransactionCardNumber) (*response.ApiResponseTransactionYearStatusFailed, bool)
	SetYearTransactionStatusFailedByCardCache(ctx context.Context, req *requests.YearStatusTransactionCardNumber, data *response.ApiResponseTransactionYearStatusFailed)
}
