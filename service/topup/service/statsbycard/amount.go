package topupstatsbycardservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-topup/redis/statsbycard"
	repository "github.com/MamangRust/monolith-payment-gateway-topup/repository/statsbycard"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type topupStatsByCardAmountDeps struct {
	Cache mencache.TopupStatsAmountByCardCache

	Repository repository.TopupStatsByCardAmountRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type topupStatsByCardAmountService struct {
	cache mencache.TopupStatsAmountByCardCache

	repository repository.TopupStatsByCardAmountRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewTopupStatsByCardAmountService(params *topupStatsByCardAmountDeps) TopupStatsByCardAmountService {

	return &topupStatsByCardAmountService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *topupStatsByCardAmountService) FindMonthlyTopupAmountsByCardNumber(ctx context.Context, req *requests.YearMonthMethod) ([]*db.GetMonthlyTopupAmountsByCardNumberRow, error) {
	const method = "FindMonthlyTopupAmountsByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTopupAmountsByCardNumberCache(ctx, req); found {
		logSuccess("Successfully fetched monthly topup amounts by card number", zap.String("card_number", observability.MaskIdentifier(req.CardNumber)), zap.Int("year", req.Year))
		return data, nil
	}

	s.logger.Debug("Cache miss for monthly topup amounts by card number, fetching from DB",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	dbRows, err := s.repository.GetMonthlyTopupAmountsByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTopupAmountsByCardNumberRow](
			s.logger,
			topup_errors.ErrFailedFindMonthlyTopupAmountsByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyTopupAmountsByCardNumberCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly topup amounts by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}

func (s *topupStatsByCardAmountService) FindYearlyTopupAmountsByCardNumber(ctx context.Context, req *requests.YearMonthMethod) ([]*db.GetYearlyTopupAmountsByCardNumberRow, error) {
	const method = "FindYearlyTopupAmountsByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTopupAmountsByCardNumberCache(ctx, req); found {
		logSuccess("Successfully fetched yearly topup amounts by card number", zap.String("card_number", observability.MaskIdentifier(req.CardNumber)), zap.Int("year", req.Year))
		return data, nil
	}

	dbRows, err := s.repository.GetYearlyTopupAmountsByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTopupAmountsByCardNumberRow](
			s.logger,
			topup_errors.ErrFailedFindYearlyTopupAmountsByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.logger.Debug("Setting cache for yearly topup amounts by card number",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	s.cache.SetYearlyTopupAmountsByCardNumberCache(ctx, req, dbRows)

	logSuccess("Successfully fetched yearly topup amounts by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}
