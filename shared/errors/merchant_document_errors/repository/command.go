package merchantdocumentrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrCreateMerchantDocumentFailed is returned when failing to create a new merchant document.
	ErrCreateMerchantDocumentFailed = errors.ErrInternal.WithMessage("Failed to create merchant document")

	// ErrUpdateMerchantDocumentFailed is returned when failing to update an existing merchant document.
	ErrUpdateMerchantDocumentFailed = errors.ErrInternal.WithMessage("Failed to update merchant document")

	// ErrUpdateMerchantDocumentStatusFailed is returned when failing to update the status of a merchant document.
	ErrUpdateMerchantDocumentStatusFailed = errors.ErrInternal.WithMessage("Failed to update merchant document status")

	// ErrTrashedMerchantDocumentFailed is returned when failing to move a merchant document to trash.
	ErrTrashedMerchantDocumentFailed = errors.ErrInternal.WithMessage("Failed to move merchant document to trash")

	// ErrRestoreMerchantDocumentFailed is returned when failing to restore a trashed merchant document.
	ErrRestoreMerchantDocumentFailed = errors.ErrInternal.WithMessage("Failed to restore merchant document from trash")

	// ErrDeleteMerchantDocumentPermanentFailed is returned when failing to permanently delete a merchant document.
	ErrDeleteMerchantDocumentPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete merchant document")

	// ErrRestoreAllMerchantDocumentsFailed is returned when failing to restore all trashed merchant documents.
	ErrRestoreAllMerchantDocumentsFailed = errors.ErrInternal.WithMessage("Failed to restore all merchant documents")

	// ErrDeleteAllMerchantDocumentsPermanentFailed is returned when failing to permanently delete all merchant documents.
	ErrDeleteAllMerchantDocumentsPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete all merchant documents")
)
