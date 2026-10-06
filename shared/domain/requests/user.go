package requests

import (
	"github.com/go-playground/validator/v10"
)

// FindAllUsers represents parameters for searching and paginating user records.
// Used in user administration interfaces to list and filter users.
type FindAllUsers struct {
	Search   string `json:"search" validate:"required"`         // Search term to filter users (matches name or email)
	Page     int    `json:"page" validate:"min=1"`              // Page number for pagination (1-based index)
	PageSize int    `json:"page_size" validate:"min=1,max=100"` // Number of items per page (1-100)
}

// CreateUserRequest represents the payload for registering a new user.
// Contains all required information for user account creation.
type CreateUserRequest struct {
	FirstName       string `json:"firstname" validate:"required,alpha"`                   // User's first name (alphabetic characters only)
	LastName        string `json:"lastname" validate:"required,alpha"`                    // User's last name (alphabetic characters only)
	Email           string `json:"email" validate:"required,email"`                       // User's email address (must be valid format)
	Password        string `json:"password" validate:"required,min=6"`                    // Account password (minimum 6 characters)
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"` // Password confirmation (must match Password field)
}

// UpdateUserRequest represents the payload for modifying an existing user's information.
// Used when updating user profile details or credentials.
type UpdateUserRequest struct {
	UserID          *int   `json:"user_id"`                                               // ID of user to update (optional in some flows)
	FirstName       string `json:"firstname" validate:"required,alpha"`                   // Updated first name
	LastName        string `json:"lastname" validate:"required,alpha"`                    // Updated last name
	Email           string `json:"email" validate:"required,email"`                       // Updated email address
	Password        string `json:"password" validate:"required,min=6"`                    // New password
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"` // Password confirmation
}

// Validate performs structural validation of CreateUserRequest fields.
// Ensures all required fields are present and meet format requirements.
// Returns:
//   - error: if validation fails according to struct tag rules
func (r *CreateUserRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

// Validate performs structural validation of UpdateUserRequest fields.
// Ensures all required fields are present and meet format requirements.
// Returns:
//   - error: if validation fails according to struct tag rules
func (r *UpdateUserRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}
