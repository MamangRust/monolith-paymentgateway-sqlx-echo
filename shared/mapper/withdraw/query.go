package withdrawapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper"
)

// withdrawQueryResponseMapper provides methods to map gRPC withdraw responses to HTTP API responses
type withdrawQueryResponseMapper struct {
}

// NewWithdrawResponseMapper creates and returns a new instance of withdrawQueryResponseMapper.
// This instance provides methods to map gRPC withdraw responses to HTTP API responses.
func NewWithdrawQueryResponseMapper() WithdrawQueryResponseMapper {
	return &withdrawQueryResponseMapper{}
}

// ToApiResponseWithdraw maps a single withdraw gRPC response to an API response.
//
// Args:
//   - pbResponse: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to an ApiResponseWithdraw containing the mapped data.
func (m *withdrawQueryResponseMapper) ToApiResponseWithdraw(pbResponse *pb.ApiResponseWithdraw) *response.ApiResponseWithdraw {
	return &response.ApiResponseWithdraw{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseWithdrawal(pbResponse.Data),
	}
}

// ToApiResponsesWithdraw maps a list of withdraw gRPC responses to an API response.
//
// Args:
//   - pbResponse: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to an ApiResponsesWithdraw containing the mapped data.
func (m *withdrawQueryResponseMapper) ToApiResponsesWithdraw(pbResponse *pb.ApiResponsesWithdraw) *response.ApiResponsesWithdraw {
	return &response.ApiResponsesWithdraw{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesWithdrawal(pbResponse.Data),
	}
}

// ToApiResponsePaginationWithdraw maps a pagination meta, status, message, and a list of WithdrawResponse
// to a response.ApiResponsePaginationWithdraw proto message.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponsePaginationWithdraw containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationWithdraw containing the mapped data.
func (m *withdrawQueryResponseMapper) ToApiResponsePaginationWithdraw(pbResponse *pb.ApiResponsePaginationWithdraw) *response.ApiResponsePaginationWithdraw {
	return &response.ApiResponsePaginationWithdraw{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       m.mapResponsesWithdrawal(pbResponse.Data),
		Pagination: apimapper.MapPaginationMeta(pbResponse.PaginationMeta),
	}
}

// ToApiResponsePaginationWithdrawDeleteAt maps a pagination meta, status, message, and a list of WithdrawResponseDeleteAt
// to a response.ApiResponsePaginationWithdrawDeleteAt proto message.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponsePaginationWithdrawDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationWithdrawDeleteAt containing the mapped data.
func (m *withdrawQueryResponseMapper) ToApiResponsePaginationWithdrawDeleteAt(pbResponse *pb.ApiResponsePaginationWithdrawDeleteAt) *response.ApiResponsePaginationWithdrawDeleteAt {
	return &response.ApiResponsePaginationWithdrawDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       m.mapResponsesWithdrawalDeleteAt(pbResponse.Data),
		Pagination: apimapper.MapPaginationMeta(pbResponse.PaginationMeta),
	}
}

// mapResponseWithdrawal maps a single withdraw gRPC response to an API response.
//
// Args:
//   - withdraw: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponse containing the mapped data.
func (w *withdrawQueryResponseMapper) mapResponseWithdrawal(withdraw *pb.WithdrawResponse) *response.WithdrawResponse {
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

// mapResponsesWithdrawal maps a slice of WithdrawResponse to a slice of WithdrawResponse.
//
// It takes a slice of WithdrawResponse as input and returns a slice of corresponding WithdrawResponse.
// The mapping includes fields like WithdrawId, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, and UpdatedAt.
func (w *withdrawQueryResponseMapper) mapResponsesWithdrawal(withdraws []*pb.WithdrawResponse) []*response.WithdrawResponse {
	var responseWithdraws []*response.WithdrawResponse

	for _, withdraw := range withdraws {
		responseWithdraws = append(responseWithdraws, w.mapResponseWithdrawal(withdraw))
	}

	return responseWithdraws
}

// mapResponseWithdrawalDeleteAt maps a single WithdrawResponseDeleteAt to an API response.
//
// Args:
//   - withdraw: The WithdrawResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseDeleteAt containing the mapped data, including id, withdraw_no, card_number,
//     withdraw_amount, withdraw_time, created_at, updated_at, and deleted_at.
func (w *withdrawQueryResponseMapper) mapResponseWithdrawalDeleteAt(withdraw *pb.WithdrawResponseDeleteAt) *response.WithdrawResponseDeleteAt {
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

// mapResponsesWithdrawalDeleteAt maps a slice of WithdrawResponseDeleteAt gRPC messages
// to a slice of API response WithdrawResponseDeleteAt objects.
//
// Args:
//   - withdraws: A slice of WithdrawResponseDeleteAt gRPC messages that need to be converted.
//
// Returns:
//   - A slice of WithdrawResponseDeleteAt API responses containing the mapped data.
func (w *withdrawQueryResponseMapper) mapResponsesWithdrawalDeleteAt(withdraws []*pb.WithdrawResponseDeleteAt) []*response.WithdrawResponseDeleteAt {
	var responseWithdraws []*response.WithdrawResponseDeleteAt

	for _, withdraw := range withdraws {
		responseWithdraws = append(responseWithdraws, w.mapResponseWithdrawalDeleteAt(withdraw))
	}

	return responseWithdraws
}
