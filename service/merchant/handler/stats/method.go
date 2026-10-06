package merchantstatshandler

import (
	"context"

	stats "github.com/MamangRust/monolith-payment-gateway-merchant/service/stats"
	statsbyapikey "github.com/MamangRust/monolith-payment-gateway-merchant/service/statsbyapikey"
	statsbymerchant "github.com/MamangRust/monolith-payment-gateway-merchant/service/statsbymerchant"
	pbmerchant "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/grpc"
)

type merchantStatsMethodHandleGrpc struct {
	pb.UnimplementedMerchantStatsMethodServiceServer

	methodstats           stats.MerchantStatsMethodService
	methodstatsbymerchant statsbymerchant.MerchantStatsByMerchantMethodService
	methodstatsbyapikey   statsbyapikey.MerchantStatsByApiKeyMethodService
}

func NewMerchantStatsMethodHandler(
	methodstats stats.MerchantStatsMethodService,
	methodstatsbymerchant statsbymerchant.MerchantStatsByMerchantMethodService,
	methodstatsbyapikey statsbyapikey.MerchantStatsByApiKeyMethodService,

) MerchantStatsMethodHandleGrpc {
	return &merchantStatsMethodHandleGrpc{
		methodstats:           methodstats,
		methodstatsbymerchant: methodstatsbymerchant,
		methodstatsbyapikey:   methodstatsbyapikey,
	}
}

func (s *merchantStatsMethodHandleGrpc) FindMonthlyPaymentMethodsMerchant(ctx context.Context, req *pbmerchant.FindYearMerchant) (*pb.ApiResponseMerchantMonthlyPaymentMethod, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	res, err := s.methodstats.FindMonthlyPaymentMethodsMerchant(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseMonthlyPaymentMethod, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseMonthlyPaymentMethod{
			Month:         item.Month,
			PaymentMethod: item.PaymentMethod,
			TotalAmount:   int64(item.TotalAmount),
		}
	}

	return &pb.ApiResponseMerchantMonthlyPaymentMethod{
		Status:  "success",
		Message: "Successfully fetched monthly payment methods for merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsMethodHandleGrpc) FindYearlyPaymentMethodMerchant(ctx context.Context, req *pbmerchant.FindYearMerchant) (*pb.ApiResponseMerchantYearlyPaymentMethod, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	res, err := s.methodstats.FindYearlyPaymentMethodMerchant(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseYearlyPaymentMethod, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseYearlyPaymentMethod{
			Year:          item.Year.Int.String(),
			PaymentMethod: item.PaymentMethod,
			TotalAmount:   item.TotalAmount,
		}
	}

	return &pb.ApiResponseMerchantYearlyPaymentMethod{
		Status:  "success",
		Message: "Successfully fetched yearly payment methods for merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsMethodHandleGrpc) FindMonthlyPaymentMethodByMerchants(ctx context.Context, req *pbmerchant.FindYearMerchantById) (*pb.ApiResponseMerchantMonthlyPaymentMethod, error) {
	merchantId := req.GetMerchantId()
	year := req.GetYear()

	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	if merchantId <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidID
	}

	reqService := requests.MonthYearPaymentMethodMerchant{
		MerchantID: int(req.MerchantId),
		Year:       int(year),
	}

	res, err := s.methodstatsbymerchant.FindMonthlyPaymentMethodByMerchants(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseMonthlyPaymentMethod, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseMonthlyPaymentMethod{
			Month:         item.Month,
			PaymentMethod: item.PaymentMethod,
			TotalAmount:   int64(item.TotalAmount),
		}
	}

	return &pb.ApiResponseMerchantMonthlyPaymentMethod{
		Status:  "success",
		Message: "Successfully fetched monthly payment methods by merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsMethodHandleGrpc) FindYearlyPaymentMethodByMerchants(ctx context.Context, req *pbmerchant.FindYearMerchantById) (*pb.ApiResponseMerchantYearlyPaymentMethod, error) {
	merchantId := req.GetMerchantId()
	year := req.GetYear()

	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	if merchantId <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidID
	}

	reqService := requests.MonthYearPaymentMethodMerchant{
		MerchantID: int(req.MerchantId),
		Year:       int(year),
	}

	res, err := s.methodstatsbymerchant.FindYearlyPaymentMethodByMerchants(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseYearlyPaymentMethod, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseYearlyPaymentMethod{
			Year:          item.Year.Int.String(),
			PaymentMethod: item.PaymentMethod,
			TotalAmount:   item.TotalAmount,
		}
	}

	return &pb.ApiResponseMerchantYearlyPaymentMethod{
		Status:  "success",
		Message: "Successfully fetched yearly payment methods by merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsMethodHandleGrpc) FindMonthlyPaymentMethodByApikey(ctx context.Context, req *pbmerchant.FindYearMerchantByApikey) (*pb.ApiResponseMerchantMonthlyPaymentMethod, error) {
	api_key := req.GetApiKey()
	year := req.GetYear()

	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	if api_key == "" {
		return nil, merchant_errors.ErrGrpcMerchantInvalidApiKey
	}

	reqService := requests.MonthYearPaymentMethodApiKey{
		Year:   int(year),
		Apikey: api_key,
	}

	res, err := s.methodstatsbyapikey.FindMonthlyPaymentMethodByApikey(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseMonthlyPaymentMethod, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseMonthlyPaymentMethod{
			Month:         item.Month,
			PaymentMethod: item.PaymentMethod,
			TotalAmount:   int64(item.TotalAmount),
		}
	}

	return &pb.ApiResponseMerchantMonthlyPaymentMethod{
		Status:  "success",
		Message: "Successfully fetched monthly payment methods by merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsMethodHandleGrpc) FindYearlyPaymentMethodByApikey(ctx context.Context, req *pbmerchant.FindYearMerchantByApikey) (*pb.ApiResponseMerchantYearlyPaymentMethod, error) {
	api_key := req.GetApiKey()
	year := req.GetYear()

	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	if api_key == "" {
		return nil, merchant_errors.ErrGrpcMerchantInvalidApiKey
	}

	reqService := requests.MonthYearPaymentMethodApiKey{
		Year:   int(year),
		Apikey: api_key,
	}

	res, err := s.methodstatsbyapikey.FindYearlyPaymentMethodByApikey(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseYearlyPaymentMethod, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseYearlyPaymentMethod{
			Year:          item.Year.Int.String(),
			PaymentMethod: item.PaymentMethod,
			TotalAmount:   item.TotalAmount,
		}
	}

	return &pb.ApiResponseMerchantYearlyPaymentMethod{
		Status:  "success",
		Message: "Successfully fetched yearly payment methods by merchant",
		Data:    protoData,
	}, nil
}
