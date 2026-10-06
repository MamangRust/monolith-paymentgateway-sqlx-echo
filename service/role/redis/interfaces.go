package mencache

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type RoleQueryCache interface {
	SetCachedRoles(ctx context.Context, req *requests.FindAllRoles, data []*db.GetRolesRow, total *int)
	GetCachedRoles(ctx context.Context, req *requests.FindAllRoles) ([]*db.GetRolesRow, *int, bool)

	GetCachedRoleById(ctx context.Context, id int) (*db.Role, bool)
	SetCachedRoleById(ctx context.Context, id int, data *db.Role)

	GetCachedRoleByUserId(ctx context.Context, userId int) ([]*db.Role, bool)
	SetCachedRoleByUserId(ctx context.Context, userId int, data []*db.Role)

	GetCachedRoleActive(ctx context.Context, req *requests.FindAllRoles) ([]*db.GetActiveRolesRow, *int, bool)
	SetCachedRoleActive(ctx context.Context, req *requests.FindAllRoles, data []*db.GetActiveRolesRow, total *int)

	GetCachedRoleTrashed(ctx context.Context, req *requests.FindAllRoles) ([]*db.GetTrashedRolesRow, *int, bool)
	SetCachedRoleTrashed(ctx context.Context, req *requests.FindAllRoles, data []*db.GetTrashedRolesRow, total *int)
}

type RoleCommandCache interface {
	DeleteCachedRole(ctx context.Context, id int)
}
