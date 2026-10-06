package transferstatshandler

import (
	"context"

	pbtransfer "github.com/MamangRust/monolith-payment-gateway-pb/transfer"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transfer/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	transfer_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transfer_errors/grpc"
	"github.com/MamangRust/monolith-payment-gateway-transfer/service"
)

type transferStatsAmountHandleGrpc struct {
	pb.UnimplementedTransferStatsAmountServiceServer

	service service.Service
}

func NewTransferStatsAmountHandler(service service.Service) TransferStatsAmountHandleGrpc {
	return &transferStatsAmountHandleGrpc{
		service: service,
	}
}

func (s *transferStatsAmountHandleGrpc) FindMonthlyTransferAmounts(ctx context.Context, req *pbtransfer.FindYearTransferStatus) (*pb.ApiResponseTransferMonthAmount, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidYear
	}

	amounts, err := s.service.FindMonthlyTransferAmounts(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransferMonthAmountResponse, len(amounts))
	for i, amount := range amounts {
		dataResponses[i] = &pb.TransferMonthAmountResponse{
			Month:       amount.Month,
			TotalAmount: int32(amount.TotalTransferAmount),
		}
	}

	return &pb.ApiResponseTransferMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly transfer amounts",
		Data:    dataResponses,
	}, nil
}

func (s *transferStatsAmountHandleGrpc) FindYearlyTransferAmounts(ctx context.Context, req *pbtransfer.FindYearTransferStatus) (*pb.ApiResponseTransferYearAmount, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidYear
	}

	amounts, err := s.service.FindYearlyTransferAmounts(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransferYearAmountResponse, len(amounts))
	for i, amount := range amounts {
		dataResponses[i] = &pb.TransferYearAmountResponse{
			Year:        amount.Year.Int.String(),
			TotalAmount: int32(amount.TotalTransferAmount),
		}
	}

	return &pb.ApiResponseTransferYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly transfer amounts",
		Data:    dataResponses,
	}, nil
}

func (s *transferStatsAmountHandleGrpc) FindMonthlyTransferAmountsBySenderCardNumber(ctx context.Context, req *pbtransfer.FindByCardNumberTransferRequest) (*pb.ApiResponseTransferMonthAmount, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidYear
	}

	if cardNumber == "" {
		return nil, transfer_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthYearCardNumber{
		Year:       year,
		CardNumber: cardNumber,
	}

	amounts, err := s.service.FindMonthlyTransferAmountsBySenderCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransferMonthAmountResponse, len(amounts))
	for i, amount := range amounts {
		dataResponses[i] = &pb.TransferMonthAmountResponse{
			Month:       amount.Month,
			TotalAmount: int32(amount.TotalTransferAmount),
		}
	}

	return &pb.ApiResponseTransferMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly transfer amounts by sender card number",
		Data:    dataResponses,
	}, nil
}

func (s *transferStatsAmountHandleGrpc) FindMonthlyTransferAmountsByReceiverCardNumber(ctx context.Context, req *pbtransfer.FindByCardNumberTransferRequest) (*pb.ApiResponseTransferMonthAmount, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidYear
	}

	if cardNumber == "" {
		return nil, transfer_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthYearCardNumber{
		Year:       year,
		CardNumber: cardNumber,
	}

	amounts, err := s.service.FindMonthlyTransferAmountsByReceiverCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransferMonthAmountResponse, len(amounts))
	for i, amount := range amounts {
		dataResponses[i] = &pb.TransferMonthAmountResponse{
			Month:       amount.Month,
			TotalAmount: int32(amount.TotalTransferAmount),
		}
	}

	return &pb.ApiResponseTransferMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly transfer amounts by receiver card number",
		Data:    dataResponses,
	}, nil
}

func (s *transferStatsAmountHandleGrpc) FindYearlyTransferAmountsBySenderCardNumber(ctx context.Context, req *pbtransfer.FindByCardNumberTransferRequest) (*pb.ApiResponseTransferYearAmount, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidYear
	}

	if cardNumber == "" {
		return nil, transfer_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthYearCardNumber{
		Year:       year,
		CardNumber: cardNumber,
	}

	amounts, err := s.service.FindYearlyTransferAmountsBySenderCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransferYearAmountResponse, len(amounts))
	for i, amount := range amounts {
		dataResponses[i] = &pb.TransferYearAmountResponse{
			Year:        amount.Year.Int.String(),
			TotalAmount: int32(amount.TotalTransferAmount),
		}
	}

	return &pb.ApiResponseTransferYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly transfer amounts by sender card number",
		Data:    dataResponses,
	}, nil
}

func (s *transferStatsAmountHandleGrpc) FindYearlyTransferAmountsByReceiverCardNumber(ctx context.Context, req *pbtransfer.FindByCardNumberTransferRequest) (*pb.ApiResponseTransferYearAmount, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidYear
	}

	if cardNumber == "" {
		return nil, transfer_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthYearCardNumber{
		Year:       year,
		CardNumber: cardNumber,
	}

	amounts, err := s.service.FindYearlyTransferAmountsByReceiverCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransferYearAmountResponse, len(amounts))
	for i, amount := range amounts {
		dataResponses[i] = &pb.TransferYearAmountResponse{
			Year:        amount.Year.Int.String(),
			TotalAmount: int32(amount.TotalTransferAmount),
		}
	}

	return &pb.ApiResponseTransferYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly transfer amounts by receiver card number",
		Data:    dataResponses,
	}, nil
}
