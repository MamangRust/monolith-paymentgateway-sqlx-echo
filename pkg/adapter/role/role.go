// Package role adapts the Role service gRPC API for services that must not read
// the roles/user_roles tables directly.
//
// Two narrow adapters sit over the split role clients: RoleAdapter for role
// reads (RoleServiceClient) and UserRoleAdapter for user_roles writes
// (RoleCommandServiceClient). The role service stays the only writer of
// user_roles; consumers reach it through these adapters.
package role

import (
	"context"
	"time"

	pbrole "github.com/MamangRust/monolith-payment-gateway-pb/role"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/resilience"
	role_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/role_errors/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

// RoleAdapter is the contract consumers depend on for role reads.
type RoleAdapter interface {
	FindById(ctx context.Context, roleID int) (*db.Role, error)
	FindByName(ctx context.Context, name string) (*db.Role, error)
}

type roleGRPCAdapter struct {
	queryClient pbrole.RoleServiceClient
	guard       *resilience.DependencyGuard
}

// NewRoleAdapter wraps the generated role query client. Resilience is applied
// per call through the optional dependency guard; a nil guard is a passthrough.
func NewRoleAdapter(queryClient pbrole.RoleServiceClient, opts ...adapter.GuardOption) RoleAdapter {
	a := &roleGRPCAdapter{queryClient: queryClient}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// SetGuard implements adapter.GuardSetter.
func (a *roleGRPCAdapter) SetGuard(g *resilience.DependencyGuard) { a.guard = g }

func (a *roleGRPCAdapter) FindById(ctx context.Context, roleID int) (*db.Role, error) {
	res, err := adapter.Call(a.guard, ctx, func(ctx context.Context) (*pbrole.ApiResponseRole, error) {
		return a.queryClient.FindByIdRole(ctx, &pbrole.FindByIdRoleRequest{
			RoleId: int32(roleID),
		})
	})
	if err != nil {
		return nil, role_errors.ErrRoleNotFound.WithInternal(err)
	}
	if res == nil || res.GetData() == nil {
		return nil, role_errors.ErrRoleNotFound
	}

	return roleToModel(res.GetData()), nil
}

func (a *roleGRPCAdapter) FindByName(ctx context.Context, name string) (*db.Role, error) {
	res, err := adapter.Call(a.guard, ctx, func(ctx context.Context) (*pbrole.ApiResponseRole, error) {
		return a.queryClient.FindByNameRole(ctx, &pbrole.FindByNameRoleRequest{
			Name: name,
		})
	})
	if err != nil {
		return nil, role_errors.ErrRoleNotFound.WithInternal(err)
	}
	if res == nil || res.GetData() == nil {
		return nil, role_errors.ErrRoleNotFound
	}

	return roleToModel(res.GetData()), nil
}

func roleToModel(role *pbrole.RoleResponse) *db.Role {
	if role == nil {
		return nil
	}

	return &db.Role{
		RoleID:    role.GetId(),
		RoleName:  role.GetName(),
		CreatedAt: parseTimestamp(role.GetCreatedAt()),
		UpdatedAt: parseTimestamp(role.GetUpdatedAt()),
	}
}

// parseTimestamp converts the role handlers' wire format back into the
// pgtype.Timestamp used by the SQLC models, so consumers can treat adapter
// results like local rows. Handlers emit a date ("2006-01-02",
// service/role/handler/{query,command}.go); the other layouts keep older or
// richer emitters working.
func parseTimestamp(value string) pgtype.Timestamp {
	if value == "" {
		return pgtype.Timestamp{}
	}

	for _, layout := range []string{"2006-01-02", "2006-01-02 15:04:05", time.RFC3339} {
		if parsed, err := time.Parse(layout, value); err == nil {
			return pgtype.Timestamp{Time: parsed, Valid: true}
		}
	}

	return pgtype.Timestamp{}
}
