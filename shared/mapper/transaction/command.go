package transactionapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transactionCommandResponseMapper struct{}

func NewTransactionCommandResponseMapper() *transactionCommandResponseMapper {
	return &transactionCommandResponseMapper{}
}

// ToApiResponseTransaction maps a single gRPC transaction response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransaction containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransaction containing the mapped data, including status, message,
//     and a single mapped transaction response.
func (m *transactionCommandResponseMapper) ToApiResponseTransaction(pbResponse *pb.ApiResponseTransaction) *response.ApiResponseTransaction {
	return &response.ApiResponseTransaction{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransaction(pbResponse.Data),
	}
}

func (m *transactionCommandResponseMapper) ToApiResponseTransactionDeleteAt(pbResponse *pb.ApiResponseTransactionDeleteAt) *response.ApiResponseTransactionDeleteAt {
	return &response.ApiResponseTransactionDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransactionDeleteAt(pbResponse.Data),
	}
}

// ToApiResponseTransactionAll maps a single gRPC transaction all response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionAll containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionAll containing the mapped data, including status and message.
func (m *transactionCommandResponseMapper) ToApiResponseTransactionAll(pbResponse *pb.ApiResponseTransactionAll) *response.ApiResponseTransactionAll {
	return &response.ApiResponseTransactionAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponseTransactionDelete maps a single gRPC transaction delete response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionDelete containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionDelete containing the mapped data, including status and message.
func (m *transactionCommandResponseMapper) ToApiResponseTransactionDelete(pbResponse *pb.ApiResponseTransactionDelete) *response.ApiResponseTransactionDelete {
	return &response.ApiResponseTransactionDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// mapResponseTransaction maps a single gRPC transaction response to an HTTP API response.
//
// Args:
//   - transaction: A pointer to a pb.TransactionResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponse containing the mapped data.
func (m *transactionCommandResponseMapper) mapResponseTransaction(transaction *pb.TransactionResponse) *response.TransactionResponse {
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

// mapResponseTransactionDeleteAt maps a single gRPC transaction delete response to an HTTP API response.
//
// Args:
//   - transaction: A pointer to a pb.TransactionResponseDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponseDeleteAt containing the mapped data.
func (m *transactionCommandResponseMapper) mapResponseTransactionDeleteAt(transaction *pb.TransactionResponseDeleteAt) *response.TransactionResponseDeleteAt {
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
