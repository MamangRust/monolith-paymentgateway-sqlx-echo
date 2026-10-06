package withdrawstatshandler

import (
	"context"

	pbwithdraw "github.com/MamangRust/monolith-payment-gateway-pb/withdraw"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	withdraw_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/withdraw_errors/grpc"
	service "github.com/MamangRust/monolith-payment-gateway-withdraw/service"
)

type withdrawAmountHandleGrpc struct {
	pb.UnimplementedWithdrawStatsAmountServiceServer

	service service.Service
}

func NewWithdrawStatsAmountHandleGrpc(
	service service.Service,
) WithdrawStatsAmountHandlerGrpc {
	return &withdrawAmountHandleGrpc{
		service: service,
	}
}

func (w *withdrawAmountHandleGrpc) FindMonthlyWithdraws(ctx context.Context, req *pbwithdraw.FindYearWithdrawStatus) (*pb.ApiResponseWithdrawMonthAmount, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidYear
	}

	withdraws, err := w.service.FindMonthlyWithdraws(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.WithdrawMonthlyAmountResponse, len(withdraws))
	for i, withdraw := range withdraws {
		dataResponses[i] = &pb.WithdrawMonthlyAmountResponse{
			Month:       withdraw.Month,
			TotalAmount: int32(withdraw.TotalWithdrawAmount),
		}
	}

	return &pb.ApiResponseWithdrawMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly withdraws",
		Data:    dataResponses,
	}, nil
}

func (w *withdrawAmountHandleGrpc) FindYearlyWithdraws(ctx context.Context, req *pbwithdraw.FindYearWithdrawStatus) (*pb.ApiResponseWithdrawYearAmount, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidYear
	}

	withdraws, err := w.service.FindYearlyWithdraws(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.WithdrawYearlyAmountResponse, len(withdraws))
	for i, withdraw := range withdraws {
		dataResponses[i] = &pb.WithdrawYearlyAmountResponse{
			Year:        withdraw.Year.Int.String(),
			TotalAmount: int32(withdraw.TotalWithdrawAmount),
		}
	}

	return &pb.ApiResponseWithdrawYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly withdraws",
		Data:    dataResponses,
	}, nil
}

func (w *withdrawAmountHandleGrpc) FindMonthlyWithdrawsByCardNumber(ctx context.Context, req *pbwithdraw.FindYearWithdrawCardNumber) (*pb.ApiResponseWithdrawMonthAmount, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidYear
	}

	if cardNumber == "" {
		return nil, withdraw_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.YearMonthCardNumber{
		Year:       year,
		CardNumber: cardNumber,
	}

	withdraws, err := w.service.FindMonthlyWithdrawsByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.WithdrawMonthlyAmountResponse, len(withdraws))
	for i, withdraw := range withdraws {
		dataResponses[i] = &pb.WithdrawMonthlyAmountResponse{
			Month:       withdraw.Month,
			TotalAmount: int32(withdraw.TotalWithdrawAmount),
		}
	}

	return &pb.ApiResponseWithdrawMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly withdraws by card number",
		Data:    dataResponses,
	}, nil
}

func (w *withdrawAmountHandleGrpc) FindYearlyWithdrawsByCardNumber(ctx context.Context, req *pbwithdraw.FindYearWithdrawCardNumber) (*pb.ApiResponseWithdrawYearAmount, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidYear
	}

	if cardNumber == "" {
		return nil, withdraw_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.YearMonthCardNumber{
		Year:       year,
		CardNumber: cardNumber,
	}

	withdraws, err := w.service.FindYearlyWithdrawsByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.WithdrawYearlyAmountResponse, len(withdraws))
	for i, withdraw := range withdraws {
		dataResponses[i] = &pb.WithdrawYearlyAmountResponse{
			Year:        withdraw.Year.Int.String(),
			TotalAmount: int32(withdraw.TotalWithdrawAmount),
		}
	}

	return &pb.ApiResponseWithdrawYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly withdraws by card number",
		Data:    dataResponses,
	}, nil
}
