package handler

import (
	"context"
	"math"
	"time"

	pbhelpers "github.com/MamangRust/monolith-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/grpc"
	"github.com/MamangRust/monolith-payment-gateway-transaction/service"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type transactionQueryHandleGrpc struct {
	pb.UnimplementedTransactionQueryServiceServer

	service service.TransactionQueryService
}

func NewTransactionQueryHandleGrpc(service service.TransactionQueryService) TransactionQueryHandleGrpc {
	return &transactionQueryHandleGrpc{
		service: service,
	}
}

func (t *transactionQueryHandleGrpc) FindAllTransaction(ctx context.Context, request *pb.FindAllTransactionRequest) (*pb.ApiResponsePaginationTransaction, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllTransactions{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	transactions, totalRecords, err := t.service.FindAll(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbhelpers.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	transactionResponses := make([]*pb.TransactionResponse, len(transactions))
	for i, transaction := range transactions {
		transactionResponses[i] = &pb.TransactionResponse{
			Id:              int32(transaction.TransactionID),
			CardNumber:      transaction.CardNumber,
			TransactionNo:   transaction.TransactionNo.String(),
			Amount:          int32(transaction.Amount),
			PaymentMethod:   transaction.PaymentMethod,
			MerchantId:      int32(transaction.MerchantID),
			TransactionTime: transaction.TransactionTime.Format(time.RFC3339),
			CreatedAt:       transaction.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       transaction.UpdatedAt.Time.Format(time.RFC3339),
		}
	}

	return &pb.ApiResponsePaginationTransaction{
		Status:         "success",
		Message:        "Successfully fetched transaction records",
		Data:           transactionResponses,
		PaginationMeta: paginationMeta,
	}, nil
}

func (t *transactionQueryHandleGrpc) FindAllTransactionByCardNumber(ctx context.Context, request *pb.FindAllTransactionCardNumberRequest) (*pb.ApiResponsePaginationTransaction, error) {
	card_number := request.GetCardNumber()
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	reqService := requests.FindAllTransactionCardNumber{
		CardNumber: card_number,
		Page:       page,
		PageSize:   pageSize,
		Search:     search,
	}

	transactions, totalRecords, err := t.service.FindAllByCardNumber(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbhelpers.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	transactionResponses := make([]*pb.TransactionResponse, len(transactions))
	for i, transaction := range transactions {
		transactionResponses[i] = &pb.TransactionResponse{
			Id:              int32(transaction.TransactionID),
			CardNumber:      transaction.CardNumber,
			TransactionNo:   transaction.TransactionNo.String(),
			Amount:          int32(transaction.Amount),
			PaymentMethod:   transaction.PaymentMethod,
			MerchantId:      int32(transaction.MerchantID),
			TransactionTime: transaction.TransactionTime.Format(time.RFC3339),
			CreatedAt:       transaction.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       transaction.UpdatedAt.Time.Format(time.RFC3339),
		}
	}

	return &pb.ApiResponsePaginationTransaction{
		Status:         "success",
		Message:        "Successfully fetched transaction records",
		Data:           transactionResponses,
		PaginationMeta: paginationMeta,
	}, nil
}

func (t *transactionQueryHandleGrpc) FindByIdTransaction(ctx context.Context, req *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransaction, error) {
	id := int(req.GetTransactionId())

	if id == 0 {
		return nil, transaction_errors.ErrGrpcTransactionInvalidID
	}

	transaction, err := t.service.FindById(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Transaction fetched successfully",
		Data: &pb.TransactionResponse{
			Id:              int32(transaction.TransactionID),
			CardNumber:      transaction.CardNumber,
			TransactionNo:   transaction.TransactionNo.String(),
			Amount:          int32(transaction.Amount),
			PaymentMethod:   transaction.PaymentMethod,
			MerchantId:      int32(transaction.MerchantID),
			TransactionTime: transaction.TransactionTime.Format(time.RFC3339),
			CreatedAt:       transaction.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       transaction.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (t *transactionQueryHandleGrpc) FindByActiveTransaction(ctx context.Context, req *pb.FindAllTransactionRequest) (*pb.ApiResponsePaginationTransactionDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllTransactions{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	transactions, totalRecords, err := t.service.FindByActive(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbhelpers.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	transactionResponses := make([]*pb.TransactionResponseDeleteAt, len(transactions))
	for i, transaction := range transactions {
		transactionResponses[i] = &pb.TransactionResponseDeleteAt{
			Id:              int32(transaction.TransactionID),
			CardNumber:      transaction.CardNumber,
			TransactionNo:   transaction.TransactionNo.String(),
			Amount:          int32(transaction.Amount),
			PaymentMethod:   transaction.PaymentMethod,
			MerchantId:      int32(transaction.MerchantID),
			TransactionTime: transaction.TransactionTime.Format(time.RFC3339),
			CreatedAt:       transaction.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       transaction.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:       &wrapperspb.StringValue{Value: transaction.DeletedAt.Time.Format(time.RFC3339)},
		}
	}

	return &pb.ApiResponsePaginationTransactionDeleteAt{
		Status:         "success",
		Message:        "Successfully fetch transactions",
		Data:           transactionResponses,
		PaginationMeta: paginationMeta,
	}, nil
}

func (t *transactionQueryHandleGrpc) FindByTrashedTransaction(ctx context.Context, req *pb.FindAllTransactionRequest) (*pb.ApiResponsePaginationTransactionDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllTransactions{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	transactions, totalRecords, err := t.service.FindByTrashed(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbhelpers.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	transactionResponses := make([]*pb.TransactionResponseDeleteAt, len(transactions))
	for i, transaction := range transactions {
		transactionResponses[i] = &pb.TransactionResponseDeleteAt{
			Id:              int32(transaction.TransactionID),
			CardNumber:      transaction.CardNumber,
			TransactionNo:   transaction.TransactionNo.String(),
			Amount:          int32(transaction.Amount),
			PaymentMethod:   transaction.PaymentMethod,
			MerchantId:      int32(transaction.MerchantID),
			TransactionTime: transaction.TransactionTime.Format(time.RFC3339),
			CreatedAt:       transaction.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       transaction.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:       &wrapperspb.StringValue{Value: transaction.DeletedAt.Time.Format(time.RFC3339)},
		}
	}

	return &pb.ApiResponsePaginationTransactionDeleteAt{
		Status:         "success",
		Message:        "Successfully fetch transactions",
		Data:           transactionResponses,
		PaginationMeta: paginationMeta,
	}, nil
}
