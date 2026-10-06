package transaction_test

import (
	"net"

	card_handler "github.com/MamangRust/monolith-payment-gateway-card/handler"
	card_repo "github.com/MamangRust/monolith-payment-gateway-card/repository"
	card_service "github.com/MamangRust/monolith-payment-gateway-card/service"
	merchant_handler "github.com/MamangRust/monolith-payment-gateway-merchant/handler"
	merchant_repo "github.com/MamangRust/monolith-payment-gateway-merchant/repository"
	merchant_service "github.com/MamangRust/monolith-payment-gateway-merchant/service"
	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pbmerchant "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	pbrole "github.com/MamangRust/monolith-payment-gateway-pb/role"
	pbsaldo "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	saldo_handler "github.com/MamangRust/monolith-payment-gateway-saldo/handler"
	saldo_repo "github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	saldo_service "github.com/MamangRust/monolith-payment-gateway-saldo/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	user_repo "github.com/MamangRust/monolith-payment-gateway-user/repository"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// grpcDeps wires the in-process gRPC dependency services that the transaction
// repository now consumes through shared adapters: card, saldo and merchant.
// Every service is backed by the same test database, so seeded rows are visible
// through both the local repositories and the remote gRPC clients.
type grpcDeps struct {
	identity *tests.IdentityStack
	server   *grpc.Server
	conn     *grpc.ClientConn

	cardQueryClient     pbcard.CardQueryServiceClient
	cardCommandClient   pbcard.CardCommandServiceClient
	saldoQueryClient    pbsaldo.SaldoQueryServiceClient
	saldoCommandClient  pbsaldo.SaldoCommandServiceClient
	merchantQueryClient pbmerchant.MerchantQueryServiceClient

	// Local repositories used for seeding. They are the same implementations
	// the gRPC services above are built from.
	cardRepo     *card_repo.Repositories
	saldoRepo    saldo_repo.Repositories
	merchantRepo merchant_repo.Repositories
	userRepo     user_repo.Repositories
}

// startGRPCDeps starts the identity stack plus card, saldo and merchant gRPC
// services on a single in-process server. grpc.NewClient dials lazily, so all
// services are registered before Serve is called.
func startGRPCDeps(db *sqlx.DB, cacheStore *cache.CacheStore, log logger.LoggerInterface) (*grpcDeps, error) {
	identity, err := tests.StartIdentityStack(db, cacheStore, log)
	if err != nil {
		return nil, err
	}

	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		identity.Close()
		return nil, err
	}

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		identity.Close()
		return nil, err
	}

	cardQueryClient := pbcard.NewCardQueryServiceClient(conn)
	cardCommandClient := pbcard.NewCardCommandServiceClient(conn)
	saldoQueryClient := pbsaldo.NewSaldoQueryServiceClient(conn)
	saldoCommandClient := pbsaldo.NewSaldoCommandServiceClient(conn)
	merchantQueryClient := pbmerchant.NewMerchantQueryServiceClient(conn)
	userQueryClient := pbuser.NewUserQueryServiceClient(identity.User.Conn)

	cardRepo := card_repo.NewRepositories(db, userQueryClient)
	cardSvc := card_service.NewService(&card_service.Deps{
		Cache:        cacheStore,
		Repositories: cardRepo,
		Logger:       log,
		Kafka:        nil,
	})
	cardH := card_handler.NewHandler(cardSvc)

	saldoRepo := saldo_repo.NewRepositories(db, cardQueryClient, cardCommandClient)
	saldoSvc := saldo_service.NewService(&saldo_service.Deps{
		Cache:        cacheStore,
		Repositories: saldoRepo,
		Logger:       log,
	})
	saldoH := saldo_handler.NewHandler(saldoSvc)

	merchantRepo := merchant_repo.NewRepositories(db, userQueryClient)
	merchantSvc := merchant_service.NewService(&merchant_service.Deps{
		Cache:        cacheStore,
		Repositories: merchantRepo,
		Logger:       log,
		Kafka:        nil,
	})
	merchantH := merchant_handler.NewHandler(merchantSvc)

	server := grpc.NewServer()
	pbcard.RegisterCardQueryServiceServer(server, cardH)
	pbcard.RegisterCardCommandServiceServer(server, cardH)
	pbsaldo.RegisterSaldoQueryServiceServer(server, saldoH)
	pbsaldo.RegisterSaldoCommandServiceServer(server, saldoH)
	pbmerchant.RegisterMerchantQueryServiceServer(server, merchantH)
	pbmerchant.RegisterMerchantCommandServiceServer(server, merchantH)

	go func() { _ = server.Serve(lis) }()

	return &grpcDeps{
		identity:            identity,
		server:              server,
		conn:                conn,
		cardQueryClient:     cardQueryClient,
		cardCommandClient:   cardCommandClient,
		saldoQueryClient:    saldoQueryClient,
		saldoCommandClient:  saldoCommandClient,
		merchantQueryClient: merchantQueryClient,
		cardRepo:            cardRepo,
		saldoRepo:           saldoRepo,
		merchantRepo:        merchantRepo,
		userRepo:            user_repo.NewRepositories(db, pbrole.NewRoleServiceClient(identity.Role.Conn), pbrole.NewRoleCommandServiceClient(identity.Role.Conn)),
	}, nil
}

func (d *grpcDeps) Close() {
	if d == nil {
		return
	}
	if d.conn != nil {
		_ = d.conn.Close()
	}
	if d.server != nil {
		d.server.Stop()
	}
	d.identity.Close()
}
