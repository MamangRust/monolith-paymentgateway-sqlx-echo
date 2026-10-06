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

// merchantTransactionDeps holds dependencies for merchant transaction operations.
type MerchantTransactionDeps struct {
	Repository    repository.MerchantTransactionRepository
	Cache         mencache.MerchantTransactionCache
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
}

// merchantTransactionService handles merchant transaction operations.
type merchantTransactionService struct {
	repo          repository.MerchantTransactionRepository
	cache         mencache.MerchantTransactionCache
	logger        logger.LoggerInterface
	observability observability.TraceLoggerObservability
}

// NewMerchantTransactionService constructs a MerchantTransactionService.
func NewMerchantTransactionService(
	params *MerchantTransactionDeps,
) MerchantTransactionService {
	return &merchantTransactionService{
		repo:          params.Repository,
		cache:         params.Cache,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *merchantTransactionService) FindAllTransactions(ctx context.Context, req *requests.FindAllMerchantTransactions) ([]*db.FindAllTransactionsRow, *int, error) {
	const method = "FindAllTransactions"
	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("page", page), attribute.Int("pageSize", pageSize), attribute.String("search", search))
	defer func() { end(status) }()

	if data, total, found := s.cache.GetCacheAllMerchantTransactions(ctx, req); found {
		logSuccess("Successfully retrieved all merchant transactions from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	transactions, err := s.repo.FindAllTransactions(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.FindAllTransactionsRow](s.logger, err, method, span, zap.String("search", search))
	}

	var totalCount int
	if len(transactions) > 0 {
		totalCount = int(transactions[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCacheAllMerchantTransactions(ctx, req, transactions, &totalCount)

	logSuccess("Successfully retrieved all merchant transactions", zap.Int("totalRecords", totalCount), zap.Int("page", page), zap.Int("pageSize", pageSize))

	return transactions, &totalCount, nil
}

func (s *merchantTransactionService) FindAllTransactionsByMerchant(ctx context.Context, req *requests.FindAllMerchantTransactionsById) ([]*db.FindAllTransactionsByMerchantRow, *int, error) {
	const method = "FindAllTransactionsByMerchant"
	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("page", page), attribute.Int("pageSize", pageSize), attribute.String("search", search))
	defer func() { end(status) }()

	if data, total, found := s.cache.GetCacheMerchantTransactions(ctx, req); found {
		logSuccess("Successfully retrieved merchant transactions from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	transactions, err := s.repo.FindAllTransactionsByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.FindAllTransactionsByMerchantRow](s.logger, err, method, span, zap.String("search", search))
	}

	var totalCount int
	if len(transactions) > 0 {
		totalCount = int(transactions[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCacheMerchantTransactions(ctx, req, transactions, &totalCount)

	logSuccess("Successfully retrieved merchant transactions", zap.Int("totalRecords", totalCount), zap.Int("page", page), zap.Int("pageSize", pageSize))

	return transactions, &totalCount, nil
}

func (s *merchantTransactionService) FindAllTransactionsByApikey(ctx context.Context, req *requests.FindAllMerchantTransactionsByApiKey) ([]*db.FindAllTransactionsByApikeyRow, *int, error) {
	const method = "FindAllTransactionsByApikey"
	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("page", page), attribute.Int("pageSize", pageSize), attribute.String("search", search))
	defer func() { end(status) }()

	if data, total, found := s.cache.GetCacheMerchantTransactionApikey(ctx, req); found {
		logSuccess("Successfully retrieved merchant transactions by apikey from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	transactions, err := s.repo.FindAllTransactionsByApikey(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.FindAllTransactionsByApikeyRow](s.logger, err, method, span, zap.String("search", search))
	}

	var totalCount int
	if len(transactions) > 0 {
		totalCount = int(transactions[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCacheMerchantTransactionApikey(ctx, req, transactions, &totalCount)

	logSuccess("Successfully retrieved merchant transactions by apikey", zap.Int("totalRecords", totalCount), zap.Int("page", page), zap.Int("pageSize", pageSize))

	return transactions, &totalCount, nil
}

func (s *merchantTransactionService) normalizePagination(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return page, pageSize
}
