package authhandler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	auth_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/auth"
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	authapimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/auth"

	errors "github.com/MamangRust/monolith-payment-gateway-shared/errors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authHandleApi struct {
	client     pb.AuthServiceClient
	logger     logger.LoggerInterface
	mapping    authapimapper.AuthResponseMapper
	apiHandler errors.ApiHandler
	cache      auth_cache.AuthMencache
}

type authHandleParams struct {
	client pb.AuthServiceClient

	router *echo.Echo

	cache auth_cache.AuthMencache

	logger logger.LoggerInterface

	mapper authapimapper.AuthResponseMapper

	apiHandler errors.ApiHandler
}

func NewHandlerAuth(params *authHandleParams) *authHandleApi {
	authHandler := &authHandleApi{
		client:     params.client,
		logger:     params.logger,
		mapping:    params.mapper,
		apiHandler: params.apiHandler,
		cache:      params.cache,
	}
	routerAuth := params.router.Group("/api/auth")

	routerAuth.GET("/hello", authHandler.HandleHello)
	routerAuth.POST("/register", params.apiHandler.Handle("register", authHandler.Register))
	routerAuth.POST("/login", params.apiHandler.Handle("login", authHandler.Login))
	routerAuth.POST("/refresh-token", params.apiHandler.Handle("register", authHandler.RefreshToken))
	routerAuth.GET("/me", params.apiHandler.Handle("GetMe", authHandler.GetMe))
	routerAuth.POST("/verify-code", params.apiHandler.Handle("verify-code", authHandler.VerifyCode))
	routerAuth.POST("/forgot-password", params.apiHandler.Handle("forgot-password", authHandler.ForgotPassword))
	routerAuth.POST("/reset-password", params.apiHandler.Handle("reset-password", authHandler.ResetPassword))

	return authHandler
}

// HandleHello godoc
// @Summary Returns a "Hello" message
// @Tags Auth
// @Description Returns a simple "Hello" message for testing purposes.
// @Produce json
// @Success 200 {string} string "Hello"
// @Router /api/auth/hello [get]
func (h *authHandleApi) HandleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

// Register godoc
// @Summary Register a new user
// @Tags Auth
// @Description Registers a new user with the provided details.
// @Accept json
// @Produce json
// @Param request body requests.CreateUserRequest true "User registration data"
// @Success 200 {object} response.ApiResponseRegister "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/auth/register [post]
func (h *authHandleApi) Register(c echo.Context) error {
	var body requests.CreateUserRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request format").WithInternal(err)
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	data := &pb.RegisterRequest{
		Firstname:       body.FirstName,
		Lastname:        body.LastName,
		Email:           body.Email,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
	}

	res, err := h.client.RegisterUser(c.Request().Context(), data)

	if err != nil {
		h.logger.Error("Registration failed", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	return c.JSON(http.StatusCreated, h.mapping.ToResponseRegister(res))
}

// Login godoc
// @Summary Authenticate a user
// @Tags Auth
// @Description Authenticates a user using the provided email and password.
// @Accept json
// @Produce json
// @Param request body requests.AuthRequest true "User login credentials"
// @Success 200 {object} response.ApiResponseLogin "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/auth/login [post]
func (h *authHandleApi) Login(c echo.Context) error {
	var body requests.AuthRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request format").WithInternal(err)
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	ctx := c.Request().Context()

	cachedResponse, found := h.cache.GetCachedLogin(ctx, body.Email)
	if found {
		h.logger.Debug("Returning login response from cache", zap.String("email", body.Email))
		return c.JSON(http.StatusOK, cachedResponse)
	}

	res, err := h.client.LoginUser(c.Request().Context(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		h.logger.Error("Login failed", zap.Error(err))

		if status.Code(err) == codes.Internal && strings.Contains(err.Error(), "empty token") {
			return errors.ParseGrpcError(err)
		}

		return errors.ParseGrpcError(err)
	}

	mappedResponse := h.mapping.ToResponseLogin(res)

	h.cache.SetCachedLogin(ctx, body.Email, mappedResponse)

	return c.JSON(http.StatusOK, mappedResponse)
}

// RefreshToken godoc
// @Summary Refresh access token
// @Tags Auth
// @Description Refreshes the access token using a valid refresh token.
// @Accept json
// @Produce json
// @Param request body requests.RefreshTokenRequest true "Refresh token data"
// @Success 200 {object} response.ApiResponseRefreshToken "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/auth/refresh-token [post]
func (h *authHandleApi) RefreshToken(c echo.Context) error {
	var body requests.RefreshTokenRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request format").WithInternal(err)
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	cachedResponse, found := h.cache.GetRefreshToken(c.Request().Context(), body.RefreshToken)
	if found {
		h.logger.Debug("Returning refresh token response from cache")
		return c.JSON(http.StatusOK, cachedResponse)
	}

	res, err := h.client.RefreshToken(c.Request().Context(), &pb.RefreshTokenRequest{
		RefreshToken: body.RefreshToken,
	})
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	mappedResponse := h.mapping.ToResponseRefreshToken(res)

	h.cache.SetRefreshToken(c.Request().Context(), body.RefreshToken, mappedResponse)

	return c.JSON(http.StatusOK, mappedResponse)
}

// VerifyCode godoc
// @Summary Verify a verification code
// @Tags Auth
// @Description Verifies a verification code sent to the user's email address.
// @Accept json
// @Produce json
// @Param request body requests.VerifyCodeRequest true "Verification code data"
// @Success 200 {object} response.ApiResponseVerifyCode "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/auth/verify-code [post]
func (h *authHandleApi) VerifyCode(c echo.Context) error {
	var body requests.VerifyCodeRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request format").WithInternal(err)
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	res, err := h.client.VerifyCode(c.Request().Context(), &pb.VerifyCodeRequest{
		Code: body.Code,
	})
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	return c.JSON(http.StatusOK, h.mapping.ToResponseVerifyCode(res))
}

// ForgotPassword godoc
// @Summary Forgot password
// @Tags Auth
// @Description Initiates the forgot password process by sending a verification code to the user's email.
// @Accept json
// @Produce json
// @Param request body requests.ForgotPasswordRequest true "User email"
// @Success 200 {object} response.ApiResponseForgotPassword "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/auth/forgot-password [post]
func (h *authHandleApi) ForgotPassword(c echo.Context) error {
	var body requests.ForgotPasswordRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request format").WithInternal(err)
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	res, err := h.client.ForgotPassword(c.Request().Context(), &pb.ForgotPasswordRequest{
		Email: body.Email,
	})
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	return c.JSON(http.StatusOK, h.mapping.ToResponseForgotPassword(res))
}

// ResetPassword godoc
// @Summary Reset password
// @Tags Auth
// @Description Resets the user's password using a reset token and new password.
// @Accept json
// @Produce json
// @Param request body requests.ResetPasswordRequest true "Reset password data"
// @Success 200 {object} response.ApiResponseResetPassword "Success"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/auth/reset-password [post]
func (h *authHandleApi) ResetPassword(c echo.Context) error {
	var body requests.ResetPasswordRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request format").WithInternal(err)
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	res, err := h.client.ResetPassword(c.Request().Context(), &pb.ResetPasswordRequest{
		ResetToken:      body.ResetToken,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
	})
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	return c.JSON(http.StatusOK, h.mapping.ToResponseResetPassword(res))
}

// GetMe godoc
// @Summary Get current user information
// @Tags Auth
// @Security Bearer
// @Description Retrieves the current user's information using a valid access token from the Authorization header.
// @Produce json
// @Security BearerToken
// @Success 200 {object} response.ApiResponseGetMe "Success"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /api/auth/me [get]
func (h *authHandleApi) GetMe(c echo.Context) error {
	userId, ok := c.Get("userId").(string)
	if !ok {
		return errors.NewBadRequestError("user not authenticated")
	}

	uid, err := strconv.ParseInt(userId, 10, 32)
	if err != nil {
		return errors.NewBadRequestError("invalid user ID format")
	}
	userID := int(uid)

	if cached, found := h.cache.GetCachedUserInfo(c.Request().Context(), userId); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.GetMe(
		c.Request().Context(),
		&pb.GetMeRequest{
			UserId: int32(userID),
		},
	)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	response := h.mapping.ToResponseGetMe(res)
	h.cache.SetCachedUserInfo(c.Request().Context(), userId, response)

	return c.JSON(http.StatusOK, response)
}

func (h *authHandleApi) parseValidationErrors(err error) []errors.ValidationError {
	var validationErrs []errors.ValidationError

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			validationErrs = append(validationErrs, errors.ValidationError{
				Field:   fe.Field(),
				Message: h.getValidationMessage(fe),
			})
		}
		return validationErrs
	}

	return []errors.ValidationError{
		{
			Field:   "general",
			Message: err.Error(),
		},
	}
}

func (h *authHandleApi) getValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Must be at least %s", fe.Param())
	case "max":
		return fmt.Sprintf("Must be at most %s", fe.Param())
	case "gte":
		return fmt.Sprintf("Must be greater than or equal to %s", fe.Param())
	case "lte":
		return fmt.Sprintf("Must be less than or equal to %s", fe.Param())
	case "oneof":
		return fmt.Sprintf("Must be one of: %s", fe.Param())
	default:
		return fmt.Sprintf("Validation failed on '%s' tag", fe.Tag())
	}
}
