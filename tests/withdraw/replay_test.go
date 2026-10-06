package withdraw_test

import (
	"context"
	"testing"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-withdraw/repository"
	"github.com/MamangRust/monolith-payment-gateway-withdraw/service"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type replayWithdrawCardRepo struct{}

func (replayWithdrawCardRepo) FindUserCardByCardNumber(context.Context, string) (*db.GetUserEmailByCardNumberRow, error) {
	return &db.GetUserEmailByCardNumberRow{CardNumber: "4111111111111111", Email: "withdraw@example.com"}, nil
}

// replayWithdrawCommandRepo mirrors the repository contract for
// CreateWithdrawAtomic: the same idempotency key returns the ORIGINAL result
// (Replayed=true), never a second withdrawal.
type replayWithdrawCommandRepo struct {
	seq     int32
	created map[string]*repository.WithdrawAtomicResult
}

func (r *replayWithdrawCommandRepo) CreateWithdrawAtomic(_ context.Context, req *requests.CreateWithdrawRequest) (*repository.WithdrawAtomicResult, error) {
	if req.IdempotencyKey != "" {
		if existing, ok := r.created[req.IdempotencyKey]; ok {
			// Signal the replay exactly like the real repository: the original
			// record is returned and side-effects must be skipped.
			existing.Replayed = true
			return existing, nil
		}
	}
	res := &repository.WithdrawAtomicResult{
		Row: &db.CreateWithdrawAtomicRow{
			WithdrawID:     r.seq + 1,
			CardNumber:     req.CardNumber,
			WithdrawAmount: int32(req.WithdrawAmount),
			Status:         "success",
		},
	}
	r.seq++
	if req.IdempotencyKey != "" {
		r.created[req.IdempotencyKey] = res
	}
	return res, nil
}

// The remaining methods of the repository contract are not exercised by the
// replay unit tests. They return well-formed dummy values (never errors) so a
// future service change that starts calling them fails with a clear assertion
// instead of a misleading "not implemented" runtime error.
func (replayWithdrawCommandRepo) CreateWithdraw(context.Context, *requests.CreateWithdrawRequest) (*db.CreateWithdrawRow, error) {
	return &db.CreateWithdrawRow{}, nil
}
func (replayWithdrawCommandRepo) UpdateWithdraw(context.Context, *requests.UpdateWithdrawRequest) (*db.UpdateWithdrawRow, error) {
	return &db.UpdateWithdrawRow{}, nil
}
func (replayWithdrawCommandRepo) UpdateSaldoBalanceDelta(context.Context, string, int) error {
	return nil
}

func (replayWithdrawCommandRepo) UpdateWithdrawAtomic(context.Context, *requests.UpdateWithdrawRequest) (*db.UpdateWithdrawAtomicRow, error) {
	return &db.UpdateWithdrawAtomicRow{}, nil
}
func (replayWithdrawCommandRepo) UpdateWithdrawStatus(context.Context, *requests.UpdateWithdrawStatus) (*db.UpdateWithdrawStatusRow, error) {
	return &db.UpdateWithdrawStatusRow{}, nil
}
func (replayWithdrawCommandRepo) TrashedWithdraw(context.Context, int) (*db.Withdraw, error) {
	return &db.Withdraw{}, nil
}
func (replayWithdrawCommandRepo) RestoreWithdraw(context.Context, int) (*db.Withdraw, error) {
	return &db.Withdraw{}, nil
}
func (replayWithdrawCommandRepo) DeleteWithdrawPermanent(context.Context, int) (bool, error) {
	return true, nil
}
func (replayWithdrawCommandRepo) RestoreAllWithdraw(context.Context) (bool, error) {
	return true, nil
}
func (replayWithdrawCommandRepo) DeleteAllWithdrawPermanent(context.Context) (bool, error) {
	return true, nil
}

type replayWithdrawQueryRepo struct{}

func (replayWithdrawQueryRepo) FindAll(context.Context, *requests.FindAllWithdraws) ([]*db.GetWithdrawsRow, error) {
	return []*db.GetWithdrawsRow{}, nil
}
func (replayWithdrawQueryRepo) FindByActive(context.Context, *requests.FindAllWithdraws) ([]*db.GetActiveWithdrawsRow, error) {
	return []*db.GetActiveWithdrawsRow{}, nil
}
func (replayWithdrawQueryRepo) FindByTrashed(context.Context, *requests.FindAllWithdraws) ([]*db.GetTrashedWithdrawsRow, error) {
	return []*db.GetTrashedWithdrawsRow{}, nil
}
func (replayWithdrawQueryRepo) FindAllByCardNumber(context.Context, *requests.FindAllWithdrawCardNumber) ([]*db.GetWithdrawsByCardNumberRow, error) {
	return []*db.GetWithdrawsByCardNumberRow{}, nil
}
func (replayWithdrawQueryRepo) FindById(context.Context, int) (*db.GetWithdrawByIDRow, error) {
	return &db.GetWithdrawByIDRow{}, nil
}

type replayWithdrawSaldoRepo struct{}

func (replayWithdrawSaldoRepo) FindByCardNumber(context.Context, string) (*db.Saldo, error) {
	return &db.Saldo{}, nil
}
func (replayWithdrawSaldoRepo) UpdateSaldoBalance(context.Context, *requests.UpdateSaldoBalance) (*db.UpdateSaldoBalanceRow, error) {
	return &db.UpdateSaldoBalanceRow{}, nil
}
func (replayWithdrawSaldoRepo) UpdateSaldoWithdraw(context.Context, *requests.UpdateSaldoWithdraw) (*db.UpdateSaldoWithdrawRow, error) {
	return &db.UpdateSaldoWithdrawRow{}, nil
}

type replayWithdrawObservability struct{}

func (replayWithdrawObservability) StartTracingAndLogging(ctx context.Context, _ string, _ ...attribute.KeyValue) (context.Context, trace.Span, func(string), string, func(string, ...zap.Field)) {
	ctx, span := trace.NewNoopTracerProvider().Tracer("withdraw-replay").Start(ctx, "withdraw-replay")
	return ctx, span, func(string) { span.End() }, "success", func(string, ...zap.Field) {}
}
func (replayWithdrawObservability) RecordMetrics(context.Context, string, string, time.Time) {}

func newReplayWithdrawService(command *replayWithdrawCommandRepo) service.WithdrawCommandService {
	return service.NewWithdrawCommandService(&service.WithdrawCommandServiceDeps{
		CardRepository:    replayWithdrawCardRepo{},
		CommandRepository: command,
		QueryRepository:   replayWithdrawQueryRepo{},
		SaldoRepository:   replayWithdrawSaldoRepo{},
		Logger:            &logger.Logger{Log: zap.NewNop()},
		Observability:     replayWithdrawObservability{},
	})
}

func TestCreateWithdrawReplayReturnsOriginal(t *testing.T) {
	command := &replayWithdrawCommandRepo{created: map[string]*repository.WithdrawAtomicResult{}}
	svc := newReplayWithdrawService(command)

	req := &requests.CreateWithdrawRequest{
		CardNumber:     "4111111111111111",
		WithdrawAmount: 100000,
		WithdrawTime:   time.Now(),
		IdempotencyKey: "withdraw:create:user:42:op-1",
	}

	first, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, first)
	require.Equal(t, int32(1), first.WithdrawID)
	require.Equal(t, "success", first.Status)

	// Replay with the same key: same withdrawal id, no second record.
	second, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, second)
	require.Equal(t, first.WithdrawID, second.WithdrawID, "replay must return the original withdrawal id")
	require.Equal(t, first.WithdrawAmount, second.WithdrawAmount)
	require.Len(t, command.created, 1, "the key must be persisted exactly once")
}

func TestCreateWithdrawWithoutKeyDebitsTwice(t *testing.T) {
	command := &replayWithdrawCommandRepo{created: map[string]*repository.WithdrawAtomicResult{}}
	svc := newReplayWithdrawService(command)

	req := &requests.CreateWithdrawRequest{
		CardNumber:     "4111111111111111",
		WithdrawAmount: 100000,
		WithdrawTime:   time.Now(),
	}

	first, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	second, err := svc.Create(context.Background(), req)
	require.NoError(t, err)

	require.NotEqual(t, first.WithdrawID, second.WithdrawID, "without a key every request creates a new withdrawal")
}
