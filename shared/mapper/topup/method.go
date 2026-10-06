package topupapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type topupStatsMethodResponseMapper struct{}

func NewTopupStatsMethodResponseMapper() TopupStatsMethodResponseMapper {
	return &topupStatsMethodResponseMapper{}
}

// ToApiResponseTopupMonthMethod maps a gRPC response containing a month's worth of top-up statistics by payment method
// to an HTTP API response. It constructs an ApiResponseTopupMonthMethod by copying the status and message fields,
// mapping the TopupMonthMethod data slice to a slice of TopupMonthMethodResponse, and assigning it to the response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupMonthMethod containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupMonthMethod with mapped data.
func (t *topupStatsMethodResponseMapper) ToApiResponseTopupMonthMethod(s *pb.ApiResponseTopupMonthMethod) *response.ApiResponseTopupMonthMethod {
	return &response.ApiResponseTopupMonthMethod{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponseTopupMonthlyMethods(s.Data),
	}
}

// ToApiResponseTopupYearMethod maps a gRPC response containing a year's worth of top-up statistics by payment method
// to an HTTP API response. It constructs an ApiResponseTopupYearMethod by copying the status and message fields,
// mapping the TopupYearMethod data slice to a slice of TopupYearMethodResponse, and assigning it to the response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupYearMethod containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupYearMethod with mapped data.
func (t *topupStatsMethodResponseMapper) ToApiResponseTopupYearMethod(s *pb.ApiResponseTopupYearMethod) *response.ApiResponseTopupYearMethod {
	return &response.ApiResponseTopupYearMethod{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponseTopupYearlyMethods(s.Data),
	}
}

// mapResponseTopupMonthlyMethod maps a gRPC response containing a month's worth of top-up statistics
// by payment method to an HTTP API response format. It constructs a TopupMonthMethodResponse by
// copying the month, top-up method, total top-ups, and total amount fields from the gRPC response
// to the API response.
func (t *topupStatsMethodResponseMapper) mapResponseTopupMonthlyMethod(s *pb.TopupMonthMethodResponse) *response.TopupMonthMethodResponse {
	return &response.TopupMonthMethodResponse{
		Month:       s.Month,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTopupMonthlyMethods maps a slice of gRPC responses containing monthly top-up statistics by payment method
// to a slice of HTTP API response formats. It iterates over the gRPC response slice, mapping each individual response
// using mapResponseTopupMonthlyMethod, and returns a slice of mapped TopupMonthMethodResponse.
//
// Args:
//   - s: A slice of pointers to pb.TopupMonthMethodResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupMonthMethodResponse with the mapped data.
func (t *topupStatsMethodResponseMapper) mapResponseTopupMonthlyMethods(s []*pb.TopupMonthMethodResponse) []*response.TopupMonthMethodResponse {
	var topupProtos []*response.TopupMonthMethodResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupMonthlyMethod(topup))
	}
	return topupProtos
}

// mapResponseTopupYearlyMethod maps a gRPC response containing a year's worth of top-up statistics
// by payment method to an HTTP API response format. It constructs a TopupYearlyMethodResponse by
// copying the year, top-up method, total top-ups, and total amount fields from the gRPC response
// to the API response.
//
// Args:
//   - s: A pointer to a pb.TopupYearlyMethodResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TopupYearlyMethodResponse with the mapped data.
func (t *topupStatsMethodResponseMapper) mapResponseTopupYearlyMethod(s *pb.TopupYearlyMethodResponse) *response.TopupYearlyMethodResponse {
	return &response.TopupYearlyMethodResponse{
		Year:        s.Year,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTopupYearlyMethods maps a slice of gRPC responses containing yearly top-up statistics
// by payment method to a slice of HTTP API response formats. It iterates over the gRPC response slice,
// mapping each individual response using mapResponseTopupYearlyMethod, and returns a slice of mapped
// TopupYearlyMethodResponse.
//
// Args:
//   - s: A slice of pointers to pb.TopupYearlyMethodResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupYearlyMethodResponse with the mapped data.
func (t *topupStatsMethodResponseMapper) mapResponseTopupYearlyMethods(s []*pb.TopupYearlyMethodResponse) []*response.TopupYearlyMethodResponse {
	var topupProtos []*response.TopupYearlyMethodResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupYearlyMethod(topup))
	}
	return topupProtos
}
