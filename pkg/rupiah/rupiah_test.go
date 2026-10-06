package rupiah

import "testing"

func TestRupiahFormat(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1000", "Rp.1000"},
		{"15000.75", "Rp.15001"},
		{"999.49", "Rp.999"},
		{"abc", "Rp 0"},
		{"", "Rp 0"},
		{"0", "Rp.0"},
		{"123456789", "Rp.123456789"},
	}

	for _, test := range tests {
		result := RupiahFormat(test.input)
		if result != test.expected {
			t.Errorf("RupiahFormat(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
