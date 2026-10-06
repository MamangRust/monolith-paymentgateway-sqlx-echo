package handler

import (
	"context"
	"math"
	"time"

	pbhelpers "github.com/MamangRust/monolith-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/grpc"
	"github.com/MamangRust/monolith-payment-gateway-topup/service"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type topupQueryHandleGrpc struct {
	pb.UnimplementedTopupQueryServiceServer

	service service.TopupQueryService
}

func NewTopupQueryHandleGrpc(service service.TopupQueryService) TopupQueryHandleGrpc {
	return &topupQueryHandleGrpc{
		service: service,
	}
}

func (s *topupQueryHandleGrpc) FindAllTopup(ctx context.Context, req *pb.FindAllTopupRequest) (*pb.ApiResponsePaginationTopup, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllTopups{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	topups, totalRecords, err := s.service.FindAll(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTopups := make([]*pb.TopupResponse, len(topups))
	for i, topup := range topups {
		protoTopups[i] = &pb.TopupResponse{
			Id:          int32(topup.TopupID),
			CardNumber:  topup.CardNumber,
			TopupNo:     topup.TopupNo.String(),
			TopupAmount: int32(topup.TopupAmount),
			TopupMethod: topup.TopupMethod,
			TopupTime:   topup.TopupTime.Format(time.RFC3339),
			CreatedAt:   topup.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:   topup.UpdatedAt.Time.Format(time.RFC3339),
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))
	paginationMeta := &pbhelpers.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationTopup{
		Status:         "success",
		Message:        "Successfully fetch topups",
		Data:           protoTopups,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *topupQueryHandleGrpc) FindAllTopupByCardNumber(ctx context.Context, req *pb.FindAllTopupByCardNumberRequest) (*pb.ApiResponsePaginationTopup, error) {
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

	reqService := requests.FindAllTopupsByCardNumber{
		CardNumber: card_number,
		Page:       page,
		PageSize:   pageSize,
		Search:     search,
	}

	topups, totalRecords, err := s.service.FindAllByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTopups := make([]*pb.TopupResponse, len(topups))
	for i, topup := range topups {
		protoTopups[i] = &pb.TopupResponse{
			Id:          int32(topup.TopupID),
			CardNumber:  topup.CardNumber,
			TopupNo:     topup.TopupNo.String(),
			TopupAmount: int32(topup.TopupAmount),
			TopupMethod: topup.TopupMethod,
			TopupTime:   topup.TopupTime.Format(time.RFC3339),
			CreatedAt:   topup.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:   topup.UpdatedAt.Time.Format(time.RFC3339),
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbhelpers.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationTopup{
		Status:         "success",
		Message:        "Successfully fetch topups",
		Data:           protoTopups,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *topupQueryHandleGrpc) FindByIdTopup(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopup, error) {
	id := int(req.GetTopupId())

	if id == 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidID
	}

	topup, err := s.service.FindById(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTopup := &pb.TopupResponse{
		Id:          int32(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo.String(),
		TopupAmount: int32(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format(time.RFC3339),
		CreatedAt:   topup.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:   topup.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseTopup{
		Status:  "success",
		Message: "Successfully fetch topup",
		Data:    protoTopup,
	}, nil
}

func (s *topupQueryHandleGrpc) FindByActive(ctx context.Context, req *pb.FindAllTopupRequest) (*pb.ApiResponsePaginationTopupDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllTopups{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	res, totalRecords, err := s.service.FindByActive(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTopups := make([]*pb.TopupResponseDeleteAt, len(res))
	for i, topup := range res {
		protoTopups[i] = &pb.TopupResponseDeleteAt{
			Id:          int32(topup.TopupID),
			CardNumber:  topup.CardNumber,
			TopupNo:     topup.TopupNo.String(),
			TopupAmount: int32(topup.TopupAmount),
			TopupMethod: topup.TopupMethod,
			TopupTime:   topup.TopupTime.Format(time.RFC3339),
			CreatedAt:   topup.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:   topup.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:   wrapperspb.String(topup.DeletedAt.Time.Format(time.RFC3339)),
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbhelpers.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationTopupDeleteAt{
		Status:         "success",
		Message:        "Successfully fetch topups",
		Data:           protoTopups,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *topupQueryHandleGrpc) FindByTrashed(ctx context.Context, req *pb.FindAllTopupRequest) (*pb.ApiResponsePaginationTopupDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllTopups{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	res, totalRecords, err := s.service.FindByTrashed(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTopups := make([]*pb.TopupResponseDeleteAt, len(res))
	for i, topup := range res {
		protoTopups[i] = &pb.TopupResponseDeleteAt{
			Id:          int32(topup.TopupID),
			CardNumber:  topup.CardNumber,
			TopupNo:     topup.TopupNo.String(),
			TopupAmount: int32(topup.TopupAmount),
			TopupMethod: topup.TopupMethod,
			TopupTime:   topup.TopupTime.Format(time.RFC3339),
			CreatedAt:   topup.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:   topup.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:   wrapperspb.String(topup.DeletedAt.Time.Format(time.RFC3339)),
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbhelpers.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationTopupDeleteAt{
		Status:         "success",
		Message:        "Successfully fetch topups",
		Data:           protoTopups,
		PaginationMeta: paginationMeta,
	}, nil
}
