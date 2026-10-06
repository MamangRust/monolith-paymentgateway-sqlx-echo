package refreshtokenrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrParseDate is returned when parsing the expiration date of a token fails.
var ErrParseDate = errors.ErrInternal.WithMessage("failed to parse expiration date")
