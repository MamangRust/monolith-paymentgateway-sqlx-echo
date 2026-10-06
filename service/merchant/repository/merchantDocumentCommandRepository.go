package repository

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	merchantdocument_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_document_errors/repository"
	"github.com/jmoiron/sqlx"
)

const createMerchantDocument = `WITH valid_merchant AS (
    SELECT m.merchant_id, u.email
    FROM merchants m
    JOIN users u ON u.user_id = m.user_id AND u.deleted_at IS NULL
    WHERE m.merchant_id = $1 AND m.deleted_at IS NULL
), inserted_document AS (
    INSERT INTO merchant_documents (
        merchant_id, document_type, document_url, status, note, uploaded_at, updated_at
    )
    SELECT merchant_id, $2, $3, $4, $5, current_timestamp, current_timestamp
    FROM valid_merchant
    RETURNING document_id, merchant_id, document_type, document_url, status, note, uploaded_at, created_at, updated_at
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'merchant-document-create:' || document_id::text,
        'email-service-topic-merchant-document-create',
        document_id::text,
        jsonb_build_object(
            'event_id', gen_random_uuid()::text,
            'schema_version', 1,
            'event_type', 'merchant_document.created',
            'occurred_at', now(),
            'email', u.email,
            'subject', 'Merchant Verification Pending - Action Required',
            'body', 'Your merchant document has been submitted and is pending review.'
        )
    FROM inserted_document d
    JOIN valid_merchant u ON u.merchant_id = d.merchant_id
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT document_id, merchant_id, document_type, document_url, status, note, uploaded_at, created_at, updated_at
FROM inserted_document`

const updateMerchantDocument = `UPDATE merchant_documents
SET
    document_type = $2,
    document_url = $3,
    status = $4,
    note = $5,
    updated_at = current_timestamp
WHERE
    document_id = $1
    AND deleted_at IS NULL
RETURNING
    document_id,
    merchant_id,
    document_type,
    document_url,
    status,
    note,
    uploaded_at,
    created_at,
    updated_at`

const updateMerchantDocumentStatus = `WITH valid_merchant AS (
    SELECT m.merchant_id
    FROM merchants m
    JOIN users u ON u.user_id = m.user_id AND u.deleted_at IS NULL
    WHERE m.merchant_id = (
        SELECT md.merchant_id FROM merchant_documents md
        WHERE md.document_id = $1 AND md.deleted_at IS NULL
    ) AND m.deleted_at IS NULL
), updated_document AS (
    UPDATE merchant_documents
    SET status = $2, note = $3, event_version = event_version + 1, updated_at = current_timestamp
    WHERE merchant_documents.document_id = $1
      AND merchant_documents.deleted_at IS NULL
      AND merchant_documents.merchant_id IN (SELECT valid_merchant.merchant_id FROM valid_merchant)
    RETURNING document_id, merchant_id, document_type, document_url, status, note, uploaded_at, created_at, updated_at, event_version
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'merchant-document-status:' || d.document_id::text || ':' || d.event_version::text,
        'email-service-topic-merchant-document-update-status',
        d.document_id::text,
        jsonb_build_object(
            'event_id', gen_random_uuid()::text,
            'schema_version', 1,
            'event_type', 'merchant_document.status_updated',
            'occurred_at', now(),
            'email', u.email,
            'subject', 'Merchant Document Status: ' || initcap(d.status),
            'body', 'Your merchant document status has been updated to ' || d.status || '.' || CASE WHEN d.note IS NULL OR d.note = '' THEN '' ELSE ' Reviewer note: ' || d.note END,
            'event_version', d.event_version
        )
    FROM updated_document d
    JOIN merchants m ON m.merchant_id = d.merchant_id AND m.deleted_at IS NULL
    JOIN users u ON u.user_id = m.user_id AND u.deleted_at IS NULL
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT d.document_id, d.merchant_id, d.document_type, d.document_url, d.status, d.note, d.uploaded_at, d.created_at, d.updated_at
FROM updated_document d`

const trashMerchantDocument = `UPDATE merchant_documents
SET
    deleted_at = current_timestamp
WHERE
    document_id = $1
    AND deleted_at IS NULL
RETURNING
    document_id,
    merchant_id,
    document_type,
    document_url,
    status,
    note,
    uploaded_at,
    created_at,
    updated_at,
    deleted_at`

const restoreMerchantDocument = `UPDATE merchant_documents
SET
    deleted_at = NULL
WHERE
    document_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    document_id,
    merchant_id,
    document_type,
    document_url,
    status,
    note,
    uploaded_at,
    created_at,
    updated_at,
    deleted_at`

const deleteMerchantDocumentPermanently = `DELETE FROM merchant_documents
WHERE
    document_id = $1
    AND deleted_at IS NOT NULL`

const restoreAllMerchantDocuments = `UPDATE merchant_documents
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL`

const deleteAllPermanentMerchantDocuments = `DELETE FROM merchant_documents WHERE deleted_at IS NOT NULL`

type merchantDocumentCommandRepository struct {
	db *sqlx.DB
}

func NewMerchantDocumentCommandRepository(db *sqlx.DB) MerchantDocumentCommandRepository {
	return &merchantDocumentCommandRepository{
		db: db,
	}
}

func (r *merchantDocumentCommandRepository) CreateMerchantDocument(ctx context.Context, request *requests.CreateMerchantDocumentRequest) (*db.CreateMerchantDocumentRow, error) {
	note := ""

	var doc db.CreateMerchantDocumentRow
	err := r.db.GetContext(ctx, &doc, createMerchantDocument,
		int32(request.MerchantID),
		request.DocumentType,
		request.DocumentUrl,
		"pending",
		&note,
	)
	if err != nil {
		return nil, merchantdocument_errors.ErrCreateMerchantDocumentFailed.WithInternal(err)
	}

	return &doc, nil
}

func (r *merchantDocumentCommandRepository) UpdateMerchantDocument(ctx context.Context, request *requests.UpdateMerchantDocumentRequest) (*db.UpdateMerchantDocumentRow, error) {
	note := ""

	var doc db.UpdateMerchantDocumentRow
	err := r.db.GetContext(ctx, &doc, updateMerchantDocument,
		int32(*request.DocumentID),
		request.DocumentType,
		request.DocumentUrl,
		request.Status,
		&note,
	)
	if err != nil {
		return nil, merchantdocument_errors.ErrUpdateMerchantDocumentFailed.WithInternal(err)
	}

	return &doc, nil
}

func (r *merchantDocumentCommandRepository) UpdateMerchantDocumentStatus(ctx context.Context, request *requests.UpdateMerchantDocumentStatusRequest) (*db.UpdateMerchantDocumentStatusRow, error) {
	note := ""

	var doc db.UpdateMerchantDocumentStatusRow
	err := r.db.GetContext(ctx, &doc, updateMerchantDocumentStatus,
		int32(*request.DocumentID),
		request.Status,
		&note,
	)
	if err != nil {
		return nil, merchantdocument_errors.ErrUpdateMerchantDocumentStatusFailed.WithInternal(err)
	}

	return &doc, nil
}

func (r *merchantDocumentCommandRepository) TrashedMerchantDocument(ctx context.Context, documentID int) (*db.MerchantDocument, error) {
	var doc db.TrashMerchantDocumentRow
	err := r.db.GetContext(ctx, &doc, trashMerchantDocument, int32(documentID))
	if err != nil {
		return nil, merchantdocument_errors.ErrTrashedMerchantDocumentFailed.WithInternal(err)
	}

	return &db.MerchantDocument{
		DocumentID:   doc.DocumentID,
		MerchantID:   doc.MerchantID,
		DocumentType: doc.DocumentType,
		DocumentUrl:  doc.DocumentUrl,
		Status:       doc.Status,
		Note:         doc.Note,
		UploadedAt:   doc.UploadedAt,
		CreatedAt:    doc.CreatedAt,
		UpdatedAt:    doc.UpdatedAt,
		DeletedAt:    doc.DeletedAt,
	}, nil
}

func (r *merchantDocumentCommandRepository) RestoreMerchantDocument(ctx context.Context, documentID int) (*db.MerchantDocument, error) {
	var doc db.RestoreMerchantDocumentRow
	err := r.db.GetContext(ctx, &doc, restoreMerchantDocument, int32(documentID))
	if err != nil {
		return nil, merchantdocument_errors.ErrRestoreMerchantDocumentFailed.WithInternal(err)
	}

	return &db.MerchantDocument{
		DocumentID:   doc.DocumentID,
		MerchantID:   doc.MerchantID,
		DocumentType: doc.DocumentType,
		DocumentUrl:  doc.DocumentUrl,
		Status:       doc.Status,
		Note:         doc.Note,
		UploadedAt:   doc.UploadedAt,
		CreatedAt:    doc.CreatedAt,
		UpdatedAt:    doc.UpdatedAt,
		DeletedAt:    doc.DeletedAt,
	}, nil
}

func (r *merchantDocumentCommandRepository) DeleteMerchantDocumentPermanent(ctx context.Context, documentID int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteMerchantDocumentPermanently, int32(documentID)); err != nil {
		return false, merchantdocument_errors.ErrDeleteMerchantDocumentPermanentFailed.WithInternal(err)
	}

	return true, nil
}

func (r *merchantDocumentCommandRepository) RestoreAllMerchantDocument(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllMerchantDocuments); err != nil {
		return false, merchantdocument_errors.ErrRestoreAllMerchantDocumentsFailed.WithInternal(err)
	}

	return true, nil
}

func (r *merchantDocumentCommandRepository) DeleteAllMerchantDocumentPermanent(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentMerchantDocuments); err != nil {
		return false, merchantdocument_errors.ErrDeleteAllMerchantDocumentsPermanentFailed.WithInternal(err)
	}

	return true, nil
}
