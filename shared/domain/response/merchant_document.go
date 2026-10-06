package response

// MerchantDocumentResponse represents the basic merchant document information returned in API responses.
// Contains all relevant document metadata without deletion information.
type MerchantDocumentResponse struct {
	ID           int    `json:"id"`            // Unique document identifier
	MerchantID   int    `json:"merchant_id"`   // ID of the merchant this document belongs to
	DocumentType string `json:"document_type"` // Type of document (e.g., "license", "identity_proof")
	DocumentURL  string `json:"document_url"`  // URL or storage path where the document is saved
	Status       string `json:"status"`        // Current status ("pending", "approved", "rejected")
	Note         string `json:"note"`          // Administrative notes about the document
	CreatedAt    string `json:"created_at"`    // Timestamp when document was uploaded (RFC3339 format)
	UpdatedAt    string `json:"updated_at"`    // Timestamp when document was last updated
}

// MerchantDocumentResponseDeleteAt extends MerchantDocumentResponse with deletion information.
// Used in administrative interfaces showing soft-deleted documents.
type MerchantDocumentResponseDeleteAt struct {
	ID           int     `json:"id"`            // Unique document identifier
	MerchantID   int     `json:"merchant_id"`   // ID of the merchant
	DocumentType string  `json:"document_type"` // Type of document
	DocumentURL  string  `json:"document_url"`  // Document storage path/URL
	Status       string  `json:"status"`        // Last known status before deletion
	Note         string  `json:"note"`          // Final administrative notes
	CreatedAt    string  `json:"created_at"`    // Original creation timestamp
	UpdatedAt    string  `json:"updated_at"`    // Last update timestamp
	DeletedAt    *string `json:"deleted_at"`    // Deletion timestamp (nil if not deleted)
}

// ApiResponsesMerchantDocument is the response format for multiple document listings.
// Used when returning arrays of document data without pagination.
type ApiResponsesMerchantDocument struct {
	Status  string                      `json:"status"`  // Response status ("success" or "error")
	Message string                      `json:"message"` // Descriptive message
	Data    []*MerchantDocumentResponse `json:"data"`    // Array of document responses
}

// ApiResponseMerchantDocument is the standard response format for single document requests.
type ApiResponseMerchantDocument struct {
	Status  string                    `json:"status"`  // Response status
	Message string                    `json:"message"` // Descriptive message
	Data    *MerchantDocumentResponse `json:"data"`    // Single document response
}

type ApiResponseMerchantDocumentDeleteAt struct {
	Status  string                            `json:"status"`  // Response status
	Message string                            `json:"message"` // Descriptive message
	Data    *MerchantDocumentResponseDeleteAt `json:"data"`    // Single document response
}

// ApiResponseMerchantDocumentDelete is the response format for document deletion operations.
type ApiResponseMerchantDocumentDelete struct {
	Status  string `json:"status"`  // Operation status
	Message string `json:"message"` // Result message
}

// ApiResponseMerchantDocumentAll is the response format for bulk document operations.
type ApiResponseMerchantDocumentAll struct {
	Status  string `json:"status"`  // Operation status
	Message string `json:"message"` // Result message
}

// ApiResponsePaginationMerchantDocument is the paginated response for document lists.
type ApiResponsePaginationMerchantDocument struct {
	Status     string                      `json:"status"`     // Response status
	Message    string                      `json:"message"`    // Descriptive message
	Data       []*MerchantDocumentResponse `json:"data"`       // Array of document data
	Pagination *PaginationMeta             `json:"pagination"` // Pagination metadata
}

// ApiResponsePaginationMerchantDocumentDeleteAt is the paginated response including deleted documents.
type ApiResponsePaginationMerchantDocumentDeleteAt struct {
	Status     string                              `json:"status"`     // Response status
	Message    string                              `json:"message"`    // Descriptive message
	Data       []*MerchantDocumentResponseDeleteAt `json:"data"`       // Array of documents (with deletion info)
	Pagination *PaginationMeta                     `json:"pagination"` // Pagination metadata
}
