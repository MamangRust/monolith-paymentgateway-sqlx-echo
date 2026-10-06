# 📦 Package `traceunic`

**Source Path:** `pkg/trace_unic/`

## 🚀 Functions

### `GenerateTraceID`

GenerateTraceID generates a trace ID, given a prefix.

The trace ID is generated as {prefix}_{date}_{random 8-character UUID},
where {date} is the current date in the format "20060102".

Note that the maximum length of the generated trace ID is 24 characters.

```go
func GenerateTraceID(prefix string) string
```

