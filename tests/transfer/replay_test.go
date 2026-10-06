package transfer_test

import (
	"context"
	"testing"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-transfer/repository"
	"github.com/MamangRust/monolith-payment-gateway-transfer/service"
	"github.com/stretchr/testify/require"
)

// replayTransferCommandRepo mirrors the repository contract for
// CreateTransferAtomic: the same idempotency key returns the ORIGINAL result
// (Replayed=true), never a second settlement.
type replayTransferCommandRepo struct {
	seq     int32
	created map[string]*repository.TransferAtomicResult
}

func (r *replayTransferCommandRepo) CreateTransferAtomic(_ context.Context, req *requests.CreateTransferRequest) (*repository.TransferAtomicResult, error) {
	if req.IdempotencyKey != "" {
		if existing, ok := r.created[req.IdempotencyKey]; ok {
			// Signal the replay exactly like the real repository: the original
			// record is returned and side-effects must be skipped.
			existing.Replayed = true
			return existing, nil
		}
	}
	res := &repository.TransferAtomicResult{
		Row: &db.CreateTransferAtomicRow{
			TransferID:     r.seq + 1,
			TransferFrom:   req.TransferFrom,
			TransferTo:     req.TransferTo,
			TransferAmount: int32(req.TransferAmount),
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
func (replayTransferCommandRepo) CreateTransfer(context.Context, *requests.CreateTransferRequest) (*db.CreateTransferRow, error) {
	return &db.CreateTransferRow{}, nil
}
func (replayTransferCommandRepo) UpdateTransferStatus(context.Context, *requests.UpdateTransferStatus) (*db.UpdateTransferStatusRow, error) {
	return &db.UpdateTransferStatusRow{}, nil
}
func (replayTransferCommandRepo) UpdateTransferAtomic(context.Context, db.UpdateTransferAtomicParams) (*db.UpdateTransferAtomicRow, error) {
	return &db.UpdateTransferAtomicRow{}, nil
}
func (replayTransferCommandRepo) UpdateTransferSettlementDelta(context.Context, string, string, int, int) error {
	return nil
}
func (replayTransferCommandRepo) UpdateTransfer(context.Context, *requests.UpdateTransferRequest) (*db.UpdateTransferRow, error) {
	return &db.UpdateTransferRow{}, nil
}
func (replayTransferCommandRepo) UpdateTransferAmount(context.Context, *requests.UpdateTransferAmountRequest) (*db.UpdateTransferAmountRow, error) {
	return &db.UpdateTransferAmountRow{}, nil
}
func (replayTransferCommandRepo) TrashedTransfer(context.Context, int) (*db.Transfer, error) {
	return &db.Transfer{}, nil
}
func (replayTransferCommandRepo) RestoreTransfer(context.Context, int) (*db.Transfer, error) {
	return &db.Transfer{}, nil
}
func (replayTransferCommandRepo) DeleteTransferPermanent(context.Context, int) (bool, error) {
	return true, nil
}
func (replayTransferCommandRepo) RestoreAllTransfer(context.Context) (bool, error) {
	return true, nil
}
func (replayTransferCommandRepo) DeleteAllTransferPermanent(context.Context) (bool, error) {
	return true, nil
}

func newReplayTransferService(command repository.TransferCommandRepository) service.TransferCommandService {
	return newP0TransferService(&p0TransferSaldoRepo{}, command)
}

func TestCreateTransferReplayReturnsOriginalSettlement(t *testing.T) {
	command := &replayTransferCommandRepo{created: map[string]*repository.TransferAtomicResult{}}
	svc := newReplayTransferService(command)

	req := &requests.CreateTransferRequest{
		TransferFrom:   "sender",
		TransferTo:     "receiver",
		TransferAmount: 50000,
		IdempotencyKey: "transfer:create:user:7:op-1",
	}

	first, err := svc.CreateTransaction(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, first)
	require.Equal(t, int32(1), first.TransferID)
	require.Equal(t, "success", first.Status)

	// Replay with the same key: same settlement id, no second record.
	second, err := svc.CreateTransaction(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, second)
	require.Equal(t, first.TransferID, second.TransferID, "replay must return the original settlement id")
	require.Equal(t, first.TransferAmount, second.TransferAmount)
	require.Len(t, command.created, 1, "the key must be persisted exactly once")
}

func TestCreateTransferWithoutKeySettlesTwice(t *testing.T) {
	command := &replayTransferCommandRepo{created: map[string]*repository.TransferAtomicResult{}}
	svc := newReplayTransferService(command)

	req := &requests.CreateTransferRequest{
		TransferFrom:   "sender",
		TransferTo:     "receiver",
		TransferAmount: 50000,
	}

	first, err := svc.CreateTransaction(context.Background(), req)
	require.NoError(t, err)
	second, err := svc.CreateTransaction(context.Background(), req)
	require.NoError(t, err)

	require.NotEqual(t, first.TransferID, second.TransferID, "without a key every request creates a new settlement")
}
