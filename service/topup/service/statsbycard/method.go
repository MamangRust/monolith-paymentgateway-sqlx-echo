package topupstatsbycardservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	cache "github.com/MamangRust/monolith-payment-gateway-topup/redis/statsbycard"
	repository "github.com/MamangRust/monolith-payment-gateway-topup/repository/statsbycard"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type topupStatsByCardMethodDeps struct {
	Cache cache.TopupStatsMethodByCardCache

	Repository repository.TopupStatsByCardMethodRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type topupStatsByCardMethodService struct {
	cache cache.TopupStatsMethodByCardCache

	repository repository.TopupStatsByCardMethodRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewTopupStatsByCardMethodService(params *topupStatsByCardMethodDeps) TopupStatsByCardMethodService {

	return &topupStatsByCardMethodService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *topupStatsByCardMethodService) FindMonthlyTopupMethodsByCardNumber(ctx context.Context, req *requests.YearMonthMethod) ([]*db.GetMonthlyTopupMethodsByCardNumberRow, error) {
	const method = "FindMonthlyTopupMethodsByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTopupMethodsByCardNumberCache(ctx, req); found {
		logSuccess("Successfully fetched monthly topup methods by card number", zap.String("card_number", observability.MaskIdentifier(req.CardNumber)), zap.Int("year", req.Year))
		return data, nil
	}

	s.logger.Debug("Cache miss for monthly topup methods by card number, fetching from DB",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	dbRows, err := s.repository.GetMonthlyTopupMethodsByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTopupMethodsByCardNumberRow](
			s.logger,
			topup_errors.ErrFailedFindMonthlyTopupMethodsByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyTopupMethodsByCardNumberCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly topup methods by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}

func (s *topupStatsByCardMethodService) FindYearlyTopupMethodsByCardNumber(ctx context.Context, req *requests.YearMonthMethod) ([]*db.GetYearlyTopupMethodsByCardNumberRow, error) {
	const method = "FindYearlyTopupMethodsByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTopupMethodsByCardNumberCache(ctx, req); found {
		logSuccess("Successfully fetched yearly topup methods by card number", zap.String("card_number", observability.MaskIdentifier(req.CardNumber)), zap.Int("year", req.Year))
		return data, nil
	}

	s.logger.Debug("Cache miss for yearly topup methods by card number, fetching from DB",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	dbRows, err := s.repository.GetYearlyTopupMethodsByCardNumber(ctx, req)

	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTopupMethodsByCardNumberRow](
			s.logger,
			topup_errors.ErrFailedFindYearlyTopupMethodsByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.logger.Debug("Setting cache for yearly topup methods by card number",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	s.cache.SetYearlyTopupMethodsByCardNumberCache(ctx, req, dbRows)

	logSuccess("Successfully fetched yearly topup methods by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}
