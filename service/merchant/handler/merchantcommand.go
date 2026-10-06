package handler

import (
	"context"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-merchant/service"
	pbmerchant "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type merchantCommandHandleGrpc struct {
	pbmerchant.UnimplementedMerchantCommandServiceServer

	merchantCommand service.MerchantCommandService
}

func NewMerchantCommandHandleGrpc(merchantCommand service.MerchantCommandService) MerchantCommandHandleGrpc {
	return &merchantCommandHandleGrpc{
		merchantCommand: merchantCommand,
	}
}

func (s *merchantCommandHandleGrpc) CreateMerchant(ctx context.Context, req *pbmerchant.CreateMerchantRequest) (*pbmerchant.ApiResponseMerchant, error) {
	request := requests.CreateMerchantRequest{
		Name:   req.GetName(),
		UserID: int(req.GetUserId()),
	}

	if err := request.Validate(); err != nil {
		return nil, merchant_errors.ErrGrpcValidateCreateMerchant
	}

	merchant, err := s.merchantCommand.CreateMerchant(ctx, &request)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoMerchant := &pbmerchant.MerchantResponse{
		Id:        int32(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		Status:    merchant.Status,
		UserId:    int32(merchant.UserID),
		CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pbmerchant.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully created merchant",
		Data:    protoMerchant,
	}, nil
}

func (s *merchantCommandHandleGrpc) UpdateMerchant(ctx context.Context, req *pbmerchant.UpdateMerchantRequest) (*pbmerchant.ApiResponseMerchant, error) {
	id := int(req.GetMerchantId())

	if id <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidID
	}

	request := requests.UpdateMerchantRequest{
		MerchantID: &id,
		Name:       req.GetName(),
		UserID:     int(req.GetUserId()),
		Status:     req.GetStatus(),
	}

	if err := request.Validate(); err != nil {
		return nil, merchant_errors.ErrGrpcValidateUpdateMerchant
	}

	merchant, err := s.merchantCommand.UpdateMerchant(ctx, &request)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoMerchant := &pbmerchant.MerchantResponse{
		Id:        int32(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		Status:    merchant.Status,
		UserId:    int32(merchant.UserID),
		CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pbmerchant.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully updated merchant",
		Data:    protoMerchant,
	}, nil
}

func (s *merchantCommandHandleGrpc) UpdateMerchantStatus(ctx context.Context, req *pbmerchant.UpdateMerchantStatusRequest) (*pbmerchant.ApiResponseMerchant, error) {
	id := int(req.GetMerchantId())

	if id <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidID
	}

	request := requests.UpdateMerchantStatusRequest{
		MerchantID: &id,
		Status:     req.GetStatus(),
	}

	if err := request.Validate(); err != nil {
		return nil, merchant_errors.ErrGrpcValidateUpdateMerchantStatus
	}

	merchant, err := s.merchantCommand.UpdateMerchantStatus(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoMerchant := &pbmerchant.MerchantResponse{
		Id:        int32(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		Status:    merchant.Status,
		UserId:    int32(merchant.UserID),
		CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pbmerchant.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully updated merchant status",
		Data:    protoMerchant,
	}, nil
}

func (s *merchantCommandHandleGrpc) TrashedMerchant(ctx context.Context, req *pbmerchant.FindByIdMerchantRequest) (*pbmerchant.ApiResponseMerchantDeleteAt, error) {
	id := int(req.GetMerchantId())

	if id <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidID
	}

	merchant, err := s.merchantCommand.TrashedMerchant(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var deletedAt *wrapperspb.StringValue
	if merchant.DeletedAt.Valid {
		deletedAt = wrapperspb.String(merchant.DeletedAt.Time.Format(time.RFC3339))
	}

	protoMerchant := &pbmerchant.MerchantResponseDeleteAt{
		Id:        int32(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		Status:    merchant.Status,
		UserId:    int32(merchant.UserID),
		CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt: deletedAt,
	}

	return &pbmerchant.ApiResponseMerchantDeleteAt{
		Status:  "success",
		Message: "Successfully trashed merchant",
		Data:    protoMerchant,
	}, nil
}

func (s *merchantCommandHandleGrpc) RestoreMerchant(ctx context.Context, req *pbmerchant.FindByIdMerchantRequest) (*pbmerchant.ApiResponseMerchantDeleteAt, error) {
	id := int(req.GetMerchantId())

	if id <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidID
	}

	merchant, err := s.merchantCommand.RestoreMerchant(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var deletedAt *wrapperspb.StringValue
	if merchant.DeletedAt.Valid {
		deletedAt = wrapperspb.String(merchant.DeletedAt.Time.Format(time.RFC3339))
	}

	protoMerchant := &pbmerchant.MerchantResponseDeleteAt{
		Id:        int32(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		Status:    merchant.Status,
		UserId:    int32(merchant.UserID),
		CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt: deletedAt,
	}

	return &pbmerchant.ApiResponseMerchantDeleteAt{
		Status:  "success",
		Message: "Successfully restored merchant",
		Data:    protoMerchant,
	}, nil
}

func (s *merchantCommandHandleGrpc) DeleteMerchantPermanent(ctx context.Context, req *pbmerchant.FindByIdMerchantRequest) (*pbmerchant.ApiResponseMerchantDelete, error) {
	id := int(req.GetMerchantId())

	if id <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidID
	}

	_, err := s.merchantCommand.DeleteMerchantPermanent(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pbmerchant.ApiResponseMerchantDelete{
		Status:  "success",
		Message: "Successfully deleted merchant",
	}, nil
}

func (s *merchantCommandHandleGrpc) RestoreAllMerchant(ctx context.Context, _ *emptypb.Empty) (*pbmerchant.ApiResponseMerchantAll, error) {
	_, err := s.merchantCommand.RestoreAllMerchant(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pbmerchant.ApiResponseMerchantAll{
		Status:  "success",
		Message: "Successfully restore all merchant",
	}, nil
}

func (s *merchantCommandHandleGrpc) DeleteAllMerchantPermanent(ctx context.Context, _ *emptypb.Empty) (*pbmerchant.ApiResponseMerchantAll, error) {
	_, err := s.merchantCommand.DeleteAllMerchantPermanent(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pbmerchant.ApiResponseMerchantAll{
		Status:  "success",
		Message: "Successfully delete all merchant",
	}, nil
}
