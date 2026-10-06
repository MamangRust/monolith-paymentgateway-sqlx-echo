package handler

import (
	"github.com/MamangRust/monolith-payment-gateway-auth/service"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	Auth AuthHandleGrpc
}

// NewHandler sets up the handler for the authentication service.
//
// It takes a pointer to a Deps struct, which contains all the dependencies required
// to set up the handler.
//
// The returned Handler contains the gRPC handler for the authentication service.
func NewHandler(deps *Deps) *Handler {
	return &Handler{
		Auth: NewAuthHandleGrpc(
			deps.Service,
			deps.Logger,
		),
	}
}
