package merchantrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthlyPaymentMethodsMerchantFailed indicates failure fetching monthly payment methods for a merchant.
	ErrGetMonthlyPaymentMethodsMerchantFailed = errors.ErrInternal.WithMessage("failed to get monthly payment methods of merchant")

	// ErrGetYearlyPaymentMethodMerchantFailed indicates failure fetching yearly payment methods for a merchant.
	ErrGetYearlyPaymentMethodMerchantFailed = errors.ErrInternal.WithMessage("failed to get yearly payment method of merchant")

	// ErrGetMonthlyPaymentMethodByMerchantsFailed indicates failure fetching monthly payment methods for all merchants.
	ErrGetMonthlyPaymentMethodByMerchantsFailed = errors.ErrInternal.WithMessage("failed to get monthly payment method by merchants")

	// ErrGetYearlyPaymentMethodByMerchantsFailed indicates failure fetching yearly payment methods for all merchants.
	ErrGetYearlyPaymentMethodByMerchantsFailed = errors.ErrInternal.WithMessage("failed to get yearly payment method by merchants")

	// ErrGetMonthlyPaymentMethodByApikeyFailed indicates failure fetching monthly payment methods by API key.
	ErrGetMonthlyPaymentMethodByApikeyFailed = errors.ErrInternal.WithMessage("failed to get monthly payment method by API key")

	// ErrGetYearlyPaymentMethodByApikeyFailed indicates failure fetching yearly payment methods by API key.
	ErrGetYearlyPaymentMethodByApikeyFailed = errors.ErrInternal.WithMessage("failed to get yearly payment method by API key")
)
