package merchantrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrCreateMerchantFailed indicates failure when creating a merchant.
	ErrCreateMerchantFailed = errors.ErrInternal.WithMessage("Failed to create merchant")

	// ErrUpdateMerchantFailed indicates failure when updating a merchant.
	ErrUpdateMerchantFailed = errors.ErrInternal.WithMessage("Failed to update merchant")

	// ErrUpdateMerchantStatusFailed indicates failure when updating merchant status.
	ErrUpdateMerchantStatusFailed = errors.ErrInternal.WithMessage("Failed to update merchant status")

	// ErrTrashedMerchantFailed indicates failure when soft-deleting a merchant.
	ErrTrashedMerchantFailed = errors.ErrInternal.WithMessage("Failed to move merchant to trash")

	// ErrRestoreMerchantFailed indicates failure when restoring a trashed merchant.
	ErrRestoreMerchantFailed = errors.ErrInternal.WithMessage("Failed to restore merchant from trash")

	// ErrDeleteMerchantPermanentFailed indicates failure when permanently deleting a merchant.
	ErrDeleteMerchantPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete merchant")

	// ErrRestoreAllMerchantFailed indicates failure when restoring all trashed merchants.
	ErrRestoreAllMerchantFailed = errors.ErrInternal.WithMessage("Failed to restore all merchants")

	// ErrDeleteAllMerchantPermanentFailed indicates failure when permanently deleting all merchants.
	ErrDeleteAllMerchantPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete all merchants")
)
