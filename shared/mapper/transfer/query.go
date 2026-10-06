package transferapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transfer"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper"
)

type transferQueryResponseMapper struct {
}

func NewTransferQueryResponseMapper() TransferQueryResponseMapper {
	return &transferQueryResponseMapper{}
}

// ToApiResponseTransfer maps a single gRPC transfer response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransfer containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransfer containing the mapped data, including status, message, and a single mapped transfer response.
func (m *transferQueryResponseMapper) ToApiResponseTransfer(pbResponse *pb.ApiResponseTransfer) *response.ApiResponseTransfer {
	return &response.ApiResponseTransfer{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransfer(pbResponse.Data),
	}
}

// ToApiResponseTransfers maps a gRPC response containing multiple transfer records into an HTTP API response.
// It constructs an ApiResponseTransfers by copying the status and message fields, mapping the Transfer data slice to a slice of
// TransferResponse, and assigning it to the response's Data field.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransfers containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransfers with mapped data.
func (m *transferQueryResponseMapper) ToApiResponseTransfers(pbResponse *pb.ApiResponseTransfers) *response.ApiResponseTransfers {
	return &response.ApiResponseTransfers{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesTransfer(pbResponse.Data),
	}
}

// ToApiResponsePaginationTransfer maps a pagination meta, status, message, and a list of TransferResponse
// to a response.ApiResponsePaginationTransfer proto message.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponsePaginationTransfer containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationTransfer containing the mapped data.
func (m *transferQueryResponseMapper) ToApiResponsePaginationTransfer(pbResponse *pb.ApiResponsePaginationTransfer) *response.ApiResponsePaginationTransfer {
	return &response.ApiResponsePaginationTransfer{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       m.mapResponsesTransfer(pbResponse.Data),
		Pagination: apimapper.MapPaginationMeta(pbResponse.PaginationMeta),
	}
}

// ToApiResponsePaginationTransferDeleteAt maps a pagination meta, status, message, and a list of TransferResponseDeleteAt
// to a response.ApiResponsePaginationTransferDeleteAt proto message.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponsePaginationTransferDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationTransferDeleteAt containing the mapped data.
func (m *transferQueryResponseMapper) ToApiResponsePaginationTransferDeleteAt(pbResponse *pb.ApiResponsePaginationTransferDeleteAt) *response.ApiResponsePaginationTransferDeleteAt {
	return &response.ApiResponsePaginationTransferDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       m.mapResponsesTransferDeleteAt(pbResponse.Data),
		Pagination: apimapper.MapPaginationMeta(pbResponse.PaginationMeta),
	}
}

// mapResponseTransfer maps a TransferResponse from a protobuf message to a response.TransferResponse.
//
// Args:
//   - transfer: A pointer to a pb.TransferResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponse containing the mapped data.
func (t *transferQueryResponseMapper) mapResponseTransfer(transfer *pb.TransferResponse) *response.TransferResponse {
	return &response.TransferResponse{
		ID:             int(transfer.Id),
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
	}
}

// mapResponsesTransfer maps a slice of TransferResponse from protobuf messages to a slice of response.TransferResponse.
//
// Args:
//   - transfers: A slice of pointers to pb.TransferResponse containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferResponse containing the mapped data.
func (t *transferQueryResponseMapper) mapResponsesTransfer(transfers []*pb.TransferResponse) []*response.TransferResponse {
	var responses []*response.TransferResponse

	for _, response := range transfers {
		responses = append(responses, t.mapResponseTransfer(response))
	}

	return responses
}

// mapResponseTransferDeleteAt maps a TransferResponseDeleteAt from a protobuf message to a response.TransferResponseDeleteAt.
//
// Args:
//   - transfer: A pointer to a pb.TransferResponseDeleteAt containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponseDeleteAt containing the mapped data.
func (t *transferQueryResponseMapper) mapResponseTransferDeleteAt(transfer *pb.TransferResponseDeleteAt) *response.TransferResponseDeleteAt {
	var deletedAt string
	if transfer.DeletedAt != nil {
		deletedAt = transfer.DeletedAt.Value
	}

	return &response.TransferResponseDeleteAt{
		ID:             int(transfer.Id),
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
		DeletedAt:      &deletedAt,
	}
}

// mapResponsesTransferDeleteAt maps a slice of TransferResponseDeleteAt from protobuf messages to a slice of response.TransferResponseDeleteAt.
//
// It iterates over the input slice, converting each TransferResponseDeleteAt to its corresponding response.TransferResponseDeleteAt
// using the mapResponseTransferDeleteAt method.
//
// Args:
//   - transfers: A slice of pointers to pb.TransferResponseDeleteAt containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferResponseDeleteAt containing the mapped data.
func (t *transferQueryResponseMapper) mapResponsesTransferDeleteAt(transfers []*pb.TransferResponseDeleteAt) []*response.TransferResponseDeleteAt {
	var responses []*response.TransferResponseDeleteAt

	for _, response := range transfers {
		responses = append(responses, t.mapResponseTransferDeleteAt(response))
	}

	return responses
}
