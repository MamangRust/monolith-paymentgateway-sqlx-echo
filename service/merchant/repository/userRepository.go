package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	user_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/user_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getUserByID = `SELECT
    user_id,
    firstname,
    lastname,
    email,
    password,
    created_at,
    updated_at
FROM users
WHERE
    user_id = $1
    AND deleted_at IS NULL`

// userRepository is a struct that represents a repository for user operations.
type userRepository struct {
	db *sqlx.DB
}

// NewUserRepository creates a new instance of userRepository.
//
// Parameters:
//   - db: A pointer to the sqlx.DB object for executing database queries.
//
// Returns:
//   - A pointer to the newly created userRepository instance.
func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindById(ctx context.Context, user_id int) (*db.GetUserByIDRow, error) {
	var user db.GetUserByIDRow
	err := r.db.GetContext(ctx, &user, getUserByID, int32(user_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, user_errors.ErrUserNotFound.WithInternal(err)
		}

		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &user, nil
}
