package handler

import (
	"context"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-merchant/service"
	pbdocument "github.com/MamangRust/monolith-payment-gateway-pb/merchant_document"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	merchantdocument_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_document_errors/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type merchantDocumentCommandHandleGrpc struct {
	pbdocument.UnimplementedMerchantDocumentCommandServiceServer

	merchantDocumentCommand service.MerchantDocumentCommandService
}

func NewMerchantDocumentCommandHandleGrpc(merchantCommand service.MerchantDocumentCommandService) MerchantDocumentCommandHandleGrpc {
	return &merchantDocumentCommandHandleGrpc{
		merchantDocumentCommand: merchantCommand,
	}
}

func (s *merchantDocumentCommandHandleGrpc) Create(ctx context.Context, req *pbdocument.CreateMerchantDocumentRequest) (*pbdocument.ApiResponseMerchantDocument, error) {
	request := requests.CreateMerchantDocumentRequest{
		MerchantID:   int(req.GetMerchantId()),
		DocumentType: req.GetDocumentType(),
		DocumentUrl:  req.GetDocumentUrl(),
	}

	if err := request.Validate(); err != nil {
		return nil, merchantdocument_errors.ErrGrpcValidateCreateMerchantDocument
	}

	document, err := s.merchantDocumentCommand.CreateMerchantDocument(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoDocument := &pbdocument.MerchantDocument{
		DocumentId:   int32(document.DocumentID),
		MerchantId:   int32(document.MerchantID),
		DocumentType: document.DocumentType,
		DocumentUrl:  document.DocumentUrl,
		Status:       document.Status,
		Note:         StringValue(document.Note),
		UploadedAt:   document.UploadedAt.Time.Format(time.RFC3339),
		UpdatedAt:    document.UpdatedAt.Time.Format(time.RFC3339),
	}

	response := &pbdocument.ApiResponseMerchantDocument{
		Status:  "success",
		Message: "Successfully created merchant document",
		Data:    protoDocument,
	}

	return response, nil
}

func (s *merchantDocumentCommandHandleGrpc) Update(ctx context.Context, req *pbdocument.UpdateMerchantDocumentRequest) (*pbdocument.ApiResponseMerchantDocument, error) {
	id := int(req.GetDocumentId())

	if id == 0 {
		return nil, merchantdocument_errors.ErrGrpcMerchantInvalidID
	}

	request := requests.UpdateMerchantDocumentRequest{
		DocumentID:   &id,
		MerchantID:   int(req.GetMerchantId()),
		DocumentType: req.GetDocumentType(),
		DocumentUrl:  req.GetDocumentUrl(),
		Status:       req.GetStatus(),
		Note:         req.GetNote(),
	}

	if err := request.Validate(); err != nil {
		return nil, merchantdocument_errors.ErrGrpcValidateUpdateMerchantDocument
	}

	document, err := s.merchantDocumentCommand.UpdateMerchantDocument(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoDocument := &pbdocument.MerchantDocument{
		DocumentId:   int32(document.DocumentID),
		MerchantId:   int32(document.MerchantID),
		DocumentType: document.DocumentType,
		DocumentUrl:  document.DocumentUrl,
		Status:       document.Status,
		Note:         StringValue(document.Note),
		UploadedAt:   document.UploadedAt.Time.Format(time.RFC3339),
		UpdatedAt:    document.UpdatedAt.Time.Format(time.RFC3339),
	}

	response := &pbdocument.ApiResponseMerchantDocument{
		Status:  "success",
		Message: "Successfully updated merchant document",
		Data:    protoDocument,
	}

	return response, nil
}

func (s *merchantDocumentCommandHandleGrpc) UpdateStatus(ctx context.Context, req *pbdocument.UpdateMerchantDocumentStatusRequest) (*pbdocument.ApiResponseMerchantDocument, error) {
	id := int(req.GetDocumentId())

	if id == 0 {
		return nil, merchantdocument_errors.ErrGrpcMerchantInvalidID
	}

	request := requests.UpdateMerchantDocumentStatusRequest{
		DocumentID: &id,
		MerchantID: int(req.GetMerchantId()),
		Status:     req.GetStatus(),
		Note:       req.GetNote(),
	}

	if err := request.Validate(); err != nil {
		return nil, merchantdocument_errors.ErrGrpcValidateUpdateMerchantDocument
	}

	document, err := s.merchantDocumentCommand.UpdateMerchantDocumentStatus(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoDocument := &pbdocument.MerchantDocument{
		DocumentId:   int32(document.DocumentID),
		MerchantId:   int32(document.MerchantID),
		DocumentType: document.DocumentType,
		DocumentUrl:  document.DocumentUrl,
		Status:       document.Status,
		Note:         StringValue(document.Note),
		UploadedAt:   document.UploadedAt.Time.Format(time.RFC3339),
		UpdatedAt:    document.UpdatedAt.Time.Format(time.RFC3339),
	}

	response := &pbdocument.ApiResponseMerchantDocument{
		Status:  "success",
		Message: "Successfully updated merchant document status",
		Data:    protoDocument,
	}

	return response, nil
}

func (s *merchantDocumentCommandHandleGrpc) Trashed(ctx context.Context, req *pbdocument.FindMerchantDocumentByIdRequest) (*pbdocument.ApiResponseMerchantDocumentDeleteAt, error) {
	id := int(req.GetDocumentId())

	if id == 0 {
		return nil, merchantdocument_errors.ErrGrpcMerchantInvalidID
	}

	document, err := s.merchantDocumentCommand.TrashedMerchantDocument(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var deletedAt *wrapperspb.StringValue
	if document.DeletedAt.Valid {
		deletedAt = wrapperspb.String(document.DeletedAt.Time.Format(time.RFC3339))
	}

	protoDocument := &pbdocument.MerchantDocumentDeleteAt{
		DocumentId:   int32(document.DocumentID),
		MerchantId:   int32(document.MerchantID),
		DocumentType: document.DocumentType,
		DocumentUrl:  document.DocumentUrl,
		Status:       document.Status,
		Note:         StringValue(document.Note),
		UploadedAt:   document.UploadedAt.Time.Format(time.RFC3339),
		UpdatedAt:    document.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt:    deletedAt,
	}

	response := &pbdocument.ApiResponseMerchantDocumentDeleteAt{
		Status:  "success",
		Message: "Successfully trashed merchant document",
		Data:    protoDocument,
	}

	return response, nil
}

func (s *merchantDocumentCommandHandleGrpc) Restore(ctx context.Context, req *pbdocument.FindMerchantDocumentByIdRequest) (*pbdocument.ApiResponseMerchantDocumentDeleteAt, error) {
	id := int(req.GetDocumentId())

	if id == 0 {
		return nil, merchantdocument_errors.ErrGrpcMerchantInvalidID
	}

	document, err := s.merchantDocumentCommand.RestoreMerchantDocument(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoDocument := &pbdocument.MerchantDocumentDeleteAt{
		DocumentId:   int32(document.DocumentID),
		MerchantId:   int32(document.MerchantID),
		DocumentType: document.DocumentType,
		DocumentUrl:  document.DocumentUrl,
		Status:       document.Status,
		Note:         StringValue(document.Note),
		UploadedAt:   document.UploadedAt.Time.Format(time.RFC3339),
		UpdatedAt:    document.UpdatedAt.Time.Format(time.RFC3339),
	}

	response := &pbdocument.ApiResponseMerchantDocumentDeleteAt{
		Status:  "success",
		Message: "Successfully restored merchant document",
		Data:    protoDocument,
	}

	return response, nil
}

func (s *merchantDocumentCommandHandleGrpc) DeletePermanent(ctx context.Context, req *pbdocument.FindMerchantDocumentByIdRequest) (*pbdocument.ApiResponseMerchantDocumentDelete, error) {
	id := int(req.GetDocumentId())

	if id == 0 {
		return nil, merchantdocument_errors.ErrGrpcMerchantInvalidID
	}

	_, err := s.merchantDocumentCommand.DeleteMerchantDocumentPermanent(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	response := &pbdocument.ApiResponseMerchantDocumentDelete{
		Status:  "success",
		Message: "Successfully permanently deleted merchant document",
	}

	return response, nil
}

func (s *merchantDocumentCommandHandleGrpc) RestoreAll(ctx context.Context, _ *emptypb.Empty) (*pbdocument.ApiResponseMerchantDocumentAll, error) {
	_, err := s.merchantDocumentCommand.RestoreAllMerchantDocument(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	response := &pbdocument.ApiResponseMerchantDocumentAll{
		Status:  "success",
		Message: "Successfully restored all merchant documents",
	}

	return response, nil
}

func (s *merchantDocumentCommandHandleGrpc) DeleteAllPermanent(ctx context.Context, _ *emptypb.Empty) (*pbdocument.ApiResponseMerchantDocumentAll, error) {
	_, err := s.merchantDocumentCommand.DeleteAllMerchantDocumentPermanent(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	response := &pbdocument.ApiResponseMerchantDocumentAll{
		Status:  "success",
		Message: "Successfully permanently deleted all merchant documents",
	}

	return response, nil
}

func StringValue(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func Int32Value(v *int32) int32 {
	if v == nil {
		return 0
	}

	return *v
}
