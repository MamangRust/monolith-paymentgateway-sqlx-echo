package topupapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type topupStatsStatusResponseMapper struct{}

func NewTopupStatsStatusResponseMapper() TopupStatsStatusResponseMapper {
	return &topupStatsStatusResponseMapper{}
}

// ToApiResponseTopupMonthStatusSuccess maps a gRPC response containing a month's worth of successful top-up statistics
// to an HTTP API response. It constructs an ApiResponseTopupMonthStatusSuccess by copying the status and message fields,
// mapping the TopupMonthStatusSuccess data slice to a slice of TopupMonthStatusSuccessResponse, and assigning it to the
// response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupMonthStatusSuccess containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupMonthStatusSuccess with mapped data.
func (t *topupStatsStatusResponseMapper) ToApiResponseTopupMonthStatusSuccess(s *pb.ApiResponseTopupMonthStatusSuccess) *response.ApiResponseTopupMonthStatusSuccess {
	return &response.ApiResponseTopupMonthStatusSuccess{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponsesTopupMonthStatusSuccess(s.Data),
	}
}

// ToApiResponseTopupYearStatusSuccess maps a gRPC response containing a year's worth of successful top-up statistics
// to an HTTP API response. It constructs an ApiResponseTopupYearStatusSuccess by copying the status and message fields,
// mapping the TopupYearStatusSuccess data slice to a slice of TopupYearStatusSuccessResponse, and assigning it to the
// response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupYearStatusSuccess containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupYearStatusSuccess with mapped data.
func (t *topupStatsStatusResponseMapper) ToApiResponseTopupYearStatusSuccess(s *pb.ApiResponseTopupYearStatusSuccess) *response.ApiResponseTopupYearStatusSuccess {
	return &response.ApiResponseTopupYearStatusSuccess{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapTopupResponsesYearStatusSuccess(s.Data),
	}
}

// ToApiResponseTopupMonthStatusFailed maps a gRPC response containing a month's worth of failed top-up statistics
// to an HTTP API response. It constructs an ApiResponseTopupMonthStatusFailed by copying the status and message fields,
// mapping the TopupMonthStatusFailed data slice to a slice of TopupMonthStatusFailedResponse, and assigning it to the
// response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupMonthStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupMonthStatusFailed with mapped data.
func (t *topupStatsStatusResponseMapper) ToApiResponseTopupMonthStatusFailed(s *pb.ApiResponseTopupMonthStatusFailed) *response.ApiResponseTopupMonthStatusFailed {
	return &response.ApiResponseTopupMonthStatusFailed{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponsesTopupMonthStatusFailed(s.Data),
	}
}

// ToApiResponseTopupYearStatusFailed maps a gRPC response containing a year's worth of failed top-up statistics
// to an HTTP API response. It constructs an ApiResponseTopupYearStatusFailed by copying the status and message fields,
// mapping the TopupYearStatusFailed data slice to a slice of TopupYearStatusFailedResponse, and assigning it to the
// response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupYearStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupYearStatusFailed with mapped data.
func (t *topupStatsStatusResponseMapper) ToApiResponseTopupYearStatusFailed(s *pb.ApiResponseTopupYearStatusFailed) *response.ApiResponseTopupYearStatusFailed {
	return &response.ApiResponseTopupYearStatusFailed{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapTopupResponsesYearStatusFailed(s.Data),
	}
}

// mapResponseTopupMonthStatusSuccess maps a gRPC response containing a month's worth of successful top-up statistics
// to an HTTP API response format. It constructs a TopupResponseMonthStatusSuccess by copying the year, month,
// total success count, and total amount fields from the gRPC response to the API response.
func (t *topupStatsStatusResponseMapper) mapResponseTopupMonthStatusSuccess(s *pb.TopupMonthStatusSuccessResponse) *response.TopupResponseMonthStatusSuccess {
	return &response.TopupResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// mapResponsesTopupMonthStatusSuccess maps a slice of gRPC responses containing a month's worth of successful
// top-up statistics to a slice of HTTP API responses. It iterates over the gRPC response slice, mapping
// each individual response using mapResponseTopupMonthStatusSuccess, and returns a slice of mapped
// TopupResponseMonthStatusSuccess.
//
// Args:
//   - topups: A slice of pointers to pb.TopupMonthStatusSuccessResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupResponseMonthStatusSuccess with the mapped data.
func (t *topupStatsStatusResponseMapper) mapResponsesTopupMonthStatusSuccess(topups []*pb.TopupMonthStatusSuccessResponse) []*response.TopupResponseMonthStatusSuccess {
	var topupRecords []*response.TopupResponseMonthStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapResponseTopupMonthStatusSuccess(topup))
	}

	return topupRecords
}

// mapTopupResponseYearStatusSuccess maps a gRPC response containing a year's worth of successful top-up statistics
// to an HTTP API response format. It constructs a TopupResponseYearStatusSuccess by copying the year, total success count,
// and total amount fields from the gRPC response to the API response.
func (t *topupStatsStatusResponseMapper) mapTopupResponseYearStatusSuccess(s *pb.TopupYearStatusSuccessResponse) *response.TopupResponseYearStatusSuccess {
	return &response.TopupResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// mapTopupResponsesYearStatusSuccess maps a slice of gRPC responses containing a year's worth of successful
// top-up statistics to a slice of HTTP API responses. It iterates over the gRPC response slice, mapping
// each individual response using mapTopupResponseYearStatusSuccess, and returns a slice of mapped
// TopupResponseYearStatusSuccess.
//
// Args:
//   - topups: A slice of pointers to pb.TopupYearStatusSuccessResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupResponseYearStatusSuccess with the mapped data.
func (t *topupStatsStatusResponseMapper) mapTopupResponsesYearStatusSuccess(topups []*pb.TopupYearStatusSuccessResponse) []*response.TopupResponseYearStatusSuccess {
	var topupRecords []*response.TopupResponseYearStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapTopupResponseYearStatusSuccess(topup))
	}

	return topupRecords
}

// mapResponseTopupMonthStatusFailed maps a gRPC response containing a month's worth of failed top-up statistics
// to an HTTP API response format. It constructs a TopupResponseMonthStatusFailed by copying the year, month,
// total failed count, and total amount fields from the gRPC response to the API response.
func (t *topupStatsStatusResponseMapper) mapResponseTopupMonthStatusFailed(s *pb.TopupMonthStatusFailedResponse) *response.TopupResponseMonthStatusFailed {
	return &response.TopupResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponsesTopupMonthStatusFailed maps a slice of gRPC responses containing a month's worth of failed
// top-up statistics to a slice of HTTP API responses. It iterates over the gRPC response slice, mapping
// each individual response using mapResponseTopupMonthStatusFailed, and returns a slice of mapped
// TopupResponseMonthStatusFailed.
//
// Args:
//   - topups: A slice of pointers to pb.TopupMonthStatusFailedResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupResponseMonthStatusFailed with the mapped data.
func (t *topupStatsStatusResponseMapper) mapResponsesTopupMonthStatusFailed(topups []*pb.TopupMonthStatusFailedResponse) []*response.TopupResponseMonthStatusFailed {
	var topupRecords []*response.TopupResponseMonthStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapResponseTopupMonthStatusFailed(topup))
	}

	return topupRecords
}

// mapTopupResponseYearStatusFailed maps a gRPC response containing a year's worth
// of failed top-up statistics to an HTTP API response format. It constructs a
// TopupResponseYearStatusFailed by copying the year, total failed count, and
// total amount fields from the gRPC response to the API response.
func (t *topupStatsStatusResponseMapper) mapTopupResponseYearStatusFailed(s *pb.TopupYearStatusFailedResponse) *response.TopupResponseYearStatusFailed {
	return &response.TopupResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapTopupResponsesYearStatusFailed maps a slice of gRPC responses containing a year's worth
// of failed top-up statistics to a slice of HTTP API responses. It iterates over the gRPC
// response slice, mapping each individual response using mapTopupResponseYearStatusFailed,
// and returns a slice of mapped TopupResponseYearStatusFailed.
//
// Args:
//   - topups: A slice of pointers to pb.TopupYearStatusFailedResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupResponseYearStatusFailed with the mapped data.
func (t *topupStatsStatusResponseMapper) mapTopupResponsesYearStatusFailed(topups []*pb.TopupYearStatusFailedResponse) []*response.TopupResponseYearStatusFailed {
	var topupRecords []*response.TopupResponseYearStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapTopupResponseYearStatusFailed(topup))
	}

	return topupRecords
}
