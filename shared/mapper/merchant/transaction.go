package merchantapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper"
)

type merchantTransactionResponseMapper struct{}

func NewMerchantTransactionResponseMapper() MerchantTransactionResponseMapper {
	return &merchantTransactionResponseMapper{}
}

// ToApiResponseMerchantsTransactionResponse maps a paginated gRPC response of merchant transactions to an HTTP API response.
// It constructs an ApiResponsePaginationMerchantTransaction by copying the status and message fields,
// mapping the transaction data slice to a slice of MerchantTransactionResponse, and including pagination metadata.
//
// Args:
//
//	merchants: A pointer to a pb.ApiResponsePaginationMerchantTransaction containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsePaginationMerchantTransaction with mapped data and pagination info.
func (m *merchantTransactionResponseMapper) ToApiResponseMerchantsTransactionResponse(merchants *pb.ApiResponsePaginationMerchantTransaction) *response.ApiResponsePaginationMerchantTransaction {

	return &response.ApiResponsePaginationMerchantTransaction{
		Status:     merchants.Status,
		Message:    merchants.Message,
		Data:       m.mapMerchantTransactionResponses(merchants.Data),
		Pagination: apimapper.MapPaginationMeta(merchants.PaginationMeta),
	}
}

// mapMerchantTransactionResponse maps a gRPC MerchantTransactionResponse to an HTTP API response.
// It constructs a MerchantTransactionResponse by copying fields such as ID, CardNumber,
// Amount, PaymentMethod, MerchantID, MerchantName, TransactionTime, CreatedAt, and UpdatedAt
// from the gRPC response.
//
// Args:
//
//	merchant: A pointer to a pb.MerchantTransactionResponse containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantTransactionResponse with the mapped data.
func (m *merchantTransactionResponseMapper) mapMerchantTransactionResponse(merchant *pb.MerchantTransactionResponse) *response.MerchantTransactionResponse {

	return &response.MerchantTransactionResponse{
		ID:              int(merchant.Id),
		CardNumber:      merchant.CardNumber,
		Amount:          merchant.Amount,
		PaymentMethod:   merchant.PaymentMethod,
		MerchantID:      merchant.MerchantId,
		MerchantName:    merchant.MerchantName,
		TransactionTime: merchant.TransactionTime,
		CreatedAt:       merchant.CreatedAt,
		UpdatedAt:       merchant.UpdatedAt,
	}
}

// mapMerchantTransactionResponses maps a slice of gRPC MerchantTransactionResponses to a slice of HTTP API responses.
// It constructs a slice of MerchantTransactionResponse by copying fields such as ID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
// TransactionTime, CreatedAt, and UpdatedAt from the gRPC response.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantTransactionResponse containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantTransactionResponse with the mapped data.
func (m *merchantTransactionResponseMapper) mapMerchantTransactionResponses(r []*pb.MerchantTransactionResponse) []*response.MerchantTransactionResponse {
	var responseMerchants []*response.MerchantTransactionResponse
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapMerchantTransactionResponse(merchant))
	}

	return responseMerchants
}
