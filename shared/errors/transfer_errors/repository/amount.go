package transferrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthlyTransferAmountsFailed indicates a failure when retrieving the total amount of monthly transfers.
	ErrGetMonthlyTransferAmountsFailed = errors.ErrInternal.WithMessage("Failed to get monthly transfer amounts")

	// ErrGetYearlyTransferAmountsFailed indicates a failure when retrieving the total amount of yearly transfers.
	ErrGetYearlyTransferAmountsFailed = errors.ErrInternal.WithMessage("Failed to get yearly transfer amounts")

	// ErrGetMonthlyTransferAmountsBySenderCardFailed indicates a failure when retrieving monthly transfer amounts filtered by sender card number.
	ErrGetMonthlyTransferAmountsBySenderCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly transfer amounts by sender card number")

	// ErrGetYearlyTransferAmountsBySenderCardFailed indicates a failure when retrieving yearly transfer amounts filtered by sender card number.
	ErrGetYearlyTransferAmountsBySenderCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly transfer amounts by sender card number")

	// ErrGetMonthlyTransferAmountsByReceiverCardFailed indicates a failure when retrieving monthly transfer amounts filtered by receiver card number.
	ErrGetMonthlyTransferAmountsByReceiverCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly transfer amounts by receiver card number")

	// ErrGetYearlyTransferAmountsByReceiverCardFailed indicates a failure when retrieving yearly transfer amounts filtered by receiver card number.
	ErrGetYearlyTransferAmountsByReceiverCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly transfer amounts by receiver card number")
)
