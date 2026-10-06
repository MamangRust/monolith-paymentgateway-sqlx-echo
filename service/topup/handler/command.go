package handler

import (
	"context"
	"time"

	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/grpc"
	"github.com/MamangRust/monolith-payment-gateway-topup/service"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type topupCommandHandleGrpc struct {
	pb.UnimplementedTopupCommandServiceServer

	service service.TopupCommandService
}

func NewTopupCommandHandleGrpc(service service.TopupCommandService) TopupCommandHandleGrpc {
	return &topupCommandHandleGrpc{
		service: service,
	}
}

func (s *topupCommandHandleGrpc) CreateTopup(ctx context.Context, req *pb.CreateTopupRequest) (*pb.ApiResponseTopup, error) {
	request := requests.CreateTopupRequest{
		CardNumber:  req.GetCardNumber(),
		TopupAmount: int(req.GetTopupAmount()),
		TopupMethod: req.GetTopupMethod(),
	}

	if err := request.Validate(); err != nil {
		return nil, topup_errors.ErrGrpcValidateCreateTopup
	}

	res, err := s.service.CreateTopup(ctx, &request)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTopup := &pb.TopupResponse{
		Id:          int32(res.TopupID),
		CardNumber:  res.CardNumber,
		TopupNo:     res.TopupNo.String(),
		TopupAmount: int32(res.TopupAmount),
		TopupMethod: res.TopupMethod,
		TopupTime:   res.TopupTime.Format(time.RFC3339),
		CreatedAt:   res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:   res.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseTopup{
		Status:  "success",
		Message: "Successfully created topup",
		Data:    protoTopup,
	}, nil
}

func (s *topupCommandHandleGrpc) UpdateTopup(ctx context.Context, req *pb.UpdateTopupRequest) (*pb.ApiResponseTopup, error) {
	id := int(req.GetTopupId())

	if id <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidID
	}

	request := requests.UpdateTopupRequest{
		TopupID:     &id,
		CardNumber:  req.GetCardNumber(),
		TopupAmount: int(req.GetTopupAmount()),
		TopupMethod: req.GetTopupMethod(),
	}

	if err := request.Validate(); err != nil {
		return nil, topup_errors.ErrGrpcValidateUpdateTopup
	}

	res, err := s.service.UpdateTopup(ctx, &request)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTopup := &pb.TopupResponse{
		Id:          int32(res.TopupID),
		CardNumber:  res.CardNumber,
		TopupNo:     res.TopupNo.String(),
		TopupAmount: int32(res.TopupAmount),
		TopupMethod: res.TopupMethod,
		TopupTime:   res.TopupTime.Format(time.RFC3339),
		CreatedAt:   res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:   res.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseTopup{
		Status:  "success",
		Message: "Successfully updated topup",
		Data:    protoTopup,
	}, nil
}

func (s *topupCommandHandleGrpc) TrashedTopup(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopupDeleteAt, error) {
	id := int(req.GetTopupId())

	if id <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidID
	}

	res, err := s.service.TrashedTopup(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTopup := &pb.TopupResponseDeleteAt{
		Id:          int32(res.TopupID),
		CardNumber:  res.CardNumber,
		TopupNo:     res.TopupNo.String(),
		TopupAmount: int32(res.TopupAmount),
		TopupMethod: res.TopupMethod,
		TopupTime:   res.TopupTime.Format(time.RFC3339),
		CreatedAt:   res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:   res.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt:   wrapperspb.String(res.DeletedAt.Time.Format(time.RFC3339)),
	}

	return &pb.ApiResponseTopupDeleteAt{
		Status:  "success",
		Message: "Successfully trashed topup",
		Data:    protoTopup,
	}, nil
}

func (s *topupCommandHandleGrpc) RestoreTopup(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopupDeleteAt, error) {
	id := int(req.GetTopupId())

	if id <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidID
	}

	res, err := s.service.RestoreTopup(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTopup := &pb.TopupResponseDeleteAt{
		Id:          int32(res.TopupID),
		CardNumber:  res.CardNumber,
		TopupNo:     res.TopupNo.String(),
		TopupAmount: int32(res.TopupAmount),
		TopupMethod: res.TopupMethod,
		TopupTime:   res.TopupTime.Format(time.RFC3339),
		CreatedAt:   res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:   res.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt:   wrapperspb.String(res.DeletedAt.Time.Format(time.RFC3339)),
	}

	return &pb.ApiResponseTopupDeleteAt{
		Status:  "success",
		Message: "Successfully restored topup",
		Data:    protoTopup,
	}, nil
}

func (s *topupCommandHandleGrpc) DeleteTopupPermanent(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopupDelete, error) {
	id := int(req.GetTopupId())

	if id <= 0 {
		return nil, topup_errors.ErrGrpcTopupInvalidID
	}

	_, err := s.service.DeleteTopupPermanent(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTopupDelete{
		Status:  "success",
		Message: "Successfully deleted topup permanently",
	}, nil
}

func (s *topupCommandHandleGrpc) RestoreAllTopup(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTopupAll, error) {
	_, err := s.service.RestoreAllTopup(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTopupAll{
		Status:  "success",
		Message: "Successfully restore all topup",
	}, nil
}

func (s *topupCommandHandleGrpc) DeleteAllTopupPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTopupAll, error) {
	_, err := s.service.DeleteAllTopupPermanent(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTopupAll{
		Status:  "success",
		Message: "Successfully delete topup permanent",
	}, nil
}
