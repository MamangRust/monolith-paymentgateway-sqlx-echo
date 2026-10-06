package methodtopup

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPaymentMethodValidator(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedValid bool
	}{
		{"Valid - lowercase", "dana", true},
		{"Valid - uppercase", "OVO", true},
		{"Valid - mixed case", "GoPay", true},
		{"Valid - space case", "american express", true},
		{"Valid - card", "visa", true},
		{"Invalid - not in list", "shopeepay", false},
		{"Invalid - typo", "mandri", false},
		{"Invalid - empty", "", false},
		{"Invalid - special char", "paypal!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PaymentMethodValidator(tt.input)
			assert.Equal(t, tt.expectedValid, result)
		})
	}
}
