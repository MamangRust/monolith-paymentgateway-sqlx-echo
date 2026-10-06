package card_stats_bycard_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type CardStatsBalanceByCardCache interface {
	GetMonthlyBalanceByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseMonthlyBalance, bool)
	GetYearlyBalanceByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseYearlyBalance, bool)

	SetMonthlyBalanceByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseMonthlyBalance)
	SetYearlyBalanceByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseYearlyBalance)
}

type CardStatsTopupByCardCache interface {
	GetMonthlyTopupByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseMonthlyAmount, bool)
	GetYearlyTopupByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseYearlyAmount, bool)

	SetMonthlyTopupByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseMonthlyAmount)
	SetYearlyTopupByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseYearlyAmount)
}

type CardStatsWithdrawByCardCache interface {
	GetMonthlyWithdrawByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseMonthlyAmount, bool)
	GetYearlyWithdrawByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseYearlyAmount, bool)

	SetMonthlyWithdrawByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseMonthlyAmount)
	SetYearlyWithdrawByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseYearlyAmount)
}

type CardStatsTransactionByCardCache interface {
	GetMonthlyTransactionByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseMonthlyAmount, bool)
	GetYearlyTransactionByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseYearlyAmount, bool)

	SetMonthlyTransactionByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseMonthlyAmount)
	SetYearlyTransactionByNumberCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseYearlyAmount)
}

type CardStatsTransferByCardCache interface {
	GetMonthlyTransferBySenderCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseMonthlyAmount, bool)
	GetYearlyTransferBySenderCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseYearlyAmount, bool)

	SetMonthlyTransferBySenderCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseMonthlyAmount)
	SetYearlyTransferBySenderCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseYearlyAmount)

	GetMonthlyTransferByReceiverCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseMonthlyAmount, bool)
	GetYearlyTransferByReceiverCache(ctx context.Context, req *requests.MonthYearCardNumberCard) (*response.ApiResponseYearlyAmount, bool)

	SetMonthlyTransferByReceiverCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseMonthlyAmount)
	SetYearlyTransferByReceiverCache(ctx context.Context, req *requests.MonthYearCardNumberCard, data *response.ApiResponseYearlyAmount)
}
