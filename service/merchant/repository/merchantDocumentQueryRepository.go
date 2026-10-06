package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	merchantdocument_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_document_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMerchantDocuments = `SELECT
    document_id,
    merchant_id,
    document_type,
    document_url,
    status,
    note,
    uploaded_at,
    created_at,
    updated_at,
    COUNT(*) OVER () AS total_count
FROM merchant_documents
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR document_type ILIKE '%' || $1 || '%'
        OR status ILIKE '%' || $1 || '%'
        OR note ILIKE '%' || $1 || '%'
    )
ORDER BY document_id
LIMIT $2
OFFSET
    $3`

const getActiveMerchantDocuments = `SELECT
    document_id,
    merchant_id,
    document_type,
    document_url,
    status,
    note,
    uploaded_at,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM merchant_documents
WHERE
    deleted_at IS NULL
    AND status != 'deleted'
    AND (
        $1::TEXT IS NULL
        OR document_type ILIKE '%' || $1 || '%'
        OR status ILIKE '%' || $1 || '%'
        OR note ILIKE '%' || $1 || '%'
    )
ORDER BY document_id
LIMIT $2
OFFSET
    $3`

const getTrashedMerchantDocuments = `SELECT
    document_id,
    merchant_id,
    document_type,
    document_url,
    status,
    note,
    uploaded_at,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM merchant_documents
WHERE
    deleted_at IS NOT NULL
    AND (
        $1::TEXT IS NULL
        OR document_type ILIKE '%' || $1 || '%'
        OR status ILIKE '%' || $1 || '%'
        OR note ILIKE '%' || $1 || '%'
    )
ORDER BY document_id
LIMIT $2
OFFSET
    $3`

const getMerchantDocument = `SELECT
    document_id,
    merchant_id,
    document_type,
    document_url,
    status,
    note,
    uploaded_at,
    created_at,
    updated_at
FROM merchant_documents
WHERE
    document_id = $1
    AND deleted_at IS NULL`

// merchantDocumentQueryRepository provides methods to query merchant documents from the database.
type merchantDocumentQueryRepository struct {
	db *sqlx.DB
}

// NewMerchantDocumentQueryRepository creates a new instance of merchantDocumentQueryRepository.
func NewMerchantDocumentQueryRepository(db *sqlx.DB) MerchantDocumentQueryRepository {
	return &merchantDocumentQueryRepository{
		db: db,
	}
}

func (r *merchantDocumentQueryRepository) FindAllDocuments(ctx context.Context, req *requests.FindAllMerchantDocuments) ([]*db.GetMerchantDocumentsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var docs []*db.GetMerchantDocumentsRow
	err := r.db.SelectContext(ctx, &docs, getMerchantDocuments,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)
	if err != nil {
		return nil, merchantdocument_errors.ErrFindAllMerchantDocumentsFailed.WithInternal(err)
	}

	return docs, nil
}

func (r *merchantDocumentQueryRepository) FindByActiveDocuments(ctx context.Context, req *requests.FindAllMerchantDocuments) ([]*db.GetActiveMerchantDocumentsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var docs []*db.GetActiveMerchantDocumentsRow
	err := r.db.SelectContext(ctx, &docs, getActiveMerchantDocuments,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)
	if err != nil {
		return nil, merchantdocument_errors.ErrFindActiveMerchantDocumentsFailed.WithInternal(err)
	}

	return docs, nil
}

func (r *merchantDocumentQueryRepository) FindByTrashedDocuments(ctx context.Context, req *requests.FindAllMerchantDocuments) ([]*db.GetTrashedMerchantDocumentsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var docs []*db.GetTrashedMerchantDocumentsRow
	err := r.db.SelectContext(ctx, &docs, getTrashedMerchantDocuments,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)
	if err != nil {
		return nil, merchantdocument_errors.ErrFindTrashedMerchantDocumentsFailed.WithInternal(err)
	}

	return docs, nil
}

func (r *merchantDocumentQueryRepository) FindByIdDocument(ctx context.Context, id int) (*db.GetMerchantDocumentRow, error) {
	var doc db.GetMerchantDocumentRow
	err := r.db.GetContext(ctx, &doc, getMerchantDocument, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchantdocument_errors.ErrMerchantDocumentNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}
	return &doc, nil
}
