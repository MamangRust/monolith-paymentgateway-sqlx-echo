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

type cardStatsTransferByCardService struct {
	cache cardstatsmencache.CardStatsTransferByCardCache

	repository repository.CardStatsTransferByCardRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

type cardStatsTransferByCardServiceDeps struct {
	Cache cardstatsmencache.CardStatsTransferByCardCache

	Repository repository.CardStatsTransferByCardRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

func NewCardStatsTransferByCardService(params *cardStatsTransferByCardServiceDeps) CardStatsTransferByCardService {
	return &cardStatsTransferByCardService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *cardStatsTransferByCardService) FindMonthlyTransferAmountBySender(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetMonthlyTransferAmountBySenderRow, error) {
	const method = "FindMonthlyTransferAmountBySender"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTransferBySenderCache(ctx, req); found {
		logSuccess("Cache hit for monthly transfer sender amount card", zap.String("card_number", observability.MaskIdentifier(req.CardNumber)), zap.Int("year", req.Year))
		return data, nil
	}

	res, err := s.repository.GetMonthlyTransferAmountBySender(ctx, req)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetMonthlyTransferAmountBySenderRow](
			s.logger,
			card_errors.ErrFailedFindMonthlyTransferAmountBySender,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyTransferBySenderCache(ctx, req, res)

	logSuccess("Monthly transfer sender amount by card number retrieved successfully",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}

func (s *cardStatsTransferByCardService) FindYearlyTransferAmountBySender(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetYearlyTransferAmountBySenderRow, error) {
	const method = "FindYearlyTransferAmountBySender"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTransferBySenderCache(ctx, req); found {
		logSuccess("Cache hit for yearly transfer sender amount card", zap.String("card_number", observability.MaskIdentifier(req.CardNumber)), zap.Int("year", req.Year))
		return data, nil
	}

	res, err := s.repository.GetYearlyTransferAmountBySender(ctx, req)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetYearlyTransferAmountBySenderRow](
			s.logger,
			card_errors.ErrFailedFindYearlyTransferAmountBySender,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearlyTransferBySenderCache(ctx, req, res)

	logSuccess("Yearly transfer sender amount by card number retrieved successfully",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}

func (s *cardStatsTransferByCardService) FindMonthlyTransferAmountByReceiver(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetMonthlyTransferAmountByReceiverRow, error) {
	const method = "FindMonthlyTransferAmountByReceiver"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTransferByReceiverCache(ctx, req); found {
		logSuccess("Cache hit for monthly transfer receiver amount card", zap.String("card_number", observability.MaskIdentifier(req.CardNumber)), zap.Int("year", req.Year))
		return data, nil
	}

	res, err := s.repository.GetMonthlyTransferAmountByReceiver(ctx, req)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetMonthlyTransferAmountByReceiverRow](
			s.logger,
			card_errors.ErrFailedFindMonthlyTransferAmountByReceiver,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyTransferByReceiverCache(ctx, req, res)

	logSuccess("Monthly transfer receiver amount by card number retrieved successfully",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}

func (s *cardStatsTransferByCardService) FindYearlyTransferAmountByReceiver(ctx context.Context, req *requests.MonthYearCardNumberCard) ([]*db.GetYearlyTransferAmountByReceiverRow, error) {
	const method = "FindYearlyTransferAmountByReceiver"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTransferByReceiverCache(ctx, req); found {
		logSuccess("Cache hit for yearly transfer receiver amount card", zap.String("card_number", observability.MaskIdentifier(req.CardNumber)), zap.Int("year", req.Year))
		return data, nil
	}

	res, err := s.repository.GetYearlyTransferAmountByReceiver(ctx, req)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetYearlyTransferAmountByReceiverRow](
			s.logger,
			card_errors.ErrFailedFindYearlyTransferAmountByReceiver,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearlyTransferByReceiverCache(ctx, req, res)

	logSuccess("Yearly transfer receiver amount by card number retrieved successfully",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}
