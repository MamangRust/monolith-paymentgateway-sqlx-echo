package response

// TokenResponse represents a pair of authentication tokens.
// Used in login and token refresh operations.
type TokenResponse struct {
	AccessToken  string `json:"access_token"`  // Short-lived JWT for API access
	RefreshToken string `json:"refresh_token"` // Long-lived token for obtaining new access tokens
}

// API Response Wrappers:

// ApiResponseVerifyCode is the response format for verification code operations.
// Used in email/SMS verification flows.
type ApiResponseVerifyCode struct {
	Status  string `json:"status"`  // Operation status ("success" or "error")
	Message string `json:"message"` // Result message ("Verification successful")
}

// ApiResponseForgotPassword is the response format for password reset initiation.
// Used when requesting a password reset link.
type ApiResponseForgotPassword struct {
	Status  string `json:"status"`  // Operation status
	Message string `json:"message"` // Result message ("Reset link sent to email")
}

// ApiResponseResetPassword is the response format for password reset completion.
// Used after submitting a new password.
type ApiResponseResetPassword struct {
	Status  string `json:"status"`  // Operation status
	Message string `json:"message"` // Result message ("Password updated successfully")
}

// ApiResponseLogin is the response format for successful login attempts.
// Contains the authentication tokens.
type ApiResponseLogin struct {
	Status  string         `json:"status"`  // Always "success" for this response
	Message string         `json:"message"` // Welcome message
	Data    *TokenResponse `json:"data"`    // Authentication tokens
}

// ApiResponseRegister is the response format for new user registration.
// Contains the created user's information.
type ApiResponseRegister struct {
	Status  string        `json:"status"`  // Operation status
	Message string        `json:"message"` // Welcome/confirmation message
	Data    *UserResponse `json:"data"`    // Registered user's profile data
}

// ApiResponseRefreshToken is the response format for token refresh operations.
// Contains new authentication tokens.
type ApiResponseRefreshToken struct {
	Status  string         `json:"status"`  // Operation status
	Message string         `json:"message"` // Result message
	Data    *TokenResponse `json:"data"`    // New set of authentication tokens
}

// ApiResponseGetMe is the response format for current user profile requests.
// Contains the authenticated user's information.
type ApiResponseGetMe struct {
	Status  string        `json:"status"`  // Always "success" for valid requests
	Message string        `json:"message"` // Descriptive message
	Data    *UserResponse `json:"data"`    // Complete user profile data
}
