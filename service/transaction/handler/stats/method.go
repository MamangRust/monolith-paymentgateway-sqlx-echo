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

type transactionStatsMethodHandleGrpc struct {
	pb.UnimplementedTransactionStatsMethodServiceServer

	service service.Service
}

func NewTransactionStatsMethodHandleGrpc(
	service service.Service,
) TransactionStatsMethodHandleGrpc {
	return &transactionStatsMethodHandleGrpc{
		service: service,
	}
}

func (t *transactionStatsMethodHandleGrpc) FindMonthlyPaymentMethods(ctx context.Context, req *pbtransaction.FindYearTransactionStatus) (*pb.ApiResponseTransactionMonthMethod, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	methods, err := t.service.FindMonthlyPaymentMethods(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionMonthMethodResponse, len(methods))
	for i, method := range methods {
		dataResponses[i] = &pb.TransactionMonthMethodResponse{
			Month:             method.Month,
			PaymentMethod:     method.PaymentMethod,
			TotalTransactions: int32(method.TotalTransactions),
			TotalAmount:       int32(method.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionMonthMethod{
		Status:  "success",
		Message: "Successfully fetched monthly payment methods",
		Data:    dataResponses,
	}, nil
}

func (t *transactionStatsMethodHandleGrpc) FindYearlyPaymentMethods(ctx context.Context, req *pbtransaction.FindYearTransactionStatus) (*pb.ApiResponseTransactionYearMethod, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	methods, err := t.service.FindYearlyPaymentMethods(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionYearMethodResponse, len(methods))
	for i, method := range methods {
		dataResponses[i] = &pb.TransactionYearMethodResponse{
			Year:              method.Year.Int.String(),
			PaymentMethod:     method.PaymentMethod,
			TotalTransactions: int32(method.TotalTransactions),
			TotalAmount:       int32(method.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionYearMethod{
		Status:  "success",
		Message: "Successfully fetched yearly payment methods",
		Data:    dataResponses,
	}, nil
}

func (t *transactionStatsMethodHandleGrpc) FindMonthlyPaymentMethodsByCardNumber(ctx context.Context, req *pbtransaction.FindByYearCardNumberTransactionRequest) (*pb.ApiResponseTransactionMonthMethod, error) {
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

	methods, err := t.service.FindMonthlyPaymentMethodsByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionMonthMethodResponse, len(methods))
	for i, method := range methods {
		dataResponses[i] = &pb.TransactionMonthMethodResponse{
			Month:             method.Month,
			PaymentMethod:     method.PaymentMethod,
			TotalTransactions: int32(method.TotalTransactions),
			TotalAmount:       int32(method.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionMonthMethod{
		Status:  "success",
		Message: "Successfully fetched monthly payment methods by card number",
		Data:    dataResponses,
	}, nil
}

func (t *transactionStatsMethodHandleGrpc) FindYearlyPaymentMethodsByCardNumber(ctx context.Context, req *pbtransaction.FindByYearCardNumberTransactionRequest) (*pb.ApiResponseTransactionYearMethod, error) {
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

	methods, err := t.service.FindYearlyPaymentMethodsByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionYearMethodResponse, len(methods))
	for i, method := range methods {
		dataResponses[i] = &pb.TransactionYearMethodResponse{
			Year:              method.Year.Int.String(),
			PaymentMethod:     method.PaymentMethod,
			TotalTransactions: int32(method.TotalTransactions),
			TotalAmount:       int32(method.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionYearMethod{
		Status:  "success",
		Message: "Successfully fetched yearly payment methods by card number",
		Data:    dataResponses,
	}, nil
}
