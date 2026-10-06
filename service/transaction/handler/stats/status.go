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

type transactionStatsStatusHandleGrpc struct {
	pb.UnimplementedTransactionStatsStatusServiceServer

	service service.Service
}

func NewTransactionStatsStatusHandleGrpc(
	service service.Service,
) TransactionStatsStatusHandleGrpc {
	return &transactionStatsStatusHandleGrpc{
		service: service,
	}
}

func (s *transactionStatsStatusHandleGrpc) FindMonthlyTransactionStatusSuccess(ctx context.Context, req *pbtransaction.FindMonthlyTransactionStatus) (*pb.ApiResponseTransactionMonthStatusSuccess, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if month <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthStatusTransaction{
		Year:  year,
		Month: month,
	}

	records, err := s.service.FindMonthTransactionStatusSuccess(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionMonthStatusSuccessResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransactionMonthStatusSuccessResponse{
			Year:         record.Year,
			Month:        record.Month,
			TotalSuccess: int32(record.TotalSuccess),
			TotalAmount:  int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionMonthStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched monthly Transaction status success",
		Data:    dataResponses,
	}, nil
}

func (s *transactionStatsStatusHandleGrpc) FindYearlyTransactionStatusSuccess(ctx context.Context, req *pbtransaction.FindYearTransactionStatus) (*pb.ApiResponseTransactionYearStatusSuccess, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	records, err := s.service.FindYearlyTransactionStatusSuccess(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionYearStatusSuccessResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransactionYearStatusSuccessResponse{
			Year:         record.Year,
			TotalSuccess: int32(record.TotalSuccess),
			TotalAmount:  int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionYearStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched yearly Transaction status success",
		Data:    dataResponses,
	}, nil
}

func (s *transactionStatsStatusHandleGrpc) FindMonthlyTransactionStatusFailed(ctx context.Context, req *pbtransaction.FindMonthlyTransactionStatus) (*pb.ApiResponseTransactionMonthStatusFailed, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if month <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthStatusTransaction{
		Year:  year,
		Month: month,
	}

	records, err := s.service.FindMonthTransactionStatusFailed(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionMonthStatusFailedResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransactionMonthStatusFailedResponse{
			Year:        record.Year,
			Month:       record.Month,
			TotalFailed: int32(record.TotalFailed),
			TotalAmount: int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionMonthStatusFailed{
		Status:  "success",
		Message: "success fetched monthly Transaction status Failed",
		Data:    dataResponses,
	}, nil
}

func (s *transactionStatsStatusHandleGrpc) FindYearlyTransactionStatusFailed(ctx context.Context, req *pbtransaction.FindYearTransactionStatus) (*pb.ApiResponseTransactionYearStatusFailed, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	records, err := s.service.FindYearlyTransactionStatusFailed(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionYearStatusFailedResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransactionYearStatusFailedResponse{
			Year:        record.Year,
			TotalFailed: int32(record.TotalFailed),
			TotalAmount: int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionYearStatusFailed{
		Status:  "success",
		Message: "success fetched yearly Transaction status Failed",
		Data:    dataResponses,
	}, nil
}

func (s *transactionStatsStatusHandleGrpc) FindMonthlyTransactionStatusSuccessByCardNumber(ctx context.Context, req *pbtransaction.FindMonthlyTransactionStatusCardNumber) (*pb.ApiResponseTransactionMonthStatusSuccess, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if month <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	if cardNumber == "" {
		return nil, transaction_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthStatusTransactionCardNumber{
		CardNumber: cardNumber,
		Year:       year,
		Month:      month,
	}

	records, err := s.service.FindMonthTransactionStatusSuccessByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionMonthStatusSuccessResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransactionMonthStatusSuccessResponse{
			Year:         record.Year,
			Month:        record.Month,
			TotalSuccess: int32(record.TotalSuccess),
			TotalAmount:  int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionMonthStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched monthly Transaction status success",
		Data:    dataResponses,
	}, nil
}

func (s *transactionStatsStatusHandleGrpc) FindYearlyTransactionStatusSuccessByCardNumber(ctx context.Context, req *pbtransaction.FindYearTransactionStatusCardNumber) (*pb.ApiResponseTransactionYearStatusSuccess, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if cardNumber == "" {
		return nil, transaction_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.YearStatusTransactionCardNumber{
		Year:       year,
		CardNumber: cardNumber,
	}

	records, err := s.service.FindYearlyTransactionStatusSuccessByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionYearStatusSuccessResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransactionYearStatusSuccessResponse{
			Year:         record.Year,
			TotalSuccess: int32(record.TotalSuccess),
			TotalAmount:  int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionYearStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched yearly Transaction status success",
		Data:    dataResponses,
	}, nil
}

func (s *transactionStatsStatusHandleGrpc) FindMonthlyTransactionStatusFailedByCardNumber(ctx context.Context, req *pbtransaction.FindMonthlyTransactionStatusCardNumber) (*pb.ApiResponseTransactionMonthStatusFailed, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if month <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	if cardNumber == "" {
		return nil, transaction_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthStatusTransactionCardNumber{
		CardNumber: cardNumber,
		Year:       year,
		Month:      month,
	}

	records, err := s.service.FindMonthTransactionStatusFailedByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionMonthStatusFailedResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransactionMonthStatusFailedResponse{
			Year:        record.Year,
			Month:       record.Month,
			TotalFailed: int32(record.TotalFailed),
			TotalAmount: int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionMonthStatusFailed{
		Status:  "success",
		Message: "success fetched monthly Transaction status Failed",
		Data:    dataResponses,
	}, nil
}

func (s *transactionStatsStatusHandleGrpc) FindYearlyTransactionStatusFailedByCardNumber(ctx context.Context, req *pbtransaction.FindYearTransactionStatusCardNumber) (*pb.ApiResponseTransactionYearStatusFailed, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if cardNumber == "" {
		return nil, transaction_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.YearStatusTransactionCardNumber{
		Year:       year,
		CardNumber: cardNumber,
	}

	records, err := s.service.FindYearlyTransactionStatusFailedByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransactionYearStatusFailedResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransactionYearStatusFailedResponse{
			Year:        record.Year,
			TotalFailed: int32(record.TotalFailed),
			TotalAmount: int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransactionYearStatusFailed{
		Status:  "success",
		Message: "success fetched yearly Transaction status Failed",
		Data:    dataResponses,
	}, nil
}
