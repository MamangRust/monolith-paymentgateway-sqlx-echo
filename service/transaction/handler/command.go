package handler

import (
	"context"
	"time"

	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/grpc"
	"github.com/MamangRust/monolith-payment-gateway-transaction/service"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type transactionCommandHandleGrpc struct {
	pb.UnimplementedTransactionCommandServiceServer

	service service.TransactionCommandService
}

func NewTransactionCommandHandleGrpc(service service.TransactionCommandService) TransactionCommandHandleGrpc {
	return &transactionCommandHandleGrpc{
		service: service,
	}
}

func (t *transactionCommandHandleGrpc) CreateTransaction(ctx context.Context, request *pb.CreateTransactionRequest) (*pb.ApiResponseTransaction, error) {
	transactionTime := request.GetTransactionTime().AsTime()
	merchantID := int(request.GetMerchantId())

	req := requests.CreateTransactionRequest{
		CardNumber:      request.GetCardNumber(),
		Amount:          int(request.GetAmount()),
		PaymentMethod:   request.GetPaymentMethod(),
		MerchantID:      &merchantID,
		TransactionTime: transactionTime,
		IdempotencyKey:  request.GetIdempotencyKey(),
	}

	if err := req.Validate(); err != nil {
		return nil, transaction_errors.ErrGrpcValidateCreateTransactionRequest
	}

	res, err := t.service.Create(ctx, request.ApiKey, &req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully created transaction",
		Data: &pb.TransactionResponse{
			Id:              int32(res.TransactionID),
			CardNumber:      res.CardNumber,
			TransactionNo:   res.TransactionNo.String(),
			Amount:          int32(res.Amount),
			PaymentMethod:   res.PaymentMethod,
			MerchantId:      int32(res.MerchantID),
			TransactionTime: res.TransactionTime.Format(time.RFC3339),
			CreatedAt:       res.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       res.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (t *transactionCommandHandleGrpc) UpdateTransaction(ctx context.Context, request *pb.UpdateTransactionRequest) (*pb.ApiResponseTransaction, error) {
	id := int(request.GetTransactionId())

	if id <= 0 {
		return nil, transaction_errors.ErrGrpcTransactionInvalidID
	}

	transactionTime := request.GetTransactionTime().AsTime()
	merchantID := int(request.GetMerchantId())

	req := requests.UpdateTransactionRequest{
		TransactionID:   &id,
		CardNumber:      request.GetCardNumber(),
		Amount:          int(request.GetAmount()),
		PaymentMethod:   request.GetPaymentMethod(),
		MerchantID:      &merchantID,
		TransactionTime: transactionTime,
	}

	if err := req.Validate(); err != nil {
		return nil, transaction_errors.ErrGrpcValidateUpdateTransactionRequest
	}

	res, err := t.service.Update(ctx, request.ApiKey, &req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully updated transaction",
		Data: &pb.TransactionResponse{
			Id:              int32(res.TransactionID),
			CardNumber:      res.CardNumber,
			TransactionNo:   res.TransactionNo.String(),
			Amount:          int32(res.Amount),
			PaymentMethod:   res.PaymentMethod,
			MerchantId:      int32(res.MerchantID),
			TransactionTime: res.TransactionTime.Format(time.RFC3339),
			CreatedAt:       res.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       res.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (t *transactionCommandHandleGrpc) TrashedTransaction(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDeleteAt, error) {
	id := int(request.GetTransactionId())

	if id <= 0 {
		return nil, transaction_errors.ErrGrpcTransactionInvalidID
	}

	res, err := t.service.TrashedTransaction(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransactionDeleteAt{
		Status:  "success",
		Message: "Successfully trashed transaction",
		Data: &pb.TransactionResponseDeleteAt{
			Id:              int32(res.TransactionID),
			CardNumber:      res.CardNumber,
			TransactionNo:   res.TransactionNo.String(),
			Amount:          int32(res.Amount),
			PaymentMethod:   res.PaymentMethod,
			MerchantId:      int32(res.MerchantID),
			TransactionTime: res.TransactionTime.Format(time.RFC3339),
			CreatedAt:       res.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       res.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:       &wrapperspb.StringValue{Value: res.DeletedAt.Time.Format(time.RFC3339)},
		},
	}, nil
}

func (t *transactionCommandHandleGrpc) RestoreTransaction(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDeleteAt, error) {
	id := int(request.GetTransactionId())

	if id <= 0 {
		return nil, transaction_errors.ErrGrpcTransactionInvalidID
	}

	res, err := t.service.RestoreTransaction(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransactionDeleteAt{
		Status:  "success",
		Message: "Successfully restored transaction",
		Data: &pb.TransactionResponseDeleteAt{
			Id:              int32(res.TransactionID),
			CardNumber:      res.CardNumber,
			TransactionNo:   res.TransactionNo.String(),
			Amount:          int32(res.Amount),
			PaymentMethod:   res.PaymentMethod,
			MerchantId:      int32(res.MerchantID),
			TransactionTime: res.TransactionTime.Format(time.RFC3339),
			CreatedAt:       res.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       res.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:       &wrapperspb.StringValue{Value: res.DeletedAt.Time.Format(time.RFC3339)},
		},
	}, nil
}

func (t *transactionCommandHandleGrpc) DeleteTransactionPermanent(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDelete, error) {
	id := int(request.GetTransactionId())

	if id <= 0 {
		return nil, transaction_errors.ErrGrpcTransactionInvalidID
	}

	_, err := t.service.DeleteTransactionPermanent(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransactionDelete{
		Status:  "success",
		Message: "Successfully deleted transaction",
	}, nil
}

func (t *transactionCommandHandleGrpc) RestoreAllTransaction(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransactionAll, error) {
	_, err := t.service.RestoreAllTransaction(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransactionAll{
		Status:  "success",
		Message: "Successfully restore all transaction",
	}, nil
}

func (t *transactionCommandHandleGrpc) DeleteAllTransactionPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransactionAll, error) {
	_, err := t.service.DeleteAllTransactionPermanent(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransactionAll{
		Status:  "success",
		Message: "Successfully delete transaction permanent",
	}, nil
}

func (t *transactionCommandHandleGrpc) AuthorizeTransaction(ctx context.Context, request *pb.CreateTransactionRequest) (*pb.ApiResponseTransaction, error) {
	transactionTime := request.GetTransactionTime().AsTime()
	merchantID := int(request.GetMerchantId())

	req := requests.CreateTransactionRequest{
		CardNumber:      request.GetCardNumber(),
		Amount:          int(request.GetAmount()),
		PaymentMethod:   request.GetPaymentMethod(),
		MerchantID:      &merchantID,
		TransactionTime: transactionTime,
		IdempotencyKey:  request.GetIdempotencyKey(),
	}

	if err := req.Validate(); err != nil {
		return nil, transaction_errors.ErrGrpcValidateCreateTransactionRequest
	}

	res, err := t.service.AuthorizeTransaction(ctx, request.ApiKey, &req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully authorized transaction",
		Data: &pb.TransactionResponse{
			Id:              int32(res.TransactionID),
			CardNumber:      res.CardNumber,
			TransactionNo:   res.TransactionNo.String(),
			Amount:          int32(res.Amount),
			PaymentMethod:   res.PaymentMethod,
			MerchantId:      int32(res.MerchantID),
			TransactionTime: res.TransactionTime.Format(time.RFC3339),
			CreatedAt:       res.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       res.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (t *transactionCommandHandleGrpc) CaptureTransaction(ctx context.Context, request *pb.CaptureTransactionRequest) (*pb.ApiResponseTransaction, error) {
	id := int(request.GetTransactionId())
	if id <= 0 {
		return nil, transaction_errors.ErrGrpcTransactionInvalidID
	}

	res, err := t.service.CaptureTransaction(ctx, request.GetApiKey(), id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully captured transaction",
		Data: &pb.TransactionResponse{
			Id:              int32(res.TransactionID),
			CardNumber:      res.CardNumber,
			TransactionNo:   res.TransactionNo.String(),
			Amount:          int32(res.Amount),
			PaymentMethod:   res.PaymentMethod,
			MerchantId:      int32(res.MerchantID),
			TransactionTime: res.TransactionTime.Format(time.RFC3339),
			CreatedAt:       res.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       res.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (t *transactionCommandHandleGrpc) VoidTransaction(ctx context.Context, request *pb.VoidTransactionRequest) (*pb.ApiResponseTransaction, error) {
	id := int(request.GetTransactionId())
	if id <= 0 {
		return nil, transaction_errors.ErrGrpcTransactionInvalidID
	}

	res, err := t.service.VoidTransaction(ctx, request.GetApiKey(), id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully voided transaction",
		Data: &pb.TransactionResponse{
			Id:              int32(res.TransactionID),
			CardNumber:      res.CardNumber,
			TransactionNo:   res.TransactionNo.String(),
			Amount:          int32(res.Amount),
			PaymentMethod:   res.PaymentMethod,
			MerchantId:      int32(res.MerchantID),
			TransactionTime: res.TransactionTime.Format(time.RFC3339),
			CreatedAt:       res.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       res.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (t *transactionCommandHandleGrpc) RefundTransaction(ctx context.Context, request *pb.RefundTransactionRequest) (*pb.ApiResponseTransaction, error) {
	id := int(request.GetTransactionId())
	if id <= 0 {
		return nil, transaction_errors.ErrGrpcTransactionInvalidID
	}

	res, err := t.service.RefundTransaction(ctx, request.GetApiKey(), id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully refunded transaction",
		Data: &pb.TransactionResponse{
			Id:              int32(res.TransactionID),
			CardNumber:      res.CardNumber,
			TransactionNo:   res.TransactionNo.String(),
			Amount:          int32(res.Amount),
			PaymentMethod:   res.PaymentMethod,
			MerchantId:      int32(res.MerchantID),
			TransactionTime: res.TransactionTime.Format(time.RFC3339),
			CreatedAt:       res.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       res.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}
