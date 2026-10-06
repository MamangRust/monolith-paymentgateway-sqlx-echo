// Package user adapts the User service gRPC API into the shared
// pkg/database sqlc row types used by the consumer repositories.
package user

import (
	"context"
	"time"

	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/resilience"
	user_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/user_errors/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

// QueryRepository is the contract consumers depend on for user reads.
type QueryRepository interface {
	FindById(ctx context.Context, userID int) (*db.GetUserByIDRow, error)
}

type userGRPCAdapter struct {
	queryClient pbuser.UserQueryServiceClient
	guard       *resilience.DependencyGuard
}

// NewQueryAdapter builds a read-only user adapter for consumers that resolve a
// user remotely instead of touching the users table. Resilience is applied per
// call through the optional dependency guard; a nil guard is a passthrough.
func NewQueryAdapter(queryClient pbuser.UserQueryServiceClient, opts ...adapter.GuardOption) QueryRepository {
	a := &userGRPCAdapter{queryClient: queryClient}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// SetGuard implements adapter.GuardSetter.
func (a *userGRPCAdapter) SetGuard(g *resilience.DependencyGuard) { a.guard = g }

func (a *userGRPCAdapter) FindById(ctx context.Context, userID int) (*db.GetUserByIDRow, error) {
	res, err := adapter.Call(a.guard, ctx, func(ctx context.Context) (*pbuser.ApiResponseUser, error) {
		return a.queryClient.FindById(ctx, &pbuser.FindByIdUserRequest{Id: int32(userID)})
	})
	if err != nil {
		return nil, user_errors.ErrUserNotFound.WithInternal(err)
	}
	if res == nil || res.GetData() == nil {
		return nil, user_errors.ErrUserNotFound
	}

	return &db.GetUserByIDRow{
		UserID:    res.GetData().GetId(),
		Firstname: res.GetData().GetFirstname(),
		Lastname:  res.GetData().GetLastname(),
		Email:     res.GetData().GetEmail(),
		CreatedAt: parseTimestamp(res.GetData().GetCreatedAt()),
		UpdatedAt: parseTimestamp(res.GetData().GetUpdatedAt()),
	}, nil
}

// parseTimestamp converts the user handlers' wire format back into the
// pgtype.Timestamp used by the SQLC models.
func parseTimestamp(value string) pgtype.Timestamp {
	if value == "" {
		return pgtype.Timestamp{}
	}

	for _, layout := range []string{time.RFC3339, "2006-01-02 15:04:05", "2006-01-02"} {
		if parsed, err := time.Parse(layout, value); err == nil {
			return pgtype.Timestamp{Time: parsed, Valid: true}
		}
	}

	return pgtype.Timestamp{}
}
