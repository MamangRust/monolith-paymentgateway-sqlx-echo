package service

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

// TransactionQueryService handles queries related to transactions.
type TransactionQueryService interface {
	FindAll(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetTransactionsRow, *int, error)
	FindAllByCardNumber(ctx context.Context, req *requests.FindAllTransactionCardNumber) ([]*db.GetTransactionsByCardNumberRow, *int, error)
	FindById(ctx context.Context, transactionID int) (*db.GetTransactionByIDRow, error)
	FindByActive(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetActiveTransactionsRow, *int, error)
	FindByTrashed(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetTrashedTransactionsRow, *int, error)
	FindTransactionByMerchantId(ctx context.Context, merchant_id int) ([]*db.GetTransactionsByMerchantIDRow, error)
}

type TransactionCommandService interface {
	Create(ctx context.Context, apiKey string, request *requests.CreateTransactionRequest) (*db.UpdateTransactionStatusRow, error)
	Update(ctx context.Context, apiKey string, request *requests.UpdateTransactionRequest) (*db.UpdateTransactionRow, error)
	TrashedTransaction(ctx context.Context, transaction_id int) (*db.TrashTransactionRow, error)
	RestoreTransaction(ctx context.Context, transaction_id int) (*db.RestoreTransactionRow, error)
	DeleteTransactionPermanent(ctx context.Context, transaction_id int) (bool, error)

	RestoreAllTransaction(ctx context.Context) (bool, error)
	DeleteAllTransactionPermanent(ctx context.Context) (bool, error)

	AuthorizeTransaction(ctx context.Context, apiKey string, request *requests.CreateTransactionRequest) (*db.UpdateTransactionStatusRow, error)
	CaptureTransaction(ctx context.Context, apiKey string, transactionID int) (*db.UpdateTransactionStatusRow, error)
	VoidTransaction(ctx context.Context, apiKey string, transactionID int) (*db.UpdateTransactionStatusRow, error)
	RefundTransaction(ctx context.Context, apiKey string, transactionID int) (*db.UpdateTransactionStatusRow, error)
}
