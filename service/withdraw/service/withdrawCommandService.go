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
	mencache "github.com/MamangRust/monolith-payment-gateway-withdraw/redis"
	"github.com/MamangRust/monolith-payment-gateway-withdraw/repository"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// withdrawCommandServiceDeps defines dependencies for withdrawCommandService.
type WithdrawCommandServiceDeps struct {
	Cache mencache.WithdrawCommandCache
	Kafka *kafka.Kafka

	CardRepository    repository.CardRepository
	SaldoRepository   repository.SaldoRepository
	CommandRepository repository.WithdrawCommandRepository
	QueryRepository   repository.WithdrawQueryRepository

	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
}

// withdrawCommandService handles command-side withdraw operations.
type withdrawCommandService struct {
	cache mencache.WithdrawCommandCache
	kafka *kafka.Kafka

	cardRepository  repository.CardRepository
	saldoRepository repository.SaldoRepository

	withdrawCommandRepository repository.WithdrawCommandRepository
	withdrawQueryRepository   repository.WithdrawQueryRepository

	logger        logger.LoggerInterface
	observability observability.TraceLoggerObservability
}

func NewWithdrawCommandService(
	deps *WithdrawCommandServiceDeps,
) WithdrawCommandService {
	return &withdrawCommandService{
		kafka:                     deps.Kafka,
		cache:                     deps.Cache,
		cardRepository:            deps.CardRepository,
		saldoRepository:           deps.SaldoRepository,
		withdrawCommandRepository: deps.CommandRepository,
		withdrawQueryRepository:   deps.QueryRepository,
		logger:                    deps.Logger,
		observability:             deps.Observability,
	}
}

func (s *withdrawCommandService) Create(ctx context.Context, request *requests.CreateWithdrawRequest) (*db.UpdateWithdrawStatusRow, error) {
	const method = "CreateWithdraw"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	_, err := s.cardRepository.FindUserCardByCardNumber(ctx, request.CardNumber)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateWithdrawStatusRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	}

	// Saldo debit + withdraw record commit atomically in one SQL statement:
	// no read-modify-write race, no separate status update, and insufficient
	// balance is rejected without changing any balance.
	withdrawRecord, err := s.withdrawCommandRepository.CreateWithdrawAtomic(ctx, request)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateWithdrawStatusRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	}

	updatedWithdraw := &db.UpdateWithdrawStatusRow{
		WithdrawID:     withdrawRecord.Row.WithdrawID,
		WithdrawNo:     withdrawRecord.Row.WithdrawNo,
		CardNumber:     withdrawRecord.Row.CardNumber,
		WithdrawAmount: withdrawRecord.Row.WithdrawAmount,
		WithdrawTime:   withdrawRecord.Row.WithdrawTime,
		Status:         withdrawRecord.Row.Status,
		CreatedAt:      withdrawRecord.Row.CreatedAt,
		UpdatedAt:      withdrawRecord.Row.UpdatedAt,
	}

	if withdrawRecord.Replayed {
		s.logger.Debug("idempotent replay: returning existing withdraw", zap.String("idempotency_key", observability.MaskIdentifier(request.IdempotencyKey)))
	}

	logSuccess("Successfully created withdraw", zap.Int("withdraw.id", int(updatedWithdraw.WithdrawID)))

	return updatedWithdraw, nil
}

func (s *withdrawCommandService) Update(ctx context.Context, request *requests.UpdateWithdrawRequest) (*db.UpdateWithdrawRow, error) {
	const method = "UpdateWithdraw"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	oldWithdraw, err := s.withdrawQueryRepository.FindById(ctx, *request.WithdrawID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateWithdrawRow](s.logger, err, method, span, zap.Int("withdraw_id", *request.WithdrawID))
	}

	// The settlement delta is computed from the existing record and applied to
	// the card on file; changing the card number would move the delta to the
	// wrong account, so it is rejected explicitly.
	if request.CardNumber != oldWithdraw.CardNumber {
		status = "error"
		return errorhandler.HandleError[*db.UpdateWithdrawRow](s.logger, sharederrors.NewBadRequestError("card number cannot be changed on withdraw update"), method, span, zap.Int("withdraw_id", *request.WithdrawID))
	}

	// Enforce the lifecycle: an update may re-confirm an already-success
	// withdraw or complete a pending/failed retry; a settled withdraw can
	// never revert to pending. Validated before any balance delta moves.
	if err := statemachine.ValidateTransition(statemachine.DomainWithdraw, oldWithdraw.Status, statemachine.Success); err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateWithdrawRow](s.logger, err, method, span, zap.Int("withdraw_id", *request.WithdrawID), zap.String("current_status", oldWithdraw.Status))
	}

	// The withdraw record update and the balance settlement commit atomically in
	// a single SQL statement. The delta is computed from the locked pre-update
	// amount inside the statement, so concurrent updates can never read a stale
	// amount and produce a lost update or a negative double-apply. The
	// statement also sets status to 'success', replacing the old three-step
	// UpdateSaldoBalanceDelta + UpdateWithdraw + status flow.
	atomicWithdraw, err := s.withdrawCommandRepository.UpdateWithdrawAtomic(ctx, request)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateWithdrawRow](s.logger, err, method, span, zap.Int("withdraw_id", *request.WithdrawID))
	}

	updatedWithdraw := &db.UpdateWithdrawRow{
		WithdrawID:     atomicWithdraw.WithdrawID,
		WithdrawNo:     atomicWithdraw.WithdrawNo,
		CardNumber:     atomicWithdraw.CardNumber,
		WithdrawAmount: atomicWithdraw.WithdrawAmount,
		WithdrawTime:   atomicWithdraw.WithdrawTime,
		Status:         atomicWithdraw.Status,
		CreatedAt:      atomicWithdraw.CreatedAt,
		UpdatedAt:      atomicWithdraw.UpdatedAt,
	}

	logSuccess("Successfully updated withdraw", zap.Int("withdraw.id", int(updatedWithdraw.WithdrawID)))

	return updatedWithdraw, nil
}

func (s *withdrawCommandService) TrashedWithdraw(ctx context.Context, withdraw_id int) (*db.Withdraw, error) {
	const method = "TrashedWithdraw"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("withdraw_id", withdraw_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Trashing withdraw", zap.Int("withdraw_id", withdraw_id))

	res, err := s.withdrawCommandRepository.TrashedWithdraw(ctx, withdraw_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Withdraw](
			s.logger,
			err,
			method,
			span,

			zap.Int("withdraw_id", withdraw_id),
		)
	}

	logSuccess("Successfully trashed withdraw", zap.Int("withdraw_id", withdraw_id))

	return res, nil
}

func (s *withdrawCommandService) RestoreWithdraw(ctx context.Context, withdraw_id int) (*db.Withdraw, error) {
	const method = "RestoreWithdraw"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("withdraw_id", withdraw_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Restoring withdraw", zap.Int("withdraw_id", withdraw_id))

	res, err := s.withdrawCommandRepository.RestoreWithdraw(ctx, withdraw_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Withdraw](
			s.logger,
			err,
			method,
			span,

			zap.Int("withdraw_id", withdraw_id),
		)
	}

	logSuccess("Successfully restored withdraw", zap.Int("withdraw_id", withdraw_id))

	return res, nil
}

func (s *withdrawCommandService) DeleteWithdrawPermanent(ctx context.Context, withdraw_id int) (bool, error) {
	const method = "DeleteWithdrawPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("withdraw_id", withdraw_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Deleting withdraw permanently", zap.Int("withdraw_id", withdraw_id))

	_, err := s.withdrawCommandRepository.DeleteWithdrawPermanent(ctx, withdraw_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,

			zap.Int("withdraw_id", withdraw_id),
		)
	}

	logSuccess("Successfully deleted withdraw permanently", zap.Int("withdraw_id", withdraw_id))

	return true, nil
}

func (s *withdrawCommandService) RestoreAllWithdraw(ctx context.Context) (bool, error) {
	const method = "RestoreAllWithdraw"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	s.logger.Debug("Restoring all withdraws")

	_, err := s.withdrawCommandRepository.RestoreAllWithdraw(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,
		)
	}

	logSuccess("Successfully restored all withdraws")
	return true, nil
}

func (s *withdrawCommandService) DeleteAllWithdrawPermanent(ctx context.Context) (bool, error) {
	const method = "DeleteAllWithdrawPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	s.logger.Debug("Permanently deleting all withdraws")

	_, err := s.withdrawCommandRepository.DeleteAllWithdrawPermanent(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,
		)
	}

	logSuccess("Successfully deleted all withdraws permanently")
	return true, nil
}
