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
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// merchantQueryDeps holds dependencies for merchant queries.
type MerchantQueryDeps struct {
	Repository    repository.MerchantQueryRepository
	Cache         cache.MerchantQueryCache
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
}

// merchantQueryService handles read operations for merchants.
type merchantQueryService struct {
	queryRepo     repository.MerchantQueryRepository
	cache         cache.MerchantQueryCache
	logger        logger.LoggerInterface
	observability observability.TraceLoggerObservability
}

// NewMerchantQueryService constructs a MerchantQueryService.
func NewMerchantQueryService(params *MerchantQueryDeps) MerchantQueryService {
	return &merchantQueryService{
		queryRepo:     params.Repository,
		cache:         params.Cache,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *merchantQueryService) FindAll(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetMerchantsRow, *int, error) {
	const method = "FindAll"
	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("page", page), attribute.Int("pageSize", pageSize), attribute.String("search", search))
	defer func() { end(status) }()

	if data, total, found := s.cache.GetCachedMerchants(ctx, req); found {
		logSuccess("Successfully retrieved all merchant records from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	merchants, err := s.queryRepo.FindAllMerchants(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetMerchantsRow](s.logger, err, method, span, zap.String("search", search))
	}

	var totalCount int
	if len(merchants) > 0 {
		totalCount = int(merchants[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedMerchants(ctx, req, merchants, &totalCount)

	logSuccess("Successfully retrieved all merchant records", zap.Int("totalRecords", totalCount), zap.Int("page", page), zap.Int("pageSize", pageSize))

	return merchants, &totalCount, nil
}

func (s *merchantQueryService) FindById(ctx context.Context, merchantID int) (*db.GetMerchantByIDRow, error) {
	const method = "FindById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("merchant.id", merchantID))
	defer func() { end(status) }()

	if cachedMerchant, found := s.cache.GetCachedMerchant(ctx, merchantID); found {
		logSuccess("Successfully retrieved merchant from cache", zap.Int("merchant.id", merchantID))
		return cachedMerchant, nil
	}

	merchant, err := s.queryRepo.FindByMerchantId(ctx, merchantID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.GetMerchantByIDRow](s.logger, err, method, span, zap.Int("merchant.id", merchantID))
	}

	s.cache.SetCachedMerchant(ctx, merchant)

	logSuccess("Successfully retrieved merchant", zap.Int("merchant.id", merchantID))

	return merchant, nil
}

func (s *merchantQueryService) FindByActive(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetActiveMerchantsRow, *int, error) {
	const method = "FindByActive"
	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("page", page), attribute.Int("pageSize", pageSize), attribute.String("search", search))
	defer func() { end(status) }()

	if data, total, found := s.cache.GetCachedMerchantActive(ctx, req); found {
		logSuccess("Successfully fetched active merchants from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	merchants, err := s.queryRepo.FindByActive(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetActiveMerchantsRow](s.logger, err, method, span, zap.String("search", search))
	}

	var totalCount int
	if len(merchants) > 0 {
		totalCount = int(merchants[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedMerchantActive(ctx, req, merchants, &totalCount)

	logSuccess("Successfully fetched active merchants", zap.Int("totalRecords", totalCount), zap.Int("page", page), zap.Int("pageSize", pageSize))

	return merchants, &totalCount, nil
}

func (s *merchantQueryService) FindByTrashed(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetTrashedMerchantsRow, *int, error) {
	const method = "FindByTrashed"
	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("page", page), attribute.Int("pageSize", pageSize), attribute.String("search", search))
	defer func() { end(status) }()

	if data, total, found := s.cache.GetCachedMerchantTrashed(ctx, req); found {
		logSuccess("Successfully fetched trashed merchants from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	merchants, err := s.queryRepo.FindByTrashed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetTrashedMerchantsRow](s.logger, err, method, span, zap.String("search", search))
	}

	var totalCount int
	if len(merchants) > 0 {
		totalCount = int(merchants[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedMerchantTrashed(ctx, req, merchants, &totalCount)

	logSuccess("Successfully fetched trashed merchants", zap.Int("totalRecords", totalCount), zap.Int("page", page), zap.Int("pageSize", pageSize))

	return merchants, &totalCount, nil
}

func (s *merchantQueryService) FindByApiKey(ctx context.Context, apiKey string) (*db.GetMerchantByApiKeyRow, error) {
	const method = "FindByApiKey"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.String("api_key", observability.MaskIdentifier(apiKey)))
	defer func() { end(status) }()

	if cachedMerchant, found := s.cache.GetCachedMerchantByApiKey(ctx, apiKey); found {
		logSuccess("Successfully found merchant by API key from cache", zap.String("api_key", observability.MaskIdentifier(apiKey)))
		return cachedMerchant, nil
	}

	merchant, err := s.queryRepo.FindByApiKey(ctx, apiKey)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.GetMerchantByApiKeyRow](s.logger, err, method, span, zap.String("api_key", observability.MaskIdentifier(apiKey)))
	}

	s.cache.SetCachedMerchantByApiKey(ctx, apiKey, merchant)

	logSuccess("Successfully found merchant by API key", zap.String("api_key", observability.MaskIdentifier(apiKey)))

	return merchant, nil
}

func (s *merchantQueryService) FindByMerchantUserId(ctx context.Context, userID int) ([]*db.GetMerchantsByUserIDRow, error) {
	const method = "FindByMerchantUserId"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("user.id", userID))
	defer func() { end(status) }()

	if cachedMerchants, found := s.cache.GetCachedMerchantsByUserId(ctx, userID); found {
		logSuccess("Successfully found merchants by user ID from cache", zap.Int("user.id", userID), zap.Int("count", len(cachedMerchants)))
		return cachedMerchants, nil
	}

	merchants, err := s.queryRepo.FindByMerchantUserId(ctx, userID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMerchantsByUserIDRow](s.logger, err, method, span, zap.Int("user.id", userID))
	}

	s.cache.SetCachedMerchantsByUserId(ctx, userID, merchants)

	logSuccess("Successfully found merchants by user ID", zap.Int("user.id", userID), zap.Int("count", len(merchants)))

	return merchants, nil
}

func (s *merchantQueryService) normalizePagination(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return page, pageSize
}
