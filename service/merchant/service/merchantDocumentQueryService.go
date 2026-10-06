package service

import (
	"context"

	mencache "github.com/MamangRust/monolith-payment-gateway-merchant/redis"
	"github.com/MamangRust/monolith-payment-gateway-merchant/repository"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// merchantDocumentQueryDeps holds dependencies for merchant document queries.
type MerchantDocumentQueryDeps struct {
	Cache         mencache.MerchantDocumentQueryCache
	Repository    repository.MerchantDocumentQueryRepository
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
}

// merchantDocumentQueryService handles read operations for merchant documents.
type merchantDocumentQueryService struct {
	cache         mencache.MerchantDocumentQueryCache
	queryRepo     repository.MerchantDocumentQueryRepository
	logger        logger.LoggerInterface
	observability observability.TraceLoggerObservability
}

// NewMerchantDocumentQueryService constructs a MerchantDocumentQueryService.
func NewMerchantDocumentQueryService(
	params *MerchantDocumentQueryDeps,
) MerchantDocumentQueryService {
	return &merchantDocumentQueryService{
		cache:         params.Cache,
		queryRepo:     params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *merchantDocumentQueryService) FindAll(ctx context.Context, req *requests.FindAllMerchantDocuments) ([]*db.GetMerchantDocumentsRow, *int, error) {
	const method = "FindAll"
	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("page", page), attribute.Int("pageSize", pageSize), attribute.String("search", search))
	defer func() { end(status) }()

	if data, total, found := s.cache.GetCachedMerchantDocuments(ctx, req); found {
		logSuccess("Successfully retrieved all merchant document records from cache", zap.Int("total", *total), zap.Int("page", page), zap.Int("pageSize", pageSize), zap.String("search", search))
		return data, total, nil
	}

	merchantDocuments, err := s.queryRepo.FindAllDocuments(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetMerchantDocumentsRow](s.logger, err, method, span, zap.String("search", search))
	}

	var totalCount int

	if len(merchantDocuments) > 0 {
		totalCount = int(merchantDocuments[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedMerchantDocuments(ctx, req, merchantDocuments, &totalCount)

	logSuccess("Successfully retrieved all merchant document records", zap.Int("total", totalCount), zap.Int("page", page), zap.Int("pageSize", pageSize), zap.String("search", search))

	return merchantDocuments, &totalCount, nil
}

func (s *merchantDocumentQueryService) FindById(ctx context.Context, documentID int) (*db.GetMerchantDocumentRow, error) {
	const method = "FindById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("merchantDocument.id", documentID))
	defer func() { end(status) }()

	if cachedMerchant, found := s.cache.GetCachedMerchantDocument(ctx, documentID); found {
		logSuccess("Successfully found merchant document by ID from cache", zap.Int("merchantDocument.id", documentID))
		return cachedMerchant, nil
	}

	merchantDocument, err := s.queryRepo.FindByIdDocument(ctx, documentID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.GetMerchantDocumentRow](s.logger, err, method, span, zap.Int("merchantDocument.id", documentID))
	}

	s.cache.SetCachedMerchantDocument(ctx, documentID, merchantDocument)

	logSuccess("Successfully found merchant document by ID", zap.Int("merchantDocument.id", documentID))

	return merchantDocument, nil
}

func (s *merchantDocumentQueryService) FindByActive(ctx context.Context, req *requests.FindAllMerchantDocuments) ([]*db.GetActiveMerchantDocumentsRow, *int, error) {
	const method = "FindByActive"
	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("page", page), attribute.Int("pageSize", pageSize), attribute.String("search", search))
	defer func() { end(status) }()

	if data, total, found := s.cache.GetCachedMerchantDocumentsActive(ctx, req); found {
		logSuccess("Successfully retrieved active merchant document records from cache", zap.Int("total", *total), zap.Int("page", page), zap.Int("pageSize", pageSize), zap.String("search", search))
		return data, total, nil
	}

	merchantDocuments, err := s.queryRepo.FindByActiveDocuments(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetActiveMerchantDocumentsRow](s.logger, err, method, span, zap.String("search", search))
	}

	var totalCount int

	if len(merchantDocuments) > 0 {
		totalCount = int(merchantDocuments[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedMerchantDocumentsActive(ctx, req, merchantDocuments, &totalCount)

	logSuccess("Successfully retrieved active merchant document records", zap.Int("total", totalCount), zap.Int("page", page), zap.Int("pageSize", pageSize), zap.String("search", search))

	return merchantDocuments, &totalCount, nil
}

func (s *merchantDocumentQueryService) FindByTrashed(ctx context.Context, req *requests.FindAllMerchantDocuments) ([]*db.GetTrashedMerchantDocumentsRow, *int, error) {
	const method = "FindByTrashed"
	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("page", page), attribute.Int("pageSize", pageSize), attribute.String("search", search))
	defer func() { end(status) }()

	if data, total, found := s.cache.GetCachedMerchantDocumentsTrashed(ctx, req); found {
		logSuccess("Successfully retrieved trashed merchant document records from cache", zap.Int("total", *total), zap.Int("page", page), zap.Int("pageSize", pageSize), zap.String("search", search))
		return data, total, nil
	}

	merchantDocuments, err := s.queryRepo.FindByTrashedDocuments(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetTrashedMerchantDocumentsRow](s.logger, err, method, span, zap.String("search", search))
	}

	var totalCount int

	if len(merchantDocuments) > 0 {
		totalCount = int(merchantDocuments[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedMerchantDocumentsTrashed(ctx, req, merchantDocuments, &totalCount)

	logSuccess("Successfully retrieved trashed merchant document records", zap.Int("total", totalCount), zap.Int("page", page), zap.Int("pageSize", pageSize), zap.String("search", search))

	return merchantDocuments, &totalCount, nil
}

func (s *merchantDocumentQueryService) normalizePagination(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return page, pageSize
}
