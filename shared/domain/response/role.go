package response

// RoleResponsePayload represents the base payload structure for role validation responses.
// Used to verify role assignments and permissions.
type RoleResponsePayload struct {
	CorrelationID string   `json:"correlation_id"` // Unique ID for request tracing
	Valid         bool     `json:"valid"`          // Indicates if role validation succeeded
	RoleNames     []string `json:"role_names"`     // List of role names associated with the validation
}

// RoleResponse represents basic role information in API responses.
// Contains core role data without deletion information.
type RoleResponse struct {
	ID        int    `json:"id"`         // Unique role identifier
	Name      string `json:"name"`       // Name of the role (e.g., "admin", "user")
	CreatedAt string `json:"created_at"` // Timestamp when role was created (RFC3339 format)
	UpdatedAt string `json:"updated_at"` // Timestamp when role was last updated
}

// RoleResponseDeleteAt extends RoleResponse with deletion information.
// Used in administrative interfaces showing soft-deleted roles.
type RoleResponseDeleteAt struct {
	ID        int     `json:"id"`         // Unique role identifier
	Name      string  `json:"name"`       // Name of the role
	CreatedAt string  `json:"created_at"` // Original creation timestamp
	UpdatedAt string  `json:"updated_at"` // Last update timestamp
	DeletedAt *string `json:"deleted_at"` // Deletion timestamp (nil if not deleted)
}

// API Response Wrappers:

// ApiResponseRoleAll is the response format for bulk role operations.
type ApiResponseRoleAll struct {
	Status  string `json:"status"`  // Operation status ("success" or "error")
	Message string `json:"message"` // Result message
}

// ApiResponseRoleDelete is the response format for role deletion operations.
type ApiResponseRoleDelete struct {
	Status  string `json:"status"`  // Operation status
	Message string `json:"message"` // Result message
}

// ApiResponseRole is the standard response format for single role requests.
type ApiResponseRole struct {
	Status  string        `json:"status"`  // Response status
	Message string        `json:"message"` // Descriptive message
	Data    *RoleResponse `json:"data"`    // Role data payload
}

type ApiResponseRoleDeleteAt struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Data    *RoleResponseDeleteAt `json:"data"`
}

// ApiResponsesRole is the response format for multiple role listings (non-paginated).
type ApiResponsesRole struct {
	Status  string          `json:"status"`  // Response status
	Message string          `json:"message"` // Descriptive message
	Data    []*RoleResponse `json:"data"`    // Array of role data
}

// ApiResponsePaginationRole is the paginated response for role lists.
type ApiResponsePaginationRole struct {
	Status     string          `json:"status"`     // Response status
	Message    string          `json:"message"`    // Descriptive message
	Data       []*RoleResponse `json:"data"`       // Array of role data
	Pagination *PaginationMeta `json:"pagination"` // Pagination metadata
}

// ApiResponsePaginationRoleDeleteAt is the paginated response including deleted roles.
type ApiResponsePaginationRoleDeleteAt struct {
	Status     string                  `json:"status"`     // Response status
	Message    string                  `json:"message"`    // Descriptive message
	Data       []*RoleResponseDeleteAt `json:"data"`       // Array of role data (with deletion info)
	Pagination *PaginationMeta         `json:"pagination"` // Pagination metadata
}
