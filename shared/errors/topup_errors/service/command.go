package topupserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedCreateTopup indicates failure in creating a new top-up record.
	ErrFailedCreateTopup = errors.ErrInternal.WithMessage("Failed to create topup")

	// ErrFailedUpdateTopup indicates failure in updating an existing top-up record.
	ErrFailedUpdateTopup = errors.ErrInternal.WithMessage("Failed to update topup")

	// ErrFailedTrashTopup indicates failure in soft-deleting (trashing) a top-up.
	ErrFailedTrashTopup = errors.ErrInternal.WithMessage("Failed to move topup to trash")

	// ErrFailedRestoreTopup indicates failure in restoring a previously trashed top-up.
	ErrFailedRestoreTopup = errors.ErrInternal.WithMessage("Failed to restore topup")

	// ErrFailedDeleteTopup indicates failure in permanently deleting a top-up.
	ErrFailedDeleteTopup = errors.ErrInternal.WithMessage("Failed to delete topup permanently")

	// ErrFailedRestoreAllTopups indicates failure in restoring all trashed top-up records.
	ErrFailedRestoreAllTopups = errors.ErrInternal.WithMessage("Failed to restore all topups")

	// ErrFailedDeleteAllTopups indicates failure in permanently deleting all trashed top-up records.
	ErrFailedDeleteAllTopups = errors.ErrInternal.WithMessage("Failed to delete all topups permanently")
)
