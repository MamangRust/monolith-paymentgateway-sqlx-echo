package apps

import (
	"context"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/role"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/server"
	"github.com/MamangRust/monolith-payment-gateway-role/handler"
	myhandlerkafka "github.com/MamangRust/monolith-payment-gateway-role/kafka"
	"github.com/MamangRust/monolith-payment-gateway-role/repository"
	"github.com/MamangRust/monolith-payment-gateway-role/service"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewServer(cfg *server.Config) (*server.GRPCServer, error) {
	srv, err := server.New(cfg)
	if err != nil {
		return nil, err
	}

	repos := repository.NewRepositories(srv.DB)
	mykafka := kafka.NewKafka(srv.Logger, []string{viper.GetString("KAFKA_BROKERS")})
	srv.AddCleanupHook(mykafka.Close)

	svc := service.NewService(&service.Deps{
		Cache:        srv.CacheStore,
		Logger:       srv.Logger,
		Repositories: repos,
	})

	kafkaHandler := myhandlerkafka.NewRoleKafkaHandler(svc.RoleQuery, mykafka, srv.Logger, context.Background())
	_, err = mykafka.StartConsumersWithContext(srv.Ctx, []string{"request-role"}, "role-service-group", kafkaHandler)
	if err != nil {
		srv.Logger.Error("Failed to start kafka consumers", zap.Error(err))
	}

	h := handler.NewHandler(svc)

	srv.RegisterServices = func(gs *grpc.Server) {
		pb.RegisterRoleServiceServer(gs, h.RoleQuery)
		pb.RegisterRoleCommandServiceServer(gs, h.RoleCommand)
	}

	return srv, nil
}
