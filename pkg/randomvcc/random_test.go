package randomvcc

import (
	"regexp"
	"testing"
)

func TestRandomCardNumber(t *testing.T) {
	cardNumber, err := RandomCardNumber()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(cardNumber) != 16 {
		t.Errorf("expected card number length 16, got %d (%s)", len(cardNumber), cardNumber)
	}

	if cardNumber[0] != '4' {
		t.Errorf("expected card number to start with 4, got %s", cardNumber)
	}

	matched, _ := regexp.MatchString(`^\d{16}$`, cardNumber)
	if !matched {
		t.Errorf("card number is not numeric or not 16 digits: %s", cardNumber)
	}
}
