# 📦 Package `apps`

**Source Path:** `service/auth/internal/apps`

## 🏷️ Variables

```go
var (
	port int
)
```

## 🧩 Types

### `Server`

```go
type Server struct {
	Logger logger.LoggerInterface
	DB *db.Queries
	TokenManager *auth.Manager
	Services *service.Service
	Handlers *handler.Handler
	Ctx context.Context
}
```

#### Methods

##### `Run`

Run starts the server, including the gRPC server and metrics server.

```go
func (s *Server) Run()
```

## 🚀 Functions

### `init`

init initializes the server port for the gRPC authentication service.
It retrieves the port number from the environment configuration using Viper.
If the port is not specified, it defaults to 50051.
The port can also be overridden via a command-line flag.

```go
func init()
```

