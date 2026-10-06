package withdraw_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type WithdrawQueryCache interface {
	GetCachedWithdrawsCache(ctx context.Context, req *requests.FindAllWithdraws) (*response.ApiResponsePaginationWithdraw, bool)
	SetCachedWithdrawsCache(ctx context.Context, req *requests.FindAllWithdraws, data *response.ApiResponsePaginationWithdraw)

	GetCachedWithdrawByCardCache(ctx context.Context, req *requests.FindAllWithdrawCardNumber) (*response.ApiResponsePaginationWithdraw, bool)
	SetCachedWithdrawByCardCache(ctx context.Context, req *requests.FindAllWithdrawCardNumber, data *response.ApiResponsePaginationWithdraw)

	GetCachedWithdrawActiveCache(ctx context.Context, req *requests.FindAllWithdraws) (*response.ApiResponsePaginationWithdrawDeleteAt, bool)
	SetCachedWithdrawActiveCache(ctx context.Context, req *requests.FindAllWithdraws, data *response.ApiResponsePaginationWithdrawDeleteAt)

	GetCachedWithdrawTrashedCache(ctx context.Context, req *requests.FindAllWithdraws) (*response.ApiResponsePaginationWithdrawDeleteAt, bool)
	SetCachedWithdrawTrashedCache(ctx context.Context, req *requests.FindAllWithdraws, data *response.ApiResponsePaginationWithdrawDeleteAt)

	GetCachedWithdrawCache(ctx context.Context, id int) (*response.ApiResponseWithdraw, bool)
	SetCachedWithdrawCache(ctx context.Context, data *response.ApiResponseWithdraw)
}

type WithdrawCommandCache interface {
	DeleteCachedWithdrawCache(ctx context.Context, id int)
}
