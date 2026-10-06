package merchantrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFindAllMerchantsFailed is returned when fetching all merchants fails.
	ErrFindAllMerchantsFailed = errors.ErrInternal.WithMessage("Failed to find all merchants")

	// ErrFindActiveMerchantsFailed is returned when fetching active merchants fails.
	ErrFindActiveMerchantsFailed = errors.ErrInternal.WithMessage("Failed to find active merchants")

	// ErrFindTrashedMerchantsFailed is returned when fetching trashed merchants fails.
	ErrFindTrashedMerchantsFailed = errors.ErrInternal.WithMessage("Failed to find trashed merchants")

	// ErrFindMerchantByIdFailed is returned when a merchant cannot be found by ID.
	ErrFindMerchantByIdFailed = errors.ErrInternal.WithMessage("Failed to find merchant by ID")

	// ErrFindMerchantByApiKeyFailed is returned when a merchant cannot be found by API key.
	ErrFindMerchantByApiKeyFailed = errors.ErrInternal.WithMessage("Failed to find merchant by API key")

	// ErrFindMerchantByNameFailed is returned when a merchant cannot be found by name.
	ErrFindMerchantByNameFailed = errors.ErrInternal.WithMessage("Failed to find merchant by name")

	// ErrFindMerchantByUserIdFailed is returned when a merchant cannot be found by user ID.
	ErrFindMerchantByUserIdFailed = errors.ErrInternal.WithMessage("Failed to find merchant by user ID")

	// ErrMerchantNotFound is returned when a merchant is not found.
	ErrMerchantNotFound = errors.ErrNotFound.WithMessage("Merchant not found")
)
