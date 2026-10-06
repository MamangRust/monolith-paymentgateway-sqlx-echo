package mencache

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type TransactionQueryCache interface {
	GetCachedTransactionsCache(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetTransactionsRow, *int, bool)
	SetCachedTransactionsCache(ctx context.Context, req *requests.FindAllTransactions, data []*db.GetTransactionsRow, total *int)

	GetCachedTransactionByCardNumberCache(ctx context.Context, req *requests.FindAllTransactionCardNumber) ([]*db.GetTransactionsByCardNumberRow, *int, bool)
	SetCachedTransactionByCardNumberCache(ctx context.Context, req *requests.FindAllTransactionCardNumber, data []*db.GetTransactionsByCardNumberRow, total *int)

	GetCachedTransactionActiveCache(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetActiveTransactionsRow, *int, bool)
	SetCachedTransactionActiveCache(ctx context.Context, req *requests.FindAllTransactions, data []*db.GetActiveTransactionsRow, total *int)

	GetCachedTransactionTrashedCache(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetTrashedTransactionsRow, *int, bool)
	SetCachedTransactionTrashedCache(ctx context.Context, req *requests.FindAllTransactions, data []*db.GetTrashedTransactionsRow, total *int)

	GetCachedTransactionByMerchantIdCache(ctx context.Context, merchant_id int) ([]*db.GetTransactionsByMerchantIDRow, bool)
	SetCachedTransactionByMerchantIdCache(ctx context.Context, merchant_id int, data []*db.GetTransactionsByMerchantIDRow)

	GetCachedTransactionCache(ctx context.Context, id int) (*db.GetTransactionByIDRow, bool)
	SetCachedTransactionCache(ctx context.Context, data *db.GetTransactionByIDRow)
}

type TransactionCommandCache interface {
	DeleteTransactionCache(ctx context.Context, id int)
	InvalidateTransactionCache(ctx context.Context)
}
