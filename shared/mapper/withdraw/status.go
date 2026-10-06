package withdrawapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type withdrawStatsStatusResponseMapper struct {
}

func NewWithdrawStatsStatusResponseMapper() WithdrawStatsStatusResponseMapper {
	return &withdrawStatsStatusResponseMapper{}
}

// ToApiResponseWithdrawMonthStatusSuccess converts a gRPC response containing monthly successful withdraw statistics
// into an API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseWithdrawMonthStatusSuccess containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseWithdrawMonthStatusSuccess containing the mapped data including status, message,
//     and detailed information about the monthly successful withdraw statistics.
func (m *withdrawStatsStatusResponseMapper) ToApiResponseWithdrawMonthStatusSuccess(pbResponse *pb.ApiResponseWithdrawMonthStatusSuccess) *response.ApiResponseWithdrawMonthStatusSuccess {
	return &response.ApiResponseWithdrawMonthStatusSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesWithdrawMonthStatusSuccess(pbResponse.Data),
	}
}

// ToApiResponseWithdrawYearStatusSuccess maps a gRPC response containing yearly successful withdraw statistics
// to an HTTP API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseWithdrawYearStatusSuccess containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseWithdrawYearStatusSuccess containing the mapped data, including status, message,
//     and detailed information about the yearly successful withdraw statistics.
func (m *withdrawStatsStatusResponseMapper) ToApiResponseWithdrawYearStatusSuccess(pbResponse *pb.ApiResponseWithdrawYearStatusSuccess) *response.ApiResponseWithdrawYearStatusSuccess {
	return &response.ApiResponseWithdrawYearStatusSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapWithdrawResponsesYearStatusSuccess(pbResponse.Data),
	}
}

// ToApiResponseWithdrawMonthStatusFailed maps a gRPC response containing monthly failed withdraw statistics
// into an API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseWithdrawMonthStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseWithdrawMonthStatusFailed containing the mapped data including status, message,
//     and detailed information about the monthly failed withdraw statistics.
func (m *withdrawStatsStatusResponseMapper) ToApiResponseWithdrawMonthStatusFailed(pbResponse *pb.ApiResponseWithdrawMonthStatusFailed) *response.ApiResponseWithdrawMonthStatusFailed {
	return &response.ApiResponseWithdrawMonthStatusFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesWithdrawMonthStatusFailed(pbResponse.Data),
	}
}

// ToApiResponseWithdrawYearStatusFailed maps a gRPC response containing yearly failed withdraw statistics
// into an API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseWithdrawYearStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseWithdrawYearStatusFailed containing the mapped data including status, message,
//     and detailed information about the yearly failed withdraw statistics.
func (m *withdrawStatsStatusResponseMapper) ToApiResponseWithdrawYearStatusFailed(pbResponse *pb.ApiResponseWithdrawYearStatusFailed) *response.ApiResponseWithdrawYearStatusFailed {
	return &response.ApiResponseWithdrawYearStatusFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapWithdrawResponsesYearStatusFailed(pbResponse.Data),
	}
}

func (t *withdrawStatsStatusResponseMapper) mapResponseWithdrawMonthStatusSuccess(s *pb.WithdrawMonthStatusSuccessResponse) *response.WithdrawResponseMonthStatusSuccess {
	return &response.WithdrawResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// mapResponsesWithdrawMonthStatusSuccess maps a slice of WithdrawMonthStatusSuccessResponse gRPC messages to a slice
// of API response WithdrawResponseMonthStatusSuccess objects.
//
// Args:
//   - Withdraws: A slice of WithdrawMonthStatusSuccessResponse gRPC messages that need to be converted.
//
// Returns:
//   - A slice of WithdrawResponseMonthStatusSuccess API responses containing the mapped data.
func (t *withdrawStatsStatusResponseMapper) mapResponsesWithdrawMonthStatusSuccess(Withdraws []*pb.WithdrawMonthStatusSuccessResponse) []*response.WithdrawResponseMonthStatusSuccess {
	var WithdrawRecords []*response.WithdrawResponseMonthStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapResponseWithdrawMonthStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// mapWithdrawResponseYearStatusSuccess maps a single WithdrawYearStatusSuccessResponse gRPC message
// to an API response WithdrawResponseYearStatusSuccess object.
//
// Args:
//   - s: The WithdrawYearStatusSuccessResponse gRPC message that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseYearStatusSuccess containing the mapped data, including
//     fields like Year, TotalSuccess, and TotalAmount.
func (t *withdrawStatsStatusResponseMapper) mapWithdrawResponseYearStatusSuccess(s *pb.WithdrawYearStatusSuccessResponse) *response.WithdrawResponseYearStatusSuccess {
	return &response.WithdrawResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// mapWithdrawResponsesYearStatusSuccess maps a slice of gRPC WithdrawYearStatusSuccessResponse messages
// to a slice of API WithdrawResponseYearStatusSuccess objects.
//
// Args:
//   - Withdraws: A slice of pointers to pb.WithdrawYearStatusSuccessResponse messages that need to be converted.
//
// Returns:
//   - A slice of pointers to response.WithdrawResponseYearStatusSuccess objects containing the mapped data,
//     including fields like Year, TotalSuccess, and TotalAmount.
func (t *withdrawStatsStatusResponseMapper) mapWithdrawResponsesYearStatusSuccess(Withdraws []*pb.WithdrawYearStatusSuccessResponse) []*response.WithdrawResponseYearStatusSuccess {
	var WithdrawRecords []*response.WithdrawResponseYearStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapWithdrawResponseYearStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// mapResponseWithdrawMonthStatusFailed maps a single WithdrawMonthStatusFailedResponse gRPC message
// to an API response WithdrawResponseMonthStatusFailed object.
//
// Args:
//   - s: The WithdrawMonthStatusFailedResponse gRPC message that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseMonthStatusFailed containing the mapped data, including
//     fields like Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawStatsStatusResponseMapper) mapResponseWithdrawMonthStatusFailed(s *pb.WithdrawMonthStatusFailedResponse) *response.WithdrawResponseMonthStatusFailed {
	return &response.WithdrawResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponsesWithdrawMonthStatusFailed maps a slice of WithdrawMonthStatusFailedResponse gRPC messages
// to a slice of API response WithdrawResponseMonthStatusFailed objects.
//
// Args:
//   - Withdraws: A slice of pointers to pb.WithdrawMonthStatusFailedResponse messages that need to be converted.
//
// Returns:
//   - A slice of pointers to response.WithdrawResponseMonthStatusFailed objects containing the mapped data,
//     including fields like Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawStatsStatusResponseMapper) mapResponsesWithdrawMonthStatusFailed(Withdraws []*pb.WithdrawMonthStatusFailedResponse) []*response.WithdrawResponseMonthStatusFailed {
	var WithdrawRecords []*response.WithdrawResponseMonthStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapResponseWithdrawMonthStatusFailed(Withdraw))
	}

	return WithdrawRecords
}

// mapWithdrawResponseYearStatusFailed maps a WithdrawYearStatusFailedResponse gRPC message
// to an API response WithdrawResponseYearStatusFailed object.
//
// Args:
//   - s: The WithdrawYearStatusFailedResponse gRPC message that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseYearStatusFailed containing the mapped data, including
//     fields like Year, TotalFailed, and TotalAmount.
func (t *withdrawStatsStatusResponseMapper) mapWithdrawResponseYearStatusFailed(s *pb.WithdrawYearStatusFailedResponse) *response.WithdrawResponseYearStatusFailed {
	return &response.WithdrawResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapWithdrawResponsesYearStatusFailed maps a slice of gRPC WithdrawYearStatusFailedResponse messages
// to a slice of API WithdrawResponseYearStatusFailed objects.
//
// Args:
//   - Withdraws: A slice of pointers to pb.WithdrawYearStatusFailedResponse messages that need to be converted.
//
// Returns:
//   - A slice of pointers to response.WithdrawResponseYearStatusFailed objects containing the mapped data,
//     including fields like Year, TotalFailed, and TotalAmount.s
func (t *withdrawStatsStatusResponseMapper) mapWithdrawResponsesYearStatusFailed(Withdraws []*pb.WithdrawYearStatusFailedResponse) []*response.WithdrawResponseYearStatusFailed {
	var WithdrawRecords []*response.WithdrawResponseYearStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapWithdrawResponseYearStatusFailed(Withdraw))
	}

	return WithdrawRecords
}
