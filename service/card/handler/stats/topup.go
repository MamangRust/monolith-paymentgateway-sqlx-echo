package handlerstats

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-card/service"
	cardstatsservice "github.com/MamangRust/monolith-payment-gateway-card/service/stats"
	cardstatsbycard "github.com/MamangRust/monolith-payment-gateway-card/service/statsbycard"
	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/grpc"
)

type cardStatsTopupGrpc struct {
	pb.UnimplementedCardStatsTopupServiceServer

	cardStatsTopup cardstatsservice.CardStatsTopupService

	cardStatsTopupByCard cardstatsbycard.CardStatsTopupByCardService
}

func NewCardStatsTopupGrpc(service service.Service) CardStatsTopupService {
	return &cardStatsTopupGrpc{
		cardStatsTopup:       service,
		cardStatsTopupByCard: service,
	}
}

func (s *cardStatsTopupGrpc) FindMonthlyTopupAmount(ctx context.Context, req *pbcard.FindYearAmount) (*pbcard.ApiResponseMonthlyAmount, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, card_errors.ErrGrpcInvalidYear
	}

	res, err := s.cardStatsTopup.FindMonthlyTopupAmount(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseMonthlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseMonthlyAmount{
			Month:       item.Month,
			TotalAmount: int64(item.TotalTopupAmount),
		}
	}

	return &pbcard.ApiResponseMonthlyAmount{
		Status:  "success",
		Message: "Monthly topup amount retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsTopupGrpc) FindYearlyTopupAmount(ctx context.Context, req *pbcard.FindYearAmount) (*pbcard.ApiResponseYearlyAmount, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, card_errors.ErrGrpcInvalidYear
	}

	res, err := s.cardStatsTopup.FindYearlyTopupAmount(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseYearlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseYearlyAmount{
			Year:        item.Year.Int.String(),
			TotalAmount: item.TotalTopupAmount,
		}
	}

	return &pbcard.ApiResponseYearlyAmount{
		Status:  "success",
		Message: "Yearly topup amount retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsTopupGrpc) FindMonthlyTopupAmountByCardNumber(ctx context.Context, req *pbcard.FindYearAmountCardNumber) (*pbcard.ApiResponseMonthlyAmount, error) {
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

	res, err := s.cardStatsTopupByCard.FindMonthlyTopupAmountByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseMonthlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseMonthlyAmount{
			Month:       item.Month,
			TotalAmount: int64(item.TotalTopupAmount),
		}
	}

	return &pbcard.ApiResponseMonthlyAmount{
		Status:  "success",
		Message: "Monthly topup amount by card number retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsTopupGrpc) FindYearlyTopupAmountByCardNumber(ctx context.Context, req *pbcard.FindYearAmountCardNumber) (*pbcard.ApiResponseYearlyAmount, error) {
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

	res, err := s.cardStatsTopupByCard.FindYearlyTopupAmountByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseYearlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseYearlyAmount{
			Year:        item.Year.Int.String(),
			TotalAmount: item.TotalTopupAmount,
		}
	}

	return &pbcard.ApiResponseYearlyAmount{
		Status:  "success",
		Message: "Yearly topup amount by card number retrieved successfully",
		Data:    protoData,
	}, nil
}
