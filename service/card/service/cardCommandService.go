package service

import (
	"context"
	"time"

	mencache "github.com/MamangRust/monolith-payment-gateway-card/redis"
	"github.com/MamangRust/monolith-payment-gateway-card/repository"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharederrorhandler "github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// CardCommandServiceDeps defines dependencies for cardCommandService.
type CardCommandServiceDeps struct {
	Cache                 mencache.CardCommandCache
	UserRepository        repository.UserRepository
	CardCommandRepository repository.CardCommandRepository
	Logger                logger.LoggerInterface
	Observability         observability.TraceLoggerObservability
}

// cardCommandService implements CardCommandService.
type cardCommandService struct {
	cache                 mencache.CardCommandCache
	userRepository        repository.UserRepository
	cardCommandRepository repository.CardCommandRepository
	logger                logger.LoggerInterface
	observability         observability.TraceLoggerObservability
}

func NewCardCommandService(params *CardCommandServiceDeps) CardCommandService {

	return &cardCommandService{
		cache:                 params.Cache,
		userRepository:        params.UserRepository,
		cardCommandRepository: params.CardCommandRepository,
		logger:                params.Logger,
		observability:         params.Observability,
	}
}

func (s *cardCommandService) CreateCard(ctx context.Context, request *requests.CreateCardRequest) (*db.CreateCardRow, error) {
	const method = "CreateCard"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	_, err := s.userRepository.FindById(ctx, request.UserID)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.CreateCardRow](s.logger, err, method, span, zap.Int("user_id", request.UserID))
	}

	res, err := s.cardCommandRepository.CreateCard(ctx, request)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.CreateCardRow](s.logger, err, method, span, zap.Int("user_id", request.UserID))
	}

	logSuccess("Successfully created card", zap.Int("card.id", int(res.CardID)), zap.String("card.card_number", observability.MaskIdentifier(res.CardNumber)))

	return res, nil
}

func (s *cardCommandService) UpdateCard(ctx context.Context, request *requests.UpdateCardRequest) (*db.UpdateCardRow, error) {
	const method = "UpdateCard"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	_, err := s.userRepository.FindById(ctx, request.UserID)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.UpdateCardRow](s.logger, err, method, span, zap.Int("user_id", request.UserID))
	}

	res, err := s.cardCommandRepository.UpdateCard(ctx, request)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.UpdateCardRow](s.logger, err, method, span, zap.Int("card_id", request.CardID))
	}

	s.cache.DeleteCardCache(ctx, int(res.CardID), int(res.UserID), res.CardNumber)

	logSuccess("Successfully updated card", zap.Int("card.id", int(res.CardID)))

	return res, nil
}

func (s *cardCommandService) TrashedCard(ctx context.Context, card_id int) (*db.TrashCardRow, error) {
	const method = "TrashedCard"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("card_id", card_id))

	defer func() {
		end(status)
	}()

	res, err := s.cardCommandRepository.TrashedCard(ctx, card_id)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.TrashCardRow](
			s.logger,
			err,
			method,
			span,

			zap.Int("card_id", card_id),
		)
	}

	s.cache.DeleteCardCache(ctx, int(res.CardID), int(res.UserID), res.CardNumber)

	logSuccess("Successfully trashed card", zap.Int("card_id", int(res.CardID)))

	return res, nil
}

func (s *cardCommandService) RestoreCard(ctx context.Context, card_id int) (*db.RestoreCardRow, error) {
	const method = "RestoreCard"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("card_id", card_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Restoring card", zap.Int("card_id", card_id))

	res, err := s.cardCommandRepository.RestoreCard(ctx, card_id)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.RestoreCardRow](
			s.logger,
			err,
			method,
			span,

			zap.Int("card_id", card_id),
		)
	}

	s.cache.DeleteCardCache(ctx, int(res.CardID), int(res.UserID), res.CardNumber)

	logSuccess("Successfully restored card", zap.Int("card_id", int(res.CardID)))

	return res, nil
}

func (s *cardCommandService) DeleteCardPermanent(ctx context.Context, card_id int) (bool, error) {
	const method = "DeleteCardPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("card_id", card_id))

	defer func() {
		end(status)
	}()

	_, err := s.cardCommandRepository.DeleteCardPermanent(ctx, card_id)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,

			zap.Int("card_id", card_id),
		)
	}

	s.cache.DeleteCardCommandCache(ctx, card_id)

	logSuccess("Successfully deleted card permanently", zap.Int("card_id", card_id))

	return true, nil
}

func (s *cardCommandService) RestoreAllCard(ctx context.Context) (bool, error) {
	const method = "RestoreAllCard"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	_, err := s.cardCommandRepository.RestoreAllCard(ctx)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,
		)
	}

	logSuccess("Successfully restored all cards")

	return true, nil
}

func (s *cardCommandService) DeleteAllCardPermanent(ctx context.Context) (bool, error) {
	const method = "DeleteAllCardPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	_, err := s.cardCommandRepository.DeleteAllCardPermanent(ctx)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,
		)
	}

	logSuccess("Successfully deleted all cards permanently")

	return true, nil
}

func (s *cardCommandService) ToggleCardStatus(ctx context.Context, request *requests.ToggleCardStatusRequest) (*db.UpdateCardStatusRow, error) {
	const method = "ToggleCardStatus"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("card_id", request.CardID))
	defer func() { end(status) }()

	res, err := s.cardCommandRepository.ToggleCardStatus(ctx, request)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.UpdateCardStatusRow](s.logger, err, method, span, zap.Int("card_id", request.CardID))
	}

	s.cache.DeleteCardCache(ctx, int(res.CardID), int(res.UserID), res.CardNumber)

	logSuccess("Successfully toggled card status", zap.Int("card_id", request.CardID), zap.String("status", res.Status))

	return res, nil
}

func (s *cardCommandService) UpdateCreditLimit(ctx context.Context, request *requests.UpdateCreditLimitRequest) (*db.UpdateCreditLimitRow, error) {
	const method = "UpdateCreditLimit"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("card_id", request.CardID))
	defer func() { end(status) }()

	res, err := s.cardCommandRepository.UpdateCreditLimit(ctx, request)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.UpdateCreditLimitRow](s.logger, err, method, span, zap.Int("card_id", request.CardID), zap.Int("credit_limit", request.CreditLimit))
	}

	s.cache.DeleteCardCache(ctx, int(res.CardID), int(res.UserID), res.CardNumber)

	logSuccess("Successfully updated credit limit", zap.Int("card_id", request.CardID), zap.Int("credit_limit", request.CreditLimit))

	return res, nil
}

func (s *cardCommandService) RedeemPoints(ctx context.Context, request *requests.RedeemPointsRequest) (*db.RedeemRewardPointsRow, error) {
	const method = "RedeemPoints"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("card_id", request.CardID))
	defer func() { end(status) }()

	res, err := s.cardCommandRepository.RedeemPoints(ctx, request)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.RedeemRewardPointsRow](s.logger, err, method, span, zap.Int("card_id", request.CardID), zap.Int("points", request.Points))
	}

	s.cache.DeleteCardCache(ctx, int(res.CardID), int(res.UserID), res.CardNumber)

	logSuccess("Successfully redeemed reward points", zap.Int("card_id", request.CardID), zap.Int("points", request.Points))

	return res, nil
}

func (s *cardCommandService) ProcessBillingCycles(ctx context.Context) error {
	const method = "ProcessBillingCycles"

	ctx, _, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	_, err := s.cardCommandRepository.ProcessBillingCycles(ctx, time.Now().Day())
	if err != nil {
		status = "error"
		s.logger.Error("Failed to process billing cycles", zap.Error(err))
		return err
	}

	logSuccess("Successfully processed billing cycles")

	return nil
}
