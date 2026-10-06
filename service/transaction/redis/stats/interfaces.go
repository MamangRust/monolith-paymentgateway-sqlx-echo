package transactionstatscache

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type TransactionStatsAmountCache interface {
	GetMonthlyAmountsCache(ctx context.Context, year int) ([]*db.GetMonthlyAmountsRow, bool)
	SetMonthlyAmountsCache(ctx context.Context, year int, data []*db.GetMonthlyAmountsRow)

	GetYearlyAmountsCache(ctx context.Context, year int) ([]*db.GetYearlyAmountsRow, bool)
	SetYearlyAmountsCache(ctx context.Context, year int, data []*db.GetYearlyAmountsRow)
}

type TransactionStatsMethodCache interface {
	GetMonthlyPaymentMethodsCache(ctx context.Context, year int) ([]*db.GetMonthlyPaymentMethodsRow, bool)
	SetMonthlyPaymentMethodsCache(ctx context.Context, year int, data []*db.GetMonthlyPaymentMethodsRow)

	GetYearlyPaymentMethodsCache(ctx context.Context, year int) ([]*db.GetYearlyPaymentMethodsRow, bool)
	SetYearlyPaymentMethodsCache(ctx context.Context, year int, data []*db.GetYearlyPaymentMethodsRow)
}

type TransactionStatsStatusCache interface {
	GetMonthTransactionStatusSuccessCache(ctx context.Context, req *requests.MonthStatusTransaction) ([]*db.GetMonthTransactionStatusSuccessRow, bool)
	SetMonthTransactionStatusSuccessCache(ctx context.Context, req *requests.MonthStatusTransaction, data []*db.GetMonthTransactionStatusSuccessRow)

	GetYearTransactionStatusSuccessCache(ctx context.Context, year int) ([]*db.GetYearlyTransactionStatusSuccessRow, bool)
	SetYearTransactionStatusSuccessCache(ctx context.Context, year int, data []*db.GetYearlyTransactionStatusSuccessRow)

	GetMonthTransactionStatusFailedCache(ctx context.Context, req *requests.MonthStatusTransaction) ([]*db.GetMonthTransactionStatusFailedRow, bool)
	SetMonthTransactionStatusFailedCache(ctx context.Context, req *requests.MonthStatusTransaction, data []*db.GetMonthTransactionStatusFailedRow)

	GetYearTransactionStatusFailedCache(ctx context.Context, year int) ([]*db.GetYearlyTransactionStatusFailedRow, bool)
	SetYearTransactionStatusFailedCache(ctx context.Context, year int, data []*db.GetYearlyTransactionStatusFailedRow)
}
