package observability

import "testing"

func TestMaskIdentifier(t *testing.T) {
	cases := map[string]string{
		"":                 "",
		"123":              "****",
		"1234":             "****",
		"4111111111111111": "************1111",
		"api-key-1234":     "********1234",
	}

	for input, expected := range cases {
		if got := MaskIdentifier(input); got != expected {
			t.Errorf("MaskIdentifier(%q) = %q, want %q", input, got, expected)
		}
	}
}
