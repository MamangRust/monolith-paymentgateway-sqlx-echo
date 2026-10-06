package handler

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
)

type SaldoQueryHandleGrpc interface {
	pb.SaldoQueryServiceServer
}

type SaldoCommandHandleGrpc interface {
	pb.SaldoCommandServiceServer
}
