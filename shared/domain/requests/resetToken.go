package requests

import "github.com/go-playground/validator/v10"

// CreateResetTokenRequest represents the payload for generating a password reset token.
// Used when initiating a password reset process for a user.
type CreateResetTokenRequest struct {
	UserID     int    `json:"user_id" validate:"required"`     // ID of the user requesting password reset
	ResetToken string `json:"reset_token" validate:"required"` // The generated reset token (should be hashed before storage)
	ExpiredAt  string `json:"expired_at" validate:"required"`  // Expiration time of the token (RFC3339 format recommended)
}

// CreateResetPasswordRequest represents the payload for actually resetting a user's password.
// Used when submitting a new password during the reset flow.
type CreateResetPasswordRequest struct {
	ResetToken      string `json:"reset_token" validate:"required"`            // The reset token provided to the user
	Password        string `json:"password" validate:"required,min=6"`         // New password (minimum 6 characters)
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6"` // Password confirmation (must match new password)
}

// ForgotPasswordRequest represents the initial forgot password request payload.
// Used when a user initiates the password reset process via email.
type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"` // Email address associated with the account
}

// Validate performs validation of CreateResetPasswordRequest fields.
// Ensures the reset token is present and passwords meet complexity requirements.
// Returns:
//   - error: if validation fails, describing the validation failure
func (r *CreateResetPasswordRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

// Validate performs validation of ForgotPasswordRequest fields.
// Ensures the email field is present and properly formatted.
// Returns:
//   - error: if validation fails, describing the validation failure
func (r *ForgotPasswordRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

// ResetPasswordRequest represents the payload for resetting a user's password.
type ResetPasswordRequest struct {
	ResetToken      string `json:"reset_token" validate:"required"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6"`
}

// Validate performs validation of ResetPasswordRequest fields.
func (r *ResetPasswordRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}
