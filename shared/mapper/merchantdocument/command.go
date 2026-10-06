package merchantdocumentapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant_document"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type merchantDocumentCommandResponseMapepr struct{}

// NewMerchantDocumentCommandResponseMapper returns a new instance of
// merchantDocumentCommandResponseMapepr which provides methods to map
// gRPC merchant document responses to HTTP API responses for command operations.
func NewMerchantDocumentCommandResponseMapper() MerchantDocumentCommandResponseMapper {
	return &merchantDocumentCommandResponseMapepr{}
}

// ToApiResponseMerchantDocument maps a gRPC merchant document response to an HTTP API response.
// It constructs an ApiResponseMerchantDocument by copying the status and message fields
// and mapping the document data to a MerchantDocumentResponse.
//
// Args:
//
//	doc: A pointer to a pb.ApiResponseMerchantDocument containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantDocument with mapped data.
func (m *merchantDocumentCommandResponseMapepr) ToApiResponseMerchantDocument(doc *pb.ApiResponseMerchantDocument) *response.ApiResponseMerchantDocument {
	return &response.ApiResponseMerchantDocument{
		Status:  doc.Status,
		Message: doc.Message,
		Data:    m.mapMerchantDocument(doc.Data),
	}
}

// ToApiResponseMerchantDocumentDeleteAt maps a soft-deleted gRPC merchant document response
// to an HTTP API response. It constructs an ApiResponseMerchantDocumentDeleteAt by copying
// the status and message fields, and mapping the document data to a MerchantDocumentResponseDeleteAt.
//
// Args:
//
//	doc: A pointer to a pb.ApiResponseMerchantDocumentDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantDocumentDeleteAt with mapped data.
func (m *merchantDocumentCommandResponseMapepr) ToApiResponseMerchantDocumentDeleteAt(doc *pb.ApiResponseMerchantDocumentDeleteAt) *response.ApiResponseMerchantDocumentDeleteAt {
	return &response.ApiResponseMerchantDocumentDeleteAt{
		Status:  doc.Status,
		Message: doc.Message,
		Data:    m.mapMerchantDocumentDeletedAt(doc.Data),
	}
}

// ToApiResponseMerchantDocumentDeleteAt maps a soft-deleted gRPC merchant document response
// to an HTTP API response. It constructs an ApiResponseMerchantDocumentDelete by copying
// the status and message fields from the gRPC response.
//
// Args:
//
//	resp: A pointer to a pb.ApiResponseMerchantDocumentDelete containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantDocumentDelete with the status and message set.
func (m *merchantDocumentCommandResponseMapepr) ToApiResponseMerchantDocumentDelete(resp *pb.ApiResponseMerchantDocumentDelete) *response.ApiResponseMerchantDocumentDelete {
	return &response.ApiResponseMerchantDocumentDelete{
		Status:  resp.Status,
		Message: resp.Message,
	}
}

// ToApiResponseMerchantDocumentAll maps a gRPC response containing all merchant documents
// to an HTTP API response. It constructs an ApiResponseMerchantDocumentAll by copying
// the status and message fields from the gRPC response.
//
// Args:
//
//	resp: A pointer to a pb.ApiResponseMerchantDocumentAll containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantDocumentAll with the status and message set.
func (m *merchantDocumentCommandResponseMapepr) ToApiResponseMerchantDocumentAll(resp *pb.ApiResponseMerchantDocumentAll) *response.ApiResponseMerchantDocumentAll {
	return &response.ApiResponseMerchantDocumentAll{
		Status:  resp.Status,
		Message: resp.Message,
	}
}

// mapMerchantDocument maps a gRPC merchant document response to an HTTP API response.
// It constructs a MerchantDocumentResponse by copying the fields from the gRPC response.
//
// Args:
//
//	doc: A pointer to a pb.MerchantDocument containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantDocumentResponse with the mapped data.
func (m *merchantDocumentCommandResponseMapepr) mapMerchantDocument(doc *pb.MerchantDocument) *response.MerchantDocumentResponse {
	return &response.MerchantDocumentResponse{
		ID:           int(doc.DocumentId),
		MerchantID:   int(doc.MerchantId),
		DocumentType: doc.DocumentType,
		DocumentURL:  doc.DocumentUrl,
		Status:       doc.Status,
		Note:         doc.Note,
		CreatedAt:    doc.UploadedAt,
		UpdatedAt:    doc.UpdatedAt,
	}
}

// mapMerchantDocumentDeletedAt maps a gRPC merchant document response for soft-deleted documents
// to an HTTP API response. It constructs a MerchantDocumentResponseDeleteAt by copying
// the fields from the gRPC response and handling the optional DeletedAt field.
//
// Args:
//
//	doc: A pointer to a pb.MerchantDocumentDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantDocumentResponseDeleteAt with the mapped data.
func (m *merchantDocumentCommandResponseMapepr) mapMerchantDocumentDeletedAt(doc *pb.MerchantDocumentDeleteAt) *response.MerchantDocumentResponseDeleteAt {
	var deletedAt *string

	if doc.DeletedAt != nil {
		deletedAt = &doc.DeletedAt.Value
	}

	return &response.MerchantDocumentResponseDeleteAt{
		ID:           int(doc.DocumentId),
		MerchantID:   int(doc.MerchantId),
		DocumentType: doc.DocumentType,
		DocumentURL:  doc.DocumentUrl,
		Status:       doc.Status,
		Note:         doc.Note,
		CreatedAt:    doc.UploadedAt,
		UpdatedAt:    doc.UpdatedAt,
		DeletedAt:    deletedAt,
	}
}
