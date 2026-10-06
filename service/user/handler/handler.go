package handler

import (
	"github.com/MamangRust/monolith-payment-gateway-user/service"
)

// Handler groups all user gRPC handlers.
type Handler interface {
	UserQueryHandleGrpc
	UserCommandHandleGrpc
}

type handler struct {
	UserQueryHandleGrpc
	UserCommandHandleGrpc
}

// NewHandler initializes user gRPC handlers.
func NewHandler(service service.Service) Handler {
	return &handler{
		UserQueryHandleGrpc:   NewUserQueryHandleGrpc(service),
		UserCommandHandleGrpc: NewUserCommandHandleGrpc(service),
	}
}
