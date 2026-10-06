package handler

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-card/service"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/grpc"

	"google.golang.org/protobuf/types/known/emptypb"
)

type cardDashboardHandleGrpc struct {
	pb.UnimplementedCardDashboardServiceServer

	cardDashboard service.CardDashboardService
}

func NewCardDashboardHandleGrpc(service service.Service) CardDashboardService {
	return &cardDashboardHandleGrpc{
		cardDashboard: service,
	}
}

func (s *cardDashboardHandleGrpc) DashboardCard(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseDashboardCard, error) {
	dashboardCard, err := s.cardDashboard.DashboardCard(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseDashboardCard{
		Status:  "success",
		Message: "Dashboard card retrieved successfully",
		Data:    s.ToGrpcDashboardCard(dashboardCard),
	}, nil
}

func (s *cardDashboardHandleGrpc) DashboardCardNumber(ctx context.Context, req *pb.FindByCardNumberRequest) (*pb.ApiResponseDashboardCardNumber, error) {
	cardNumber := req.GetCardNumber()

	if cardNumber == "" {
		return nil, card_errors.ErrGrpcInvalidCardNumber
	}

	dashboardCard, err := s.cardDashboard.DashboardCardCardNumber(ctx, cardNumber)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseDashboardCardNumber{
		Status:  "success",
		Message: "Dashboard card for card number retrieved successfully",
		Data:    s.ToGrpcDashboardCardNumber(dashboardCard),
	}, nil
}

func (s *cardDashboardHandleGrpc) ToGrpcDashboardCard(d *response.DashboardCard) *pb.CardResponseDashboard {
	if d == nil {
		return &pb.CardResponseDashboard{}
	}

	return &pb.CardResponseDashboard{
		TotalBalance:     int64Value(d.TotalBalance),
		TotalTopup:       int64Value(d.TotalTopup),
		TotalTransaction: int64Value(d.TotalTransaction),
		TotalTransfer:    int64Value(d.TotalTransfer),
		TotalWithdraw:    int64Value(d.TotalWithdraw),
	}
}

func (s *cardDashboardHandleGrpc) ToGrpcDashboardCardNumber(d *response.DashboardCardCardNumber) *pb.CardResponseDashboardCardNumber {
	if d == nil {
		return &pb.CardResponseDashboardCardNumber{}
	}

	return &pb.CardResponseDashboardCardNumber{
		TotalBalance:          int64Value(d.TotalBalance),
		TotalTopup:            int64Value(d.TotalTopup),
		TotalTransaction:      int64Value(d.TotalTransaction),
		TotalTransferSend:     int64Value(d.TotalTransferSend),
		TotalTransferReceiver: int64Value(d.TotalTransferReceiver),
		TotalWithdraw:         int64Value(d.TotalWithdraw),
	}
}

func int64Value(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}
