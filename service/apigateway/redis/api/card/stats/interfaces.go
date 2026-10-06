package card_stats_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type CardStatsBalanceCache interface {
	GetMonthlyBalanceCache(ctx context.Context, year int) (*response.ApiResponseMonthlyBalance, bool)
	SetMonthlyBalanceCache(ctx context.Context, year int, data *response.ApiResponseMonthlyBalance)

	GetYearlyBalanceCache(ctx context.Context, year int) (*response.ApiResponseYearlyBalance, bool)
	SetYearlyBalanceCache(ctx context.Context, year int, data *response.ApiResponseYearlyBalance)
}

type CardStatsTopupCache interface {
	GetMonthlyTopupCache(ctx context.Context, year int) (*response.ApiResponseMonthlyAmount, bool)
	SetMonthlyTopupCache(ctx context.Context, year int, data *response.ApiResponseMonthlyAmount)

	GetYearlyTopupCache(ctx context.Context, year int) (*response.ApiResponseYearlyAmount, bool)
	SetYearlyTopupCache(ctx context.Context, year int, data *response.ApiResponseYearlyAmount)
}

type CardStatsWithdrawCache interface {
	GetMonthlyWithdrawCache(ctx context.Context, year int) (*response.ApiResponseMonthlyAmount, bool)
	SetMonthlyWithdrawCache(ctx context.Context, year int, data *response.ApiResponseMonthlyAmount)

	GetYearlyWithdrawCache(ctx context.Context, year int) (*response.ApiResponseYearlyAmount, bool)
	SetYearlyWithdrawCache(ctx context.Context, year int, data *response.ApiResponseYearlyAmount)
}

type CardStatsTransactionCache interface {
	GetMonthlyTransactionCache(ctx context.Context, year int) (*response.ApiResponseMonthlyAmount, bool)
	SetMonthlyTransactionCache(ctx context.Context, year int, data *response.ApiResponseMonthlyAmount)

	GetYearlyTransactionCache(ctx context.Context, year int) (*response.ApiResponseYearlyAmount, bool)
	SetYearlyTransactionCache(ctx context.Context, year int, data *response.ApiResponseYearlyAmount)
}

type CardStatsTransferCache interface {
	GetMonthlyTransferSenderCache(ctx context.Context, year int) (*response.ApiResponseMonthlyAmount, bool)
	SetMonthlyTransferSenderCache(ctx context.Context, year int, data *response.ApiResponseMonthlyAmount)

	GetYearlyTransferSenderCache(ctx context.Context, year int) (*response.ApiResponseYearlyAmount, bool)
	SetYearlyTransferSenderCache(ctx context.Context, year int, data *response.ApiResponseYearlyAmount)

	GetMonthlyTransferReceiverCache(ctx context.Context, year int) (*response.ApiResponseMonthlyAmount, bool)
	SetMonthlyTransferReceiverCache(ctx context.Context, year int, data *response.ApiResponseMonthlyAmount)

	GetYearlyTransferReceiverCache(ctx context.Context, year int) (*response.ApiResponseYearlyAmount, bool)
	SetYearlyTransferReceiverCache(ctx context.Context, year int, data *response.ApiResponseYearlyAmount)
}
