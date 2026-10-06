package topup_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type TopupQueryCache interface {
	GetCachedTopupsCache(ctx context.Context, req *requests.FindAllTopups) (*response.ApiResponsePaginationTopup, bool)
	SetCachedTopupsCache(ctx context.Context, req *requests.FindAllTopups, data *response.ApiResponsePaginationTopup)

	GetCacheTopupByCardCache(ctx context.Context, req *requests.FindAllTopupsByCardNumber) (*response.ApiResponsePaginationTopup, bool)
	SetCacheTopupByCardCache(ctx context.Context, req *requests.FindAllTopupsByCardNumber, data *response.ApiResponsePaginationTopup)

	GetCachedTopupActiveCache(ctx context.Context, req *requests.FindAllTopups) (*response.ApiResponsePaginationTopupDeleteAt, bool)
	SetCachedTopupActiveCache(ctx context.Context, req *requests.FindAllTopups, data *response.ApiResponsePaginationTopupDeleteAt)

	GetCachedTopupTrashedCache(ctx context.Context, req *requests.FindAllTopups) (*response.ApiResponsePaginationTopupDeleteAt, bool)
	SetCachedTopupTrashedCache(ctx context.Context, req *requests.FindAllTopups, data *response.ApiResponsePaginationTopupDeleteAt)

	GetCachedTopupCache(ctx context.Context, id int) (*response.ApiResponseTopup, bool)
	SetCachedTopupCache(ctx context.Context, data *response.ApiResponseTopup)
}

type TopupCommandCache interface {
	DeleteCachedTopupCache(ctx context.Context, id int)
}
