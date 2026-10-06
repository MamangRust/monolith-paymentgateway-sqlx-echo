package authapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type AuthResponseMapper interface {
	// ToResponseVerifyCode maps gRPC response to API response for verify code operation.
	ToResponseVerifyCode(res *pb.ApiResponseVerifyCode) *response.ApiResponseVerifyCode

	// ToResponseForgotPassword maps gRPC response to API response for forgot password.
	ToResponseForgotPassword(res *pb.ApiResponseForgotPassword) *response.ApiResponseForgotPassword

	// ToResponseResetPassword maps gRPC response to API response for reset password.
	ToResponseResetPassword(res *pb.ApiResponseResetPassword) *response.ApiResponseResetPassword

	// ToResponseLogin maps gRPC response to API response for login.
	ToResponseLogin(res *pb.ApiResponseLogin) *response.ApiResponseLogin

	// ToResponseRegister maps gRPC response to API response for register.
	ToResponseRegister(res *pb.ApiResponseRegister) *response.ApiResponseRegister

	// ToResponseRefreshToken maps gRPC response to API response for refresh token.
	ToResponseRefreshToken(res *pb.ApiResponseRefreshToken) *response.ApiResponseRefreshToken

	// ToResponseGetMe maps gRPC response to API response for get user info.
	ToResponseGetMe(res *pb.ApiResponseGetMe) *response.ApiResponseGetMe
}
