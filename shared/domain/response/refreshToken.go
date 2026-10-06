package response

// RefreshTokenResponse represents a refresh token record in API responses.
// Contains the token details along with timestamps for tracking.
type RefreshTokenResponse struct {
	UserID    int    `json:"user_id"`    // ID of the user this token belongs to
	Token     string `json:"token"`      // The refresh token value (should be hashed in storage)
	ExpiredAt string `json:"expired_at"` // Expiration timestamp (RFC3339 format)
	CreatedAt string `json:"created_at"` // Timestamp when token was created
	UpdatedAt string `json:"updated_at"` // Timestamp when token was last updated
}
