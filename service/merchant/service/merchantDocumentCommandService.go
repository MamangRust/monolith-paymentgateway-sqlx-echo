package service

import (
	"context"

	cache "github.com/MamangRust/monolith-payment-gateway-merchant/redis"
	"github.com/MamangRust/monolith-payment-gateway-merchant/repository"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.uber.org/zap"
)

// merchantDocumentCommandDeps groups dependencies for merchant document commands.
type MerchantDocumentCommandDeps struct {
	Cache                   cache.MerchantDocumentCommandCache
	CommandRepository       repository.MerchantDocumentCommandRepository
	MerchantQueryRepository repository.MerchantQueryRepository
	UserRepository          repository.UserRepository
	Logger                  logger.LoggerInterface
	Observability           observability.TraceLoggerObservability
}

// merchantDocumentCommandService implements command operations for merchant documents.
type merchantDocumentCommandService struct {
	cache         cache.MerchantDocumentCommandCache
	commandRepo   repository.MerchantDocumentCommandRepository
	merchantRepo  repository.MerchantQueryRepository
	userRepo      repository.UserRepository
	logger        logger.LoggerInterface
	observability observability.TraceLoggerObservability
}

// NewMerchantDocumentCommandService constructs a MerchantDocumentCommandService.
func NewMerchantDocumentCommandService(
	params *MerchantDocumentCommandDeps,
) MerchantDocumentCommandService {
	return &merchantDocumentCommandService{
		cache:         params.Cache,
		commandRepo:   params.CommandRepository,
		merchantRepo:  params.MerchantQueryRepository,
		userRepo:      params.UserRepository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *merchantDocumentCommandService) CreateMerchantDocument(ctx context.Context, request *requests.CreateMerchantDocumentRequest) (*db.CreateMerchantDocumentRow, error) {
	const method = "CreateMerchantDocument"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	merchantDocument, err := s.commandRepo.CreateMerchantDocument(ctx, request)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateMerchantDocumentRow](s.logger, err, method, span, zap.Int("merchant_id", request.MerchantID))
	}

	logSuccess("Successfully created merchant document", zap.Int("document_id", int(merchantDocument.DocumentID)))

	return merchantDocument, nil
}

func (s *merchantDocumentCommandService) UpdateMerchantDocument(ctx context.Context, request *requests.UpdateMerchantDocumentRequest) (*db.UpdateMerchantDocumentRow, error) {
	const method = "UpdateMerchantDocument"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	merchantDocument, err := s.commandRepo.UpdateMerchantDocument(ctx, request)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateMerchantDocumentRow](s.logger, err, method, span, zap.Int("document_id", *request.DocumentID))
	}

	s.cache.DeleteCachedMerchantDocuments(ctx, int(merchantDocument.DocumentID))

	logSuccess("Successfully updated merchant document", zap.Int("document_id", int(merchantDocument.DocumentID)))

	return merchantDocument, nil
}

func (s *merchantDocumentCommandService) UpdateMerchantDocumentStatus(ctx context.Context, request *requests.UpdateMerchantDocumentStatusRequest) (*db.UpdateMerchantDocumentStatusRow, error) {
	const method = "UpdateMerchantDocumentStatus"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	merchantDocument, err := s.commandRepo.UpdateMerchantDocumentStatus(ctx, request)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateMerchantDocumentStatusRow](s.logger, err, method, span, zap.Int("merchant_id", request.MerchantID))
	}

	s.cache.DeleteCachedMerchantDocuments(ctx, int(merchantDocument.DocumentID))

	logSuccess("Successfully updated merchant document status", zap.Int("document_id", int(merchantDocument.DocumentID)))

	return merchantDocument, nil
}

func (s *merchantDocumentCommandService) TrashedMerchantDocument(ctx context.Context, documentID int) (*db.MerchantDocument, error) {
	const method = "TrashedMerchantDocument"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	res, err := s.commandRepo.TrashedMerchantDocument(ctx, documentID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.MerchantDocument](s.logger, err, method, span, zap.Int("document_id", documentID))
	}

	s.cache.DeleteCachedMerchantDocuments(ctx, documentID)

	logSuccess("Successfully trashed document", zap.Int("document_id", documentID))

	return res, nil
}

func (s *merchantDocumentCommandService) RestoreMerchantDocument(ctx context.Context, documentID int) (*db.MerchantDocument, error) {
	const method = "RestoreMerchantDocument"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	res, err := s.commandRepo.RestoreMerchantDocument(ctx, documentID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.MerchantDocument](s.logger, err, method, span, zap.Int("document_id", documentID))
	}

	logSuccess("Successfully restored document", zap.Int("document_id", documentID))

	return res, nil
}

func (s *merchantDocumentCommandService) DeleteMerchantDocumentPermanent(ctx context.Context, documentID int) (bool, error) {
	const method = "DeleteMerchantDocumentPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	_, err := s.commandRepo.DeleteMerchantDocumentPermanent(ctx, documentID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](s.logger, err, method, span, zap.Int("document_id", documentID))
	}

	logSuccess("Successfully deleted document permanently", zap.Int("document_id", documentID))

	return true, nil
}

func (s *merchantDocumentCommandService) RestoreAllMerchantDocument(ctx context.Context) (bool, error) {
	const method = "RestoreAllMerchantDocument"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	_, err := s.commandRepo.RestoreAllMerchantDocument(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](s.logger, err, method, span)
	}

	logSuccess("Successfully restored all documents")

	return true, nil
}

func (s *merchantDocumentCommandService) DeleteAllMerchantDocumentPermanent(ctx context.Context) (bool, error) {
	const method = "DeleteAllMerchantDocumentPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	_, err := s.commandRepo.DeleteAllMerchantDocumentPermanent(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](s.logger, err, method, span)
	}

	logSuccess("Successfully deleted all documents permanently")

	return true, nil
}
