package requests

import "github.com/go-playground/validator/v10"

// FindAllMerchantDocuments represents the request parameters for searching and paginating merchant documents.
// Used in document listing operations with search functionality.
type FindAllMerchantDocuments struct {
	Search   string `json:"search" validate:"required"`         // Search term to filter documents (matches against document type or status)
	Page     int    `json:"page" validate:"min=1"`              // Page number for pagination (1-based index)
	PageSize int    `json:"page_size" validate:"min=1,max=100"` // Number of items per page (1-100)
}

// CreateMerchantDocumentRequest represents the payload for uploading a new merchant document.
// Contains all necessary information to register a document for a merchant.
type CreateMerchantDocumentRequest struct {
	MerchantID   int    `json:"merchant_id" validate:"required,min=1"` // ID of the merchant this document belongs to
	DocumentType string `json:"document_type" validate:"required"`     // Type of document (e.g., "license", "identity_proof")
	DocumentUrl  string `json:"document_url" validate:"required"`      // URL or storage path where the document is saved
}

// UpdateMerchantDocumentRequest represents the payload for updating an existing merchant document.
// Contains all modifiable fields for a merchant document record.
type UpdateMerchantDocumentRequest struct {
	DocumentID   *int   `json:"document_id"`                           // ID of the document being updated (optional in some flows)
	MerchantID   int    `json:"merchant_id" validate:"required,min=1"` // ID of the merchant (for verification)
	DocumentType string `json:"document_type" validate:"required"`     // Updated document type
	DocumentUrl  string `json:"document_url" validate:"required"`      // Updated document URL/path
	Status       string `json:"status" validate:"required"`            // Updated review status (e.g., "pending", "approved")
	Note         string `json:"note" validate:"required"`              // Administrative notes about the document
}

// UpdateMerchantDocumentStatusRequest represents the payload for changing a document's review status.
// Used specifically for status updates during document review processes.
type UpdateMerchantDocumentStatusRequest struct {
	DocumentID *int   `json:"document_id"`                           // ID of the document being updated (optional in some flows)
	MerchantID int    `json:"merchant_id" validate:"required,min=1"` // ID of the merchant (for verification)
	Status     string `json:"status" validate:"required"`            // New review status (e.g., "approved", "rejected")
	Note       string `json:"note" validate:"required"`              // Explanation for status change
}

// Validate performs basic validation of CreateMerchantDocumentRequest fields using struct tags.
// Returns:
//   - error: if validation fails according to struct tag rules
func (r *CreateMerchantDocumentRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

// Validate performs basic validation of UpdateMerchantDocumentRequest fields using struct tags.
// Returns:
//   - error: if validation fails according to struct tag rules
func (r *UpdateMerchantDocumentRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

// Validate performs basic validation of UpdateMerchantDocumentStatusRequest fields using struct tags.
// Returns:
//   - error: if validation fails according to struct tag rules
func (r *UpdateMerchantDocumentStatusRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}
