package rolegrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var ErrGrpcRoleInvalidId = errors.NewGrpcError("Invalid Role ID", http.StatusBadRequest)

var ErrGrpcRoleInvalidName = errors.NewGrpcError("Invalid Role name", http.StatusBadRequest)
