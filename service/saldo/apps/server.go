package apps

import (
	"context"
	"fmt"
	"time"

	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	pbstats "github.com/MamangRust/monolith-payment-gateway-pb/saldo/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/resilience"
	"github.com/MamangRust/monolith-payment-gateway-pkg/server"
	"github.com/MamangRust/monolith-payment-gateway-saldo/handler"
	saldokafka "github.com/MamangRust/monolith-payment-gateway-saldo/kafka"
	"github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	"github.com/MamangRust/monolith-payment-gateway-saldo/service"
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

	// gRPC client to the card service so saldo reads card data through the
	// shared adapter instead of the card service's tables.
	connCard, err := grpc.NewClient(
		viper.GetString("GRPC_CARD_ADDR"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		resilience.GuardDialOption("card", srv.Logger),
	)
	if err != nil {
		return nil, fmt.Errorf("connect to card service: %w", err)
	}
	srv.AddCleanupHook(func() error {
		return connCard.Close()
	})

	cardQueryClient := pbcard.NewCardQueryServiceClient(connCard)
	cardCommandClient := pbcard.NewCardCommandServiceClient(connCard)

	repos := repository.NewRepositories(
		srv.DB,
		cardQueryClient,
		cardCommandClient,
		repository.GuardOptions{
			Card: []adapter.GuardOption{
				adapter.WithDependencyGuard(resilience.NewDependencyGuard("card", 5, 30, 100, 3*time.Second, srv.Logger)),
			},
		},
	)

	mykafka := kafka.NewKafka(srv.Logger, []string{viper.GetString("KAFKA_BROKERS")})
	srv.AddCleanupHook(mykafka.Close)

	svc := service.NewService(&service.Deps{
		Cache:        srv.CacheStore,
		Logger:       srv.Logger,
		Repositories: repos,
	})

	kafkaHandler := saldokafka.NewSaldoKafkaHandler(svc, srv.Logger, context.Background())
	_, err = mykafka.StartConsumersWithContext(srv.Ctx, []string{
		"saldo-service-topic-create-saldo",
		"saldo-service-topic-invalidate-cache",
	}, "saldo-service-group", kafkaHandler)
	if err != nil {
		srv.Logger.Error("Failed to start kafka consumers", zap.Error(err))
	}

	h := handler.NewHandler(svc)

	srv.RegisterServices = func(gs *grpc.Server) {
		pb.RegisterSaldoQueryServiceServer(gs, h)
		pb.RegisterSaldoCommandServiceServer(gs, h)
		pbstats.RegisterSaldoStatsBalanceServiceServer(gs, h)
		pbstats.RegisterSaldoStatsTotalBalanceServer(gs, h)
	}

	return srv, nil
}
