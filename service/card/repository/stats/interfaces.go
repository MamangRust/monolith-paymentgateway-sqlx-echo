package repositorystats

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
)

type CardStatsBalanceRepository interface {
	GetMonthlyBalance(ctx context.Context, year int) ([]*db.GetMonthlyBalancesRow, error)
	GetYearlyBalance(ctx context.Context, year int) ([]*db.GetYearlyBalancesRow, error)
}

type CardStatsTopupRepository interface {
	GetMonthlyTopupAmount(ctx context.Context, year int) ([]*db.GetMonthlyTopupAmountRow, error)
	GetYearlyTopupAmount(ctx context.Context, year int) ([]*db.GetYearlyTopupAmountRow, error)
}

type CardStatsWithdrawRepository interface {
	GetMonthlyWithdrawAmount(ctx context.Context, year int) ([]*db.GetMonthlyWithdrawAmountRow, error)
	GetYearlyWithdrawAmount(ctx context.Context, year int) ([]*db.GetYearlyWithdrawAmountRow, error)
}

type CardStatsTransactionRepository interface {
	GetMonthlyTransactionAmount(ctx context.Context, year int) ([]*db.GetMonthlyTransactionAmountRow, error)
	GetYearlyTransactionAmount(ctx context.Context, year int) ([]*db.GetYearlyTransactionAmountRow, error)
}

type CardStatsTransferRepository interface {
	GetMonthlyTransferAmountSender(ctx context.Context, year int) ([]*db.GetMonthlyTransferAmountSenderRow, error)
	GetYearlyTransferAmountSender(ctx context.Context, year int) ([]*db.GetYearlyTransferAmountSenderRow, error)
	GetMonthlyTransferAmountReceiver(ctx context.Context, year int) ([]*db.GetMonthlyTransferAmountReceiverRow, error)
	GetYearlyTransferAmountReceiver(ctx context.Context, year int) ([]*db.GetYearlyTransferAmountReceiverRow, error)
}
