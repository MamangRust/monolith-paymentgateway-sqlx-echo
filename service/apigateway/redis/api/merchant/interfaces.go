package merchant_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type MerchantQueryCache interface {
	GetCachedMerchants(ctx context.Context, req *requests.FindAllMerchants) (*response.ApiResponsePaginationMerchant, bool)
	SetCachedMerchants(ctx context.Context, req *requests.FindAllMerchants, data *response.ApiResponsePaginationMerchant)

	GetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchants) (*response.ApiResponsePaginationMerchantDeleteAt, bool)
	SetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchants, data *response.ApiResponsePaginationMerchantDeleteAt)

	GetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchants) (*response.ApiResponsePaginationMerchantDeleteAt, bool)
	SetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchants, data *response.ApiResponsePaginationMerchantDeleteAt)

	GetCachedMerchant(ctx context.Context, id int) (*response.ApiResponseMerchant, bool)
	SetCachedMerchant(ctx context.Context, data *response.ApiResponseMerchant)

	GetCachedMerchantsByUserId(ctx context.Context, userId int) (*response.ApiResponsesMerchant, bool)
	SetCachedMerchantsByUserId(ctx context.Context, userId int, data *response.ApiResponsesMerchant)

	GetCachedMerchantByApiKey(ctx context.Context, apiKey string) (*response.ApiResponseMerchant, bool)
	SetCachedMerchantByApiKey(ctx context.Context, apiKey string, data *response.ApiResponseMerchant)
}

type MerchantTransactionCache interface {
	GetCacheAllMerchantTransactions(ctx context.Context, req *requests.FindAllMerchantTransactions) (*response.ApiResponsePaginationMerchantTransaction, bool)
	SetCacheAllMerchantTransactions(ctx context.Context, req *requests.FindAllMerchantTransactions, data *response.ApiResponsePaginationMerchantTransaction)

	GetCacheMerchantTransactions(ctx context.Context, req *requests.FindAllMerchantTransactionsById) (*response.ApiResponsePaginationMerchantTransaction, bool)
	SetCacheMerchantTransactions(ctx context.Context, req *requests.FindAllMerchantTransactionsById, data *response.ApiResponsePaginationMerchantTransaction)

	GetCacheMerchantTransactionApikey(ctx context.Context, req *requests.FindAllMerchantTransactionsByApiKey) (*response.ApiResponsePaginationMerchantTransaction, bool)
	SetCacheMerchantTransactionApikey(ctx context.Context, req *requests.FindAllMerchantTransactionsByApiKey, data *response.ApiResponsePaginationMerchantTransaction)
}

type MerchantCommandCache interface {
	DeleteCachedMerchant(ctx context.Context, id int)
}
