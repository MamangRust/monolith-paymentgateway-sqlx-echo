package seeder

import (
	"context"
	"fmt"
	"math/rand"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/hash"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const seedCreateUser = `WITH inserted_user AS (
    INSERT INTO users (
        firstname,
        lastname,
        email,
        password,
        verification_code,
        is_verified,
        created_at,
        updated_at
    )
    VALUES ($1, $2, $3, $4, $5, $6, current_timestamp, current_timestamp)
    RETURNING user_id, firstname, lastname, email, password, verification_code, created_at, updated_at
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'auth-register:' || user_id::text,
        'email-service-topic-auth-register',
        user_id::text,
        jsonb_build_object(
            'event_id', gen_random_uuid()::text,
            'schema_version', 1,
            'event_type', 'auth.register',
            'occurred_at', now(),
            'email', email,
            'subject', 'Welcome to SanEdge',
            'body', 'Your account has been successfully created. Verify your account here: https://sanedge.example.com/login?verify_code=' || verification_code
        )
    FROM inserted_user
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT user_id, firstname, lastname, email, password, created_at, updated_at
FROM inserted_user
`

const seedTrashUser = `UPDATE users
SET
    deleted_at = current_timestamp,
    updated_at = current_timestamp
WHERE
    user_id = $1
    AND deleted_at IS NULL
RETURNING
    user_id,
    firstname,
    lastname,
    email,
    created_at,
    updated_at,
    deleted_at
`

const seedAssignRoleToUser = `INSERT INTO user_roles (
    user_id,
    role_id,
    created_at,
    updated_at
) VALUES (
    $1,
    $2,
    current_timestamp,
    current_timestamp
) RETURNING
    user_role_id,
    user_id,
    role_id,
    created_at,
    updated_at,
    deleted_at
`

// userSeeder is a struct that represents a seeder for the users table.
type userSeeder struct {
	db     *sqlx.DB
	hash   hash.HashPassword
	ctx    context.Context
	logger logger.LoggerInterface
}

// NewUserSeeder creates a new instance of the userSeeder, which is
// responsible for populating the users table with fake data.
//
// Args:
// db: a pointer to the database pool
// ctx: a context.Context object
// hash: a hash.HashPassword object
// logger: a logger.LoggerInterface object
//
// Returns:
// a pointer to the userSeeder struct
func NewUserSeeder(db *sqlx.DB, ctx context.Context, hash hash.HashPassword, logger logger.LoggerInterface) *userSeeder {
	return &userSeeder{
		db:     db,
		hash:   hash,
		ctx:    ctx,
		logger: logger,
	}
}

// Seed populates the users table with fake data.
//
// It creates a total of 15 users, with 10 of them being active and 5 of them being
// trashed. The active users have a status of "active", while the trashed users have a
// status of "deactive". The user names are generated randomly from the combination of
// the adjectives and nouns provided. The API key is generated using the GenerateApiKey
// function from the api-key package.
//
// If any errors occur during the seeding process, an error is returned.
//
// Returns:
// an error if any of the users fail to be created, otherwise nil
func (r *userSeeder) Seed() error {
	totalUsers := 15
	activeUsers := 10
	trashedUsers := totalUsers - activeUsers

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

	for i := 1; i <= totalUsers; i++ {
		email := fmt.Sprintf("user_%s@example.com", uuid.NewString())

		hashedPassword, err := r.hash.HashPassword("password")
		if err != nil {
			r.logger.Error("failed to hash password", zap.Int("user", i), zap.Error(err))
			return fmt.Errorf("failed to hash password for user %d: %w", i, err)
		}

		isVerified := true

		user := db.CreateUserParams{
			Firstname:        fmt.Sprintf("User%d", i),
			Lastname:         fmt.Sprintf("Last%d", i),
			Email:            email,
			Password:         hashedPassword,
			VerificationCode: uuid.NewString(),
			IsVerified:       &isVerified,
		}

		var createdUser db.CreateUserRow
		err = r.db.GetContext(r.ctx, &createdUser, seedCreateUser,
			user.Firstname,
			user.Lastname,
			user.Email,
			user.Password,
			user.VerificationCode,
			user.IsVerified,
		)
		if err != nil {
			r.logger.Error("failed to create user", zap.Int("user", i), zap.Error(err))
			return fmt.Errorf("failed to create user %d: %w", i, err)
		}

		randomRole := randomRoles[rand.Intn(len(randomRoles))]
		var role db.Role
		err = r.db.GetContext(r.ctx, &role, seedGetRoleByName, randomRole)
		if err != nil {
			r.logger.Error("failed to get role", zap.String("role", randomRole), zap.Error(err))
			return fmt.Errorf("failed to get role %s: %w", randomRole, err)
		}

		var userRole db.UserRole
		err = r.db.GetContext(r.ctx, &userRole, seedAssignRoleToUser, createdUser.UserID, role.RoleID)
		if err != nil {
			r.logger.Error("failed to assign role to user", zap.Int("userID", int(createdUser.UserID)), zap.String("role", randomRole), zap.Error(err))
			return fmt.Errorf("failed to assign role %s to user %d: %w", randomRole, createdUser.UserID, err)
		}

		if i > activeUsers {
			var trashed db.TrashUserRow
			err := r.db.GetContext(r.ctx, &trashed, seedTrashUser, createdUser.UserID)
			if err != nil {
				r.logger.Error("failed to trash user", zap.Int("user", i), zap.Error(err))
				return fmt.Errorf("failed to trash user %d: %w", i, err)
			}
		}
	}

	r.logger.Info("user seeded successfully", zap.Int("totalUsers", totalUsers), zap.Int("activeUsers", activeUsers), zap.Int("trashedUsers", trashedUsers))

	return nil
}
