package saldoapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type saldoStatsBalanceResponseMapper struct{}

func NewSaldoStatsBalanceResponseMapper() SaldoStatsBalanceResponseMapper {
	return &saldoStatsBalanceResponseMapper{}
}

// ToApiResponseMonthSaldoBalances maps a gRPC response containing monthly saldo
// balance statistics to an HTTP API response. It constructs an
// ApiResponseMonthSaldoBalances by copying the status and message fields and mapping
// the saldo statistics to a SaldoMonthBalanceResponse slice.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponseMonthSaldoBalances containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMonthSaldoBalances with mapped data.
func (s *saldoStatsBalanceResponseMapper) ToApiResponseMonthSaldoBalances(pbResponse *pb.ApiResponseMonthSaldoBalances) *response.ApiResponseMonthSaldoBalances {
	return &response.ApiResponseMonthSaldoBalances{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapSaldoMonthBalanceResponses(pbResponse.Data),
	}
}

// ToApiResponseYearSaldoBalances maps a gRPC response containing yearly saldo
// balance statistics to an HTTP API response. It constructs an
// ApiResponseYearSaldoBalances by copying the status and message fields and
// mapping the saldo statistics to a SaldoYearBalanceResponse slice.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponseYearSaldoBalances containing the
//	gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseYearSaldoBalances with mapped data.
func (s *saldoStatsBalanceResponseMapper) ToApiResponseYearSaldoBalances(pbResponse *pb.ApiResponseYearSaldoBalances) *response.ApiResponseYearSaldoBalances {
	return &response.ApiResponseYearSaldoBalances{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapSaldoYearBalanceResponses(pbResponse.Data),
	}
}

// mapSaldoMonthBalanceResponse maps a gRPC SaldoMonthBalanceResponse to an HTTP API SaldoMonthBalanceResponse.
// It constructs a SaldoMonthBalanceResponse by copying the Month field and mapping the TotalBalance field to an integer.
//
// Args:
//
//	ss: A pointer to a pb.SaldoMonthBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.SaldoMonthBalanceResponse with mapped data.
func (s *saldoStatsBalanceResponseMapper) mapSaldoMonthBalanceResponse(ss *pb.SaldoMonthBalanceResponse) *response.SaldoMonthBalanceResponse {
	return &response.SaldoMonthBalanceResponse{
		Month:        ss.Month,
		TotalBalance: int(ss.TotalBalance),
	}
}

// mapSaldoMonthBalanceResponses maps a slice of gRPC SaldoMonthBalanceResponses to a slice of HTTP API SaldoMonthBalanceResponses.
// It constructs a slice of SaldoMonthBalanceResponse by copying the Month field and mapping the TotalBalance field to an integer from each gRPC SaldoMonthBalanceResponse.
//
// Args:
//
//	ss: A slice of pointers to pb.SaldoMonthBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.SaldoMonthBalanceResponse with mapped data.
func (s *saldoStatsBalanceResponseMapper) mapSaldoMonthBalanceResponses(ss []*pb.SaldoMonthBalanceResponse) []*response.SaldoMonthBalanceResponse {
	var saldoProtos []*response.SaldoMonthBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoMonthBalanceResponse(saldo))
	}
	return saldoProtos
}

// mapSaldoYearBalanceResponse maps a gRPC SaldoYearBalanceResponse to an HTTP API SaldoYearBalanceResponse.
// It constructs a SaldoYearBalanceResponse by copying the Year field and mapping the TotalBalance field to an integer.
//
// Args:
//
//	ss: A pointer to a pb.SaldoYearBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.SaldoYearBalanceResponse with mapped data.
func (s *saldoStatsBalanceResponseMapper) mapSaldoYearBalanceResponse(ss *pb.SaldoYearBalanceResponse) *response.SaldoYearBalanceResponse {
	return &response.SaldoYearBalanceResponse{
		Year:         ss.Year,
		TotalBalance: int(ss.TotalBalance),
	}
}

// mapSaldoYearBalanceResponses maps a slice of gRPC SaldoYearBalanceResponses to a slice of HTTP API SaldoYearBalanceResponses.
// It constructs a slice of SaldoYearBalanceResponse by copying the Year field and mapping the TotalBalance field to an integer from each gRPC SaldoYearBalanceResponse.
//
// Args:
//
//	ss: A slice of pointers to pb.SaldoYearBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.SaldoYearBalanceResponse with mapped data.
func (s *saldoStatsBalanceResponseMapper) mapSaldoYearBalanceResponses(ss []*pb.SaldoYearBalanceResponse) []*response.SaldoYearBalanceResponse {
	var saldoProtos []*response.SaldoYearBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoYearBalanceResponse(saldo))
	}
	return saldoProtos
}
