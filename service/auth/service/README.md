# 📦 Package `service`

**Source Path:** `service/auth/internal/service`

## 🧩 Types

### `Deps`

Deps holds all dependencies required to initialize the Service.

```go
type Deps struct {
	Context context.Context
	ErrorHandler *errorhandler.ErrorHandler
	Mencache *mencache.Mencache
	Repositories *repository.Repositories
	Token auth.TokenManager
	Hash hash.HashPassword
	Logger logger.LoggerInterface
	Kafka *kafka.Kafka
	Mapper responseservice.UserResponseMapper
}
```

### `IdentifyService`

IdentifyService defines the contract for identity verification and token management.
It provides methods to refresh tokens and get user information.

```go
type IdentifyService interface {
	RefreshToken func(token string) (*response.TokenResponse, *response.ErrorResponse)
	GetMe func(token string) (*response.UserResponse, *response.ErrorResponse)
}
```

### `IdentityServiceParams`

IdentityServiceParams holds the parameters for the identity service.

```go
type IdentityServiceParams struct {
	Ctx context.Context
	ErrorHandler errorhandler.IdentityErrorHandler
	ErrorToken errorhandler.TokenErrorHandler
	Mencache mencache.IdentityCache
	Token auth.TokenManager
	RefreshToken repository.RefreshTokenRepository
	User repository.UserRepository
	Logger logger.LoggerInterface
	Mapping responseservice.UserResponseMapper
	TokenService *tokenService
}
```

### `LoginService`

LoginService defines the contract for user authentication operations.
It provides methods to authenticate existing users.

```go
type LoginService interface {
	Login func(request *requests.AuthRequest) (*response.TokenResponse, *response.ErrorResponse)
}
```

### `LoginServiceParams`

LoginServiceParams groups all dependencies required to initialize a new loginService.

```go
type LoginServiceParams struct {
	Ctx context.Context
	ErrorPassword errorhandler.PasswordErrorHandler
	ErrorToken errorhandler.TokenErrorHandler
	ErrorHandler errorhandler.LoginErrorHandler
	Mencache mencache.LoginCache
	Logger logger.LoggerInterface
	Hash hash.HashPassword
	UserRepository repository.UserRepository
	RefreshToken repository.RefreshTokenRepository
	Token auth.TokenManager
	TokenService *tokenService
}
```

### `PasswordResetService`

PasswordResetService defines the contract for password recovery operations.
It provides methods to handle the complete password reset flow.

```go
type PasswordResetService interface {
	ForgotPassword func(email string) (bool, *response.ErrorResponse)
	ResetPassword func(request *requests.CreateResetPasswordRequest) (bool, *response.ErrorResponse)
	VerifyCode func(code string) (bool, *response.ErrorResponse)
}
```

### `PasswordResetServiceParams`

```go
type PasswordResetServiceParams struct {
	Ctx context.Context
	ErrorHandler errorhandler.PasswordResetErrorHandler
	ErrorRandomString errorhandler.RandomStringErrorHandler
	ErrorMarshal errorhandler.MarshalErrorHandler
	ErrorPassword errorhandler.PasswordErrorHandler
	ErrorKafka errorhandler.KafkaErrorHandler
	Mencache mencache.PasswordResetCache
	Kafka *kafka.Kafka
	Logger logger.LoggerInterface
	User repository.UserRepository
	ResetToken repository.ResetTokenRepository
}
```

### `RegisterServiceParams`

```go
type RegisterServiceParams struct {
	Ctx context.Context
	ErrorHandler errorhandler.RegisterErrorHandler
	ErrorPassword errorhandler.PasswordErrorHandler
	ErrorRandomString errorhandler.RandomStringErrorHandler
	ErrorMarshal errorhandler.MarshalErrorHandler
	ErrorKafka errorhandler.KafkaErrorHandler
	Mencache mencache.RegisterCache
	User repository.UserRepository
	Role repository.RoleRepository
	UserRole repository.UserRoleRepository
	Hash hash.HashPassword
	Kafka *kafka.Kafka
	Logger logger.LoggerInterface
	Mapping responseservice.UserResponseMapper
}
```

### `RegistrationService`

RegistrationService defines the contract for user registration operations.
It provides a method to register new users in the system.

```go
type RegistrationService interface {
	Register func(request *requests.RegisterRequest) (*response.UserResponse, *response.ErrorResponse)
}
```

### `Service`

Service contains the core user authentication and identity services.

```go
type Service struct {
	Login LoginService
	Register RegistrationService
	PasswordReset PasswordResetService
	Identify IdentifyService
}
```

### `identityService`

identityService is the implementation of the identity service.

```go
type identityService struct {
	ctx context.Context
	errorhandler errorhandler.IdentityErrorHandler
	errorToken errorhandler.TokenErrorHandler
	mencache mencache.IdentityCache
	trace trace.Tracer
	logger logger.LoggerInterface
	token auth.TokenManager
	refreshToken repository.RefreshTokenRepository
	user repository.UserRepository
	mapping responseservice.UserResponseMapper
	tokenService *tokenService
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `GetMe`

GetMe retrieves user information using a valid access token.
It validates the access token, fetches user details from the cache or database, and returns the user's information.
Parameters:
- token: The access token string used to authenticate the user.
Returns:
- UserResponse: Contains the authenticated user's details if successful.
- ErrorResponse: Contains error details if the token is invalid or user retrieval fails.

```go
func (s *identityService) GetMe(token string) (*response.UserResponse, *response.ErrorResponse)
```

##### `RefreshToken`

RefreshToken generates new access tokens using a valid refresh token.
Takes the refresh token string and returns:
- TokenResponse with new access and refresh tokens
- ErrorResponse if token refresh fails

```go
func (s *identityService) RefreshToken(token string) (*response.TokenResponse, *response.ErrorResponse)
```

##### `recordMetrics`

recordMetrics records a Prometheus metric for the given method and status.
It increments a counter and records the duration since the provided start time.

```go
func (s *identityService) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

startTracingAndLogging initiates tracing and logging for the specified method.
It returns a trace.Span, an end function to conclude the tracing with a status,
the initial status string, and a logSuccess function to record successful events.

Parameters:
- method: The name of the method being traced and logged.
- attrs: Optional attributes to set on the span for additional context.

Returns:
- trace.Span: The span object associated with the tracing.
- func(string): A function to end the tracing with the given status.
- string: The initial status, set to "success".
- func(string, ...zap.Field): A function to log successful events with additional fields.

```go
func (s *identityService) startTracingAndLogging(method string, attrs ...attribute.KeyValue) (trace.Span, func(string), string, func(string, ...zap.Field))
```

### `loginService`

loginService is the implementation of the LoginService interface.
It handles the logic for authenticating users, validating passwords,
issuing tokens, and caching login sessions.

```go
type loginService struct {
	ctx context.Context
	errorPassword errorhandler.PasswordErrorHandler
	errorToken errorhandler.TokenErrorHandler
	errorHandler errorhandler.LoginErrorHandler
	mencache mencache.LoginCache
	logger logger.LoggerInterface
	hash hash.HashPassword
	user repository.UserRepository
	refreshToken repository.RefreshTokenRepository
	token auth.TokenManager
	trace trace.Tracer
	tokenService *tokenService
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `Login`

Login authenticates the user using the provided credentials.
It first checks for a cached login token; if found, it returns the cached token.
If not cached, it verifies the user's email and password, and upon success,
generates new access and refresh tokens. The tokens are then cached and returned.
Tracing, logging, and error handling are integrated for detailed monitoring and response.
Parameters:
- request: AuthRequest containing the user's email and password.
Returns:
- TokenResponse with the access and refresh tokens if successful.
- ErrorResponse if any authentication step fails.

```go
func (s *loginService) Login(request *requests.AuthRequest) (*response.TokenResponse, *response.ErrorResponse)
```

##### `recordMetrics`

recordMetrics records a Prometheus metric for the given method and status.
It increments a counter and records the duration since the provided start time.

```go
func (s *loginService) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

startTracingAndLogging initializes a tracing span and logging for the specified method.
It returns the trace span, a function to end the span with a status, the initial status
of the operation, and a function to log a success message.

The `end` function should be called with the status of the operation, which can be
either "success" or "error". The `logSuccess` function can be used to log a success
message, adding the message as an event to the trace span and logging it with any
provided zap fields.

```go
func (s *loginService) startTracingAndLogging(method string, attrs ...attribute.KeyValue) (trace.Span, func(string), string, func(string, ...zap.Field))
```

### `passwordResetService`

```go
type passwordResetService struct {
	ctx context.Context
	errorhandler errorhandler.PasswordResetErrorHandler
	errorRandomString errorhandler.RandomStringErrorHandler
	errorMarshal errorhandler.MarshalErrorHandler
	errorPassword errorhandler.PasswordErrorHandler
	errorKafka errorhandler.KafkaErrorHandler
	mencache mencache.PasswordResetCache
	trace trace.Tracer
	kafka *kafka.Kafka
	logger logger.LoggerInterface
	user repository.UserRepository
	resetToken repository.ResetTokenRepository
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `ForgotPassword`

ForgotPassword initiates the password reset process.
Takes an email address as a string and returns boolean status and ErrorResponse.
Status is true if the password reset was successful, and false if it failed.
ErrorResponse is nil if the request was successful, and set to the error response
if the request failed.

```go
func (s *passwordResetService) ForgotPassword(email string) (bool, *response.ErrorResponse)
```

##### `ResetPassword`

ResetPassword completes the password reset process.
Takes a CreateResetPasswordRequest containing the new password and returns:
- boolean indicating if the password was successfully reset
- ErrorResponse if the reset fails

```go
func (s *passwordResetService) ResetPassword(req *requests.CreateResetPasswordRequest) (bool, *response.ErrorResponse)
```

##### `VerifyCode`

VerifyCode checks the validity of a password reset verification code.
It performs the following operations:
1. Finds the user associated with the given verification code.
2. Updates the user's verification status to true.
3. Deletes the verification code from the cache.
4. Sends a verification success email to the user via Kafka messaging.

Returns:
- A boolean indicating if the code verification was successful.
- An ErrorResponse if any step in the process fails.

```go
func (s *passwordResetService) VerifyCode(code string) (bool, *response.ErrorResponse)
```

##### `recordMetrics`

recordMetrics records Prometheus metrics for the passwordResetService.
It increments the request counter and observes the request duration
for the given method and status, using the provided start time.

```go
func (s *passwordResetService) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

startTracingAndLogging starts a tracing span and logging for a given method,
and returns the span, a function to end the span, the initial status of the
operation, and a function to log a message as a success event.

The returned end function takes a status string and updates the span's status
and ends the span. The status string can be one of "success" or "error".

The returned logSuccess function takes a message and any number of zap fields
and logs the message as a success event, and adds the event to the span.

```go
func (s *passwordResetService) startTracingAndLogging(method string, attrs ...attribute.KeyValue) (trace.Span, func(string), string, func(string, ...zap.Field))
```

### `registerService`

```go
type registerService struct {
	ctx context.Context
	trace trace.Tracer
	errohandler errorhandler.RegisterErrorHandler
	errorPassword errorhandler.PasswordErrorHandler
	errorRandomString errorhandler.RandomStringErrorHandler
	errorMarshal errorhandler.MarshalErrorHandler
	errorKafka errorhandler.KafkaErrorHandler
	mencache mencache.RegisterCache
	user repository.UserRepository
	role repository.RoleRepository
	userRole repository.UserRoleRepository
	hash hash.HashPassword
	kafka *kafka.Kafka
	logger logger.LoggerInterface
	mapping responseservice.UserResponseMapper
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `Register`

Register handles new user registration.
Takes a RegisterRequest containing user details and returns:
- UserResponse with registered user data if successful
- ErrorResponse if registration fails

If the email address is already in use, the method will return an ErrorResponse.
If the role "Admin_Admin_14" is not found in the database, the method will return an ErrorResponse.
If the random string for the verification code cannot be generated, the method will return an ErrorResponse.
If the user cannot be created, the method will return an ErrorResponse.
If the email cannot be sent, the method will return an ErrorResponse.
If the user role cannot be assigned, the method will return an ErrorResponse.

```go
func (s *registerService) Register(request *requests.RegisterRequest) (*response.UserResponse, *response.ErrorResponse)
```

##### `recordMetrics`

recordMetrics records Prometheus metrics for the given method and status.
It increments the request counter and observes the request duration
for the given method and status, using the provided start time.

```go
func (s *registerService) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

startTracingAndLogging starts a trace and logging for the given method.
It returns the trace span, a function to end the trace and log the status,
the initial status which is success, and a function to log a success message.
The end function should be called with the status of the operation.
The status can be either success or error.
The logSuccess function can be called with a message and fields to log a success message.
The message will be added as an event to the trace span and logged with the given fields.

```go
func (s *registerService) startTracingAndLogging(method string, attrs ...attribute.KeyValue) (trace.Span, func(string), string, func(string, ...zap.Field))
```

### `tokenService`

```go
type tokenService struct {
	ctx context.Context
	refreshToken repository.RefreshTokenRepository
	token auth.TokenManager
	logger logger.LoggerInterface
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `createAccessToken`

createAccessToken generates an access token for a given user ID.
It initiates tracing and logging for the token creation process.
The function returns the generated token as a string if successful,
or an error if the token generation fails. Tracing and logging are
used to record the success or failure of the operation.

```go
func (s *tokenService) createAccessToken(id int) (string, error)
```

##### `createRefreshToken`

createRefreshToken generates a refresh token for a given user ID.
It initiates tracing and logging for the token creation process.
The function deletes any existing refresh tokens for the user before
creating a new one.
The function returns the generated token as a string if successful,
or an error if the token generation fails. Tracing and logging are
used to record the success or failure of the operation.

```go
func (s *tokenService) createRefreshToken(id int) (string, error)
```

##### `recordMetrics`

recordMetrics records a Prometheus metric for the given method and status.
It increments a counter and records the duration since the provided start time.

```go
func (s *tokenService) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

startTracingAndLogging starts a tracing span and logging for a given method,
and returns the span, a function to end the span, the initial status of the
operation, and a function to log a message as a success event.

The returned end function takes a status string and updates the span's status
and ends the span. The status string can be one of "success" or "error".

The returned logSuccess function takes a message and any number of zap fields
and logs the message as a success event, and adds the event to the span.

The returned logError function takes a trace ID, a message, an error, and any
number of zap fields, and logs the message as an error event, and adds the event
to the span. The trace ID is used to identify the span in the error message.

```go
func (s *tokenService) startTracingAndLogging(method string, attrs ...attribute.KeyValue) (end func(string), logSuccess func(string, ...zap.Field), status string, logError func(traceID string, msg string, err error, fields ...zap.Field))
```

