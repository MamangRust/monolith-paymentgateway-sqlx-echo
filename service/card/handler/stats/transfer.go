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

type cardStatsTransferGrpc struct {
	pb.UnimplementedCardStatsTransferServiceServer

	cardStatsTransfer cardstatsservice.CardStatsTransferService

	cardStatsTransferByCard cardstatsbycard.CardStatsTransferByCardService
}

func NewCardStatsTransferGrpc(service service.Service) CardStatsTransferService {
	return &cardStatsTransferGrpc{
		cardStatsTransfer:       service,
		cardStatsTransferByCard: service,
	}
}

func (s *cardStatsTransferGrpc) FindMonthlyTransferSenderAmount(ctx context.Context, req *pbcard.FindYearAmount) (*pbcard.ApiResponseMonthlyAmount, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, card_errors.ErrGrpcInvalidYear
	}

	res, err := s.cardStatsTransfer.FindMonthlyTransferAmountSender(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseMonthlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseMonthlyAmount{
			Month:       item.Month,
			TotalAmount: int64(item.TotalSentAmount),
		}
	}

	return &pbcard.ApiResponseMonthlyAmount{
		Status:  "success",
		Message: "Monthly transfer sender amount retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsTransferGrpc) FindYearlyTransferSenderAmount(ctx context.Context, req *pbcard.FindYearAmount) (*pbcard.ApiResponseYearlyAmount, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, card_errors.ErrGrpcInvalidYear
	}

	res, err := s.cardStatsTransfer.FindYearlyTransferAmountSender(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseYearlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseYearlyAmount{
			Year:        item.Year.Int.String(),
			TotalAmount: item.TotalSentAmount,
		}
	}

	return &pbcard.ApiResponseYearlyAmount{
		Status:  "success",
		Message: "transfer sender amount retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsTransferGrpc) FindMonthlyTransferReceiverAmount(ctx context.Context, req *pbcard.FindYearAmount) (*pbcard.ApiResponseMonthlyAmount, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, card_errors.ErrGrpcInvalidYear
	}

	res, err := s.cardStatsTransfer.FindMonthlyTransferAmountReceiver(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseMonthlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseMonthlyAmount{
			Month:       item.Month,
			TotalAmount: int64(item.TotalReceivedAmount),
		}
	}

	return &pbcard.ApiResponseMonthlyAmount{
		Status:  "success",
		Message: "Monthly transfer receiver amount retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsTransferGrpc) FindYearlyTransferReceiverAmount(ctx context.Context, req *pbcard.FindYearAmount) (*pbcard.ApiResponseYearlyAmount, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, card_errors.ErrGrpcInvalidYear
	}

	res, err := s.cardStatsTransfer.FindYearlyTransferAmountReceiver(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseYearlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseYearlyAmount{
			Year:        item.Year.Int.String(),
			TotalAmount: item.TotalReceivedAmount,
		}
	}

	return &pbcard.ApiResponseYearlyAmount{
		Status:  "success",
		Message: "Yearly transfer receiver amount retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsTransferGrpc) FindMonthlyTransferSenderAmountByCardNumber(ctx context.Context, req *pbcard.FindYearAmountCardNumber) (*pbcard.ApiResponseMonthlyAmount, error) {
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

	res, err := s.cardStatsTransferByCard.FindMonthlyTransferAmountBySender(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseMonthlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseMonthlyAmount{
			Month:       item.Month,
			TotalAmount: int64(item.TotalSentAmount),
		}
	}

	return &pbcard.ApiResponseMonthlyAmount{
		Status:  "success",
		Message: "Monthly transfer sender amount by card number retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsTransferGrpc) FindYearlyTransferSenderAmountByCardNumber(ctx context.Context, req *pbcard.FindYearAmountCardNumber) (*pbcard.ApiResponseYearlyAmount, error) {
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

	res, err := s.cardStatsTransferByCard.FindYearlyTransferAmountBySender(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseYearlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseYearlyAmount{
			Year:        item.Year.Int.String(),
			TotalAmount: item.TotalSentAmount,
		}
	}

	return &pbcard.ApiResponseYearlyAmount{
		Status:  "success",
		Message: "Yearly transfer sender amount by card number retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsTransferGrpc) FindMonthlyTransferReceiverAmountByCardNumber(ctx context.Context, req *pbcard.FindYearAmountCardNumber) (*pbcard.ApiResponseMonthlyAmount, error) {
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

	res, err := s.cardStatsTransferByCard.FindMonthlyTransferAmountByReceiver(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseMonthlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseMonthlyAmount{
			Month:       item.Month,
			TotalAmount: int64(item.TotalReceivedAmount),
		}
	}

	return &pbcard.ApiResponseMonthlyAmount{
		Status:  "success",
		Message: "Monthly transfer receiver amount by card number retrieved successfully",
		Data:    protoData,
	}, nil
}

func (s *cardStatsTransferGrpc) FindYearlyTransferReceiverAmountByCardNumber(ctx context.Context, req *pbcard.FindYearAmountCardNumber) (*pbcard.ApiResponseYearlyAmount, error) {
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

	res, err := s.cardStatsTransferByCard.FindYearlyTransferAmountByReceiver(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pbcard.CardResponseYearlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pbcard.CardResponseYearlyAmount{
			Year:        item.Year.Int.String(),
			TotalAmount: item.TotalReceivedAmount,
		}
	}

	return &pbcard.ApiResponseYearlyAmount{
		Status:  "success",
		Message: "Yearly transfer receiver amount by card number retrieved successfully",
		Data:    protoData,
	}, nil
}
