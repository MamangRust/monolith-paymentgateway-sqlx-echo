package authhandler

import (
	auth_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/auth"
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	authapimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/auth"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type DepsAuth struct {
	Client     *grpc.ClientConn
	E          *echo.Echo
	Logger     logger.LoggerInterface
	Cache      *cache.CacheStore
	ApiHandler errors.ApiHandler
}

func RegisterAuthHandler(deps *DepsAuth) {
	mapper := authapimapper.NewAuthResponseMapper()

	cache := auth_cache.NewMencache(deps.Cache)

	NewHandlerAuth(&authHandleParams{
		client:     pb.NewAuthServiceClient(deps.Client),
		router:     deps.E,
		logger:     deps.Logger,
		mapper:     mapper,
		cache:      cache,
		apiHandler: deps.ApiHandler,
	})
}
