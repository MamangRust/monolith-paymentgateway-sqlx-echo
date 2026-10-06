package merchantserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedCreateMerchant indicates failure when creating a new merchant.
	ErrFailedCreateMerchant = errors.ErrInternal.WithMessage("Failed to create merchant")

	// ErrFailedUpdateMerchant indicates failure when updating a merchant.
	ErrFailedUpdateMerchant = errors.ErrInternal.WithMessage("Failed to update merchant")

	// ErrFailedTrashMerchant indicates failure when soft-deleting (trashing) a merchant.
	ErrFailedTrashMerchant = errors.ErrInternal.WithMessage("Failed to move merchant to trash")

	// ErrFailedRestoreMerchant indicates failure when restoring a trashed merchant.
	ErrFailedRestoreMerchant = errors.ErrInternal.WithMessage("Failed to restore merchant")

	// ErrFailedDeleteMerchant indicates failure when permanently deleting a merchant.
	ErrFailedDeleteMerchant = errors.ErrInternal.WithMessage("Failed to delete merchant permanently")

	// ErrFailedRestoreAllMerchants indicates failure when restoring all trashed merchants.
	ErrFailedRestoreAllMerchants = errors.ErrInternal.WithMessage("Failed to restore all merchants")

	// ErrFailedDeleteAllMerchants indicates failure when permanently deleting all trashed merchants.
	ErrFailedDeleteAllMerchants = errors.ErrInternal.WithMessage("Failed to delete all merchants permanently")
)
