package repository

import (
	"context"

	useradapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/user"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	user_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/user_errors/repository"
)

// userRepository validates card ownership against the canonical user store via
// the shared user gRPC adapter instead of a local copy of the users table.
type userRepository struct {
	adapter useradapter.QueryRepository
}

func NewUserRepository(a useradapter.QueryRepository) UserRepository {
	return &userRepository{adapter: a}
}

func (r *userRepository) FindById(ctx context.Context, user_id int) (*db.GetUserByIDRow, error) {
	res, err := r.adapter.FindById(ctx, user_id)
	if err != nil {
		return nil, user_errors.ErrUserNotFound.WithInternal(err)
	}

	return res, nil
}
