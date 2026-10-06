package transfer_stats_bycard_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type TransferStatsByCardAmountCache interface {
	GetMonthlyTransferAmountsBySenderCard(ctx context.Context, req *requests.MonthYearCardNumber) (*response.ApiResponseTransferMonthAmount, bool)
	SetMonthlyTransferAmountsBySenderCard(ctx context.Context, req *requests.MonthYearCardNumber, data *response.ApiResponseTransferMonthAmount)

	GetMonthlyTransferAmountsByReceiverCard(ctx context.Context, req *requests.MonthYearCardNumber) (*response.ApiResponseTransferMonthAmount, bool)
	SetMonthlyTransferAmountsByReceiverCard(ctx context.Context, req *requests.MonthYearCardNumber, data *response.ApiResponseTransferMonthAmount)

	GetYearlyTransferAmountsBySenderCard(ctx context.Context, req *requests.MonthYearCardNumber) (*response.ApiResponseTransferYearAmount, bool)
	SetYearlyTransferAmountsBySenderCard(ctx context.Context, req *requests.MonthYearCardNumber, data *response.ApiResponseTransferYearAmount)

	GetYearlyTransferAmountsByReceiverCard(ctx context.Context, req *requests.MonthYearCardNumber) (*response.ApiResponseTransferYearAmount, bool)
	SetYearlyTransferAmountsByReceiverCard(ctx context.Context, req *requests.MonthYearCardNumber, data *response.ApiResponseTransferYearAmount)
}

type TransferStatsByCardStatusCache interface {
	GetMonthTransferStatusSuccessByCard(ctx context.Context, req *requests.MonthStatusTransferCardNumber) (*response.ApiResponseTransferMonthStatusSuccess, bool)
	SetMonthTransferStatusSuccessByCard(ctx context.Context, req *requests.MonthStatusTransferCardNumber, data *response.ApiResponseTransferMonthStatusSuccess)

	GetYearlyTransferStatusSuccessByCard(ctx context.Context, req *requests.YearStatusTransferCardNumber) (*response.ApiResponseTransferYearStatusSuccess, bool)
	SetYearlyTransferStatusSuccessByCard(ctx context.Context, req *requests.YearStatusTransferCardNumber, data *response.ApiResponseTransferYearStatusSuccess)

	GetMonthTransferStatusFailedByCard(ctx context.Context, req *requests.MonthStatusTransferCardNumber) (*response.ApiResponseTransferMonthStatusFailed, bool)
	SetMonthTransferStatusFailedByCard(ctx context.Context, req *requests.MonthStatusTransferCardNumber, data *response.ApiResponseTransferMonthStatusFailed)

	GetYearlyTransferStatusFailedByCard(ctx context.Context, req *requests.YearStatusTransferCardNumber) (*response.ApiResponseTransferYearStatusFailed, bool)
	SetYearlyTransferStatusFailedByCard(ctx context.Context, req *requests.YearStatusTransferCardNumber, data *response.ApiResponseTransferYearStatusFailed)
}
