package repository

import (
	"context"

	apikey "github.com/MamangRust/monolith-payment-gateway-pkg/api-key"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/repository"
	"github.com/jmoiron/sqlx"
)

const createMerchant = `WITH valid_user AS (
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
FROM inserted_merchant`

const updateMerchant = `UPDATE merchants
SET
    name = $2,
    user_id = $3,
    status = $4,
    updated_at = current_timestamp
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
    updated_at`

const updateMerchantStatus = `WITH valid_owner AS (
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
FROM updated_merchant`

const trashMerchant = `UPDATE merchants
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
    deleted_at`

const restoreMerchant = `UPDATE merchants
SET
    deleted_at = NULL
WHERE
    merchant_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    merchant_id,
    merchant_no,
    name,
    api_key,
    user_id,
    status,
    created_at,
    updated_at,
    deleted_at`

const deleteMerchantPermanently = `DELETE FROM merchants
WHERE
    merchant_id = $1
    AND deleted_at IS NOT NULL`

const restoreAllMerchants = `UPDATE merchants
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL`

const deleteAllPermanentMerchants = `DELETE FROM merchants WHERE deleted_at IS NOT NULL`

type merchantCommandRepository struct {
	db *sqlx.DB
}

func NewMerchantCommandRepository(db *sqlx.DB) MerchantCommandRepository {
	return &merchantCommandRepository{
		db: db,
	}
}

func (r *merchantCommandRepository) CreateMerchant(ctx context.Context, request *requests.CreateMerchantRequest) (*db.CreateMerchantRow, error) {
	apiKey, err := apikey.GenerateApiKey()

	if err != nil {
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	var merchant db.CreateMerchantRow
	err = r.db.GetContext(ctx, &merchant, createMerchant,
		request.Name,
		apiKey,
		int32(request.UserID),
		"inactive",
	)

	if err != nil {
		return nil, sharedErrors.ErrConstraintOrFailed(err, "Merchant", "create merchant")
	}

	return &merchant, nil
}

func (r *merchantCommandRepository) UpdateMerchant(ctx context.Context, request *requests.UpdateMerchantRequest) (*db.UpdateMerchantRow, error) {
	var merchant db.UpdateMerchantRow
	err := r.db.GetContext(ctx, &merchant, updateMerchant,
		int32(*request.MerchantID),
		request.Name,
		int32(request.UserID),
		request.Status,
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Merchant", "update merchant")
	}

	return &merchant, nil
}

func (r *merchantCommandRepository) UpdateMerchantStatus(ctx context.Context, request *requests.UpdateMerchantStatusRequest) (*db.UpdateMerchantStatusRow, error) {
	var merchant db.UpdateMerchantStatusRow
	err := r.db.GetContext(ctx, &merchant, updateMerchantStatus,
		int32(*request.MerchantID),
		request.Status,
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Merchant", "update merchant status")
	}

	return &merchant, nil
}

func (r *merchantCommandRepository) TrashedMerchant(ctx context.Context, merchant_id int) (*db.Merchant, error) {
	var merchant db.TrashMerchantRow
	err := r.db.GetContext(ctx, &merchant, trashMerchant, int32(merchant_id))

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Merchant", "trash merchant")
	}

	return &db.Merchant{
		MerchantID: merchant.MerchantID,
		MerchantNo: merchant.MerchantNo,
		Name:       merchant.Name,
		ApiKey:     merchant.ApiKey,
		UserID:     merchant.UserID,
		Status:     merchant.Status,
		CreatedAt:  merchant.CreatedAt,
		UpdatedAt:  merchant.UpdatedAt,
		DeletedAt:  merchant.DeletedAt,
	}, nil
}

func (r *merchantCommandRepository) RestoreMerchant(ctx context.Context, merchant_id int) (*db.Merchant, error) {
	var merchant db.RestoreMerchantRow
	err := r.db.GetContext(ctx, &merchant, restoreMerchant, int32(merchant_id))

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Merchant", "restore merchant")
	}

	return &db.Merchant{
		MerchantID: merchant.MerchantID,
		MerchantNo: merchant.MerchantNo,
		Name:       merchant.Name,
		ApiKey:     merchant.ApiKey,
		UserID:     merchant.UserID,
		Status:     merchant.Status,
		CreatedAt:  merchant.CreatedAt,
		UpdatedAt:  merchant.UpdatedAt,
		DeletedAt:  merchant.DeletedAt,
	}, nil
}

func (r *merchantCommandRepository) DeleteMerchantPermanent(ctx context.Context, merchant_id int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteMerchantPermanently, int32(merchant_id)); err != nil {
		return false, sharedErrors.ErrNoRowsOrFailed(err, "Merchant", "delete merchant permanently")
	}

	return true, nil
}

func (r *merchantCommandRepository) RestoreAllMerchant(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllMerchants); err != nil {
		return false, merchant_errors.ErrRestoreAllMerchantFailed.WithInternal(err)
	}

	return true, nil
}

func (r *merchantCommandRepository) DeleteAllMerchantPermanent(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentMerchants); err != nil {
		return false, merchant_errors.ErrDeleteAllMerchantPermanentFailed.WithInternal(err)
	}

	return true, nil
}
