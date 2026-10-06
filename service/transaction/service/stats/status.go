package transactionstatsservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-transaction/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-transaction/repository/stats"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type transactionStatsStatusServiceDeps struct {
	Cache mencache.TransactionStatsStatusCache

	Repository repository.TransactionStatsStatusRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type transactionStatsStatusService struct {
	cache mencache.TransactionStatsStatusCache

	repository repository.TransactionStatsStatusRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewTransactionStatsStatusService(params *transactionStatsStatusServiceDeps) TransactionStatsStatusService {
	return &transactionStatsStatusService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *transactionStatsStatusService) FindMonthTransactionStatusSuccess(ctx context.Context, req *requests.MonthStatusTransaction) ([]*db.GetMonthTransactionStatusSuccessRow, error) {
	const method = "FindMonthTransactionStatusSuccess"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetMonthTransactionStatusSuccessCache(ctx, req); found {
		logSuccess("Successfully fetched monthly transaction status success (from cache)", zap.Int("year", req.Year), zap.Int("month", req.Month))
		return dbRows, nil
	}

	s.logger.Debug("Cache miss for monthly transaction status success, fetching from DB", zap.Int("year", req.Year), zap.Int("month", req.Month))

	dbRows, err := s.repository.GetMonthTransactionStatusSuccess(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthTransactionStatusSuccessRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthTransactionSuccess,
			method,
			span,

			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
		)
	}

	s.cache.SetMonthTransactionStatusSuccessCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly transaction status success (from DB)", zap.Int("year", req.Year), zap.Int("month", req.Month))

	return dbRows, nil
}

func (s *transactionStatsStatusService) FindYearlyTransactionStatusSuccess(ctx context.Context, year int) ([]*db.GetYearlyTransactionStatusSuccessRow, error) {
	const method = "FindYearlyTransactionStatusSuccess"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Checking cache for yearly transaction status success", zap.Int("year", year))
	dbRows, found := s.cache.GetYearTransactionStatusSuccessCache(ctx, year)
	if found {
		s.logger.Info("Cache hit for yearly transaction status success", zap.Int("year", year))
		status = "ok"
		logSuccess("Successfully fetched yearly transaction status success (from cache)", zap.Int("year", year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyTransactionStatusSuccess(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTransactionStatusSuccessRow](
			s.logger,
			transaction_errors.ErrFailedFindYearTransactionSuccess,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetYearTransactionStatusSuccessCache(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly transaction status success (from DB)", zap.Int("year", year))

	return dbRows, nil
}

func (s *transactionStatsStatusService) FindMonthTransactionStatusFailed(ctx context.Context, req *requests.MonthStatusTransaction) ([]*db.GetMonthTransactionStatusFailedRow, error) {
	const method = "FindMonthTransactionStatusFailed"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetMonthTransactionStatusFailedCache(ctx, req); found {
		logSuccess("Successfully fetched monthly transaction status failed (from cache)", zap.Int("year", req.Year), zap.Int("month", req.Month))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetMonthTransactionStatusFailed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthTransactionStatusFailedRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthTransactionFailed,
			method,
			span,

			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
		)
	}

	s.cache.SetMonthTransactionStatusFailedCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly transaction status failed (from DB)", zap.Int("year", req.Year), zap.Int("month", req.Month))

	return dbRows, nil
}

func (s *transactionStatsStatusService) FindYearlyTransactionStatusFailed(ctx context.Context, year int) ([]*db.GetYearlyTransactionStatusFailedRow, error) {
	const method = "FindYearlyTransactionStatusFailed"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetYearTransactionStatusFailedCache(ctx, year); found {
		s.logger.Info("Cache hit for yearly transaction status failed", zap.Int("year", year))
		status = "ok"
		logSuccess("Successfully fetched yearly transaction status failed (from cache)", zap.Int("year", year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyTransactionStatusFailed(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTransactionStatusFailedRow](
			s.logger,
			transaction_errors.ErrFailedFindYearTransactionFailed,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetYearTransactionStatusFailedCache(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly transaction status failed (from DB)", zap.Int("year", year))

	return dbRows, nil
}
