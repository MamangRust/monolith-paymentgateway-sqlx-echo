package withdrawstatshandler

import (
	"github.com/MamangRust/monolith-payment-gateway-withdraw/service"
)

type HandleStats interface {
	WithdrawStatsAmountHandlerGrpc
	WithdrawStatsStatusHandleGrpc
}

type handleStats struct {
	WithdrawStatsAmountHandlerGrpc
	WithdrawStatsStatusHandleGrpc
}

func NewWithdrawStatsHandleGrpc(service service.Service) HandleStats {
	return &handleStats{
		WithdrawStatsAmountHandlerGrpc: NewWithdrawStatsAmountHandleGrpc(service),
		WithdrawStatsStatusHandleGrpc:  NewWithdrawStatsStatusHandleGrpc(service),
	}
}
