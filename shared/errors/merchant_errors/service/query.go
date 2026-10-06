package merchantserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrMerchantNotFoundRes is returned when the merchant is not found.
	ErrMerchantNotFoundRes = errors.ErrNotFound.WithMessage("Merchant not found")

	// ErrFailedFindAllMerchants indicates failure in fetching all merchants.
	ErrFailedFindAllMerchants = errors.ErrInternal.WithMessage("Failed to fetch merchants")

	// ErrFailedFindActiveMerchants indicates failure in fetching active merchants.
	ErrFailedFindActiveMerchants = errors.ErrInternal.WithMessage("Failed to fetch active merchants")

	// ErrFailedFindTrashedMerchants indicates failure in fetching trashed merchants.
	ErrFailedFindTrashedMerchants = errors.ErrInternal.WithMessage("Failed to fetch trashed merchants")

	// ErrFailedFindMerchantById is returned when a merchant cannot be found by ID.
	ErrFailedFindMerchantById = errors.ErrInternal.WithMessage("Failed to find merchant by ID")

	// ErrFailedFindByApiKey is returned when a merchant cannot be found by API key.
	ErrFailedFindByApiKey = errors.ErrInternal.WithMessage("Failed to find merchant by API key")

	// ErrFailedFindByMerchantUserId is returned when a merchant cannot be found by user ID.
	ErrFailedFindByMerchantUserId = errors.ErrInternal.WithMessage("Failed to find merchant by user ID")
)
