package handler

import (
	authhandler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/auth"
	cardhandler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/card"
	merchanthandler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/merchant"
	merchantdocumenthandler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/merchantdocument"
	rolehandler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/role"
	saldohandler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/saldo"
	topuphandler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/topup"
	transactionhandler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/transaction"
	transferhandler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/transfer"
	userhandler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/user"
	withdrawhandler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/withdraw"
	mencache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis"
	"github.com/MamangRust/monolith-payment-gateway-pkg/auth"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

// ServiceConnections aggregates gRPC connections to backend services.
type ServiceConnections struct {
	Auth        *grpc.ClientConn
	Role        *grpc.ClientConn
	Card        *grpc.ClientConn
	Merchant    *grpc.ClientConn
	User        *grpc.ClientConn
	Saldo       *grpc.ClientConn
	Topup       *grpc.ClientConn
	Transaction *grpc.ClientConn
	Transfer    *grpc.ClientConn
	Withdraw    *grpc.ClientConn
}

type Deps struct {
	Kafka              *kafka.Kafka
	Token              auth.TokenManager
	E                  *echo.Echo
	Logger             logger.LoggerInterface
	ServiceConnections *ServiceConnections
	Cache              *cache.CacheStore
}

func NewHandler(deps *Deps) {
	observability, _ := observability.NewObservability("apigateway", deps.Logger)

	apiHandler := errors.NewApiHandler(observability, deps.Logger)

	cache_apigateway := mencache.NewCacheApiGateway(deps.Cache)

	authhandler.RegisterAuthHandler(&authhandler.DepsAuth{
		Client:     deps.ServiceConnections.Auth,
		E:          deps.E,
		Logger:     deps.Logger,
		Cache:      deps.Cache,
		ApiHandler: apiHandler,
	})

	cardhandler.RegisterCardHandler(&cardhandler.DepsCard{
		Client:     deps.ServiceConnections.Card,
		E:          deps.E,
		Logger:     deps.Logger,
		Cache:      deps.Cache,
		ApiHandler: apiHandler,
	})

	merchanthandler.RegisterMerchantHandler(&merchanthandler.DepsMerchant{
		Client:     deps.ServiceConnections.Merchant,
		E:          deps.E,
		Logger:     deps.Logger,
		Cache:      deps.Cache,
		ApiHandler: apiHandler,
	})

	merchantdocumenthandler.RegisterMerchantDocumentHandler(&merchantdocumenthandler.DepsMerchantDocument{
		Client:     deps.ServiceConnections.Merchant,
		E:          deps.E,
		Logger:     deps.Logger,
		Cache:      deps.Cache,
		ApiHandler: apiHandler,
	})

	rolehandler.RegisterRoleHandler(&rolehandler.DepsRole{
		Kafka:      deps.Kafka,
		Client:     deps.ServiceConnections.Role,
		E:          deps.E,
		Logger:     deps.Logger,
		CacheStore: deps.Cache,
		Cache:      cache_apigateway,
		ApiHandler: apiHandler,
	})

	saldohandler.RegisterSaldoHandler(&saldohandler.DepsSaldo{
		Client:     deps.ServiceConnections.Saldo,
		E:          deps.E,
		Logger:     deps.Logger,
		Cache:      deps.Cache,
		ApiHandler: apiHandler,
	})

	topuphandler.RegisterTopupHandler(&topuphandler.DepsTopup{
		Client:     deps.ServiceConnections.Topup,
		E:          deps.E,
		Logger:     deps.Logger,
		Cache:      deps.Cache,
		ApiHandler: apiHandler,
	})

	transactionhandler.RegisterTransactionHandler(&transactionhandler.DepsTransaction{
		Kafka:           deps.Kafka,
		Client:          deps.ServiceConnections.Transaction,
		E:               deps.E,
		Logger:          deps.Logger,
		Cache:           deps.Cache,
		ApiHandler:      apiHandler,
		CacheApiGateway: cache_apigateway,
	})

	transferhandler.RegisterTransferHandler(&transferhandler.DepsTransfer{
		Client:     deps.ServiceConnections.Transfer,
		E:          deps.E,
		Logger:     deps.Logger,
		Cache:      deps.Cache,
		ApiHandler: apiHandler,
	})

	userhandler.RegisterUserHandler(&userhandler.DepsUser{
		Client:     deps.ServiceConnections.User,
		E:          deps.E,
		Logger:     deps.Logger,
		Cache:      deps.Cache,
		ApiHandler: apiHandler,
	})

	withdrawhandler.RegisterWithdrawHandler(&withdrawhandler.DepsWithdraw{
		Client:     deps.ServiceConnections.Withdraw,
		E:          deps.E,
		Logger:     deps.Logger,
		Cache:      deps.Cache,
		ApiHandler: apiHandler,
	})
}
