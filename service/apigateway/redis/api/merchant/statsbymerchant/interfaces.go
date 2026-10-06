package merchant_stats_bymerchant_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type MerchantStatsMethodByMerchantCache interface {
	GetMonthlyPaymentMethodByMerchantsCache(ctx context.Context, req *requests.MonthYearPaymentMethodMerchant) (*response.ApiResponseMerchantMonthlyPaymentMethod, bool)
	SetMonthlyPaymentMethodByMerchantsCache(ctx context.Context, req *requests.MonthYearPaymentMethodMerchant, data *response.ApiResponseMerchantMonthlyPaymentMethod)

	GetYearlyPaymentMethodByMerchantsCache(ctx context.Context, req *requests.MonthYearPaymentMethodMerchant) (*response.ApiResponseMerchantYearlyPaymentMethod, bool)
	SetYearlyPaymentMethodByMerchantsCache(ctx context.Context, req *requests.MonthYearPaymentMethodMerchant, data *response.ApiResponseMerchantYearlyPaymentMethod)
}

type MerchantStatsAmountByMerchantCache interface {
	GetMonthlyAmountByMerchantsCache(ctx context.Context, req *requests.MonthYearAmountMerchant) (*response.ApiResponseMerchantMonthlyAmount, bool)
	SetMonthlyAmountByMerchantsCache(ctx context.Context, req *requests.MonthYearAmountMerchant, data *response.ApiResponseMerchantMonthlyAmount)

	GetYearlyAmountByMerchantsCache(ctx context.Context, req *requests.MonthYearAmountMerchant) (*response.ApiResponseMerchantYearlyAmount, bool)
	SetYearlyAmountByMerchantsCache(ctx context.Context, req *requests.MonthYearAmountMerchant, data *response.ApiResponseMerchantYearlyAmount)
}

type MerchantStatsTotalAmountByMerchantCache interface {
	GetMonthlyTotalAmountByMerchantsCache(ctx context.Context, req *requests.MonthYearTotalAmountMerchant) (*response.ApiResponseMerchantMonthlyTotalAmount, bool)
	SetMonthlyTotalAmountByMerchantsCache(ctx context.Context, req *requests.MonthYearTotalAmountMerchant, data *response.ApiResponseMerchantMonthlyTotalAmount)

	GetYearlyTotalAmountByMerchantsCache(ctx context.Context, req *requests.MonthYearTotalAmountMerchant) (*response.ApiResponseMerchantYearlyTotalAmount, bool)
	SetYearlyTotalAmountByMerchantsCache(ctx context.Context, req *requests.MonthYearTotalAmountMerchant, data *response.ApiResponseMerchantYearlyTotalAmount)
}
