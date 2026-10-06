package handler

import (
	topupstatshandler "github.com/MamangRust/monolith-payment-gateway-topup/handler/stats"
	"github.com/MamangRust/monolith-payment-gateway-topup/service"
)

type Handler interface {
	TopupQueryHandleGrpc
	TopupCommandHandleGrpc
	topupstatshandler.HandleStats
}

type handler struct {
	TopupQueryHandleGrpc
	TopupCommandHandleGrpc
	topupstatshandler.HandleStats
}

func NewHandler(service service.Service) Handler {
	return &handler{
		TopupQueryHandleGrpc:   NewTopupQueryHandleGrpc(service),
		TopupCommandHandleGrpc: NewTopupCommandHandleGrpc(service),
		HandleStats:            topupstatshandler.NewTopupStatsHandleGrpc(service),
	}
}
