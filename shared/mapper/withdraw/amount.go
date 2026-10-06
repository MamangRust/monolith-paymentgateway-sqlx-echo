package withdrawapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type withdrawStatsAmountResponseMapper struct {
}

func NewWithdrawStatsAmountResponseMapper() WithdrawStatsAmountResponseMapper {
	return &withdrawStatsAmountResponseMapper{}
}

// ToApiResponseWithdrawMonthAmount converts a gRPC response containing monthly withdraw amounts
// into an API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseWithdrawMonthAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseWithdrawMonthAmount containing the mapped data, including status, message,
//     and detailed information about the monthly withdraw amounts.
func (m *withdrawStatsAmountResponseMapper) ToApiResponseWithdrawMonthAmount(pbResponse *pb.ApiResponseWithdrawMonthAmount) *response.ApiResponseWithdrawMonthAmount {
	return &response.ApiResponseWithdrawMonthAmount{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseWithdrawMonthlyAmounts(pbResponse.Data),
	}
}

// ToApiResponseWithdrawYearAmount maps a gRPC response containing yearly withdraw amounts
// into an API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseWithdrawYearAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseWithdrawYearAmount containing the mapped data, including status, message,
//     and detailed information about the yearly withdraw amounts.
func (m *withdrawStatsAmountResponseMapper) ToApiResponseWithdrawYearAmount(pbResponse *pb.ApiResponseWithdrawYearAmount) *response.ApiResponseWithdrawYearAmount {
	return &response.ApiResponseWithdrawYearAmount{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseWithdrawYearlyAmounts(pbResponse.Data),
	}
}

// mapResponseWithdrawMonthlyAmount maps a single WithdrawMonthlyAmountResponse gRPC message
// to an API response WithdrawMonthlyAmountResponse object.
//
// Args:
//   - s: The WithdrawMonthlyAmountResponse gRPC message that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawMonthlyAmountResponse containing the mapped data, including
//     fields like Month and TotalAmount.
func (m *withdrawStatsAmountResponseMapper) mapResponseWithdrawMonthlyAmount(s *pb.WithdrawMonthlyAmountResponse) *response.WithdrawMonthlyAmountResponse {
	return &response.WithdrawMonthlyAmountResponse{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseWithdrawMonthlyAmounts maps a slice of gRPC WithdrawMonthlyAmountResponse messages
// to a slice of API response WithdrawMonthlyAmountResponse objects.
//
// Args:
//   - s: A slice of pointers to pb.WithdrawMonthlyAmountResponse messages that need to be converted.
//
// Returns:
//   - A slice of pointers to response.WithdrawMonthlyAmountResponse objects containing the mapped data,
//     including fields like Month and TotalAmount.
func (m *withdrawStatsAmountResponseMapper) mapResponseWithdrawMonthlyAmounts(s []*pb.WithdrawMonthlyAmountResponse) []*response.WithdrawMonthlyAmountResponse {
	var protoResponses []*response.WithdrawMonthlyAmountResponse
	for _, withdraw := range s {
		protoResponses = append(protoResponses, m.mapResponseWithdrawMonthlyAmount(withdraw))
	}
	return protoResponses
}

// mapResponseWithdrawYearlyAmount maps a single WithdrawYearlyAmountResponse gRPC message
// to an API response WithdrawYearlyAmountResponse object.
//
// Args:
//   - s: The WithdrawYearlyAmountResponse gRPC message that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawYearlyAmountResponse containing the mapped data, including
//     fields like Year and TotalAmount.
func (m *withdrawStatsAmountResponseMapper) mapResponseWithdrawYearlyAmount(s *pb.WithdrawYearlyAmountResponse) *response.WithdrawYearlyAmountResponse {
	return &response.WithdrawYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseWithdrawYearlyAmounts maps a slice of gRPC WithdrawYearlyAmountResponse messages
// to a slice of API response WithdrawYearlyAmountResponse objects.
//
// Args:
//   - s: A slice of pointers to pb.WithdrawYearlyAmountResponse messages that need to be converted.
//
// Returns:
//   - A slice of pointers to response.WithdrawYearlyAmountResponse objects containing the mapped data,
//     including fields like Year and TotalAmount.
func (m *withdrawStatsAmountResponseMapper) mapResponseWithdrawYearlyAmounts(s []*pb.WithdrawYearlyAmountResponse) []*response.WithdrawYearlyAmountResponse {
	var protoResponses []*response.WithdrawYearlyAmountResponse
	for _, withdraw := range s {
		protoResponses = append(protoResponses, m.mapResponseWithdrawYearlyAmount(withdraw))
	}
	return protoResponses
}
