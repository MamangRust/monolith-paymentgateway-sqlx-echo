package handler

import (
	"context"
	"math"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-merchant/service"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/common"
	pbdocument "github.com/MamangRust/monolith-payment-gateway-pb/merchant_document"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	merchantdocument_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_document_errors/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type merchantDocumentQueryHandleGrpc struct {
	pbdocument.UnimplementedMerchantDocumentQueryServiceServer

	merchantDocumentQuery service.MerchantDocumentQueryService
}

func NewMerchantDocumentQueryHandleGrpc(merchantQuery service.MerchantDocumentQueryService) MerchantDocumentQueryHandleGrpc {
	return &merchantDocumentQueryHandleGrpc{
		merchantDocumentQuery: merchantQuery,
	}
}

func (s *merchantDocumentQueryHandleGrpc) FindAll(ctx context.Context, req *pbdocument.FindAllMerchantDocumentsRequest) (*pbdocument.ApiResponsePaginationMerchantDocument, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchantDocuments{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	documents, totalRecords, err := s.merchantDocumentQuery.FindAll(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var protoDocuments []*pbdocument.MerchantDocument
	for _, doc := range documents {
		protoDocuments = append(protoDocuments, &pbdocument.MerchantDocument{
			DocumentId:   int32(doc.DocumentID),
			MerchantId:   int32(doc.MerchantID),
			DocumentType: doc.DocumentType,
			DocumentUrl:  doc.DocumentUrl,
			Status:       doc.Status,
			Note:         StringValue(doc.Note),
			UploadedAt:   doc.UploadedAt.Time.Format(time.RFC3339),
			UpdatedAt:    doc.UpdatedAt.Time.Format(time.RFC3339),
		})
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))
	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	response := &pbdocument.ApiResponsePaginationMerchantDocument{
		Status:         "success",
		Message:        "Successfully fetched merchant documents",
		Data:           protoDocuments,
		PaginationMeta: paginationMeta,
	}

	return response, nil
}

func (s *merchantDocumentQueryHandleGrpc) FindById(ctx context.Context, req *pbdocument.FindMerchantDocumentByIdRequest) (*pbdocument.ApiResponseMerchantDocument, error) {
	id := int(req.GetDocumentId())

	if id == 0 {
		return nil, merchantdocument_errors.ErrGrpcMerchantInvalidID
	}

	doc, err := s.merchantDocumentQuery.FindById(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoDocument := &pbdocument.MerchantDocument{
		DocumentId:   int32(doc.DocumentID),
		MerchantId:   int32(doc.MerchantID),
		DocumentType: doc.DocumentType,
		DocumentUrl:  doc.DocumentUrl,
		Status:       doc.Status,
		Note:         StringValue(doc.Note),
		UploadedAt:   doc.UploadedAt.Time.Format(time.RFC3339),
		UpdatedAt:    doc.UpdatedAt.Time.Format(time.RFC3339),
	}

	response := &pbdocument.ApiResponseMerchantDocument{
		Status:  "success",
		Message: "Successfully fetched merchant document",
		Data:    protoDocument,
	}

	return response, nil
}

func (s *merchantDocumentQueryHandleGrpc) FindAllActive(ctx context.Context, req *pbdocument.FindAllMerchantDocumentsRequest) (*pbdocument.ApiResponsePaginationMerchantDocumentAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchantDocuments{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	documents, totalRecords, err := s.merchantDocumentQuery.FindByActive(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var protoDocuments []*pbdocument.MerchantDocumentDeleteAt
	for _, doc := range documents {
		protoDocuments = append(protoDocuments, &pbdocument.MerchantDocumentDeleteAt{
			DocumentId:   int32(doc.DocumentID),
			MerchantId:   int32(doc.MerchantID),
			DocumentType: doc.DocumentType,
			DocumentUrl:  doc.DocumentUrl,
			Status:       doc.Status,
			Note:         StringValue(doc.Note),
			UploadedAt:   doc.UploadedAt.Time.Format(time.RFC3339),
			UpdatedAt:    doc.UpdatedAt.Time.Format(time.RFC3339),
		})
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))
	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	response := &pbdocument.ApiResponsePaginationMerchantDocumentAt{
		Status:         "success",
		Message:        "Successfully fetched active merchant documents",
		Data:           protoDocuments,
		PaginationMeta: paginationMeta,
	}

	return response, nil
}

func (s *merchantDocumentQueryHandleGrpc) FindAllTrashed(ctx context.Context, req *pbdocument.FindAllMerchantDocumentsRequest) (*pbdocument.ApiResponsePaginationMerchantDocumentAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchantDocuments{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	documents, totalRecords, err := s.merchantDocumentQuery.FindByTrashed(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var protoDocuments []*pbdocument.MerchantDocumentDeleteAt
	for _, doc := range documents {
		var deletedAt *wrapperspb.StringValue
		if doc.DeletedAt.Valid {
			deletedAt = wrapperspb.String(doc.DeletedAt.Time.String())
		}

		protoDocuments = append(protoDocuments, &pbdocument.MerchantDocumentDeleteAt{
			DocumentId:   int32(doc.DocumentID),
			MerchantId:   int32(doc.MerchantID),
			DocumentType: doc.DocumentType,
			DocumentUrl:  doc.DocumentUrl,
			Status:       doc.Status,
			Note:         StringValue(doc.Note),
			UploadedAt:   doc.UploadedAt.Time.Format(time.RFC3339),
			UpdatedAt:    doc.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:    deletedAt,
		})
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))
	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	response := &pbdocument.ApiResponsePaginationMerchantDocumentAt{
		Status:         "success",
		Message:        "Successfully fetched trashed merchant documents",
		Data:           protoDocuments,
		PaginationMeta: paginationMeta,
	}

	return response, nil
}
