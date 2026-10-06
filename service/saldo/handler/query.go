package handler

import (
	"context"
	"math"
	"time"

	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pbhelpers "github.com/MamangRust/monolith-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	"github.com/MamangRust/monolith-payment-gateway-saldo/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	saldo_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/saldo_errors/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type saldoQueryHandleGrpc struct {
	pb.UnimplementedSaldoQueryServiceServer

	service service.SaldoQueryService
}

func NewSaldoQueryHandleGrpc(query service.SaldoQueryService) SaldoQueryHandleGrpc {
	return &saldoQueryHandleGrpc{
		service: query,
	}
}

func (s *saldoQueryHandleGrpc) FindAllSaldo(ctx context.Context, req *pb.FindAllSaldoRequest) (*pb.ApiResponsePaginationSaldo, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllSaldos{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	res, totalRecords, err := s.service.FindAll(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoSaldos := make([]*pb.SaldoResponse, len(res))
	for i, saldo := range res {

		protoSaldos[i] = &pb.SaldoResponse{
			SaldoId:        int32(saldo.SaldoID),
			CardNumber:     saldo.CardNumber,
			TotalBalance:   saldo.TotalBalance,
			WithdrawTime:   saldo.WithdrawTime.Time.Format(time.RFC3339),
			WithdrawAmount: Int32Value(saldo.WithdrawAmount),
			CreatedAt:      saldo.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      saldo.UpdatedAt.Time.Format(time.RFC3339),
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))
	paginationMeta := &pbhelpers.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationSaldo{
		Status:         "success",
		Message:        "Successfully fetched saldo record",
		Data:           protoSaldos,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *saldoQueryHandleGrpc) FindByIdSaldo(ctx context.Context, req *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldo, error) {
	id := int(req.GetSaldoId())
	if id <= 0 {
		return nil, saldo_errors.ErrGrpcSaldoInvalidID
	}

	saldo, err := s.service.FindById(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoSaldo := &pb.SaldoResponse{
		SaldoId:        int32(saldo.SaldoID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   saldo.TotalBalance,
		WithdrawTime:   saldo.WithdrawTime.Time.Format(time.RFC3339),
		WithdrawAmount: Int32Value(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:      saldo.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseSaldo{
		Status:  "success",
		Message: "Successfully fetched saldo record",
		Data:    protoSaldo,
	}, nil
}

func (s *saldoQueryHandleGrpc) FindByCardNumber(ctx context.Context, req *pbcard.FindByCardNumberRequest) (*pb.ApiResponseSaldo, error) {
	cardNumber := req.GetCardNumber()
	if cardNumber == "" {
		return nil, saldo_errors.ErrGrpcSaldoInvalidCardNumber
	}

	saldo, err := s.service.FindByCardNumber(ctx, cardNumber)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoSaldo := &pb.SaldoResponse{
		SaldoId:        int32(saldo.SaldoID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   saldo.TotalBalance,
		WithdrawTime:   saldo.WithdrawTime.Time.Format(time.RFC3339),
		WithdrawAmount: Int32Value(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:      saldo.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseSaldo{
		Status:  "success",
		Message: "Successfully fetched saldo record",
		Data:    protoSaldo,
	}, nil
}

func (s *saldoQueryHandleGrpc) FindByActive(ctx context.Context, req *pb.FindAllSaldoRequest) (*pb.ApiResponsePaginationSaldoDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllSaldos{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	res, totalRecords, err := s.service.FindByActive(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoSaldos := make([]*pb.SaldoResponseDeleteAt, len(res))
	for i, saldo := range res {
		protoSaldos[i] = &pb.SaldoResponseDeleteAt{
			SaldoId:        int32(saldo.SaldoID),
			CardNumber:     saldo.CardNumber,
			TotalBalance:   saldo.TotalBalance,
			WithdrawTime:   saldo.WithdrawTime.Time.Format(time.RFC3339),
			WithdrawAmount: Int32Value(saldo.WithdrawAmount),
			CreatedAt:      saldo.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      saldo.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:      wrapperspb.String(saldo.DeletedAt.Time.Format(time.RFC3339)),
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))
	paginationMeta := &pbhelpers.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationSaldoDeleteAt{
		Status:         "success",
		Message:        "Successfully fetched saldo record",
		Data:           protoSaldos,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *saldoQueryHandleGrpc) FindByTrashed(ctx context.Context, req *pb.FindAllSaldoRequest) (*pb.ApiResponsePaginationSaldoDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllSaldos{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	res, totalRecords, err := s.service.FindByTrashed(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoSaldos := make([]*pb.SaldoResponseDeleteAt, len(res))
	for i, saldo := range res {
		protoSaldos[i] = &pb.SaldoResponseDeleteAt{
			SaldoId:        int32(saldo.SaldoID),
			CardNumber:     saldo.CardNumber,
			TotalBalance:   saldo.TotalBalance,
			WithdrawTime:   saldo.WithdrawTime.Time.Format(time.RFC3339),
			WithdrawAmount: Int32Value(saldo.WithdrawAmount),
			CreatedAt:      saldo.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      saldo.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:      wrapperspb.String(saldo.DeletedAt.Time.Format(time.RFC3339)),
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))
	paginationMeta := &pbhelpers.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationSaldoDeleteAt{
		Status:         "success",
		Message:        "Successfully fetched saldo record",
		Data:           protoSaldos,
		PaginationMeta: paginationMeta,
	}, nil
}

func Int32Value(v *int32) int32 {
	if v == nil {
		return 0
	}

	return *v
}
