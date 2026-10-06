package merchantdocument_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type MerchantDocumentQueryCache interface {
	GetCachedMerchants(ctx context.Context, req *requests.FindAllMerchantDocuments) (*response.ApiResponsePaginationMerchantDocument, bool)
	SetCachedMerchants(ctx context.Context, req *requests.FindAllMerchantDocuments, data *response.ApiResponsePaginationMerchantDocument)

	GetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchantDocuments) (*response.ApiResponsePaginationMerchantDocumentDeleteAt, bool)
	SetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchantDocuments, data *response.ApiResponsePaginationMerchantDocumentDeleteAt)

	GetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchantDocuments) (*response.ApiResponsePaginationMerchantDocumentDeleteAt, bool)
	SetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchantDocuments, data *response.ApiResponsePaginationMerchantDocumentDeleteAt)

	GetCachedMerchant(ctx context.Context, id int) (*response.ApiResponseMerchantDocument, bool)
	SetCachedMerchant(ctx context.Context, data *response.ApiResponseMerchantDocument)
}

type MerchantDocumentCommandCache interface {
	DeleteCachedMerchantDocument(ctx context.Context, id int)
}
