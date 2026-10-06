package merchant_stats_byapikey_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type MerchantStatsMethodByApiKeyCache interface {
	GetMonthlyPaymentMethodByApikeysCache(ctx context.Context, req *requests.MonthYearPaymentMethodApiKey) (*response.ApiResponseMerchantMonthlyPaymentMethod, bool)
	SetMonthlyPaymentMethodByApikeysCache(ctx context.Context, req *requests.MonthYearPaymentMethodApiKey, data *response.ApiResponseMerchantMonthlyPaymentMethod)

	GetYearlyPaymentMethodByApikeysCache(ctx context.Context, req *requests.MonthYearPaymentMethodApiKey) (*response.ApiResponseMerchantYearlyPaymentMethod, bool)
	SetYearlyPaymentMethodByApikeysCache(ctx context.Context, req *requests.MonthYearPaymentMethodApiKey, data *response.ApiResponseMerchantYearlyPaymentMethod)
}

type MerchantStatsAmountByApiKeyCache interface {
	GetMonthlyAmountByApikeysCache(ctx context.Context, req *requests.MonthYearAmountApiKey) (*response.ApiResponseMerchantMonthlyAmount, bool)
	SetMonthlyAmountByApikeysCache(ctx context.Context, req *requests.MonthYearAmountApiKey, data *response.ApiResponseMerchantMonthlyAmount)

	GetYearlyAmountByApikeysCache(ctx context.Context, req *requests.MonthYearAmountApiKey) (*response.ApiResponseMerchantYearlyAmount, bool)
	SetYearlyAmountByApikeysCache(ctx context.Context, req *requests.MonthYearAmountApiKey, data *response.ApiResponseMerchantYearlyAmount)
}

type MerchantStatsTotalAmountByApiKeyCache interface {
	GetMonthlyTotalAmountByApikeysCache(ctx context.Context, req *requests.MonthYearTotalAmountApiKey) (*response.ApiResponseMerchantMonthlyTotalAmount, bool)
	SetMonthlyTotalAmountByApikeysCache(ctx context.Context, req *requests.MonthYearTotalAmountApiKey, data *response.ApiResponseMerchantMonthlyTotalAmount)

	GetYearlyTotalAmountByApikeysCache(ctx context.Context, req *requests.MonthYearTotalAmountApiKey) (*response.ApiResponseMerchantYearlyTotalAmount, bool)
	SetYearlyTotalAmountByApikeysCache(ctx context.Context, req *requests.MonthYearTotalAmountApiKey, data *response.ApiResponseMerchantYearlyTotalAmount)
}
