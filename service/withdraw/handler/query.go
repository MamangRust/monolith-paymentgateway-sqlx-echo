package handler

import (
	"context"
	"math"
	"time"

	pbhelper "github.com/MamangRust/monolith-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	withdraw_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/withdraw_errors/grpc"
	"github.com/MamangRust/monolith-payment-gateway-withdraw/service"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type withdrawQueryHandleGrpc struct {
	pb.UnimplementedWithdrawQueryServiceServer

	withdrawQuery service.WithdrawQueryService
}

func NewWithdrawQueryHandleGrpc(
	withdrawQuery service.WithdrawQueryService,
) WithdrawQueryHandlerGrpc {
	return &withdrawQueryHandleGrpc{
		withdrawQuery: withdrawQuery,
	}
}

func (w *withdrawQueryHandleGrpc) FindAllWithdraw(ctx context.Context, req *pb.FindAllWithdrawRequest) (*pb.ApiResponsePaginationWithdraw, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllWithdraws{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	withdraws, totalRecords, err := w.withdrawQuery.FindAll(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbhelper.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	withdrawResponses := make([]*pb.WithdrawResponse, len(withdraws))
	for i, withdraw := range withdraws {
		withdrawResponses[i] = &pb.WithdrawResponse{
			WithdrawId:     int32(withdraw.WithdrawID),
			WithdrawNo:     withdraw.WithdrawNo.String(),
			CardNumber:     withdraw.CardNumber,
			WithdrawAmount: int32(withdraw.WithdrawAmount),
			WithdrawTime:   withdraw.WithdrawTime.Format(time.RFC3339),
			CreatedAt:      withdraw.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      withdraw.UpdatedAt.Time.Format(time.RFC3339),
		}
	}

	return &pb.ApiResponsePaginationWithdraw{
		Status:         "success",
		Message:        "withdraw",
		Data:           withdrawResponses,
		PaginationMeta: paginationMeta,
	}, nil
}

func (w *withdrawQueryHandleGrpc) FindAllWithdrawByCardNumber(ctx context.Context, req *pb.FindAllWithdrawByCardNumberRequest) (*pb.ApiResponsePaginationWithdraw, error) {
	card_number := req.GetCardNumber()
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllWithdrawCardNumber{
		CardNumber: card_number,
		Page:       page,
		PageSize:   pageSize,
		Search:     search,
	}

	withdraws, totalRecords, err := w.withdrawQuery.FindAllByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbhelper.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	withdrawResponses := make([]*pb.WithdrawResponse, len(withdraws))
	for i, withdraw := range withdraws {
		withdrawResponses[i] = &pb.WithdrawResponse{
			WithdrawId:     int32(withdraw.WithdrawID),
			WithdrawNo:     withdraw.WithdrawNo.String(),
			CardNumber:     withdraw.CardNumber,
			WithdrawAmount: int32(withdraw.WithdrawAmount),
			WithdrawTime:   withdraw.WithdrawTime.Format(time.RFC3339),
			CreatedAt:      withdraw.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      withdraw.UpdatedAt.Time.Format(time.RFC3339),
		}
	}

	return &pb.ApiResponsePaginationWithdraw{
		Status:         "success",
		Message:        "Withdraws fetched successfully",
		Data:           withdrawResponses,
		PaginationMeta: paginationMeta,
	}, nil
}

func (w *withdrawQueryHandleGrpc) FindByIdWithdraw(ctx context.Context, req *pb.FindByIdWithdrawRequest) (*pb.ApiResponseWithdraw, error) {
	id := int(req.GetWithdrawId())

	if id == 0 {
		return nil, withdraw_errors.ErrGrpcWithdrawInvalidID
	}

	withdraw, err := w.withdrawQuery.FindById(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseWithdraw{
		Status:  "success",
		Message: "Successfully fetched withdraw",
		Data: &pb.WithdrawResponse{
			WithdrawId:     int32(withdraw.WithdrawID),
			WithdrawNo:     withdraw.WithdrawNo.String(),
			CardNumber:     withdraw.CardNumber,
			WithdrawAmount: int32(withdraw.WithdrawAmount),
			WithdrawTime:   withdraw.WithdrawTime.Format(time.RFC3339),
			CreatedAt:      withdraw.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      withdraw.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (w *withdrawQueryHandleGrpc) FindByActive(ctx context.Context, req *pb.FindAllWithdrawRequest) (*pb.ApiResponsePaginationWithdrawDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllWithdraws{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	withdraws, totalRecords, err := w.withdrawQuery.FindByActive(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbhelper.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	withdrawResponses := make([]*pb.WithdrawResponseDeleteAt, len(withdraws))
	for i, withdraw := range withdraws {
		withdrawResponses[i] = &pb.WithdrawResponseDeleteAt{
			WithdrawId:     int32(withdraw.WithdrawID),
			WithdrawNo:     withdraw.WithdrawNo.String(),
			CardNumber:     withdraw.CardNumber,
			WithdrawAmount: int32(withdraw.WithdrawAmount),
			WithdrawTime:   withdraw.WithdrawTime.Format(time.RFC3339),
			CreatedAt:      withdraw.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      withdraw.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:      &wrapperspb.StringValue{Value: withdraw.DeletedAt.Time.Format(time.RFC3339)},
		}
	}

	return &pb.ApiResponsePaginationWithdrawDeleteAt{
		Status:         "success",
		Message:        "Successfully fetched withdraws",
		Data:           withdrawResponses,
		PaginationMeta: paginationMeta,
	}, nil
}

func (w *withdrawQueryHandleGrpc) FindByTrashed(ctx context.Context, req *pb.FindAllWithdrawRequest) (*pb.ApiResponsePaginationWithdrawDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllWithdraws{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	withdraws, totalRecords, err := w.withdrawQuery.FindByTrashed(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbhelper.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	withdrawResponses := make([]*pb.WithdrawResponseDeleteAt, len(withdraws))
	for i, withdraw := range withdraws {
		withdrawResponses[i] = &pb.WithdrawResponseDeleteAt{
			WithdrawId:     int32(withdraw.WithdrawID),
			WithdrawNo:     withdraw.WithdrawNo.String(),
			CardNumber:     withdraw.CardNumber,
			WithdrawAmount: int32(withdraw.WithdrawAmount),
			WithdrawTime:   withdraw.WithdrawTime.Format(time.RFC3339),
			CreatedAt:      withdraw.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      withdraw.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:      &wrapperspb.StringValue{Value: withdraw.DeletedAt.Time.Format(time.RFC3339)},
		}
	}

	return &pb.ApiResponsePaginationWithdrawDeleteAt{
		Status:         "success",
		Message:        "Successfully fetched withdraws",
		Data:           withdrawResponses,
		PaginationMeta: paginationMeta,
	}, nil
}
