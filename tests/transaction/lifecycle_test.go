package transaction_test

import (
	"context"
	"testing"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	statemachine "github.com/MamangRust/monolith-payment-gateway-shared/domain/status"
	"github.com/MamangRust/monolith-payment-gateway-transaction/repository"
	"github.com/MamangRust/monolith-payment-gateway-transaction/service"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// lifecycleQueryRepo returns a transaction in the given state so the lifecycle
// state machine is exercised at the service boundary.
type lifecycleQueryRepo struct {
	currentStatus string
}

func (r *lifecycleQueryRepo) FindById(_ context.Context, id int) (*db.GetTransactionByIDRow, error) {
	return &db.GetTransactionByIDRow{TransactionID: int32(id), Status: r.currentStatus}, nil
}

// The remaining repository methods are not exercised by the lifecycle unit
// tests. They return well-formed dummy values (never errors) so a future
// service change that starts calling them fails with a clear assertion
// instead of a misleading "not implemented" runtime error.
func (lifecycleQueryRepo) FindAllTransactions(context.Context, *requests.FindAllTransactions) ([]*db.GetTransactionsRow, error) {
	return []*db.GetTransactionsRow{}, nil
}
func (lifecycleQueryRepo) FindByActive(context.Context, *requests.FindAllTransactions) ([]*db.GetActiveTransactionsRow, error) {
	return []*db.GetActiveTransactionsRow{}, nil
}
func (lifecycleQueryRepo) FindByTrashed(context.Context, *requests.FindAllTransactions) ([]*db.GetTrashedTransactionsRow, error) {
	return []*db.GetTrashedTransactionsRow{}, nil
}
func (lifecycleQueryRepo) FindAllTransactionByCardNumber(context.Context, *requests.FindAllTransactionCardNumber) ([]*db.GetTransactionsByCardNumberRow, error) {
	return []*db.GetTransactionsByCardNumberRow{}, nil
}
func (lifecycleQueryRepo) FindTransactionByMerchantId(context.Context, int) ([]*db.GetTransactionsByMerchantIDRow, error) {
	return []*db.GetTransactionsByMerchantIDRow{}, nil
}

// lifecycleCommandRepo counts balance deltas so tests can assert that a
// rejected lifecycle operation never moves money.
type lifecycleCommandRepo struct {
	deltaCalls int
}

func (r *lifecycleCommandRepo) UpdateSaldoBalanceDelta(context.Context, string, int) error {
	r.deltaCalls++
	return nil
}

func (lifecycleCommandRepo) CreateTransactionAtomic(context.Context, db.CreateTransactionAtomicParams) (*db.CreateTransactionAtomicRow, error) {
	return &db.CreateTransactionAtomicRow{}, nil
}
func (lifecycleCommandRepo) CreateTransactionAtomicIdempotent(context.Context, db.CreateTransactionAtomicParams) (*repository.TransactionAtomicResult, error) {
	return &repository.TransactionAtomicResult{Row: &db.CreateTransactionAtomicRow{}}, nil
}
func (lifecycleCommandRepo) CreateTransaction(context.Context, *requests.CreateTransactionRequest) (*db.CreateTransactionRow, error) {
	return &db.CreateTransactionRow{}, nil
}
func (lifecycleCommandRepo) UpdateTransaction(context.Context, *requests.UpdateTransactionRequest) (*db.UpdateTransactionRow, error) {
	return &db.UpdateTransactionRow{}, nil
}
func (lifecycleCommandRepo) UpdateTransactionStatus(context.Context, *requests.UpdateTransactionStatus) (*db.UpdateTransactionStatusRow, error) {
	return &db.UpdateTransactionStatusRow{}, nil
}

func (lifecycleCommandRepo) AuthorizeTransactionAtomic(context.Context, db.AuthorizeTransactionAtomicParams) (*db.AuthorizeTransactionAtomicRow, error) {
	return &db.AuthorizeTransactionAtomicRow{}, nil
}
func (lifecycleCommandRepo) AuthorizeTransactionAtomicIdempotent(context.Context, db.AuthorizeTransactionAtomicParams) (*repository.AuthorizeTransactionResult, error) {
	return &repository.AuthorizeTransactionResult{Row: &db.AuthorizeTransactionAtomicRow{}}, nil
}

func (lifecycleCommandRepo) UpdateTransactionAtomic(context.Context, db.UpdateTransactionAtomicParams) (*db.UpdateTransactionAtomicRow, error) {
	return &db.UpdateTransactionAtomicRow{}, nil
}

func (lifecycleCommandRepo) CaptureTransactionAtomic(context.Context, db.CaptureTransactionAtomicParams) (*db.CaptureTransactionAtomicRow, error) {
	return &db.CaptureTransactionAtomicRow{}, nil
}

func (lifecycleCommandRepo) VoidTransactionAtomic(context.Context, int32) (*db.VoidTransactionAtomicRow, error) {
	return &db.VoidTransactionAtomicRow{}, nil
}

func (lifecycleCommandRepo) RefundTransactionAtomic(context.Context, db.RefundTransactionAtomicParams) (*db.RefundTransactionAtomicRow, error) {
	return &db.RefundTransactionAtomicRow{}, nil
}
func (lifecycleCommandRepo) TrashedTransaction(context.Context, int) (*db.TrashTransactionRow, error) {
	return &db.TrashTransactionRow{}, nil
}
func (lifecycleCommandRepo) RestoreTransaction(context.Context, int) (*db.RestoreTransactionRow, error) {
	return &db.RestoreTransactionRow{}, nil
}
func (lifecycleCommandRepo) DeleteTransactionPermanent(context.Context, int) (bool, error) {
	return true, nil
}
func (lifecycleCommandRepo) RestoreAllTransaction(context.Context) (bool, error) {
	return true, nil
}
func (lifecycleCommandRepo) DeleteAllTransactionPermanent(context.Context) (bool, error) {
	return true, nil
}

func newLifecycleService(command *lifecycleCommandRepo, query *lifecycleQueryRepo) service.TransactionCommandService {
	return service.NewTransactionCommandService(&service.TransactionCommandServiceDeps{
		MerchantRepository:           replayTransactionMerchantRepo{},
		CardRepository:               replayTransactionCardRepo{},
		SaldoRepository:              replayTransactionSaldoRepo{},
		TransactionQueryRepository:   query,
		TransactionCommandRepository: command,
		Logger:                       &logger.Logger{Log: zap.NewNop()},
		Observability:                replayTransactionObservability{},
	})
}

func TestLifecycleOpsRejectDoubleApplication(t *testing.T) {
	cases := []struct {
		name    string
		current string
		call    func(service.TransactionCommandService) (*db.UpdateTransactionStatusRow, error)
	}{
		{"double-capture", statemachine.Captured, func(s service.TransactionCommandService) (*db.UpdateTransactionStatusRow, error) {
			return s.CaptureTransaction(context.Background(), "api-key", 7)
		}},
		{"double-void", statemachine.Voided, func(s service.TransactionCommandService) (*db.UpdateTransactionStatusRow, error) {
			return s.VoidTransaction(context.Background(), "api-key", 7)
		}},
		{"double-refund", statemachine.Refunded, func(s service.TransactionCommandService) (*db.UpdateTransactionStatusRow, error) {
			return s.RefundTransaction(context.Background(), "api-key", 7)
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			command := &lifecycleCommandRepo{}
			svc := newLifecycleService(command, &lifecycleQueryRepo{currentStatus: tc.current})

			_, err := tc.call(svc)
			require.Error(t, err, "double %s must be rejected", tc.name)
			require.Contains(t, err.Error(), "invalid status transition", "expected a 400 state-machine error")
			require.Zero(t, command.deltaCalls, "double %s must not move any balance", tc.name)
		})
	}
}

func TestLifecycleOpsRejectWrongSourceState(t *testing.T) {
	// A refund directly from "authorized" (without capture) is invalid and must
	// not touch balances.
	command := &lifecycleCommandRepo{}
	svc := newLifecycleService(command, &lifecycleQueryRepo{currentStatus: statemachine.Authorized})

	_, err := svc.RefundTransaction(context.Background(), "api-key", 7)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid status transition")
	require.Zero(t, command.deltaCalls)
}
