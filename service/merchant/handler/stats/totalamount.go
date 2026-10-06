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

type merchantStatsTotalAmountHandleGrpc struct {
	pb.UnimplementedMerchantStatsTotalAmountServiceServer

	amountstats           stats.MerchantStatsTotalAmountService
	amountstatsbymerchant statsbymerchant.MerchantStatsByMerchantTotalAmountService
	amountstatsbyapikey   statsbyapikey.MerchantStatsByApiKeyTotalAmountService
}

func NewMerchantStatsTotalAmountHandler(
	amountstats stats.MerchantStatsTotalAmountService,
	amountstatsbymerchant statsbymerchant.MerchantStatsByMerchantTotalAmountService,
	amountstatsbyapikey statsbyapikey.MerchantStatsByApiKeyTotalAmountService,

) MerchantStatsTotalAmountHandleGrpc {
	return &merchantStatsTotalAmountHandleGrpc{
		amountstats:           amountstats,
		amountstatsbymerchant: amountstatsbymerchant,
		amountstatsbyapikey:   amountstatsbyapikey,
	}
}

func (s *merchantStatsTotalAmountHandleGrpc) FindMonthlyTotalAmountMerchant(ctx context.Context, req *pbmerchant.FindYearMerchant) (*pb.ApiResponseMerchantMonthlyTotalAmount, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	res, err := s.amountstats.FindMonthlyTotalAmountMerchant(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseMonthlyTotalAmount, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseMonthlyTotalAmount{
			Month:       item.Month,
			Year:        item.Year,
			TotalAmount: int64(item.TotalAmount),
		}
	}

	return &pb.ApiResponseMerchantMonthlyTotalAmount{
		Status:  "success",
		Message: "Successfully fetched monthly amount for merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsTotalAmountHandleGrpc) FindYearlyTotalAmountMerchant(ctx context.Context, req *pbmerchant.FindYearMerchant) (*pb.ApiResponseMerchantYearlyTotalAmount, error) {
	year := int(req.GetYear())
	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	res, err := s.amountstats.FindYearlyTotalAmountMerchant(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseYearlyTotalAmount, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseYearlyTotalAmount{
			Year:        item.Year,
			TotalAmount: int64(item.TotalAmount),
		}
	}

	return &pb.ApiResponseMerchantYearlyTotalAmount{
		Status:  "success",
		Message: "Successfully fetched yearly amount for merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsTotalAmountHandleGrpc) FindMonthlyTotalAmountByMerchants(ctx context.Context, req *pbmerchant.FindYearMerchantById) (*pb.ApiResponseMerchantMonthlyTotalAmount, error) {
	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	if id <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidID
	}

	res, err := s.amountstatsbymerchant.FindMonthlyTotalAmountByMerchants(ctx, &requests.MonthYearTotalAmountMerchant{
		MerchantID: id,
		Year:       year,
	})

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseMonthlyTotalAmount, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseMonthlyTotalAmount{
			Month:       item.Month,
			Year:        item.Year,
			TotalAmount: int64(item.TotalAmount),
		}
	}

	return &pb.ApiResponseMerchantMonthlyTotalAmount{
		Status:  "success",
		Message: "Successfully fetched monthly amount for merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsTotalAmountHandleGrpc) FindYearlyTotalAmountByMerchants(ctx context.Context, req *pbmerchant.FindYearMerchantById) (*pb.ApiResponseMerchantYearlyTotalAmount, error) {
	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	if id <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidID
	}

	res, err := s.amountstatsbymerchant.FindYearlyTotalAmountByMerchants(ctx, &requests.MonthYearTotalAmountMerchant{
		MerchantID: id,
		Year:       year,
	})

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseYearlyTotalAmount, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseYearlyTotalAmount{
			Year:        item.Year,
			TotalAmount: int64(item.TotalAmount),
		}
	}

	return &pb.ApiResponseMerchantYearlyTotalAmount{
		Status:  "success",
		Message: "Successfully fetched yearly amount for merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsTotalAmountHandleGrpc) FindMonthlyTotalAmountByApikey(ctx context.Context, req *pbmerchant.FindYearMerchantByApikey) (*pb.ApiResponseMerchantMonthlyTotalAmount, error) {
	year := int(req.GetYear())
	apikey := req.GetApiKey()

	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	if apikey == "" {
		return nil, merchant_errors.ErrGrpcMerchantInvalidApiKey
	}

	res, err := s.amountstatsbyapikey.FindMonthlyTotalAmountByApikey(ctx, &requests.MonthYearTotalAmountApiKey{
		Year:   year,
		Apikey: apikey,
	})

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseMonthlyTotalAmount, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseMonthlyTotalAmount{
			Month:       item.Month,
			Year:        item.Year,
			TotalAmount: int64(item.TotalAmount),
		}
	}

	return &pb.ApiResponseMerchantMonthlyTotalAmount{
		Status:  "success",
		Message: "Successfully fetched monthly amount for merchant",
		Data:    protoData,
	}, nil
}

func (s *merchantStatsTotalAmountHandleGrpc) FindYearlyTotalAmountByApikey(ctx context.Context, req *pbmerchant.FindYearMerchantByApikey) (*pb.ApiResponseMerchantYearlyTotalAmount, error) {
	year := int(req.GetYear())
	apikey := req.GetApiKey()

	if year <= 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidYear
	}

	if apikey == "" {
		return nil, merchant_errors.ErrGrpcMerchantInvalidApiKey
	}

	res, err := s.amountstatsbyapikey.FindYearlyTotalAmountByApikey(ctx, &requests.MonthYearTotalAmountApiKey{
		Apikey: apikey,
		Year:   year,
	})

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoData := make([]*pb.MerchantResponseYearlyTotalAmount, len(res))
	for i, item := range res {
		protoData[i] = &pb.MerchantResponseYearlyTotalAmount{
			Year:        item.Year,
			TotalAmount: int64(item.TotalAmount),
		}
	}

	return &pb.ApiResponseMerchantYearlyTotalAmount{
		Status:  "success",
		Message: "Successfully fetched yearly amount for merchant",
		Data:    protoData,
	}, nil
}
