package response

// UserResponse represents a user account in the system
type UserResponse struct {
	ID         int    `json:"id"`          // Unique user identifier
	FirstName  string `json:"firstname"`   // User's first name
	LastName   string `json:"lastname"`    // User's last name
	Email      string `json:"email"`       // User's email address (used for authentication)
	IsVerified bool   `json:"is_verified"` // Whether the user's email has been verified
	CreatedAt  string `json:"created_at"`  // Timestamp when the user account was created
	UpdatedAt  string `json:"updated_at"`  // Timestamp when the user account was last updated
}

// UserResponseDeleteAt extends UserResponse with soft delete capability
type UserResponseDeleteAt struct {
	DeletedAt *string `json:"deleted_at"` // Timestamp when user was soft deleted (nil if active)
	// Embedded UserResponse fields
	ID         int    `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

// ApiResponseUser is the API response for a single user
type ApiResponseUser struct {
	Status  string        `json:"status"`  // Response status ("success" or "error")
	Message string        `json:"message"` // Descriptive response message
	Data    *UserResponse `json:"data"`    // Contains the user data
}

// ApiResponseUserDeleteAt is the API response for a single user including deletion info
type ApiResponseUserDeleteAt struct {
	Status  string                `json:"status"`  // Response status
	Message string                `json:"message"` // Response message
	Data    *UserResponseDeleteAt `json:"data"`    // User data including deletion timestamp
}

// ApiResponsesUser is the API response for multiple users
type ApiResponsesUser struct {
	Status  string          `json:"status"`  // Response status
	Message string          `json:"message"` // Response message
	Data    []*UserResponse `json:"data"`    // Array of user data
}

// ApiResponseUserDelete is the API response for user deletion
type ApiResponseUserDelete struct {
	Status  string `json:"status"`  // Response status
	Message string `json:"message"` // Operation result message
}

// ApiResponseUserAll is a generic API response for user operations
type ApiResponseUserAll struct {
	Status  string `json:"status"`  // Response status
	Message string `json:"message"` // Operation message
}

// ApiResponsePaginationUserDeleteAt is a paginated API response for soft-deleted users
type ApiResponsePaginationUserDeleteAt struct {
	Status     string                  `json:"status"`     // Response status
	Message    string                  `json:"message"`    // Response message
	Data       []*UserResponseDeleteAt `json:"data"`       // Array of user data with deletion info
	Pagination *PaginationMeta         `json:"pagination"` // Pagination metadata
}

// ApiResponsePaginationUser is a paginated API response for active users
type ApiResponsePaginationUser struct {
	Status     string          `json:"status"`     // Response status
	Message    string          `json:"message"`    // Response message
	Data       []*UserResponse `json:"data"`       // Array of user data
	Pagination *PaginationMeta `json:"pagination"` // Pagination metadata
}
