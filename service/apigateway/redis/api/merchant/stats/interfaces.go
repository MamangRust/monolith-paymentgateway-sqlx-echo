package merchant_stats_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type MerchantStatsMethodCache interface {
	GetMonthlyPaymentMethodsMerchantCache(ctx context.Context, year int) (*response.ApiResponseMerchantMonthlyPaymentMethod, bool)
	SetMonthlyPaymentMethodsMerchantCache(ctx context.Context, year int, data *response.ApiResponseMerchantMonthlyPaymentMethod)

	GetYearlyPaymentMethodMerchantCache(ctx context.Context, year int) (*response.ApiResponseMerchantYearlyPaymentMethod, bool)
	SetYearlyPaymentMethodMerchantCache(ctx context.Context, year int, data *response.ApiResponseMerchantYearlyPaymentMethod)
}

type MerchantStatsAmountCache interface {
	GetMonthlyAmountMerchantCache(ctx context.Context, year int) (*response.ApiResponseMerchantMonthlyAmount, bool)
	SetMonthlyAmountMerchantCache(ctx context.Context, year int, data *response.ApiResponseMerchantMonthlyAmount)

	GetYearlyAmountMerchantCache(ctx context.Context, year int) (*response.ApiResponseMerchantYearlyAmount, bool)
	SetYearlyAmountMerchantCache(ctx context.Context, year int, data *response.ApiResponseMerchantYearlyAmount)
}

type MerchantStatsTotalAmountCache interface {
	GetMonthlyTotalAmountMerchantCache(ctx context.Context, year int) (*response.ApiResponseMerchantMonthlyTotalAmount, bool)
	SetMonthlyTotalAmountMerchantCache(ctx context.Context, year int, data *response.ApiResponseMerchantMonthlyTotalAmount)

	GetYearlyTotalAmountMerchantCache(ctx context.Context, year int) (*response.ApiResponseMerchantYearlyTotalAmount, bool)
	SetYearlyTotalAmountMerchantCache(ctx context.Context, year int, data *response.ApiResponseMerchantYearlyTotalAmount)
}
