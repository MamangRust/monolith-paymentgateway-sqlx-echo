# 📦 Package `app`

**Source Path:** `service/apigateway/internal/app`

## 🧩 Types

### `Client`

@securityDefinitions.apikey BearerAuth
@in Header
@name Authorization

```go
type Client struct {
	App *echo.Echo
	Logger logger.LoggerInterface
}
```

#### Methods

##### `Shutdown`

Shutdown the server.

This method is idempotent and will not return an error if the server is already shutdown.

The context is used to receive any error encountered that caused the shutdown.

The client will not be usable after this method is called.

```go
func (c *Client) Shutdown(ctx context.Context) error
```

### `ServiceAddresses`

ServiceAddresses holds the addresses of all the microservices.

```go
type ServiceAddresses struct {
	Auth string
	Role string
	Card string
	Merchant string
	User string
	Saldo string
	Topup string
	Transaction string
	Transfer string
	Withdraw string
}
```

## 🚀 Functions

### `closeConnections`

closeConnections gracefully closes all gRPC connections contained within the
ServiceConnections struct. It logs any errors encountered during the closure
of each connection using the provided logger.

Parameters:
  - conns: A pointer to a ServiceConnections struct holding the gRPC connections
    to be closed.
  - log: An instance of LoggerInterface used to log errors if any connection fails
    to close.

The function iterates over each connection in the ServiceConnections struct.
If a connection is not nil, it attempts to close it and logs any errors that
occur during the closure process.

```go
func closeConnections(conns *handler.ServiceConnections, log logger.LoggerInterface)
```

### `createConnection`

createConnection creates a gRPC client connection to a service given its address and
service name. It logs the connection attempt and any errors that occur.

The service name is used for logging purposes only.

The function uses the insecure gRPC transport credentials, as the services are assumed
to be running in the same network and no encryption is needed.

```go
func createConnection(address, serviceName string, logger logger.LoggerInterface) (*grpc.ClientConn, error)
```

### `createServiceConnections`

createServiceConnections creates a set of gRPC connections to the microservices.

It takes a pointer to a ServiceAddresses struct and a logger interface as arguments.
The ServiceAddresses struct contains the addresses of all the microservices.

The function returns a pointer to a ServiceConnections struct which contains the gRPC connections
to the microservices or an error if any of the connections fail.

The connections are created using the createConnection function which returns a gRPC connection
to the specified microservice or an error if the connection fails.

The function will return an error if any of the connections fail.

```go
func createServiceConnections(addresses *ServiceAddresses, logger logger.LoggerInterface) (*handler.ServiceConnections, error)
```

### `getEnvOrDefault`

getEnvOrDefault retrieves the value of an environment variable and returns a default
value if it is not set.

```go
func getEnvOrDefault(key, defaultValue string) string
```

### `setupEcho`

setupEcho sets up the Echo web framework with the necessary middleware and routes.

It returns a pointer to the Echo instance.

The middleware used are:
  - middleware.Logger(): logs each request
  - middleware.Recover(): recovers from panics and logs the error
  - limiter.Limit: limits the number of requests per second to 20 and the number of requests per minute to 50
  - middleware.CORSWithConfig: enables CORS with the specified configuration
  - middlewares.WebSecurityConfig: sets up the web security configuration for the Echo instance
  - echoSwagger.WrapHandler: sets up the Swagger UI handler

```go
func setupEcho() *echo.Echo
```

