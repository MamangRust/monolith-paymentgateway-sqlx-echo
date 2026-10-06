package withdrawapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// WithdrawCommandResponseMapper provides methods to map gRPC withdraw responses to HTTP API responses
type withdrawCommandResponseMapper struct {
}

// NewWithdrawCommandResponseMapper creates and returns a new instance of withdrawCommandResponseMapper.
// This mapper provides methods to convert withdraw gRPC responses into API-compatible responses for command operations.
func NewWithdrawCommandResponseMapper() WithdrawCommandResponseMapper {
	return &withdrawCommandResponseMapper{}
}

// ToApiResponseWithdraw maps a single withdraw gRPC response to an API response.
//
// Args:
//   - pbResponse: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to an ApiResponseWithdraw containing the mapped data.
func (m *withdrawCommandResponseMapper) ToApiResponseWithdraw(pbResponse *pb.ApiResponseWithdraw) *response.ApiResponseWithdraw {
	return &response.ApiResponseWithdraw{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseWithdrawal(pbResponse.Data),
	}
}

// ToApiResponseWithdrawDeleteAt maps a single withdraw gRPC response to an API response.
//
// Args:
//   - pbResponse: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to an ApiResponseWithdrawDeleteAt containing the mapped data.
func (m *withdrawCommandResponseMapper) ToApiResponseWithdrawDeleteAt(pbResponse *pb.ApIResponseWithdrawDeleteAt) *response.ApiResponseWithdrawDeleteAt {
	return &response.ApiResponseWithdrawDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseWithdrawalDeleteAt(pbResponse.Data),
	}
}

// ToApiResponseWithdrawDelete maps a gRPC response indicating a withdraw has been deleted to an API response.
//
// Args:
//   - pbResponse: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to an ApiResponseWithdrawDelete containing the mapped data.
func (m *withdrawCommandResponseMapper) ToApiResponseWithdrawDelete(pbResponse *pb.ApiResponseWithdrawDelete) *response.ApiResponseWithdrawDelete {
	return &response.ApiResponseWithdrawDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponseWithdrawAll maps a gRPC response containing all withdraw records to an API response.
//
// Args:
//   - pbResponse: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to an ApiResponseWithdrawAll containing the mapped data.
func (m *withdrawCommandResponseMapper) ToApiResponseWithdrawAll(pbResponse *pb.ApiResponseWithdrawAll) *response.ApiResponseWithdrawAll {
	return &response.ApiResponseWithdrawAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// mapResponseWithdrawal maps a single withdraw gRPC response to an API response.
//
// Args:
//   - withdraw: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponse containing the mapped data.
func (w *withdrawCommandResponseMapper) mapResponseWithdrawal(withdraw *pb.WithdrawResponse) *response.WithdrawResponse {
	return &response.WithdrawResponse{
		ID:             int(withdraw.WithdrawId),
		WithdrawNo:     withdraw.WithdrawNo,
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime,
		CreatedAt:      withdraw.CreatedAt,
		UpdatedAt:      withdraw.UpdatedAt,
	}
}

// mapResponseWithdrawalDeleteAt maps a single WithdrawResponseDeleteAt to an API response.
//
// Args:
//   - withdraw: The WithdrawResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseDeleteAt containing the mapped data, including id, withdraw_no, card_number,
//     withdraw_amount, withdraw_time, created_at, updated_at, and deleted_at.
func (w *withdrawCommandResponseMapper) mapResponseWithdrawalDeleteAt(withdraw *pb.WithdrawResponseDeleteAt) *response.WithdrawResponseDeleteAt {
	var deletedAt string
	if withdraw.DeletedAt != nil {
		deletedAt = withdraw.DeletedAt.Value
	}

	return &response.WithdrawResponseDeleteAt{
		ID:             int(withdraw.WithdrawId),
		WithdrawNo:     withdraw.WithdrawNo,
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime,
		CreatedAt:      withdraw.CreatedAt,
		UpdatedAt:      withdraw.UpdatedAt,
		DeletedAt:      &deletedAt,
	}
}
