package handler

import (
	"context"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-card/service"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type cardCommandService struct {
	pb.UnimplementedCardCommandServiceServer

	cardCommand service.CardCommandService
}

func NewCardCommandHandleGrpc(cardCommand service.CardCommandService) CardCommandService {
	return &cardCommandService{
		cardCommand: cardCommand,
	}
}

func (s *cardCommandService) CreateCard(ctx context.Context, req *pb.CreateCardRequest) (*pb.ApiResponseCard, error) {
	request := requests.CreateCardRequest{
		UserID:       int(req.UserId),
		CardType:     req.CardType,
		ExpireDate:   req.ExpireDate.AsTime(),
		CVV:          req.Cvv,
		CardProvider: req.CardProvider,
	}

	if err := request.Validate(); err != nil {
		return nil, card_errors.ErrGrpcValidateCreateCardRequest
	}

	res, err := s.cardCommand.CreateCard(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCard := &pb.CardResponse{
		Id:         int32(res.CardID),
		UserId:     int32(res.UserID),
		CardNumber: res.CardNumber,
		CardType:   res.CardType,
		Cvv:        res.Cvv,
		ExpireDate: res.ExpireDate.Time.Format(time.RFC3339),
		CreatedAt:  res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:  res.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully created card",
		Data:    protoCard,
	}, nil
}

func (s *cardCommandService) UpdateCard(ctx context.Context, req *pb.UpdateCardRequest) (*pb.ApiResponseCard, error) {
	request := requests.UpdateCardRequest{
		CardID:       int(req.CardId),
		UserID:       int(req.UserId),
		CardType:     req.CardType,
		ExpireDate:   req.ExpireDate.AsTime(),
		CVV:          req.Cvv,
		CardProvider: req.CardProvider,
	}

	if err := request.Validate(); err != nil {
		return nil, card_errors.ErrGrpcValidateUpdateCardRequest
	}

	res, err := s.cardCommand.UpdateCard(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCard := &pb.CardResponse{
		Id:         int32(res.CardID),
		UserId:     int32(res.UserID),
		CardNumber: res.CardNumber,
		CardType:   res.CardType,
		Cvv:        res.Cvv,
		ExpireDate: res.ExpireDate.Time.Format(time.RFC3339),
		CreatedAt:  res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:  res.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully updated card",
		Data:    protoCard,
	}, nil
}

func (s *cardCommandService) TrashedCard(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCardDeleteAt, error) {
	id := int(req.GetCardId())
	if id <= 0 {
		return nil, card_errors.ErrGrpcInvalidCardID
	}

	res, err := s.cardCommand.TrashedCard(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCard := &pb.CardResponseDeleteAt{
		Id:         int32(res.CardID),
		UserId:     int32(res.UserID),
		CardNumber: res.CardNumber,
		CardType:   res.CardType,
		Cvv:        res.Cvv,
		ExpireDate: res.ExpireDate.Time.Format(time.RFC3339),
		CreatedAt:  res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:  res.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt:  wrapperspb.String(res.DeletedAt.Time.Format(time.RFC3339)),
	}

	return &pb.ApiResponseCardDeleteAt{
		Status:  "success",
		Message: "Successfully trashed card",
		Data:    protoCard,
	}, nil
}

func (s *cardCommandService) RestoreCard(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCardDeleteAt, error) {
	id := int(req.GetCardId())
	if id <= 0 {
		return nil, card_errors.ErrGrpcInvalidCardID
	}

	res, err := s.cardCommand.RestoreCard(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCard := &pb.CardResponseDeleteAt{
		Id:         int32(res.CardID),
		UserId:     int32(res.UserID),
		CardNumber: res.CardNumber,
		CardType:   res.CardType,
		Cvv:        res.Cvv,
		ExpireDate: res.ExpireDate.Time.Format(time.RFC3339),
		CreatedAt:  res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:  res.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt:  wrapperspb.String(res.DeletedAt.Time.Format(time.RFC3339)),
	}

	return &pb.ApiResponseCardDeleteAt{
		Status:  "success",
		Message: "Successfully restored card",
		Data:    protoCard,
	}, nil
}

func (s *cardCommandService) DeleteCardPermanent(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCardDelete, error) {
	id := int(req.GetCardId())
	if id <= 0 {
		return nil, card_errors.ErrGrpcInvalidCardID
	}

	_, err := s.cardCommand.DeleteCardPermanent(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCardDelete{
		Status:  "success",
		Message: "Successfully deleted card",
	}, nil
}

func (s *cardCommandService) RestoreAllCard(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCardAll, error) {
	_, err := s.cardCommand.RestoreAllCard(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCardAll{
		Status:  "success",
		Message: "Successfully restore card",
	}, nil
}

func (s *cardCommandService) DeleteAllCardPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCardAll, error) {
	_, err := s.cardCommand.DeleteAllCardPermanent(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCardAll{
		Status:  "success",
		Message: "Successfully delete card permanent",
	}, nil
}

func (s *cardCommandService) ToggleCardStatus(ctx context.Context, req *pb.ToggleCardStatusRequest) (*pb.ApiResponseCard, error) {
	id := int(req.GetCardId())
	if id <= 0 {
		return nil, card_errors.ErrGrpcInvalidCardID
	}

	request := requests.ToggleCardStatusRequest{
		CardID: id,
	}

	if err := request.Validate(); err != nil {
		return nil, card_errors.ErrGrpcValidateToggleCardStatus
	}

	res, err := s.cardCommand.ToggleCardStatus(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCard := &pb.CardResponse{
		Id:                 int32(res.CardID),
		UserId:             int32(res.UserID),
		CardNumber:         res.CardNumber,
		CardType:           res.CardType,
		Cvv:                res.Cvv,
		ExpireDate:         res.ExpireDate.Time.Format(time.RFC3339),
		CardProvider:       res.CardProvider,
		Status:             res.Status,
		CreditLimit:        int32(res.CreditLimit),
		OutstandingBalance: int32(res.OutstandingBalance),
		RewardPoints:       int32(res.RewardPoints),
		CreatedAt:          res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:          res.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully toggled card status",
		Data:    protoCard,
	}, nil
}

func (s *cardCommandService) UpdateCreditLimit(ctx context.Context, req *pb.UpdateCreditLimitRequest) (*pb.ApiResponseCard, error) {
	id := int(req.GetCardId())
	if id <= 0 {
		return nil, card_errors.ErrGrpcInvalidCardID
	}

	request := requests.UpdateCreditLimitRequest{
		CardID:      id,
		CreditLimit: int(req.GetCreditLimit()),
	}

	if err := request.Validate(); err != nil {
		return nil, card_errors.ErrGrpcValidateUpdateCreditLimit
	}

	res, err := s.cardCommand.UpdateCreditLimit(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCard := &pb.CardResponse{
		Id:                 int32(res.CardID),
		UserId:             int32(res.UserID),
		CardNumber:         res.CardNumber,
		CardType:           res.CardType,
		Cvv:                res.Cvv,
		ExpireDate:         res.ExpireDate.Time.Format(time.RFC3339),
		CardProvider:       res.CardProvider,
		Status:             res.Status,
		CreditLimit:        int32(res.CreditLimit),
		OutstandingBalance: int32(res.OutstandingBalance),
		RewardPoints:       int32(res.RewardPoints),
		CreatedAt:          res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:          res.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully updated credit limit",
		Data:    protoCard,
	}, nil
}

func (s *cardCommandService) RedeemPoints(ctx context.Context, req *pb.RedeemPointsRequest) (*pb.ApiResponseCard, error) {
	id := int(req.GetCardId())
	if id <= 0 {
		return nil, card_errors.ErrGrpcInvalidCardID
	}

	request := requests.RedeemPointsRequest{
		CardID: id,
		Points: int(req.GetPoints()),
	}

	if err := request.Validate(); err != nil {
		return nil, card_errors.ErrGrpcValidateRedeemPoints
	}

	res, err := s.cardCommand.RedeemPoints(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCard := &pb.CardResponse{
		Id:                 int32(res.CardID),
		UserId:             int32(res.UserID),
		CardNumber:         res.CardNumber,
		CardType:           res.CardType,
		Cvv:                res.Cvv,
		ExpireDate:         res.ExpireDate.Time.Format(time.RFC3339),
		CardProvider:       res.CardProvider,
		Status:             res.Status,
		CreditLimit:        int32(res.CreditLimit),
		OutstandingBalance: int32(res.OutstandingBalance),
		RewardPoints:       int32(res.RewardPoints),
		CreatedAt:          res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:          res.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully redeemed reward points",
		Data:    protoCard,
	}, nil
}

func (s *cardCommandService) ProcessBillingCycles(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.cardCommand.ProcessBillingCycles(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &emptypb.Empty{}, nil
}
