package handler

import (
	"context"
	"math"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-card/service"

	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pbutils "github.com/MamangRust/monolith-payment-gateway-pb/common"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type cardQueryHandleGrpc struct {
	pb.UnimplementedCardQueryServiceServer

	cardQuery service.CardQueryService
}

func NewCardQueryHandleGrpc(cardQuery service.CardQueryService) CardQueryService {
	return &cardQueryHandleGrpc{
		cardQuery: cardQuery,
	}
}

func (s *cardQueryHandleGrpc) FindAllCard(ctx context.Context, req *pb.FindAllCardRequest) (*pb.ApiResponsePaginationCard, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllCards{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	cards, totalRecords, err := s.cardQuery.FindAll(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCards := make([]*pb.CardResponse, len(cards))
	for i, card := range cards {
		protoCards[i] = &pb.CardResponse{
			Id:         int32(card.CardID),
			UserId:     int32(card.UserID),
			CardNumber: card.CardNumber,
			CardType:   card.CardType,
			Cvv:        card.Cvv,
			ExpireDate: card.ExpireDate.Time.Format(time.RFC3339),
			CreatedAt:  card.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:  card.UpdatedAt.Time.Format(time.RFC3339),
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationCard{
		Status:         "success",
		Message:        "Successfully fetched card records",
		Data:           protoCards,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *cardQueryHandleGrpc) FindByIdCard(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCard, error) {
	id := int(req.GetCardId())

	if id == 0 {

		return nil, card_errors.ErrGrpcInvalidCardID
	}

	card, err := s.cardQuery.FindById(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	res := &pb.ApiResponseCard{
		Message: "successfully",
		Status:  "success",
		Data: &pb.CardResponse{
			Id:           int32(card.CardID),
			UserId:       int32(card.UserID),
			CardNumber:   card.CardNumber,
			CardType:     card.CardType,
			CardProvider: card.CardProvider,
			Cvv:          card.Cvv,
			ExpireDate:   card.ExpireDate.Time.Format(time.RFC3339),
			CreatedAt:    card.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:    card.UpdatedAt.Time.Format(time.RFC3339),
		},
	}

	return res, nil
}

func (s *cardQueryHandleGrpc) FindByCardNumber(ctx context.Context, req *pb.FindByCardNumberRequest) (*pb.ApiResponseCard, error) {
	card_number := req.GetCardNumber()
	if card_number == "" {
		return nil, card_errors.ErrGrpcInvalidCardNumber
	}

	res, err := s.cardQuery.FindByCardNumber(ctx, card_number)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCard := &pb.CardResponse{
		Id:         int32(res.CardID),
		UserId:     int32(res.UserID),
		CardNumber: res.CardNumber,
		CardType:     res.CardType,
		CardProvider: res.CardProvider,
		Cvv:          res.Cvv,
		ExpireDate: res.ExpireDate.Time.Format(time.RFC3339),
		CreatedAt:  res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:  res.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully fetched card record",
		Data:    protoCard,
	}, nil
}

func (s *cardQueryHandleGrpc) FindByUserIdCard(ctx context.Context, req *pb.FindByUserIdCardRequest) (*pb.ApiResponseCard, error) {
	id := int(req.GetUserId())

	if id == 0 {
		return nil, card_errors.ErrGrpcInvalidUserID
	}
	res, err := s.cardQuery.FindByUserID(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	Pbres := &pb.ApiResponseCard{
		Message: "successfully",
		Status:  "success",
		Data: &pb.CardResponse{
			Id:           int32(res.CardID),
			UserId:       int32(res.UserID),
			CardNumber:   res.CardNumber,
			CardType:     res.CardType,
			CardProvider: res.CardProvider,
			Cvv:          res.Cvv,
			ExpireDate:   res.ExpireDate.Time.Format(time.RFC3339),
			CreatedAt:    res.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:    res.UpdatedAt.Time.Format(time.RFC3339),
		},
	}

	return Pbres, nil
}

func (s *cardQueryHandleGrpc) FindByActiveCard(ctx context.Context, req *pb.FindAllCardRequest) (*pb.ApiResponsePaginationCardDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllCards{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	cards, totalRecords, err := s.cardQuery.FindByActive(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCards := make([]*pb.CardResponseDeleteAt, len(cards))
	for i, card := range cards {
		var deletedAt *wrapperspb.StringValue
		if card.DeletedAt.Valid {
			deletedAt = wrapperspb.String(card.DeletedAt.Time.Format(time.RFC3339))
		}
		protoCards[i] = &pb.CardResponseDeleteAt{
			Id:           int32(card.CardID),
			UserId:       int32(card.UserID),
			CardNumber:   card.CardNumber,
			CardType:     card.CardType,
			Cvv:          card.Cvv,
			CardProvider: card.CardProvider,
			ExpireDate:   card.ExpireDate.Time.Format(time.RFC3339),
			CreatedAt:    card.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:    card.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:    deletedAt,
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationCardDeleteAt{
		Status:         "success",
		Message:        "Successfully fetched active card records",
		Data:           protoCards,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *cardQueryHandleGrpc) FindByTrashedCard(ctx context.Context, req *pb.FindAllCardRequest) (*pb.ApiResponsePaginationCardDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllCards{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	cards, totalRecords, err := s.cardQuery.FindByTrashed(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCards := make([]*pb.CardResponseDeleteAt, len(cards))
	for i, card := range cards {
		var deletedAt *wrapperspb.StringValue
		if card.DeletedAt.Valid {
			deletedAt = wrapperspb.String(card.DeletedAt.Time.Format(time.RFC3339))
		}
		protoCards[i] = &pb.CardResponseDeleteAt{
			Id:           int32(card.CardID),
			UserId:       int32(card.UserID),
			CardNumber:   card.CardNumber,
			CardType:     card.CardType,
			Cvv:          card.Cvv,
			CardProvider: card.CardProvider,
			ExpireDate:   card.ExpireDate.Time.Format(time.RFC3339),
			CreatedAt:    card.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:    card.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:    deletedAt,
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationCardDeleteAt{
		Status:         "success",
		Message:        "Successfully fetched trashed card records",
		Data:           protoCards,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *cardQueryHandleGrpc) FindUserCardByCardNumber(ctx context.Context, req *pb.FindByCardNumberRequest) (*pb.CardWithEmailResponse, error) {
	card_number := req.GetCardNumber()
	if card_number == "" {
		return nil, card_errors.ErrGrpcInvalidCardNumber
	}

	res, err := s.cardQuery.FindUserCardByCardNumber(ctx, card_number)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCard := &pb.CardWithEmailResponse{
		Id:         int32(res.CardID),
		UserId:     int32(res.UserID),
		Email:      res.Email,
		CardNumber: res.CardNumber,
		CardType:     res.CardType,
		Cvv:          res.Cvv,
		CardProvider: res.CardProvider,
		ExpireDate:   res.ExpireDate.Time.Format(time.RFC3339),
		CreatedAt:  res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:  res.UpdatedAt.Time.Format(time.RFC3339),
	}

	return protoCard, nil
}
