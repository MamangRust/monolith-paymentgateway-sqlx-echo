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

	mencache "github.com/MamangRust/monolith-payment-gateway-transaction/redis"
	"github.com/MamangRust/monolith-payment-gateway-transaction/repository"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// TransactionCommandServiceDeps groups dependencies for transaction command service.
type TransactionCommandServiceDeps struct {
	Kafka                        *kafka.Kafka
	Mencache                     mencache.TransactionCommandCache
	Tracer                       trace.Tracer
	MerchantRepository           repository.MerchantRepository
	CardRepository               repository.CardRepository
	SaldoRepository              repository.SaldoRepository
	TransactionQueryRepository   repository.TransactionQueryRepository
	TransactionCommandRepository repository.TransactionCommandRepository
	Logger                       logger.LoggerInterface
	Observability                observability.TraceLoggerObservability
}

// transactionCommandService handles transaction write operations.
type transactionCommandService struct {
	kafka                        *kafka.Kafka
	cache                        mencache.TransactionCommandCache
	merchantRepository           repository.MerchantRepository
	cardRepository               repository.CardRepository
	saldoRepository              repository.SaldoRepository
	transactionQueryRepository   repository.TransactionQueryRepository
	transactionCommandRepository repository.TransactionCommandRepository
	logger                       logger.LoggerInterface
	observability                observability.TraceLoggerObservability
}

func NewTransactionCommandService(
	params *TransactionCommandServiceDeps,
) TransactionCommandService {
	return &transactionCommandService{
		kafka:                        params.Kafka,
		cache:                        params.Mencache,
		merchantRepository:           params.MerchantRepository,
		cardRepository:               params.CardRepository,
		saldoRepository:              params.SaldoRepository,
		transactionCommandRepository: params.TransactionCommandRepository,
		transactionQueryRepository:   params.TransactionQueryRepository,
		logger:                       params.Logger,
		observability:                params.Observability,
	}
}

func (s *transactionCommandService) Create(ctx context.Context, apiKey string, request *requests.CreateTransactionRequest) (*db.UpdateTransactionStatusRow, error) {
	const method = "CreateTransaction"

	if request == nil {
		return nil, sharederrors.NewBadRequestError("transaction request cannot be nil")
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.String("apikey", observability.MaskIdentifier(apiKey)))
	defer func() { end(status) }()

	merchant, err := s.merchantRepository.FindByApiKey(ctx, apiKey)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.String("api_key", observability.MaskIdentifier(apiKey)))
	}

	card, err := s.cardRepository.FindUserCardByCardNumber(ctx, request.CardNumber)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	}

	merchantCard, err := s.cardRepository.FindCardByUserId(ctx, int(merchant.UserID))
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.Int("merchant_id", int(merchant.MerchantID)))
	}

	merchantID := int32(merchant.MerchantID)
	atomicResult, err := s.transactionCommandRepository.CreateTransactionAtomicIdempotent(ctx, db.CreateTransactionAtomicParams{
		CardNumber:      card.CardNumber,
		Amount:          int32(request.Amount),
		CardNumber2:     merchantCard.CardNumber,
		PaymentMethod:   request.PaymentMethod,
		MerchantID:      merchantID,
		TransactionTime: request.TransactionTime,
		IdempotencyKey:  request.IdempotencyKey,
	})
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span,
			zap.String("card_number", observability.MaskIdentifier(card.CardNumber)),
			zap.String("merchant_card_number", observability.MaskIdentifier(merchantCard.CardNumber)),
			zap.Int("amount", request.Amount))
	}

	atomicTransaction := atomicResult.Row
	updatedTransaction := &db.UpdateTransactionStatusRow{
		TransactionID:   atomicTransaction.TransactionID,
		TransactionNo:   atomicTransaction.TransactionNo,
		CardNumber:      atomicTransaction.CardNumber,
		Amount:          atomicTransaction.Amount,
		PaymentMethod:   atomicTransaction.PaymentMethod,
		MerchantID:      atomicTransaction.MerchantID,
		TransactionTime: atomicTransaction.TransactionTime,
		Status:          atomicTransaction.Status,
		CreatedAt:       atomicTransaction.CreatedAt,
		UpdatedAt:       atomicTransaction.UpdatedAt,
	}

	if s.cache != nil && !atomicResult.Replayed {
		s.cache.InvalidateTransactionCache(ctx)
	}

	if atomicResult.Replayed {
		span.SetAttributes(attribute.Bool("idempotent_replay", true))
		logSuccess("Idempotent transaction replay returned existing transaction", zap.Int("transaction.id", int(updatedTransaction.TransactionID)))
	} else {
		logSuccess("Successfully created transaction", zap.Int("transaction.id", int(updatedTransaction.TransactionID)))
	}

	return updatedTransaction, nil
}

func (s *transactionCommandService) Update(ctx context.Context, apiKey string, request *requests.UpdateTransactionRequest) (*db.UpdateTransactionRow, error) {
	const method = "UpdateTransaction"

	if request == nil || request.TransactionID == nil {
		return nil, sharederrors.NewBadRequestError("transaction update request is invalid")
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	transaction, err := s.transactionQueryRepository.FindById(ctx, *request.TransactionID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionRow](s.logger, err, method, span, zap.Int("transaction_id", *request.TransactionID))
	}

	merchant, err := s.merchantRepository.FindByApiKey(ctx, apiKey)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionRow](s.logger, err, method, span, zap.String("api_key", observability.MaskIdentifier(apiKey)))
	}
	if transaction.MerchantID != merchant.MerchantID {
		status = "error"
		err := sharederrors.NewForbiddenError("merchant does not own transaction")
		return errorhandler.HandleError[*db.UpdateTransactionRow](s.logger, err, method, span, zap.String("api_key", observability.MaskIdentifier(apiKey)))
	}

	merchantCard, err := s.cardRepository.FindCardByUserId(ctx, int(merchant.UserID))
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionRow](s.logger, err, method, span, zap.Int("merchant_id", int(merchant.MerchantID)))
	}

	// Enforce the lifecycle: an update may re-confirm an already-success
	// transaction or complete a pending/failed retry; a settled transaction
	// can never revert to pending. Validated before any balance delta moves.
	if err := statemachine.ValidateTransition(statemachine.DomainTransaction, transaction.Status, statemachine.Success); err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionRow](s.logger, err, method, span, zap.Int("transaction_id", *request.TransactionID), zap.String("current_status", transaction.Status))
	}

	// The transaction update and the balance settlement commit atomically in a
	// single SQL statement. The delta is computed from the locked pre-update
	// amount inside the statement, so concurrent updates can never read a stale
	// amount and produce a lost update or a negative double-apply. The statement
	// also sets status to 'success', replacing the old three-step
	// UpdateSaldoBalanceDelta + UpdateTransaction + status flow.
	atomicTransaction, err := s.transactionCommandRepository.UpdateTransactionAtomic(ctx, db.UpdateTransactionAtomicParams{
		TransactionID:      int32(*request.TransactionID),
		CardNumber:         transaction.CardNumber,
		Amount:             int32(request.Amount),
		PaymentMethod:      request.PaymentMethod,
		MerchantID:         int32(transaction.MerchantID),
		TransactionTime:    transaction.TransactionTime,
		MerchantCardNumber: merchantCard.CardNumber,
	})
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionRow](s.logger, err, method, span, zap.Int("transaction_id", *request.TransactionID))
	}

	res := &db.UpdateTransactionRow{
		TransactionID:   atomicTransaction.TransactionID,
		TransactionNo:   atomicTransaction.TransactionNo,
		CardNumber:      atomicTransaction.CardNumber,
		Amount:          atomicTransaction.Amount,
		PaymentMethod:   atomicTransaction.PaymentMethod,
		MerchantID:      atomicTransaction.MerchantID,
		TransactionTime: atomicTransaction.TransactionTime,
		Status:          atomicTransaction.Status,
		CreatedAt:       atomicTransaction.CreatedAt,
		UpdatedAt:       atomicTransaction.UpdatedAt,
	}

	logSuccess("Successfully updated transaction", zap.Int("transaction.id", int(res.TransactionID)))

	return res, nil
}

func (s *transactionCommandService) TrashedTransaction(ctx context.Context, transaction_id int) (*db.TrashTransactionRow, error) {
	const method = "TrashedTransaction"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("transaction_id", transaction_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Starting TrashedTransaction process", zap.Int("transaction_id", transaction_id))

	res, err := s.transactionCommandRepository.TrashedTransaction(ctx, transaction_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.TrashTransactionRow](
			s.logger,
			err,
			method,
			span,

			zap.Int("transaction_id", transaction_id),
		)
	}

	logSuccess("Successfully trashed transaction", zap.Int("transaction_id", transaction_id))

	return res, nil
}

func (s *transactionCommandService) RestoreTransaction(ctx context.Context, transaction_id int) (*db.RestoreTransactionRow, error) {
	const method = "RestoreTransaction"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("transaction_id", transaction_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Starting RestoreTransaction process", zap.Int("transaction_id", transaction_id))

	res, err := s.transactionCommandRepository.RestoreTransaction(ctx, transaction_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.RestoreTransactionRow](
			s.logger,
			err,
			method,
			span,

			zap.Int("transaction_id", transaction_id),
		)
	}

	logSuccess("Successfully restored transaction", zap.Int("transaction_id", transaction_id))

	return res, nil
}

func (s *transactionCommandService) DeleteTransactionPermanent(ctx context.Context, transaction_id int) (bool, error) {
	const method = "DeleteTransactionPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("transaction_id", transaction_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Starting DeleteTransactionPermanent process", zap.Int("transaction_id", transaction_id))

	_, err := s.transactionCommandRepository.DeleteTransactionPermanent(ctx, transaction_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,

			zap.Int("transaction_id", transaction_id),
		)
	}

	logSuccess("Successfully permanently deleted transaction", zap.Int("transaction_id", transaction_id))

	return true, nil
}

func (s *transactionCommandService) RestoreAllTransaction(ctx context.Context) (bool, error) {
	const method = "RestoreAllTransaction"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	s.logger.Debug("Restoring all transactions")

	_, err := s.transactionCommandRepository.RestoreAllTransaction(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,
		)
	}

	logSuccess("Successfully restored all transactions")
	return true, nil
}

func (s *transactionCommandService) DeleteAllTransactionPermanent(ctx context.Context) (bool, error) {
	const method = "DeleteAllTransactionPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	s.logger.Debug("Permanently deleting all transactions")

	_, err := s.transactionCommandRepository.DeleteAllTransactionPermanent(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,
		)
	}

	logSuccess("Successfully deleted all transactions permanently")
	return true, nil
}

func (s *transactionCommandService) AuthorizeTransaction(ctx context.Context, apiKey string, request *requests.CreateTransactionRequest) (*db.UpdateTransactionStatusRow, error) {
	const method = "AuthorizeTransaction"

	if request == nil {
		return nil, sharederrors.NewBadRequestError("transaction authorization request cannot be nil")
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.String("apikey", observability.MaskIdentifier(apiKey)))
	defer func() { end(status) }()

	merchant, err := s.merchantRepository.FindByApiKey(ctx, apiKey)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.String("api_key", observability.MaskIdentifier(apiKey)))
	}

	merchantID := int32(merchant.MerchantID)
	authorizeResult, err := s.transactionCommandRepository.AuthorizeTransactionAtomicIdempotent(ctx, db.AuthorizeTransactionAtomicParams{
		CardNumber:      request.CardNumber,
		Amount:          int32(request.Amount),
		PaymentMethod:   request.PaymentMethod,
		MerchantID:      merchantID,
		TransactionTime: request.TransactionTime,
		IdempotencyKey:  request.IdempotencyKey,
	})
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span,
			zap.String("card_number", observability.MaskIdentifier(request.CardNumber)),
			zap.Int("amount", request.Amount))
	}

	updatedTransaction := &db.UpdateTransactionStatusRow{
		TransactionID:   authorizeResult.Row.TransactionID,
		TransactionNo:   authorizeResult.Row.TransactionNo,
		CardNumber:      authorizeResult.Row.CardNumber,
		Amount:          authorizeResult.Row.Amount,
		PaymentMethod:   authorizeResult.Row.PaymentMethod,
		MerchantID:      authorizeResult.Row.MerchantID,
		TransactionTime: authorizeResult.Row.TransactionTime,
		Status:          authorizeResult.Row.Status,
		CreatedAt:       authorizeResult.Row.CreatedAt,
		UpdatedAt:       authorizeResult.Row.UpdatedAt,
	}

	logSuccess("Successfully authorized transaction", zap.Int("transaction.id", int(updatedTransaction.TransactionID)))

	return updatedTransaction, nil
}

func (s *transactionCommandService) CaptureTransaction(ctx context.Context, apiKey string, transactionID int) (*db.UpdateTransactionStatusRow, error) {
	const method = "CaptureTransaction"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("transaction_id", transactionID))
	defer func() { end(status) }()

	transaction, err := s.transactionQueryRepository.FindById(ctx, transactionID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.Int("transaction_id", transactionID))
	}

	// Only authorized transactions may be captured (canonical state machine).
	if err := statemachine.ValidateTransition(statemachine.DomainTransaction, transaction.Status, statemachine.Captured); err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.String("current_status", transaction.Status))
	}

	merchant, err := s.merchantRepository.FindByApiKey(ctx, apiKey)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.Int("transaction_id", transactionID))
	}
	if transaction.MerchantID != merchant.MerchantID {
		status = "error"
		err := sharederrors.NewForbiddenError("merchant does not own transaction")
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.Int("transaction_id", transactionID))
	}

	// Get merchant card for settlement
	merchantCard, err := s.cardRepository.FindCardByUserId(ctx, int(merchant.UserID))
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.Int("merchant_id", int(merchant.MerchantID)))
	}

	// Credit the merchant and mark the transaction captured in one atomic SQL
	// statement. Only an 'authorized' transaction can be captured; the row lock
	// plus the status guard make a concurrent double-capture impossible (no
	// double merchant credit).
	atomicTransaction, err := s.transactionCommandRepository.CaptureTransactionAtomic(ctx, db.CaptureTransactionAtomicParams{
		TransactionID:      int32(transactionID),
		MerchantCardNumber: merchantCard.CardNumber,
	})
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.Int("transaction_id", transactionID))
	}

	updatedTransaction := &db.UpdateTransactionStatusRow{
		TransactionID:   atomicTransaction.TransactionID,
		TransactionNo:   atomicTransaction.TransactionNo,
		CardNumber:      atomicTransaction.CardNumber,
		Amount:          atomicTransaction.Amount,
		PaymentMethod:   atomicTransaction.PaymentMethod,
		MerchantID:      atomicTransaction.MerchantID,
		TransactionTime: atomicTransaction.TransactionTime,
		Status:          atomicTransaction.Status,
		CreatedAt:       atomicTransaction.CreatedAt,
		UpdatedAt:       atomicTransaction.UpdatedAt,
	}

	logSuccess("Successfully captured transaction", zap.Int("transaction.id", transactionID))

	return updatedTransaction, nil
}

func (s *transactionCommandService) VoidTransaction(ctx context.Context, apiKey string, transactionID int) (*db.UpdateTransactionStatusRow, error) {
	const method = "VoidTransaction"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("transaction_id", transactionID))
	defer func() { end(status) }()

	transaction, err := s.transactionQueryRepository.FindById(ctx, transactionID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.Int("transaction_id", transactionID))
	}

	// Only pending or authorized transactions may be voided (canonical state
	// machine); settled or already-terminal transactions cannot.
	if err := statemachine.ValidateTransition(statemachine.DomainTransaction, transaction.Status, statemachine.Voided); err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.String("current_status", transaction.Status))
	}

	merchant, err := s.merchantRepository.FindByApiKey(ctx, apiKey)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.String("api_key", observability.MaskIdentifier(apiKey)), zap.Int("transaction_id", transactionID))
	}
	if transaction.MerchantID != merchant.MerchantID {
		status = "error"
		err := sharederrors.NewForbiddenError("merchant does not own transaction")
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.Int("transaction_id", transactionID))
	}

	// Void the transaction and release the hold in one atomic SQL statement.
	// Only a 'pending' or 'authorized' transaction can be voided; the row lock
	// plus the status guard make a concurrent double-void impossible. For debit
	// cards the held amount is credited back to the card saldo; credit cards
	// only transition status.
	atomicTransaction, err := s.transactionCommandRepository.VoidTransactionAtomic(ctx, int32(transactionID))
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.Int("transaction_id", transactionID))
	}

	updatedTransaction := &db.UpdateTransactionStatusRow{
		TransactionID:   atomicTransaction.TransactionID,
		TransactionNo:   atomicTransaction.TransactionNo,
		CardNumber:      atomicTransaction.CardNumber,
		Amount:          atomicTransaction.Amount,
		PaymentMethod:   atomicTransaction.PaymentMethod,
		MerchantID:      atomicTransaction.MerchantID,
		TransactionTime: atomicTransaction.TransactionTime,
		Status:          atomicTransaction.Status,
		CreatedAt:       atomicTransaction.CreatedAt,
		UpdatedAt:       atomicTransaction.UpdatedAt,
	}

	logSuccess("Successfully voided transaction", zap.Int("transaction.id", transactionID))

	return updatedTransaction, nil
}

func (s *transactionCommandService) RefundTransaction(ctx context.Context, apiKey string, transactionID int) (*db.UpdateTransactionStatusRow, error) {
	const method = "RefundTransaction"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method, attribute.Int("transaction_id", transactionID))
	defer func() { end(status) }()

	transaction, err := s.transactionQueryRepository.FindById(ctx, transactionID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.Int("transaction_id", transactionID))
	}

	// Only captured transactions may be refunded (canonical state machine).
	if err := statemachine.ValidateTransition(statemachine.DomainTransaction, transaction.Status, statemachine.Refunded); err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.String("current_status", transaction.Status))
	}

	merchant, err := s.merchantRepository.FindByApiKey(ctx, apiKey)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.String("api_key", observability.MaskIdentifier(apiKey)), zap.Int("transaction_id", transactionID))
	}
	if transaction.MerchantID != merchant.MerchantID {
		status = "error"
		err := sharederrors.NewForbiddenError("merchant does not own transaction")
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.Int("transaction_id", transactionID))
	}

	// Debit merchant (reverse settlement)
	merchantCard, err := s.cardRepository.FindCardByUserId(ctx, int(merchant.UserID))
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.Int("merchant_id", int(merchant.MerchantID)))
	}

	// Refund the customer and reverse the merchant settlement in one atomic SQL
	// statement. Only a 'captured' transaction can be refunded; the row lock
	// plus the status guard make a concurrent double-refund impossible. For
	// debit cards the captured amount is credited back to the customer saldo;
	// for credit cards the outstanding balance is reduced. The merchant debit
	// has a non-negative guard (400 when the merchant balance would go
	// negative).
	atomicTransaction, err := s.transactionCommandRepository.RefundTransactionAtomic(ctx, db.RefundTransactionAtomicParams{
		TransactionID:      int32(transactionID),
		MerchantCardNumber: merchantCard.CardNumber,
	})
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionStatusRow](s.logger, err, method, span, zap.String("merchant_card_number", observability.MaskIdentifier(merchantCard.CardNumber)))
	}

	updatedTransaction := &db.UpdateTransactionStatusRow{
		TransactionID:   atomicTransaction.TransactionID,
		TransactionNo:   atomicTransaction.TransactionNo,
		CardNumber:      atomicTransaction.CardNumber,
		Amount:          atomicTransaction.Amount,
		PaymentMethod:   atomicTransaction.PaymentMethod,
		MerchantID:      atomicTransaction.MerchantID,
		TransactionTime: atomicTransaction.TransactionTime,
		Status:          atomicTransaction.Status,
		CreatedAt:       atomicTransaction.CreatedAt,
		UpdatedAt:       atomicTransaction.UpdatedAt,
	}

	logSuccess("Successfully refunded transaction", zap.Int("transaction.id", transactionID))

	return updatedTransaction, nil
}
