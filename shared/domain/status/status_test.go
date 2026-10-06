package status_test

import (
	"errors"
	"testing"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/status"
	sharederrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	"github.com/stretchr/testify/require"
)

func TestValidateTransition_TopupLifecycle(t *testing.T) {
	valid := [][2]string{
		{status.Pending, status.Success}, // settlement completes
		{status.Pending, status.Failed},  // settlement fails / compensation
		{status.Failed, status.Pending},  // retry reopens
		{status.Failed, status.Success},  // retry completes
		{status.Success, status.Success}, // same-state re-assertion (update flow)
		{status.Success, status.Failed},  // system compensation only
	}
	for _, tr := range valid {
		require.NoError(t, status.ValidateTransition(status.DomainTopup, tr[0], tr[1]),
			"topup %s -> %s should be valid", tr[0], tr[1])
	}

	invalid := [][2]string{
		{status.Success, status.Pending}, // a settled topup can never be reverted
		{status.Success, status.Authorized},
		{status.Authorized, status.Success},
		{status.Voided, status.Success},
	}
	for _, tr := range invalid {
		err := status.ValidateTransition(status.DomainTopup, tr[0], tr[1])
		require.Error(t, err, "topup %s -> %s should be invalid", tr[0], tr[1])
		assertBadRequest(t, err)
	}
}

func TestValidateTransition_TransferAndWithdrawMirrorTopup(t *testing.T) {
	for _, domain := range []status.Domain{status.DomainTransfer, status.DomainWithdraw} {
		require.NoError(t, status.ValidateTransition(domain, status.Pending, status.Success))
		require.NoError(t, status.ValidateTransition(domain, status.Failed, status.Success))
		require.Error(t, status.ValidateTransition(domain, status.Success, status.Pending))
		require.Error(t, status.ValidateTransition(domain, status.Pending, status.Authorized))
	}
}

func TestValidateTransition_TransactionLifecycle(t *testing.T) {
	// Happy path: pending -> authorized -> captured -> refunded.
	happyPath := [][2]string{
		{status.Pending, status.Authorized},
		{status.Authorized, status.Captured},
		{status.Captured, status.Refunded},
	}
	for _, tr := range happyPath {
		require.NoError(t, status.ValidateTransition(status.DomainTransaction, tr[0], tr[1]),
			"transaction %s -> %s should be valid", tr[0], tr[1])
	}

	// Void is valid from pending and authorized only.
	require.NoError(t, status.ValidateTransition(status.DomainTransaction, status.Pending, status.Voided))
	require.NoError(t, status.ValidateTransition(status.DomainTransaction, status.Authorized, status.Voided))

	// Update flow re-confirms success; failed records can be reopened.
	require.NoError(t, status.ValidateTransition(status.DomainTransaction, status.Success, status.Success))
	require.NoError(t, status.ValidateTransition(status.DomainTransaction, status.Failed, status.Success))

	invalid := [][2]string{
		{status.Captured, status.Authorized}, // cannot un-capture
		{status.Voided, status.Captured},     // voided is terminal
		{status.Refunded, status.Captured},   // refunded is terminal
		{status.Success, status.Captured},    // success -> capture is not a lifecycle move
		{status.Authorized, status.Refunded}, // must capture before refunding
		{status.Pending, status.Refunded},    // must be captured first
		// Same-state lifecycle ops must be rejected: a second Capture/Void/
		// Refund on an already-terminal record would double-apply balances.
		{status.Captured, status.Captured},
		{status.Voided, status.Voided},
		{status.Refunded, status.Refunded},
	}
	for _, tr := range invalid {
		err := status.ValidateTransition(status.DomainTransaction, tr[0], tr[1])
		require.Error(t, err, "transaction %s -> %s should be invalid", tr[0], tr[1])
		assertBadRequest(t, err)
	}
}

func TestValidateTransition_CardBillingReference(t *testing.T) {
	require.NoError(t, status.ValidateTransition(status.DomainCardBilling, status.Pending, status.Unpaid))
	require.NoError(t, status.ValidateTransition(status.DomainCardBilling, status.Unpaid, status.Paid))
	require.NoError(t, status.ValidateTransition(status.DomainCardBilling, status.Paid, status.Completed))
	require.Error(t, status.ValidateTransition(status.DomainCardBilling, status.Completed, status.Pending))
}

func TestValidateTransition_RequiresTargetStatus(t *testing.T) {
	err := status.ValidateTransition(status.DomainTopup, status.Pending, "")
	require.Error(t, err)
	assertBadRequest(t, err)
	require.Contains(t, err.Error(), "status is required")
}

func TestValidateTransition_UnknownDomain(t *testing.T) {
	err := status.ValidateTransition("unknown_domain", status.Pending, status.Success)
	require.Error(t, err)
}

func TestCanTransition_SameStateOnlyWhereExplicit(t *testing.T) {
	// Update re-confirmation of an already-success record is explicitly allowed.
	require.True(t, status.CanTransition(status.DomainTopup, status.Success, status.Success))
	require.True(t, status.CanTransition(status.DomainTransaction, status.Success, status.Success))

	// Lifecycle operations that mutate balances must NOT be same-state safe:
	// re-capture/re-void/re-refund would double-apply balances.
	require.False(t, status.CanTransition(status.DomainTransaction, status.Captured, status.Captured))
	require.False(t, status.CanTransition(status.DomainTransaction, status.Voided, status.Voided))
	require.False(t, status.CanTransition(status.DomainTransaction, status.Refunded, status.Refunded))
	require.False(t, status.CanTransition(status.DomainTopup, status.Success, status.Pending))
}

func assertBadRequest(t *testing.T, err error) {
	t.Helper()
	var appErr *sharederrors.AppError
	require.True(t, errors.As(err, &appErr), "expected *AppError, got %T", err)
	require.Equal(t, sharederrors.ErrorTypeBadRequest, appErr.Type)
	require.Equal(t, 400, appErr.Code)
}
