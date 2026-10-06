package saldoapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type saldoCommandResponseMapper struct {
}

func NewSaldoCommandResponseMapper() SaldoCommandResponseMapper {
	return &saldoCommandResponseMapper{}
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
func (s *saldoCommandResponseMapper) ToApiResponseSaldo(pbResponse *pb.ApiResponseSaldo) *response.ApiResponseSaldo {
	return &response.ApiResponseSaldo{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseSaldo(pbResponse.Data),
	}
}

func (s *saldoCommandResponseMapper) ToApiResponseSaldoDeleteAt(pbResponse *pb.ApiResponseSaldoDeleteAt) *response.ApiResponseSaldoDeleteAt {
	return &response.ApiResponseSaldoDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseSaldoDeleteAt(pbResponse.Data),
	}
}

// ToApiResponseSaldoDelete maps a gRPC response containing a deleted saldo record to an HTTP API response.
// It constructs an ApiResponseSaldoDelete by copying the status and message fields from the gRPC response.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponseSaldoDelete containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseSaldoDelete with mapped data.
func (s *saldoCommandResponseMapper) ToApiResponseSaldoDelete(pbResponse *pb.ApiResponseSaldoDelete) *response.ApiResponseSaldoDelete {
	return &response.ApiResponseSaldoDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponseSaldoAll maps a gRPC response containing all saldo records to an HTTP API response. It constructs an ApiResponseSaldoAll by copying the status and message fields from the gRPC response.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponseSaldoAll containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseSaldoAll with mapped data.
func (s *saldoCommandResponseMapper) ToApiResponseSaldoAll(pbResponse *pb.ApiResponseSaldoAll) *response.ApiResponseSaldoAll {
	return &response.ApiResponseSaldoAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
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
func (s *saldoCommandResponseMapper) mapResponseSaldo(saldo *pb.SaldoResponse) *response.SaldoResponse {
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
func (s *saldoCommandResponseMapper) mapResponseSaldoDeleteAt(saldo *pb.SaldoResponseDeleteAt) *response.SaldoResponseDeleteAt {
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
