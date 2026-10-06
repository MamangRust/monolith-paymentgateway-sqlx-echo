package withdrawserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedCreateWithdraw is used when failed to create withdraw
	ErrFailedCreateWithdraw = errors.ErrInternal.WithMessage("Failed to create withdraw")

	// ErrFailedUpdateWithdraw is used when failed to update withdraw
	ErrFailedUpdateWithdraw = errors.ErrInternal.WithMessage("Failed to update withdraw")

	// ErrFailedTrashedWithdraw is used when failed to trash withdraw
	ErrFailedTrashedWithdraw = errors.ErrInternal.WithMessage("Failed to move withdraw to trash")

	// ErrFailedRestoreWithdraw is used when failed to restore withdraw
	ErrFailedRestoreWithdraw = errors.ErrInternal.WithMessage("Failed to restore withdraw")

	// ErrFailedDeleteWithdrawPermanent is used when failed to permanently delete withdraw
	ErrFailedDeleteWithdrawPermanent = errors.ErrInternal.WithMessage("Failed to delete withdraw permanently")

	// ErrFailedRestoreAllWithdraw is used when failed to restore all withdraws
	ErrFailedRestoreAllWithdraw = errors.ErrInternal.WithMessage("Failed to restore all withdraws")

	// ErrFailedDeleteAllWithdrawPermanent is used when failed to permanently delete all withdraws
	ErrFailedDeleteAllWithdrawPermanent = errors.ErrInternal.WithMessage("Failed to delete all withdraws permanently")
)
