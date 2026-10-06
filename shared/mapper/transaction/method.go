package transactionapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transactionStatsMethodResponseMapper struct{}

func NewTransactionStatsMethodResponseMapper() TransactionStatsMethodResponseMapper {
	return &transactionStatsMethodResponseMapper{}
}

// ToApiResponseTransactionMonthMethod maps a single gRPC transaction month method response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionMonthMethod containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionMonthMethod containing the mapped data, including status, message,
//     and a single mapped transaction month method response.
func (m *transactionStatsMethodResponseMapper) ToApiResponseTransactionMonthMethod(pbResponse *pb.ApiResponseTransactionMonthMethod) *response.ApiResponseTransactionMonthMethod {
	return &response.ApiResponseTransactionMonthMethod{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransactionMonthMethods(pbResponse.Data),
	}
}

// ToApiResponseTransactionYearMethod maps a single gRPC transaction year method response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionYearMethod containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionYearMethod containing the mapped data, including status, message,
//     and a single mapped transaction year method response.
func (m *transactionStatsMethodResponseMapper) ToApiResponseTransactionYearMethod(pbResponse *pb.ApiResponseTransactionYearMethod) *response.ApiResponseTransactionYearMethod {
	return &response.ApiResponseTransactionYearMethod{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransactionYearMethods(pbResponse.Data),
	}
}

// mapResponseTransactionMonthMethod maps a gRPC transaction month method response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionMonthMethodResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionMonthMethodResponse containing the mapped data,
//     including the month, payment method, total transactions, and total amount.
func (m *transactionStatsMethodResponseMapper) mapResponseTransactionMonthMethod(s *pb.TransactionMonthMethodResponse) *response.TransactionMonthMethodResponse {
	return &response.TransactionMonthMethodResponse{
		Month:             s.Month,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int(s.TotalTransactions),
		TotalAmount:       int(s.TotalAmount),
	}
}

// mapResponseTransactionMonthMethods maps a slice of gRPC transaction month method responses
// to a slice of HTTP API responses.
//
// Args:
//   - s: A slice of pointers to pb.TransactionMonthMethodResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionMonthMethodResponse containing the mapped data,
//     including the month, payment method, total transactions, and total amount.
func (m *transactionStatsMethodResponseMapper) mapResponseTransactionMonthMethods(s []*pb.TransactionMonthMethodResponse) []*response.TransactionMonthMethodResponse {
	var responses []*response.TransactionMonthMethodResponse
	for _, transaction := range s {
		responses = append(responses, m.mapResponseTransactionMonthMethod(transaction))
	}
	return responses
}

// mapResponseTransactionYearMethod maps a single gRPC transaction year method response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionYearMethodResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionYearMethodResponse containing the mapped data,
//     including the year, payment method, total transactions, and total amount.
func (m *transactionStatsMethodResponseMapper) mapResponseTransactionYearMethod(s *pb.TransactionYearMethodResponse) *response.TransactionYearMethodResponse {
	return &response.TransactionYearMethodResponse{
		Year:              s.Year,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int(s.TotalTransactions),
		TotalAmount:       int(s.TotalAmount),
	}
}

// mapResponseTransactionYearMethods maps a slice of gRPC transaction year method responses
// to a slice of HTTP API responses.
//
// Args:
//   - s: A slice of pointers to pb.TransactionYearMethodResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionYearMethodResponse containing the mapped data,
//     including the year, payment method, total transactions, and total amount.
func (m *transactionStatsMethodResponseMapper) mapResponseTransactionYearMethods(s []*pb.TransactionYearMethodResponse) []*response.TransactionYearMethodResponse {
	var responses []*response.TransactionYearMethodResponse
	for _, transaction := range s {
		responses = append(responses, m.mapResponseTransactionYearMethod(transaction))
	}
	return responses
}
