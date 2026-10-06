package transactionapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transactionStatsStatusResponseMapper struct {
}

func NewTransactionStatsStatusResponseMapper() TransactionStatsStatusResponseMapper {
	return &transactionStatsStatusResponseMapper{}
}

// ToApiResponseTransactionMonthStatusSuccess maps a single gRPC transaction month status response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionMonthStatusSuccess containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionMonthStatusSuccess containing the mapped data, including status, message,
//     and a single mapped transaction month status response.
func (m *transactionStatsStatusResponseMapper) ToApiResponseTransactionMonthStatusSuccess(pbResponse *pb.ApiResponseTransactionMonthStatusSuccess) *response.ApiResponseTransactionMonthStatusSuccess {

	return &response.ApiResponseTransactionMonthStatusSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.ToResponsesTransactionMonthStatusSuccess(pbResponse.Data),
	}
}

// ToApiResponseTransactionYearStatusSuccess maps a single gRPC transaction year status response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionYearStatusSuccess containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionYearStatusSuccess containing the mapped data, including status, message,
//     and a single mapped transaction year status response.
func (m *transactionStatsStatusResponseMapper) ToApiResponseTransactionYearStatusSuccess(pbResponse *pb.ApiResponseTransactionYearStatusSuccess) *response.ApiResponseTransactionYearStatusSuccess {

	return &response.ApiResponseTransactionYearStatusSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapTransactionResponsesYearStatusSuccess(pbResponse.Data),
	}
}

// ToApiResponseTransactionMonthStatusFailed maps a single gRPC transaction month status failed response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionMonthStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionMonthStatusFailed containing the mapped data, including status, message,
//     and a single mapped transaction month status failed response.
func (m *transactionStatsStatusResponseMapper) ToApiResponseTransactionMonthStatusFailed(pbResponse *pb.ApiResponseTransactionMonthStatusFailed) *response.ApiResponseTransactionMonthStatusFailed {

	return &response.ApiResponseTransactionMonthStatusFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesTransactionMonthStatusFailed(pbResponse.Data),
	}
}

// ToApiResponseTransactionYearStatusFailed maps a single gRPC transaction year status failed response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionYearStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionYearStatusFailed containing the mapped data, including status, message,
//     and a single mapped transaction year status failed response.
func (m *transactionStatsStatusResponseMapper) ToApiResponseTransactionYearStatusFailed(pbResponse *pb.ApiResponseTransactionYearStatusFailed) *response.ApiResponseTransactionYearStatusFailed {

	return &response.ApiResponseTransactionYearStatusFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapTransactionResponsesYearStatusFailed(pbResponse.Data),
	}
}

// mapResponseTransactionMonthStatusSuccess maps a gRPC transaction month status success response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionMonthStatusSuccessResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponseMonthStatusSuccess containing the mapped data,
//     including the year, month, total successful transactions, and total amount.
func (m *transactionStatsStatusResponseMapper) mapResponseTransactionMonthStatusSuccess(s *pb.TransactionMonthStatusSuccessResponse) *response.TransactionResponseMonthStatusSuccess {
	return &response.TransactionResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToResponsesTransactionMonthStatusSuccess maps a slice of gRPC transaction month status success responses
// to a slice of HTTP API responses.
//
// Args:
//   - transactions: A slice of pointers to pb.TransactionMonthStatusSuccessResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionResponseMonthStatusSuccess containing the mapped data.
func (m *transactionStatsStatusResponseMapper) ToResponsesTransactionMonthStatusSuccess(transactions []*pb.TransactionMonthStatusSuccessResponse) []*response.TransactionResponseMonthStatusSuccess {
	var transactionRecords []*response.TransactionResponseMonthStatusSuccess
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, m.mapResponseTransactionMonthStatusSuccess(transaction))
	}
	return transactionRecords
}

// mapTransactionResponseYearStatusSuccess maps a gRPC transaction year status success response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionYearStatusSuccessResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponseYearStatusSuccess containing the mapped data,
//     including the year, total successful transactions, and total amount.
func (m *transactionStatsStatusResponseMapper) mapTransactionResponseYearStatusSuccess(s *pb.TransactionYearStatusSuccessResponse) *response.TransactionResponseYearStatusSuccess {
	return &response.TransactionResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// mapTransactionResponsesYearStatusSuccess maps a slice of gRPC transaction year status success responses
// to a slice of HTTP API responses.
//
// Args:
//   - transactions: A slice of pointers to pb.TransactionYearStatusSuccessResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionResponseYearStatusSuccess containing the mapped data.
func (m *transactionStatsStatusResponseMapper) mapTransactionResponsesYearStatusSuccess(transactions []*pb.TransactionYearStatusSuccessResponse) []*response.TransactionResponseYearStatusSuccess {
	var transactionRecords []*response.TransactionResponseYearStatusSuccess
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, m.mapTransactionResponseYearStatusSuccess(transaction))
	}
	return transactionRecords
}

// mapResponseTransactionMonthStatusFailed maps a gRPC transaction month status failed response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionMonthStatusFailedResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponseMonthStatusFailed containing the mapped data,
//     including the year, month, total failed transactions, and total amount.
func (m *transactionStatsStatusResponseMapper) mapResponseTransactionMonthStatusFailed(s *pb.TransactionMonthStatusFailedResponse) *response.TransactionResponseMonthStatusFailed {
	return &response.TransactionResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponsesTransactionMonthStatusFailed maps a slice of gRPC transaction month status failed responses
// to a slice of HTTP API responses.
//
// Args:
//   - transactions: A slice of pointers to pb.TransactionMonthStatusFailedResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionResponseMonthStatusFailed containing the mapped data.
func (m *transactionStatsStatusResponseMapper) mapResponsesTransactionMonthStatusFailed(transactions []*pb.TransactionMonthStatusFailedResponse) []*response.TransactionResponseMonthStatusFailed {
	var transactionRecords []*response.TransactionResponseMonthStatusFailed
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, m.mapResponseTransactionMonthStatusFailed(transaction))
	}
	return transactionRecords
}

// mapTransactionResponseYearStatusFailed maps a gRPC transaction year status failed response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionYearStatusFailedResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponseYearStatusFailed containing the mapped data,
//     including the year, total failed transactions, and total amount.
func (m *transactionStatsStatusResponseMapper) mapTransactionResponseYearStatusFailed(s *pb.TransactionYearStatusFailedResponse) *response.TransactionResponseYearStatusFailed {
	return &response.TransactionResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapTransactionResponsesYearStatusFailed maps a slice of gRPC transaction year status failed responses
// to a slice of HTTP API responses.
//
// Args:
//   - transactions: A slice of pointers to pb.TransactionYearStatusFailedResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionResponseYearStatusFailed containing the mapped data.
func (m *transactionStatsStatusResponseMapper) mapTransactionResponsesYearStatusFailed(transactions []*pb.TransactionYearStatusFailedResponse) []*response.TransactionResponseYearStatusFailed {
	var transactionRecords []*response.TransactionResponseYearStatusFailed
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, m.mapTransactionResponseYearStatusFailed(transaction))
	}
	return transactionRecords
}
