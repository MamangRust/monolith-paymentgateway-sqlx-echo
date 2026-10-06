# 📦 Package `logger`

**Source Path:** `pkg/logger`

## 🏷️ Variables

```go
var once sync.Once
```

## 🧩 Types

### `Logger`

```go
type Logger struct {
	Log *zap.Logger
}
```

#### Methods

##### `Debug`

Debug logs a message at the Debug level. It accepts a message string
and optional zap fields to include additional context in the log entry.

```go
func (l *Logger) Debug(message string, fields ...zap.Field)
```

##### `Error`

Error logs a message at the Error level. It accepts a message string
and optional zap fields to include additional context in the log entry.

```go
func (l *Logger) Error(message string, fields ...zap.Field)
```

##### `Fatal`

Fatal logs a message at the Fatal level. It accepts a message string
and optional zap fields to include additional context in the log entry.
The function will log the message and then terminate the application.

```go
func (l *Logger) Fatal(message string, fields ...zap.Field)
```

##### `Info`

Info logs a message at the Info level. It accepts a message string
and optional zap fields to include additional context in the log entry.

```go
func (l *Logger) Info(message string, fields ...zap.Field)
```

### `LoggerInterface`

```go
type LoggerInterface interface {
	Info func(message string, fields ...zap.Field)
	Fatal func(message string, fields ...zap.Field)
	Debug func(message string, fields ...zap.Field)
	Error func(message string, fields ...zap.Field)
}
```

