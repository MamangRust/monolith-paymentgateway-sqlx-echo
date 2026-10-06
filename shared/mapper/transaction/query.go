package transactionapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper"
)

type transactionQueryResponseMapper struct {
}

func NewTransactionQueryResponseMapper() TransactionQueryResponseMapper {
	return &transactionQueryResponseMapper{}
}

// ToApiResponseTransaction maps a single gRPC transaction response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransaction containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransaction containing the mapped data, including status, message,
//     and a single mapped transaction response.
func (m *transactionQueryResponseMapper) ToApiResponseTransaction(pbResponse *pb.ApiResponseTransaction) *response.ApiResponseTransaction {
	return &response.ApiResponseTransaction{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransaction(pbResponse.Data),
	}
}

// ToApiResponseTransactions maps multiple gRPC transaction responses into an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactions containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactions containing the mapped data, including status, message,
//     and a slice of mapped transaction responses.
func (m *transactionQueryResponseMapper) ToApiResponseTransactions(pbResponse *pb.ApiResponseTransactions) *response.ApiResponseTransactions {

	return &response.ApiResponseTransactions{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesTransaction(pbResponse.Data),
	}
}

// ToApiResponsePaginationTransaction maps a paginated gRPC response of transactions to an HTTP API response.
// It constructs an ApiResponsePaginationTransaction by copying the status and message fields,
// mapping the transaction data slice to a slice of TransactionResponse, and including pagination metadata.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponsePaginationTransaction containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsePaginationTransaction with mapped data and pagination info.
func (m *transactionQueryResponseMapper) ToApiResponsePaginationTransaction(pbResponse *pb.ApiResponsePaginationTransaction) *response.ApiResponsePaginationTransaction {

	return &response.ApiResponsePaginationTransaction{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       m.mapResponsesTransaction(pbResponse.Data),
		Pagination: apimapper.MapPaginationMeta(pbResponse.PaginationMeta),
	}
}

// ToApiResponsePaginationTransactionDeleteAt maps a paginated gRPC response of
// soft-deleted transactions to an HTTP API response. It constructs an
// ApiResponsePaginationTransactionDeleteAt by copying the status and message fields,
// mapping the transaction data slice to a slice of TransactionResponseDeleteAt, and
// including pagination metadata.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponsePaginationTransactionDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationTransactionDeleteAt with mapped data and pagination info.
func (m *transactionQueryResponseMapper) ToApiResponsePaginationTransactionDeleteAt(pbResponse *pb.ApiResponsePaginationTransactionDeleteAt) *response.ApiResponsePaginationTransactionDeleteAt {

	return &response.ApiResponsePaginationTransactionDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       m.ToResponsesTransactionDeleteAt(pbResponse.Data),
		Pagination: apimapper.MapPaginationMeta(pbResponse.PaginationMeta),
	}
}

// mapResponseTransaction maps a single gRPC transaction response to an HTTP API response.
//
// Args:
//   - transaction: A pointer to a pb.TransactionResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponse containing the mapped data.
func (m *transactionQueryResponseMapper) mapResponseTransaction(transaction *pb.TransactionResponse) *response.TransactionResponse {
	return &response.TransactionResponse{
		ID:              int(transaction.Id),
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          int(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime,
		MerchantID:      int(transaction.MerchantId),
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	}
}

// mapResponsesTransaction maps a slice of gRPC transaction responses to a slice of HTTP API responses.
//
// Args:
//   - transactions: A slice of pointers to pb.TransactionResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionResponse containing the mapped data.
func (m *transactionQueryResponseMapper) mapResponsesTransaction(transactions []*pb.TransactionResponse) []*response.TransactionResponse {
	var result []*response.TransactionResponse
	for _, transaction := range transactions {
		result = append(result, m.mapResponseTransaction(transaction))
	}
	return result

}

// mapResponseTransactionDeleteAt maps a single gRPC transaction delete response to an HTTP API response.
//
// Args:
//   - transaction: A pointer to a pb.TransactionResponseDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponseDeleteAt containing the mapped data.
func (m *transactionQueryResponseMapper) mapResponseTransactionDeleteAt(transaction *pb.TransactionResponseDeleteAt) *response.TransactionResponseDeleteAt {
	var deletedAt string
	if transaction.DeletedAt != nil {
		deletedAt = transaction.DeletedAt.Value
	}

	return &response.TransactionResponseDeleteAt{
		ID:              int(transaction.Id),
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          int(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime,
		MerchantID:      int(transaction.MerchantId),
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
		DeletedAt:       &deletedAt,
	}
}

// ToResponsesTransactionDeleteAt maps a slice of gRPC transaction delete responses to a slice of HTTP API responses.
//
// Args:
//   - transactions: A slice of pointers to pb.TransactionResponseDeleteAt containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionResponseDeleteAt containing the mapped data.
func (m *transactionQueryResponseMapper) ToResponsesTransactionDeleteAt(transactions []*pb.TransactionResponseDeleteAt) []*response.TransactionResponseDeleteAt {
	var result []*response.TransactionResponseDeleteAt
	for _, transaction := range transactions {
		result = append(result, m.mapResponseTransactionDeleteAt(transaction))
	}
	return result
}
