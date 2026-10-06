package transactionapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// transactionStatsAmountResponseMapper is a struct that implements the TransactionStatsAmountResponseMapper interface.
type transactionStatsAmountResponseMapper struct {
}

// NewTransactionStatsAmountResponseMapper creates a new instance of transactionStatsAmountResponseMapper.
// It returns a pointer to a transactionStatsAmountResponseMapper which implements the
// TransactionStatsAmountResponseMapper interface for mapping database rows to domain
// models.
func NewTransactionStatsAmountResponseMapper() *transactionStatsAmountResponseMapper {
	return &transactionStatsAmountResponseMapper{}
}

// ToApiResponseTransactionMonthAmount maps a gRPC transaction month amount response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionMonthAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionMonthAmount containing the mapped data, including status, message,
//     and a single mapped transaction month amount response.
func (m *transactionStatsAmountResponseMapper) ToApiResponseTransactionMonthAmount(pbResponse *pb.ApiResponseTransactionMonthAmount) *response.ApiResponseTransactionMonthAmount {

	return &response.ApiResponseTransactionMonthAmount{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransactionMonthAmounts(pbResponse.Data),
	}
}

// ToApiResponseTransactionYearAmount maps a single gRPC transaction year amount response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionYearAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionYearAmount containing the mapped data, including status, message,
//     and a single mapped transaction year amount response.
func (m *transactionStatsAmountResponseMapper) ToApiResponseTransactionYearAmount(pbResponse *pb.ApiResponseTransactionYearAmount) *response.ApiResponseTransactionYearAmount {

	return &response.ApiResponseTransactionYearAmount{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransactionYearlyAmounts(pbResponse.Data),
	}
}

// mapResponseTransactionMonthAmount maps a gRPC transaction month amount response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionMonthAmountResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionMonthAmountResponse containing the mapped data,
//     including the month and total amount.
func (m *transactionStatsAmountResponseMapper) mapResponseTransactionMonthAmount(s *pb.TransactionMonthAmountResponse) *response.TransactionMonthAmountResponse {
	return &response.TransactionMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTransactionMonthAmounts maps a slice of gRPC transaction month amount responses
// to a slice of HTTP API responses.
//
// Args:
//   - s: A slice of pointers to pb.TransactionMonthAmountResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionMonthAmountResponse containing the mapped data,
//     including the month and total amount.
func (m *transactionStatsAmountResponseMapper) mapResponseTransactionMonthAmounts(s []*pb.TransactionMonthAmountResponse) []*response.TransactionMonthAmountResponse {
	var responses []*response.TransactionMonthAmountResponse
	for _, transaction := range s {
		responses = append(responses, m.mapResponseTransactionMonthAmount(transaction))
	}
	return responses
}

// mapResponseTransactionYearlyAmount maps a gRPC transaction yearly amount response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionYearlyAmountResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionYearlyAmountResponse containing the mapped data,
//     including the year and total amount.
func (m *transactionStatsAmountResponseMapper) mapResponseTransactionYearlyAmount(s *pb.TransactionYearlyAmountResponse) *response.TransactionYearlyAmountResponse {
	return &response.TransactionYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTransactionYearlyAmounts maps a slice of gRPC transaction yearly amount responses
// to a slice of HTTP API responses.
//
// Args:
//   - s: A slice of pointers to pb.TransactionYearlyAmountResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionYearlyAmountResponse containing the mapped data,
//     including the year and total amount.
func (m *transactionStatsAmountResponseMapper) mapResponseTransactionYearlyAmounts(s []*pb.TransactionYearlyAmountResponse) []*response.TransactionYearlyAmountResponse {
	var responses []*response.TransactionYearlyAmountResponse
	for _, transaction := range s {
		responses = append(responses, m.mapResponseTransactionYearlyAmount(transaction))
	}
	return responses
}
