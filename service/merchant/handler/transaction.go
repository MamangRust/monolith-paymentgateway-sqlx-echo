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

type merchantTransactionHandleGrpc struct {
	pb.UnimplementedMerchantTransactionServiceServer

	merchantTransaction service.MerchantTransactionService
}

func NewMerchantTransactionHandleGrpc(merchantTransaction service.MerchantTransactionService) MerchantTransactionHandleGrpc {
	return &merchantTransactionHandleGrpc{
		merchantTransaction: merchantTransaction,
	}
}

func (s *merchantTransactionHandleGrpc) FindAllTransactionMerchant(ctx context.Context, req *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchantTransaction, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchantTransactions{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	transactions, totalRecords, err := s.merchantTransaction.FindAllTransactions(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTransactions := make([]*pb.MerchantTransactionResponse, len(transactions))
	for i, txn := range transactions {
		var deletedAt *wrapperspb.StringValue
		if txn.DeletedAt.Valid {
			deletedAt = wrapperspb.String(txn.DeletedAt.Time.Format(time.RFC3339))
		}

		protoTransactions[i] = &pb.MerchantTransactionResponse{
			Id:              int32(txn.TransactionID),
			CardNumber:      txn.CardNumber,
			Amount:          int32(txn.Amount),
			PaymentMethod:   txn.PaymentMethod,
			MerchantId:      int32(txn.MerchantID),
			MerchantName:    txn.MerchantName,
			TransactionTime: txn.TransactionTime.Format(time.RFC3339),
			CreatedAt:       txn.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       txn.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:       deletedAt,
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))
	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationMerchantTransaction{
		Status:         "success",
		Message:        "Successfully fetched merchant record",
		Data:           protoTransactions,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *merchantTransactionHandleGrpc) FindAllTransactionByMerchant(ctx context.Context, req *pb.FindAllMerchantTransaction) (*pb.ApiResponsePaginationMerchantTransaction, error) {
	merchant_id := int(req.MerchantId)
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if merchant_id <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidID
	}

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchantTransactionsById{
		Page:       page,
		PageSize:   pageSize,
		Search:     search,
		MerchantID: merchant_id,
	}

	transactions, totalRecords, err := s.merchantTransaction.FindAllTransactionsByMerchant(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTransactions := make([]*pb.MerchantTransactionResponse, len(transactions))
	for i, txn := range transactions {
		var deletedAt *wrapperspb.StringValue
		if txn.DeletedAt.Valid {
			deletedAt = wrapperspb.String(txn.DeletedAt.Time.Format(time.RFC3339))
		}

		protoTransactions[i] = &pb.MerchantTransactionResponse{
			Id:              int32(txn.TransactionID),
			CardNumber:      txn.CardNumber,
			Amount:          int32(txn.Amount),
			PaymentMethod:   txn.PaymentMethod,
			MerchantId:      int32(txn.MerchantID),
			MerchantName:    txn.MerchantName,
			TransactionTime: txn.TransactionTime.Format(time.RFC3339),
			CreatedAt:       txn.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       txn.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:       deletedAt,
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))
	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationMerchantTransaction{
		Status:         "success",
		Message:        "Successfully fetched merchant record",
		Data:           protoTransactions,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *merchantTransactionHandleGrpc) FindAllTransactionByApikey(ctx context.Context, req *pb.FindAllMerchantApikey) (*pb.ApiResponsePaginationMerchantTransaction, error) {
	api_key := req.GetApiKey()
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if api_key == "" {
		return nil, merchant_errors.ErrGrpcMerchantInvalidApiKey
	}

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchantTransactionsByApiKey{
		ApiKey:   api_key,
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	transactions, totalRecords, err := s.merchantTransaction.FindAllTransactionsByApikey(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTransactions := make([]*pb.MerchantTransactionResponse, len(transactions))
	for i, txn := range transactions {
		var deletedAt *wrapperspb.StringValue
		if txn.DeletedAt.Valid {
			deletedAt = wrapperspb.String(txn.DeletedAt.Time.Format(time.RFC3339))
		}

		protoTransactions[i] = &pb.MerchantTransactionResponse{
			Id:              int32(txn.TransactionID),
			CardNumber:      txn.CardNumber,
			Amount:          int32(txn.Amount),
			PaymentMethod:   txn.PaymentMethod,
			MerchantId:      int32(txn.MerchantID),
			MerchantName:    txn.MerchantName,
			TransactionTime: txn.TransactionTime.Format(time.RFC3339),
			CreatedAt:       txn.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       txn.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:       deletedAt,
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))
	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationMerchantTransaction{
		Status:         "success",
		Message:        "Successfully fetched merchant record",
		Data:           protoTransactions,
		PaginationMeta: paginationMeta,
	}, nil
}
