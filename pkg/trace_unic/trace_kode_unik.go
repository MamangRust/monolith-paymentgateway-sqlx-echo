package traceunic

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// GenerateTraceID generates a trace ID, given a prefix.
//
// The trace ID is generated as {prefix}_{date}_{random 8-character UUID},
// where {date} is the current date in the format "20060102".
//
// Note that the maximum length of the generated trace ID is 24 characters.
func GenerateTraceID(prefix string) string {
	date := time.Now().Format("20060102")
	uid := uuid.New().String()[:8]
	return fmt.Sprintf("%s_%s_%s", prefix, date, uid)
}
