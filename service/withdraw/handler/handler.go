package handler

import (
	withdrawstatshandler "github.com/MamangRust/monolith-payment-gateway-withdraw/handler/stats"
	"github.com/MamangRust/monolith-payment-gateway-withdraw/service"
)

type Handler interface {
	WithdrawQueryHandlerGrpc
	WithdrawCommandHandlerGrpc
	withdrawstatshandler.HandleStats
}

type handler struct {
	WithdrawQueryHandlerGrpc
	WithdrawCommandHandlerGrpc
	withdrawstatshandler.HandleStats
}

func NewHandler(service service.Service) Handler {
	return &handler{
		WithdrawQueryHandlerGrpc:   NewWithdrawQueryHandleGrpc(service),
		WithdrawCommandHandlerGrpc: NewWithdrawCommandHandleGrpc(service),
		HandleStats:                withdrawstatshandler.NewWithdrawStatsHandleGrpc(service),
	}
}
