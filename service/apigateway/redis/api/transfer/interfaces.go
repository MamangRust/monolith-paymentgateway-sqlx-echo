package transfer_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type TransferQueryCache interface {
	GetCachedTransfersCache(ctx context.Context, req *requests.FindAllTransfers) (*response.ApiResponsePaginationTransfer, bool)
	SetCachedTransfersCache(ctx context.Context, req *requests.FindAllTransfers, data *response.ApiResponsePaginationTransfer)

	GetCachedTransferActiveCache(ctx context.Context, req *requests.FindAllTransfers) (*response.ApiResponsePaginationTransferDeleteAt, bool)
	SetCachedTransferActiveCache(ctx context.Context, req *requests.FindAllTransfers, data *response.ApiResponsePaginationTransferDeleteAt)

	GetCachedTransferTrashedCache(ctx context.Context, req *requests.FindAllTransfers) (*response.ApiResponsePaginationTransferDeleteAt, bool)
	SetCachedTransferTrashedCache(ctx context.Context, req *requests.FindAllTransfers, data *response.ApiResponsePaginationTransferDeleteAt)

	GetCachedTransferCache(ctx context.Context, id int) (*response.ApiResponseTransfer, bool)
	SetCachedTransferCache(ctx context.Context, data *response.ApiResponseTransfer)

	GetCachedTransferByFrom(ctx context.Context, from string) (*response.ApiResponseTransfers, bool)
	SetCachedTransferByFrom(ctx context.Context, from string, data *response.ApiResponseTransfers)

	GetCachedTransferByTo(ctx context.Context, to string) (*response.ApiResponseTransfers, bool)
	SetCachedTransferByTo(ctx context.Context, to string, data *response.ApiResponseTransfers)
}

type TransferCommandCache interface {
	DeleteTransferCache(ctx context.Context, id int)
}
