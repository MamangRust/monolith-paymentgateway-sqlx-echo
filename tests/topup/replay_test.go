package topup_test

import (
	"context"
	"testing"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-topup/repository"
	"github.com/MamangRust/monolith-payment-gateway-topup/service"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// replayTopupCardRepo records card metadata side-effects so tests can assert
// they are not repeated on idempotent replays.
type replayTopupCardRepo struct {
	updateCalls int
}

func (r *replayTopupCardRepo) FindUserCardByCardNumber(context.Context, string) (*db.GetUserEmailByCardNumberRow, error) {
	return &db.GetUserEmailByCardNumberRow{CardNumber: "4111111111111111", Email: "topup@example.com"}, nil
}

func (replayTopupCardRepo) FindCardByCardNumber(context.Context, string) (*db.GetCardByCardNumberRow, error) {
	return &db.GetCardByCardNumberRow{CardNumber: "4111111111111111"}, nil
}

func (r *replayTopupCardRepo) UpdateCard(context.Context, *requests.UpdateCardRequest) (*db.UpdateCardRow, error) {
	r.updateCalls++
	return &db.UpdateCardRow{}, nil
}

// replayTopupCommandRepo mirrors the repository contract for CreateTopupAtomic:
// the same idempotency key returns the ORIGINAL result (Replayed=true), never
// a second record.
type replayTopupCommandRepo struct {
	seq     int32
	created map[string]*repository.TopupAtomicResult
}

func (r *replayTopupCommandRepo) CreateTopupAtomic(_ context.Context, request *requests.CreateTopupRequest) (*repository.TopupAtomicResult, error) {
	if request.IdempotencyKey != "" {
		if existing, ok := r.created[request.IdempotencyKey]; ok {
			// Signal the replay exactly like the real repository: the original
			// record is returned and side-effects must be skipped.
			existing.Replayed = true
			return existing, nil
		}
	}
	res := &repository.TopupAtomicResult{
		Row: &db.CreateTopupAtomicRow{
			TopupID:     r.seq + 1,
			CardNumber:  request.CardNumber,
			TopupAmount: int32(request.TopupAmount),
			TopupMethod: request.TopupMethod,
			Status:      "success",
		},
	}
	r.seq++
	if request.IdempotencyKey != "" {
		r.created[request.IdempotencyKey] = res
	}
	return res, nil
}

// The remaining methods of the repository contract are not exercised by the
// replay unit tests. They return well-formed dummy values (never errors) so a
// future service change that starts calling them fails with a clear assertion
// instead of a misleading "not implemented" runtime error.
func (replayTopupCommandRepo) CreateTopup(context.Context, *requests.CreateTopupRequest) (*db.CreateTopupRow, error) {
	return &db.CreateTopupRow{}, nil
}
func (replayTopupCommandRepo) UpdateTopup(context.Context, *requests.UpdateTopupRequest) (*db.UpdateTopupRow, error) {
	return &db.UpdateTopupRow{}, nil
}
func (replayTopupCommandRepo) UpdateTopupAtomic(context.Context, *requests.UpdateTopupRequest) (*db.UpdateTopupAtomicRow, error) {
	return &db.UpdateTopupAtomicRow{}, nil
}
func (replayTopupCommandRepo) UpdateSaldoBalanceDelta(context.Context, string, int) error {
	return nil
}
func (replayTopupCommandRepo) UpdateTopupAmount(context.Context, *requests.UpdateTopupAmount) (*db.UpdateTopupAmountRow, error) {
	return &db.UpdateTopupAmountRow{}, nil
}
func (replayTopupCommandRepo) UpdateTopupStatus(context.Context, *requests.UpdateTopupStatus) (*db.UpdateTopupStatusRow, error) {
	return &db.UpdateTopupStatusRow{}, nil
}
func (replayTopupCommandRepo) TrashedTopup(context.Context, int) (*db.Topup, error) {
	return &db.Topup{}, nil
}
func (replayTopupCommandRepo) RestoreTopup(context.Context, int) (*db.Topup, error) {
	return &db.Topup{}, nil
}
func (replayTopupCommandRepo) DeleteTopupPermanent(context.Context, int) (bool, error) {
	return true, nil
}
func (replayTopupCommandRepo) RestoreAllTopup(context.Context) (bool, error) {
	return true, nil
}
func (replayTopupCommandRepo) DeleteAllTopupPermanent(context.Context) (bool, error) {
	return true, nil
}

type replayTopupQueryRepo struct{}

func (replayTopupQueryRepo) FindAllTopups(context.Context, *requests.FindAllTopups) ([]*db.GetTopupsRow, error) {
	return []*db.GetTopupsRow{}, nil
}
func (replayTopupQueryRepo) FindByActive(context.Context, *requests.FindAllTopups) ([]*db.GetActiveTopupsRow, error) {
	return []*db.GetActiveTopupsRow{}, nil
}
func (replayTopupQueryRepo) FindByTrashed(context.Context, *requests.FindAllTopups) ([]*db.GetTrashedTopupsRow, error) {
	return []*db.GetTrashedTopupsRow{}, nil
}
func (replayTopupQueryRepo) FindAllTopupByCardNumber(context.Context, *requests.FindAllTopupsByCardNumber) ([]*db.GetTopupsByCardNumberRow, error) {
	return []*db.GetTopupsByCardNumberRow{}, nil
}
func (replayTopupQueryRepo) FindById(context.Context, int) (*db.GetTopupByIDRow, error) {
	return &db.GetTopupByIDRow{}, nil
}

type replayTopupSaldoRepo struct{}

func (replayTopupSaldoRepo) FindByCardNumber(context.Context, string) (*db.Saldo, error) {
	return &db.Saldo{}, nil
}
func (replayTopupSaldoRepo) UpdateSaldoBalance(context.Context, *requests.UpdateSaldoBalance) (*db.UpdateSaldoBalanceRow, error) {
	return &db.UpdateSaldoBalanceRow{}, nil
}

type replayTopupObservability struct{}

func (replayTopupObservability) StartTracingAndLogging(ctx context.Context, _ string, _ ...attribute.KeyValue) (context.Context, trace.Span, func(string), string, func(string, ...zap.Field)) {
	ctx, span := trace.NewNoopTracerProvider().Tracer("topup-replay").Start(ctx, "topup-replay")
	return ctx, span, func(string) { span.End() }, "success", func(string, ...zap.Field) {}
}
func (replayTopupObservability) RecordMetrics(context.Context, string, string, time.Time) {}

func newReplayTopupService(card *replayTopupCardRepo, command *replayTopupCommandRepo) service.TopupCommandService {
	return service.NewTopupCommandService(&service.TopupCommandDeps{
		CardRepository:         card,
		TopupQueryRepository:   replayTopupQueryRepo{},
		TopupCommandRepository: command,
		SaldoRepository:        replayTopupSaldoRepo{},
		Logger:                 &logger.Logger{Log: zap.NewNop()},
		Observability:          replayTopupObservability{},
	})
}

func TestCreateTopupReplayReturnsOriginalAndSkipsSideEffects(t *testing.T) {
	card := &replayTopupCardRepo{}
	command := &replayTopupCommandRepo{created: map[string]*repository.TopupAtomicResult{}}
	svc := newReplayTopupService(card, command)

	req := &requests.CreateTopupRequest{
		CardNumber:     "4111111111111111",
		TopupAmount:    100000,
		TopupMethod:    "gopay",
		IdempotencyKey: "topup:create:user:42:op-1",
	}

	first, err := svc.CreateTopup(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, first)
	require.Equal(t, int32(1), first.TopupID)
	require.Equal(t, "success", first.Status)
	require.Equal(t, 1, card.updateCalls, "first create must run the card metadata side-effect once")

	// Replay with the same key: same response, and the card metadata
	// side-effect must NOT run again.
	second, err := svc.CreateTopup(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, second)
	require.Equal(t, first.TopupID, second.TopupID, "replay must return the original record id")
	require.Equal(t, first.TopupAmount, second.TopupAmount)
	require.Equal(t, "success", second.Status)
	require.Equal(t, 1, card.updateCalls, "replay must not repeat the card metadata side-effect")
	require.Len(t, command.created, 1, "the key must be persisted exactly once")
}

func TestCreateTopupWithoutKeyCreatesIndependently(t *testing.T) {
	card := &replayTopupCardRepo{}
	command := &replayTopupCommandRepo{created: map[string]*repository.TopupAtomicResult{}}
	svc := newReplayTopupService(card, command)

	req := &requests.CreateTopupRequest{
		CardNumber:  "4111111111111111",
		TopupAmount: 100000,
		TopupMethod: "gopay",
	}

	first, err := svc.CreateTopup(context.Background(), req)
	require.NoError(t, err)
	second, err := svc.CreateTopup(context.Background(), req)
	require.NoError(t, err)

	require.NotEqual(t, first.TopupID, second.TopupID, "without a key every request creates a new record")
	require.Equal(t, 2, card.updateCalls)
}
