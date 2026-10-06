package requests

import "github.com/go-playground/validator/v10"

// AuthRequest represents the payload for user authentication (login).
// Used when a user attempts to log in to the system.
type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`    // User's email address (must be valid email format)
	Password string `json:"password" validate:"required,min=6"` // User's password (minimum 6 characters)
}

// RegisterRequest represents the payload for user registration.
// Contains all necessary information to create a new user account.
type RegisterRequest struct {
	FirstName       string `json:"firstname"`                                  // User's first name
	LastName        string `json:"lastname"`                                   // User's last name
	Email           string `json:"email" validate:"required,email"`            // User's email address (must be valid email format)
	Password        string `json:"password" validate:"required,min=6"`         // User's password (minimum 6 characters)
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6"` // Password confirmation (must match password)
	VerifiedCode    string `json:"verified_code"`                              // Verification code for email confirmation (optional during registration)
	IsVerified      bool   `json:"is_verified"`                                // Flag indicating if user has verified their email (typically false initially)
}

// Validate validates the AuthRequest struct using the validator.v10 package.
//
// Returns an error if any of the validation rules fail.
func (r *AuthRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

// Validate validates the RegisterRequest struct using the validator.v10 package.
//
// Returns an error if any of the validation rules fail.
func (r *RegisterRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

type JWTToken struct {
	Token string `json:"token"`
}

type VerifyCodeRequest struct {
	Code string `json:"code" validate:"required"`
}

func (r *VerifyCodeRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}
