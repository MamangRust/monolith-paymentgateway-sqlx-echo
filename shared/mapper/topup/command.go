package topupapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type topupCommandResponseMapper struct{}

func NewTopupCommandResponseMapper() TopupCommandResponseMapper {
	return &topupCommandResponseMapper{}
}

// ToApiResponseTopup maps a single gRPC top-up response to an HTTP API response.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopup containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopup containing the mapped data, including status, message, and a single
//     mapped top-up response.
func (t *topupCommandResponseMapper) ToApiResponseTopup(s *pb.ApiResponseTopup) *response.ApiResponseTopup {
	return &response.ApiResponseTopup{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponseTopup(s.Data),
	}
}

// ToApiResponseTopupDeleteAt maps a single gRPC soft-deleted top-up response to an HTTP API response.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupDeleteAt containing the mapped data, including status, message, and a single
//     mapped soft-deleted top-up response.
func (t *topupCommandResponseMapper) ToApiResponseTopupDeleteAt(s *pb.ApiResponseTopupDeleteAt) *response.ApiResponseTopupDeleteAt {
	return &response.ApiResponseTopupDeleteAt{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponseTopupDeleteAt(s.Data),
	}
}

// ToApiResponseTopupAll maps a gRPC response containing all top-up records to an HTTP API response.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupAll containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupAll containing the mapped data, including status and message.
func (t *topupCommandResponseMapper) ToApiResponseTopupAll(s *pb.ApiResponseTopupAll) *response.ApiResponseTopupAll {
	return &response.ApiResponseTopupAll{
		Status:  s.Status,
		Message: s.Message,
	}
}

// ToApiResponseTopupDelete maps a single gRPC top-up delete response to an HTTP API response.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupDelete containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupDelete containing the mapped data, including status and message.
func (t *topupCommandResponseMapper) ToApiResponseTopupDelete(s *pb.ApiResponseTopupDelete) *response.ApiResponseTopupDelete {
	return &response.ApiResponseTopupDelete{
		Status:  s.Status,
		Message: s.Message,
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
func (t *topupCommandResponseMapper) mapResponseTopup(topup *pb.TopupResponse) *response.TopupResponse {
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

// mapResponseTopupDeleteAt maps a gRPC soft-deleted top-up response to an HTTP API response format.
// It constructs a TopupResponseDeleteAt by copying the fields from the gRPC response, converting
// the ID and TopupAmount fields to integers, and handling the potentially nil DeletedAt field.
//
// Args:
//   - topup: A pointer to a pb.TopupResponseDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TopupResponseDeleteAt with the mapped data, including a non-nil DeletedAt field if present.
func (t *topupCommandResponseMapper) mapResponseTopupDeleteAt(topup *pb.TopupResponseDeleteAt) *response.TopupResponseDeleteAt {
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
