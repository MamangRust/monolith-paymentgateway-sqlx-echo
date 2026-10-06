package topupstatshandler

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup/stats"
)

type TopupStatsAmountHandleGrpc interface {
	pb.TopupStatsAmountServiceServer
}

type TopupStatsMethodHandleGrpc interface {
	pb.TopupStatsMethodServiceServer
}

type TopupStatsStatusHandleGrpc interface {
	pb.TopupStatsStatusServiceServer
}
