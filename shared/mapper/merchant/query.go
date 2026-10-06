package merchantapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper"
)

type merchantQueryResponseMapper struct {
}

func NewMerchantQueryResponseMapper() MerchantQueryResponseMapper {
	return &merchantQueryResponseMapper{}
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
func (m *merchantQueryResponseMapper) ToApiResponseMerchant(merchants *pb.ApiResponseMerchant) *response.ApiResponseMerchant {
	return &response.ApiResponseMerchant{
		Status:  merchants.Status,
		Message: merchants.Message,
		Data:    m.mapMerchantResponse(merchants.Data),
	}
}

// ToApiResponseMerchants maps a gRPC merchant response slice to an HTTP API response. It constructs
// an ApiResponsesMerchant by copying the status and message fields and mapping the merchant
// data to a slice of MerchantResponse.
func (m *merchantQueryResponseMapper) ToApiResponseMerchants(merchants *pb.ApiResponsesMerchant) *response.ApiResponsesMerchant {
	return &response.ApiResponsesMerchant{
		Status:  merchants.Status,
		Message: merchants.Message,
		Data:    m.mapMerchantResponses(merchants.Data),
	}
}

// ToApiResponsesMerchant maps a paginated list of merchants to a paginated HTTP API response.
// It constructs an ApiResponsePaginationMerchant by copying the status and message fields
// and mapping the merchant data slice to a slice of MerchantResponse, and including pagination
// metadata.
//
// Args:
//
//	merchants: A pointer to a pb.ApiResponsePaginationMerchant containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsePaginationMerchant with mapped data and pagination info.
func (m *merchantQueryResponseMapper) ToApiResponsesMerchant(merchants *pb.ApiResponsePaginationMerchant) *response.ApiResponsePaginationMerchant {
	return &response.ApiResponsePaginationMerchant{
		Status:     merchants.Status,
		Message:    merchants.Message,
		Data:       m.mapMerchantResponses(merchants.Data),
		Pagination: apimapper.MapPaginationMeta(merchants.PaginationMeta),
	}
}

// ToApiResponsesMerchantDeleteAt maps a paginated list of soft-deleted merchants to an HTTP API response.
// It constructs an ApiResponsePaginationMerchantDeleteAt by copying the status and message fields,
// mapping the merchant data slice to a slice of MerchantResponseDeleteAt, and including pagination metadata.
//
// Args:
//
//	merchants: A pointer to a pb.ApiResponsePaginationMerchantDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsePaginationMerchantDeleteAt with mapped data and pagination info.
func (m *merchantQueryResponseMapper) ToApiResponsesMerchantDeleteAt(merchants *pb.ApiResponsePaginationMerchantDeleteAt) *response.ApiResponsePaginationMerchantDeleteAt {
	return &response.ApiResponsePaginationMerchantDeleteAt{
		Status:     merchants.Status,
		Message:    merchants.Message,
		Data:       m.mapMerchantResponsesDeleteAt(merchants.Data),
		Pagination: apimapper.MapPaginationMeta(merchants.PaginationMeta),
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
func (m *merchantQueryResponseMapper) mapMerchantResponse(merchant *pb.MerchantResponse) *response.MerchantResponse {
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

// mapMerchantResponses maps a slice of gRPC MerchantResponse to a slice of HTTP API responses. It
// constructs a slice of ApiResponseMerchant by copying the status, message, ID, name,
// status, API key, user ID, created at, and updated at fields from the gRPC
// response.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantResponse containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponse with the mapped data.
func (m *merchantQueryResponseMapper) mapMerchantResponses(r []*pb.MerchantResponse) []*response.MerchantResponse {
	var responseMerchants []*response.MerchantResponse
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapMerchantResponse(merchant))
	}

	return responseMerchants
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
func (m *merchantQueryResponseMapper) mapMerchantResponseDeleteAt(merchant *pb.MerchantResponseDeleteAt) *response.MerchantResponseDeleteAt {
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

// mapMerchantResponsesDeleteAt maps a slice of gRPC MerchantResponseDeleteAt to a slice of HTTP API responses.
// It constructs a slice of ApiResponseMerchantDeleteAt by iterating over each MerchantResponseDeleteAt
// in the input slice and mapping them using mapMerchantResponseDeleteAt. It handles the optional DeletedAt field.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantResponseDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponseDeleteAt with the mapped data.
func (m *merchantQueryResponseMapper) mapMerchantResponsesDeleteAt(r []*pb.MerchantResponseDeleteAt) []*response.MerchantResponseDeleteAt {
	var responseMerchants []*response.MerchantResponseDeleteAt
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapMerchantResponseDeleteAt(merchant))
	}

	return responseMerchants
}
