package repository

import (
	pbrole "github.com/MamangRust/monolith-payment-gateway-pb/role"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	roleadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/role"
	"github.com/jmoiron/sqlx"
)

type GuardOptions struct {
	Role     []adapter.GuardOption
	UserRole []adapter.GuardOption
}

// Repositories embeds the user query/command repositories and exposes the role
// adapters through accessors. roleadapter.RoleAdapter declares FindById, which
// would collide with UserQueryRepository.FindById if it were embedded directly.
type Repositories interface {
	UserQuery() UserQueryRepository
	UserCommand() UserCommandRepository
	RoleAdapter() roleadapter.RoleAdapter
	UserRoleAdapter() roleadapter.UserRoleAdapter
}

type repositories struct {
	userQuery       UserQueryRepository
	userCommand     UserCommandRepository
	roleAdapter     roleadapter.RoleAdapter
	userRoleAdapter roleadapter.UserRoleAdapter
}

func NewRepositories(
	db *sqlx.DB,
	roleQueryClient pbrole.RoleServiceClient,
	roleCommandClient pbrole.RoleCommandServiceClient,
	guards ...GuardOptions,
) Repositories {
	var g GuardOptions

	if len(guards) > 0 {
		g = guards[0]
	}

	return &repositories{
		userQuery:       NewUserQueryRepository(db),
		userCommand:     NewUserCommandRepository(db),
		roleAdapter:     roleadapter.NewRoleAdapter(roleQueryClient, g.Role...),
		userRoleAdapter: roleadapter.NewUserRoleAdapter(roleCommandClient, g.UserRole...),
	}
}

func (r *repositories) UserQuery() UserQueryRepository {
	return r.userQuery
}

func (r *repositories) UserCommand() UserCommandRepository {
	return r.userCommand
}

func (r *repositories) RoleAdapter() roleadapter.RoleAdapter {
	return r.roleAdapter
}

func (r *repositories) UserRoleAdapter() roleadapter.UserRoleAdapter {
	return r.userRoleAdapter
}
