package repository

import (
	pbrole "github.com/MamangRust/monolith-payment-gateway-pb/role"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	roleadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/role"
	"github.com/jmoiron/sqlx"
)

// Repositories bundles all auth repositories. It uses named fields (not
// embedding) because UserRepository and RoleRepository both declare FindById,
// and RefreshTokenRepository and ResetTokenRepository both declare FindByToken.
type Repositories struct {
	User         UserRepository
	RefreshToken RefreshTokenRepository
	UserRole     roleadapter.UserRoleAdapter
	Role         roleadapter.RoleAdapter
	ResetToken   ResetTokenRepository
}

// GuardOptions collects the dependency guards for each remote dependency. Role
// reads and user_role writes get separate breakers so one failure mode cannot
// starve the other.
type GuardOptions struct {
	Role     []adapter.GuardOption
	UserRole []adapter.GuardOption
}

type Deps struct {
	Db                *sqlx.DB
	RoleQueryClient   pbrole.RoleServiceClient
	RoleCommandClient pbrole.RoleCommandServiceClient
	Guard             GuardOptions
}

func NewRepositories(
	deps *Deps,
) *Repositories {
	return &Repositories{
		User:         NewUserRepository(deps.Db),
		RefreshToken: NewRefreshTokenRepository(deps.Db),
		UserRole:     roleadapter.NewUserRoleAdapter(deps.RoleCommandClient, deps.Guard.UserRole...),
		Role:         roleadapter.NewRoleAdapter(deps.RoleQueryClient, deps.Guard.Role...),
		ResetToken:   NewResetTokenRepository(deps.Db),
	}
}
