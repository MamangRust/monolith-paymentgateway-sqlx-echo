# 📦 Package `handler`

**Source Path:** `apigateway/internal/handler`

## 🧩 Types

### `Deps`

Deps holds dependencies required to initialize HTTP API handlers.

```go
type Deps struct {
	Kafka *kafka.Kafka
	Token auth.TokenManager
	E *echo.Echo
	Logger logger.LoggerInterface
	ServiceConnections *ServiceConnections
}
```

### `ServiceConnections`

ServiceConnections holds gRPC connections to external monolith.

```go
type ServiceConnections struct {
	Auth *grpc.ClientConn
	Role *grpc.ClientConn
	Card *grpc.ClientConn
	Merchant *grpc.ClientConn
	User *grpc.ClientConn
	Saldo *grpc.ClientConn
	Topup *grpc.ClientConn
	Transaction *grpc.ClientConn
	Transfer *grpc.ClientConn
	Withdraw *grpc.ClientConn
}
```

## 🚀 Functions

### `NewHandler`

NewHandler sets up all the handlers for the API Gateway.
It takes a pointer to a Deps struct, which contains all the dependencies
required to set up the handlers.

```go
func NewHandler(deps *Deps)
```

