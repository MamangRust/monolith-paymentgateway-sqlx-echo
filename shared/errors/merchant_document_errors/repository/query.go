package merchantdocumentrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFindAllMerchantDocumentsFailed is returned when failing to retrieve all merchant documents.
	ErrFindAllMerchantDocumentsFailed = errors.ErrInternal.WithMessage("Failed to find all merchant documents")

	// ErrFindActiveMerchantDocumentsFailed is returned when failing to retrieve active merchant documents.
	ErrFindActiveMerchantDocumentsFailed = errors.ErrInternal.WithMessage("Failed to find active merchant documents")

	// ErrFindTrashedMerchantDocumentsFailed is returned when failing to retrieve trashed merchant documents.
	ErrFindTrashedMerchantDocumentsFailed = errors.ErrInternal.WithMessage("Failed to find trashed merchant documents")

	// ErrFindMerchantDocumentByIdFailed is returned when failing to retrieve a merchant document by ID.
	ErrFindMerchantDocumentByIdFailed = errors.ErrInternal.WithMessage("Failed to find merchant document by ID")

	// ErrMerchantDocumentNotFound is returned when a merchant document is not found.
	ErrMerchantDocumentNotFound = errors.ErrNotFound.WithMessage("Merchant document not found")
)
