package cardstatsservice

import (
	"context"

	cardstatsmencache "github.com/MamangRust/monolith-payment-gateway-card/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-card/repository/stats"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	sharederrorhandler "github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type cardStatsTransferService struct {
	cache cardstatsmencache.CardStatsTransferCache

	repository repository.CardStatsTransferRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

type cardStatsTransferServiceDeps struct {
	Cache cardstatsmencache.CardStatsTransferCache

	Repository repository.CardStatsTransferRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

func NewCardStatsTransferService(params *cardStatsTransferServiceDeps) CardStatsTransferService {

	return &cardStatsTransferService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *cardStatsTransferService) FindMonthlyTransferAmountSender(ctx context.Context, year int) ([]*db.GetMonthlyTransferAmountSenderRow, error) {
	const method = "FindMonthlyTransferAmountSender"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTransferSenderCache(ctx, year); found {
		logSuccess("Monthly transfer sender amount cache hit", zap.Int("year", year))
		return data, nil
	}

	res, err := s.repository.GetMonthlyTransferAmountSender(ctx, year)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetMonthlyTransferAmountSenderRow](
			s.logger,
			card_errors.ErrFailedFindMonthlyTransferAmountSender,
			method,
			span,
			zap.Int("year", year),
		)
	}

	s.cache.SetMonthlyTransferSenderCache(ctx, year, res)

	logSuccess("Monthly transfer sender amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}

func (s *cardStatsTransferService) FindYearlyTransferAmountSender(ctx context.Context, year int) ([]*db.GetYearlyTransferAmountSenderRow, error) {
	const method = "FindYearlyTransferAmountSender"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTransferSenderCache(ctx, year); found {
		logSuccess("Yearly transfer sender amount cache hit", zap.Int("year", year))
		return data, nil
	}

	res, err := s.repository.GetYearlyTransferAmountSender(ctx, year)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetYearlyTransferAmountSenderRow](
			s.logger,
			card_errors.ErrFailedFindYearlyTransferAmountSender,
			method,
			span,
			zap.Int("year", year),
		)
	}

	s.cache.SetYearlyTransferSenderCache(ctx, year, res)

	logSuccess("Yearly transfer sender amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}

func (s *cardStatsTransferService) FindMonthlyTransferAmountReceiver(ctx context.Context, year int) ([]*db.GetMonthlyTransferAmountReceiverRow, error) {
	const method = "FindMonthlyTransferAmountReceiver"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTransferReceiverCache(ctx, year); found {
		logSuccess("Monthly transfer receiver amount cache hit", zap.Int("year", year))
		return data, nil
	}

	res, err := s.repository.GetMonthlyTransferAmountReceiver(ctx, year)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetMonthlyTransferAmountReceiverRow](
			s.logger,
			card_errors.ErrFailedFindMonthlyTransferAmountReceiver,
			method,
			span,
			zap.Int("year", year),
		)
	}

	s.cache.SetMonthlyTransferReceiverCache(ctx, year, res)

	logSuccess("Monthly transfer receiver amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}

func (s *cardStatsTransferService) FindYearlyTransferAmountReceiver(ctx context.Context, year int) ([]*db.GetYearlyTransferAmountReceiverRow, error) {
	const method = "FindYearlyTransferAmountReceiver"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTransferReceiverCache(ctx, year); found {
		logSuccess("Yearly transfer receiver amount cache hit", zap.Int("year", year))
		return data, nil
	}

	res, err := s.repository.GetYearlyTransferAmountReceiver(ctx, year)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetYearlyTransferAmountReceiverRow](
			s.logger,
			card_errors.ErrFailedFindYearlyTransferAmountReceiver,
			method,
			span,
			zap.Int("year", year),
		)
	}

	s.cache.SetYearlyTransferReceiverCache(ctx, year, res)

	logSuccess("Yearly transfer receiver amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}
