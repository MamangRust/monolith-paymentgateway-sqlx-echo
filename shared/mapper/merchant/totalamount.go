package merchantapimapper

import (
	pbstats "github.com/MamangRust/monolith-payment-gateway-pb/merchant/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type merchantStatsTotalAmountResponseMapper struct{}

// NewMerchantStatsTotalAmountResponseMapper creates a new instance of merchantStatsTotalAmountResponseMapper.
// It returns a MerchantStatsTotalAmountResponseMapper interface for mapping monthly and yearly
// total financial statistics of merchants to API responses.
func NewMerchantStatsTotalAmountResponseMapper() MerchantStatsTotalAmountResponseMapper {
	return &merchantStatsTotalAmountResponseMapper{}
}

// ToApiResponseMonthlyTotalAmounts converts monthly total amount data from a gRPC response
// into an HTTP API response format. It constructs an ApiResponseMerchantMonthlyTotalAmount by
// copying the status and message fields and mapping the data to a slice of
// MerchantResponseMonthlyTotalAmount.
//
// Args:
//
//	ms: A pointer to a pbstats.ApiResponseMerchantMonthlyTotalAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantMonthlyTotalAmount with mapped data.
func (m *merchantStatsTotalAmountResponseMapper) ToApiResponseMonthlyTotalAmounts(ms *pbstats.ApiResponseMerchantMonthlyTotalAmount) *response.ApiResponseMerchantMonthlyTotalAmount {
	return &response.ApiResponseMerchantMonthlyTotalAmount{
		Status:  ms.Status,
		Message: ms.Message,
		Data:    m.mapResponsesMonthlyTotalAmount(ms.Data),
	}
}

// ToApiResponseYearlyTotalAmounts converts yearly total amount data from a gRPC response
// into an HTTP API response format. It constructs an ApiResponseMerchantYearlyTotalAmount by
// copying the status and message fields and mapping the data to a slice of
// MerchantResponseYearlyTotalAmount.
//
// Args:
//
//	ms: A pointer to a pbstats.ApiResponseMerchantYearlyTotalAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantYearlyTotalAmount with mapped data.
func (m *merchantStatsTotalAmountResponseMapper) ToApiResponseYearlyTotalAmounts(ms *pbstats.ApiResponseMerchantYearlyTotalAmount) *response.ApiResponseMerchantYearlyTotalAmount {
	return &response.ApiResponseMerchantYearlyTotalAmount{
		Status:  ms.Status,
		Message: ms.Message,
		Data:    m.mapResponsesYearlyTotalAmount(ms.Data),
	}
}

// mapResponseMonthlyTotalAmount maps a single gRPC MerchantResponseMonthlyTotalAmount to an HTTP API response.
// It constructs a MerchantResponseMonthlyTotalAmount by copying the month, year, and total amount from the gRPC response.
//
// Args:
//
//	ms: A pointer to a pbstats.MerchantResponseMonthlyTotalAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponseMonthlyTotalAmount with the mapped data.
func (m *merchantStatsTotalAmountResponseMapper) mapResponseMonthlyTotalAmount(ms *pbstats.MerchantResponseMonthlyTotalAmount) *response.MerchantResponseMonthlyTotalAmount {
	return &response.MerchantResponseMonthlyTotalAmount{
		Month:       ms.Month,
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// mapResponsesMonthlyTotalAmount maps a slice of gRPC MerchantResponseMonthlyTotalAmount
// to a slice of HTTP API responses. It constructs a slice of MerchantResponseMonthlyTotalAmount
// by iterating over each gRPC response and mapping them using mapResponseMonthlyTotalAmount.
//
// Args:
//
//	r: A slice of pointers to pbstats.MerchantResponseMonthlyTotalAmount containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponseMonthlyTotalAmount with the mapped data.
func (m *merchantStatsTotalAmountResponseMapper) mapResponsesMonthlyTotalAmount(r []*pbstats.MerchantResponseMonthlyTotalAmount) []*response.MerchantResponseMonthlyTotalAmount {
	var responseMerchants []*response.MerchantResponseMonthlyTotalAmount
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseMonthlyTotalAmount(merchant))
	}

	return responseMerchants
}

// mapResponseYearlyTotalAmount maps a single gRPC MerchantResponseYearlyTotalAmount to an HTTP API response.
// It constructs a MerchantResponseYearlyTotalAmount by copying the year and total amount from the gRPC response.
//
// Args:
//
//	ms: A pointer to a pbstats.MerchantResponseYearlyTotalAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponseYearlyTotalAmount with the mapped data.
func (m *merchantStatsTotalAmountResponseMapper) mapResponseYearlyTotalAmount(ms *pbstats.MerchantResponseYearlyTotalAmount) *response.MerchantResponseYearlyTotalAmount {
	return &response.MerchantResponseYearlyTotalAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// mapResponsesYearlyTotalAmount maps a slice of gRPC MerchantResponseYearlyTotalAmount
// to a slice of HTTP API responses. It constructs a slice of MerchantResponseYearlyTotalAmount
// by copying the year and total amount from each gRPC response.
//
// Args:
//
//	r: A slice of pointers to pbstats.MerchantResponseYearlyTotalAmount containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponseYearlyTotalAmount with the mapped data.
func (m *merchantStatsTotalAmountResponseMapper) mapResponsesYearlyTotalAmount(r []*pbstats.MerchantResponseYearlyTotalAmount) []*response.MerchantResponseYearlyTotalAmount {
	var responseMerchants []*response.MerchantResponseYearlyTotalAmount
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseYearlyTotalAmount(merchant))
	}

	return responseMerchants
}
