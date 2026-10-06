package handlerstats

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-card/service"
	cardstatsservice "github.com/MamangRust/monolith-payment-gateway-card/service/stats"
	cardstatsbycard "github.com/MamangRust/monolith-payment-gateway-card/service/statsbycard"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/grpc"
)

type cardStatsBalanceGrpc struct {
	pb.UnimplementedCardStatsBalanceServiceServer

	cardStatsBalance cardstatsservice.CardStatsBalanceService

	cardStatsBalanceByCard cardstatsbycard.CardStatsBalanceByCardService
}

func NewCardStatsBalanceGrpc(service service.Service) CardStatsBalanceService {
	return &cardStatsBalanceGrpc{
		cardStatsBalance:       service,
		cardStatsBalanceByCard: service,
	}
}

func (s *cardStatsBalanceGrpc) FindMonthlyBalance(ctx context.Context, req *pb.FindYearBalance) (*pb.ApiResponseMonthlyBalance, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, card_errors.ErrGrpcInvalidYear
	}

	res, err := s.cardStatsBalance.FindMonthlyBalance(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.CardResponseMonthlyBalance, len(res))
	for i, item := range res {
		protoData[i] = &pb.CardResponseMonthlyBalance{
			Month:        item.Month,
			TotalBalance: int64(item.TotalBalance),
		}
	}

	return &pb.ApiResponseMonthlyBalance{
		Status:  "success",
		Message: "Monthly balance retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsBalanceGrpc) FindYearlyBalance(ctx context.Context, req *pb.FindYearBalance) (*pb.ApiResponseYearlyBalance, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, card_errors.ErrGrpcInvalidYear
	}

	res, err := s.cardStatsBalance.FindYearlyBalance(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.CardResponseYearlyBalance, len(res))
	for i, item := range res {
		protoData[i] = &pb.CardResponseYearlyBalance{
			Year:         item.Year.Int.String(),
			TotalBalance: item.TotalBalance,
		}
	}

	return &pb.ApiResponseYearlyBalance{
		Status:  "success",
		Message: "Yearly balance retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsBalanceGrpc) FindMonthlyBalanceByCardNumber(ctx context.Context, req *pb.FindYearBalanceCardNumber) (*pb.ApiResponseMonthlyBalance, error) {
	card_number := req.GetCardNumber()
	year := int(req.GetYear())

	if year <= 0 {
		return nil, card_errors.ErrGrpcInvalidYear
	}
	if card_number == "" {
		return nil, card_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthYearCardNumberCard{
		CardNumber: card_number,
		Year:       year,
	}

	res, err := s.cardStatsBalanceByCard.FindMonthlyBalancesByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.CardResponseMonthlyBalance, len(res))
	for i, item := range res {
		protoData[i] = &pb.CardResponseMonthlyBalance{
			Month:        item.Month,
			TotalBalance: int64(item.TotalBalance),
		}
	}

	return &pb.ApiResponseMonthlyBalance{
		Status:  "success",
		Message: "Monthly balance retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsBalanceGrpc) FindYearlyBalanceByCardNumber(ctx context.Context, req *pb.FindYearBalanceCardNumber) (*pb.ApiResponseYearlyBalance, error) {
	card_number := req.GetCardNumber()
	year := int(req.GetYear())

	if year <= 0 {
		return nil, card_errors.ErrGrpcInvalidYear
	}
	if card_number == "" {
		return nil, card_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthYearCardNumberCard{
		CardNumber: card_number,
		Year:       year,
	}

	res, err := s.cardStatsBalanceByCard.FindYearlyBalanceByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.CardResponseYearlyBalance, len(res))
	for i, item := range res {
		protoData[i] = &pb.CardResponseYearlyBalance{
			Year:         item.Year.Int.String(),
			TotalBalance: item.TotalBalance,
		}
	}

	return &pb.ApiResponseYearlyBalance{
		Status:  "success",
		Message: "Yearly balance retrieved successfully",
		Data:    protoData,
	}, nil
}
