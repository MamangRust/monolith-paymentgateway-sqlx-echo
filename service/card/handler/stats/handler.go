package handlerstats

import (
	"github.com/MamangRust/monolith-payment-gateway-card/service"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
)

type DepsStats struct {
	Service service.Service
	Logger  logger.LoggerInterface
}

type HandlerStats interface {
	CardStatsBalanceService
	CardStatsTransferService
	CardStatsWithdrawService
	CardStatsTopupService
	CardStatsTransactionService
}

type handlerStats struct {
	CardStatsBalanceService
	CardStatsTransferService
	CardStatsWithdrawService
	CardStatsTopupService
	CardStatsTransactionService
}

func NewHandlerStats(service service.Service) HandlerStats {
	return &handlerStats{
		NewCardStatsBalanceGrpc(service),
		NewCardStatsTransferGrpc(service),
		NewCardStatsWithdrawGrpc(service),
		NewCardStatsTopupGrpc(service),
		NewCardStatsTransactionGrpc(service),
	}
}
