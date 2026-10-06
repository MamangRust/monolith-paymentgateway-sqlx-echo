package merchantserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedSendEmail indicates a failure when sending an email related to merchant operations.
var ErrFailedSendEmail = errors.NewErrorResponse("Failed to send email", http.StatusInternalServerError)
