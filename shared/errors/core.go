package errors

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type ErrorType string

const (
	ErrorTypeInternal     ErrorType = "INTERNAL"
	ErrorTypeNotFound     ErrorType = "NOT_FOUND"
	ErrorTypeBadRequest   ErrorType = "BAD_REQUEST"
	ErrorTypeConflict     ErrorType = "CONFLICT"
	ErrorTypeUnauthorized ErrorType = "UNAUTHORIZED"
	ErrorTypeForbidden    ErrorType = "FORBIDDEN"
	ErrorTypeTimeout      ErrorType = "TIMEOUT"
	ErrorTypeUnavailable  ErrorType = "UNAVAILABLE"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type AppError struct {
	Type        ErrorType         `json:"type"`
	Code        int               `json:"-"`
	Message     string            `json:"message"`
	Retryable   bool              `json:"retryable,omitempty"`
	Validations []ValidationError `json:"validations,omitempty"`
	Internal    error             `json:"-"`
}

func (e *AppError) Error() string {
	if e.Internal != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Type, e.Message, e.Internal)
	}
	return fmt.Sprintf("[%s] %s", e.Type, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Internal
}

func (e *AppError) Is(target error) bool {
	t, ok := target.(*AppError)
	if !ok {
		return false
	}
	return e.Type == t.Type && e.Message == t.Message
}

func (e *AppError) WithInternal(err error) *AppError {
	copy := *e
	copy.Internal = err
	return &copy
}

func (e *AppError) WithMessage(message string) *AppError {
	copy := *e
	copy.Message = message
	return &copy
}

func (e *AppError) WithValidations(validations []ValidationError) *AppError {
	copy := *e
	copy.Validations = validations
	return &copy
}

func (e *AppError) AsRetryable() *AppError {
	copy := *e
	copy.Retryable = true
	return &copy
}

func NewValidationError(validations []ValidationError) *AppError {
	return ErrValidationFailed.WithValidations(validations)
}

// --- Sentinel Errors (Repository Layer) ---

// ErrUserNotFound is a sentinel error for when a user is not found.
var ErrUserNotFound = ErrNotFound.WithMessage("user not found")

// ErrRoleNotFound is a sentinel error for when a role is not found.
var ErrRoleNotFound = ErrNotFound.WithMessage("role not found")

// ErrTokenNotFound is a sentinel error for when a refresh token is not found.
var ErrTokenNotFound = ErrNotFound.WithMessage("refresh token not found")

// ErrFindByUserID is a sentinel error for when a refresh token by user ID is not found.
var ErrFindByUserID = ErrNotFound.WithMessage("refresh token not found by user ID")

// ErrParseDate is returned when parsing the expiration date of a token fails.
var ErrParseDate = ErrInternal.WithMessage("failed to parse expiration date")

// ErrFailed creates a generic repository error with the given operation description.
// This single function replaces dozens of individual error variables across all services.
func ErrFailed(operation string) *AppError {
	return ErrInternal.WithMessage("Failed to " + operation)
}

// ErrNotFoundResponse creates a not-found error response with the given entity name.
// The entity name should be capitalized (e.g., "User", "Transaction") for proper response formatting.
// This replaces individual ErrXxxNotFound variables across all services.
func ErrNotFoundResponse(entity string) *AppError {
	return ErrNotFound.WithMessage(entity + " not found")
}

// ErrNoRowsOrFailed maps a single-row mutation error: a "no rows in result set"
// error (sql.ErrNoRows) becomes a 404 not-found AppError for the given entity,
// and any other error becomes a generic failure for the given operation. This
// avoids leaking a misleading 500 INTERNAL when a mutation targets a row that
// does not exist (or is in the wrong state).
func ErrNoRowsOrFailed(err error, entity, operation string) *AppError {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFoundResponse(entity).WithInternal(err)
	}
	return ErrFailed(operation).WithInternal(err)
}

// ErrConstraintOrFailed maps database constraint violations to a stable
// conflict response while retaining the original database error for logs.
// PostgreSQL SQLSTATE 23505 is a unique violation and 23503 is a foreign-key
// violation. Neither code nor the database message is exposed to API clients.
func ErrConstraintOrFailed(err error, entity, operation string) *AppError {
	var stateErr interface{ SQLState() string }
	if errors.As(err, &stateErr) {
		switch stateErr.SQLState() {
		case "23505":
			return ErrConflict.WithMessage(entity + " already exists").WithInternal(err)
		case "23503":
			return ErrBadRequest.WithMessage("invalid related " + entity + " reference").WithInternal(err)
		}
	}

	return ErrNoRowsOrFailed(err, entity, operation)
}

var (
	ErrBadRequest = &AppError{
		Type:    ErrorTypeBadRequest,
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}

	ErrValidationFailed = &AppError{
		Type:    ErrorTypeBadRequest,
		Code:    http.StatusBadRequest,
		Message: "Validation failed",
	}

	ErrUnauthorized = &AppError{
		Type:    ErrorTypeUnauthorized,
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}

	ErrForbidden = &AppError{
		Type:    ErrorTypeForbidden,
		Code:    http.StatusForbidden,
		Message: "Forbidden",
	}

	ErrNotFound = &AppError{
		Type:    ErrorTypeNotFound,
		Code:    http.StatusNotFound,
		Message: "Resource not found",
	}

	ErrConflict = &AppError{
		Type:    ErrorTypeConflict,
		Code:    http.StatusConflict,
		Message: "Resource conflict",
	}

	ErrTooManyRequests = &AppError{
		Type:      ErrorTypeBadRequest, // Or a dedicated RATE_LIMIT type
		Code:      http.StatusTooManyRequests,
		Message:   "Too many requests",
		Retryable: true,
	}

	ErrInternal = &AppError{
		Type:    ErrorTypeInternal,
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	}

	ErrServiceUnavailable = &AppError{
		Type:      ErrorTypeUnavailable,
		Code:      http.StatusServiceUnavailable,
		Message:   "Service unavailable",
		Retryable: true,
	}

	ErrTimeout = &AppError{
		Type:      ErrorTypeTimeout,
		Code:      http.StatusGatewayTimeout,
		Message:   "Request timeout",
		Retryable: true,
	}
)
