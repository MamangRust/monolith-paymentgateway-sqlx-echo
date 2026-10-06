package handler

import (
	transactionstatshandler "github.com/MamangRust/monolith-payment-gateway-transaction/handler/stats"
	"github.com/MamangRust/monolith-payment-gateway-transaction/service"
)

type Handler interface {
	TransactionQueryHandleGrpc
	TransactionCommandHandleGrpc
	transactionstatshandler.HandleStats
}

type handler struct {
	TransactionQueryHandleGrpc
	TransactionCommandHandleGrpc
	transactionstatshandler.HandleStats
}

func NewHandler(service service.Service) Handler {
	return &handler{
		TransactionQueryHandleGrpc:   NewTransactionQueryHandleGrpc(service),
		TransactionCommandHandleGrpc: NewTransactionCommandHandleGrpc(service),
		HandleStats:                  transactionstatshandler.NewTransactionStatsHandleGrpc(service),
	}
}
