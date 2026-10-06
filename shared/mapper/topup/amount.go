package topupapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type topupStatsAmountResponseMapper struct {
}

func NewTopupStatsAmountResponseMapper() TopupStatsAmountResponseMapper {
	return &topupStatsAmountResponseMapper{}
}

// ToApiResponseTopupMonthAmount maps a gRPC response containing a month's worth of top-up amounts
// to an HTTP API response. It constructs an ApiResponseTopupMonthAmount by copying the status and message fields,
// mapping the TopupMonthAmount data slice to a slice of TopupMonthAmountResponse, and assigning it to the response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupMonthAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupMonthAmount with mapped data.
func (t *topupStatsAmountResponseMapper) ToApiResponseTopupMonthAmount(s *pb.ApiResponseTopupMonthAmount) *response.ApiResponseTopupMonthAmount {
	return &response.ApiResponseTopupMonthAmount{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponseTopupMonthlyAmounts(s.Data),
	}
}

// ToApiResponseTopupYearAmount maps a gRPC response containing a year's worth of top-up amounts
// to an HTTP API response. It constructs an ApiResponseTopupYearAmount by copying the status and message fields,
// mapping the TopupYearAmount data slice to a slice of TopupYearAmountResponse, and assigning it to the response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupYearAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupYearAmount with mapped data.
func (t *topupStatsAmountResponseMapper) ToApiResponseTopupYearAmount(s *pb.ApiResponseTopupYearAmount) *response.ApiResponseTopupYearAmount {
	return &response.ApiResponseTopupYearAmount{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponseTopupYearlyAmounts(s.Data),
	}
}

// mapResponseTopupMonthlyAmount maps a gRPC response containing a month's worth of top-up
// amount statistics to an HTTP API response format. It constructs a TopupMonthAmountResponse
// by copying the month and total amount fields from the gRPC response to the API response.
func (t *topupStatsAmountResponseMapper) mapResponseTopupMonthlyAmount(s *pb.TopupMonthAmountResponse) *response.TopupMonthAmountResponse {
	return &response.TopupMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTopupMonthlyAmounts maps a slice of gRPC responses containing monthly top-up amount statistics
// to a slice of HTTP API response formats. It iterates over the gRPC response slice, mapping each individual response
// using mapResponseTopupMonthlyAmount, and returns a slice of mapped TopupMonthAmountResponse.
//
// Args:
//   - s: A slice of pointers to pb.TopupMonthAmountResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupMonthAmountResponse with the mapped data.
func (t *topupStatsAmountResponseMapper) mapResponseTopupMonthlyAmounts(s []*pb.TopupMonthAmountResponse) []*response.TopupMonthAmountResponse {
	var topupProtos []*response.TopupMonthAmountResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupMonthlyAmount(topup))
	}
	return topupProtos
}

// mapResponseTopupYearlyAmount maps a gRPC response containing a year's worth of top-up amount statistics
// to an HTTP API response format. It constructs a TopupYearlyAmountResponse by copying the year and
// total amount fields from the gRPC response to the API response.
func (t *topupStatsAmountResponseMapper) mapResponseTopupYearlyAmount(s *pb.TopupYearlyAmountResponse) *response.TopupYearlyAmountResponse {
	return &response.TopupYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTopupYearlyAmounts maps a slice of gRPC responses containing yearly top-up amount statistics
// to a slice of HTTP API response formats. It iterates over the gRPC response slice, mapping each individual response
// using mapResponseTopupYearlyAmount, and returns a slice of mapped TopupYearlyAmountResponse.
//
// Args:
//   - s: A slice of pointers to pb.TopupYearlyAmountResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupYearlyAmountResponse with the mapped data.
func (t *topupStatsAmountResponseMapper) mapResponseTopupYearlyAmounts(s []*pb.TopupYearlyAmountResponse) []*response.TopupYearlyAmountResponse {
	var topupProtos []*response.TopupYearlyAmountResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupYearlyAmount(topup))
	}
	return topupProtos
}
