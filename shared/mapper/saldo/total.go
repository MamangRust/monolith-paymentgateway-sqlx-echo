package saldoapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type saldoStatsTotalResponseMapper struct{}

func NewSaldoStatsTotalResponseMapper() SaldoStatsTotalResponseMapper {
	return &saldoStatsTotalResponseMapper{}
}

// ToApiResponseMonthTotalSaldo maps a gRPC response containing monthly total saldo
// statistics to an HTTP API response. It constructs an ApiResponseMonthTotalSaldo by
// copying the status and message fields and mapping the saldo statistics to a
// SaldoMonthTotalBalanceResponse slice.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponseMonthTotalSaldo containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMonthTotalSaldo with mapped data.
func (s *saldoStatsTotalResponseMapper) ToApiResponseMonthTotalSaldo(pbResponse *pb.ApiResponseMonthTotalSaldo) *response.ApiResponseMonthTotalSaldo {
	return &response.ApiResponseMonthTotalSaldo{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapSaldoMonthTotalBalanceResponses(pbResponse.Data),
	}
}

// ToApiResponseYearTotalSaldo maps a gRPC response containing yearly total saldo
// statistics to an HTTP API response. It constructs an ApiResponseYearTotalSaldo by
// copying the status and message fields and mapping the saldo statistics to a
// SaldoYearTotalBalanceResponse slice.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponseYearTotalSaldo containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseYearTotalSaldo with mapped data.
func (s *saldoStatsTotalResponseMapper) ToApiResponseYearTotalSaldo(pbResponse *pb.ApiResponseYearTotalSaldo) *response.ApiResponseYearTotalSaldo {
	return &response.ApiResponseYearTotalSaldo{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapSaldoYearTotalBalanceResponses(pbResponse.Data),
	}
}

// mapSaldoMonthTotalBalanceResponse maps a gRPC SaldoMonthTotalBalanceResponse to an HTTP API SaldoMonthTotalBalanceResponse.
// It constructs a SaldoMonthTotalBalanceResponse by copying the Month and Year fields and mapping the TotalBalance field to an integer.
//
// Args:
//
//	ss: A pointer to a pb.SaldoMonthTotalBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.SaldoMonthTotalBalanceResponse with mapped data.
func (s *saldoStatsTotalResponseMapper) mapSaldoMonthTotalBalanceResponse(ss *pb.SaldoMonthTotalBalanceResponse) *response.SaldoMonthTotalBalanceResponse {
	totalBalance := 0

	if ss.TotalBalance != 0 {
		totalBalance = int(ss.TotalBalance)
	}

	return &response.SaldoMonthTotalBalanceResponse{
		Month:        ss.Month,
		Year:         ss.Year,
		TotalBalance: totalBalance,
	}
}

// mapSaldoMonthTotalBalanceResponses maps a slice of gRPC SaldoMonthTotalBalanceResponses to a slice of HTTP API SaldoMonthTotalBalanceResponses.
// It constructs a slice of SaldoMonthTotalBalanceResponse by copying the Month and Year fields and mapping the TotalBalance field to an integer from each gRPC SaldoMonthTotalBalanceResponse.
//
// Args:
//
//	ss: A slice of pointers to pb.SaldoMonthTotalBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.SaldoMonthTotalBalanceResponse with mapped data.
func (s *saldoStatsTotalResponseMapper) mapSaldoMonthTotalBalanceResponses(ss []*pb.SaldoMonthTotalBalanceResponse) []*response.SaldoMonthTotalBalanceResponse {
	var saldoProtos []*response.SaldoMonthTotalBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoMonthTotalBalanceResponse(saldo))
	}
	return saldoProtos
}

// mapSaldoYearTotalBalanceResponse maps a gRPC SaldoYearTotalBalanceResponse to an HTTP API SaldoYearTotalBalanceResponse.
// It constructs a SaldoYearTotalBalanceResponse by converting and copying the Year
// field and mapping the TotalBalance field to an integer.
//
// Args:
//
//	ss: A pointer to a pb.SaldoYearTotalBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.SaldoYearTotalBalanceResponse with mapped data.
func (s *saldoStatsTotalResponseMapper) mapSaldoYearTotalBalanceResponse(ss *pb.SaldoYearTotalBalanceResponse) *response.SaldoYearTotalBalanceResponse {
	totalBalance := 0

	if ss.TotalBalance != 0 {
		totalBalance = int(ss.TotalBalance)
	}

	return &response.SaldoYearTotalBalanceResponse{
		Year:         ss.Year,
		TotalBalance: totalBalance,
	}
}

// mapSaldoYearTotalBalanceResponses maps a slice of gRPC SaldoYearTotalBalanceResponses to a slice of HTTP API SaldoYearTotalBalanceResponses.
// It constructs a slice of SaldoYearTotalBalanceResponse by copying the Year field and mapping the TotalBalance field to an integer from each gRPC SaldoYearTotalBalanceResponse.
//
// Args:
//
//	ss: A slice of pointers to pb.SaldoYearTotalBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.SaldoYearTotalBalanceResponse with mapped data.
func (s *saldoStatsTotalResponseMapper) mapSaldoYearTotalBalanceResponses(ss []*pb.SaldoYearTotalBalanceResponse) []*response.SaldoYearTotalBalanceResponse {
	var saldoProtos []*response.SaldoYearTotalBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoYearTotalBalanceResponse(saldo))
	}
	return saldoProtos
}
