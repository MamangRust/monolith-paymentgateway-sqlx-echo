package repository

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type MerchantRepository interface {
	FindByApiKey(ctx context.Context, api_key string) (*db.GetMerchantByApiKeyRow, error)
}

type SaldoRepository interface {
	FindByCardNumber(ctx context.Context, card_number string) (*db.Saldo, error)

	UpdateSaldoBalance(ctx context.Context, request *requests.UpdateSaldoBalance) (*db.UpdateSaldoBalanceRow, error)
}

type CardRepository interface {
	FindCardByUserId(ctx context.Context, user_id int) (*db.GetCardByUserIDRow, error)

	FindUserCardByCardNumber(ctx context.Context, card_number string) (*db.GetUserEmailByCardNumberRow, error)

	FindCardByCardNumber(ctx context.Context, card_number string) (*db.GetCardByCardNumberRow, error)

	UpdateCard(ctx context.Context, request *requests.UpdateCardRequest) (*db.UpdateCardRow, error)

	UpdateCardOutstandingBalance(ctx context.Context, cardID int, outstandingBalance int) (*db.UpdateOutstandingBalanceRow, error)

	AddRewardPoints(ctx context.Context, cardID int, points int) (*db.AddRewardPointsRow, error)
}

type TransactionQueryRepository interface {
	FindAllTransactions(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetTransactionsRow, error)
	FindByActive(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetActiveTransactionsRow, error)
	FindByTrashed(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetTrashedTransactionsRow, error)
	FindAllTransactionByCardNumber(ctx context.Context, req *requests.FindAllTransactionCardNumber) ([]*db.GetTransactionsByCardNumberRow, error)
	FindById(ctx context.Context, transaction_id int) (*db.GetTransactionByIDRow, error)
	FindTransactionByMerchantId(ctx context.Context, merchant_id int) ([]*db.GetTransactionsByMerchantIDRow, error)
}

type TransactionCommandRepository interface {
	// CreateTransactionAtomic debits the customer, credits the merchant, and records
	// the transaction in one database statement. It returns no row when the
	// settlement preconditions are not met.
	CreateTransactionAtomic(ctx context.Context, params db.CreateTransactionAtomicParams) (*db.CreateTransactionAtomicRow, error)
	AuthorizeTransactionAtomic(ctx context.Context, params db.AuthorizeTransactionAtomicParams) (*db.AuthorizeTransactionAtomicRow, error)
	AuthorizeTransactionAtomicIdempotent(ctx context.Context, params db.AuthorizeTransactionAtomicParams) (*AuthorizeTransactionResult, error)
	// CreateTransactionAtomicIdempotent adds safe replay semantics: a previously
	// persisted idempotency key returns the original record, and a concurrent
	// duplicate is resolved by returning the existing record (the atomic
	// statement rolls back fully on insert failure).
	CreateTransactionAtomicIdempotent(ctx context.Context, params db.CreateTransactionAtomicParams) (*TransactionAtomicResult, error)
	// UpdateSaldoBalanceDelta applies an atomic signed delta to a card balance
	// (400 when the resulting balance would be negative).
	UpdateSaldoBalanceDelta(ctx context.Context, cardNumber string, delta int) error
	// UpdateTransactionAtomic updates a transaction and settles its card balance
	// in one SQL statement (no lost update under concurrency).
	UpdateTransactionAtomic(ctx context.Context, params db.UpdateTransactionAtomicParams) (*db.UpdateTransactionAtomicRow, error)
	// CaptureTransactionAtomic captures an authorized transaction and credits the
	// merchant atomically (double-capture impossible).
	CaptureTransactionAtomic(ctx context.Context, params db.CaptureTransactionAtomicParams) (*db.CaptureTransactionAtomicRow, error)
	// VoidTransactionAtomic voids a pending/authorized transaction and releases
	// the hold atomically (double-void impossible).
	VoidTransactionAtomic(ctx context.Context, transactionID int32) (*db.VoidTransactionAtomicRow, error)
	// RefundTransactionAtomic refunds a captured transaction and reverses the
	// settlement atomically (double-refund impossible).
	RefundTransactionAtomic(ctx context.Context, params db.RefundTransactionAtomicParams) (*db.RefundTransactionAtomicRow, error)
	CreateTransaction(ctx context.Context, request *requests.CreateTransactionRequest) (*db.CreateTransactionRow, error)
	UpdateTransaction(ctx context.Context, request *requests.UpdateTransactionRequest) (*db.UpdateTransactionRow, error)
	UpdateTransactionStatus(ctx context.Context, request *requests.UpdateTransactionStatus) (*db.UpdateTransactionStatusRow, error)
	TrashedTransaction(ctx context.Context, transaction_id int) (*db.TrashTransactionRow, error)
	RestoreTransaction(ctx context.Context, topup_id int) (*db.RestoreTransactionRow, error)
	DeleteTransactionPermanent(ctx context.Context, topup_id int) (bool, error)

	RestoreAllTransaction(ctx context.Context) (bool, error)
	DeleteAllTransactionPermanent(ctx context.Context) (bool, error)
}
