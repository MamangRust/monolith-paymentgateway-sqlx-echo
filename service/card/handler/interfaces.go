package handler

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
)

type CardQueryService interface {
	pb.CardQueryServiceServer
}

type CardCommandService interface {
	pb.CardCommandServiceServer
}

type CardDashboardService interface {
	pb.CardDashboardServiceServer
}
