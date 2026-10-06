package requests

import "github.com/go-playground/validator/v10"

// CreateRefreshToken represents the payload for generating a new refresh token.
// Used when creating persistent authentication tokens for users.
type CreateRefreshToken struct {
	UserId    int    `json:"user_id" validate:"required,min=1"`    // ID of the user the token belongs to
	Token     string `json:"token" validate:"required,min=1"`      // The actual refresh token value (should be hashed)
	ExpiresAt string `json:"expires_at" validate:"required,min=1"` // Expiration timestamp (RFC3339 format recommended)
}

// UpdateRefreshToken represents the payload for updating an existing refresh token.
// Used when rotating or extending refresh token validity.
type UpdateRefreshToken struct {
	UserId    int    `json:"user_id" validate:"required,min=1"`    // ID of the user the token belongs to
	Token     string `json:"token" validate:"required,min=1"`      // The new refresh token value (should be hashed)
	ExpiresAt string `json:"expires_at" validate:"required,min=1"` // New expiration timestamp (RFC3339 format recommended)
}

// RefreshTokenRequest represents the payload for refresh token operations.
// Used when exchanging a refresh token for new access credentials.
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required,min=1"` // The refresh token to validate/redeem
}

// Validate performs structural validation of CreateRefreshToken fields.
// Checks that all required fields meet minimum requirements.
// Returns:
//   - error: if validation fails according to struct tag rules
func (r *CreateRefreshToken) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

// Validate performs structural validation of UpdateRefreshToken fields.
// Checks that all required fields meet minimum requirements.
// Returns:
//   - error: if validation fails according to struct tag rules
func (r *UpdateRefreshToken) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

// Validate performs structural validation of RefreshTokenRequest fields.
// Checks that the refresh token is present and meets minimum requirements.
// Returns:
//   - error: if validation fails according to struct tag rules
func (r *RefreshTokenRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}
