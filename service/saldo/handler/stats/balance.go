package saldostatshandler

import (
	"context"

	pbsaldo "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo/stats"
	saldostatsservice "github.com/MamangRust/monolith-payment-gateway-saldo/service/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	saldo_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/saldo_errors/grpc"
)

type saldoStatsBalanceHandleGrpc struct {
	pb.UnimplementedSaldoStatsBalanceServiceServer

	service saldostatsservice.SaldoStatsBalanceService
}

func NewSaldoStatsBalanceHandleGrpc(service saldostatsservice.SaldoStatsBalanceService) SaldoStatsBalanceHandleGrpc {
	return &saldoStatsBalanceHandleGrpc{
		service: service,
	}
}

func (s *saldoStatsBalanceHandleGrpc) FindMonthlySaldoBalances(ctx context.Context, req *pbsaldo.FindYearlySaldo) (*pb.ApiResponseMonthSaldoBalances, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, saldo_errors.ErrGrpcSaldoInvalidYear
	}

	res, err := s.service.FindMonthlySaldoBalances(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.SaldoMonthBalanceResponse, len(res))
	for i, item := range res {
		protoData[i] = &pb.SaldoMonthBalanceResponse{
			Month:        item.Month,
			TotalBalance: int32(item.TotalBalance),
		}
	}

	return &pb.ApiResponseMonthSaldoBalances{
		Status:  "success",
		Message: "Successfully fetched monthly saldo balances",
		Data:    protoData,
	}, nil
}

func (s *saldoStatsBalanceHandleGrpc) FindYearlySaldoBalances(ctx context.Context, req *pbsaldo.FindYearlySaldo) (*pb.ApiResponseYearSaldoBalances, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, saldo_errors.ErrGrpcSaldoInvalidYear
	}

	res, err := s.service.FindYearlySaldoBalances(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.SaldoYearBalanceResponse, len(res))
	for i, item := range res {
		protoData[i] = &pb.SaldoYearBalanceResponse{
			Year:         item.YYear,
			TotalBalance: int32(item.TotalBalance),
		}
	}

	return &pb.ApiResponseYearSaldoBalances{
		Status:  "success",
		Message: "Successfully fetched yearly saldo balances",
		Data:    protoData,
	}, nil
}
