package transferapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transfer/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transferStatsStatusResponseMapper struct{}

func NewTransferStatsStatusResponseMapper() TransferStatsStatusResponseMapper {
	return &transferStatsStatusResponseMapper{}
}

// ToApiResponseTransferMonthStatusSuccess converts a gRPC response containing a month's worth
// of successful transfer statistics into an HTTP API response. It constructs an
// ApiResponseTransferMonthStatusSuccess by copying the status and message fields, mapping the
// TransferMonthStatusSuccess data slice to a slice of TransferResponseMonthStatusSuccess, and
// assigning it to the response's Data field.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferMonthStatusSuccess containing the gRPC
//     response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferMonthStatusSuccess with mapped data.
func (m *transferStatsStatusResponseMapper) ToApiResponseTransferMonthStatusSuccess(pbResponse *pb.ApiResponseTransferMonthStatusSuccess) *response.ApiResponseTransferMonthStatusSuccess {
	return &response.ApiResponseTransferMonthStatusSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesTransferMonthStatusSuccess(pbResponse.Data),
	}
}

// ToApiResponseTransferYearStatusSuccess maps a gRPC response containing a year's worth of successful transfer
// statistics to an HTTP API response. It constructs an ApiResponseTransferYearStatusSuccess by copying the
// status and message fields, mapping the TransferYearStatusSuccess data slice to a slice of
// TransferResponseYearStatusSuccess, and assigning it to the response's Data field.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferYearStatusSuccess containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferYearStatusSuccess with mapped data.
func (m *transferStatsStatusResponseMapper) ToApiResponseTransferYearStatusSuccess(pbResponse *pb.ApiResponseTransferYearStatusSuccess) *response.ApiResponseTransferYearStatusSuccess {
	return &response.ApiResponseTransferYearStatusSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapTransferResponsesYearStatusSuccess(pbResponse.Data),
	}
}

// ToApiResponseTransferMonthStatusFailed maps a gRPC response containing a month's worth of failed transfer
// statistics into an HTTP API response. It constructs an ApiResponseTransferMonthStatusFailed by copying the
// status and message fields, mapping the TransferMonthStatusFailed data slice to a slice of
// TransferResponseMonthStatusFailed, and assigning it to the response's Data field.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferMonthStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferMonthStatusFailed with mapped data.
func (m *transferStatsStatusResponseMapper) ToApiResponseTransferMonthStatusFailed(pbResponse *pb.ApiResponseTransferMonthStatusFailed) *response.ApiResponseTransferMonthStatusFailed {
	return &response.ApiResponseTransferMonthStatusFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesTransferMonthStatusFailed(pbResponse.Data),
	}
}

// ToApiResponseTransferYearStatusFailed maps a gRPC response containing a year's worth of failed transfer
// statistics into an HTTP API response. It constructs an ApiResponseTransferYearStatusFailed by copying the
// status and message fields, mapping the TransferYearStatusFailed data slice to a slice of
// TransferResponseYearStatusFailed, and assigning it to the response's Data field.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferYearStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferYearStatusFailed with mapped data.
func (m *transferStatsStatusResponseMapper) ToApiResponseTransferYearStatusFailed(pbResponse *pb.ApiResponseTransferYearStatusFailed) *response.ApiResponseTransferYearStatusFailed {
	return &response.ApiResponseTransferYearStatusFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapTransferResponsesYearStatusFailed(pbResponse.Data),
	}
}

// mapResponseTransferMonthStatusSuccess maps a TransferMonthStatusSuccessResponse protobuf message to a
// response.TransferResponseMonthStatusSuccess.
//
// Args:
//   - s: A pointer to a pb.TransferMonthStatusSuccessResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponseMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transferStatsStatusResponseMapper) mapResponseTransferMonthStatusSuccess(s *pb.TransferMonthStatusSuccessResponse) *response.TransferResponseMonthStatusSuccess {
	return &response.TransferResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// mapResponsesTransferMonthStatusSuccess maps a slice of TransferMonthStatusSuccessResponse
// protobuf messages to a slice of TransferResponseMonthStatusSuccess.
//
// It iterates over the input slice, converting each TransferMonthStatusSuccessResponse
// to its corresponding response representation using the mapResponseTransferMonthStatusSuccess method.
//
// Args:
//   - Transfers: A slice of pointers to pb.TransferMonthStatusSuccessResponse containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferResponseMonthStatusSuccess containing the mapped data.
func (t *transferStatsStatusResponseMapper) mapResponsesTransferMonthStatusSuccess(Transfers []*pb.TransferMonthStatusSuccessResponse) []*response.TransferResponseMonthStatusSuccess {
	var TransferRecords []*response.TransferResponseMonthStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapResponseTransferMonthStatusSuccess(Transfer))
	}

	return TransferRecords
}

// mapTransferResponseYearStatusSuccess maps a TransferYearStatusSuccessResponse protobuf message to a
// response.TransferResponseYearStatusSuccess.
//
// Args:
//   - s: A pointer to a pb.TransferYearStatusSuccessResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponseYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transferStatsStatusResponseMapper) mapTransferResponseYearStatusSuccess(s *pb.TransferYearStatusSuccessResponse) *response.TransferResponseYearStatusSuccess {
	return &response.TransferResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// mapTransferResponsesYearStatusSuccess maps a slice of TransferYearStatusSuccessResponse
// protobuf messages to a slice of TransferResponseYearStatusSuccess.
//
// It iterates over the input slice, converting each TransferYearStatusSuccessResponse
// to its corresponding response representation using the mapTransferResponseYearStatusSuccess method.
//
// Args:
//   - Transfers: A slice of pointers to pb.TransferYearStatusSuccessResponse containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferResponseYearStatusSuccess containing the mapped data.
func (t *transferStatsStatusResponseMapper) mapTransferResponsesYearStatusSuccess(Transfers []*pb.TransferYearStatusSuccessResponse) []*response.TransferResponseYearStatusSuccess {
	var TransferRecords []*response.TransferResponseYearStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapTransferResponseYearStatusSuccess(Transfer))
	}

	return TransferRecords
}

// mapResponseTransferMonthStatusFailed maps a TransferMonthStatusFailedResponse protobuf message to a
// response.TransferResponseMonthStatusFailed.
//
// Args:
//   - s: A pointer to a pb.TransferMonthStatusFailedResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponseMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transferStatsStatusResponseMapper) mapResponseTransferMonthStatusFailed(s *pb.TransferMonthStatusFailedResponse) *response.TransferResponseMonthStatusFailed {
	return &response.TransferResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponsesTransferMonthStatusFailed maps a slice of TransferMonthStatusFailedResponse
// protobuf messages to a slice of TransferResponseMonthStatusFailed.
//
// It iterates over the input slice, converting each TransferMonthStatusFailedResponse
// to its corresponding response representation using the mapResponseTransferMonthStatusFailed method.
//
// Args:
//   - Transfers: A slice of pointers to pb.TransferMonthStatusFailedResponse containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferResponseMonthStatusFailed containing the mapped data.
func (t *transferStatsStatusResponseMapper) mapResponsesTransferMonthStatusFailed(Transfers []*pb.TransferMonthStatusFailedResponse) []*response.TransferResponseMonthStatusFailed {
	var TransferRecords []*response.TransferResponseMonthStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapResponseTransferMonthStatusFailed(Transfer))
	}

	return TransferRecords
}

// mapTransferResponseYearStatusFailed maps a TransferYearStatusFailedResponse protobuf message to a
// response.TransferResponseYearStatusFailed.
//
// Args:
//   - s: A pointer to a pb.TransferYearStatusFailedResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponseYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transferStatsStatusResponseMapper) mapTransferResponseYearStatusFailed(s *pb.TransferYearStatusFailedResponse) *response.TransferResponseYearStatusFailed {
	return &response.TransferResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapTransferResponsesYearStatusFailed maps a slice of TransferYearStatusFailedResponse
// protobuf messages to a slice of TransferResponseYearStatusFailed.
//
// It iterates over the input slice, converting each TransferYearStatusFailedResponse
// to its corresponding response representation using the mapTransferResponseYearStatusFailed method.
//
// Args:
//   - Transfers: A slice of pointers to pb.TransferYearStatusFailedResponse containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferResponseYearStatusFailed containing the mapped data.
func (t *transferStatsStatusResponseMapper) mapTransferResponsesYearStatusFailed(Transfers []*pb.TransferYearStatusFailedResponse) []*response.TransferResponseYearStatusFailed {
	var TransferRecords []*response.TransferResponseYearStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapTransferResponseYearStatusFailed(Transfer))
	}

	return TransferRecords
}
