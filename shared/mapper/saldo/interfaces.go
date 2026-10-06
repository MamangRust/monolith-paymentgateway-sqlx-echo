package saldoapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	pbstats "github.com/MamangRust/monolith-payment-gateway-pb/saldo/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type SaldoBaseResponseMapper interface {
	// Converts a single saldo gRPC response to an API response.
	ToApiResponseSaldo(pbResponse *pb.ApiResponseSaldo) *response.ApiResponseSaldo
}

type SaldoQueryResponseMapper interface {
	SaldoBaseResponseMapper

	// Converts paginated saldo responses into an API response.
	ToApiResponsePaginationSaldo(pbResponse *pb.ApiResponsePaginationSaldo) *response.ApiResponsePaginationSaldo

	// Converts paginated soft-deleted saldo responses into an API response.
	ToApiResponsePaginationSaldoDeleteAt(pbResponse *pb.ApiResponsePaginationSaldoDeleteAt) *response.ApiResponsePaginationSaldoDeleteAt
}

type SaldoCommandResponseMapper interface {
	SaldoBaseResponseMapper

	ToApiResponseSaldoDeleteAt(pbResponse *pb.ApiResponseSaldoDeleteAt) *response.ApiResponseSaldoDeleteAt

	// Converts a deleted saldo response into an API response.
	ToApiResponseSaldoDelete(pbResponse *pb.ApiResponseSaldoDelete) *response.ApiResponseSaldoDelete

	// Converts all saldo records into a comprehensive API response.
	ToApiResponseSaldoAll(pbResponse *pb.ApiResponseSaldoAll) *response.ApiResponseSaldoAll
}

type SaldoStatsTotalResponseMapper interface {
	// Converts monthly total saldo values into an API response.
	ToApiResponseMonthTotalSaldo(pbResponse *pbstats.ApiResponseMonthTotalSaldo) *response.ApiResponseMonthTotalSaldo

	// Converts yearly total saldo values into an API response.
	ToApiResponseYearTotalSaldo(pbResponse *pbstats.ApiResponseYearTotalSaldo) *response.ApiResponseYearTotalSaldo
}

type SaldoStatsBalanceResponseMapper interface {
	// Converts monthly saldo balances into an API response.
	ToApiResponseMonthSaldoBalances(pbResponse *pbstats.ApiResponseMonthSaldoBalances) *response.ApiResponseMonthSaldoBalances

	// Converts yearly saldo balances into an API response.
	ToApiResponseYearSaldoBalances(pbResponse *pbstats.ApiResponseYearSaldoBalances) *response.ApiResponseYearSaldoBalances
}
