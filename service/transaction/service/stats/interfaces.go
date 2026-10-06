package transactionstatsservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type TransactionStatsAmountService interface {
	FindMonthlyAmounts(ctx context.Context, year int) ([]*db.GetMonthlyAmountsRow, error)
	FindYearlyAmounts(ctx context.Context, year int) ([]*db.GetYearlyAmountsRow, error)
}

type TransactionStatsMethodService interface {
	FindMonthlyPaymentMethods(ctx context.Context, year int) ([]*db.GetMonthlyPaymentMethodsRow, error)
	FindYearlyPaymentMethods(ctx context.Context, year int) ([]*db.GetYearlyPaymentMethodsRow, error)
}

type TransactionStatsStatusService interface {
	FindMonthTransactionStatusSuccess(ctx context.Context, req *requests.MonthStatusTransaction) ([]*db.GetMonthTransactionStatusSuccessRow, error)
	FindYearlyTransactionStatusSuccess(ctx context.Context, year int) ([]*db.GetYearlyTransactionStatusSuccessRow, error)
	FindMonthTransactionStatusFailed(ctx context.Context, req *requests.MonthStatusTransaction) ([]*db.GetMonthTransactionStatusFailedRow, error)
	FindYearlyTransactionStatusFailed(ctx context.Context, year int) ([]*db.GetYearlyTransactionStatusFailedRow, error)
}
