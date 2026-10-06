package repository

import (
	"context"

	roleadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/role"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

// RoleRepository is the role read contract the auth service depends on. The
// role service owns the roles table, so reads go through the shared adapter.
type RoleRepository = roleadapter.RoleAdapter

// UserRoleRepository is the user_roles write contract. The role service is the
// only writer of user_roles; the auth service reaches it through the adapter.
type UserRoleRepository = roleadapter.UserRoleAdapter

// UserRepository defines the data access layer for user-related operations.
//
//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go
type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*db.GetUserByEmailRow, error)

	FindByEmailAndVerify(ctx context.Context, email string) (*db.GetUserByEmailAndVerifiedRow, error)

	FindById(ctx context.Context, user_id int) (*db.GetUserByIDRow, error)

	CreateUser(ctx context.Context, request *requests.RegisterRequest) (*db.CreateUserRow, error)

	UpdateUserIsVerified(ctx context.Context, user_id int, is_verified bool) (*db.UpdateUserIsVerifiedRow, error)

	UpdateUserPassword(ctx context.Context, user_id int, password string) (*db.UpdateUserPasswordRow, error)

	FindByVerificationCode(ctx context.Context, verification_code string) (*db.GetUserByVerificationCodeRow, error)
}

type ResetTokenRepository interface {
	FindByToken(ctx context.Context, code string) (*db.GetResetTokenRow, error)

	CreateResetToken(ctx context.Context, req *requests.CreateResetTokenRequest) (*db.CreateResetTokenRow, error)

	DeleteResetToken(ctx context.Context, user_id int) error
}

type RefreshTokenRepository interface {
	FindByToken(ctx context.Context, token string) (*db.RefreshToken, error)

	FindByUserId(ctx context.Context, user_id int) (*db.RefreshToken, error)

	CreateRefreshToken(ctx context.Context, req *requests.CreateRefreshToken) (*db.RefreshToken, error)

	UpdateRefreshToken(ctx context.Context, req *requests.UpdateRefreshToken) (*db.RefreshToken, error)

	DeleteRefreshToken(ctx context.Context, token string) error

	DeleteRefreshTokenByUserId(ctx context.Context, user_id int) error
}
