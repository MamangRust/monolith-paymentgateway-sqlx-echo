package mencache

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type MerchantQueryCache interface {
	GetCachedMerchants(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetMerchantsRow, *int, bool)
	SetCachedMerchants(ctx context.Context, req *requests.FindAllMerchants, data []*db.GetMerchantsRow, total *int)

	GetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetActiveMerchantsRow, *int, bool)
	SetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchants, data []*db.GetActiveMerchantsRow, total *int)

	GetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetTrashedMerchantsRow, *int, bool)
	SetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchants, data []*db.GetTrashedMerchantsRow, total *int)

	GetCachedMerchant(ctx context.Context, id int) (*db.GetMerchantByIDRow, bool)
	SetCachedMerchant(ctx context.Context, data *db.GetMerchantByIDRow)

	GetCachedMerchantsByUserId(ctx context.Context, userId int) ([]*db.GetMerchantsByUserIDRow, bool)
	SetCachedMerchantsByUserId(ctx context.Context, userId int, data []*db.GetMerchantsByUserIDRow)

	GetCachedMerchantByApiKey(ctx context.Context, apiKey string) (*db.GetMerchantByApiKeyRow, bool)
	SetCachedMerchantByApiKey(ctx context.Context, apiKey string, data *db.GetMerchantByApiKeyRow)
}

type MerchantDocumentQueryCache interface {
	GetCachedMerchantDocuments(ctx context.Context, req *requests.FindAllMerchantDocuments) ([]*db.GetMerchantDocumentsRow, *int, bool)
	SetCachedMerchantDocuments(ctx context.Context, req *requests.FindAllMerchantDocuments, data []*db.GetMerchantDocumentsRow, total *int)

	GetCachedMerchantDocumentsActive(ctx context.Context, req *requests.FindAllMerchantDocuments) ([]*db.GetActiveMerchantDocumentsRow, *int, bool)
	SetCachedMerchantDocumentsActive(ctx context.Context, req *requests.FindAllMerchantDocuments, data []*db.GetActiveMerchantDocumentsRow, total *int)

	GetCachedMerchantDocumentsTrashed(ctx context.Context, req *requests.FindAllMerchantDocuments) ([]*db.GetTrashedMerchantDocumentsRow, *int, bool)
	SetCachedMerchantDocumentsTrashed(ctx context.Context, req *requests.FindAllMerchantDocuments, data []*db.GetTrashedMerchantDocumentsRow, total *int)

	GetCachedMerchantDocument(ctx context.Context, id int) (*db.GetMerchantDocumentRow, bool)
	SetCachedMerchantDocument(ctx context.Context, id int, data *db.GetMerchantDocumentRow)
}

type MerchantCommandCache interface {
	DeleteCachedMerchant(ctx context.Context, id int)
	InvalidateMerchantListCaches(ctx context.Context)
	DeleteCachedMerchantByUserID(ctx context.Context, userID int)
	DeleteCachedMerchantByAPIKey(ctx context.Context, apiKey string)
}

type MerchantDocumentCommandCache interface {
	DeleteCachedMerchantDocuments(ctx context.Context, id int)
}

type MerchantTransactionCache interface {
	GetCacheAllMerchantTransactions(ctx context.Context, req *requests.FindAllMerchantTransactions) ([]*db.FindAllTransactionsRow, *int, bool)
	SetCacheAllMerchantTransactions(ctx context.Context, req *requests.FindAllMerchantTransactions, data []*db.FindAllTransactionsRow, total *int)

	GetCacheMerchantTransactions(ctx context.Context, req *requests.FindAllMerchantTransactionsById) ([]*db.FindAllTransactionsByMerchantRow, *int, bool)
	SetCacheMerchantTransactions(ctx context.Context, req *requests.FindAllMerchantTransactionsById, data []*db.FindAllTransactionsByMerchantRow, total *int)

	GetCacheMerchantTransactionApikey(ctx context.Context, req *requests.FindAllMerchantTransactionsByApiKey) ([]*db.FindAllTransactionsByApikeyRow, *int, bool)
	SetCacheMerchantTransactionApikey(ctx context.Context, req *requests.FindAllMerchantTransactionsByApiKey, data []*db.FindAllTransactionsByApikeyRow, total *int)
}
