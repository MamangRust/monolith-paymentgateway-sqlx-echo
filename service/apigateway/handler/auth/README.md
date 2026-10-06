# 📦 Package `authhandler`

**Source Path:** `service/apigateway/internal/handler/auth`

## 🧩 Types

### `DepsAuth`

```go
type DepsAuth struct {
	Client *grpc.ClientConn
	E *echo.Echo
	Logger logger.LoggerInterface
}
```

### `authHandleApi`

authHandleApi is a handler for the authentication service HTTP endpoints.

It encapsulates the gRPC client for the auth service, logger, response mapper,
OpenTelemetry tracing, and Prometheus metrics instrumentation.

```go
type authHandleApi struct {
	client pb.AuthServiceClient
	logger logger.LoggerInterface
	mapper apimapper.AuthResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `ForgotPassword`

ForgotPassword godoc
@Summary Sends a reset token to the user's email
@Tags Auth
@Description Initiates password reset by sending a reset token to the provided email.
@Accept json
@Produce json
@Param request body requests.ForgotPasswordRequest true "Forgot Password Request"
@Success 200 {object} response.ApiResponseForgotPassword
@Failure 400 {object} response.ErrorResponse
@Router /auth/forgot-password [post]

```go
func (h *authHandleApi) ForgotPassword(c echo.Context) error
```

##### `GetMe`

GetMe godoc
@Summary Get current user information
@Tags Auth
@Security Bearer
@Description Retrieves the current user's information using a valid access token from the Authorization header.
@Produce json
@Security BearerToken
@Success 200 {object} response.ApiResponseGetMe "Success"
@Failure 401 {object} response.ErrorResponse "Unauthorized"
@Failure 500 {object} response.ErrorResponse "Internal Server Error"
@Router /api/auth/me [get]

```go
func (h *authHandleApi) GetMe(c echo.Context) error
```

##### `HandleHello`

HandleHello godoc
@Summary Returns a "Hello" message
@Tags Auth
@Description Returns a simple "Hello" message for testing purposes.
@Produce json
@Success 200 {string} string "Hello"
@Router /auth/hello [get]

```go
func (h *authHandleApi) HandleHello(c echo.Context) error
```

##### `Login`

Login godoc
@Summary Authenticate a user
@Tags Auth
@Description Authenticates a user using the provided email and password.
@Accept json
@Produce json
@Param request body requests.AuthRequest true "User login credentials"
@Success 200 {object} response.ApiResponseLogin "Success"
@Failure 400 {object} response.ErrorResponse "Bad Request"
@Failure 500 {object} response.ErrorResponse "Internal Server Error"
@Router /api/auth/login [post]

```go
func (h *authHandleApi) Login(c echo.Context) error
```

##### `RefreshToken`

RefreshToken godoc
@Summary Refresh access token
@Tags Auth
@Security Bearer
@Description Refreshes the access token using a valid refresh token.
@Accept json
@Produce json
@Param request body requests.RefreshTokenRequest true "Refresh token data"
@Success 200 {object} response.ApiResponseRefreshToken "Success"
@Failure 400 {object} response.ErrorResponse "Bad Request"
@Failure 500 {object} response.ErrorResponse "Internal Server Error"
@Router /api/auth/refresh-token [post]

```go
func (h *authHandleApi) RefreshToken(c echo.Context) error
```

##### `Register`

Register godoc
@Summary Register a new user
@Tags Auth
@Description Registers a new user with the provided details.
@Accept json
@Produce json
@Param request body requests.CreateUserRequest true "User registration data"
@Success 200 {object} response.ApiResponseRegister "Success"
@Failure 400 {object} response.ErrorResponse "Bad Request"
@Failure 500 {object} response.ErrorResponse "Internal Server Error"
@Router /api/auth/register [post]

```go
func (h *authHandleApi) Register(c echo.Context) error
```

##### `ResetPassword`

ResetPassword godoc
@Summary Resets the user's password using a reset token
@Tags Auth
@Description Allows user to reset their password using a valid reset token.
@Accept json
@Produce json
@Param request body requests.CreateResetPasswordRequest true "Reset Password Request"
@Success 200 {object} response.ApiResponseResetPassword
@Failure 400 {object} response.ErrorResponse
@Router /auth/reset-password [post]

```go
func (h *authHandleApi) ResetPassword(c echo.Context) error
```

##### `VerifyCode`

VerifyCode godoc
@Summary Verifies the user using a verification code
@Tags Auth
@Description Verifies the user's email using the verification code provided in the query parameter.
@Produce json
@Param verify_code query string true "Verification Code"
@Success 200 {object} response.ApiResponseVerifyCode
@Failure 400 {object} response.ErrorResponse
@Router /auth/verify-code [get]

```go
func (h *authHandleApi) VerifyCode(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *authHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *authHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `authHandleParams`

authHandleParams contains the dependencies required to initialize auth HTTP handlers.

```go
type authHandleParams struct {
	client pb.AuthServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.AuthResponseMapper
}
```

## 🚀 Functions

### `RegisterAuthHandler`

```go
func RegisterAuthHandler(deps *DepsAuth)
```

