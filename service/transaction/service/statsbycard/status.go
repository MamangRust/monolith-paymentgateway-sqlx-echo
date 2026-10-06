package transactionstatsbycardservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	cache "github.com/MamangRust/monolith-payment-gateway-transaction/redis/statsbycard"
	repository "github.com/MamangRust/monolith-payment-gateway-transaction/repository/statsbycard"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type transactionStatsByCardStatusServiceDeps struct {
	Cache cache.TransactionStatsByCardStatusCache

	Repository repository.TransactionStatsByCardStatusRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type transactionStatsByCardStatusService struct {
	cache cache.TransactionStatsByCardStatusCache

	repository repository.TransactionStatsByCardStatusRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewTransactionStatsByCardStatusService(params *transactionStatsByCardStatusServiceDeps) TransactionStatsByCardStatusService {
	return &transactionStatsByCardStatusService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *transactionStatsByCardStatusService) FindMonthTransactionStatusSuccessByCardNumber(ctx context.Context, req *requests.MonthStatusTransactionCardNumber) ([]*db.GetMonthTransactionStatusSuccessCardNumberRow, error) {
	const method = "FindMonthTransactionStatusSuccessByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetMonthTransactionStatusSuccessByCardCache(ctx, req); found {
		logSuccess("Successfully fetched monthly transaction status success by card number (from cache)",
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
			zap.Int("month", req.Month))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetMonthTransactionStatusSuccessByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthTransactionStatusSuccessCardNumberRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthTransactionSuccessByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
		)
	}

	s.cache.SetMonthTransactionStatusSuccessByCardCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly transaction status success by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year),
		zap.Int("month", req.Month))

	return dbRows, nil
}

func (s *transactionStatsByCardStatusService) FindYearlyTransactionStatusSuccessByCardNumber(ctx context.Context, req *requests.YearStatusTransactionCardNumber) ([]*db.GetYearlyTransactionStatusSuccessCardNumberRow, error) {
	const method = "FindYearlyTransactionStatusSuccessByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetYearTransactionStatusSuccessByCardCache(ctx, req); found {
		logSuccess("Successfully fetched yearly transaction status success by card number (from cache)",
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyTransactionStatusSuccessByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTransactionStatusSuccessCardNumberRow](
			s.logger,
			transaction_errors.ErrFailedFindYearTransactionSuccessByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearTransactionStatusSuccessByCardCache(ctx, req, dbRows)

	logSuccess("Successfully fetched yearly transaction status success by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}

func (s *transactionStatsByCardStatusService) FindMonthTransactionStatusFailedByCardNumber(ctx context.Context, req *requests.MonthStatusTransactionCardNumber) ([]*db.GetMonthTransactionStatusFailedCardNumberRow, error) {
	const method = "FindMonthTransactionStatusFailedByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetMonthTransactionStatusFailedByCardCache(ctx, req); found {
		logSuccess("Successfully fetched monthly transaction status failed by card number (from cache)",
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
			zap.Int("month", req.Month))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetMonthTransactionStatusFailedByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthTransactionStatusFailedCardNumberRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthTransactionFailedByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
		)
	}

	s.cache.SetMonthTransactionStatusFailedByCardCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly transaction status failed by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year),
		zap.Int("month", req.Month))

	return dbRows, nil
}

func (s *transactionStatsByCardStatusService) FindYearlyTransactionStatusFailedByCardNumber(ctx context.Context, req *requests.YearStatusTransactionCardNumber) ([]*db.GetYearlyTransactionStatusFailedCardNumberRow, error) {
	const method = "FindYearlyTransactionStatusFailedByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetYearTransactionStatusFailedByCardCache(ctx, req); found {
		logSuccess("Successfully fetched yearly transaction status failed by card number (from cache)",
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyTransactionStatusFailedByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTransactionStatusFailedCardNumberRow](
			s.logger,
			transaction_errors.ErrFailedFindYearTransactionFailedByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearTransactionStatusFailedByCardCache(ctx, req, dbRows)

	logSuccess("Successfully fetched yearly transaction status failed by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}
