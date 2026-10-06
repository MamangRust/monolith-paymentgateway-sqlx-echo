package merchantdocumentserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrMerchantDocumentNotFoundRes is returned when the requested merchant document cannot be found.
	ErrMerchantDocumentNotFoundRes = errors.ErrNotFound.WithMessage("Merchant document not found")

	// ErrFailedFindAllMerchantDocuments is returned when failing to fetch all merchant documents.
	ErrFailedFindAllMerchantDocuments = errors.ErrInternal.WithMessage("Failed to fetch merchant documents")

	// ErrFailedFindActiveMerchantDocuments is returned when failing to fetch active merchant documents.
	ErrFailedFindActiveMerchantDocuments = errors.ErrInternal.WithMessage("Failed to fetch active merchant documents")

	// ErrFailedFindTrashedMerchantDocuments is returned when failing to fetch trashed merchant documents.
	ErrFailedFindTrashedMerchantDocuments = errors.ErrInternal.WithMessage("Failed to fetch trashed merchant documents")

	// ErrFailedFindMerchantDocumentById is returned when failing to find a merchant document by ID.
	ErrFailedFindMerchantDocumentById = errors.ErrInternal.WithMessage("Failed to find merchant document by ID")
)
