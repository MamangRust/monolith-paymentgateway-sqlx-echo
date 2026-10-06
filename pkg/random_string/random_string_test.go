package randomstring

import (
	"regexp"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	allowed := regexp.MustCompile(`^[A-Za-z0-9]*$`)

	tests := []struct {
		name       string
		length     int
		wantErr    bool
		wantLength int
	}{
		{"zero length", 0, false, 0},
		{"small length", 1, false, 1},
		{"medium length", 16, false, 16},
		{"large length", 100, false, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := GenerateRandomString(tt.length)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GenerateRandomString(%d) error = %v, wantErr %v", tt.length, err, tt.wantErr)
			}
			if len(s) != tt.wantLength {
				t.Errorf("GenerateRandomString(%d) length = %d, want %d", tt.length, len(s), tt.wantLength)
			}
			if !allowed.MatchString(s) {
				t.Errorf("GenerateRandomString(%d) contains invalid characters: %q", tt.length, s)
			}
		})
	}

	t.Run("randomness", func(t *testing.T) {
		s1, err1 := GenerateRandomString(32)
		s2, err2 := GenerateRandomString(32)
		if err1 != nil || err2 != nil {
			t.Fatalf("unexpected error: %v, %v", err1, err2)
		}
		if s1 == s2 {
			t.Errorf("expected two different random strings, but got identical:\n%s", s1)
		}
	})
}
