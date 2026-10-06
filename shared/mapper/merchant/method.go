package merchantapimapper

import (
	pbstats "github.com/MamangRust/monolith-payment-gateway-pb/merchant/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type merchantStatsMethodResponseMapper struct{}

// NewMerchantStatsMethodResponseMapper returns a pointer to a new merchantStatsMethodResponseMapper
// instance. This struct provides methods to map payment method statistics of a merchant to API
// responses.
func NewMerchantStatsMethodResponseMapper() MerchantStatsMethodResponseMapper {
	return &merchantStatsMethodResponseMapper{}
}

// ToApiResponseMonthlyPaymentMethods converts monthly payment method statistics from a gRPC
// response to an HTTP API response. It constructs an ApiResponseMerchantMonthlyPaymentMethod
// by copying the status and message fields and mapping the data to a slice of
// MerchantResponseMonthlyPaymentMethod.
//
// Args:
//
//	ms: A pointer to a pb.ApiResponseMerchantMonthlyPaymentMethod containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantMonthlyPaymentMethod with mapped data.
func (m *merchantStatsMethodResponseMapper) ToApiResponseMonthlyPaymentMethods(ms *pbstats.ApiResponseMerchantMonthlyPaymentMethod) *response.ApiResponseMerchantMonthlyPaymentMethod {
	return &response.ApiResponseMerchantMonthlyPaymentMethod{
		Status:  ms.Status,
		Message: ms.Message,
		Data:    m.mapResponsesMonthlyPaymentMethod(ms.Data),
	}
}

// ToApiResponseYearlyPaymentMethods converts yearly payment method statistics from
// a gRPC response to an HTTP API response. It constructs an ApiResponseMerchantYearlyPaymentMethod
// by copying the status and message fields and mapping the data to a slice of
// MerchantResponseYearlyPaymentMethod.
//
// Args:
//
//	ms: A pointer to a pb.ApiResponseMerchantYearlyPaymentMethod containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantYearlyPaymentMethod with mapped data.
func (m *merchantStatsMethodResponseMapper) ToApiResponseYearlyPaymentMethods(ms *pbstats.ApiResponseMerchantYearlyPaymentMethod) *response.ApiResponseMerchantYearlyPaymentMethod {
	return &response.ApiResponseMerchantYearlyPaymentMethod{
		Status:  ms.Status,
		Message: ms.Message,
		Data:    m.mapResponsesYearlyPaymentMethod(ms.Data),
	}
}

// mapResponseMonthlyPaymentMethod maps a single gRPC MerchantResponseMonthlyPaymentMethod to an HTTP API response.
// It constructs a MerchantResponseMonthlyPaymentMethod by copying the month, payment method, and total amount
// from the gRPC response.
//
// Args:
//
//	ms: A pointer to a pb.MerchantResponseMonthlyPaymentMethod containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponseMonthlyPaymentMethod with the mapped data.
func (m *merchantStatsMethodResponseMapper) mapResponseMonthlyPaymentMethod(ms *pbstats.MerchantResponseMonthlyPaymentMethod) *response.MerchantResponseMonthlyPaymentMethod {
	return &response.MerchantResponseMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// mapResponsesMonthlyPaymentMethod maps a slice of gRPC MerchantResponseMonthlyPaymentMethods to a slice of HTTP API responses.
// It constructs a slice of MerchantResponseMonthlyPaymentMethod by copying the month, payment method, and total amount
// from the gRPC response.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantResponseMonthlyPaymentMethod containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponseMonthlyPaymentMethod with the mapped data.
func (m *merchantStatsMethodResponseMapper) mapResponsesMonthlyPaymentMethod(r []*pbstats.MerchantResponseMonthlyPaymentMethod) []*response.MerchantResponseMonthlyPaymentMethod {
	var responseMerchants []*response.MerchantResponseMonthlyPaymentMethod
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseMonthlyPaymentMethod(merchant))
	}

	return responseMerchants
}

// mapResponseYearlyPaymentMethod maps a single gRPC MerchantResponseYearlyPaymentMethod to an HTTP API response.
// It constructs a MerchantResponseYearlyPaymentMethod by copying the year, payment method, and total amount
// from the gRPC response.
//
// Args:
//
//	ms: A pointer to a pb.MerchantResponseYearlyPaymentMethod containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponseYearlyPaymentMethod with the mapped data.
func (m *merchantStatsMethodResponseMapper) mapResponseYearlyPaymentMethod(ms *pbstats.MerchantResponseYearlyPaymentMethod) *response.MerchantResponseYearlyPaymentMethod {
	return &response.MerchantResponseYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// mapResponsesYearlyPaymentMethod maps a slice of gRPC MerchantResponseYearlyPaymentMethod
// to a slice of HTTP API responses. It constructs a slice of MerchantResponseYearlyPaymentMethod
// by copying the year, payment method, and total amount from each gRPC response.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantResponseYearlyPaymentMethod containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponseYearlyPaymentMethod with the mapped data.
func (m *merchantStatsMethodResponseMapper) mapResponsesYearlyPaymentMethod(r []*pbstats.MerchantResponseYearlyPaymentMethod) []*response.MerchantResponseYearlyPaymentMethod {
	var responseMerchants []*response.MerchantResponseYearlyPaymentMethod
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseYearlyPaymentMethod(merchant))
	}

	return responseMerchants
}
