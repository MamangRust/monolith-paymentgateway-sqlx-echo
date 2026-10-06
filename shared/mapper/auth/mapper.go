package authapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// authResponseMapper provides methods to map the auth response from the domain
// to the api response.
type authResponseMapper struct {
}

// NewAuthResponseMapper returns a new instance of authResponseMapper.
func NewAuthResponseMapper() AuthResponseMapper {
	return &authResponseMapper{}
}

// ToResponseVerifyCode maps the gRPC ApiResponseVerifyCode from the domain layer
// to the HTTP-compatible ApiResponseVerifyCode used in the API layer.
//
// Args:
//   - res: A pointer to a pb.ApiResponseVerifyCode containing the status and message
//     from the domain layer.
//
// Returns:
//   - A pointer to a response.ApiResponseVerifyCode with the mapped status and message
//     for the API layer.
func (s *authResponseMapper) ToResponseVerifyCode(res *pb.ApiResponseVerifyCode) *response.ApiResponseVerifyCode {
	return &response.ApiResponseVerifyCode{
		Status:  res.Status,
		Message: res.Message,
	}
}

// ToResponseForgotPassword maps the ApiResponseForgotPassword from the domain to the
// ApiResponseForgotPassword of the api.
//
// Args:
//   - res: A pointer to a pb.ApiResponseForgotPassword containing the status and message
//     from the domain layer.
//
// Returns:
//   - A pointer to a response.ApiResponseForgotPassword with the mapped status and message
//     for the API layer.
func (s *authResponseMapper) ToResponseForgotPassword(res *pb.ApiResponseForgotPassword) *response.ApiResponseForgotPassword {
	return &response.ApiResponseForgotPassword{
		Status:  res.Status,
		Message: res.Message,
	}
}

// ToResponseResetPassword maps the gRPC ApiResponseResetPassword from the domain layer
// to the HTTP-compatible ApiResponseResetPassword used in the API layer.
//
// Args:
//   - res: A pointer to a pb.ApiResponseResetPassword containing the status and message
//     from the domain layer.
//
// Returns:
//   - A pointer to a response.ApiResponseResetPassword with the mapped status and message
//     for the API layer.
func (s *authResponseMapper) ToResponseResetPassword(res *pb.ApiResponseResetPassword) *response.ApiResponseResetPassword {
	return &response.ApiResponseResetPassword{
		Status:  res.Status,
		Message: res.Message,
	}
}

// ToResponseLogin maps the ApiResponseLogin from the domain to the ApiResponseLogin of the api.
//
// Args:
//   - res: A pointer to a pb.ApiResponseLogin representing the domain response.
//
// Returns:
//   - A pointer to a response.ApiResponseLogin containing the mapped data, including status,
//     message, and token data such as access and refresh tokens.
func (s *authResponseMapper) ToResponseLogin(res *pb.ApiResponseLogin) *response.ApiResponseLogin {
	return &response.ApiResponseLogin{
		Status:  res.Status,
		Message: res.Message,
		Data: &response.TokenResponse{
			AccessToken:  res.Data.AccessToken,
			RefreshToken: res.Data.RefreshToken,
		},
	}
}

// ToResponseRegister maps the ApiResponseRegister from the domain to the
// ApiResponseRegister of the api.
//
// Args:
//   - res: A pointer to a pb.ApiResponseRegister representing the ApiResponseRegister from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseRegister containing the mapped data, including status, message,
//     and user data such as ID, FirstName, LastName, Email, CreatedAt, and UpdatedAt.
func (s *authResponseMapper) ToResponseRegister(res *pb.ApiResponseRegister) *response.ApiResponseRegister {
	return &response.ApiResponseRegister{
		Status:  res.Status,
		Message: res.Message,
		Data: &response.UserResponse{
			ID:        int(res.Data.Id),
			FirstName: res.Data.Firstname,
			LastName:  res.Data.Lastname,
			Email:     res.Data.Email,
			CreatedAt: res.Data.CreatedAt,
			UpdatedAt: res.Data.UpdatedAt,
		},
	}
}

// ToResponseRefreshToken maps the ApiResponseRefreshToken from the domain to the
// ApiResponseRefreshToken of the api.
//
// Args:
//   - res: A pointer to a pb.ApiResponseRefreshToken representing the ApiResponseRefreshToken from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseRefreshToken containing the mapped data, including status, message,
//     and token data such as access and refresh tokens.
func (s *authResponseMapper) ToResponseRefreshToken(res *pb.ApiResponseRefreshToken) *response.ApiResponseRefreshToken {
	return &response.ApiResponseRefreshToken{
		Status:  res.Status,
		Message: res.Message,
		Data: &response.TokenResponse{
			AccessToken:  res.Data.AccessToken,
			RefreshToken: res.Data.RefreshToken,
		},
	}
}

// ToResponseGetMe maps the ApiResponseGetMe from the domain to the ApiResponseGetMe of the api.
//
// Args:
//   - res: A pointer to a pb.ApiResponseGetMe representing the ApiResponseGetMe from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseGetMe containing the mapped data, including status, message,
//     and user data such as ID, FirstName, LastName, Email, CreatedAt, and UpdatedAt.
func (s *authResponseMapper) ToResponseGetMe(res *pb.ApiResponseGetMe) *response.ApiResponseGetMe {
	return &response.ApiResponseGetMe{
		Status:  res.Status,
		Message: res.Message,
		Data: &response.UserResponse{
			ID:        int(res.Data.Id),
			FirstName: res.Data.Firstname,
			LastName:  res.Data.Lastname,
			Email:     res.Data.Email,
			CreatedAt: res.Data.CreatedAt,
			UpdatedAt: res.Data.UpdatedAt,
		},
	}
}
