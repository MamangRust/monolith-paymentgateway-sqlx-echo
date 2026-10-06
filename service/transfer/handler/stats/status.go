package transferstatshandler

import (
	"context"

	pb "github.com/MamangRust/monolith-payment-gateway-pb/transfer/stats"

	pbtransfer "github.com/MamangRust/monolith-payment-gateway-pb/transfer"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	transfer_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transfer_errors/grpc"
	"github.com/MamangRust/monolith-payment-gateway-transfer/service"
)

type transferStatsStatusHandleGrpc struct {
	pb.UnimplementedTransferStatsStatusServiceServer

	service service.Service
}

func NewTransferStatsStatusHandler(service service.Service) TransferStatsStatusHandleGrpc {
	return &transferStatsStatusHandleGrpc{
		service: service,
	}
}

func (s *transferStatsStatusHandleGrpc) FindMonthlyTransferStatusSuccess(ctx context.Context, req *pbtransfer.FindMonthlyTransferStatus) (*pb.ApiResponseTransferMonthStatusSuccess, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidYear
	}

	if month <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthStatusTransfer{
		Year:  year,
		Month: month,
	}

	records, err := s.service.FindMonthTransferStatusSuccess(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransferMonthStatusSuccessResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransferMonthStatusSuccessResponse{
			Year:         record.Year,
			Month:        record.Month,
			TotalSuccess: int32(record.TotalSuccess),
			TotalAmount:  int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransferMonthStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched monthly Transfer status success",
		Data:    dataResponses,
	}, nil
}

func (s *transferStatsStatusHandleGrpc) FindYearlyTransferStatusSuccess(ctx context.Context, req *pbtransfer.FindYearTransferStatus) (*pb.ApiResponseTransferYearStatusSuccess, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidYear
	}

	records, err := s.service.FindYearlyTransferStatusSuccess(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransferYearStatusSuccessResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransferYearStatusSuccessResponse{
			Year:         record.Year,
			TotalSuccess: int32(record.TotalSuccess),
			TotalAmount:  int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransferYearStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched yearly Transfer status success",
		Data:    dataResponses,
	}, nil
}

func (s *transferStatsStatusHandleGrpc) FindMonthlyTransferStatusFailed(ctx context.Context, req *pbtransfer.FindMonthlyTransferStatus) (*pb.ApiResponseTransferMonthStatusFailed, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidYear
	}

	if month <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthStatusTransfer{
		Year:  year,
		Month: month,
	}

	records, err := s.service.FindMonthTransferStatusFailed(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransferMonthStatusFailedResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransferMonthStatusFailedResponse{
			Year:        record.Year,
			Month:       record.Month,
			TotalFailed: int32(record.TotalFailed),
			TotalAmount: int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransferMonthStatusFailed{
		Status:  "success",
		Message: "success fetched monthly Transfer status Failed",
		Data:    dataResponses,
	}, nil
}

func (s *transferStatsStatusHandleGrpc) FindYearlyTransferStatusFailed(ctx context.Context, req *pbtransfer.FindYearTransferStatus) (*pb.ApiResponseTransferYearStatusFailed, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidYear
	}

	records, err := s.service.FindYearlyTransferStatusFailed(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransferYearStatusFailedResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransferYearStatusFailedResponse{
			Year:        record.Year,
			TotalFailed: int32(record.TotalFailed),
			TotalAmount: int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransferYearStatusFailed{
		Status:  "success",
		Message: "success fetched yearly Transfer status Failed",
		Data:    dataResponses,
	}, nil
}

func (s *transferStatsStatusHandleGrpc) FindMonthlyTransferStatusSuccessByCardNumber(ctx context.Context, req *pbtransfer.FindMonthlyTransferStatusCardNumber) (*pb.ApiResponseTransferMonthStatusSuccess, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidYear
	}

	if month <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidMonth
	}

	if cardNumber == "" {
		return nil, transfer_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthStatusTransferCardNumber{
		Year:       year,
		Month:      month,
		CardNumber: cardNumber,
	}

	records, err := s.service.FindMonthTransferStatusSuccessByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransferMonthStatusSuccessResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransferMonthStatusSuccessResponse{
			Year:         record.Year,
			Month:        record.Month,
			TotalSuccess: int32(record.TotalSuccess),
			TotalAmount:  int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransferMonthStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched monthly Transfer status success",
		Data:    dataResponses,
	}, nil
}

func (s *transferStatsStatusHandleGrpc) FindYearlyTransferStatusSuccessByCardNumber(ctx context.Context, req *pbtransfer.FindYearTransferStatusCardNumber) (*pb.ApiResponseTransferYearStatusSuccess, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidYear
	}

	if cardNumber == "" {
		return nil, transfer_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.YearStatusTransferCardNumber{
		Year:       year,
		CardNumber: cardNumber,
	}

	records, err := s.service.FindYearlyTransferStatusSuccessByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransferYearStatusSuccessResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransferYearStatusSuccessResponse{
			Year:         record.Year,
			TotalSuccess: int32(record.TotalSuccess),
			TotalAmount:  int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransferYearStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched yearly Transfer status success",
		Data:    dataResponses,
	}, nil
}

func (s *transferStatsStatusHandleGrpc) FindMonthlyTransferStatusFailedByCardNumber(ctx context.Context, req *pbtransfer.FindMonthlyTransferStatusCardNumber) (*pb.ApiResponseTransferMonthStatusFailed, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidYear
	}

	if month <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidMonth
	}

	if cardNumber == "" {
		return nil, transfer_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthStatusTransferCardNumber{
		Year:       year,
		Month:      month,
		CardNumber: cardNumber,
	}

	records, err := s.service.FindMonthTransferStatusFailedByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransferMonthStatusFailedResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransferMonthStatusFailedResponse{
			Year:        record.Year,
			Month:       record.Month,
			TotalFailed: int32(record.TotalFailed),
			TotalAmount: int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransferMonthStatusFailed{
		Status:  "success",
		Message: "success fetched monthly Transfer status Failed",
		Data:    dataResponses,
	}, nil
}

func (s *transferStatsStatusHandleGrpc) FindYearlyTransferStatusFailedByCardNumber(ctx context.Context, req *pbtransfer.FindYearTransferStatusCardNumber) (*pb.ApiResponseTransferYearStatusFailed, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, transfer_errors.ErrGrpcInvalidYear
	}

	if cardNumber == "" {
		return nil, transfer_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.YearStatusTransferCardNumber{
		Year:       year,
		CardNumber: cardNumber,
	}

	records, err := s.service.FindYearlyTransferStatusFailedByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.TransferYearStatusFailedResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.TransferYearStatusFailedResponse{
			Year:        record.Year,
			TotalFailed: int32(record.TotalFailed),
			TotalAmount: int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseTransferYearStatusFailed{
		Status:  "success",
		Message: "success fetched yearly Transfer status Failed",
		Data:    dataResponses,
	}, nil
}
