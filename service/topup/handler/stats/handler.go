package topupstatshandler

import (
	"github.com/MamangRust/monolith-payment-gateway-topup/service"
)

type HandleStats interface {
	TopupStatsAmountHandleGrpc
	TopupStatsMethodHandleGrpc
	TopupStatsStatusHandleGrpc
}

type handlerStats struct {
	TopupStatsAmountHandleGrpc
	TopupStatsMethodHandleGrpc
	TopupStatsStatusHandleGrpc
}

func NewTopupStatsHandleGrpc(service service.Service) HandleStats {
	return &handlerStats{
		TopupStatsAmountHandleGrpc: NewTopupStatsAmountHandleGrpc(service),
		TopupStatsMethodHandleGrpc: NewTopupStatsMethodHandleGrpc(service),
		TopupStatsStatusHandleGrpc: NewTopupStatsStatusHandleGrpc(service),
	}
}
