package seeder

import (
	"context"
	"fmt"

	apikey "github.com/MamangRust/monolith-payment-gateway-pkg/api-key"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const seedCreateMerchant = `WITH valid_user AS (
    SELECT user_id, email
    FROM users
    WHERE users.user_id = $3 AND users.deleted_at IS NULL
), inserted_merchant AS (
    INSERT INTO merchants (name, api_key, user_id, status, created_at, updated_at)
    SELECT $1, $2, valid_user.user_id, $4, current_timestamp, current_timestamp
    FROM valid_user
    RETURNING merchant_id, merchant_no, name, api_key, user_id, status, created_at, updated_at
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'merchant-create:' || merchant_id::text,
        'email-service-topic-merchant-create',
        merchant_id::text,
        jsonb_build_object(
            'event_id', gen_random_uuid()::text,
            'schema_version', 1,
            'event_type', 'merchant.created',
            'occurred_at', now(),
            'email', u.email,
            'subject', 'Initial Verification - SanEdge',
            'body', 'Your merchant account has been created successfully. Please upload the required documents for verification.'
        )
    FROM inserted_merchant m
    JOIN valid_user u ON u.user_id = m.user_id
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT merchant_id, merchant_no, name, api_key, user_id, status, created_at, updated_at
FROM inserted_merchant
`

const seedUpdateMerchantStatus = `WITH valid_owner AS (
    SELECT m.merchant_id
    FROM merchants m
    JOIN users u ON u.user_id = m.user_id AND u.deleted_at IS NULL
    WHERE m.merchant_id = $1 AND m.deleted_at IS NULL
), updated_merchant AS (
    UPDATE merchants
    SET status = $2, event_version = event_version + 1, updated_at = current_timestamp
    WHERE merchant_id IN (SELECT merchant_id FROM valid_owner)
    RETURNING merchant_id, merchant_no, name, api_key, user_id, status, created_at, updated_at, event_version
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'merchant-status:' || merchant_id::text || ':' || event_version::text,
        'email-service-topic-merchant-update-status',
        merchant_id::text,
        jsonb_build_object(
            'event_id', gen_random_uuid()::text,
            'schema_version', 1,
            'event_type', 'merchant.status_updated',
            'occurred_at', now(),
            'email', u.email,
            'subject', CASE status
                WHEN 'active' THEN 'Your Merchant Account is Now Active'
                WHEN 'inactive' THEN 'Merchant Account Set to Inactive'
                WHEN 'rejected' THEN 'Merchant Account Rejected'
                ELSE 'Merchant Account Status Updated'
            END,
            'body', 'Your merchant account status has been updated to ' || status || '.',
            'event_version', event_version
        )
    FROM updated_merchant m
    JOIN users u ON u.user_id = m.user_id AND u.deleted_at IS NULL
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT merchant_id, merchant_no, name, api_key, user_id, status, created_at, updated_at
FROM updated_merchant
`

const seedTrashMerchant = `UPDATE merchants
SET
    deleted_at = current_timestamp
WHERE
    merchant_id = $1
    AND deleted_at IS NULL
RETURNING
    merchant_id,
    merchant_no,
    name,
    api_key,
    user_id,
    status,
    created_at,
    updated_at,
    deleted_at
`

// merchantSeeder is a struct that represents a seeder for the merchants table.
type merchantSeeder struct {
	db     *sqlx.DB
	ctx    context.Context
	logger logger.LoggerInterface
}

// NewMerchantSeeder creates a new instance of the merchantSeeder, which is
// responsible for populating the merchants table with fake data.
//
// Args:
// db: a pointer to the database pool
// ctx: a context.Context object
// logger: a logger.LoggerInterface object
//
// Returns:
// a pointer to the merchantSeeder struct
func NewMerchantSeeder(db *sqlx.DB, ctx context.Context, logger logger.LoggerInterface) *merchantSeeder {
	return &merchantSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

// Seed populates the merchants table with fake data.
//
// It creates a total of 10 merchants, with 5 of them being active and 5 of them being
// trashed. The active merchants have a status of "active", while the trashed merchants
// have a status of "deactive". The merchant names are generated randomly from the
// combination of the adjectives and nouns provided. The API key is generated using
// the GenerateApiKey function from the api-key package.
func (r *merchantSeeder) Seed() error {
	adjectives := []string{"Blue", "Green", "Red", "Yellow", "Fast"}
	nouns := []string{"Shop", "Store", "Mart", "Market", "Hub"}

	totalMerchants := 10
	activeMerchants := 5
	trashedMerchants := 5

	for i := 0; i < totalMerchants; i++ {
		adjective := adjectives[i%len(adjectives)]
		noun := nouns[i%len(nouns)]
		merchantName := fmt.Sprintf("%s %s", adjective, noun)

		apiKey, _ := apikey.GenerateApiKey()

		req := db.CreateMerchantParams{
			Name:   merchantName,
			UserID: int32((i % 5) + 1),
			ApiKey: apiKey,
		}

		var merchant db.CreateMerchantRow
		err := r.db.GetContext(r.ctx, &merchant, seedCreateMerchant,
			req.Name,
			req.ApiKey,
			req.UserID,
			req.Status,
		)
		if err != nil {
			r.logger.Error("failed to seed merchant", zap.Int("merchant", i+1), zap.Error(err))
			return fmt.Errorf("failed to seed merchant %d: %w", i+1, err)
		}

		var status string
		if i < activeMerchants {
			status = "active"
		} else {
			status = "deactive"
		}

		var updated db.UpdateMerchantStatusRow
		err = r.db.GetContext(r.ctx, &updated, seedUpdateMerchantStatus, merchant.MerchantID, status)
		if err != nil {
			r.logger.Error("failed to update merchant status", zap.Int("merchantID", int(merchant.MerchantID)), zap.String("status", status), zap.Error(err))
			return fmt.Errorf("failed to update status for merchant ID %d: %w", merchant.MerchantID, err)
		}

		if i >= activeMerchants {
			var trashed db.TrashMerchantRow
			err = r.db.GetContext(r.ctx, &trashed, seedTrashMerchant, merchant.MerchantID)
			if err != nil {
				r.logger.Error("failed to trash merchant", zap.Int("merchant", i+1), zap.Error(err))
				return fmt.Errorf("failed to trash merchant %d: %w", i+1, err)
			}
		}
	}

	r.logger.Info("merchant seeded successfully",
		zap.Int("totalMerchants", totalMerchants),
		zap.Int("activeMerchants", activeMerchants),
		zap.Int("trashedMerchants", trashedMerchants))

	return nil
}
