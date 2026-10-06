package saldoapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper"
)

type saldoQueryResponseMapper struct {
}

func NewSaldoQueryResponseMapper() SaldoQueryResponseMapper {
	return &saldoQueryResponseMapper{}
}

// ToApiResponseSaldo maps a gRPC saldo response to an HTTP API response. It constructs an ApiResponseSaldo by copying the status and message fields and mapping the saldo data to a SaldoResponse.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponseSaldo containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseSaldo with mapped data.
func (s *saldoQueryResponseMapper) ToApiResponseSaldo(pbResponse *pb.ApiResponseSaldo) *response.ApiResponseSaldo {
	return &response.ApiResponseSaldo{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseSaldo(pbResponse.Data),
	}
}

// ToApiResponsePaginationSaldo maps a paginated gRPC response of saldo records to an HTTP API response.
// It constructs an ApiResponsePaginationSaldo by copying the status and message fields,
// mapping the saldo data slice to a slice of SaldoResponse, and including pagination metadata.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponsePaginationSaldo containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsePaginationSaldo with mapped data and pagination info.
func (s *saldoQueryResponseMapper) ToApiResponsePaginationSaldo(pbResponse *pb.ApiResponsePaginationSaldo) *response.ApiResponsePaginationSaldo {
	return &response.ApiResponsePaginationSaldo{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesSaldo(pbResponse.Data),
		Pagination: apimapper.MapPaginationMeta(pbResponse.PaginationMeta),
	}
}

// ToApiResponsePaginationSaldoDeleteAt maps a paginated gRPC response of soft-deleted saldo records
// to an HTTP API response. It constructs an ApiResponsePaginationSaldoDeleteAt by copying the
// status and message fields, mapping the saldo data slice to a slice of SaldoResponseDeleteAt,
// and including pagination metadata.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponsePaginationSaldoDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsePaginationSaldoDeleteAt with mapped data and pagination info.
func (s *saldoQueryResponseMapper) ToApiResponsePaginationSaldoDeleteAt(pbResponse *pb.ApiResponsePaginationSaldoDeleteAt) *response.ApiResponsePaginationSaldoDeleteAt {
	return &response.ApiResponsePaginationSaldoDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesSaldoDeleteAt(pbResponse.Data),
		Pagination: apimapper.MapPaginationMeta(pbResponse.PaginationMeta),
	}
}

// mapResponseSaldo maps a gRPC SaldoResponse to an HTTP API SaldoResponse.
// It constructs a SaldoResponse by converting and copying relevant fields
// such as ID, CardNumber, TotalBalance, WithdrawTime, WithdrawAmount,
// CreatedAt, and UpdatedAt.
//
// Args:
//
//	saldo: A pointer to a pb.SaldoResponse containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.SaldoResponse with mapped data.
func (s *saldoQueryResponseMapper) mapResponseSaldo(saldo *pb.SaldoResponse) *response.SaldoResponse {
	return &response.SaldoResponse{
		ID:             int(saldo.SaldoId),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime,
		WithdrawAmount: int(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
	}
}

// mapResponsesSaldo maps a slice of gRPC SaldoResponses to a slice of HTTP API SaldoResponses.
// It constructs a slice of SaldoResponse by converting and copying relevant fields
// such as ID, CardNumber, TotalBalance, WithdrawTime, WithdrawAmount,
// CreatedAt, and UpdatedAt from each gRPC SaldoResponse.
//
// Args:
//
//	saldos: A slice of pointers to pb.SaldoResponse containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.SaldoResponse with mapped data.
func (s *saldoQueryResponseMapper) mapResponsesSaldo(saldos []*pb.SaldoResponse) []*response.SaldoResponse {
	var responseSaldos []*response.SaldoResponse

	for _, saldo := range saldos {
		responseSaldos = append(responseSaldos, s.mapResponseSaldo(saldo))
	}

	return responseSaldos
}

// mapResponseSaldoDeleteAt maps a gRPC SaldoResponseDeleteAt to an HTTP API SaldoResponseDeleteAt.
// It constructs a SaldoResponseDeleteAt by converting and copying relevant fields
// such as ID, CardNumber, TotalBalance, WithdrawTime, WithdrawAmount,
// CreatedAt, UpdatedAt, and DeletedAt.
//
// Args:
//
//	saldo: A pointer to a pb.SaldoResponseDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.SaldoResponseDeleteAt with mapped data.
func (s *saldoQueryResponseMapper) mapResponseSaldoDeleteAt(saldo *pb.SaldoResponseDeleteAt) *response.SaldoResponseDeleteAt {
	var deletedAt string
	if saldo.DeletedAt != nil {
		deletedAt = saldo.DeletedAt.Value
	}

	return &response.SaldoResponseDeleteAt{
		ID:             int(saldo.SaldoId),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime,
		WithdrawAmount: int(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
		DeletedAt:      &deletedAt,
	}
}

// mapResponsesSaldoDeleteAt maps a slice of gRPC SaldoResponseDeleteAt to a slice of HTTP API SaldoResponseDeleteAt.
// It constructs a slice of SaldoResponseDeleteAt by converting and copying relevant fields
// such as ID, CardNumber, TotalBalance, WithdrawTime, WithdrawAmount,
// CreatedAt, UpdatedAt, and DeletedAt from each gRPC SaldoResponseDeleteAt.
//
// Args:
//
//	saldos: A slice of pointers to pb.SaldoResponseDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.SaldoResponseDeleteAt with mapped data.
func (s *saldoQueryResponseMapper) mapResponsesSaldoDeleteAt(saldos []*pb.SaldoResponseDeleteAt) []*response.SaldoResponseDeleteAt {
	var responseSaldos []*response.SaldoResponseDeleteAt

	for _, saldo := range saldos {
		responseSaldos = append(responseSaldos, s.mapResponseSaldoDeleteAt(saldo))
	}

	return responseSaldos
}
