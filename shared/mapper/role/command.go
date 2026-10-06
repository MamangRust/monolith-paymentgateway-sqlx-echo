package roleapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/role"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type roleCommandResponseMapper struct {
}

// NewRoleCommandResponseMapper returns a new instance of roleCommandResponseMapper, which provides methods to map RoleRecord domain models to RoleResponse API-compatible response types for command operations.
func NewRoleCommandResponseMapper() RoleCommandResponseMapper {
	return &roleCommandResponseMapper{}
}

// ToApiResponseRole maps a single gRPC role response into an HTTP API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseRole containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseRole containing the mapped status, message, and data.
func (s *roleCommandResponseMapper) ToApiResponseRole(pbResponse *pb.ApiResponseRole) *response.ApiResponseRole {
	return &response.ApiResponseRole{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseRole(pbResponse.Data),
	}
}

// ToApiResponseRole maps a single gRPC role response into an HTTP API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseRole containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseRole containing the mapped status, message, and data.
func (s *roleCommandResponseMapper) ToApiResponseRoleDeleteAt(pbResponse *pb.ApiResponseRoleDeleteAt) *response.ApiResponseRoleDeleteAt {
	return &response.ApiResponseRoleDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseRoleDeleteAt(pbResponse.Data),
	}
}

// ToApiResponseRoleDelete maps a gRPC delete role response to an HTTP API response format.
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseRoleDelete containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseRoleDelete containing the mapped status and message.
func (s *roleCommandResponseMapper) ToApiResponseRoleDelete(pbResponse *pb.ApiResponseRoleDelete) *response.ApiResponseRoleDelete {
	return &response.ApiResponseRoleDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponseRoleAll maps a gRPC response containing all roles to an HTTP API response format.
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseRoleAll containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseRoleAll containing the mapped data, including status and message.
func (s *roleCommandResponseMapper) ToApiResponseRoleAll(pbResponse *pb.ApiResponseRoleAll) *response.ApiResponseRoleAll {
	return &response.ApiResponseRoleAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// mapResponseRole maps a gRPC RoleResponse to an HTTP-compatible RoleResponse.
//
// Args:
//   - role: A pointer to a pb.RoleResponse containing the gRPC role response data.
//
// Returns:
//   - A pointer to a response.RoleResponse containing the mapped data, with fields ID, Name,
//     CreatedAt, and UpdatedAt extracted from the gRPC response.
func (s *roleCommandResponseMapper) mapResponseRole(role *pb.RoleResponse) *response.RoleResponse {
	return &response.RoleResponse{
		ID:        int(role.Id),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

// mapResponseRoleDeleteAt maps a gRPC RoleResponseDeleteAt to an HTTP-compatible RoleResponseDeleteAt.
//
// Args:
//   - role: A pointer to a pb.RoleResponseDeleteAt containing the gRPC role response data including deletion info.
//
// Returns:
//   - A pointer to a response.RoleResponseDeleteAt containing the mapped data, with fields ID, Name,
//     CreatedAt, UpdatedAt, and DeletedAt extracted from the gRPC response. If DeletedAt is present,
//     it is mapped to a string; otherwise, it is set to nil.
func (s *roleCommandResponseMapper) mapResponseRoleDeleteAt(role *pb.RoleResponseDeleteAt) *response.RoleResponseDeleteAt {
	var deletedAt string
	if role.DeletedAt != nil {
		deletedAt = role.DeletedAt.Value
	}

	return &response.RoleResponseDeleteAt{
		ID:        int(role.Id),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
		DeletedAt: &deletedAt,
	}
}
