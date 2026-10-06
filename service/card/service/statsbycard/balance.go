package cardstatsbycard

import (
	"context"

	cardstatsmencache "github.com/MamangRust/monolith-payment-gateway-card/redis/statsbycard"
	repository "github.com/MamangRust/monolith-payment-gateway-card/repository/statsbycard"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"

	sharederrorhandler "github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
)

type cardStatsBalanceByCardService struct {
	cache cardstatsmencache.CardStatsBalanceByCardCache

	repository repository.CardStatsBalanceByCardRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

type cardStatsBalanceByCardServiceDeps struct {
	Cache cardstatsmencache.CardStatsBalanceByCardCache

	Repository repository.CardStatsBalanceByCardRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

func NewCardStatsBalanceByCardService(params *cardStatsBalanceByCardServiceDeps) CardStatsBalanceByCardService {
	return &cardStatsBalanceByCardService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *cardStatsBalanceByCardService) FindMonthlyBalancesByCardNumber(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetMonthlyBalancesByCardNumberRow, error) {
	const method = "FindMonthlyBalancesByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyBalanceByNumberCache(ctx, req); found {
		logSuccess("Cache hit for monthly balance card", zap.Int("year", req.Year))
		return data, nil
	}

	res, err := s.repository.GetMonthlyBalancesByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetMonthlyBalancesByCardNumberRow](
			s.logger,
			card_errors.ErrFailedFindMonthlyBalanceByCard,
			method,
			span,
			zap.Int("year", req.Year),
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		)
	}

	s.cache.SetMonthlyBalanceByNumberCache(ctx, req, res)

	logSuccess("Monthly balance retrieved successfully",
		zap.Int("year", req.Year),
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}

func (s *cardStatsBalanceByCardService) FindYearlyBalanceByCardNumber(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetYearlyBalancesByCardNumberRow, error) {
	const method = "FindYearlyBalanceByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyBalanceByNumberCache(ctx, req); found {
		logSuccess("Cache hit for yearly balance card", zap.Int("year", req.Year))
		return data, nil
	}

	res, err := s.repository.GetYearlyBalanceByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetYearlyBalancesByCardNumberRow](
			s.logger,
			card_errors.ErrFailedFindYearlyBalanceByCard,
			method,
			span,
			zap.Int("year", req.Year),
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		)
	}

	s.cache.SetYearlyBalanceByNumberCache(ctx, req, res)

	logSuccess("Yearly balance retrieved successfully",
		zap.Int("year", req.Year),
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}
