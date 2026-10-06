package handler

import (
	"github.com/MamangRust/monolith-payment-gateway-role/service"
)

// Handler is a struct that holds the dependencies for the role service.
type Handler struct {
	RoleQuery   RoleQueryHandlerGrpc
	RoleCommand RoleCommandHandlerGrpc
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		RoleQuery:   NewRoleQueryHandleGrpc(service.RoleQuery),
		RoleCommand: NewRoleCommandHandleGrpc(service.RoleCommand),
	}
}
