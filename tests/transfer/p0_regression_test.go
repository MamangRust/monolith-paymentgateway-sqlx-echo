package transfer_test

import (
	"context"
	"errors"
	"testing"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-transfer/repository"
	"github.com/MamangRust/monolith-payment-gateway-transfer/service"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type p0TransferCardRepo struct{}

func (p0TransferCardRepo) FindUserCardByCardNumber(context.Context, string) (*db.GetUserEmailByCardNumberRow, error) {
	return &db.GetUserEmailByCardNumberRow{CardNumber: "sender", Email: "sender@example.com"}, nil
}

func (p0TransferCardRepo) FindCardByCardNumber(context.Context, string) (*db.GetCardByCardNumberRow, error) {
	return &db.GetCardByCardNumberRow{CardNumber: "receiver"}, nil
}

type p0TransferSaldoRepo struct {
	balances map[string]int32
}

func (r *p0TransferSaldoRepo) FindByCardNumber(_ context.Context, cardNumber string) (*db.Saldo, error) {
	balance, ok := r.balances[cardNumber]
	if !ok {
		return nil, errors.New("saldo not found")
	}
	return &db.Saldo{CardNumber: cardNumber, TotalBalance: balance}, nil
}

func (r *p0TransferSaldoRepo) UpdateSaldoBalance(_ context.Context, request *requests.UpdateSaldoBalance) (*db.UpdateSaldoBalanceRow, error) {
	if _, ok := r.balances[request.CardNumber]; !ok {
		return nil, errors.New("saldo not found")
	}
	r.balances[request.CardNumber] = int32(request.TotalBalance)
	return &db.UpdateSaldoBalanceRow{CardNumber: request.CardNumber, TotalBalance: int32(request.TotalBalance)}, nil
}

type p0TransferCommandRepo struct {
	createErr        error
	statusErr        error
	createTransferID int32
	statusCalls      int
	// balances simulates the atomic SQL settlement: debit sender, credit
	// receiver in one call, mirroring CreateTransferAtomic.
	balances map[string]int32
}

func (r *p0TransferCommandRepo) CreateTransfer(context.Context, *requests.CreateTransferRequest) (*db.CreateTransferRow, error) {
	if r.createErr != nil {
		return nil, r.createErr
	}
	return &db.CreateTransferRow{
		TransferID:     r.createTransferID,
		Status:         "success",
		TransferAmount: 50000,
	}, nil
}

func (r *p0TransferCommandRepo) CreateTransferAtomic(_ context.Context, req *requests.CreateTransferRequest) (*repository.TransferAtomicResult, error) {
	if r.createErr != nil {
		return nil, r.createErr
	}
	if r.balances != nil {
		r.balances[req.TransferFrom] -= int32(req.TransferAmount)
		r.balances[req.TransferTo] += int32(req.TransferAmount)
	}
	return &repository.TransferAtomicResult{
		Row: &db.CreateTransferAtomicRow{
			TransferID:     r.createTransferID,
			TransferFrom:   req.TransferFrom,
			TransferTo:     req.TransferTo,
			TransferAmount: int32(req.TransferAmount),
			Status:         "success",
		},
	}, nil
}

func (r *p0TransferCommandRepo) UpdateTransferStatus(context.Context, *requests.UpdateTransferStatus) (*db.UpdateTransferStatusRow, error) {
	r.statusCalls++
	if r.statusErr != nil {
		return nil, r.statusErr
	}
	return &db.UpdateTransferStatusRow{}, nil
}

func (r *p0TransferCommandRepo) UpdateTransferAtomic(context.Context, db.UpdateTransferAtomicParams) (*db.UpdateTransferAtomicRow, error) {
	return &db.UpdateTransferAtomicRow{}, nil
}

func (r *p0TransferCommandRepo) UpdateTransferSettlementDelta(_ context.Context, senderCard, receiverCard string, senderDelta, receiverDelta int) error {
	if r.balances != nil {
		r.balances[senderCard] += int32(senderDelta)
		r.balances[receiverCard] += int32(receiverDelta)
	}
	return nil
}

// The remaining repository methods are not exercised by the P0 regression
// tests. They return well-formed dummy values (never errors) so a future
// service change that starts calling them fails with a clear assertion
// instead of a misleading "not implemented" runtime error.
func (p0TransferCommandRepo) UpdateTransfer(context.Context, *requests.UpdateTransferRequest) (*db.UpdateTransferRow, error) {
	return &db.UpdateTransferRow{}, nil
}
func (p0TransferCommandRepo) UpdateTransferAmount(context.Context, *requests.UpdateTransferAmountRequest) (*db.UpdateTransferAmountRow, error) {
	return &db.UpdateTransferAmountRow{}, nil
}
func (p0TransferCommandRepo) TrashedTransfer(context.Context, int) (*db.Transfer, error) {
	return &db.Transfer{}, nil
}
func (p0TransferCommandRepo) RestoreTransfer(context.Context, int) (*db.Transfer, error) {
	return &db.Transfer{}, nil
}
func (p0TransferCommandRepo) DeleteTransferPermanent(context.Context, int) (bool, error) {
	return true, nil
}
func (p0TransferCommandRepo) RestoreAllTransfer(context.Context) (bool, error) {
	return true, nil
}
func (p0TransferCommandRepo) DeleteAllTransferPermanent(context.Context) (bool, error) {
	return true, nil
}

type p0TransferQueryRepo struct{}

func (p0TransferQueryRepo) FindAll(context.Context, *requests.FindAllTransfers) ([]*db.GetTransfersRow, error) {
	return []*db.GetTransfersRow{}, nil
}
func (p0TransferQueryRepo) FindByActive(context.Context, *requests.FindAllTransfers) ([]*db.GetActiveTransfersRow, error) {
	return []*db.GetActiveTransfersRow{}, nil
}
func (p0TransferQueryRepo) FindByTrashed(context.Context, *requests.FindAllTransfers) ([]*db.GetTrashedTransfersRow, error) {
	return []*db.GetTrashedTransfersRow{}, nil
}
func (p0TransferQueryRepo) FindById(context.Context, int) (*db.GetTransferByIDRow, error) {
	return &db.GetTransferByIDRow{}, nil
}
func (p0TransferQueryRepo) FindTransferByTransferFrom(context.Context, string) ([]*db.GetTransfersBySourceCardRow, error) {
	return []*db.GetTransfersBySourceCardRow{}, nil
}
func (p0TransferQueryRepo) FindTransferByTransferTo(context.Context, string) ([]*db.GetTransfersByDestinationCardRow, error) {
	return []*db.GetTransfersByDestinationCardRow{}, nil
}

type p0TransferObservability struct{}

func (p0TransferObservability) StartTracingAndLogging(ctx context.Context, _ string, _ ...attribute.KeyValue) (context.Context, trace.Span, func(string), string, func(string, ...zap.Field)) {
	ctx, span := trace.NewNoopTracerProvider().Tracer("transfer-p0").Start(ctx, "transfer-p0")
	return ctx, span, func(string) { span.End() }, "success", func(string, ...zap.Field) {}
}
func (p0TransferObservability) RecordMetrics(context.Context, string, string, time.Time) {}

func newP0TransferService(saldo *p0TransferSaldoRepo, command repository.TransferCommandRepository) service.TransferCommandService {
	return service.NewTransferCommandService(&service.TransferCommandDeps{
		CardRepository:            p0TransferCardRepo{},
		SaldoRepository:           saldo,
		TransferQueryRepository:   p0TransferQueryRepo{},
		TransferCommandRepository: command,
		Logger:                    &logger.Logger{Log: zap.NewNop()},
		Observability:             p0TransferObservability{},
	})
}

func TestCreateTransferAtomicFailureLeavesBalancesUntouched(t *testing.T) {
	// The atomic operation either settles fully or not at all. When the
	// repository call fails, the service must not touch any balance and must
	// not panic (regression: previous non-atomic flow needed manual rollback).
	balances := map[string]int32{"sender": 100000, "receiver": 50000}
	saldo := &p0TransferSaldoRepo{balances: balances}
	svc := newP0TransferService(saldo, &p0TransferCommandRepo{createErr: errors.New("create transfer failed")})

	result, err := svc.CreateTransaction(context.Background(), &requests.CreateTransferRequest{
		TransferFrom:   "sender",
		TransferTo:     "receiver",
		TransferAmount: 50000,
	})

	require.Error(t, err)
	require.Nil(t, result)
	require.Equal(t, int32(100000), balances["sender"])
	require.Equal(t, int32(50000), balances["receiver"])
}

func TestCreateTransferAtomicSettlesOnceWithoutStatusUpdate(t *testing.T) {
	balances := map[string]int32{"sender": 100000, "receiver": 50000}
	saldo := &p0TransferSaldoRepo{balances: balances}
	command := &p0TransferCommandRepo{
		createTransferID: 7,
		statusErr:        errors.New("status update must not be called"),
		balances:         balances,
	}
	svc := newP0TransferService(saldo, command)

	result, err := svc.CreateTransaction(context.Background(), &requests.CreateTransferRequest{
		TransferFrom:   "sender",
		TransferTo:     "receiver",
		TransferAmount: 50000,
	})

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, int32(7), result.TransferID)
	require.Equal(t, "success", result.Status)
	require.Zero(t, command.statusCalls)
	require.Equal(t, int32(50000), balances["sender"])
	require.Equal(t, int32(100000), balances["receiver"])
}
