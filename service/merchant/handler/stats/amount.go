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

type merchantStatsAmountHandleGrpc struct {
	pb.UnimplementedMerchantStatsAmountServiceServer

	amountstats           stats.MerchantStatsAmountService
	amountstatsbymerchant statsbymerchant.MerchantStatsByMerchantAmountService
	amountstatsbyapikey   statsbyapikey.MerchantStatsByApiKeyAmountService
}

func NewMerchantStatsAmountHandler(
	amountstats stats.MerchantStatsAmountService,
	amountstatsbymerchant statsbymerchant.MerchantStatsByMerchantAmountService,
	amountstatsbyapikey statsbyapikey.MerchantStatsByApiKeyAmountService,
) MerchantStatsAmountHandleGrpc {
	return &merchantStatsAmountHandleGrpc{
		amountstats:           amountstats,
		amountstatsbymerchant: amountstatsbymerchant,
		amountstatsbyapikey:   amountstatsbyapikey,
	}
}

func (s *merchantStatsAmountHandleGrpc) FindMonthlyAmountMerchant(ctx context.Context, req *pbmerchant.FindYearMerchant) (*pb.ApiResponseMerchantMonthlyAmount, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	res, err := s.amountstats.FindMonthlyAmountMerchant(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseMonthlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseMonthlyAmount{
			Month:       item.Month,
			TotalAmount: int64(item.TotalAmount),
		}
	}

	return &pb.ApiResponseMerchantMonthlyAmount{
		Status:  "success",
		Message: "Successfully fetched monthly amount for merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsAmountHandleGrpc) FindYearlyAmountMerchant(ctx context.Context, req *pbmerchant.FindYearMerchant) (*pb.ApiResponseMerchantYearlyAmount, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	res, err := s.amountstats.FindYearlyAmountMerchant(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseYearlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseYearlyAmount{
			Year:        item.Year.Int.String(),
			TotalAmount: item.TotalAmount,
		}
	}

	return &pb.ApiResponseMerchantYearlyAmount{
		Status:  "success",
		Message: "Successfully fetched yearly amount for merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsAmountHandleGrpc) FindMonthlyAmountByMerchants(ctx context.Context, req *pbmerchant.FindYearMerchantById) (*pb.ApiResponseMerchantMonthlyAmount, error) {
	merchantId := req.GetMerchantId()
	year := req.GetYear()

	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	if merchantId <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidID
	}

	reqService := requests.MonthYearAmountMerchant{
		MerchantID: int(req.MerchantId),
		Year:       int(year),
	}

	res, err := s.amountstatsbymerchant.FindMonthlyAmountByMerchants(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseMonthlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseMonthlyAmount{
			Month:       item.Month,
			TotalAmount: int64(item.TotalAmount),
		}
	}

	return &pb.ApiResponseMerchantMonthlyAmount{
		Status:  "success",
		Message: "Successfully fetched monthly amount by merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsAmountHandleGrpc) FindYearlyAmountByMerchants(ctx context.Context, req *pbmerchant.FindYearMerchantById) (*pb.ApiResponseMerchantYearlyAmount, error) {
	merchantId := req.GetMerchantId()
	year := req.GetYear()

	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	if merchantId <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidID
	}

	reqService := requests.MonthYearAmountMerchant{
		MerchantID: int(req.MerchantId),
		Year:       int(year),
	}
	res, err := s.amountstatsbymerchant.FindYearlyAmountByMerchants(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseYearlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseYearlyAmount{
			Year:        item.Year.Int.String(),
			TotalAmount: item.TotalAmount,
		}
	}

	return &pb.ApiResponseMerchantYearlyAmount{
		Status:  "success",
		Message: "Successfully fetched yearly amount by merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsAmountHandleGrpc) FindMonthlyAmountByApikey(ctx context.Context, req *pbmerchant.FindYearMerchantByApikey) (*pb.ApiResponseMerchantMonthlyAmount, error) {
	api_key := req.GetApiKey()
	year := req.GetYear()

	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	if api_key == "" {
		return nil, merchant_errors.ErrGrpcMerchantInvalidApiKey
	}

	reqService := requests.MonthYearAmountApiKey{
		Apikey: api_key,
		Year:   int(year),
	}

	res, err := s.amountstatsbyapikey.FindMonthlyAmountByApikey(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseMonthlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseMonthlyAmount{
			Month:       item.Month,
			TotalAmount: int64(item.TotalAmount),
		}
	}

	return &pb.ApiResponseMerchantMonthlyAmount{
		Status:  "success",
		Message: "Successfully fetched monthly amount by merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsAmountHandleGrpc) FindYearlyAmountByApikey(ctx context.Context, req *pbmerchant.FindYearMerchantByApikey) (*pb.ApiResponseMerchantYearlyAmount, error) {
	api_key := req.GetApiKey()
	year := req.GetYear()

	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	if api_key == "" {
		return nil, merchant_errors.ErrGrpcMerchantInvalidApiKey
	}

	reqService := requests.MonthYearAmountApiKey{
		Apikey: api_key,
		Year:   int(year),
	}

	res, err := s.amountstatsbyapikey.FindYearlyAmountByApikey(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseYearlyAmount, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseYearlyAmount{
			Year:        item.Year.Int.String(),
			TotalAmount: item.TotalAmount,
		}
	}

	return &pb.ApiResponseMerchantYearlyAmount{
		Status:  "success",
		Message: "Successfully fetched yearly amount by merchant",
		Data:    protoData,
	}, nil
}
