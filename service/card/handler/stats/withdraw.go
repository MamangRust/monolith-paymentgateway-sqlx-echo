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

type cardStatsWithdrawGrpc struct {
	pb.UnimplementedCardStatsWithdrawServiceServer

	cardStatsWithdraw cardstatsservice.CardStatsWithdrawService

	cardStatsWithdrawByCard cardstatsbycard.CardStatsWithdrawByCardService
}

func NewCardStatsWithdrawGrpc(service service.Service) CardStatsWithdrawService {
	return &cardStatsWithdrawGrpc{
		cardStatsWithdraw:       service,
		cardStatsWithdrawByCard: service,
	}
}

func (s *cardStatsWithdrawGrpc) FindMonthlyWithdrawAmount(ctx context.Context, req *pbcard.FindYearAmount) (*pbcard.ApiResponseMonthlyAmount, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, card_errors.ErrGrpcInvalidYear
	}

	res, err := s.cardStatsWithdraw.FindMonthlyWithdrawAmount(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseMonthlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseMonthlyAmount{
			Month:       item.Month,
			TotalAmount: int64(item.TotalWithdrawAmount),
		}
	}

	return &pbcard.ApiResponseMonthlyAmount{
		Status:  "success",
		Message: "Monthly withdraw amount retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsWithdrawGrpc) FindYearlyWithdrawAmount(ctx context.Context, req *pbcard.FindYearAmount) (*pbcard.ApiResponseYearlyAmount, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, card_errors.ErrGrpcInvalidYear
	}

	res, err := s.cardStatsWithdraw.FindYearlyWithdrawAmount(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseYearlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseYearlyAmount{
			Year:        item.Year.Int.String(),
			TotalAmount: item.TotalWithdrawAmount,
		}
	}

	return &pbcard.ApiResponseYearlyAmount{
		Status:  "success",
		Message: "Yearly withdraw amount retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsWithdrawGrpc) FindMonthlyWithdrawAmountByCardNumber(ctx context.Context, req *pbcard.FindYearAmountCardNumber) (*pbcard.ApiResponseMonthlyAmount, error) {
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

	res, err := s.cardStatsWithdrawByCard.FindMonthlyWithdrawAmountByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseMonthlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseMonthlyAmount{
			Month:       item.Month,
			TotalAmount: int64(item.TotalWithdrawAmount),
		}
	}

	return &pbcard.ApiResponseMonthlyAmount{
		Status:  "success",
		Message: "Monthly withdraw amount by card number retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsWithdrawGrpc) FindYearlyWithdrawAmountByCardNumber(ctx context.Context, req *pbcard.FindYearAmountCardNumber) (*pbcard.ApiResponseYearlyAmount, error) {
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

	res, err := s.cardStatsWithdrawByCard.FindYearlyWithdrawAmountByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseYearlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseYearlyAmount{
			Year:        item.Year.Int.String(),
			TotalAmount: item.TotalWithdrawAmount,
		}
	}

	return &pbcard.ApiResponseYearlyAmount{
		Status:  "success",
		Message: "Yearly withdraw amount by card number retrieved successfully",
		Data:    protoData,
	}, nil
}
