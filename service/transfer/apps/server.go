package apps

import (
	"context"
	"errors"
	"fmt"
	"time"

	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pbsaldo "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transfer"
	pbstats "github.com/MamangRust/monolith-payment-gateway-pb/transfer/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/outbox"
	"github.com/MamangRust/monolith-payment-gateway-pkg/resilience"
	"github.com/MamangRust/monolith-payment-gateway-pkg/server"
	"github.com/MamangRust/monolith-payment-gateway-transfer/handler"
	transferstatshandler "github.com/MamangRust/monolith-payment-gateway-transfer/handler/stats"
	"github.com/MamangRust/monolith-payment-gateway-transfer/repository"
	"github.com/MamangRust/monolith-payment-gateway-transfer/service"
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

	// gRPC clients for cross-service communication. Fail during bootstrap rather
	// than constructing adapters around a nil connection and panicking on the
	// first request that needs saldo/card data.
	connSaldo, err := grpc.NewClient(
		viper.GetString("GRPC_SALDO_ADDR"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("connect to saldo service: %w", err)
	}

	connCard, err := grpc.NewClient(
		viper.GetString("GRPC_CARD_ADDR"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		_ = connSaldo.Close()
		return nil, fmt.Errorf("connect to card service: %w", err)
	}

	srv.AddCleanupHook(func() error {
		return connSaldo.Close()
	})
	srv.AddCleanupHook(func() error {
		return connCard.Close()
	})

	saldoClientQuery := pbsaldo.NewSaldoQueryServiceClient(connSaldo)
	saldoClientCmd := pbsaldo.NewSaldoCommandServiceClient(connSaldo)
	cardClientQuery := pbcard.NewCardQueryServiceClient(connCard)
	cardClientCmd := pbcard.NewCardCommandServiceClient(connCard)

	repos := repository.NewRepositories(
		srv.DB,
		cardClientQuery,
		cardClientCmd,
		saldoClientQuery,
		saldoClientCmd,
		repository.GuardOptions{
			Card: []adapter.GuardOption{
				adapter.WithDependencyGuard(resilience.NewDependencyGuard("card", 5, 30, 100, 3*time.Second, srv.Logger)),
			},
			Saldo: []adapter.GuardOption{
				adapter.WithDependencyGuard(resilience.NewDependencyGuard("saldo", 5, 30, 100, 3*time.Second, srv.Logger)),
			},
		},
	)

	myKafka := kafka.NewKafka(srv.Logger, []string{viper.GetString("KAFKA_BROKERS")})
	srv.AddCleanupHook(myKafka.Close)
	relay, relayErr := outbox.NewRelay(srv.DB, myKafka, outbox.RelayConfig{})
	if relayErr != nil {
		srv.Cleanup()
		return nil, fmt.Errorf("initialize transfer outbox relay: %w", relayErr)
	}
	{
		relayDone := make(chan struct{})
		go func() {
			defer close(relayDone)
			if err := relay.Run(srv.Ctx); err != nil && !errors.Is(err, context.Canceled) {
				srv.Logger.Error("transfer outbox relay stopped", zap.Error(err))
			}
		}()
		srv.AddCleanupHook(func() error {
			select {
			case <-relayDone:
				return nil
			case <-time.After(5 * time.Second):
				return errors.New("timed out waiting for transfer outbox relay")
			}
		})
	}
	svc := service.NewService(&service.Deps{
		Kafka:        myKafka,
		Cache:        srv.CacheStore,
		Logger:       srv.Logger,
		Repositories: repos,
	})
	h := handler.NewHandler(svc)
	statsH := transferstatshandler.NewTransferStatsHandleGrpc(svc)

	srv.RegisterServices = func(gs *grpc.Server) {
		pb.RegisterTransferQueryServiceServer(gs, h)
		pb.RegisterTransferCommandServiceServer(gs, h)
		pbstats.RegisterTransferStatsAmountServiceServer(gs, statsH)
		pbstats.RegisterTransferStatsStatusServiceServer(gs, statsH)
	}

	return srv, nil
}
