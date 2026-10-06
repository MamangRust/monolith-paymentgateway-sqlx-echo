package rupiah

import (
	"fmt"
	"strconv"
)

// RupiahFormat formats a given number string into a string
// with a 'Rp. ' prefix and thousand separator.
//
// If the given number string is invalid, it returns "Rp 0".
func RupiahFormat(digit string) string {
	digitNumber, err := strconv.ParseFloat(digit, 64)

	if err != nil {
		return "Rp 0"
	}

	formatter := fmt.Sprintf("Rp.%.0f", digitNumber)
	return formatter
}
