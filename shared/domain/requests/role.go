package requests

import (
	"github.com/go-playground/validator/v10"
)

// RoleRequestPayload represents the base payload structure for role management requests.
// Contains common fields used across role-related operations.
type RoleRequestPayload struct {
	UserID        int    `json:"user_id"`        // ID of the user performing the role operation
	CorrelationID string `json:"correlation_id"` // Unique identifier for request tracing
	ReplyTopic    string `json:"reply_topic"`    // Topic name for asynchronous response delivery
}

// FindAllRoles represents parameters for searching and paginating role records.
// Used in role administration and listing operations.
type FindAllRoles struct {
	Search   string `json:"search" validate:"required"`         // Search term to filter roles (matches against role name)
	Page     int    `json:"page" validate:"min=1"`              // Page number for pagination (1-based index)
	PageSize int    `json:"page_size" validate:"min=1,max=100"` // Number of items per page (1-100)
}

// CreateRoleRequest represents the payload for creating a new role.
// Used when defining new access control roles in the system.
type CreateRoleRequest struct {
	Name string `json:"name" validate:"required"` // Name of the new role (e.g., "admin", "user")
}

// UpdateRoleRequest represents the payload for modifying an existing role.
// Used when renaming or updating role definitions.
type UpdateRoleRequest struct {
	ID   *int   `json:"id"`                       // ID of the role to update (optional in some flows)
	Name string `json:"name" validate:"required"` // New name for the role
}

// Validate performs validation of CreateRoleRequest fields.
// Ensures the role name is provided and meets requirements.
// Returns:
//   - error: if validation fails, describing the validation failure
func (r *CreateRoleRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

// Validate performs validation of UpdateRoleRequest fields.
// Ensures the role ID (when provided) and name meet requirements.
// Returns:
//   - error: if validation fails, describing the validation failure
func (r *UpdateRoleRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}
