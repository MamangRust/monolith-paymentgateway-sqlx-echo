package topupapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper"
)

type topupQueryResponseMapper struct{}

func NewTopupQueryResponseMapper() TopupQueryResponseMapper {
	return &topupQueryResponseMapper{}
}

// ToApiResponseTopup maps a single gRPC top-up response to an HTTP API response.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopup containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopup containing the mapped data, including status, message, and a single
//     mapped top-up response.
func (t *topupQueryResponseMapper) ToApiResponseTopup(s *pb.ApiResponseTopup) *response.ApiResponseTopup {
	return &response.ApiResponseTopup{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponseTopup(s.Data),
	}
}

// ToApiResponsePaginationTopup maps a paginated gRPC top-up response to an HTTP API response. It constructs an
// ApiResponsePaginationTopup by copying the status and message fields, mapping the top-up data slice to a slice
// of TopupResponse, and including pagination metadata.
//
// Args:
//   - s: A pointer to a pb.ApiResponsePaginationTopup containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationTopup with mapped data and pagination info.
func (t *topupQueryResponseMapper) ToApiResponsePaginationTopup(s *pb.ApiResponsePaginationTopup) *response.ApiResponsePaginationTopup {
	return &response.ApiResponsePaginationTopup{
		Status:     s.Status,
		Message:    s.Message,
		Data:       t.mapResponsesTopup(s.Data),
		Pagination: apimapper.MapPaginationMeta(s.PaginationMeta),
	}
}

// ToApiResponsePaginationTopupDeleteAt maps a paginated gRPC response of soft-deleted top-ups
// to an HTTP API response. It constructs an ApiResponsePaginationTopupDeleteAt by copying
// the status and message fields, mapping the top-up data slice to a slice of TopupDeleteAtResponse,
// and including pagination metadata.
//
// Args:
//   - s: A pointer to a pb.ApiResponsePaginationTopupDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationTopupDeleteAt with mapped data and pagination info.
func (t *topupQueryResponseMapper) ToApiResponsePaginationTopupDeleteAt(s *pb.ApiResponsePaginationTopupDeleteAt) *response.ApiResponsePaginationTopupDeleteAt {
	return &response.ApiResponsePaginationTopupDeleteAt{
		Status:     s.Status,
		Message:    s.Message,
		Data:       t.mapResponsesTopupDeleteAt(s.Data),
		Pagination: apimapper.MapPaginationMeta(s.PaginationMeta),
	}
}

// mapResponseTopup maps a single gRPC top-up response to an HTTP API response.
// It constructs a TopupResponse by copying the fields from the gRPC response.
//
// Args:
//   - topup: A pointer to a pb.TopupResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TopupResponse with mapped data.
func (t *topupQueryResponseMapper) mapResponseTopup(topup *pb.TopupResponse) *response.TopupResponse {
	return &response.TopupResponse{
		ID:          int(topup.Id),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
	}
}

// mapResponsesTopup maps a slice of gRPC top-up responses to a slice of HTTP API responses.
// It iterates over the slice of gRPC responses and maps each one to an HTTP API response
// using mapResponseTopup, returning the slice of mapped responses.
//
// Args:
//   - topups: A slice of pointers to pb.TopupResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupResponse with the mapped data.
func (t *topupQueryResponseMapper) mapResponsesTopup(topups []*pb.TopupResponse) []*response.TopupResponse {
	var responses []*response.TopupResponse

	for _, topup := range topups {
		responses = append(responses, t.mapResponseTopup(topup))
	}

	return responses
}

// mapResponseTopupDeleteAt maps a gRPC soft-deleted top-up response to an HTTP API response format.
// It constructs a TopupResponseDeleteAt by copying the fields from the gRPC response, converting
// the ID and TopupAmount fields to integers, and handling the potentially nil DeletedAt field.
//
// Args:
//   - topup: A pointer to a pb.TopupResponseDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TopupResponseDeleteAt with the mapped data, including a non-nil DeletedAt field if present.
func (t *topupQueryResponseMapper) mapResponseTopupDeleteAt(topup *pb.TopupResponseDeleteAt) *response.TopupResponseDeleteAt {
	var deletedAt string
	if topup.DeletedAt != nil {
		deletedAt = topup.DeletedAt.Value
	}

	return &response.TopupResponseDeleteAt{
		ID:          int(topup.Id),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
		DeletedAt:   &deletedAt,
	}
}

// mapResponsesTopupDeleteAt maps a slice of gRPC soft-deleted top-up responses to a slice
// of HTTP API responses. It iterates over the gRPC response slice, mapping each individual
// response using mapResponseTopupDeleteAt, and returns a slice of mapped TopupResponseDeleteAt.
//
// Args:
//   - topups: A slice of pointers to pb.TopupResponseDeleteAt containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupResponseDeleteAt with the mapped data.
func (t *topupQueryResponseMapper) mapResponsesTopupDeleteAt(topups []*pb.TopupResponseDeleteAt) []*response.TopupResponseDeleteAt {
	var responses []*response.TopupResponseDeleteAt

	for _, topup := range topups {
		responses = append(responses, t.mapResponseTopupDeleteAt(topup))
	}

	return responses
}
