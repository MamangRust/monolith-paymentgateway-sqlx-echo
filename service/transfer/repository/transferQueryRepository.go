package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	transfer_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transfer_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getTransfers = `SELECT
    transfer_id,
    transfer_no,
    transfer_from,
    transfer_to,
    transfer_amount,
    transfer_time,
    status,
    created_at,
    updated_at,
    COUNT(*) OVER () AS total_count
FROM transfers
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR transfer_from ILIKE '%' || $1 || '%'
        OR transfer_to ILIKE '%' || $1 || '%'
    )
ORDER BY transfer_time DESC
LIMIT $2
OFFSET
    $3`

const getTransferByID = `SELECT
    transfer_id,
    transfer_no,
    transfer_from,
    transfer_to,
    transfer_amount,
    transfer_time,
    status,
    created_at,
    updated_at
FROM transfers
WHERE
    transfer_id = $1
    AND deleted_at IS NULL`

const getActiveTransfers = `SELECT
    transfer_id,
    transfer_no,
    transfer_from,
    transfer_to,
    transfer_amount,
    transfer_time,
    status,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM transfers
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR transfer_from ILIKE '%' || $1 || '%'
        OR transfer_to ILIKE '%' || $1 || '%'
    )
ORDER BY transfer_time DESC
LIMIT $2
OFFSET
    $3`

const getTrashedTransfers = `SELECT
    transfer_id,
    transfer_no,
    transfer_from,
    transfer_to,
    transfer_amount,
    transfer_time,
    status,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM transfers
WHERE
    deleted_at IS NOT NULL
    AND (
        $1::TEXT IS NULL
        OR transfer_from ILIKE '%' || $1 || '%'
        OR transfer_to ILIKE '%' || $1 || '%'
    )
ORDER BY transfer_time DESC
LIMIT $2
OFFSET
    $3`

const getTransfersBySourceCard = `SELECT
    transfer_id,
    transfer_no,
    transfer_from,
    transfer_to,
    transfer_amount,
    transfer_time,
    status,
    created_at,
    updated_at
FROM transfers
WHERE
    deleted_at IS NULL
    AND transfer_from = $1
ORDER BY transfer_time DESC`

const getTransfersByDestinationCard = `SELECT
    transfer_id,
    transfer_no,
    transfer_from,
    transfer_to,
    transfer_amount,
    transfer_time,
    status,
    created_at,
    updated_at
FROM transfers
WHERE
    deleted_at IS NULL
    AND transfer_to = $1
ORDER BY transfer_time DESC`

type transferQueryRepository struct {
	db *sqlx.DB
}

func NewTransferQueryRepository(db *sqlx.DB) TransferQueryRepository {
	return &transferQueryRepository{
		db: db,
	}
}

func (r *transferQueryRepository) FindAll(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetTransfersRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var transfers []*db.GetTransfersRow
	err := r.db.SelectContext(ctx, &transfers, getTransfers,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, transfer_errors.ErrFindAllTransfersFailed.WithInternal(err)
	}

	return transfers, nil
}

func (r *transferQueryRepository) FindByActive(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetActiveTransfersRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var transfers []*db.GetActiveTransfersRow
	err := r.db.SelectContext(ctx, &transfers, getActiveTransfers,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, transfer_errors.ErrFindActiveTransfersFailed.WithInternal(err)
	}

	return transfers, nil
}

func (r *transferQueryRepository) FindByTrashed(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetTrashedTransfersRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var transfers []*db.GetTrashedTransfersRow
	err := r.db.SelectContext(ctx, &transfers, getTrashedTransfers,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, transfer_errors.ErrFindTrashedTransfersFailed.WithInternal(err)
	}

	return transfers, nil
}

func (r *transferQueryRepository) FindById(ctx context.Context, id int) (*db.GetTransferByIDRow, error) {
	if id <= 0 {
		return nil, sharedErrors.NewBadRequestError("transfer ID must be greater than zero")
	}

	var transfer db.GetTransferByIDRow
	err := r.db.GetContext(ctx, &transfer, getTransferByID, int32(id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, transfer_errors.ErrTransferNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &transfer, nil
}

func (r *transferQueryRepository) FindTransferByTransferFrom(ctx context.Context, transfer_from string) ([]*db.GetTransfersBySourceCardRow, error) {
	var transfers []*db.GetTransfersBySourceCardRow
	err := r.db.SelectContext(ctx, &transfers, getTransfersBySourceCard, transfer_from)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, transfer_errors.ErrFindTransferByTransferFromFailed.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return transfers, nil
}

func (r *transferQueryRepository) FindTransferByTransferTo(ctx context.Context, transfer_to string) ([]*db.GetTransfersByDestinationCardRow, error) {
	var transfers []*db.GetTransfersByDestinationCardRow
	err := r.db.SelectContext(ctx, &transfers, getTransfersByDestinationCard, transfer_to)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, transfer_errors.ErrFindTransferByTransferToFailed.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}
	return transfers, nil
}
