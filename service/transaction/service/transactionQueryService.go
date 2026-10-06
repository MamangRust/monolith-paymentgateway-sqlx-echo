package service

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-transaction/redis"
	"github.com/MamangRust/monolith-payment-gateway-transaction/repository"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// TransactionQueryServiceDeps groups dependencies for transaction query service.
type TransactionQueryServiceDeps struct {
	Cache                      mencache.TransactionQueryCache
	TransactionQueryRepository repository.TransactionQueryRepository
	Logger                     logger.LoggerInterface
	Observability              observability.TraceLoggerObservability
}

// transactionQueryService handles transaction read operations.
type transactionQueryService struct {
	cache                      mencache.TransactionQueryCache
	transactionQueryRepository repository.TransactionQueryRepository
	logger                     logger.LoggerInterface
	observability              observability.TraceLoggerObservability
}

func NewTransactionQueryService(
	params *TransactionQueryServiceDeps,
) TransactionQueryService {
	return &transactionQueryService{
		cache:                      params.Cache,
		transactionQueryRepository: params.TransactionQueryRepository,
		logger:                     params.Logger,
		observability:              params.Observability,
	}
}

func (s *transactionQueryService) FindAll(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetTransactionsRow, *int, error) {
	const method = "FindAll"

	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedTransactionsCache(ctx, req); found {
		logSuccess("Successfully fetched card records from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	transactions, err := s.transactionQueryRepository.FindAllTransactions(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetTransactionsRow](
			s.logger,
			transaction_errors.ErrFailedFindAllTransactions,
			method,
			span,

			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search),
		)
	}

	var totalCount int

	if len(transactions) > 0 {
		totalCount = int(transactions[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedTransactionsCache(ctx, req, transactions, &totalCount)

	logSuccess("Successfully fetched transaction",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return transactions, &totalCount, nil
}

func (s *transactionQueryService) FindAllByCardNumber(ctx context.Context, req *requests.FindAllTransactionCardNumber) ([]*db.GetTransactionsByCardNumberRow, *int, error) {
	const method = "FindAllByCardNumber"

	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search),
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedTransactionByCardNumberCache(ctx, req); found {
		logSuccess("Successfully retrieved all transaction records from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	transactions, err := s.transactionQueryRepository.FindAllTransactionByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetTransactionsByCardNumberRow](
			s.logger,
			transaction_errors.ErrFailedFindAllByCardNumber,
			method,
			span,

			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search),
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		)
	}

	var totalCount int

	if len(transactions) > 0 {
		totalCount = int(transactions[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedTransactionByCardNumberCache(ctx, req, transactions, &totalCount)

	logSuccess("Successfully fetched transaction by card number",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)))

	return transactions, &totalCount, nil
}

func (s *transactionQueryService) FindById(ctx context.Context, transactionID int) (*db.GetTransactionByIDRow, error) {
	const method = "FindById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("transaction_id", transactionID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedTransactionCache(ctx, transactionID); found {
		logSuccess("Successfully fetched transaction from cache", zap.Int("transaction.id", transactionID))
		return data, nil
	}

	transaction, err := s.transactionQueryRepository.FindById(ctx, transactionID)

	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.GetTransactionByIDRow](
			s.logger,
			transaction_errors.ErrTransactionNotFound,
			method,
			span,

			zap.Int("transaction_id", transactionID),
		)
	}

	s.cache.SetCachedTransactionCache(ctx, transaction)

	logSuccess("Successfully fetched transaction", zap.Int("transaction_id", transactionID))

	return transaction, nil
}

func (s *transactionQueryService) FindByActive(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetActiveTransactionsRow, *int, error) {
	const method = "FindByActive"

	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", req.Page),
		attribute.Int("pageSize", req.PageSize),
		attribute.String("search", req.Search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedTransactionActiveCache(ctx, req); found {
		logSuccess("Successfully fetched active transaction from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	transactions, err := s.transactionQueryRepository.FindByActive(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetActiveTransactionsRow](
			s.logger,
			transaction_errors.ErrFailedFindByActiveTransactions,
			method,
			span,

			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search),
		)
	}

	var totalCount int

	if len(transactions) > 0 {
		totalCount = int(transactions[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedTransactionActiveCache(ctx, req, transactions, &totalCount)

	logSuccess("Successfully fetched active transaction",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return transactions, &totalCount, nil
}

func (s *transactionQueryService) FindByTrashed(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetTrashedTransactionsRow, *int, error) {
	const method = "FindByTrashed"

	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedTransactionTrashedCache(ctx, req); found {
		logSuccess("Successfully fetched trashed transaction from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	transactions, err := s.transactionQueryRepository.FindByTrashed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetTrashedTransactionsRow](
			s.logger,
			transaction_errors.ErrFailedFindByTrashedTransactions,
			method,
			span,

			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search),
		)
	}

	var totalCount int

	if len(transactions) > 0 {
		totalCount = int(transactions[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedTransactionTrashedCache(ctx, req, transactions, &totalCount)

	logSuccess("Successfully fetched trashed transaction",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return transactions, &totalCount, nil
}

func (s *transactionQueryService) FindTransactionByMerchantId(ctx context.Context, merchant_id int) ([]*db.GetTransactionsByMerchantIDRow, error) {
	const method = "FindTransactionByMerchantId"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("merchant_id", merchant_id))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedTransactionByMerchantIdCache(ctx, merchant_id); found {
		logSuccess("Successfully fetched transaction by merchant ID from cache", zap.Int("merchant.id", merchant_id))
		return data, nil
	}

	res, err := s.transactionQueryRepository.FindTransactionByMerchantId(ctx, merchant_id)

	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetTransactionsByMerchantIDRow](
			s.logger,
			transaction_errors.ErrFailedFindByMerchantID,
			method,
			span,

			zap.Int("merchant_id", merchant_id),
		)
	}
	s.cache.SetCachedTransactionByMerchantIdCache(ctx, merchant_id, res)

	logSuccess("Successfully fetched transaction by merchant ID", zap.Int("merchant_id", merchant_id))

	return res, nil
}

func (s *transactionQueryService) normalizePagination(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return page, pageSize
}
