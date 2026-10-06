package transaction_test

import (
	"context"
	"testing"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-transaction/repository"
	"github.com/MamangRust/monolith-payment-gateway-transaction/service"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type replayTransactionMerchantRepo struct{}

func (replayTransactionMerchantRepo) FindByApiKey(context.Context, string) (*db.GetMerchantByApiKeyRow, error) {
	return &db.GetMerchantByApiKeyRow{MerchantID: 1, UserID: 7, ApiKey: "key"}, nil
}

type replayTransactionCardRepo struct{}

func (replayTransactionCardRepo) FindUserCardByCardNumber(context.Context, string) (*db.GetUserEmailByCardNumberRow, error) {
	return &db.GetUserEmailByCardNumberRow{CardNumber: "4111111111111111", Email: "user@example.com"}, nil
}

func (replayTransactionCardRepo) FindCardByUserId(context.Context, int) (*db.GetCardByUserIDRow, error) {
	return &db.GetCardByUserIDRow{CardNumber: "5111111111111111"}, nil
}

func (replayTransactionCardRepo) FindCardByCardNumber(context.Context, string) (*db.GetCardByCardNumberRow, error) {
	return &db.GetCardByCardNumberRow{}, nil
}
func (replayTransactionCardRepo) UpdateCard(context.Context, *requests.UpdateCardRequest) (*db.UpdateCardRow, error) {
	return &db.UpdateCardRow{}, nil
}
func (replayTransactionCardRepo) UpdateCardOutstandingBalance(context.Context, int, int) (*db.UpdateOutstandingBalanceRow, error) {
	return &db.UpdateOutstandingBalanceRow{}, nil
}
func (replayTransactionCardRepo) AddRewardPoints(context.Context, int, int) (*db.AddRewardPointsRow, error) {
	return &db.AddRewardPointsRow{}, nil
}

// replayTransactionCommandRepo mirrors the repository contract for
// CreateTransactionAtomicIdempotent: the same idempotency key returns the
// ORIGINAL record, never a second settlement.
type replayTransactionCommandRepo struct {
	seq     int32
	created map[string]*db.CreateTransactionAtomicRow
}

func (replayTransactionCommandRepo) AuthorizeTransactionAtomic(context.Context, db.AuthorizeTransactionAtomicParams) (*db.AuthorizeTransactionAtomicRow, error) {
	return &db.AuthorizeTransactionAtomicRow{}, nil
}
func (replayTransactionCommandRepo) AuthorizeTransactionAtomicIdempotent(context.Context, db.AuthorizeTransactionAtomicParams) (*repository.AuthorizeTransactionResult, error) {
	return &repository.AuthorizeTransactionResult{Row: &db.AuthorizeTransactionAtomicRow{}}, nil
}

func (r *replayTransactionCommandRepo) CreateTransactionAtomicIdempotent(_ context.Context, params db.CreateTransactionAtomicParams) (*repository.TransactionAtomicResult, error) {
	if params.IdempotencyKey != "" {
		if existing, ok := r.created[params.IdempotencyKey]; ok {
			return &repository.TransactionAtomicResult{Row: existing, Replayed: true}, nil
		}
	}
	res := &db.CreateTransactionAtomicRow{
		TransactionID:   r.seq + 1,
		CardNumber:      params.CardNumber,
		Amount:          params.Amount,
		PaymentMethod:   params.PaymentMethod,
		MerchantID:      params.MerchantID,
		TransactionTime: params.TransactionTime,
		Status:          "success",
	}
	r.seq++
	if params.IdempotencyKey != "" {
		r.created[params.IdempotencyKey] = res
	}
	return &repository.TransactionAtomicResult{Row: res}, nil
}

// The remaining methods of the repository contract are not exercised by the
// replay unit tests. They return well-formed dummy values (never errors) so a
// future service change that starts calling them fails with a clear assertion
// instead of a misleading "not implemented" runtime error.
func (replayTransactionCommandRepo) CreateTransactionAtomic(context.Context, db.CreateTransactionAtomicParams) (*db.CreateTransactionAtomicRow, error) {
	return &db.CreateTransactionAtomicRow{}, nil
}
func (replayTransactionCommandRepo) UpdateSaldoBalanceDelta(context.Context, string, int) error {
	return nil
}
func (replayTransactionCommandRepo) CreateTransaction(context.Context, *requests.CreateTransactionRequest) (*db.CreateTransactionRow, error) {
	return &db.CreateTransactionRow{}, nil
}
func (replayTransactionCommandRepo) UpdateTransaction(context.Context, *requests.UpdateTransactionRequest) (*db.UpdateTransactionRow, error) {
	return &db.UpdateTransactionRow{}, nil
}
func (replayTransactionCommandRepo) UpdateTransactionStatus(context.Context, *requests.UpdateTransactionStatus) (*db.UpdateTransactionStatusRow, error) {
	return &db.UpdateTransactionStatusRow{}, nil
}

func (replayTransactionCommandRepo) UpdateTransactionAtomic(context.Context, db.UpdateTransactionAtomicParams) (*db.UpdateTransactionAtomicRow, error) {
	return &db.UpdateTransactionAtomicRow{}, nil
}

func (replayTransactionCommandRepo) CaptureTransactionAtomic(context.Context, db.CaptureTransactionAtomicParams) (*db.CaptureTransactionAtomicRow, error) {
	return &db.CaptureTransactionAtomicRow{}, nil
}

func (replayTransactionCommandRepo) VoidTransactionAtomic(context.Context, int32) (*db.VoidTransactionAtomicRow, error) {
	return &db.VoidTransactionAtomicRow{}, nil
}

func (replayTransactionCommandRepo) RefundTransactionAtomic(context.Context, db.RefundTransactionAtomicParams) (*db.RefundTransactionAtomicRow, error) {
	return &db.RefundTransactionAtomicRow{}, nil
}
func (replayTransactionCommandRepo) TrashedTransaction(context.Context, int) (*db.TrashTransactionRow, error) {
	return &db.TrashTransactionRow{}, nil
}
func (replayTransactionCommandRepo) RestoreTransaction(context.Context, int) (*db.RestoreTransactionRow, error) {
	return &db.RestoreTransactionRow{}, nil
}
func (replayTransactionCommandRepo) DeleteTransactionPermanent(context.Context, int) (bool, error) {
	return true, nil
}
func (replayTransactionCommandRepo) RestoreAllTransaction(context.Context) (bool, error) {
	return true, nil
}
func (replayTransactionCommandRepo) DeleteAllTransactionPermanent(context.Context) (bool, error) {
	return true, nil
}

type replayTransactionQueryRepo struct{}

func (replayTransactionQueryRepo) FindAllTransactions(context.Context, *requests.FindAllTransactions) ([]*db.GetTransactionsRow, error) {
	return []*db.GetTransactionsRow{}, nil
}
func (replayTransactionQueryRepo) FindByActive(context.Context, *requests.FindAllTransactions) ([]*db.GetActiveTransactionsRow, error) {
	return []*db.GetActiveTransactionsRow{}, nil
}
func (replayTransactionQueryRepo) FindByTrashed(context.Context, *requests.FindAllTransactions) ([]*db.GetTrashedTransactionsRow, error) {
	return []*db.GetTrashedTransactionsRow{}, nil
}
func (replayTransactionQueryRepo) FindAllTransactionByCardNumber(context.Context, *requests.FindAllTransactionCardNumber) ([]*db.GetTransactionsByCardNumberRow, error) {
	return []*db.GetTransactionsByCardNumberRow{}, nil
}
func (replayTransactionQueryRepo) FindById(context.Context, int) (*db.GetTransactionByIDRow, error) {
	return &db.GetTransactionByIDRow{}, nil
}
func (replayTransactionQueryRepo) FindTransactionByMerchantId(context.Context, int) ([]*db.GetTransactionsByMerchantIDRow, error) {
	return []*db.GetTransactionsByMerchantIDRow{}, nil
}

type replayTransactionSaldoRepo struct{}

func (replayTransactionSaldoRepo) FindByCardNumber(context.Context, string) (*db.Saldo, error) {
	return &db.Saldo{}, nil
}
func (replayTransactionSaldoRepo) UpdateSaldoBalance(context.Context, *requests.UpdateSaldoBalance) (*db.UpdateSaldoBalanceRow, error) {
	return &db.UpdateSaldoBalanceRow{}, nil
}

type replayTransactionObservability struct{}

func (replayTransactionObservability) StartTracingAndLogging(ctx context.Context, _ string, _ ...attribute.KeyValue) (context.Context, trace.Span, func(string), string, func(string, ...zap.Field)) {
	ctx, span := trace.NewNoopTracerProvider().Tracer("transaction-replay").Start(ctx, "transaction-replay")
	return ctx, span, func(string) { span.End() }, "success", func(string, ...zap.Field) {}
}
func (replayTransactionObservability) RecordMetrics(context.Context, string, string, time.Time) {}

func newReplayTransactionService(command *replayTransactionCommandRepo) service.TransactionCommandService {
	return service.NewTransactionCommandService(&service.TransactionCommandServiceDeps{
		MerchantRepository:           replayTransactionMerchantRepo{},
		CardRepository:               replayTransactionCardRepo{},
		SaldoRepository:              replayTransactionSaldoRepo{},
		TransactionQueryRepository:   replayTransactionQueryRepo{},
		TransactionCommandRepository: command,
		Logger:                       &logger.Logger{Log: zap.NewNop()},
		Observability:                replayTransactionObservability{},
	})
}

func TestCreateTransactionReplayReturnsOriginal(t *testing.T) {
	command := &replayTransactionCommandRepo{created: map[string]*db.CreateTransactionAtomicRow{}}
	svc := newReplayTransactionService(command)

	req := &requests.CreateTransactionRequest{
		CardNumber:      "4111111111111111",
		Amount:          100000,
		PaymentMethod:   "gopay",
		TransactionTime: time.Now(),
		IdempotencyKey:  "transaction:create:merchant:1:op-1",
	}

	first, err := svc.Create(context.Background(), "api-key", req)
	require.NoError(t, err)
	require.NotNil(t, first)
	require.Equal(t, int32(1), first.TransactionID)
	require.Equal(t, "success", first.Status)

	// Replay with the same key: same transaction id, no second record.
	second, err := svc.Create(context.Background(), "api-key", req)
	require.NoError(t, err)
	require.NotNil(t, second)
	require.Equal(t, first.TransactionID, second.TransactionID, "replay must return the original transaction id")
	require.Equal(t, first.Amount, second.Amount)
	require.Len(t, command.created, 1, "the key must be persisted exactly once")
}

func TestCreateTransactionWithoutKeyCreatesTwice(t *testing.T) {
	command := &replayTransactionCommandRepo{created: map[string]*db.CreateTransactionAtomicRow{}}
	svc := newReplayTransactionService(command)

	req := &requests.CreateTransactionRequest{
		CardNumber:      "4111111111111111",
		Amount:          100000,
		PaymentMethod:   "gopay",
		TransactionTime: time.Now(),
	}

	first, err := svc.Create(context.Background(), "api-key", req)
	require.NoError(t, err)
	second, err := svc.Create(context.Background(), "api-key", req)
	require.NoError(t, err)

	require.NotEqual(t, first.TransactionID, second.TransactionID, "without a key every request creates a new transaction")
}
