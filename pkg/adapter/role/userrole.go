package role

import (
	"context"

	pbrole "github.com/MamangRust/monolith-payment-gateway-pb/role"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/resilience"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	userrole_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/user_role_errors/repository"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UserRoleAdapter is the write path consumers use to (un)assign roles. The role
// service owns the user_roles table, so every write goes through these RPCs.
type UserRoleAdapter interface {
	AssignRoleToUser(ctx context.Context, req *requests.CreateUserRoleRequest) (*db.UserRole, error)
	RemoveRoleFromUser(ctx context.Context, req *requests.RemoveUserRoleRequest) error
}

type userRoleGRPCAdapter struct {
	commandClient pbrole.RoleCommandServiceClient
	guard         *resilience.DependencyGuard
}

// NewUserRoleAdapter wraps the generated role command client. Resilience is
// applied per call through the optional dependency guard; a nil guard is a
// passthrough.
func NewUserRoleAdapter(commandClient pbrole.RoleCommandServiceClient, opts ...adapter.GuardOption) UserRoleAdapter {
	a := &userRoleGRPCAdapter{commandClient: commandClient}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// SetGuard implements adapter.GuardSetter.
func (a *userRoleGRPCAdapter) SetGuard(g *resilience.DependencyGuard) { a.guard = g }

// AssignRoleToUser implements UserRoleAdapter.
func (a *userRoleGRPCAdapter) AssignRoleToUser(ctx context.Context, req *requests.CreateUserRoleRequest) (*db.UserRole, error) {
	res, err := adapter.Call(a.guard, ctx, func(ctx context.Context) (*pbrole.ApiResponseUserRole, error) {
		return a.commandClient.AssignRoleToUser(ctx, &pbrole.AssignRoleToUserRequest{
			UserId: int32(req.UserId),
			RoleId: int32(req.RoleId),
		})
	})
	if err != nil || res == nil || res.GetData() == nil {
		return nil, userrole_errors.ErrAssignRoleToUser.WithInternal(err)
	}

	return &db.UserRole{
		UserRoleID: res.GetData().GetUserRoleId(),
		UserID:     res.GetData().GetUserId(),
		RoleID:     res.GetData().GetRoleId(),
	}, nil
}

// RemoveRoleFromUser implements UserRoleAdapter.
func (a *userRoleGRPCAdapter) RemoveRoleFromUser(ctx context.Context, req *requests.RemoveUserRoleRequest) error {
	_, err := adapter.Call(a.guard, ctx, func(ctx context.Context) (*emptypb.Empty, error) {
		return a.commandClient.RemoveRoleFromUser(ctx, &pbrole.RemoveRoleFromUserRequest{
			UserId: int32(req.UserId),
			RoleId: int32(req.RoleId),
		})
	})
	if err != nil {
		return userrole_errors.ErrRemoveRole.WithInternal(err)
	}

	return nil
}
