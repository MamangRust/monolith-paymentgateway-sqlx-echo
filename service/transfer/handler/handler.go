package handler

import (
	transferstatshandler "github.com/MamangRust/monolith-payment-gateway-transfer/handler/stats"
	"github.com/MamangRust/monolith-payment-gateway-transfer/service"
)

type Handler interface {
	TransferQueryHandleGrpc
	TransferCommandHandleGrpc
	transferstatshandler.HandleStats
}

type handler struct {
	TransferQueryHandleGrpc
	TransferCommandHandleGrpc
	transferstatshandler.HandleStats
}

func NewHandler(service service.Service) Handler {
	return &handler{
		TransferQueryHandleGrpc:   NewTransferQueryHandler(service),
		TransferCommandHandleGrpc: NewTransferCommandHandler(service),
		HandleStats:               transferstatshandler.NewTransferStatsHandleGrpc(service),
	}
}
