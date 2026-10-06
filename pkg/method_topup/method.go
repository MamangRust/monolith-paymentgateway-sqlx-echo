package methodtopup

import "strings"

// PaymentMethodValidator will validate the payment method string with list of known payment methods.
//
// It will check if the given paymentMethod is in the list of known payment methods.
// If the payment method is in the list, it will return true, otherwise it will return false.
func PaymentMethodValidator(paymentMethod string) bool {
	paymentRules := []string{
		"alfamart",
		"indomart",
		"lawson",
		"dana",
		"ovo",
		"gopay",
		"linkaja",
		"jenius",
		"fastpay",
		"kudo",
		"bri",
		"mandiri",
		"bca",
		"bni",
		"bukopin",
		"e-banking",
		"visa",
		"mastercard",
		"discover",
		"american express",
		"paypal",
	}

	paymentMethodLower := strings.ToLower(paymentMethod)
	for _, rule := range paymentRules {
		if paymentMethodLower == rule {
			return true
		}
	}

	return false
}
