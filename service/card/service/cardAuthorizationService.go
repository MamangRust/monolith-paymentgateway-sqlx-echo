package service

import (
	"context"

	mencache "github.com/MamangRust/monolith-payment-gateway-card/redis"
	cardauthmencache "github.com/MamangRust/monolith-payment-gateway-card/redis/auth"
	"github.com/MamangRust/monolith-payment-gateway-card/repository"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharederrorhandler "github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// CardAuthorizationServiceDeps defines dependencies for cardAuthorizationService.
type CardAuthorizationServiceDeps struct {
	Cache                  mencache.CardCommandCache
	AuthCache              cardauthmencache.CardAuthCache
	CardCommandRepository  repository.CardCommandRepository
	CardAuthRepository     repository.CardAuthTransactionRepository
	CardPaymentRepository  repository.CardPaymentRepository
	CardRewardRepository   repository.CardRewardRepository
	BillingCycleRepository repository.BillingCycleRepository
	CardQueryRepository    repository.CardQueryRepository
	Logger                 logger.LoggerInterface
	Observability          observability.TraceLoggerObservability
}

// cardAuthorizationService implements CardAuthorizationService.
type cardAuthorizationService struct {
	cache                       mencache.CardCommandCache
	authCache                   cardauthmencache.CardAuthCache
	cardCommandRepository       repository.CardCommandRepository
	cardAuthRepository          repository.CardAuthTransactionRepository
	cardCreditAccountRepository repository.CardCommandRepository
	cardPaymentRepository       repository.CardPaymentRepository
	cardRewardRepository        repository.CardRewardRepository
	billingCycleRepository      repository.BillingCycleRepository
	cardQueryRepository         repository.CardQueryRepository
	logger                      logger.LoggerInterface
	observability               observability.TraceLoggerObservability
}

// NewCardAuthorizationService creates a new cardAuthorizationService.
func NewCardAuthorizationService(deps *CardAuthorizationServiceDeps) CardAuthorizationService {
	return &cardAuthorizationService{
		cache:                       deps.Cache,
		authCache:                   deps.AuthCache,
		cardCommandRepository:       deps.CardCommandRepository,
		cardAuthRepository:          deps.CardAuthRepository,
		cardCreditAccountRepository: deps.CardCommandRepository,
		cardPaymentRepository:       deps.CardPaymentRepository,
		cardRewardRepository:        deps.CardRewardRepository,
		billingCycleRepository:      deps.BillingCycleRepository,
		cardQueryRepository:         deps.CardQueryRepository,
		logger:                      deps.Logger,
		observability:               deps.Observability,
	}
}

// Authorize authorizes a card transaction.
func (s *cardAuthorizationService) Authorize(ctx context.Context, request *requests.AuthorizeCardRequest) (*db.CardAuthTransaction, error) {
	const method = "Authorize"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(request.CardNumber)),
		attribute.Int64("amount", request.Amount),
		attribute.String("mcc", request.Mcc),
		attribute.Int("merchant_id", request.MerchantID),
	)

	defer func() {
		end(status)
	}()

	// Idempotency check
	if request.IdempotencyKey != "" {
		existing, err := s.cardAuthRepository.FindByIdempotencyKey(ctx, request.IdempotencyKey)
		if err == nil && existing != nil {
			span.SetAttributes(attribute.String("idempotent", "true"))
			logSuccess("Idempotent request — returning existing auth transaction", zap.String("txn_id", existing.TxnID), zap.String("status", existing.Status))
			return existing, nil
		}
	}

	// Find card by card number
	card, err := s.cardQueryRepository.FindCardByCardNumber(ctx, request.CardNumber)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.CardAuthTransaction](
			s.logger,
			card_errors.ErrFailedCreateCard,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(request.CardNumber)),
		)
	}

	// Verify card is active
	if card.Status != "active" {
		status = "error"
		return sharederrorhandler.HandleError[*db.CardAuthTransaction](
			s.logger,
			card_errors.ErrFailedCreateCard,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(request.CardNumber)),
			zap.String("card_status", card.Status),
		)
	}

	// For credit cards: check credit limit
	if card.CardType == "credit" {
		availableCredit := int64(card.CreditLimit) - int64(card.OutstandingBalance)
		if availableCredit < request.Amount {
			status = "error"
			return sharederrorhandler.HandleError[*db.CardAuthTransaction](
				s.logger,
				card_errors.ErrFailedCreateCard,
				method,
				span,
				zap.String("card_number", observability.MaskIdentifier(request.CardNumber)),
				zap.Int64("available_credit", availableCredit),
				zap.Int64("requested_amount", request.Amount),
			)
		}
	}

	// Insert pending auth transaction
	result, err := s.cardAuthRepository.InsertPending(ctx, request)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.CardAuthTransaction](
			s.logger,
			err,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(request.CardNumber)),
		)
	}

	txnID := result.TxnID

	logSuccess("Successfully authorized card transaction",
		zap.String("txn_id", txnID),
		zap.String("card_number", observability.MaskIdentifier(request.CardNumber)),
		zap.Int64("amount", request.Amount),
	)

	return result, nil
}

// Reverse reverses an authorized transaction.
func (s *cardAuthorizationService) Reverse(ctx context.Context, request *requests.ReverseTransactionRequest) (*db.CardAuthTransaction, error) {
	const method = "Reverse"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("txn_id", request.TxnID),
		attribute.String("card_number", observability.MaskIdentifier(request.CardNumber)),
		attribute.Int64("amount", request.Amount),
	)

	defer func() {
		end(status)
	}()

	// Idempotency check
	if request.IdempotencyKey != "" {
		existing, err := s.cardAuthRepository.FindByIdempotencyKey(ctx, request.IdempotencyKey)
		if err == nil && existing != nil {
			span.SetAttributes(attribute.String("idempotent", "true"))
			logSuccess("Idempotent request — returning existing reverse transaction", zap.String("txn_id", existing.TxnID), zap.String("status", existing.Status))
			return existing, nil
		}
	}

	// Verify the transaction exists and is approved
	existing, err := s.cardAuthRepository.FindByTxnID(ctx, request.TxnID)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.CardAuthTransaction](
			s.logger,
			err,
			method,
			span,
			zap.String("txn_id", request.TxnID),
		)
	}

	if existing.Status != "approved" {
		status = "error"
		return sharederrorhandler.HandleError[*db.CardAuthTransaction](
			s.logger,
			card_errors.ErrFailedCreateCard,
			method,
			span,
			zap.String("txn_id", request.TxnID),
			zap.String("current_status", existing.Status),
		)
	}

	// Reverse the transaction
	result, err := s.cardAuthRepository.Reverse(ctx, request.TxnID)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.CardAuthTransaction](
			s.logger,
			err,
			method,
			span,
			zap.String("txn_id", request.TxnID),
		)
	}

	if s.authCache != nil {
		s.authCache.SetByTxnID(ctx, request.TxnID, result)
	}

	logSuccess("Successfully reversed card transaction",
		zap.String("txn_id", request.TxnID),
		zap.String("card_number", observability.MaskIdentifier(request.CardNumber)),
		zap.Int64("amount", request.Amount),
	)

	return result, nil
}

// GetAuthTransaction retrieves an auth transaction by its transaction ID.
func (s *cardAuthorizationService) GetAuthTransaction(ctx context.Context, txnID string) (*db.CardAuthTransaction, error) {
	const method = "GetAuthTransaction"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("txn_id", txnID),
	)

	defer func() {
		end(status)
	}()

	if s.authCache != nil {
		if cached, found := s.authCache.GetByTxnID(ctx, txnID); found {
			logSuccess("Successfully retrieved auth transaction from cache", zap.String("txn_id", txnID))
			return cached, nil
		}
	}

	result, err := s.cardAuthRepository.FindByTxnID(ctx, txnID)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.CardAuthTransaction](
			s.logger,
			err,
			method,
			span,
			zap.String("txn_id", txnID),
		)
	}

	if s.authCache != nil {
		s.authCache.SetByTxnID(ctx, txnID, result)
	}

	logSuccess("Successfully retrieved auth transaction", zap.String("txn_id", txnID))

	return result, nil
}

// GetAuthTransactionsByCardNumber retrieves auth transactions by card number with pagination.
func (s *cardAuthorizationService) GetAuthTransactionsByCardNumber(ctx context.Context, cardNumber string, page, pageSize int) ([]*db.GetAuthTransactionsByCardNumberRow, error) {
	const method = "GetAuthTransactionsByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(cardNumber)),
		attribute.Int("page", page),
		attribute.Int("page_size", pageSize),
	)

	defer func() {
		end(status)
	}()

	results, err := s.cardAuthRepository.FindByCardNumber(ctx, cardNumber, page, pageSize)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetAuthTransactionsByCardNumberRow](
			s.logger,
			err,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		)
	}

	logSuccess("Successfully retrieved auth transactions by card number",
		zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		zap.Int("count", len(results)),
	)

	return results, nil
}
