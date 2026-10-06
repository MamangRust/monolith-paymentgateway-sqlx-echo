package withdrawserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedSendEmail is used when failed to send email
var ErrFailedSendEmail = errors.NewErrorResponse("Failed to send email", http.StatusInternalServerError)
