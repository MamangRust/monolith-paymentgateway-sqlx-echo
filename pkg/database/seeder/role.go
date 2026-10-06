package seeder

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const seedCreateRole = `INSERT INTO
    roles (
        role_name,
        created_at,
        updated_at
    )
VALUES (
        $1,
        current_timestamp,
        current_timestamp
    )
RETURNING
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at
`

const seedGetRoleByName = `SELECT
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at
FROM roles
WHERE
    role_name = $1
`

// roleSeeder is a struct that represents a seeder for the roles table.
type roleSeeder struct {
	db     *sqlx.DB
	ctx    context.Context
	logger logger.LoggerInterface
}

// NewRoleSeeder creates a new instance of the roleSeeder, which is
// responsible for populating the roles table with fake data.
//
// Args:
// db: a pointer to the database pool
// ctx: a context.Context object
// logger: a logger.LoggerInterface object
//
// Returns:
// a pointer to the roleSeeder struct
func NewRoleSeeder(db *sqlx.DB, ctx context.Context, logger logger.LoggerInterface) *roleSeeder {
	return &roleSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

// Seed populates the roles table with fake data.
//
// It creates a total of 10 roles, with names randomly selected from the
// following list:
//
// - Super Admin
// - Admin
// - Merchant Admin
// - Merchant Operator
// - Finance
// - Compliance
// - Auditor
// - Support
// - Viewer
// - User
//
// If any of the roles fail to be created, the function returns an error.
//
// Returns:
// an error if any of the roles fail to be created, otherwise nil
func (r *roleSeeder) Seed() error {
	randomRoles := []string{
		"Super Admin",
		"Admin",
		"Merchant Admin",
		"Merchant Operator",
		"Finance",
		"Compliance",
		"Auditor",
		"Support",
		"Viewer",
		"User",
	}

	totalRoles := len(randomRoles)

	for i, roleName := range randomRoles {
		var role db.Role
		err := r.db.GetContext(r.ctx, &role, seedCreateRole, roleName)
		if err != nil {
			r.logger.Error("failed to seed role", zap.Int("role", i+1), zap.String("roleName", roleName), zap.Error(err))
			return fmt.Errorf("failed to seed role %d (%s): %w", i+1, roleName, err)
		}
	}

	r.logger.Debug("role seeded successfully", zap.Int("totalRoles", totalRoles))
	return nil
}
