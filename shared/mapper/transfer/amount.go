package transferapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transfer/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transferStatsAmountResponseMapper struct {
}

func NewTransferStatsAmountResponseMapper() TransferStatsAmountResponseMapper {
	return &transferStatsAmountResponseMapper{}
}

// ToApiResponseTransferMonthAmount maps a gRPC response containing a month's worth of transfer amounts
// into an HTTP API response. It constructs an ApiResponseTransferMonthAmount by copying the status and message fields,
// mapping the TransferMonthAmount data slice to a slice of TransferMonthAmountResponse, and assigning it to the response's Data field.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferMonthAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferMonthAmount with mapped data.
func (m *transferStatsAmountResponseMapper) ToApiResponseTransferMonthAmount(pbResponse *pb.ApiResponseTransferMonthAmount) *response.ApiResponseTransferMonthAmount {
	return &response.ApiResponseTransferMonthAmount{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransferMonthAmounts(pbResponse.Data),
	}
}

// ToApiResponseTransferYearAmount maps a gRPC response containing a year's worth of transfer amounts
// into an HTTP API response. It constructs an ApiResponseTransferYearAmount by copying the status and message fields,
// mapping the TransferYearAmount data slice to a slice of TransferYearAmountResponse, and assigning it to the response's Data field.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferYearAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferYearAmount with mapped data.
func (m *transferStatsAmountResponseMapper) ToApiResponseTransferYearAmount(pbResponse *pb.ApiResponseTransferYearAmount) *response.ApiResponseTransferYearAmount {
	return &response.ApiResponseTransferYearAmount{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransferYearAmounts(pbResponse.Data),
	}
}

// mapResponseTransferMonthAmount maps a TransferMonthAmountResponse protobuf message to a
// response.TransferMonthAmountResponse.
//
// Args:
//   - s: A pointer to a pb.TransferMonthAmountResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferMonthAmountResponse containing the mapped data, including Month,
//     and TotalAmount.
func (m *transferStatsAmountResponseMapper) mapResponseTransferMonthAmount(s *pb.TransferMonthAmountResponse) *response.TransferMonthAmountResponse {
	return &response.TransferMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTransferMonthAmounts maps a slice of TransferMonthAmountResponse
// protobuf messages to a slice of TransferMonthAmountResponse.
//
// It iterates over the input slice, converting each TransferMonthAmountResponse
// to its corresponding response representation using the mapResponseTransferMonthAmount method.
//
// Args:
//   - s: A slice of pointers to pb.TransferMonthAmountResponse containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferMonthAmountResponse containing the mapped data.
func (m *transferStatsAmountResponseMapper) mapResponseTransferMonthAmounts(s []*pb.TransferMonthAmountResponse) []*response.TransferMonthAmountResponse {
	var protoResponses []*response.TransferMonthAmountResponse
	for _, transfer := range s {
		protoResponses = append(protoResponses, m.mapResponseTransferMonthAmount(transfer))
	}
	return protoResponses
}

// mapResponseTransferYearAmount maps a TransferYearAmountResponse protobuf message to a
// response.TransferYearAmountResponse.
//
// Args:
//   - s: A pointer to a pb.TransferYearAmountResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferYearAmountResponse containing the mapped data, including Year
//     and TotalAmount.
func (m *transferStatsAmountResponseMapper) mapResponseTransferYearAmount(s *pb.TransferYearAmountResponse) *response.TransferYearAmountResponse {
	return &response.TransferYearAmountResponse{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTransferYearAmounts maps a slice of TransferYearAmountResponse
// protobuf messages to a slice of TransferYearAmountResponse.
//
// It iterates over the input slice, converting each TransferYearAmountResponse
// to its corresponding response representation using the mapResponseTransferYearAmount method.
//
// Args:
//   - s: A slice of pointers to pb.TransferYearAmountResponse containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferYearAmountResponse containing the mapped data.
func (m *transferStatsAmountResponseMapper) mapResponseTransferYearAmounts(s []*pb.TransferYearAmountResponse) []*response.TransferYearAmountResponse {
	var protoResponses []*response.TransferYearAmountResponse
	for _, transfer := range s {
		protoResponses = append(protoResponses, m.mapResponseTransferYearAmount(transfer))
	}
	return protoResponses
}
