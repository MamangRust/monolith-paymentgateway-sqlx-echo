package convert

import (
	"time"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// FormatTimePtr converts *time.Time to formatted string (RFC3339).
// Returns empty string if nil.
func FormatTimePtr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

// FormatTimeRFC3339 converts *time.Time to RFC3339 string.
// Returns empty string if nil.
func FormatTimeRFC3339(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(time.RFC3339)
}

// StrVal converts *string to string.
// Returns empty string if nil.
func StrVal(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// StrValToWrappers converts *string to *wrapperspb.StringValue.
// Returns nil if nil.
func StrValToWrappers(s *string) *wrapperspb.StringValue {
	if s == nil {
		return nil
	}
	return wrapperspb.String(*s)
}

// TimeToWrappers converts *time.Time to *wrapperspb.StringValue.
// Returns nil if nil.
func TimeToWrappers(t *time.Time) *wrapperspb.StringValue {
	if t == nil {
		return nil
	}
	return wrapperspb.String(t.Format(time.RFC3339))
}

// Int32Val converts *int32 to int32.
// Returns 0 if nil.
func Int32Val(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}
