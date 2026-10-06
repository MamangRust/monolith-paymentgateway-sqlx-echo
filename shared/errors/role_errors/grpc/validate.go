package rolegrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcValidateCreateRole = errors.NewGrpcError("validation failed: invalid create Role request", http.StatusBadRequest)
	ErrGrpcValidateUpdateRole = errors.NewGrpcError("validation failed: invalid update Role request", http.StatusBadRequest)
)
