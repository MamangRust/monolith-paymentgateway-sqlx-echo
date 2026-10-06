package transferapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transfer"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transferCommandResponseMapper struct{}

func NewTransferCommandResponseMapper() TransferCommandResponseMapper {
	return &transferCommandResponseMapper{}
}

// ToApiResponseTransfer maps a single gRPC transfer response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransfer containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransfer containing the mapped data, including status, message, and a single mapped transfer response.
func (m *transferCommandResponseMapper) ToApiResponseTransfer(pbResponse *pb.ApiResponseTransfer) *response.ApiResponseTransfer {
	return &response.ApiResponseTransfer{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransfer(pbResponse.Data),
	}
}

func (m *transferCommandResponseMapper) ToApiResponseTransferDeleteAt(pbResponse *pb.ApIResponseTransferDeleteAt) *response.ApiResponseTransferDeleteAt {
	return &response.ApiResponseTransferDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransferDeleteAt(pbResponse.Data),
	}
}

// ToApiResponseTransferDelete maps a gRPC transfer delete response to an HTTP API response.
// It constructs an ApiResponseTransferDelete by copying the status and message fields from
// the gRPC response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferDelete containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferDelete with the mapped status and message.
func (m *transferCommandResponseMapper) ToApiResponseTransferDelete(pbResponse *pb.ApiResponseTransferDelete) *response.ApiResponseTransferDelete {
	return &response.ApiResponseTransferDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponseTransferAll maps a gRPC response containing all transfer records into an HTTP API response.
// It constructs an ApiResponseTransferAll by copying the status and message fields from the gRPC response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferAll containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferAll with the mapped status and message.
func (m *transferCommandResponseMapper) ToApiResponseTransferAll(pbResponse *pb.ApiResponseTransferAll) *response.ApiResponseTransferAll {
	return &response.ApiResponseTransferAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// mapResponseTransfer maps a TransferResponse from a protobuf message to a response.TransferResponse.
//
// Args:
//   - transfer: A pointer to a pb.TransferResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponse containing the mapped data.
func (t *transferCommandResponseMapper) mapResponseTransfer(transfer *pb.TransferResponse) *response.TransferResponse {
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

// mapResponseTransferDeleteAt maps a TransferResponseDeleteAt from a protobuf message to a response.TransferResponseDeleteAt.
//
// Args:
//   - transfer: A pointer to a pb.TransferResponseDeleteAt containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponseDeleteAt containing the mapped data.
func (t *transferCommandResponseMapper) mapResponseTransferDeleteAt(transfer *pb.TransferResponseDeleteAt) *response.TransferResponseDeleteAt {
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
