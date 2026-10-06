package withdrawserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedFindAllWithdraws is used when failed to fetch all withdraws
	ErrFailedFindAllWithdraws = errors.ErrInternal.WithMessage("Failed to fetch all withdraws")

	// ErrWithdrawNotFound is used when withdraw is not found
	ErrWithdrawNotFound = errors.ErrNotFound.WithMessage("Withdraw not found")

	// ErrFailedFindAllWithdrawsByCard is used when failed to fetch all withdraws by card number
	ErrFailedFindAllWithdrawsByCard = errors.ErrInternal.WithMessage("Failed to fetch all withdraws by card number")

	// ErrFailedFindActiveWithdraws is used when failed to fetch active withdraws
	ErrFailedFindActiveWithdraws = errors.ErrInternal.WithMessage("Failed to fetch active withdraws")

	// ErrFailedFindTrashedWithdraws is used when failed to fetch trashed withdraws
	ErrFailedFindTrashedWithdraws = errors.ErrInternal.WithMessage("Failed to fetch trashed withdraws")
)
