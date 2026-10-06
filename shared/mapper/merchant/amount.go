package merchantapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type merchantStatsAmountResponseMapper struct{}

// NewMerchantStatsAmountResponseMapper returns a pointer to a new
// merchantStatsAmountResponseMapper instance. This struct provides methods to map
// monthly and yearly financial amounts of a merchant to API responses.
func NewMerchantStatsAmountResponseMapper() MerchantStatsAmountResponseMapper {
	return &merchantStatsAmountResponseMapper{}
}

// ToApiResponseMonthlyAmounts converts monthly financial amounts data from a gRPC response
// into an HTTP API response format. It constructs an ApiResponseMerchantMonthlyAmount by
// copying the status and message fields and mapping the data to a slice of
// MerchantResponseMonthlyAmount.
//
// Args:
//
//	ms: A pointer to a pb.ApiResponseMerchantMonthlyAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantMonthlyAmount with mapped data.
func (m *merchantStatsAmountResponseMapper) ToApiResponseMonthlyAmounts(ms *pb.ApiResponseMerchantMonthlyAmount) *response.ApiResponseMerchantMonthlyAmount {
	return &response.ApiResponseMerchantMonthlyAmount{
		Status:  ms.Status,
		Message: ms.Message,
		Data:    m.mapResponsesMonthlyAmount(ms.Data),
	}
}

// ToApiResponseYearlyAmounts converts yearly financial amounts data from a gRPC response
// into an HTTP API response format. It constructs an ApiResponseMerchantYearlyAmount by
// copying the status and message fields and mapping the data to a slice of
// MerchantResponseYearlyAmount.
//
// Args:
//
//	ms: A pointer to a pb.ApiResponseMerchantYearlyAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantYearlyAmount with mapped data.
func (m *merchantStatsAmountResponseMapper) ToApiResponseYearlyAmounts(ms *pb.ApiResponseMerchantYearlyAmount) *response.ApiResponseMerchantYearlyAmount {
	return &response.ApiResponseMerchantYearlyAmount{
		Status:  ms.Status,
		Message: ms.Message,
		Data:    m.mapResponsesYearlyAmount(ms.Data),
	}
}

// mapResponseMonthlyAmount maps a single gRPC MerchantResponseMonthlyAmount to an HTTP API response.
// It constructs a MerchantResponseMonthlyAmount by copying the month and total amount from the gRPC response.
//
// Args:
//
//	ms: A pointer to a pb.MerchantResponseMonthlyAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponseMonthlyAmount with the mapped data.
func (m *merchantStatsAmountResponseMapper) mapResponseMonthlyAmount(ms *pb.MerchantResponseMonthlyAmount) *response.MerchantResponseMonthlyAmount {
	return &response.MerchantResponseMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: int(ms.TotalAmount),
	}
}

// mapResponsesMonthlyAmount maps a slice of gRPC MerchantResponseMonthlyAmount
// to a slice of HTTP API responses. It constructs a slice of MerchantResponseMonthlyAmount
// by copying the month and total amount from each gRPC response.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantResponseMonthlyAmount containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponseMonthlyAmount with the mapped data.
func (m *merchantStatsAmountResponseMapper) mapResponsesMonthlyAmount(r []*pb.MerchantResponseMonthlyAmount) []*response.MerchantResponseMonthlyAmount {
	var responseMerchants []*response.MerchantResponseMonthlyAmount
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseMonthlyAmount(merchant))
	}

	return responseMerchants
}

// mapResponseYearlyAmount maps a single gRPC MerchantResponseYearlyAmount to an HTTP API response.
// It constructs a MerchantResponseYearlyAmount by copying the year and total amount from the gRPC response.
//
// Args:
//
//	ms: A pointer to a pb.MerchantResponseYearlyAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponseYearlyAmount with the mapped data.
func (m *merchantStatsAmountResponseMapper) mapResponseYearlyAmount(ms *pb.MerchantResponseYearlyAmount) *response.MerchantResponseYearlyAmount {
	return &response.MerchantResponseYearlyAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// mapResponsesYearlyAmount maps a slice of gRPC MerchantResponseYearlyAmount
// to a slice of HTTP API responses. It constructs a slice of MerchantResponseYearlyAmount
// by copying the year and total amount from each gRPC response.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantResponseYearlyAmount containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponseYearlyAmount with the mapped data.
func (m *merchantStatsAmountResponseMapper) mapResponsesYearlyAmount(r []*pb.MerchantResponseYearlyAmount) []*response.MerchantResponseYearlyAmount {
	var responseMerchants []*response.MerchantResponseYearlyAmount
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseYearlyAmount(merchant))
	}

	return responseMerchants
}
