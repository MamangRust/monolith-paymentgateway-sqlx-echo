package cardstatsservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
)

// CardStatsBalanceService handles balance statistics globally and per specific card number.
type CardStatsBalanceService interface {
	FindMonthlyBalance(ctx context.Context, year int) ([]*db.GetMonthlyBalancesRow, error)
	FindYearlyBalance(ctx context.Context, year int) ([]*db.GetYearlyBalancesRow, error)
}

// CardStatsTopupService handles top-up statistics globally and per specific card number.
type CardStatsTopupService interface {
	FindMonthlyTopupAmount(ctx context.Context, year int) ([]*db.GetMonthlyTopupAmountRow, error)
	FindYearlyTopupAmount(ctx context.Context, year int) ([]*db.GetYearlyTopupAmountRow, error)
}

// CardStatsWithdrawService handles withdraw statistics globally and per specific card number.
type CardStatsWithdrawService interface {
	FindMonthlyWithdrawAmount(ctx context.Context, year int) ([]*db.GetMonthlyWithdrawAmountRow, error)
	FindYearlyWithdrawAmount(ctx context.Context, year int) ([]*db.GetYearlyWithdrawAmountRow, error)
}

// CardStatsTransactionService handles transaction statistics globally and per specific card number.
type CardStatsTransactionService interface {
	FindMonthlyTransactionAmount(ctx context.Context, year int) ([]*db.GetMonthlyTransactionAmountRow, error)
	FindYearlyTransactionAmount(ctx context.Context, year int) ([]*db.GetYearlyTransactionAmountRow, error)
}

// CardStatsTransferService handles transfer statistics globally and per specific card number (as sender or receiver).
type CardStatsTransferService interface {
	FindMonthlyTransferAmountSender(ctx context.Context, year int) ([]*db.GetMonthlyTransferAmountSenderRow, error)
	FindYearlyTransferAmountSender(ctx context.Context, year int) ([]*db.GetYearlyTransferAmountSenderRow, error)
	FindMonthlyTransferAmountReceiver(ctx context.Context, year int) ([]*db.GetMonthlyTransferAmountReceiverRow, error)
	FindYearlyTransferAmountReceiver(ctx context.Context, year int) ([]*db.GetYearlyTransferAmountReceiverRow, error)
}
