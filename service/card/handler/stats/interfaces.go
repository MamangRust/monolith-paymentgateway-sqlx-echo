package handlerstats

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card/stats"
)

type CardStatsBalanceService interface {
	pb.CardStatsBalanceServiceServer
}

type CardStatsTopupService interface {
	pb.CardStatsTopupServiceServer
}

type CardStatsTransactionService interface {
	pb.CardStatsTransactionServiceServer
}

type CardStatsTransferService interface {
	pb.CardStatsTransferServiceServer
}

type CardStatsWithdrawService interface {
	pb.CardStatsWithdrawServiceServer
}
