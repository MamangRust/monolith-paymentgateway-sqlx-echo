package merchantrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthlyAmountMerchantFailed indicates failure fetching monthly amount for a merchant.
	ErrGetMonthlyAmountMerchantFailed = errors.ErrInternal.WithMessage("failed to get monthly amount of merchant")

	// ErrGetYearlyAmountMerchantFailed indicates failure fetching yearly amount for a merchant.
	ErrGetYearlyAmountMerchantFailed = errors.ErrInternal.WithMessage("failed to get yearly amount of merchant")

	// ErrGetMonthlyAmountByMerchantsFailed indicates failure fetching monthly amount for all merchants.
	ErrGetMonthlyAmountByMerchantsFailed = errors.ErrInternal.WithMessage("failed to get monthly amount by merchants")

	// ErrGetYearlyAmountByMerchantsFailed indicates failure fetching yearly amount for all merchants.
	ErrGetYearlyAmountByMerchantsFailed = errors.ErrInternal.WithMessage("failed to get yearly amount by merchants")

	// ErrGetMonthlyAmountByApikeyFailed indicates failure fetching monthly amount by API key.
	ErrGetMonthlyAmountByApikeyFailed = errors.ErrInternal.WithMessage("failed to get monthly amount by API key")

	// ErrGetYearlyAmountByApikeyFailed indicates failure fetching yearly amount by API key.
	ErrGetYearlyAmountByApikeyFailed = errors.ErrInternal.WithMessage("failed to get yearly amount by API key")
)
