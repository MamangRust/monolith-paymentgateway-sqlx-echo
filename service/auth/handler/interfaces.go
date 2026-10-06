package handler

import (
	"context"

	pb "github.com/MamangRust/monolith-payment-gateway-pb"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/handler.go
type AuthHandleGrpc interface {
	pb.AuthServiceServer

	LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.ApiResponseLogin, error)

	RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.ApiResponseRegister, error)

	ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ApiResponseForgotPassword, error)

	VerifyCode(ctx context.Context, req *pb.VerifyCodeRequest) (*pb.ApiResponseVerifyCode, error)

	ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ApiResponseResetPassword, error)

	GetMe(ctx context.Context, req *pb.GetMeRequest) (*pb.ApiResponseGetMe, error)

	RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.ApiResponseRefreshToken, error)
}
