package transactionstatshandler

import (
	"context"

	pbtransaction "github.com/MamangRust/monolith-payment-gateway-pb/transaction"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/grpc"
	"github.com/MamangRust/monolith-payment-gateway-transaction/service"
)

type transactionStatsAmountHandleGrpc struct {
	pb.UnimplementedTransactionStatsAmountServiceServer

	service service.Service
}

func NewTransactionStatsAmountHandleGrpc(
	service service.Service,
) TransactionStatsAmountHandlerGrpc {
	return &transactionStatsAmountHandleGrpc{
		service: service,
	}
}

func (t *transactionStatsAmountHandleGrpc) FindMonthlyAmounts(ctx context.Context, req *pbtransaction.FindYearTransactionStatus) (*pb.ApiResponseTransactionMonthAmount, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	amounts, err := t.service.FindMonthlyAmounts(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionMonthAmountResponse, len(amounts))
	for i, amount := range amounts {
		dataResponses[i] = &pb.TransactionMonthAmountResponse{
			Month:       amount.Month,
			TotalAmount: int32(amount.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly amounts",
		Data:    dataResponses,
	}, nil
}

func (t *transactionStatsAmountHandleGrpc) FindYearlyAmounts(ctx context.Context, req *pbtransaction.FindYearTransactionStatus) (*pb.ApiResponseTransactionYearAmount, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	amounts, err := t.service.FindYearlyAmounts(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionYearlyAmountResponse, len(amounts))
	for i, amount := range amounts {
		dataResponses[i] = &pb.TransactionYearlyAmountResponse{
			Year:        amount.Year.Int.String(),
			TotalAmount: int32(amount.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly amounts",
		Data:    dataResponses,
	}, nil
}

func (t *transactionStatsAmountHandleGrpc) FindMonthlyAmountsByCardNumber(ctx context.Context, req *pbtransaction.FindByYearCardNumberTransactionRequest) (*pb.ApiResponseTransactionMonthAmount, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if cardNumber == "" {
		return nil, transaction_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthYearPaymentMethod{
		Year:       year,
		CardNumber: cardNumber,
	}

	amounts, err := t.service.FindMonthlyAmountsByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionMonthAmountResponse, len(amounts))
	for i, amount := range amounts {
		dataResponses[i] = &pb.TransactionMonthAmountResponse{
			Month:       amount.Month,
			TotalAmount: int32(amount.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionMonthAmount{
		Status:  "success",
		Message: "Successfully fetched monthly amounts by card number",
		Data:    dataResponses,
	}, nil
}

func (t *transactionStatsAmountHandleGrpc) FindYearlyAmountsByCardNumber(ctx context.Context, req *pbtransaction.FindByYearCardNumberTransactionRequest) (*pb.ApiResponseTransactionYearAmount, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if cardNumber == "" {
		return nil, transaction_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthYearPaymentMethod{
		Year:       year,
		CardNumber: cardNumber,
	}

	amounts, err := t.service.FindYearlyAmountsByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionYearlyAmountResponse, len(amounts))
	for i, amount := range amounts {
		dataResponses[i] = &pb.TransactionYearlyAmountResponse{
			Year:        amount.Year.Int.String(),
			TotalAmount: int32(amount.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionYearAmount{
		Status:  "success",
		Message: "Successfully fetched yearly amounts by card number",
		Data:    dataResponses,
	}, nil
}
