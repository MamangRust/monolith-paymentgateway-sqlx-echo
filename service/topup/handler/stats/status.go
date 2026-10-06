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

type topupStatusHandleGrpc struct {
	pb.UnimplementedTopupStatsStatusServiceServer

	service service.Service
}

func NewTopupStatsStatusHandleGrpc(
	service service.Service,
) TopupStatsStatusHandleGrpc {
	return &topupStatusHandleGrpc{
		service: service,
	}
}

func (s *topupStatusHandleGrpc) FindMonthlyTopupStatusSuccess(ctx context.Context, req *pbtopup.FindMonthlyTopupStatus) (*pb.ApiResponseTopupMonthStatusSuccess, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidYear
	}
	if month <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidMonth
	}

	reqService := requests.MonthTopupStatus{
		Year:  year,
		Month: month,
	}

	records, err := s.service.FindMonthTopupStatusSuccess(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupMonthStatusSuccessResponse, len(records))
	for i, item := range records {
		protoData[i] = &pb.TopupMonthStatusSuccessResponse{
			Year:         item.Year,
			Month:        item.Month,
			TotalSuccess: int32(item.TotalSuccess),
			TotalAmount:  item.TotalAmount,
		}
	}

	return &pb.ApiResponseTopupMonthStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched monthly topup status success",
		Data:    protoData,
	}, nil
}

func (s *topupStatusHandleGrpc) FindYearlyTopupStatusSuccess(ctx context.Context, req *pbtopup.FindYearTopupStatus) (*pb.ApiResponseTopupYearStatusSuccess, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidYear
	}

	records, err := s.service.FindYearlyTopupStatusSuccess(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupYearStatusSuccessResponse, len(records))
	for i, item := range records {
		protoData[i] = &pb.TopupYearStatusSuccessResponse{
			Year:         item.Year,
			TotalSuccess: item.TotalSuccess,
			TotalAmount:  item.TotalAmount,
		}
	}

	return &pb.ApiResponseTopupYearStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched yearly topup status success",
		Data:    protoData,
	}, nil
}

func (s *topupStatusHandleGrpc) FindMonthlyTopupStatusFailed(ctx context.Context, req *pbtopup.FindMonthlyTopupStatus) (*pb.ApiResponseTopupMonthStatusFailed, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidYear
	}
	if month <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidMonth
	}

	reqService := requests.MonthTopupStatus{
		Year:  year,
		Month: month,
	}

	records, err := s.service.FindMonthTopupStatusFailed(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupMonthStatusFailedResponse, len(records))
	for i, item := range records {
		protoData[i] = &pb.TopupMonthStatusFailedResponse{
			Year:        item.Year,
			Month:       item.Month,
			TotalFailed: int32(item.TotalFailed),
			TotalAmount: item.TotalAmount,
		}
	}

	return &pb.ApiResponseTopupMonthStatusFailed{
		Status:  "success",
		Message: "Successfully fetched monthly topup status failed",
		Data:    protoData,
	}, nil
}

func (s *topupStatusHandleGrpc) FindYearlyTopupStatusFailed(ctx context.Context, req *pbtopup.FindYearTopupStatus) (*pb.ApiResponseTopupYearStatusFailed, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidYear
	}

	records, err := s.service.FindYearlyTopupStatusFailed(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupYearStatusFailedResponse, len(records))
	for i, item := range records {
		protoData[i] = &pb.TopupYearStatusFailedResponse{
			Year:        item.Year,
			TotalFailed: item.TotalFailed,
			TotalAmount: item.TotalAmount,
		}
	}

	return &pb.ApiResponseTopupYearStatusFailed{
		Status:  "success",
		Message: "Successfully fetched yearly topup status failed",
		Data:    protoData,
	}, nil
}

func (s *topupStatusHandleGrpc) FindMonthlyTopupStatusSuccessByCardNumber(ctx context.Context, req *pbtopup.FindMonthlyTopupStatusCardNumber) (*pb.ApiResponseTopupMonthStatusSuccess, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidYear
	}
	if month <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidMonth
	}
	if cardNumber == "" {
		return nil, topup_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthTopupStatusCardNumber{
		Year:       year,
		Month:      month,
		CardNumber: cardNumber,
	}

	records, err := s.service.FindMonthTopupStatusSuccessByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupMonthStatusSuccessResponse, len(records))
	for i, item := range records {
		protoData[i] = &pb.TopupMonthStatusSuccessResponse{
			Year:         item.Year,
			Month:        item.Month,
			TotalSuccess: int32(item.TotalSuccess),
			TotalAmount:  item.TotalAmount,
		}
	}

	return &pb.ApiResponseTopupMonthStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched monthly topup status success",
		Data:    protoData,
	}, nil
}

func (s *topupStatusHandleGrpc) FindYearlyTopupStatusSuccessByCardNumber(ctx context.Context, req *pbtopup.FindYearTopupStatusCardNumber) (*pb.ApiResponseTopupYearStatusSuccess, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidYear
	}
	if cardNumber == "" {
		return nil, topup_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.YearTopupStatusCardNumber{
		Year:       year,
		CardNumber: cardNumber,
	}

	records, err := s.service.FindYearlyTopupStatusSuccessByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupYearStatusSuccessResponse, len(records))
	for i, item := range records {
		protoData[i] = &pb.TopupYearStatusSuccessResponse{
			Year:         item.Year,
			TotalSuccess: item.TotalSuccess,
			TotalAmount:  item.TotalAmount,
		}
	}

	return &pb.ApiResponseTopupYearStatusSuccess{
		Status:  "success",
		Message: "Successfully fetched yearly topup status success",
		Data:    protoData,
	}, nil
}

func (s *topupStatusHandleGrpc) FindMonthlyTopupStatusFailedByCardNumber(ctx context.Context, req *pbtopup.FindMonthlyTopupStatusCardNumber) (*pb.ApiResponseTopupMonthStatusFailed, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidYear
	}
	if month <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidMonth
	}
	if cardNumber == "" {
		return nil, topup_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.MonthTopupStatusCardNumber{
		Year:       year,
		Month:      month,
		CardNumber: cardNumber,
	}

	records, err := s.service.FindMonthTopupStatusFailedByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupMonthStatusFailedResponse, len(records))
	for i, item := range records {
		protoData[i] = &pb.TopupMonthStatusFailedResponse{
			Year:        item.Year,
			Month:       item.Month,
			TotalFailed: int32(item.TotalFailed),
			TotalAmount: item.TotalAmount,
		}
	}

	return &pb.ApiResponseTopupMonthStatusFailed{
		Status:  "success",
		Message: "Successfully fetched monthly topup status failed",
		Data:    protoData,
	}, nil
}

func (s *topupStatusHandleGrpc) FindYearlyTopupStatusFailedByCardNumber(ctx context.Context, req *pbtopup.FindYearTopupStatusCardNumber) (*pb.ApiResponseTopupYearStatusFailed, error) {
	year := int(req.GetYear())
	cardNumber := req.GetCardNumber()

	if year <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidYear
	}
	if cardNumber == "" {
		return nil, topup_errors.ErrGrpcInvalidCardNumber
	}

	reqService := requests.YearTopupStatusCardNumber{
		Year:       year,
		CardNumber: cardNumber,
	}

	records, err := s.service.FindYearlyTopupStatusFailedByCardNumber(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupYearStatusFailedResponse, len(records))
	for i, item := range records {
		protoData[i] = &pb.TopupYearStatusFailedResponse{
			Year:        item.Year,
			TotalFailed: item.TotalFailed,
			TotalAmount: item.TotalAmount,
		}
	}

	return &pb.ApiResponseTopupYearStatusFailed{
		Status:  "success",
		Message: "Successfully fetched yearly topup status failed",
		Data:    protoData,
	}, nil
}
