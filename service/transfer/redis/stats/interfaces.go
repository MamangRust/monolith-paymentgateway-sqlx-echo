package transferstatscache

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type TransferStatsAmountCache interface {
	GetCachedMonthTransferAmounts(ctx context.Context, year int) ([]*db.GetMonthlyTransferAmountsRow, bool)
	SetCachedMonthTransferAmounts(ctx context.Context, year int, data []*db.GetMonthlyTransferAmountsRow)

	GetCachedYearlyTransferAmounts(ctx context.Context, year int) ([]*db.GetYearlyTransferAmountsRow, bool)
	SetCachedYearlyTransferAmounts(ctx context.Context, year int, data []*db.GetYearlyTransferAmountsRow)
}

type TransferStatsStatusCache interface {
	GetCachedMonthTransferStatusSuccess(ctx context.Context, req *requests.MonthStatusTransfer) ([]*db.GetMonthTransferStatusSuccessRow, bool)
	SetCachedMonthTransferStatusSuccess(ctx context.Context, req *requests.MonthStatusTransfer, data []*db.GetMonthTransferStatusSuccessRow)

	GetCachedYearlyTransferStatusSuccess(ctx context.Context, year int) ([]*db.GetYearlyTransferStatusSuccessRow, bool)
	SetCachedYearlyTransferStatusSuccess(ctx context.Context, year int, data []*db.GetYearlyTransferStatusSuccessRow)

	GetCachedMonthTransferStatusFailed(ctx context.Context, req *requests.MonthStatusTransfer) ([]*db.GetMonthTransferStatusFailedRow, bool)
	SetCachedMonthTransferStatusFailed(ctx context.Context, req *requests.MonthStatusTransfer, data []*db.GetMonthTransferStatusFailedRow)

	GetCachedYearlyTransferStatusFailed(ctx context.Context, year int) ([]*db.GetYearlyTransferStatusFailedRow, bool)
	SetCachedYearlyTransferStatusFailed(ctx context.Context, year int, data []*db.GetYearlyTransferStatusFailedRow)
}
