package withdrawstatsbycardservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	withdraw_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/withdraw_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-withdraw/redis/statsbycard"
	repository "github.com/MamangRust/monolith-payment-gateway-withdraw/repository/statsbycard"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type WithdrawStatsByCardAmountDeps struct {
	Cache mencache.WithdrawStatsByCardAmountCache

	Repository repository.WithdrawStatsByCardAmountRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type withdrawStatsByCardAmountService struct {
	cache mencache.WithdrawStatsByCardAmountCache

	repository repository.WithdrawStatsByCardAmountRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewWithdrawStatsByCardAmountService(deps *WithdrawStatsByCardAmountDeps) WithdrawStatsByCardAmountService {

	return &withdrawStatsByCardAmountService{
		cache:         deps.Cache,
		repository:    deps.Repository,
		logger:        deps.Logger,
		observability: deps.Observability,
	}
}

func (s *withdrawStatsByCardAmountService) FindMonthlyWithdrawsByCardNumber(ctx context.Context, req *requests.YearMonthCardNumber) ([]*db.GetMonthlyWithdrawsByCardNumberRow, error) {
	const method = "FindMonthlyWithdrawsByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetCachedMonthlyWithdrawsByCardNumber(ctx, req); found {
		logSuccess("Successfully fetched monthly withdraws by card number (from cache)",
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetMonthlyWithdrawsByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyWithdrawsByCardNumberRow](
			s.logger,
			withdraw_errors.ErrFailedFindMonthlyWithdraws,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetCachedMonthlyWithdrawsByCardNumber(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly withdraws by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}

func (s *withdrawStatsByCardAmountService) FindYearlyWithdrawsByCardNumber(ctx context.Context, req *requests.YearMonthCardNumber) ([]*db.GetYearlyWithdrawsByCardNumberRow, error) {
	const method = "FindYearlyWithdrawsByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetCachedYearlyWithdrawsByCardNumber(ctx, req); found {
		logSuccess("Successfully fetched yearly withdraws by card number (from cache)",
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyWithdrawsByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyWithdrawsByCardNumberRow](
			s.logger,
			withdraw_errors.ErrFailedFindYearlyWithdraws,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetCachedYearlyWithdrawsByCardNumber(ctx, req, dbRows)

	logSuccess("Successfully fetched yearly withdraws by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}
