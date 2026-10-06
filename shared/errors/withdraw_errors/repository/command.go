package withdrawrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrCreateWithdrawFailed indicates a failure when creating a new withdraw record.
	ErrCreateWithdrawFailed = errors.ErrInternal.WithMessage("Failed to create withdraw")

	// ErrUpdateWithdrawFailed indicates a failure when updating a withdraw record.
	ErrUpdateWithdrawFailed = errors.ErrInternal.WithMessage("Failed to update withdraw")

	// ErrUpdateWithdrawStatusFailed indicates a failure when updating the status of a withdraw record.
	ErrUpdateWithdrawStatusFailed = errors.ErrInternal.WithMessage("Failed to update withdraw status")

	// ErrTrashedWithdrawFailed indicates a failure when soft-deleting (trashing) a withdraw record.
	ErrTrashedWithdrawFailed = errors.ErrInternal.WithMessage("Failed to move withdraw to trash")

	// ErrRestoreWithdrawFailed indicates a failure when restoring a previously trashed withdraw record.
	ErrRestoreWithdrawFailed = errors.ErrInternal.WithMessage("Failed to restore withdraw from trash")

	// ErrDeleteWithdrawPermanentFailed indicates a failure when permanently deleting a withdraw record.
	ErrDeleteWithdrawPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete withdraw")

	// ErrRestoreAllWithdrawsFailed indicates a failure when restoring all trashed withdraw records.
	ErrRestoreAllWithdrawsFailed = errors.ErrInternal.WithMessage("Failed to restore all withdraws")

	// ErrDeleteAllWithdrawsPermanentFailed indicates a failure when permanently deleting all withdraw records.
	ErrDeleteAllWithdrawsPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete all withdraws")
)
