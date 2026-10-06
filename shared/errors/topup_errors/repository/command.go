package topuprepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrCreateTopupFailed indicates failure in creating a new top-up record.
	ErrCreateTopupFailed = errors.ErrInternal.WithMessage("Failed to create topup")

	// ErrUpdateTopupFailed indicates failure in updating an existing top-up record.
	ErrUpdateTopupFailed = errors.ErrInternal.WithMessage("Failed to update topup")

	// ErrUpdateTopupAmountFailed indicates failure in updating only the top-up amount.
	ErrUpdateTopupAmountFailed = errors.ErrInternal.WithMessage("Failed to update topup amount")

	// ErrUpdateTopupStatusFailed indicates failure in updating the top-up status (e.g., success/failed).
	ErrUpdateTopupStatusFailed = errors.ErrInternal.WithMessage("Failed to update topup status")

	// ErrTrashedTopupFailed indicates failure in soft-deleting (trashing) a top-up.
	ErrTrashedTopupFailed = errors.ErrInternal.WithMessage("Failed to move topup to trash")

	// ErrRestoreTopupFailed indicates failure in restoring a previously trashed top-up.
	ErrRestoreTopupFailed = errors.ErrInternal.WithMessage("Failed to restore topup from trash")

	// ErrDeleteTopupPermanentFailed indicates failure in permanently deleting a top-up.
	ErrDeleteTopupPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete topup")

	// ErrRestoreAllTopupFailed indicates failure in restoring all trashed top-ups.
	ErrRestoreAllTopupFailed = errors.ErrInternal.WithMessage("Failed to restore all topups")

	// ErrDeleteAllTopupPermanentFailed indicates failure in permanently deleting all trashed top-ups.
	ErrDeleteAllTopupPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete all topups")
)
