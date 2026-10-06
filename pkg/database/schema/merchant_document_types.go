// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: merchant_document.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateMerchantDocumentParams struct {
	MerchantID   int32   `db:"merchant_id" json:"merchant_id"`
	DocumentType string  `db:"document_type" json:"document_type"`
	DocumentUrl  string  `db:"document_url" json:"document_url"`
	Status       string  `db:"status" json:"status"`
	Note         *string `db:"note" json:"note"`
}

type CreateMerchantDocumentRow struct {
	DocumentID   int32            `db:"document_id" json:"document_id"`
	MerchantID   int32            `db:"merchant_id" json:"merchant_id"`
	DocumentType string           `db:"document_type" json:"document_type"`
	DocumentUrl  string           `db:"document_url" json:"document_url"`
	Status       string           `db:"status" json:"status"`
	Note         *string          `db:"note" json:"note"`
	UploadedAt   pgtype.Timestamp `db:"uploaded_at" json:"uploaded_at"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetActiveMerchantDocumentsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetActiveMerchantDocumentsRow struct {
	DocumentID   int32            `db:"document_id" json:"document_id"`
	MerchantID   int32            `db:"merchant_id" json:"merchant_id"`
	DocumentType string           `db:"document_type" json:"document_type"`
	DocumentUrl  string           `db:"document_url" json:"document_url"`
	Status       string           `db:"status" json:"status"`
	Note         *string          `db:"note" json:"note"`
	UploadedAt   pgtype.Timestamp `db:"uploaded_at" json:"uploaded_at"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt    pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount   int64            `db:"total_count" json:"total_count"`
}

type GetMerchantDocumentRow struct {
	DocumentID   int32            `db:"document_id" json:"document_id"`
	MerchantID   int32            `db:"merchant_id" json:"merchant_id"`
	DocumentType string           `db:"document_type" json:"document_type"`
	DocumentUrl  string           `db:"document_url" json:"document_url"`
	Status       string           `db:"status" json:"status"`
	Note         *string          `db:"note" json:"note"`
	UploadedAt   pgtype.Timestamp `db:"uploaded_at" json:"uploaded_at"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetMerchantDocumentsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetMerchantDocumentsRow struct {
	DocumentID   int32            `db:"document_id" json:"document_id"`
	MerchantID   int32            `db:"merchant_id" json:"merchant_id"`
	DocumentType string           `db:"document_type" json:"document_type"`
	DocumentUrl  string           `db:"document_url" json:"document_url"`
	Status       string           `db:"status" json:"status"`
	Note         *string          `db:"note" json:"note"`
	UploadedAt   pgtype.Timestamp `db:"uploaded_at" json:"uploaded_at"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount   int64            `db:"total_count" json:"total_count"`
}

type GetTrashedMerchantDocumentsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetTrashedMerchantDocumentsRow struct {
	DocumentID   int32            `db:"document_id" json:"document_id"`
	MerchantID   int32            `db:"merchant_id" json:"merchant_id"`
	DocumentType string           `db:"document_type" json:"document_type"`
	DocumentUrl  string           `db:"document_url" json:"document_url"`
	Status       string           `db:"status" json:"status"`
	Note         *string          `db:"note" json:"note"`
	UploadedAt   pgtype.Timestamp `db:"uploaded_at" json:"uploaded_at"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt    pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount   int64            `db:"total_count" json:"total_count"`
}

type RestoreMerchantDocumentRow struct {
	DocumentID   int32            `db:"document_id" json:"document_id"`
	MerchantID   int32            `db:"merchant_id" json:"merchant_id"`
	DocumentType string           `db:"document_type" json:"document_type"`
	DocumentUrl  string           `db:"document_url" json:"document_url"`
	Status       string           `db:"status" json:"status"`
	Note         *string          `db:"note" json:"note"`
	UploadedAt   pgtype.Timestamp `db:"uploaded_at" json:"uploaded_at"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt    pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type TrashMerchantDocumentRow struct {
	DocumentID   int32            `db:"document_id" json:"document_id"`
	MerchantID   int32            `db:"merchant_id" json:"merchant_id"`
	DocumentType string           `db:"document_type" json:"document_type"`
	DocumentUrl  string           `db:"document_url" json:"document_url"`
	Status       string           `db:"status" json:"status"`
	Note         *string          `db:"note" json:"note"`
	UploadedAt   pgtype.Timestamp `db:"uploaded_at" json:"uploaded_at"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt    pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type UpdateMerchantDocumentParams struct {
	DocumentID   int32   `db:"document_id" json:"document_id"`
	DocumentType string  `db:"document_type" json:"document_type"`
	DocumentUrl  string  `db:"document_url" json:"document_url"`
	Status       string  `db:"status" json:"status"`
	Note         *string `db:"note" json:"note"`
}

type UpdateMerchantDocumentRow struct {
	DocumentID   int32            `db:"document_id" json:"document_id"`
	MerchantID   int32            `db:"merchant_id" json:"merchant_id"`
	DocumentType string           `db:"document_type" json:"document_type"`
	DocumentUrl  string           `db:"document_url" json:"document_url"`
	Status       string           `db:"status" json:"status"`
	Note         *string          `db:"note" json:"note"`
	UploadedAt   pgtype.Timestamp `db:"uploaded_at" json:"uploaded_at"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateMerchantDocumentStatusParams struct {
	DocumentID int32   `db:"document_id" json:"document_id"`
	Status     string  `db:"status" json:"status"`
	Note       *string `db:"note" json:"note"`
}

type UpdateMerchantDocumentStatusRow struct {
	DocumentID   int32            `db:"document_id" json:"document_id"`
	MerchantID   int32            `db:"merchant_id" json:"merchant_id"`
	DocumentType string           `db:"document_type" json:"document_type"`
	DocumentUrl  string           `db:"document_url" json:"document_url"`
	Status       string           `db:"status" json:"status"`
	Note         *string          `db:"note" json:"note"`
	UploadedAt   pgtype.Timestamp `db:"uploaded_at" json:"uploaded_at"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
