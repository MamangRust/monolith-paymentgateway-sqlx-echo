package service

import (
	"context"

	mencache "github.com/MamangRust/monolith-payment-gateway-card/redis/dashboard"
	repository "github.com/MamangRust/monolith-payment-gateway-card/repository/dashboard"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	sharederrorhandler "github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// cardDashboardDeps defines dependencies for cardDashboardService.
type cardDashboardDeps struct {
	Cache                   mencache.CardDashboardCache
	CardDashboardRepository repository.CardDashboardRepository
	Logger                  logger.LoggerInterface
	Observability           observability.TraceLoggerObservability
}

// cardDashboardService implements CardDashboardService.
type cardDashboardService struct {
	cache                   mencache.CardDashboardCache
	cardDashboardRepository repository.CardDashboardRepository
	logger                  logger.LoggerInterface
	observability           observability.TraceLoggerObservability
}

// NewCardDashboardService initializes a new instance of cardDashboardService.
//
// It sets up Prometheus metrics for tracking request counts and durations,
// and registers them for monitoring. This constructor function requires a set
// of parameters encapsulated in cardDashboardDeps, which include the context,
// error handler, cache, repository, logger, and mapper.
//
// Parameters:
//   - params: A pointer to cardDashboardDeps containing the dependencies
//     needed to initialize the service.
//
// Returns:
//   - A pointer to a cardDashboardService struct, fully initialized and ready to handle
//     dashboard card operations.
func NewCardDashboardService(
	params *cardDashboardDeps,
) CardDashboardService {
	return &cardDashboardService{
		cache:                   params.Cache,
		cardDashboardRepository: params.CardDashboardRepository,
		logger:                  params.Logger,
		observability:           params.Observability,
	}
}

func (s *cardDashboardService) DashboardCard(ctx context.Context) (*response.DashboardCard, error) {
	const method = "DashboardCard"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetDashboardCardCache(ctx); found {
		s.logger.Debug("DashboardCard cache hit")
		return data, nil
	}

	totalBalance, err := s.cardDashboardRepository.GetTotalBalances(ctx)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*response.DashboardCard](
			s.logger,
			card_errors.ErrFailedFindTotalBalances,
			method,
			span,
		)
	}

	totalTopup, err := s.cardDashboardRepository.GetTotalTopAmount(ctx)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*response.DashboardCard](
			s.logger,
			card_errors.ErrFailedFindTotalTopAmount,
			method,
			span,
		)
	}

	totalWithdraw, err := s.cardDashboardRepository.GetTotalWithdrawAmount(ctx)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*response.DashboardCard](
			s.logger,
			card_errors.ErrFailedFindTotalWithdrawAmount,
			method,
			span,
		)
	}

	totalTransaction, err := s.cardDashboardRepository.GetTotalTransactionAmount(ctx)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*response.DashboardCard](
			s.logger,
			card_errors.ErrFailedFindTotalTransactionAmount,
			method,
			span,
		)
	}

	totalTransfer, err := s.cardDashboardRepository.GetTotalTransferAmount(ctx)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*response.DashboardCard](
			s.logger,
			card_errors.ErrFailedFindTotalTransferAmount,
			method,
			span,
		)
	}

	result := &response.DashboardCard{
		TotalBalance:     totalBalance,
		TotalTopup:       totalTopup,
		TotalWithdraw:    totalWithdraw,
		TotalTransaction: totalTransaction,
		TotalTransfer:    totalTransfer,
	}

	s.cache.SetDashboardCardCache(ctx, result)

	logSuccess("Completed DashboardCard service",
		zap.Int64("total_balance", *totalBalance),
		zap.Int64("total_topup", *totalTopup),
		zap.Int64("total_withdraw", *totalWithdraw),
		zap.Int64("total_transaction", *totalTransaction),
		zap.Int64("total_transfer", *totalTransfer),
	)

	return result, nil
}

func (s *cardDashboardService) DashboardCardCardNumber(ctx context.Context, cardNumber string) (*response.DashboardCardCardNumber, error) {
	const method = "DashboardCardCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(cardNumber)))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetDashboardCardCardNumberCache(ctx, cardNumber); found {
		s.logger.Debug("DashboardCardCardNumber cache hit", zap.String("card_number", observability.MaskIdentifier(cardNumber)))
		return data, nil
	}

	totalBalance, err := s.cardDashboardRepository.GetTotalBalanceByCardNumber(ctx, cardNumber)

	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*response.DashboardCardCardNumber](
			s.logger,
			card_errors.ErrFailedFindTotalBalanceByCard,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		)
	}

	totalTopup, err := s.cardDashboardRepository.GetTotalTopupAmountByCardNumber(ctx, cardNumber)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*response.DashboardCardCardNumber](
			s.logger,
			card_errors.ErrFailedFindTotalTopupAmountByCard,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		)
	}

	totalWithdraw, err := s.cardDashboardRepository.GetTotalWithdrawAmountByCardNumber(ctx, cardNumber)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*response.DashboardCardCardNumber](
			s.logger,
			card_errors.ErrFailedFindTotalWithdrawAmountByCard,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		)
	}

	totalTransaction, err := s.cardDashboardRepository.GetTotalTransactionAmountByCardNumber(ctx, cardNumber)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*response.DashboardCardCardNumber](
			s.logger,
			card_errors.ErrFailedFindTotalTransactionAmountByCard,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		)
	}

	totalTransferSent, err := s.cardDashboardRepository.GetTotalTransferAmountBySender(ctx, cardNumber)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*response.DashboardCardCardNumber](
			s.logger,
			card_errors.ErrFailedFindTotalTransferAmountBySender,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		)
	}

	totalTransferReceived, err := s.cardDashboardRepository.GetTotalTransferAmountByReceiver(ctx, cardNumber)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*response.DashboardCardCardNumber](
			s.logger,
			card_errors.ErrFailedFindTotalTransferAmountByReceiver,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		)
	}

	result := &response.DashboardCardCardNumber{
		TotalBalance:          totalBalance,
		TotalTopup:            totalTopup,
		TotalWithdraw:         totalWithdraw,
		TotalTransaction:      totalTransaction,
		TotalTransferSend:     totalTransferSent,
		TotalTransferReceiver: totalTransferReceived,
	}

	s.cache.SetDashboardCardCardNumberCache(ctx, cardNumber, result)

	logSuccess("Completed DashboardCardCardNumber service",
		zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		zap.Int64("total_balance", *totalBalance),
		zap.Int64("total_topup", *totalTopup),
		zap.Int64("total_withdraw", *totalWithdraw),
		zap.Int64("total_transaction", *totalTransaction),
		zap.Int64("total_transfer_sent", *totalTransferSent),
		zap.Int64("total_transfer_received", *totalTransferReceived),
	)

	return result, nil
}
