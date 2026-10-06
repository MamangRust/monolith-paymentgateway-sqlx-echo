package service

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	mencache "github.com/MamangRust/monolith-payment-gateway-saldo/redis"
	"github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type SaldoCommandParams struct {
	Cache                  mencache.SaldoCommandCache
	SaldoCommandRepository repository.SaldoCommandRepository
	CardRepository         repository.CardRepository
	Logger                 logger.LoggerInterface
	Observability          observability.TraceLoggerObservability
}

type saldoCommandService struct {
	ctx                    context.Context
	cache                  mencache.SaldoCommandCache
	cardRepository         repository.CardRepository
	logger                 logger.LoggerInterface
	SaldoCommandRepository repository.SaldoCommandRepository
	observability          observability.TraceLoggerObservability
}

func NewSaldoCommandService(params *SaldoCommandParams) SaldoCommandService {
	return &saldoCommandService{
		cache:                  params.Cache,
		SaldoCommandRepository: params.SaldoCommandRepository,
		cardRepository:         params.CardRepository,
		logger:                 params.Logger,
		observability:          params.Observability,
	}
}

func (s *saldoCommandService) CreateSaldo(ctx context.Context, request *requests.CreateSaldoRequest) (*db.CreateSaldoRow, error) {
	const method = "CreateSaldo"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()
	if request == nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateSaldoRow](s.logger, sharedErrors.NewBadRequestError("saldo request cannot be nil"), method, span)
	}

	_, err := s.cardRepository.FindCardByCardNumber(ctx, request.CardNumber)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateSaldoRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	}

	res, err := s.SaldoCommandRepository.CreateSaldo(ctx, request)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateSaldoRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	}

	s.cache.InvalidateSaldoCache(ctx)

	logSuccess("Successfully created saldo record", zap.String("card_number", observability.MaskIdentifier(request.CardNumber)), zap.Float64("amount", float64(request.TotalBalance)))

	return res, nil
}

func (s *saldoCommandService) InvalidateSaldoCache(ctx context.Context) {
	if s.cache != nil {
		s.cache.InvalidateSaldoCache(ctx)
	}
}

func (s *saldoCommandService) CreateSaldoIfNotExists(ctx context.Context, request *requests.CreateSaldoRequest) error {
	const method = "CreateSaldoIfNotExists"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()

	if request == nil {
		status = "error"
		_, err := errorhandler.HandleError[*db.CreateSaldoRow](s.logger, sharedErrors.NewBadRequestError("saldo request cannot be nil"), method, span)
		return err
	}

	if _, err := s.cardRepository.FindCardByCardNumber(ctx, request.CardNumber); err != nil {
		status = "error"
		_, err = errorhandler.HandleError[*db.CreateSaldoRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
		return err
	}

	if err := s.SaldoCommandRepository.CreateSaldoIfNotExists(ctx, request); err != nil {
		status = "error"
		_, err = errorhandler.HandleError[*db.CreateSaldoRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
		return err
	}

	logSuccess("Ensured saldo record exists", zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	return nil
}

func (s *saldoCommandService) UpdateSaldo(ctx context.Context, request *requests.UpdateSaldoRequest) (*db.UpdateSaldoRow, error) {
	const method = "UpdateSaldo"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()
	if request == nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateSaldoRow](s.logger, sharedErrors.NewBadRequestError("saldo request cannot be nil"), method, span)
	}

	_, err := s.cardRepository.FindCardByCardNumber(ctx, request.CardNumber)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateSaldoRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	}

	res, err := s.SaldoCommandRepository.UpdateSaldo(ctx, request)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateSaldoRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	}

	s.cache.InvalidateSaldoCache(ctx)

	logSuccess("Successfully updated saldo record", zap.String("card_number", observability.MaskIdentifier(request.CardNumber)), zap.Float64("amount", float64(request.TotalBalance)))

	return res, nil
}

func (s *saldoCommandService) UpdateSaldoBalance(ctx context.Context, request *requests.UpdateSaldoBalance) (*db.UpdateSaldoBalanceRow, error) {
	const method = "UpdateSaldoBalance"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()
	if request == nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateSaldoBalanceRow](s.logger, sharedErrors.NewBadRequestError("saldo request cannot be nil"), method, span)
	}

	_, err := s.cardRepository.FindCardByCardNumber(ctx, request.CardNumber)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateSaldoBalanceRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	}

	res, err := s.SaldoCommandRepository.UpdateSaldoBalance(ctx, request)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateSaldoBalanceRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	}

	s.cache.InvalidateSaldoCache(ctx)

	logSuccess("Successfully updated saldo balance", zap.String("card_number", observability.MaskIdentifier(request.CardNumber)), zap.Float64("amount", float64(request.TotalBalance)))

	return res, nil
}

func (s *saldoCommandService) UpdateSaldoWithdraw(ctx context.Context, request *requests.UpdateSaldoWithdraw) (*db.UpdateSaldoWithdrawRow, error) {
	const method = "UpdateSaldoWithdraw"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)
	defer func() { end(status) }()
	if request == nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateSaldoWithdrawRow](s.logger, sharedErrors.NewBadRequestError("saldo request cannot be nil"), method, span)
	}

	_, err := s.cardRepository.FindCardByCardNumber(ctx, request.CardNumber)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateSaldoWithdrawRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	}

	res, err := s.SaldoCommandRepository.UpdateSaldoWithdraw(ctx, request)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateSaldoWithdrawRow](s.logger, err, method, span, zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))
	}

	s.cache.InvalidateSaldoCache(ctx)

	logSuccess("Successfully updated saldo withdraw record", zap.String("card_number", observability.MaskIdentifier(request.CardNumber)))

	return res, nil
}

func (s *saldoCommandService) TrashSaldo(ctx context.Context, saldo_id int) (*db.Saldo, error) {
	const method = "TrashSaldo"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("saldo_id", saldo_id))

	defer func() {
		end(status)
	}()

	res, err := s.SaldoCommandRepository.TrashedSaldo(ctx, saldo_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Saldo](
			s.logger,
			err,
			method,
			span,

			zap.Int("saldo_id", saldo_id),
		)
	}

	logSuccess("Successfully trashed saldo", zap.Int("saldo_id", saldo_id))

	return res, nil
}

func (s *saldoCommandService) RestoreSaldo(ctx context.Context, saldo_id int) (*db.Saldo, error) {
	const method = "RestoreSaldo"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("saldo_id", saldo_id))

	defer func() {
		end(status)
	}()

	res, err := s.SaldoCommandRepository.RestoreSaldo(ctx, saldo_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Saldo](
			s.logger,
			err,
			method,
			span,

			zap.Int("saldo_id", saldo_id),
		)
	}

	logSuccess("Successfully restored saldo", zap.Int("saldo_id", saldo_id))

	return res, nil
}

func (s *saldoCommandService) DeleteSaldoPermanent(ctx context.Context, saldo_id int) (bool, error) {
	const method = "DeleteSaldoPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("saldo_id", saldo_id))

	defer func() {
		end(status)
	}()

	_, err := s.SaldoCommandRepository.DeleteSaldoPermanent(ctx, saldo_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,

			zap.Int("saldo_id", saldo_id),
		)
	}

	logSuccess("Successfully deleted saldo permanently", zap.Int("saldo_id", saldo_id))

	return true, nil
}

func (s *saldoCommandService) RestoreAllSaldo(ctx context.Context) (bool, error) {
	const method = "RestoreAllSaldo"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	_, err := s.SaldoCommandRepository.RestoreAllSaldo(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,
		)
	}

	logSuccess("Successfully restored all saldo")
	return true, nil
}

func (s *saldoCommandService) DeleteAllSaldoPermanent(ctx context.Context) (bool, error) {
	const method = "DeleteAllSaldoPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	_, err := s.SaldoCommandRepository.DeleteAllSaldoPermanent(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,
		)
	}

	logSuccess("Successfully deleted all saldo permanently")
	return true, nil
}
