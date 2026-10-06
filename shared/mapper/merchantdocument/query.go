package merchantdocumentapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant_document"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper"
)

type merchantDocumentQueryResponseMapper struct{}

func NewMerchantDocumentQueryResponseMapper() MerchantDocumentQueryResponseMapper {
	return &merchantDocumentQueryResponseMapper{}
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
func (m *merchantDocumentQueryResponseMapper) ToApiResponseMerchantDocument(doc *pb.ApiResponseMerchantDocument) *response.ApiResponseMerchantDocument {
	return &response.ApiResponseMerchantDocument{
		Status:  doc.Status,
		Message: doc.Message,
		Data:    m.mapMerchantDocument(doc.Data),
	}
}

// ToApiResponsePaginationMerchantDocument maps a paginated gRPC merchant document response
// to an HTTP API response. It constructs an ApiResponsePaginationMerchantDocument by copying
// the status and message fields, mapping the document data slice to a slice of MerchantDocumentResponse,
// and including pagination metadata.
//
// Args:
//
//	docs: A pointer to a pb.ApiResponsePaginationMerchantDocument containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsePaginationMerchantDocument with mapped data and pagination info.
func (m *merchantDocumentQueryResponseMapper) ToApiResponsePaginationMerchantDocument(docs *pb.ApiResponsePaginationMerchantDocument) *response.ApiResponsePaginationMerchantDocument {
	return &response.ApiResponsePaginationMerchantDocument{
		Status:     docs.Status,
		Message:    docs.Message,
		Data:       m.mapMerchantDocuments(docs.Data),
		Pagination: apimapper.MapPaginationMeta(docs.PaginationMeta),
	}
}

// ToApiResponsePaginationMerchantDocumentDeleteAt maps a paginated gRPC response of
// soft-deleted merchant documents to an HTTP API response. It constructs an
// ApiResponsePaginationMerchantDocumentDeleteAt by copying the status and message fields,
// mapping the document data slice to a slice of MerchantDocumentResponseDeleteAt, and including
// pagination metadata.
//
// Args:
//
//	docs: A pointer to a pb.ApiResponsePaginationMerchantDocumentAt containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsePaginationMerchantDocumentDeleteAt with mapped data and pagination info.
func (m *merchantDocumentQueryResponseMapper) ToApiResponsePaginationMerchantDocumentDeleteAt(docs *pb.ApiResponsePaginationMerchantDocumentAt) *response.ApiResponsePaginationMerchantDocumentDeleteAt {
	return &response.ApiResponsePaginationMerchantDocumentDeleteAt{
		Status:     docs.Status,
		Message:    docs.Message,
		Data:       m.mapMerchantDocumentsDeletedAt(docs.Data),
		Pagination: apimapper.MapPaginationMeta(docs.PaginationMeta),
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
func (m *merchantDocumentQueryResponseMapper) ToApiResponseMerchantDocumentDelete(resp *pb.ApiResponseMerchantDocumentDelete) *response.ApiResponseMerchantDocumentDelete {
	return &response.ApiResponseMerchantDocumentDelete{
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
func (m *merchantDocumentQueryResponseMapper) mapMerchantDocument(doc *pb.MerchantDocument) *response.MerchantDocumentResponse {
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

// mapMerchantDocuments maps a slice of gRPC merchant document responses to a slice of HTTP API responses.
// It constructs a slice of MerchantDocumentResponse by copying the fields from the gRPC responses.
//
// Args:
//
//	docs: A slice of pointers to pb.MerchantDocument containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantDocumentResponse with the mapped data.
func (m *merchantDocumentQueryResponseMapper) mapMerchantDocuments(docs []*pb.MerchantDocument) []*response.MerchantDocumentResponse {
	var responses []*response.MerchantDocumentResponse
	for _, doc := range docs {
		responses = append(responses, m.mapMerchantDocument(doc))
	}
	return responses
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
func (m *merchantDocumentQueryResponseMapper) mapMerchantDocumentDeletedAt(doc *pb.MerchantDocumentDeleteAt) *response.MerchantDocumentResponseDeleteAt {
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

// mapMerchantDocumentsDeletedAt maps a slice of gRPC merchant document responses for soft-deleted documents
// to a slice of HTTP API responses. It constructs a slice of MerchantDocumentResponseDeleteAt by iterating
// over each MerchantDocumentDeleteAt in the input slice, mapping them using mapMerchantDocumentDeletedAt,
// and handling the optional DeletedAt field.
//
// Args:
//
//	docs: A slice of pointers to pb.MerchantDocumentDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantDocumentResponseDeleteAt with the mapped data.
func (m *merchantDocumentQueryResponseMapper) mapMerchantDocumentsDeletedAt(docs []*pb.MerchantDocumentDeleteAt) []*response.MerchantDocumentResponseDeleteAt {
	var responses []*response.MerchantDocumentResponseDeleteAt
	for _, doc := range docs {
		responses = append(responses, m.mapMerchantDocumentDeletedAt(doc))
	}
	return responses
}
