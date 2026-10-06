package service

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	statemachine "github.com/MamangRust/monolith-payment-gateway-shared/domain/status"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	sharederrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"

	mencache "github.com/MamangRust/monolith-payment-gateway-topup/redis"
	"github.com/MamangRust/monolith-payment-gateway-topup/repository"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// topupCommandDeps groups dependencies for top-up command service.
type TopupCommandDeps struct {
	Kafka                  *kafka.Kafka
	Cache                  mencache.TopupCommandCache
	CardRepository         repository.CardRepository
	TopupQueryRepository   repository.TopupQueryRepository
	TopupCommandRepository repository.TopupCommandRepository
	SaldoRepository        repository.SaldoRepository
	Logger                 logger.LoggerInterface
	Observability          observability.TraceLoggerObservability
}

// topupCommandService handles top-up command operations.
type topupCommandService struct {
	kafka                  *kafka.Kafka
	cache                  mencache.TopupCommandCache
	topupQueryRepository   repository.TopupQueryRepository
	cardRepository         repository.CardRepository
	topupCommandRepository repository.TopupCommandRepository
	saldoRepository        repository.SaldoRepository
	logger                 logger.LoggerInterface
	observability          observability.TraceLoggerObservability
}

func NewTopupCommandService(
	params *TopupCommandDeps,
) TopupCommandService {
	return &topupCommandService{
		kafka:                  params.Kafka,
		cache:                  params.Cache,
		topupQueryRepository:   params.TopupQueryRepository,
		topupCommandRepository: params.TopupCommandRepository,
		saldoRepository:        params.SaldoRepository,
		cardRepository:         params.CardRepository,
		logger:                 params.Logger,
		observability:          params.Observability,
	}
}

func (s *topupCommandService) CreateTopup(ctx context.Context, request *requests.CreateTopupRequest) (*db.UpdateTopupStatusRow, error) {
	const method = "CreateTopup"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	card, err := s.cardRepository.FindUserCardByCardNumber(ctx, request.CardNumber)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTopupStatusRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	}

	// Saldo credit + topup record commit atomically in one SQL statement:
	// no read-modify-write race, no separate status update, and a missing
	// saldo is rejected without changing any balance.
	topup, err := s.topupCommandRepository.CreateTopupAtomic(ctx, request)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTopupStatusRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	}

	expireDate := card.ExpireDate.Time

	// On a replay the original side-effects (card metadata update, email) were
	// already performed when the key was first persisted; skip them so a retry
	// never re-sends notifications or re-runs metadata writes.
	if !topup.Replayed {
		if _, err := s.cardRepository.UpdateCard(ctx, &requests.UpdateCardRequest{
			CardID:       int(card.CardID),
			UserID:       int(card.UserID),
			CardType:     card.CardType,
			ExpireDate:   expireDate,
			CVV:          card.Cvv,
			CardProvider: card.CardProvider,
		}); err != nil {
			// The financial operation already committed atomically (balance credited
			// and topup recorded as success). A card metadata update failure must
			// not surface as a financial error, otherwise clients may retry the
			// topup and double-credit.
			s.logger.Error("topup committed but card update failed", zap.Error(err), zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
		}
	}

	updatedTopup := &db.UpdateTopupStatusRow{
		TopupID:     topup.Row.TopupID,
		TopupNo:     topup.Row.TopupNo,
		CardNumber:  topup.Row.CardNumber,
		TopupAmount: topup.Row.TopupAmount,
		TopupMethod: topup.Row.TopupMethod,
		TopupTime:   topup.Row.TopupTime,
		Status:      topup.Row.Status,
		CreatedAt:   topup.Row.CreatedAt,
		UpdatedAt:   topup.Row.UpdatedAt,
	}

	// The atomic SQL command persists the notification in outbox_events in
	// the same transaction as the balance credit and topup row. The relay owns
	// Kafka delivery, so this path cannot lose the event after commit and does
	// not publish a duplicate directly from the request goroutine.

	if topup.Replayed {
		s.logger.Debug("idempotent replay: returning existing topup", zap.String("idempotency_key", observability.MaskIdentifier(request.IdempotencyKey)))
	}

	logSuccess("Topup created successfully", zap.String("cardNumber", observability.MaskIdentifier(request.CardNumber)), zap.Int("topupID", int(topup.Row.TopupID)), zap.Float64("topupAmount", float64(request.TopupAmount)))

	return updatedTopup, nil
}

func (s *topupCommandService) UpdateTopup(ctx context.Context, request *requests.UpdateTopupRequest) (*db.UpdateTopupStatusRow, error) {
	const method = "UpdateTopup"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	_, err := s.cardRepository.FindCardByCardNumber(ctx, request.CardNumber)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTopupStatusRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	}

	existingTopup, err := s.topupQueryRepository.FindById(ctx, *request.TopupID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTopupStatusRow](s.logger, err, method, span, zap.Int("topup_id", *request.TopupID))
	}

	// The settlement delta is computed from the existing record and applied to
	// the card on file; changing the card number would move the delta to the
	// wrong account, so it is rejected explicitly.
	if request.CardNumber != existingTopup.CardNumber {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTopupStatusRow](s.logger, sharederrors.NewBadRequestError("card number cannot be changed on topup update"), method, span, zap.Int("topup_id", *request.TopupID))
	}

	// Enforce the lifecycle: an update may re-confirm an already-success topup
	// or complete a pending/failed retry, but it can never revert a settled
	// record to pending. Validated before any balance delta is applied.
	if err := statemachine.ValidateTransition(statemachine.DomainTopup, existingTopup.Status, statemachine.Success); err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTopupStatusRow](s.logger, err, method, span, zap.Int("topup_id", *request.TopupID), zap.String("current_status", existingTopup.Status))
	}

	// The topup record update and the balance settlement commit atomically in a
	// single SQL statement. The delta is computed from the locked pre-update
	// amount inside the statement, so concurrent updates can never read a stale
	// amount and produce a lost update or a negative double-apply. The
	// statement also sets status to 'success', replacing the old three-step
	// UpdateTopup + delta + status flow.
	atomicTopup, err := s.topupCommandRepository.UpdateTopupAtomic(ctx, request)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTopupStatusRow](s.logger, err, method, span, zap.Int("topup_id", *request.TopupID))
	}

	updatedTopup := &db.UpdateTopupStatusRow{
		TopupID:     atomicTopup.TopupID,
		TopupNo:     atomicTopup.TopupNo,
		CardNumber:  atomicTopup.CardNumber,
		TopupAmount: atomicTopup.TopupAmount,
		TopupMethod: atomicTopup.TopupMethod,
		TopupTime:   atomicTopup.TopupTime,
		Status:      atomicTopup.Status,
		CreatedAt:   atomicTopup.CreatedAt,
		UpdatedAt:   atomicTopup.UpdatedAt,
	}

	logSuccess("UpdateTopup process completed", zap.Int("topup_id", *request.TopupID))

	return updatedTopup, nil
}

func (s *topupCommandService) TrashedTopup(ctx context.Context, topup_id int) (*db.Topup, error) {
	const method = "TrashedTopup"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("topup_id", topup_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Starting TrashedTopup process", zap.Int("topup_id", topup_id))

	res, err := s.topupCommandRepository.TrashedTopup(ctx, topup_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Topup](
			s.logger,
			err,
			method,
			span,

			zap.Int("topup_id", topup_id),
		)
	}

	logSuccess("TrashedTopup process completed", zap.Int("topup_id", topup_id))

	return res, nil
}

func (s *topupCommandService) RestoreTopup(ctx context.Context, topup_id int) (*db.Topup, error) {
	const method = "RestoreTopup"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("topup_id", topup_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Starting RestoreTopup process", zap.Int("topup_id", topup_id))

	res, err := s.topupCommandRepository.RestoreTopup(ctx, topup_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Topup](
			s.logger,
			err,
			method,
			span,

			zap.Int("topup_id", topup_id),
		)
	}

	logSuccess("RestoreTopup process completed", zap.Int("topup_id", topup_id))

	return res, nil
}

func (s *topupCommandService) DeleteTopupPermanent(ctx context.Context, topup_id int) (bool, error) {
	const method = "DeleteTopupPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("topup_id", topup_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Starting DeleteTopupPermanent process", zap.Int("topup_id", topup_id))

	_, err := s.topupCommandRepository.DeleteTopupPermanent(ctx, topup_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,

			zap.Int("topup_id", topup_id),
		)
	}

	logSuccess("DeleteTopupPermanent process completed", zap.Int("topup_id", topup_id))

	return true, nil
}

func (s *topupCommandService) RestoreAllTopup(ctx context.Context) (bool, error) {
	const method = "RestoreAllTopup"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	s.logger.Debug("Restoring all topups")

	_, err := s.topupCommandRepository.RestoreAllTopup(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,
		)
	}

	logSuccess("Successfully restored all topups")
	return true, nil
}

func (s *topupCommandService) DeleteAllTopupPermanent(ctx context.Context) (bool, error) {
	const method = "DeleteAllTopupPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	s.logger.Debug("Permanently deleting all topups")

	_, err := s.topupCommandRepository.DeleteAllTopupPermanent(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,
		)
	}

	logSuccess("Successfully deleted all topups permanently")
	return true, nil
}
