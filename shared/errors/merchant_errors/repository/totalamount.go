package merchantrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthlyTotalAmountMerchantFailed indicates failure fetching monthly total amount for a merchant.
	ErrGetMonthlyTotalAmountMerchantFailed = errors.ErrInternal.WithMessage("failed to get monthly total amount of merchant")

	// ErrGetYearlyTotalAmountMerchantFailed indicates failure fetching yearly total amount for a merchant.
	ErrGetYearlyTotalAmountMerchantFailed = errors.ErrInternal.WithMessage("failed to get yearly total amount of merchant")

	// ErrGetMonthlyTotalAmountByMerchantsFailed indicates failure fetching monthly total amount for all merchants.
	ErrGetMonthlyTotalAmountByMerchantsFailed = errors.ErrInternal.WithMessage("failed to get monthly total amount by merchants")

	// ErrGetYearlyTotalAmountByMerchantsFailed indicates failure fetching yearly total amount for all merchants.
	ErrGetYearlyTotalAmountByMerchantsFailed = errors.ErrInternal.WithMessage("failed to get yearly total amount by merchants")

	// ErrGetMonthlyTotalAmountByApikeyFailed indicates failure fetching monthly total amount by API key.
	ErrGetMonthlyTotalAmountByApikeyFailed = errors.ErrInternal.WithMessage("failed to get monthly total amount by API key")

	// ErrGetYearlyTotalAmountByApikeyFailed indicates failure fetching yearly total amount by API key.
	ErrGetYearlyTotalAmountByApikeyFailed = errors.ErrInternal.WithMessage("failed to get yearly total amount by API key")
)
