package traceunic

import (
	"strings"
	"testing"
	"time"
)

func TestGenerateTraceID(t *testing.T) {
	prefix := "TX"
	traceID := GenerateTraceID(prefix)

	parts := strings.Split(traceID, "_")
	if len(parts) != 3 {
		t.Fatalf("expected 3 parts in traceID, got %d: %v", len(parts), traceID)
	}

	if parts[0] != prefix {
		t.Errorf("expected prefix %s, got %s", prefix, parts[0])
	}

	today := time.Now().Format("20060102")
	if parts[1] != today {
		t.Errorf("expected date part %s, got %s", today, parts[1])
	}

	if len(parts[2]) != 8 {
		t.Errorf("expected UID part to be 8 characters, got %d: %s", len(parts[2]), parts[2])
	}
}
