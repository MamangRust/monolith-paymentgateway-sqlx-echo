package saldo_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type SaldoQueryCache interface {
	GetCachedSaldos(ctx context.Context, req *requests.FindAllSaldos) (*response.ApiResponsePaginationSaldo, bool)
	SetCachedSaldos(ctx context.Context, req *requests.FindAllSaldos, data *response.ApiResponsePaginationSaldo)

	GetCachedSaldoById(ctx context.Context, saldo_id int) (*response.ApiResponseSaldo, bool)
	SetCachedSaldoById(ctx context.Context, saldo_id int, data *response.ApiResponseSaldo)

	GetCachedSaldoByCardNumber(ctx context.Context, card_number string) (*response.ApiResponseSaldo, bool)
	SetCachedSaldoByCardNumber(ctx context.Context, card_number string, data *response.ApiResponseSaldo)

	GetCachedSaldoByActive(ctx context.Context, req *requests.FindAllSaldos) (*response.ApiResponsePaginationSaldoDeleteAt, bool)
	SetCachedSaldoByActive(ctx context.Context, req *requests.FindAllSaldos, data *response.ApiResponsePaginationSaldoDeleteAt)

	GetCachedSaldoByTrashed(ctx context.Context, req *requests.FindAllSaldos) (*response.ApiResponsePaginationSaldoDeleteAt, bool)
	SetCachedSaldoByTrashed(ctx context.Context, req *requests.FindAllSaldos, data *response.ApiResponsePaginationSaldoDeleteAt)
}

type SaldoCommandCache interface {
	DeleteSaldoCache(ctx context.Context, saldo_id int)
}
