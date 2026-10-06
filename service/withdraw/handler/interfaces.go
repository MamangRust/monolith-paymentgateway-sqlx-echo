package handler

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw"
)

type WithdrawQueryHandlerGrpc interface {
	pb.WithdrawQueryServiceServer
}

type WithdrawCommandHandlerGrpc interface {
	pb.WithdrawCommandServiceServer
}
