package response

// PaginationMeta contains metadata for paginated responses.
// Used to provide navigation information in list endpoints.
type PaginationMeta struct {
	CurrentPage  int `json:"current_page"`  // Current page number (1-based index)
	PageSize     int `json:"page_size"`     // Number of items per page
	TotalPages   int `json:"total_pages"`   // Total number of pages available
	TotalRecords int `json:"total_records"` // Total number of items across all pages
}

// ApiResponse is a generic response wrapper for successful API responses.
// The type parameter T allows this to be used with any data type.
type ApiResponse[T any] struct {
	Status  string `json:"status"`  // Response status ("success", "error", etc.)
	Message string `json:"message"` // Descriptive message about the response
	Data    T      `json:"data"`    // The actual response payload of type T
}

// APIResponsePagination is a generic paginated response wrapper.
// Combines data payload with pagination metadata.
type APIResponsePagination[T any] struct {
	Status  string         `json:"status"`     // Response status
	Message string         `json:"message"`    // Descriptive message
	Data    T              `json:"data"`       // The paginated data payload of type T
	Meta    PaginationMeta `json:"pagination"` // Pagination metadata
}

// ErrorResponse represents standardized error responses from the API.
// Used to maintain consistent error reporting across endpoints.
type ErrorResponse struct {
	Status  string `json:"status"`  // Error status ("error", "fail", etc.)
	Message string `json:"message"` // Human-readable error message
	Code    int    `json:"code"`    // HTTP status code or application error code
}
