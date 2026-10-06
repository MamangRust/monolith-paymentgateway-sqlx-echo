package transactionstatshandler

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction/stats"
)

type TransactionStatsAmountHandlerGrpc interface {
	pb.TransactionStatsAmountServiceServer
}

type TransactionStatsMethodHandleGrpc interface {
	pb.TransactionStatsMethodServiceServer
}

type TransactionStatsStatusHandleGrpc interface {
	pb.TransactionStatsStatusServiceServer
}
