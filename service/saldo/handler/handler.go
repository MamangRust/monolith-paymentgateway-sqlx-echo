package handler

import (
	saldostatshandler "github.com/MamangRust/monolith-payment-gateway-saldo/handler/stats"
	"github.com/MamangRust/monolith-payment-gateway-saldo/service"
)

type handler struct {
	SaldoQueryHandleGrpc
	SaldoCommandHandleGrpc
	saldostatshandler.HandleStats
}

type Handler interface {
	SaldoQueryHandleGrpc
	SaldoCommandHandleGrpc
	saldostatshandler.HandleStats
}

func NewHandler(service service.Service) Handler {
	return &handler{
		SaldoQueryHandleGrpc:   NewSaldoQueryHandleGrpc(service),
		SaldoCommandHandleGrpc: NewSaldoCommandHandleGrpc(service),
		HandleStats:            saldostatshandler.NewSaldoStatsHandle(service),
	}
}
