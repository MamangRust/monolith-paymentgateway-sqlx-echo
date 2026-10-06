package handler

import (
	"context"
	"math"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-merchant/service"
	pbutils "github.com/MamangRust/monolith-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type merchantQueryHandleGrpc struct {
	pb.UnimplementedMerchantQueryServiceServer

	merchantQuery service.MerchantQueryService
}

func NewMerchantQueryHandleGrpc(merchantQuery service.MerchantQueryService) MerchantQueryHandleGrpc {
	return &merchantQueryHandleGrpc{
		merchantQuery: merchantQuery,
	}
}

func (s *merchantQueryHandleGrpc) FindAllMerchant(ctx context.Context, req *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchant, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchants{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	merchants, totalRecords, err := s.merchantQuery.FindAll(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoMerchants := make([]*pb.MerchantResponse, len(merchants))
	for i, merchant := range merchants {
		protoMerchants[i] = &pb.MerchantResponse{
			Id:        int32(merchant.MerchantID),
			Name:      merchant.Name,
			ApiKey:    merchant.ApiKey,
			Status:    merchant.Status,
			UserId:    int32(merchant.UserID),
			CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))
	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationMerchant{
		Status:         "success",
		Message:        "Successfully fetched merchant record",
		Data:           protoMerchants,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *merchantQueryHandleGrpc) FindByIdMerchant(ctx context.Context, req *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchant, error) {
	id := int(req.GetMerchantId())
	if id <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidID
	}

	merchant, err := s.merchantQuery.FindById(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoMerchant := &pb.MerchantResponse{
		Id:        int32(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		Status:    merchant.Status,
		UserId:    int32(merchant.UserID),
		CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully fetched merchant record",
		Data:    protoMerchant,
	}, nil
}

func (s *merchantQueryHandleGrpc) FindByApiKey(ctx context.Context, req *pb.FindByApiKeyRequest) (*pb.ApiResponseMerchant, error) {
	api_key := req.GetApiKey()

	if api_key == "" {
		return nil, merchant_errors.ErrGrpcMerchantInvalidApiKey
	}

	merchant, err := s.merchantQuery.FindByApiKey(ctx, api_key)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoMerchant := &pb.MerchantResponse{
		Id:        int32(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		Status:    merchant.Status,
		UserId:    int32(merchant.UserID),
		CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully fetched merchant record",
		Data:    protoMerchant,
	}, nil
}

func (s *merchantQueryHandleGrpc) FindByMerchantUserId(ctx context.Context, req *pb.FindByMerchantUserIdRequest) (*pb.ApiResponsesMerchant, error) {
	user_id := req.GetUserId()

	if user_id <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidUserID
	}

	res, err := s.merchantQuery.FindByMerchantUserId(ctx, int(user_id))

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoMerchant := make([]*pb.MerchantResponse, 0, len(res))

	for _, merchant := range res {
		protoMerchant = append(protoMerchant, &pb.MerchantResponse{
			Id:        int32(merchant.MerchantID),
			Name:      merchant.Name,
			ApiKey:    merchant.ApiKey,
			Status:    merchant.Status,
			UserId:    int32(merchant.UserID),
			CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
		})
	}

	return &pb.ApiResponsesMerchant{
		Status:  "success",
		Message: "Successfully fetched merchant record",
		Data:    protoMerchant,
	}, nil
}

func (s *merchantQueryHandleGrpc) FindByActive(ctx context.Context, req *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchantDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchants{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	res, totalRecords, err := s.merchantQuery.FindByActive(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoMerchants := make([]*pb.MerchantResponseDeleteAt, len(res))
	for i, merchant := range res {
		var deletedAt *wrapperspb.StringValue
		if merchant.DeletedAt.Valid {
			deletedAt = wrapperspb.String(merchant.DeletedAt.Time.Format(time.RFC3339))
		}

		protoMerchants[i] = &pb.MerchantResponseDeleteAt{
			Id:        int32(merchant.MerchantID),
			Name:      merchant.Name,
			ApiKey:    merchant.ApiKey,
			Status:    merchant.Status,
			UserId:    int32(merchant.UserID),
			CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt: deletedAt,
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))
	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationMerchantDeleteAt{
		Status:         "success",
		Message:        "Successfully fetched merchant record",
		Data:           protoMerchants,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *merchantQueryHandleGrpc) FindByTrashed(ctx context.Context, req *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchantDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchants{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	res, totalRecords, err := s.merchantQuery.FindByTrashed(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoMerchants := make([]*pb.MerchantResponseDeleteAt, len(res))
	for i, merchant := range res {
		var deletedAt *wrapperspb.StringValue
		if merchant.DeletedAt.Valid {
			deletedAt = wrapperspb.String(merchant.DeletedAt.Time.Format(time.RFC3339))
		}

		protoMerchants[i] = &pb.MerchantResponseDeleteAt{
			Id:        int32(merchant.MerchantID),
			Name:      merchant.Name,
			ApiKey:    merchant.ApiKey,
			Status:    merchant.Status,
			UserId:    int32(merchant.UserID),
			CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt: deletedAt,
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))
	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationMerchantDeleteAt{
		Status:         "success",
		Message:        "Successfully fetched merchant record",
		Data:           protoMerchants,
		PaginationMeta: paginationMeta,
	}, nil
}
