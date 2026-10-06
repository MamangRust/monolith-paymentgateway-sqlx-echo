package transferstatshandler

import (
	"github.com/MamangRust/monolith-payment-gateway-transfer/service"
)

type HandleStats interface {
	TransferStatsAmountHandleGrpc
	TransferStatsStatusHandleGrpc
}

type handleStats struct {
	TransferStatsAmountHandleGrpc
	TransferStatsStatusHandleGrpc
}

func NewTransferStatsHandleGrpc(service service.Service) HandleStats {

	return &handleStats{
		TransferStatsAmountHandleGrpc: NewTransferStatsAmountHandler(service),
		TransferStatsStatusHandleGrpc: NewTransferStatsStatusHandler(service),
	}
}
