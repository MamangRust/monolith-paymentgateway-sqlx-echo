package service

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	statemachine "github.com/MamangRust/monolith-payment-gateway-shared/domain/status"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	sharederrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"

	mencache "github.com/MamangRust/monolith-payment-gateway-transfer/redis"
	"github.com/MamangRust/monolith-payment-gateway-transfer/repository"
	"go.uber.org/zap"
)

// transferCommandDeps defines dependencies for transferCommandService.
type TransferCommandDeps struct {
	Kafka *kafka.Kafka
	Cache mencache.TransferCommandCache

	CardRepository  repository.CardRepository
	SaldoRepository repository.SaldoRepository

	TransferQueryRepository   repository.TransferQueryRepository
	TransferCommandRepository repository.TransferCommandRepository

	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
}

// transferCommandService handles command-side transfer operations.
type transferCommandService struct {
	kafka *kafka.Kafka
	cache mencache.TransferCommandCache

	cardRepository  repository.CardRepository
	saldoRepository repository.SaldoRepository

	transferQueryRepository   repository.TransferQueryRepository
	transferCommandRepository repository.TransferCommandRepository

	logger        logger.LoggerInterface
	observability observability.TraceLoggerObservability
}

func NewTransferCommandService(
	params *TransferCommandDeps,
) TransferCommandService {
	return &transferCommandService{
		kafka:                     params.Kafka,
		cache:                     params.Cache,
		cardRepository:            params.CardRepository,
		saldoRepository:           params.SaldoRepository,
		transferQueryRepository:   params.TransferQueryRepository,
		transferCommandRepository: params.TransferCommandRepository,
		logger:                    params.Logger,
		observability:             params.Observability,
	}
}

func (s *transferCommandService) CreateTransaction(ctx context.Context, request *requests.CreateTransferRequest) (*db.UpdateTransferStatusRow, error) {
	const method = "CreateTransaction"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	_, err := s.cardRepository.FindUserCardByCardNumber(ctx, request.TransferFrom)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransferStatusRow](s.logger, err, method, span, zap.String("from_card", observability.MaskIdentifier(request.TransferFrom)))
	}

	_, err = s.cardRepository.FindCardByCardNumber(ctx, request.TransferTo)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransferStatusRow](s.logger, err, method, span, zap.String("to_card", observability.MaskIdentifier(request.TransferTo)))
	}

	// Sender debit + receiver credit + transfer record commit atomically in
	// one SQL statement: no read-modify-write race, no manual rollback, and
	// insufficient balance is rejected without changing any balance.
	transfer, err := s.transferCommandRepository.CreateTransferAtomic(ctx, request)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransferStatusRow](s.logger, err, method, span, zap.String("from_card", observability.MaskIdentifier(request.TransferFrom)), zap.String("to_card", observability.MaskIdentifier(request.TransferTo)), zap.Float64("amount", float64(request.TransferAmount)))
	}

	updatedTransfer := &db.UpdateTransferStatusRow{
		TransferID:     transfer.Row.TransferID,
		TransferNo:     transfer.Row.TransferNo,
		TransferFrom:   transfer.Row.TransferFrom,
		TransferTo:     transfer.Row.TransferTo,
		TransferAmount: transfer.Row.TransferAmount,
		TransferTime:   transfer.Row.TransferTime,
		Status:         transfer.Row.Status,
		CreatedAt:      transfer.Row.CreatedAt,
		UpdatedAt:      transfer.Row.UpdatedAt,
	}

	if transfer.Replayed {
		s.logger.Debug("idempotent replay: returning existing transfer", zap.String("idempotency_key", observability.MaskIdentifier(request.IdempotencyKey)))
	}

	logSuccess("Transfer created successfully", zap.Int("transfer_id", int(updatedTransfer.TransferID)), zap.String("from", observability.MaskIdentifier(request.TransferFrom)), zap.String("to", observability.MaskIdentifier(request.TransferTo)))

	return updatedTransfer, nil
}

// UpdateTransaction updates an existing transfer transaction.
//
// Parameters:
//   - ctx: The context for timeout and cancellation.
//   - request: The request containing updated transfer details.
//
// Returns:
//   - *response.TransferResponse: The updated transfer data.
//   - *response.ErrorResponse: Error details if operation fails.
func (s *transferCommandService) UpdateTransaction(ctx context.Context, request *requests.UpdateTransferRequest) (*db.UpdateTransferRow, error) {
	const method = "UpdateTransaction"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	// 1. Dapatkan data transfer yang ada
	transfer, err := s.transferQueryRepository.FindById(ctx, *request.TransferID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransferRow](s.logger, err, method, span, zap.Int("transfer_id", *request.TransferID))
	}

	// Settlement deltas are applied to the accounts on the existing record;
	// changing sender/receiver on update would move money to the wrong
	// accounts, so it is rejected explicitly.
	if request.TransferFrom != transfer.TransferFrom || request.TransferTo != transfer.TransferTo {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransferRow](s.logger, sharederrors.NewBadRequestError("transfer sender/receiver cannot be changed on update"), method, span, zap.Int("transfer_id", *request.TransferID))
	}

	// Enforce the lifecycle: an update may re-confirm an already-success
	// transfer or complete a pending/failed retry; a settled transfer can
	// never revert to pending. Validated before any settlement delta moves.
	if err := statemachine.ValidateTransition(statemachine.DomainTransfer, transfer.Status, statemachine.Success); err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransferRow](s.logger, err, method, span, zap.Int("transfer_id", *request.TransferID), zap.String("current_status", transfer.Status))
	}

	atomicTransfer, err := s.transferCommandRepository.UpdateTransferAtomic(ctx, db.UpdateTransferAtomicParams{
		TransferID:     int32(*request.TransferID),
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int32(request.TransferAmount),
		TransferTime:   time.Now(),
	})
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransferRow](s.logger, err, method, span, zap.Int("transfer_id", *request.TransferID))
	}

	updatedTransfer := &db.UpdateTransferRow{
		TransferID:     atomicTransfer.TransferID,
		TransferNo:     atomicTransfer.TransferNo,
		TransferFrom:   atomicTransfer.TransferFrom,
		TransferTo:     atomicTransfer.TransferTo,
		TransferAmount: atomicTransfer.TransferAmount,
		TransferTime:   atomicTransfer.TransferTime,
		Status:         atomicTransfer.Status,
		CreatedAt:      atomicTransfer.CreatedAt,
		UpdatedAt:      atomicTransfer.UpdatedAt,
	}

	logSuccess("Successfully updated transfer", zap.Int("transfer.id", *request.TransferID))

	return updatedTransfer, nil
}

func (s *transferCommandService) TrashedTransfer(ctx context.Context, transfer_id int) (*db.Transfer, error) {
	const method = "TrashedTransfer"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("transfer_id", transfer_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Starting trashed transfer process", zap.Int("transfer_id", transfer_id))

	res, err := s.transferCommandRepository.TrashedTransfer(ctx, transfer_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Transfer](
			s.logger,
			err,
			method,
			span,

			zap.Int("transfer_id", transfer_id),
		)
	}

	logSuccess("Successfully trashed transfer", zap.Int("transfer_id", transfer_id))

	return res, nil
}

func (s *transferCommandService) RestoreTransfer(ctx context.Context, transfer_id int) (*db.Transfer, error) {
	const method = "RestoreTransfer"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("transfer_id", transfer_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Starting restore transfer process", zap.Int("transfer_id", transfer_id))

	res, err := s.transferCommandRepository.RestoreTransfer(ctx, transfer_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Transfer](
			s.logger,
			err,
			method,
			span,

			zap.Int("transfer_id", transfer_id),
		)
	}

	logSuccess("Successfully restored transfer", zap.Int("transfer_id", transfer_id))

	return res, nil
}

func (s *transferCommandService) DeleteTransferPermanent(ctx context.Context, transfer_id int) (bool, error) {
	const method = "DeleteTransferPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("transfer_id", transfer_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Starting delete transfer permanent process", zap.Int("transfer_id", transfer_id))

	_, err := s.transferCommandRepository.DeleteTransferPermanent(ctx, transfer_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,

			zap.Int("transfer_id", transfer_id),
		)
	}

	logSuccess("Successfully deleted transfer permanently", zap.Int("transfer_id", transfer_id))

	return true, nil
}

func (s *transferCommandService) RestoreAllTransfer(ctx context.Context) (bool, error) {
	const method = "RestoreAllTransfer"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	s.logger.Debug("Restoring all transfers")

	_, err := s.transferCommandRepository.RestoreAllTransfer(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,
		)
	}

	logSuccess("Successfully restored all transfers")
	return true, nil
}

func (s *transferCommandService) DeleteAllTransferPermanent(ctx context.Context) (bool, error) {
	const method = "DeleteAllTransferPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	s.logger.Debug("Permanently deleting all transfers")

	_, err := s.transferCommandRepository.DeleteAllTransferPermanent(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,
		)
	}

	logSuccess("Successfully deleted all transfers permanently")
	return true, nil
}
