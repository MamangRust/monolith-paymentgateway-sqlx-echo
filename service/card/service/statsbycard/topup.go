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

type cardStatsTopupByCardService struct {
	cache cardstatsmencache.CardStatsTopupByCardCache

	repository repository.CardStatsTopupByCardRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

type cardStatsTopupByCardServiceDeps struct {
	Cache cardstatsmencache.CardStatsTopupByCardCache

	Repository repository.CardStatsTopupByCardRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

func NewCardStatsTopupByCardService(params *cardStatsTopupByCardServiceDeps) CardStatsTopupByCardService {

	return &cardStatsTopupByCardService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *cardStatsTopupByCardService) FindMonthlyTopupAmountByCardNumber(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetMonthlyTopupAmountByCardNumberRow, error) {
	const method = "FindMonthlyTopupAmountByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTopupByNumberCache(ctx, req); found {
		logSuccess("Cache hit for monthly topup amount card", zap.String("card_number", observability.MaskIdentifier(req.CardNumber)), zap.Int("year", req.Year))
		return data, nil
	}

	res, err := s.repository.GetMonthlyTopupAmountByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetMonthlyTopupAmountByCardNumberRow](
			s.logger,
			card_errors.ErrFailedFindMonthlyTopupAmountByCard,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyTopupByNumberCache(ctx, req, res)

	logSuccess("Monthly topup amount by card number retrieved successfully",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}

func (s *cardStatsTopupByCardService) FindYearlyTopupAmountByCardNumber(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetYearlyTopupAmountByCardNumberRow, error) {
	const method = "FindYearlyTopupAmountByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTopupByNumberCache(ctx, req); found {
		logSuccess("Cache hit for yearly topup amount card", zap.String("card_number", observability.MaskIdentifier(req.CardNumber)), zap.Int("year", req.Year))
		return data, nil
	}

	res, err := s.repository.GetYearlyTopupAmountByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetYearlyTopupAmountByCardNumberRow](
			s.logger,
			card_errors.ErrFailedFindYearlyTopupAmountByCard,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearlyTopupByNumberCache(ctx, req, res)

	logSuccess("Yearly topup amount by card number retrieved successfully",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}
