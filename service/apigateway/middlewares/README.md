# 📦 Package `middlewares`

**Source Path:** `service/apigateway/internal/middlewares/`

## 🏷️ Variables

**Var:**

whiteListPaths defines a list of HTTP paths that are excluded from authentication or middleware checks.

These paths are typically public endpoints such as login, registration, or documentation routes.
Middleware such as JWT authentication should skip these paths to allow anonymous access.

```go
var whiteListPaths = []string{"/api/auth/login", "/api/auth/register", "/api/auth/hello", "/api/auth/verify-code", "/docs/", "/docs", "/swagger"}
```

## 🧩 Types

### `ApiKeyValidator`

ApiKeyValidator is responsible for validating merchant API keys via a Kafka request-response pattern.

It sends API key validation requests to a Kafka topic and listens for responses on a separate topic.
The validator tracks pending responses using correlation IDs and response channels, ensuring thread-safe access
with a mutex and supports timeouts for unresponsive requests.

```go
type ApiKeyValidator struct {
	kafka *kafka.Kafka
	logger logger.LoggerInterface
	requestTopic string
	responseTopic string
	timeout time.Duration
	responseChans map[string]chan []byte
	mu sync.Mutex
}
```

#### Methods

##### `Middleware`

Middleware returns an Echo middleware that validates the API Key in the request
header by publishing a message to a Kafka topic and waiting for a response.
If the validation is successful, the merchant ID is stored in the Echo context
and the request is passed to the next handler. If the validation fails or
times out, an HTTP error is returned.

```go
func (v *ApiKeyValidator) Middleware() echo.MiddlewareFunc
```

### `RateLimiter`

RateLimiter wraps a token bucket rate limiter to control the rate of operations,
such as API requests or background jobs.

It helps to prevent abuse and ensure fair usage of system resources by limiting
how frequently certain actions can be performed.

```go
type RateLimiter struct {
	limiter *rate.Limiter
}
```

#### Methods

##### `Limit`

Limit limits the number of requests to the given handler to the specified
rate. If the rate is exceeded, it returns a 429 (Too Many Requests) error.
The rate is specified in requests per second, and the burst parameter
specifies how many requests can be made before the rate limiting kicks in.

```go
func (rl *RateLimiter) Limit(next echo.HandlerFunc) echo.HandlerFunc
```

### `RoleValidator`

RoleValidator is responsible for validating user roles via Kafka-based request-response messaging.

This struct sends role validation requests through Kafka and listens for asynchronous responses.
It includes logic to manage concurrent access to response channels and timeout control.

```go
type RoleValidator struct {
	kafka *kafka.Kafka
	logger logger.LoggerInterface
	requestTopic string
	responseTopic string
	timeout time.Duration
	responseChans map[string]chan *response.RoleResponsePayload
	mu sync.RWMutex
}
```

#### Methods

##### `Middleware`

Middleware returns an Echo middleware that validates the user role by publishing
a message to a Kafka topic and waiting for a response. If the validation is
successful, the role names are stored in the Echo context and the request is
passed to the next handler. If the validation fails or times out, an HTTP error
is returned.

```go
func (v *RoleValidator) Middleware() echo.MiddlewareFunc
```

##### `extractUserID`

```go
func (v *RoleValidator) extractUserID(userIDVal interface{}) (int, error)
```

##### `sendValidationRequest`

```go
func (v *RoleValidator) sendValidationRequest(userID int, correlationID string) error
```

### `merchantResponseHandler`

merchantResponseHandler handles Kafka consumer group lifecycle events for merchant-related API key validation responses.

It implements the sarama.ConsumerGroupHandler interface and is primarily responsible
for handling the setup and cleanup phases of the Kafka consumer group session.
The actual message processing (ConsumeClaim) should be implemented to process incoming responses.

```go
type merchantResponseHandler struct {
	validator *ApiKeyValidator
}
```

#### Methods

##### `Cleanup`

Cleanup is called when a Kafka consumer group session ends.

It can be used to perform cleanup tasks, such as closing resources or finalizing state.
This implementation performs no cleanup.

```go
func (h *merchantResponseHandler) Cleanup(sarama.ConsumerGroupSession) error
```

##### `ConsumeClaim`

ConsumeClaim processes incoming Kafka messages from the specified consumer group claim.
It unmarshals each message into a payload map and retrieves the correlation ID.
If a valid correlation ID is found, it sends the message value to the corresponding
response channel managed by the validator. Each message is marked as processed
in the consumer group session.

```go
func (h *merchantResponseHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error
```

##### `Setup`

Setup is called when a new Kafka consumer group session begins.

This method can be used to perform any necessary initialization before message consumption starts.
In this implementation, no setup is required.

```go
func (h *merchantResponseHandler) Setup(sarama.ConsumerGroupSession) error
```

### `roleResponseHandler`

roleResponseHandler handles Kafka consumer group events related to role validation responses.

It implements the sarama.ConsumerGroupHandler interface, and is responsible for
setting up and cleaning up the Kafka consumer session when consuming role validation responses.

```go
type roleResponseHandler struct {
	validator *RoleValidator
}
```

#### Methods

##### `Cleanup`

Cleanup is called at the end of a Kafka consumer group session.

It can be used to release resources allocated during Setup or message consumption.
In this implementation, it performs no cleanup and returns nil.

```go
func (h *roleResponseHandler) Cleanup(sarama.ConsumerGroupSession) error
```

##### `ConsumeClaim`

ConsumeClaim processes incoming Kafka messages from the specified consumer group claim.
It unmarshals each message into a payload map and retrieves the correlation ID.
If a valid correlation ID is found, it sends the message value to the corresponding
response channel managed by the validator. Each message is marked as processed
in the consumer group session.

```go
func (h *roleResponseHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error
```

##### `Setup`

Setup is called at the beginning of a new Kafka consumer group session.

It can be used to initialize resources or state before message consumption begins.
In this implementation, it performs no setup and returns nil.

```go
func (h *roleResponseHandler) Setup(sarama.ConsumerGroupSession) error
```

## 🚀 Functions

### `RequireRoles`

RequireRoles is an Echo middleware that checks if the value of the key "role_names"
(set by the RoleValidator middleware) contains any of the given allowedRoles. If
it does, the request is allowed to proceed. Otherwise, it returns a 403 status code
with an error message indicating that the role is not permitted.

Example:

	e.GET("/admin", RequireRoles("admin", "superadmin"), func(c echo.Context) error {
		// only users with role "admin" or "superadmin" can access this route
		return c.String(http.StatusOK, "Hello, "+c.Get("user_id").(string))
	})

```go
func RequireRoles(allowedRoles ...string) echo.MiddlewareFunc
```

### `WebSecurityConfig`

WebSecurityConfig adds JWT middleware to an echo router.

The middleware uses the SigningKey from the config file. It also sets the
Skipper to skipAuth, which allows the following paths to be accessed without
a valid JWT:

- /api/auth/login
- /api/auth/register
- /api/auth/hello
- /api/auth/verify-code
- /docs/
- /docs
- /swagger

The SuccessHandler is used to add the subject of the JWT to the context
under the key "user_id".

The ErrorHandler is used to return a 401 Unauthorized status code in case
of a JWT error.

```go
func WebSecurityConfig(e *echo.Echo)
```

### `skipAuth`

skipAuth is the Skipper used in the JWT middleware.

It returns true for the following paths, which are skipped by the JWT middleware:

- /api/auth/login
- /api/auth/register
- /api/auth/hello
- /api/auth/verify-code
- /docs/
- /docs
- /swagger

```go
func skipAuth(e echo.Context) bool
```

