package withdrawrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFindAllWithdrawsFailed is used when the system fails to find all withdraws
	ErrFindAllWithdrawsFailed = errors.ErrInternal.WithMessage("Failed to find all withdraws")

	// ErrFindActiveWithdrawsFailed is used when the system fails to find active withdraws
	ErrFindActiveWithdrawsFailed = errors.ErrInternal.WithMessage("Failed to find active withdraws")

	// ErrFindTrashedWithdrawsFailed is used when the system fails to find trashed withdraws
	ErrFindTrashedWithdrawsFailed = errors.ErrInternal.WithMessage("Failed to find trashed withdraws")

	// ErrFindWithdrawsByCardNumberFailed is used when the system fails to find withdraws by card number
	ErrFindWithdrawsByCardNumberFailed = errors.ErrInternal.WithMessage("Failed to find withdraws by card number")

	// ErrFindWithdrawByIdFailed is used when the system fails to find a withdraw by ID
	ErrFindWithdrawByIdFailed = errors.ErrInternal.WithMessage("Failed to find withdraw by ID")

	// ErrWithdrawNotFound is returned when a withdraw is not found.
	ErrWithdrawNotFound = errors.ErrNotFound.WithMessage("Withdraw not found")
)
