package errors

func NewErrorResponse(message string, code int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}
