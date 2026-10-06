package cardstatsbycard

import (
	"context"

	cardstatsmencache "github.com/MamangRust/monolith-payment-gateway-card/redis/statsbycard"
	repository "github.com/MamangRust/monolith-payment-gateway-card/repository/statsbycard"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharederrorhandler "github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type cardStatsWithdrawByCardService struct {
	cache cardstatsmencache.CardStatsWithdrawByCardCache

	repository repository.CardStatsWithdrawByCardRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

type cardStatsWithdrawByCardServiceDeps struct {
	Cache cardstatsmencache.CardStatsWithdrawByCardCache

	Repository repository.CardStatsWithdrawByCardRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

func NewCardStatsWithdrawByCardService(params *cardStatsWithdrawByCardServiceDeps) CardStatsWithdrawByCardService {
	return &cardStatsWithdrawByCardService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *cardStatsWithdrawByCardService) FindMonthlyWithdrawAmountByCardNumber(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetMonthlyWithdrawAmountByCardNumberRow, error) {
	const method = "FindMonthlyWithdrawAmountByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyWithdrawByNumberCache(ctx, req); found {
		logSuccess("Cache hit for monthly withdraw amount card", zap.String("card_number", observability.MaskIdentifier(req.CardNumber)), zap.Int("year", req.Year))
		return data, nil
	}

	res, err := s.repository.GetMonthlyWithdrawAmountByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetMonthlyWithdrawAmountByCardNumberRow](
			s.logger,
			card_errors.ErrFailedFindMonthlyWithdrawAmountByCard,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyWithdrawByNumberCache(ctx, req, res)

	logSuccess("Monthly withdraw amount by card number retrieved successfully",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}

func (s *cardStatsWithdrawByCardService) FindYearlyWithdrawAmountByCardNumber(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetYearlyWithdrawAmountByCardNumberRow, error) {
	const method = "FindYearlyWithdrawAmountByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyWithdrawByNumberCache(ctx, req); found {
		logSuccess("Cache hit for yearly withdraw amount card", zap.String("card_number", observability.MaskIdentifier(req.CardNumber)), zap.Int("year", req.Year))
		return data, nil
	}

	res, err := s.repository.GetYearlyWithdrawAmountByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetYearlyWithdrawAmountByCardNumberRow](
			s.logger,
			card_errors.ErrFailedFindYearlyWithdrawAmountByCard,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearlyWithdrawByNumberCache(ctx, req, res)

	logSuccess("Yearly withdraw amount by card number retrieved successfully",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}
