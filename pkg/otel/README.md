# 📦 Package `otel_pkg`

**Source Path:** `pkg/otel`

## 🧩 Types

### `ExporterFactory`

ExporterFactory is a function that creates a trace exporter.

```go
type ExporterFactory func(ctx context.Context) (sdktrace.SpanExporter, error)
```

### `mockExporter`

mockExporter is a no-op SpanExporter for testing purposes.

```go
type mockExporter struct {
}
```

#### Methods

##### `ExportSpans`

ExportSpans implements sdktrace.SpanExporter.ExportSpans.

It is a no-op, and always returns nil.

```go
func (m *mockExporter) ExportSpans(_ context.Context, _ []sdktrace.ReadOnlySpan) error
```

##### `Shutdown`

Shutdown does nothing and returns no error.

```go
func (m *mockExporter) Shutdown(_ context.Context) error
```

## 🚀 Functions

### `DefaultExporterFactory`

DefaultExporterFactory creates an OTLP GRPC trace exporter based on the
environment variable OTEL_ENDPOINT. The exporter is configured to be
insecure, meaning it does not validate the server's certificate.

This function should be used as a fallback when no other exporter factory
is specified.

```go
func DefaultExporterFactory(ctx context.Context) (sdktrace.SpanExporter, error)
```

### `InitTracerProvider`

InitTracerProvider initializes an OpenTelemetry TracerProvider for a given service.
It sets up a trace exporter using the OTLP gRPC protocol and configures the TracerProvider
with resource attributes such as service name, version, and environment. It also sets up
a batch span processor and a composite text map propagator for context propagation.

Parameters:
  - service: The name of the service for which the TracerProvider is being initialized.
  - ctx: The context for managing the lifecycle of the TracerProvider and trace exporter.

Returns:
  - A function to shut down the TracerProvider, releasing any resources held.
  - An error if there was a failure during the initialization of the trace exporter or resource.

```go
func InitTracerProvider(service string, ctx context.Context) (func(context.Context) error, error)
```

### `TestInitTracerProvider_WithMockExporter`

TestInitTracerProvider_WithMockExporter tests that initTracerProviderWithExporter returns
a valid shutdown function that can be used to shut down the TracerProvider.

It also checks that the shutdown function returns no error when called.

```go
func TestInitTracerProvider_WithMockExporter(t *testing.T)
```

### `fakeExporterFactory`

fakeExporterFactory is a fake ExporterFactory for testing purposes, which
returns a no-op SpanExporter.

```go
func fakeExporterFactory(_ context.Context) (sdktrace.SpanExporter, error)
```

### `initTracerProviderWithExporter`

initTracerProviderWithExporter initializes an OpenTelemetry TracerProvider for a given service using the given trace exporter factory.
It sets up a trace exporter using the OTLP gRPC protocol and configures the TracerProvider
with resource attributes such as service name, version, and environment. It also sets up
a batch span processor and a composite text map propagator for context propagation.

Parameters:
  - service: The name of the service for which the TracerProvider is being initialized.
  - ctx: The context for managing the lifecycle of the TracerProvider and trace exporter.
  - factory: A function that creates a trace exporter.

Returns:
  - A function to shut down the TracerProvider, releasing any resources held.
  - An error if there was a failure during the initialization of the trace exporter or resource.

```go
func initTracerProviderWithExporter(service string, ctx context.Context, factory ExporterFactory) (func(context.Context) error, error)
```

