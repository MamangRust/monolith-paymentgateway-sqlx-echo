package transactionstatsbycardservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-transaction/redis/statsbycard"
	repository "github.com/MamangRust/monolith-payment-gateway-transaction/repository/statsbycard"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type transactionStatsByCardAmountServiceDeps struct {
	Cache mencache.TransactionStatsByCardAmountCache

	Repository repository.TransactionStatsByCardAmountRepository

	Observability observability.TraceLoggerObservability

	Logger logger.LoggerInterface
}

type transactionStatsByCardAmountService struct {
	cache         mencache.TransactionStatsByCardAmountCache
	repository    repository.TransactionStatsByCardAmountRepository
	logger        logger.LoggerInterface
	observability observability.TraceLoggerObservability
}

func NewTransactionStatsByCardAmountService(params *transactionStatsByCardAmountServiceDeps) TransactionStatsByCardAmountService {
	return &transactionStatsByCardAmountService{
		cache:         params.Cache,
		logger:        params.Logger,
		repository:    params.Repository,
		observability: params.Observability,
	}
}

func (s *transactionStatsByCardAmountService) FindMonthlyAmountsByCardNumber(ctx context.Context, req *requests.MonthYearPaymentMethod) ([]*db.GetMonthlyAmountsByCardNumberRow, error) {
	const method = "FindMonthlyAmountsByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetMonthlyAmountsByCardCache(ctx, req); found {
		logSuccess("Successfully fetched monthly amounts by card number (from cache)",
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetMonthlyAmountsByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyAmountsByCardNumberRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthlyAmountsByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyAmountsByCardCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly amounts by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}

func (s *transactionStatsByCardAmountService) FindYearlyAmountsByCardNumber(ctx context.Context, req *requests.MonthYearPaymentMethod) ([]*db.GetYearlyAmountsByCardNumberRow, error) {
	const method = "FindYearlyAmountsByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetYearlyAmountsByCardCache(ctx, req); found {
		logSuccess("Successfully fetched yearly amounts by card number (from cache)",
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyAmountsByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyAmountsByCardNumberRow](
			s.logger,
			transaction_errors.ErrFailedFindYearlyAmountsByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearlyAmountsByCardCache(ctx, req, dbRows)

	logSuccess("Successfully fetched yearly amounts by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}
