package saldostatshandler

import (
	"github.com/MamangRust/monolith-payment-gateway-saldo/service"
)

type HandleStats interface {
	SaldoStatsBalanceHandleGrpc
	SaldoStatsTotalBalanceHandleGrpc
}

type handleStats struct {
	SaldoStatsBalanceHandleGrpc
	SaldoStatsTotalBalanceHandleGrpc
}

func NewSaldoStatsHandle(service service.Service) HandleStats {
	return &handleStats{
		SaldoStatsBalanceHandleGrpc:      NewSaldoStatsBalanceHandleGrpc(service),
		SaldoStatsTotalBalanceHandleGrpc: NewSaldoStatsTotalBalanceHandleGrpc(service),
	}
}
