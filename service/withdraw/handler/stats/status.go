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

type withdrawStatusHandleGrpc struct {
	pb.UnimplementedWithdrawStatsStatusServiceServer

	service service.Service
}

func NewWithdrawStatsStatusHandleGrpc(
	service service.Service,
) WithdrawStatsStatusHandleGrpc {
	return &withdrawStatusHandleGrpc{
		service: service,
	}
}

func (s *withdrawStatusHandleGrpc) FindMonthlyWithdrawStatusSuccess(ctx context.Context, req *pbwithdraw.FindMonthlyWithdrawStatus) (*pb.ApiResponseWithdrawMonthStatusSuccess, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidYear
	}

	if month <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthStatusWithdraw{
		Year:  year,
		Month: month,
	}

	records, err := s.service.FindMonthWithdrawStatusSuccess(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.WithdrawMonthStatusSuccessResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.WithdrawMonthStatusSuccessResponse{
			Year:         record.Year,
			Month:        record.Month,
			TotalSuccess: int32(record.TotalSuccess),
			TotalAmount:  int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseWithdrawMonthStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched withdraw",
		Data:    dataResponses,
	}, nil
}

func (s *withdrawStatusHandleGrpc) FindYearlyWithdrawStatusSuccess(ctx context.Context, req *pbwithdraw.FindYearWithdrawStatus) (*pb.ApiResponseWithdrawYearStatusSuccess, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidYear
	}

	records, err := s.service.FindYearlyWithdrawStatusSuccess(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.WithdrawYearStatusSuccessResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.WithdrawYearStatusSuccessResponse{
			Year:         record.Year,
			TotalSuccess: int32(record.TotalSuccess),
			TotalAmount:  int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseWithdrawYearStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched yearly Withdraw status success",
		Data:    dataResponses,
	}, nil
}

func (s *withdrawStatusHandleGrpc) FindMonthlyWithdrawStatusFailed(ctx context.Context, req *pbwithdraw.FindMonthlyWithdrawStatus) (*pb.ApiResponseWithdrawMonthStatusFailed, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidYear
	}

	if month <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthStatusWithdraw{
		Year:  year,
		Month: month,
	}

	records, err := s.service.FindMonthWithdrawStatusFailed(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.WithdrawMonthStatusFailedResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.WithdrawMonthStatusFailedResponse{
			Year:        record.Year,
			Month:       record.Month,
			TotalFailed: int32(record.TotalFailed),
			TotalAmount: int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseWithdrawMonthStatusFailed{
		Status:  "success",
		Message: "success fetched monthly Withdraw status Failed",
		Data:    dataResponses,
	}, nil
}

func (s *withdrawStatusHandleGrpc) FindYearlyWithdrawStatusFailed(ctx context.Context, req *pbwithdraw.FindYearWithdrawStatus) (*pb.ApiResponseWithdrawYearStatusFailed, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidYear
	}

	records, err := s.service.FindYearlyWithdrawStatusFailed(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.WithdrawYearStatusFailedResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.WithdrawYearStatusFailedResponse{
			Year:        record.Year,
			TotalFailed: int32(record.TotalFailed),
			TotalAmount: int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseWithdrawYearStatusFailed{
		Status:  "success",
		Message: "success fetched yearly Withdraw status Failed",
		Data:    dataResponses,
	}, nil
}

func (s *withdrawStatusHandleGrpc) FindMonthlyWithdrawStatusSuccessCardNumber(ctx context.Context, req *pbwithdraw.FindMonthlyWithdrawStatusCardNumber) (*pb.ApiResponseWithdrawMonthStatusSuccess, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidYear
	}

	if month <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidMonth
	}

	if cardNumber == "" {
		return nil, withdraw_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthStatusWithdrawCardNumber{
		Year:       year,
		Month:      month,
		CardNumber: cardNumber,
	}

	records, err := s.service.FindMonthWithdrawStatusSuccessByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.WithdrawMonthStatusSuccessResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.WithdrawMonthStatusSuccessResponse{
			Year:         record.Year,
			Month:        record.Month,
			TotalSuccess: int32(record.TotalSuccess),
			TotalAmount:  int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseWithdrawMonthStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched withdraw",
		Data:    dataResponses,
	}, nil
}

func (s *withdrawStatusHandleGrpc) FindYearlyWithdrawStatusSuccessCardNumber(ctx context.Context, req *pbwithdraw.FindYearWithdrawStatusCardNumber) (*pb.ApiResponseWithdrawYearStatusSuccess, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidYear
	}

	if cardNumber == "" {
		return nil, withdraw_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.YearStatusWithdrawCardNumber{
		Year:       year,
		CardNumber: cardNumber,
	}

	records, err := s.service.FindYearlyWithdrawStatusSuccessByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.WithdrawYearStatusSuccessResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.WithdrawYearStatusSuccessResponse{
			Year:         record.Year,
			TotalSuccess: int32(record.TotalSuccess),
			TotalAmount:  int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseWithdrawYearStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched yearly Withdraw status success",
		Data:    dataResponses,
	}, nil
}

func (s *withdrawStatusHandleGrpc) FindMonthlyWithdrawStatusFailedCardNumber(ctx context.Context, req *pbwithdraw.FindMonthlyWithdrawStatusCardNumber) (*pb.ApiResponseWithdrawMonthStatusFailed, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidYear
	}

	if month <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidMonth
	}

	if cardNumber == "" {
		return nil, withdraw_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthStatusWithdrawCardNumber{
		Year:       year,
		Month:      month,
		CardNumber: cardNumber,
	}

	records, err := s.service.FindMonthWithdrawStatusFailedByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.WithdrawMonthStatusFailedResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.WithdrawMonthStatusFailedResponse{
			Year:        record.Year,
			Month:       record.Month,
			TotalFailed: int32(record.TotalFailed),
			TotalAmount: int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseWithdrawMonthStatusFailed{
		Status:  "success",
		Message: "Successfully fetched monthly Withdraw status failed",
		Data:    dataResponses,
	}, nil
}

func (s *withdrawStatusHandleGrpc) FindYearlyWithdrawStatusFailedCardNumber(ctx context.Context, req *pbwithdraw.FindYearWithdrawStatusCardNumber) (*pb.ApiResponseWithdrawYearStatusFailed, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, withdraw_errors.ErrGrpcInvalidYear
	}

	reqService := requests.YearStatusWithdrawCardNumber{
		Year:       year,
		CardNumber: cardNumber,
	}

	records, err := s.service.FindYearlyWithdrawStatusFailedByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	dataResponses := make([]*pb.WithdrawYearStatusFailedResponse, len(records))
	for i, record := range records {
		dataResponses[i] = &pb.WithdrawYearStatusFailedResponse{
			Year:        record.Year,
			TotalFailed: int32(record.TotalFailed),
			TotalAmount: int32(record.TotalAmount),
		}
	}

	return &pb.ApiResponseWithdrawYearStatusFailed{
		Status:  "success",
		Message: "Successfully fetched yearly Withdraw status failed",
		Data:    dataResponses,
	}, nil
}
