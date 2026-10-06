package refreshtokenserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedParseExpirationDate indicates failure when parsing the expiration date of a token.
var ErrFailedParseExpirationDate = errors.NewErrorResponse("Failed to parse expiration date", http.StatusBadRequest)

// ErrFailedInValidToken is returned when an access token is invalid.
var ErrFailedInvalidToken = errors.NewErrorResponse("Failed to invalid access token", http.StatusInternalServerError)

// ErrFailedInValidUserId is returned when a user ID is invalid.
var ErrFailedInvalidUserId = errors.NewErrorResponse("Failed to invalid user id", http.StatusInternalServerError)

// ErrFailedExpire occurs when expiring a refresh token fails.
var ErrFailedExpire = errors.NewErrorResponse("Failed to find refresh token by token", http.StatusInternalServerError)
