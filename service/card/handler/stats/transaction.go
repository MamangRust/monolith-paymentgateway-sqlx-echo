package handlerstats

import (
	"context"

	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card/stats"

	"github.com/MamangRust/monolith-payment-gateway-card/service"
	cardstatsservice "github.com/MamangRust/monolith-payment-gateway-card/service/stats"
	cardstatsbycard "github.com/MamangRust/monolith-payment-gateway-card/service/statsbycard"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/grpc"
)

type cardStatsTransactionGrpc struct {
	pb.UnimplementedCardStatsTransactionServiceServer

	cardStatsTransaction cardstatsservice.CardStatsTransactionService

	cardStatsTransactionByCard cardstatsbycard.CardStatsTransactionByCardService
}

func NewCardStatsTransactionGrpc(service service.Service) CardStatsTransactionService {
	return &cardStatsTransactionGrpc{
		cardStatsTransaction:       service,
		cardStatsTransactionByCard: service,
	}
}

func (s *cardStatsTransactionGrpc) FindMonthlyTransactionAmount(ctx context.Context, req *pbcard.FindYearAmount) (*pbcard.ApiResponseMonthlyAmount, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, card_errors.ErrGrpcInvalidYear
	}

	res, err := s.cardStatsTransaction.FindMonthlyTransactionAmount(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseMonthlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseMonthlyAmount{
			Month:       item.Month,
			TotalAmount: int64(item.TotalTransactionAmount),
		}
	}

	return &pbcard.ApiResponseMonthlyAmount{
		Status:  "success",
		Message: "Monthly transaction amount retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsTransactionGrpc) FindYearlyTransactionAmount(ctx context.Context, req *pbcard.FindYearAmount) (*pbcard.ApiResponseYearlyAmount, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, card_errors.ErrGrpcInvalidYear
	}

	res, err := s.cardStatsTransaction.FindYearlyTransactionAmount(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseYearlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseYearlyAmount{
			Year:        item.Year.Int.String(),
			TotalAmount: item.TotalTransactionAmount,
		}
	}

	return &pbcard.ApiResponseYearlyAmount{
		Status:  "success",
		Message: "Yearly transaction amount retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsTransactionGrpc) FindMonthlyTransactionAmountByCardNumber(ctx context.Context, req *pbcard.FindYearAmountCardNumber) (*pbcard.ApiResponseMonthlyAmount, error) {
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

	res, err := s.cardStatsTransactionByCard.FindMonthlyTransactionAmountByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseMonthlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseMonthlyAmount{
			Month:       item.Month,
			TotalAmount: int64(item.TotalTransactionAmount),
		}
	}

	return &pbcard.ApiResponseMonthlyAmount{
		Status:  "success",
		Message: "Monthly transaction amount by card number retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsTransactionGrpc) FindYearlyTransactionAmountByCardNumber(ctx context.Context, req *pbcard.FindYearAmountCardNumber) (*pbcard.ApiResponseYearlyAmount, error) {
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

	res, err := s.cardStatsTransactionByCard.FindYearlyTransactionAmountByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseYearlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseYearlyAmount{
			Year:        item.Year.Int.String(),
			TotalAmount: item.TotalTransactionAmount,
		}
	}

	return &pbcard.ApiResponseYearlyAmount{
		Status:  "success",
		Message: "Yearly transaction amount by card number retrieved successfully",
		Data:    protoData,
	}, nil
}
