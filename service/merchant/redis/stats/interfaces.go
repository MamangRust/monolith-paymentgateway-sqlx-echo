package merchantstatscache

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
)

type MerchantStatsMethodCache interface {
	GetMonthlyPaymentMethodsMerchantCache(ctx context.Context, year int) ([]*db.GetMonthlyPaymentMethodsMerchantRow, bool)
	SetMonthlyPaymentMethodsMerchantCache(ctx context.Context, year int, data []*db.GetMonthlyPaymentMethodsMerchantRow)

	GetYearlyPaymentMethodMerchantCache(ctx context.Context, year int) ([]*db.GetYearlyPaymentMethodMerchantRow, bool)
	SetYearlyPaymentMethodMerchantCache(ctx context.Context, year int, data []*db.GetYearlyPaymentMethodMerchantRow)
}

type MerchantStatsAmountCache interface {
	GetMonthlyAmountMerchantCache(ctx context.Context, year int) ([]*db.GetMonthlyAmountMerchantRow, bool)
	SetMonthlyAmountMerchantCache(ctx context.Context, year int, data []*db.GetMonthlyAmountMerchantRow)

	GetYearlyAmountMerchantCache(ctx context.Context, year int) ([]*db.GetYearlyAmountMerchantRow, bool)
	SetYearlyAmountMerchantCache(ctx context.Context, year int, data []*db.GetYearlyAmountMerchantRow)
}

type MerchantStatsTotalAmountCache interface {
	GetMonthlyTotalAmountMerchantCache(ctx context.Context, year int) ([]*db.GetMonthlyTotalAmountMerchantRow, bool)
	SetMonthlyTotalAmountMerchantCache(ctx context.Context, year int, data []*db.GetMonthlyTotalAmountMerchantRow)

	GetYearlyTotalAmountMerchantCache(ctx context.Context, year int) ([]*db.GetYearlyTotalAmountMerchantRow, bool)
	SetYearlyTotalAmountMerchantCache(ctx context.Context, year int, data []*db.GetYearlyTotalAmountMerchantRow)
}
