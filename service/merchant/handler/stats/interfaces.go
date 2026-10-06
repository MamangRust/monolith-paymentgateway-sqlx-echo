package merchantstatshandler

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant/stats"
)

type MerchantStatsAmountHandleGrpc interface {
	pb.MerchantStatsAmountServiceServer
}

type MerchantStatsMethodHandleGrpc interface {
	pb.MerchantStatsMethodServiceServer
}

type MerchantStatsTotalAmountHandleGrpc interface {
	pb.MerchantStatsTotalAmountServiceServer
}
