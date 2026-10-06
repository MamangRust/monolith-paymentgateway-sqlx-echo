package service

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-card/repository"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharederrorhandler "github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// CardRewardServiceDeps defines dependencies for cardRewardService.
type CardRewardServiceDeps struct {
	CardRewardRepository  repository.CardRewardRepository
	CardCommandRepository repository.CardCommandRepository
	Logger                logger.LoggerInterface
	Observability         observability.TraceLoggerObservability
}

// cardRewardService implements CardRewardService.
type cardRewardService struct {
	cardRewardRepository  repository.CardRewardRepository
	cardCommandRepository repository.CardCommandRepository
	logger                logger.LoggerInterface
	observability         observability.TraceLoggerObservability
}

func NewCardRewardService(deps *CardRewardServiceDeps) CardRewardService {
	return &cardRewardService{
		cardRewardRepository:  deps.CardRewardRepository,
		cardCommandRepository: deps.CardCommandRepository,
		logger:                deps.Logger,
		observability:         deps.Observability,
	}
}

func (s *cardRewardService) EarnRewards(ctx context.Context, request *requests.EarnRewardsRequest) (*db.CardReward, error) {
	const method = "EarnRewards"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	res, err := s.cardRewardRepository.EarnRewards(ctx, request)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.CardReward](
			s.logger,
			card_errors.ErrFailedRedeemPoints,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(request.CardNumber)),
			zap.Int64("amount", request.Amount),
		)
	}

	logSuccess("Successfully earned rewards", zap.String("card_number", observability.MaskIdentifier(request.CardNumber)), zap.Int64("amount", request.Amount))

	return res, nil
}

func (s *cardRewardService) GetBalance(ctx context.Context, cardNumber string) (int64, error) {
	const method = "GetBalance"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(cardNumber)))

	defer func() {
		end(status)
	}()

	res, err := s.cardRewardRepository.GetBalance(ctx, cardNumber)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[int64](
			s.logger,
			err,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		)
	}

	logSuccess("Successfully retrieved reward balance", zap.String("card_number", observability.MaskIdentifier(cardNumber)), zap.Int64("balance", res))

	return res, nil
}

func (s *cardRewardService) GetHistory(ctx context.Context, cardNumber string) ([]*db.CardReward, error) {
	const method = "GetHistory"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(cardNumber)))

	defer func() {
		end(status)
	}()

	res, err := s.cardRewardRepository.GetHistory(ctx, cardNumber)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.CardReward](
			s.logger,
			err,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		)
	}

	logSuccess("Successfully retrieved reward history", zap.String("card_number", observability.MaskIdentifier(cardNumber)), zap.Int("count", len(res)))

	return res, nil
}

func (s *cardRewardService) RedeemRewards(ctx context.Context, cardNumber string, points int64) (int64, error) {
	const method = "RedeemRewards"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(cardNumber)),
		attribute.Int64("points", points))

	defer func() {
		end(status)
	}()

	balance, err := s.cardRewardRepository.GetBalance(ctx, cardNumber)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[int64](
			s.logger,
			err,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
			zap.Int64("points", points),
		)
	}

	if balance < points {
		status = "error"
		return sharederrorhandler.HandleError[int64](
			s.logger,
			sharedErrors.ErrBadRequest.WithMessage("Insufficient reward balance"),
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
			zap.Int64("balance", balance),
			zap.Int64("points", points),
		)
	}

	res, err := s.cardRewardRepository.RedeemRewards(ctx, cardNumber, points)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[int64](
			s.logger,
			card_errors.ErrFailedRedeemPoints,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
			zap.Int64("points", points),
		)
	}

	logSuccess("Successfully redeemed rewards", zap.String("card_number", observability.MaskIdentifier(cardNumber)), zap.Int64("points", points), zap.Int64("remaining_balance", res))

	return res, nil
}
