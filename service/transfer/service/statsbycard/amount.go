package transferstatsbycardservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	transfer_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transfer_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-transfer/redis/statsbycard"
	repository "github.com/MamangRust/monolith-payment-gateway-transfer/repository/statsbycard"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type TransferStatsByCardAmountDeps struct {
	Cache mencache.TransferStatsByCardAmountCache

	Sender repository.TransferStatsByCardAmountSenderRepository

	Receiver repository.TransferStatsByCardAmountReceiverRepository

	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
}

type transferStatsByCardAmountService struct {
	cache mencache.TransferStatsByCardAmountCache

	sender repository.TransferStatsByCardAmountSenderRepository

	receiver repository.TransferStatsByCardAmountReceiverRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewTransferStatsByCardAmountService(params *TransferStatsByCardAmountDeps) TransferStatsByCardAmountService {
	return &transferStatsByCardAmountService{
		cache:         params.Cache,
		sender:        params.Sender,
		receiver:      params.Receiver,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *transferStatsByCardAmountService) FindMonthlyTransferAmountsBySenderCardNumber(ctx context.Context, req *requests.MonthYearCardNumber) ([]*db.GetMonthlyTransferAmountsBySenderCardNumberRow, error) {
	const method = "FindMonthlyTransferAmountsBySenderCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetMonthlyTransferAmountsBySenderCard(ctx, req); found {
		logSuccess("Successfully fetched monthly transfer amounts by sender card number (from cache)",
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year))
		return dbRows, nil
	}

	dbRows, err := s.sender.GetMonthlyTransferAmountsBySenderCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTransferAmountsBySenderCardNumberRow](
			s.logger,
			transfer_errors.ErrFailedFindMonthlyTransferAmountsBySenderCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyTransferAmountsBySenderCard(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly transfer amounts by sender card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}

func (s *transferStatsByCardAmountService) FindMonthlyTransferAmountsByReceiverCardNumber(ctx context.Context, req *requests.MonthYearCardNumber) ([]*db.GetMonthlyTransferAmountsByReceiverCardNumberRow, error) {
	const method = "FindMonthlyTransferAmountsByReceiverCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetMonthlyTransferAmountsByReceiverCard(ctx, req); found {
		logSuccess("Successfully fetched monthly transfer amounts by receiver card number (from cache)",
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year))
		return dbRows, nil
	}

	dbRows, err := s.receiver.GetMonthlyTransferAmountsByReceiverCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTransferAmountsByReceiverCardNumberRow](
			s.logger,
			transfer_errors.ErrFailedFindMonthlyTransferAmountsByReceiverCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyTransferAmountsByReceiverCard(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly transfer amounts by receiver card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}

func (s *transferStatsByCardAmountService) FindYearlyTransferAmountsBySenderCardNumber(ctx context.Context, req *requests.MonthYearCardNumber) ([]*db.GetYearlyTransferAmountsBySenderCardNumberRow, error) {
	const method = "FindYearlyTransferAmountsBySenderCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetYearlyTransferAmountsBySenderCard(ctx, req); found {
		logSuccess("Successfully fetched yearly transfer amounts by sender card number (from cache)",
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year))
		return dbRows, nil
	}

	dbRows, err := s.sender.GetYearlyTransferAmountsBySenderCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTransferAmountsBySenderCardNumberRow](
			s.logger,
			transfer_errors.ErrFailedFindYearlyTransferAmountsBySenderCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearlyTransferAmountsBySenderCard(ctx, req, dbRows)

	logSuccess("Successfully fetched yearly transfer amounts by sender card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}

func (s *transferStatsByCardAmountService) FindYearlyTransferAmountsByReceiverCardNumber(ctx context.Context, req *requests.MonthYearCardNumber) ([]*db.GetYearlyTransferAmountsByReceiverCardNumberRow, error) {
	const method = "FindYearlyTransferAmountsByReceiverCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetYearlyTransferAmountsByReceiverCard(ctx, req); found {
		logSuccess("Successfully fetched yearly transfer amounts by receiver card number (from cache)",
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year))
		return dbRows, nil
	}

	dbRows, err := s.receiver.GetYearlyTransferAmountsByReceiverCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTransferAmountsByReceiverCardNumberRow](
			s.logger,
			transfer_errors.ErrFailedFindYearlyTransferAmountsByReceiverCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearlyTransferAmountsByReceiverCard(ctx, req, dbRows)

	logSuccess("Successfully fetched yearly transfer amounts by receiver card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}
