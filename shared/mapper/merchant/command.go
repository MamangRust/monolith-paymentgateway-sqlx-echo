package merchantapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type merchantCommandResponse struct {
}

func NewMerchantCommandResponseMapper() MerchantCommandResponseMapper {
	return &merchantCommandResponse{}
}

// ToApiResponseMerchant maps a gRPC merchant response to an HTTP API response. It constructs
// an ApiResponseMerchant by copying the status and message fields and mapping the merchant
// data to a MerchantResponse.
//
// Args:
//
//	merchants: A pointer to a pb.ApiResponseMerchant containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchant with mapped data.
func (m *merchantCommandResponse) ToApiResponseMerchant(merchants *pb.ApiResponseMerchant) *response.ApiResponseMerchant {
	return &response.ApiResponseMerchant{
		Status:  merchants.Status,
		Message: merchants.Message,
		Data:    m.mapMerchantResponse(merchants.Data),
	}
}

func (m *merchantCommandResponse) ToApiResponseMerchantDeleteAt(merchants *pb.ApiResponseMerchantDeleteAt) *response.ApiResponseMerchantDeleteAt {
	return &response.ApiResponseMerchantDeleteAt{
		Status:  merchants.Status,
		Message: merchants.Message,
		Data:    m.mapMerchantResponseDeleteAt(merchants.Data),
	}
}

func (s *merchantCommandResponse) ToApiResponseMerchantAll(card *pb.ApiResponseMerchantAll) *response.ApiResponseMerchantAll {
	return &response.ApiResponseMerchantAll{
		Status:  card.Status,
		Message: card.Message,
	}
}

func (s *merchantCommandResponse) ToApiResponseMerchantDelete(card *pb.ApiResponseMerchantDelete) *response.ApiResponseMerchantDelete {
	return &response.ApiResponseMerchantDelete{
		Status:  card.Status,
		Message: card.Message,
	}
}

// mapMerchantResponse maps a gRPC MerchantResponse to an HTTP API response. It
// constructs an ApiResponseMerchant by copying the status, message, ID, name,
// status, API key, user ID, created at, and updated at fields from the gRPC
// response.
//
// Args:
//
//	merchant: A pointer to a pb.MerchantResponse containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponse with the mapped data.
func (m *merchantCommandResponse) mapMerchantResponse(merchant *pb.MerchantResponse) *response.MerchantResponse {
	return &response.MerchantResponse{
		ID:        int(merchant.Id),
		Name:      merchant.Name,
		Status:    merchant.Status,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserId),
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
	}
}

// mapMerchantResponseDeleteAt maps a gRPC MerchantResponseDeleteAt to an HTTP API response. It
// constructs an ApiResponseMerchantDeleteAt by copying the status, message, ID, name,
// status, API key, user ID, created at, updated at, and deleted at fields from the gRPC
// response.
//
// Args:
//
//	merchant: A pointer to a pb.MerchantResponseDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponseDeleteAt with the mapped data.
func (m *merchantCommandResponse) mapMerchantResponseDeleteAt(merchant *pb.MerchantResponseDeleteAt) *response.MerchantResponseDeleteAt {
	var deletedAt string
	if merchant.DeletedAt != nil {
		deletedAt = merchant.DeletedAt.Value
	}

	return &response.MerchantResponseDeleteAt{
		ID:        int(merchant.Id),
		Name:      merchant.Name,
		Status:    merchant.Status,
		UserID:    int(merchant.UserId),
		ApiKey:    merchant.ApiKey,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
		DeletedAt: &deletedAt,
	}
}
