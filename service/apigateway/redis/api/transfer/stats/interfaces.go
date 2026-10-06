package transfer_stats_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type TransferStatsAmountCache interface {
	GetCachedMonthTransferAmounts(ctx context.Context, year int) (*response.ApiResponseTransferMonthAmount, bool)
	SetCachedMonthTransferAmounts(ctx context.Context, year int, data *response.ApiResponseTransferMonthAmount)

	GetCachedYearlyTransferAmounts(ctx context.Context, year int) (*response.ApiResponseTransferYearAmount, bool)
	SetCachedYearlyTransferAmounts(ctx context.Context, year int, data *response.ApiResponseTransferYearAmount)
}

type TransferStatsStatusCache interface {
	GetCachedMonthTransferStatusSuccess(ctx context.Context, req *requests.MonthStatusTransfer) (*response.ApiResponseTransferMonthStatusSuccess, bool)
	SetCachedMonthTransferStatusSuccess(ctx context.Context, req *requests.MonthStatusTransfer, data *response.ApiResponseTransferMonthStatusSuccess)

	GetCachedYearlyTransferStatusSuccess(ctx context.Context, year int) (*response.ApiResponseTransferYearStatusSuccess, bool)
	SetCachedYearlyTransferStatusSuccess(ctx context.Context, year int, data *response.ApiResponseTransferYearStatusSuccess)

	GetCachedMonthTransferStatusFailed(ctx context.Context, req *requests.MonthStatusTransfer) (*response.ApiResponseTransferMonthStatusFailed, bool)
	SetCachedMonthTransferStatusFailed(ctx context.Context, req *requests.MonthStatusTransfer, data *response.ApiResponseTransferMonthStatusFailed)

	GetCachedYearlyTransferStatusFailed(ctx context.Context, year int) (*response.ApiResponseTransferYearStatusFailed, bool)
	SetCachedYearlyTransferStatusFailed(ctx context.Context, year int, data *response.ApiResponseTransferYearStatusFailed)
}
