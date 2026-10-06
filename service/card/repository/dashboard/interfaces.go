package repositorydashboard

import "context"

// CardDashboardBalanceRepository handles total balance data for dashboard usage,
// providing both global totals and per-card-specific totals.
type CardDashboardBalanceRepository interface {
	// GetTotalBalances retrieves the total balance amount across all cards.
	//
	// Parameters:
	//   - ctx: the context for the database operation
	//
	// Returns:
	//   - A pointer to the total balance as int64, or an error if the operation fails.
	GetTotalBalances(ctx context.Context) (*int64, error)

	// GetTotalBalanceByCardNumber retrieves the total balance for a specific card number.
	//
	// Parameters:
	//   - ctx: the context for the database operation
	//   - cardNumber: the card number for which to retrieve the balance
	//
	// Returns:
	//   - A pointer to the total balance as int64, or an error if the operation fails.
	GetTotalBalanceByCardNumber(ctx context.Context, cardNumber string) (*int64, error)
}

// CardDashboardTopupRepository handles total top-up data for dashboard usage,
// providing both global totals and per-card-specific totals.
type CardDashboardTopupRepository interface {
	// GetTotalTopAmount retrieves the total top-up amount across all cards.
	//
	// Parameters:
	//   - ctx: the context for the database operation
	//
	// Returns:
	//   - A pointer to the total top-up amount as int64, or an error if the operation fails.
	GetTotalTopAmount(ctx context.Context) (*int64, error)

	// GetTotalTopupAmountByCardNumber retrieves the total top-up amount for a specific card number.
	//
	// Parameters:
	//   - ctx: the context for the database operation
	//   - cardNumber: the card number for which to retrieve the top-up total
	//
	// Returns:
	//   - A pointer to the total top-up amount as int64, or an error if the operation fails.
	GetTotalTopupAmountByCardNumber(ctx context.Context, cardNumber string) (*int64, error)
}

// CardDashboardWithdrawRepository handles total withdraw data for dashboard (global and per card).
type CardDashboardWithdrawRepository interface {
	// GetTotalWithdrawAmount retrieves the total withdraw amount across all cards.
	//
	// Parameters:
	//   - ctx: the context for the database operation
	//
	// Returns:
	//   - A pointer to the total withdraw amount as int64, or an error if the operation fails.
	GetTotalWithdrawAmount(ctx context.Context) (*int64, error)

	// GetTotalWithdrawAmountByCardNumber retrieves the total withdraw amount for a specific card number.
	//
	// Parameters:
	//   - ctx: the context for the database operation
	//   - cardNumber: the card number for which to retrieve the withdraw total
	//
	// Returns:
	//   - A pointer to the total withdraw amount as int64, or an error if the operation fails.
	GetTotalWithdrawAmountByCardNumber(ctx context.Context, cardNumber string) (*int64, error)
}

// CardDashboardTransactionRepository handles total transaction data for dashboard (global and per card).
type CardDashboardTransactionRepository interface {
	// GetTotalTransactionAmount retrieves the total transaction amount across all cards.
	//
	// Parameters:
	//   - ctx: the context for the database operation
	//
	// Returns:
	//   - A pointer to the total transaction amount as int64, or an error if the operation fails.
	GetTotalTransactionAmount(ctx context.Context) (*int64, error)

	// GetTotalTransactionAmountByCardNumber retrieves the total transaction amount for a specific card number.
	//
	// Parameters:
	//   - ctx: the context for the database operation
	//   - cardNumber: the card number for which to retrieve the transaction total
	//
	// Returns:
	//   - A pointer to the total transaction amount as int64, or an error if the operation fails.
	GetTotalTransactionAmountByCardNumber(ctx context.Context, cardNumber string) (*int64, error)
}

// CardDashboardTransferRepository handles total transfer data for dashboard (global and per sender/receiver).
type CardDashboardTransferRepository interface {
	// GetTotalTransferAmount retrieves the total transfer amount across all cards.
	//
	// Parameters:
	//   - ctx: the context for the database operation
	//
	// Returns:
	//   - A pointer to the total transfer amount as int64, or an error if the operation fails.
	GetTotalTransferAmount(ctx context.Context) (*int64, error)

	// GetTotalTransferAmountBySender retrieves the total amount transferred by a specific sender card number.
	//
	// Parameters:
	//   - ctx: the context for the database operation
	//   - senderCardNumber: the card number of the sender
	//
	// Returns:
	//   - A pointer to the total transferred amount as int64, or an error if the operation fails.
	GetTotalTransferAmountBySender(ctx context.Context, senderCardNumber string) (*int64, error)

	// GetTotalTransferAmountByReceiver retrieves the total amount received by a specific receiver card number.
	//
	// Parameters:
	//   - ctx: the context for the database operation
	//   - receiverCardNumber: the card number of the receiver
	//
	// Returns:
	//   - A pointer to the total received amount as int64, or an error if the operation fails.
	GetTotalTransferAmountByReceiver(ctx context.Context, receiverCardNumber string) (*int64, error)
}
