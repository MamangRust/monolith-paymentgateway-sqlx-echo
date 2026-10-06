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

type topupMethodHandleGrpc struct {
	pb.UnimplementedTopupStatsMethodServiceServer

	service service.Service
}

func NewTopupStatsMethodHandleGrpc(
	service service.Service,
) TopupStatsMethodHandleGrpc {
	return &topupMethodHandleGrpc{
		service: service,
	}
}

func (s *topupMethodHandleGrpc) FindMonthlyTopupMethods(ctx context.Context, req *pbtopup.FindYearTopupStatus) (*pb.ApiResponseTopupMonthMethod, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidYear
	}

	methods, err := s.service.FindMonthlyTopupMethods(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupMonthMethodResponse, len(methods))
	for i, item := range methods {
		protoData[i] = &pb.TopupMonthMethodResponse{
			Month:       item.Month,
			TopupMethod: item.TopupMethod,
			TotalTopups: item.TotalTopups,
			TotalAmount: item.TotalAmount,
		}
	}

	return &pb.ApiResponseTopupMonthMethod{
		Status:  "success",
		Message: "Successfully fetched monthly topup methods",
		Data:    protoData,
	}, nil
}

func (s *topupMethodHandleGrpc) FindYearlyTopupMethods(ctx context.Context, req *pbtopup.FindYearTopupStatus) (*pb.ApiResponseTopupYearMethod, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidYear
	}

	methods, err := s.service.FindYearlyTopupMethods(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupYearlyMethodResponse, len(methods))
	for i, item := range methods {
		protoData[i] = &pb.TopupYearlyMethodResponse{
			Year:        item.Year.Int.String(),
			TopupMethod: item.TopupMethod,
			TotalTopups: int32(item.TotalTopups),
			TotalAmount: int32(item.TotalAmount),
		}
	}

	return &pb.ApiResponseTopupYearMethod{
		Status:  "success",
		Message: "Successfully fetched yearly topup methods",
		Data:    protoData,
	}, nil
}

func (s *topupMethodHandleGrpc) FindMonthlyTopupMethodsByCardNumber(ctx context.Context, req *pbtopup.FindYearTopupCardNumber) (*pb.ApiResponseTopupMonthMethod, error) {
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

	methods, err := s.service.FindMonthlyTopupMethodsByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupMonthMethodResponse, len(methods))
	for i, item := range methods {
		protoData[i] = &pb.TopupMonthMethodResponse{
			Month:       item.Month,
			TopupMethod: item.TopupMethod,
			TotalTopups: item.TotalTopups,
			TotalAmount: item.TotalAmount,
		}
	}

	return &pb.ApiResponseTopupMonthMethod{
		Status:  "success",
		Message: "Successfully fetched monthly topup methods by card number",
		Data:    protoData,
	}, nil
}

func (s *topupMethodHandleGrpc) FindYearlyTopupMethodsByCardNumber(ctx context.Context, req *pbtopup.FindYearTopupCardNumber) (*pb.ApiResponseTopupYearMethod, error) {
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

	methods, err := s.service.FindYearlyTopupMethodsByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.TopupYearlyMethodResponse, len(methods))
	for i, item := range methods {
		protoData[i] = &pb.TopupYearlyMethodResponse{
			Year:        item.Year.Int.String(),
			TopupMethod: item.TopupMethod,
			TotalTopups: int32(item.TotalTopups),
			TotalAmount: int32(item.TotalAmount),
		}
	}

	return &pb.ApiResponseTopupYearMethod{
		Status:  "success",
		Message: "Successfully fetched yearly topup methods by card number",
		Data:    protoData,
	}, nil
}
