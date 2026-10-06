# 📦 Package `handler`

**Source Path:** `service/auth/internal/handler`

## 🧩 Types

### `AuthHandleGrpc`

AuthHandleGrpc defines the gRPC service interface for authentication operations.
It combines the generated gRPC server interface with concrete method signatures
for all authentication-related operations.

This interface serves as the contract for the gRPC authentication service handler,
providing methods for user authentication, registration, password management,
and token operations.
```go
type AuthHandleGrpc interface {
	pb.AuthServiceServer
	LoginUser func(ctx context.Context, req *pb.LoginRequest) (*pb.ApiResponseLogin, error)
	RegisterUser func(ctx context.Context, req *pb.RegisterRequest) (*pb.ApiResponseRegister, error)
	ForgotPassword func(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ApiResponseForgotPassword, error)
	VerifyCode func(ctx context.Context, req *pb.VerifyCodeRequest) (*pb.ApiResponseVerifyCode, error)
	ResetPassword func(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ApiResponseResetPassword, error)
	GetMe func(ctx context.Context, req *pb.GetMeRequest) (*pb.ApiResponseGetMe, error)
	RefreshToken func(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.ApiResponseRefreshToken, error)
}
```

### `Deps`

```go
type Deps struct {
	Service *service.Service
	Logger logger.LoggerInterface
}
```

### `Handler`

```go
type Handler struct {
	Auth AuthHandleGrpc
}
```

### `authHandleGrpc`

```go
type authHandleGrpc struct {
	pb.UnimplementedAuthServiceServer
	registerService service.RegistrationService
	loginService service.LoginService
	passwordResetService service.PasswordResetService
	identifyService service.IdentifyService
	logger logger.LoggerInterface
	mapping protomapper.AuthProtoMapper
}
```

#### Methods

##### `ForgotPassword`

ForgotPassword initiates the password reset process.
Takes a ForgotPasswordRequest with user email and returns:
- ApiResponseForgotPassword with process status
- error if the request fails

```go
func (s *authHandleGrpc) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ApiResponseForgotPassword, error)
```

##### `GetMe`

GetMe retrieves the current authenticated user's details.
Takes a GetMeRequest with access token and returns:
- ApiResponseGetMe with user information
- error if the token is invalid

```go
func (s *authHandleGrpc) GetMe(ctx context.Context, req *pb.GetMeRequest) (*pb.ApiResponseGetMe, error)
```

##### `LoginUser`

LoginUser authenticates a user and returns access tokens.
Takes a LoginRequest with user credentials and returns:
- ApiResponseLogin containing tokens if successful
- error if authentication fails

```go
func (s *authHandleGrpc) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.ApiResponseLogin, error)
```

##### `RefreshToken`

RefreshToken generates new access tokens using a refresh token.
Takes a RefreshTokenRequest with refresh token and returns:
- ApiResponseRefreshToken with new tokens
- error if token refresh fails

```go
func (s *authHandleGrpc) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.ApiResponseRefreshToken, error)
```

##### `RegisterUser`

RegisterUser creates a new user account.
Takes a RegisterRequest with user details and returns:
- ApiResponseRegister with registration status
- error if registration fails

```go
func (s *authHandleGrpc) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.ApiResponseRegister, error)
```

##### `ResetPassword`

ResetPassword completes the password reset process.
Takes a ResetPasswordRequest with new credentials and returns:
- ApiResponseResetPassword with reset status
- error if reset fails

```go
func (s *authHandleGrpc) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ApiResponseResetPassword, error)
```

##### `VerifyCode`

VerifyCode validates a password reset verification code.
Takes a VerifyCodeRequest with the code and returns:
- ApiResponseVerifyCode with validation result
- error if verification fails

```go
func (s *authHandleGrpc) VerifyCode(ctx context.Context, req *pb.VerifyCodeRequest) (*pb.ApiResponseVerifyCode, error)
```

## 🚀 Functions

### `NewAuthHandleGrpc`

```go
func NewAuthHandleGrpc(authService *service.Service, logger logger.LoggerInterface) pb.AuthServiceServer
```

