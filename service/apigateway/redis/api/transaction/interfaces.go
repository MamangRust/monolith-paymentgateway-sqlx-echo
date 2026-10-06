package transaction_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type TransactionQueryCache interface {
	GetCachedTransactionsCache(ctx context.Context, req *requests.FindAllTransactions) (*response.ApiResponsePaginationTransaction, bool)
	SetCachedTransactionsCache(ctx context.Context, req *requests.FindAllTransactions, data *response.ApiResponsePaginationTransaction)

	GetCachedTransactionByCardNumberCache(ctx context.Context, req *requests.FindAllTransactionCardNumber) (*response.ApiResponsePaginationTransaction, bool)
	SetCachedTransactionByCardNumberCache(ctx context.Context, req *requests.FindAllTransactionCardNumber, data *response.ApiResponsePaginationTransaction)

	GetCachedTransactionActiveCache(ctx context.Context, req *requests.FindAllTransactions) (*response.ApiResponsePaginationTransactionDeleteAt, bool)
	SetCachedTransactionActiveCache(ctx context.Context, req *requests.FindAllTransactions, data *response.ApiResponsePaginationTransactionDeleteAt)

	GetCachedTransactionTrashedCache(ctx context.Context, req *requests.FindAllTransactions) (*response.ApiResponsePaginationTransactionDeleteAt, bool)
	SetCachedTransactionTrashedCache(ctx context.Context, req *requests.FindAllTransactions, data *response.ApiResponsePaginationTransactionDeleteAt)

	GetCachedTransactionByMerchantIdCache(ctx context.Context, merchant_id int) (*response.ApiResponseTransactions, bool)
	SetCachedTransactionByMerchantIdCache(ctx context.Context, merchant_id int, data *response.ApiResponseTransactions)

	GetCachedTransactionCache(ctx context.Context, id int) (*response.ApiResponseTransaction, bool)
	SetCachedTransactionCache(ctx context.Context, data *response.ApiResponseTransaction)
}

type TransactionCommandCache interface {
	DeleteTransactionCache(ctx context.Context, id int)
}
