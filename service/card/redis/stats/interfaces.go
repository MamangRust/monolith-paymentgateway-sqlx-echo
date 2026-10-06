package cardstatsmencache

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
)

type CardStatsBalanceCache interface {
	GetMonthlyBalanceCache(ctx context.Context, year int) ([]*db.GetMonthlyBalancesRow, bool)
	SetMonthlyBalanceCache(ctx context.Context, year int, data []*db.GetMonthlyBalancesRow)

	GetYearlyBalanceCache(ctx context.Context, year int) ([]*db.GetYearlyBalancesRow, bool)
	SetYearlyBalanceCache(ctx context.Context, year int, data []*db.GetYearlyBalancesRow)
}

type CardStatsTopupCache interface {
	GetMonthlyTopupCache(ctx context.Context, year int) ([]*db.GetMonthlyTopupAmountRow, bool)
	SetMonthlyTopupCache(ctx context.Context, year int, data []*db.GetMonthlyTopupAmountRow)

	GetYearlyTopupCache(ctx context.Context, year int) ([]*db.GetYearlyTopupAmountRow, bool)
	SetYearlyTopupCache(ctx context.Context, year int, data []*db.GetYearlyTopupAmountRow)
}

type CardStatsWithdrawCache interface {
	GetMonthlyWithdrawCache(ctx context.Context, year int) ([]*db.GetMonthlyWithdrawAmountRow, bool)
	SetMonthlyWithdrawCache(ctx context.Context, year int, data []*db.GetMonthlyWithdrawAmountRow)

	GetYearlyWithdrawCache(ctx context.Context, year int) ([]*db.GetYearlyWithdrawAmountRow, bool)
	SetYearlyWithdrawCache(ctx context.Context, year int, data []*db.GetYearlyWithdrawAmountRow)
}

type CardStatsTransactionCache interface {
	GetMonthlyTransactionCache(ctx context.Context, year int) ([]*db.GetMonthlyTransactionAmountRow, bool)
	SetMonthlyTransactionCache(ctx context.Context, year int, data []*db.GetMonthlyTransactionAmountRow)

	GetYearlyTransactionCache(ctx context.Context, year int) ([]*db.GetYearlyTransactionAmountRow, bool)
	SetYearlyTransactionCache(ctx context.Context, year int, data []*db.GetYearlyTransactionAmountRow)
}

type CardStatsTransferCache interface {
	GetMonthlyTransferSenderCache(ctx context.Context, year int) ([]*db.GetMonthlyTransferAmountSenderRow, bool)
	SetMonthlyTransferSenderCache(ctx context.Context, year int, data []*db.GetMonthlyTransferAmountSenderRow)

	GetYearlyTransferSenderCache(ctx context.Context, year int) ([]*db.GetYearlyTransferAmountSenderRow, bool)
	SetYearlyTransferSenderCache(ctx context.Context, year int, data []*db.GetYearlyTransferAmountSenderRow)

	GetMonthlyTransferReceiverCache(ctx context.Context, year int) ([]*db.GetMonthlyTransferAmountReceiverRow, bool)
	SetMonthlyTransferReceiverCache(ctx context.Context, year int, data []*db.GetMonthlyTransferAmountReceiverRow)

	GetYearlyTransferReceiverCache(ctx context.Context, year int) ([]*db.GetYearlyTransferAmountReceiverRow, bool)
	SetYearlyTransferReceiverCache(ctx context.Context, year int, data []*db.GetYearlyTransferAmountReceiverRow)
}
