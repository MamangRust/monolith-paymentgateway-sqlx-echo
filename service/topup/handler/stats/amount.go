package topupstatshandler

import (
	"context"

	pbtopup "github.com/MamangRust/monolith-payment-gateway-pb/topup"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/grpc"
	"github.com/MamangRust/monolith-payment-gateway-topup/service"
)

type topupStatsAmountHandleGrpc struct {
	pb.UnimplementedTopupStatsAmountServiceServer

	service service.Service
}

func NewTopupStatsAmountHandleGrpc(
	service service.Service,
) TopupStatsAmountHandleGrpc {
	return &topupStatsAmountHandleGrpc{
		service: service,
	}
}

func (s *topupStatsAmountHandleGrpc) FindMonthlyTopupAmounts(ctx context.Context, req *pbtopup.FindYearTopupStatus) (*pb.ApiResponseTopupMonthAmount, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidYear
	}

	amounts, err := s.service.FindMonthlyTopupAmounts(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupMonthAmountResponse, len(amounts))
	for i, item := range amounts {
		protoData[i] = &pb.TopupMonthAmountResponse{
			Month:       item.Month,
			TotalAmount: item.TotalAmount,
		}
	}

	return &pb.ApiResponseTopupMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly topup amounts",
		Data:    protoData,
	}, nil
}

func (s *topupStatsAmountHandleGrpc) FindYearlyTopupAmounts(ctx context.Context, req *pbtopup.FindYearTopupStatus) (*pb.ApiResponseTopupYearAmount, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidYear
	}

	amounts, err := s.service.FindYearlyTopupAmounts(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupYearlyAmountResponse, len(amounts))
	for i, item := range amounts {
		protoData[i] = &pb.TopupYearlyAmountResponse{
			Year:        item.Year.Int.String(),
			TotalAmount: int32(item.TotalAmount),
		}
	}

	return &pb.ApiResponseTopupYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly topup amounts",
		Data:    protoData,
	}, nil
}

func (s *topupStatsAmountHandleGrpc) FindMonthlyTopupAmountsByCardNumber(ctx context.Context, req *pbtopup.FindYearTopupCardNumber) (*pb.ApiResponseTopupMonthAmount, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidYear
	}

	if cardNumber == "" {
		return nil, topup_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.YearMonthMethod{
		Year:       year,
		CardNumber: cardNumber,
	}

	amounts, err := s.service.FindMonthlyTopupAmountsByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupMonthAmountResponse, len(amounts))
	for i, item := range amounts {
		protoData[i] = &pb.TopupMonthAmountResponse{
			Month:       item.Month,
			TotalAmount: item.TotalAmount,
		}
	}

	return &pb.ApiResponseTopupMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly topup amounts by card number",
		Data:    protoData,
	}, nil
}

func (s *topupStatsAmountHandleGrpc) FindYearlyTopupAmountsByCardNumber(ctx context.Context, req *pbtopup.FindYearTopupCardNumber) (*pb.ApiResponseTopupYearAmount, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidYear
	}

	if cardNumber == "" {
		return nil, topup_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.YearMonthMethod{
		Year:       year,
		CardNumber: cardNumber,
	}

	amounts, err := s.service.FindYearlyTopupAmountsByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupYearlyAmountResponse, len(amounts))
	for i, item := range amounts {
		protoData[i] = &pb.TopupYearlyAmountResponse{
			Year:        item.Year.Int.String(),
			TotalAmount: int32(item.TotalAmount),
		}
	}

	return &pb.ApiResponseTopupYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly topup amounts by card number",
		Data:    protoData,
	}, nil
}
