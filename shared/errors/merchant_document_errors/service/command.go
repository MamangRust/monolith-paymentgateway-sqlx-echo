package merchantdocumentserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedCreateMerchantDocument is returned when failing to create a new merchant document.
	ErrFailedCreateMerchantDocument = errors.ErrInternal.WithMessage("Failed to create merchant document")

	// ErrFailedUpdateMerchantDocument is returned when failing to update an existing merchant document.
	ErrFailedUpdateMerchantDocument = errors.ErrInternal.WithMessage("Failed to update merchant document")

	// ErrFailedTrashMerchantDocument is returned when failing to move a merchant document to trash.
	ErrFailedTrashMerchantDocument = errors.ErrInternal.WithMessage("Failed to move merchant document to trash")

	// ErrFailedRestoreMerchantDocument is returned when failing to restore a trashed merchant document.
	ErrFailedRestoreMerchantDocument = errors.ErrInternal.WithMessage("Failed to restore merchant document")

	// ErrFailedDeleteMerchantDocument is returned when failing to permanently delete a merchant document.
	ErrFailedDeleteMerchantDocument = errors.ErrInternal.WithMessage("Failed to delete merchant document permanently")

	// ErrFailedRestoreAllMerchantDocuments is returned when failing to restore all trashed merchant documents.
	ErrFailedRestoreAllMerchantDocuments = errors.ErrInternal.WithMessage("Failed to restore all merchant documents")

	// ErrFailedDeleteAllMerchantDocuments is returned when failing to permanently delete all merchant documents.
	ErrFailedDeleteAllMerchantDocuments = errors.ErrInternal.WithMessage("Failed to delete all merchant documents permanently")
)
