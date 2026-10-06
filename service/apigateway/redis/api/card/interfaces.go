package card_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type CardQueryCache interface {
	GetByIdCache(ctx context.Context, cardID int) (*response.ApiResponseCard, bool)

	GetByUserIDCache(ctx context.Context, userID int) (*response.ApiResponseCard, bool)

	GetByCardNumberCache(ctx context.Context, cardNumber string) (*response.ApiResponseCard, bool)

	GetFindAllCache(ctx context.Context, req *requests.FindAllCards) (*response.ApiResponsePaginationCard, bool)

	GetByActiveCache(ctx context.Context, req *requests.FindAllCards) (*response.ApiResponsePaginationCardDeleteAt, bool)

	GetByTrashedCache(ctx context.Context, req *requests.FindAllCards) (*response.ApiResponsePaginationCardDeleteAt, bool)

	SetByIdCache(ctx context.Context, cardID int, data *response.ApiResponseCard)

	SetByUserIDCache(ctx context.Context, userID int, data *response.ApiResponseCard)

	SetByCardNumberCache(ctx context.Context, cardNumber string, data *response.ApiResponseCard)

	SetFindAllCache(ctx context.Context, req *requests.FindAllCards, data *response.ApiResponsePaginationCard)

	SetByActiveCache(ctx context.Context, req *requests.FindAllCards, data *response.ApiResponsePaginationCardDeleteAt)

	SetByTrashedCache(ctx context.Context, req *requests.FindAllCards, data *response.ApiResponsePaginationCardDeleteAt)

	DeleteByIdCache(ctx context.Context, cardID int)
	DeleteByUserIDCache(ctx context.Context, userID int)
	DeleteByCardNumberCache(ctx context.Context, cardNumber string)
}

type CardCommandCache interface {
	DeleteCardCommandCache(ctx context.Context, id int)
}
