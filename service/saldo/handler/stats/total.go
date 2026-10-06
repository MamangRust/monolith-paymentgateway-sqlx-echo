package saldostatshandler

import (
	"context"

	pbsaldo "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo/stats"
	saldostatsservice "github.com/MamangRust/monolith-payment-gateway-saldo/service/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	saldo_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/saldo_errors/grpc"
)

type saldoStatsTotalBalanceHandleGrpc struct {
	pb.UnimplementedSaldoStatsTotalBalanceServer

	service saldostatsservice.SaldoStatsTotalBalanceService
}

func NewSaldoStatsTotalBalanceHandleGrpc(service saldostatsservice.SaldoStatsTotalBalanceService) SaldoStatsTotalBalanceHandleGrpc {
	return &saldoStatsTotalBalanceHandleGrpc{
		service: service,
	}
}

func (s *saldoStatsTotalBalanceHandleGrpc) FindMonthlyTotalSaldoBalance(ctx context.Context, req *pbsaldo.FindMonthlySaldoTotalBalance) (*pb.ApiResponseMonthTotalSaldo, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, saldo_errors.ErrGrpcSaldoInvalidYear
	}
	if month <= 0 {
		return nil, saldo_errors.ErrGrpcSaldoInvalidMonth
	}

	reqService := requests.MonthTotalSaldoBalance{
		Year:  year,
		Month: month,
	}

	res, err := s.service.FindMonthlyTotalSaldoBalance(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.SaldoMonthTotalBalanceResponse, len(res))
	for i, item := range res {
		protoData[i] = &pb.SaldoMonthTotalBalanceResponse{
			Month:        item.Month,
			Year:         item.Year,
			TotalBalance: int32(item.TotalBalance),
		}
	}

	return &pb.ApiResponseMonthTotalSaldo{
		Status:  "success",
		Message: "Successfully fetched monthly total saldo balance",
		Data:    protoData,
	}, nil
}

func (s *saldoStatsTotalBalanceHandleGrpc) FindYearTotalSaldoBalance(ctx context.Context, req *pbsaldo.FindYearlySaldo) (*pb.ApiResponseYearTotalSaldo, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, saldo_errors.ErrGrpcSaldoInvalidYear
	}

	res, err := s.service.FindYearTotalSaldoBalance(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.SaldoYearTotalBalanceResponse, len(res))
	for i, item := range res {
		protoData[i] = &pb.SaldoYearTotalBalanceResponse{
			Year:         item.Year,
			TotalBalance: int32(item.TotalBalance),
		}
	}

	return &pb.ApiResponseYearTotalSaldo{
		Status:  "success",
		Message: "Successfully fetched yearly total saldo balance",
		Data:    protoData,
	}, nil
}
