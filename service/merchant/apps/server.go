package apps

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-merchant/handler"
	"github.com/MamangRust/monolith-payment-gateway-merchant/repository"
	"github.com/MamangRust/monolith-payment-gateway-merchant/service"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	pbstats "github.com/MamangRust/monolith-payment-gateway-pb/merchant/stats"
	pbdocument "github.com/MamangRust/monolith-payment-gateway-pb/merchant_document"
	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/outbox"
	"github.com/MamangRust/monolith-payment-gateway-pkg/resilience"
	"github.com/MamangRust/monolith-payment-gateway-pkg/server"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServer(cfg *server.Config) (*server.GRPCServer, error) {
	srv, err := server.New(cfg)
	if err != nil {
		return nil, err
	}

	myKafka := kafka.NewKafka(srv.Logger, []string{viper.GetString("KAFKA_BROKERS")})
	srv.AddCleanupHook(myKafka.Close)
	relay, relayErr := outbox.NewRelay(srv.DB, myKafka, outbox.RelayConfig{})
	if relayErr != nil {
		srv.Cleanup()
		return nil, fmt.Errorf("initialize merchant outbox relay: %w", relayErr)
	}
	{
		relayDone := make(chan struct{})
		go func() {
			defer close(relayDone)
			if relayErr := relay.Run(srv.Ctx); relayErr != nil && !errors.Is(relayErr, context.Canceled) {
				srv.Logger.Error("merchant outbox relay stopped", zap.Error(relayErr))
			}
		}()
		srv.AddCleanupHook(func() error {
			select {
			case <-relayDone:
				return nil
			case <-time.After(5 * time.Second):
				return errors.New("timed out waiting for merchant outbox relay")
			}
		})
	}

	// gRPC client to the user service so merchant ownership is validated
	// against the canonical user store via the shared user adapter.
	connUser, err := grpc.NewClient(
		viper.GetString("GRPC_USER_ADDR"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		resilience.GuardDialOption("user", srv.Logger),
	)
	if err != nil {
		return nil, fmt.Errorf("connect to user service: %w", err)
	}
	srv.AddCleanupHook(func() error {
		return connUser.Close()
	})

	repos := repository.NewRepositories(srv.DB, pbuser.NewUserQueryServiceClient(connUser),
		repository.GuardOptions{
			User: []adapter.GuardOption{
				adapter.WithDependencyGuard(resilience.NewDependencyGuard("user", 5, 30, 100, 3*time.Second, srv.Logger)),
			},
		},
	)
	svc := service.NewService(&service.Deps{
		Cache:        srv.CacheStore,
		Logger:       srv.Logger,
		Repositories: repos,
	})
	h := handler.NewHandler(svc)

	srv.RegisterServices = func(gs *grpc.Server) {
		pb.RegisterMerchantQueryServiceServer(gs, h)
		pb.RegisterMerchantCommandServiceServer(gs, h)
		pb.RegisterMerchantTransactionServiceServer(gs, h)
		pbdocument.RegisterMerchantDocumentQueryServiceServer(gs, h)
		pbdocument.RegisterMerchantDocumentCommandServiceServer(gs, h)
		pbstats.RegisterMerchantStatsAmountServiceServer(gs, h)
		pbstats.RegisterMerchantStatsMethodServiceServer(gs, h)
		pbstats.RegisterMerchantStatsTotalAmountServiceServer(gs, h)
	}

	return srv, nil
}
