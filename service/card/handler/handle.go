package handler

import (
	handlerstats "github.com/MamangRust/monolith-payment-gateway-card/handler/stats"
	"github.com/MamangRust/monolith-payment-gateway-card/service"
)

type Handler interface {
	CardQueryService
	CardCommandService
	CardDashboardService
	handlerstats.HandlerStats
}

type handler struct {
	CardQueryService
	CardCommandService
	CardDashboardService
	handlerstats.HandlerStats
}

// NewHandler creates a new CardHandler instance.
func NewHandler(service service.Service) Handler {
	return &handler{
		CardQueryService:     NewCardQueryHandleGrpc(service),
		CardCommandService:   NewCardCommandHandleGrpc(service),
		CardDashboardService: NewCardDashboardHandleGrpc(service),
		HandlerStats:         handlerstats.NewHandlerStats(service),
	}
}
