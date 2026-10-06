package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-pkg/database"
	"github.com/MamangRust/monolith-payment-gateway-pkg/dotenv"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-pkg/middleware"
	otel_pkg "github.com/MamangRust/monolith-payment-gateway-pkg/otel"
	"github.com/MamangRust/monolith-payment-gateway-pkg/resilience"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"github.com/grafana/pyroscope-go"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	Logger           logger.LoggerInterface
	DB               *sqlx.DB
	Ctx              context.Context
	Cancel           context.CancelFunc
	CacheStore       *cache.CacheStore
	Redis            *redis.Client
	Telemetry        *otel_pkg.Telemetry
	Config           *Config
	RegisterServices func(*grpc.Server)
	cleanupHooks     []func() error
}

// AddCleanupHook registers a resource cleanup function owned by the service
// application. Hooks run before the shared Redis and telemetry cleanup.
func (s *GRPCServer) AddCleanupHook(hook func() error) {
	if hook != nil {
		s.cleanupHooks = append(s.cleanupHooks, hook)
	}
}

func New(cfg *Config) (*GRPCServer, error) {
	if err := dotenv.Viper(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	if err := initPyroscope(cfg); err != nil {
		log.Printf("Warning: Failed to initialize pyroscope: %v", err)
	}

	telemetry := initTelemetry(cfg)
	if err := telemetry.Init(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to initialize telemetry: %w", err)
	}

	cacheMetrics, err := observability.NewCacheMetrics("cache")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cache metrics: %w", err)
	}

	l, err := logger.NewLogger(cfg.ServiceName, telemetry.GetLogger())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	dbConn, err := database.NewClient(l)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	redisClient, err := initRedisServer(ctx, l, cfg.ServiceName)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to initialize Redis: %w", err)
	}

	cacheStore := cache.NewCacheStore(redisClient, l, cacheMetrics)

	return &GRPCServer{
		Logger:     l,
		DB:         dbConn,
		Ctx:        ctx,
		Cancel:     cancel,
		CacheStore: cacheStore,
		Redis:      redisClient,
		Telemetry:  telemetry,
		Config:     cfg,
	}, nil
}

func (s *GRPCServer) Run() error {
	defer s.Cleanup()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Config.Port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", s.Config.Port, err)
	}

	loadMonitor := resilience.NewLoadMonitor()
	circuitBreaker := resilience.NewCircuitBreaker(100, 10, s.Logger)
	requestLimiter := resilience.NewRequestLimiter(800, s.Logger)
	resilienceHandler := middleware.NewResilienceInterceptor(loadMonitor, circuitBreaker, requestLimiter)

	grpcServer := grpc.NewServer(
		grpc.MaxConcurrentStreams(DefaultMaxConcurrentConn),
		grpc.InitialConnWindowSize(DefaultWindowSize),
		grpc.InitialWindowSize(DefaultWindowSize),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    DefaultKeepaliveTime,
			Timeout: DefaultKeepaliveTimeout,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             DefaultMinKeepaliveTime,
			PermitWithoutStream: true,
		}),
		grpc.ChainUnaryInterceptor(
			middleware.ContextMiddleware(30*time.Second, s.Logger),
			middleware.RecoveryMiddleware(s.Logger),
			middleware.PyroscopeUnaryInterceptor(),
			resilienceHandler.UnaryInterceptor(),
		),
	)

	if s.RegisterServices != nil {
		s.RegisterServices(grpcServer)
	}

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	if os.Getenv("ENABLE_REFLECTION") == "true" {
		reflection.Register(grpcServer)
		s.Logger.Info("gRPC reflection enabled")
	}

	monitoringDone := s.spawnMonitoringTask()
	cleanupDone := s.spawnCleanupTask()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	errChan := make(chan error, 1)
	go func() {
		s.Logger.Info("gRPC server starting",
			zap.Int("port", s.Config.Port),
			zap.String("address", lis.Addr().String()),
		)
		if err := grpcServer.Serve(lis); err != nil {
			errChan <- fmt.Errorf("failed to serve: %w", err)
		}
	}()

	select {
	case sig := <-sigChan:
		s.Logger.Info("Received shutdown signal", zap.String("signal", sig.String()))
	case err := <-errChan:
		s.Logger.Error("Server error", zap.Error(err))
		return err
	}

	return s.gracefulShutdown(grpcServer, healthServer, monitoringDone, cleanupDone)
}

func (s *GRPCServer) gracefulShutdown(
	grpcServer *grpc.Server,
	healthServer *health.Server,
	monitoringDone, cleanupDone <-chan struct{},
) error {
	s.Logger.Info("Starting graceful shutdown...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer shutdownCancel()

	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_NOT_SERVING)

	s.Cancel()

	tasksDone := make(chan struct{})
	go func() {
		<-monitoringDone
		<-cleanupDone
		close(tasksDone)
	}()

	select {
	case <-tasksDone:
		s.Logger.Info("Background tasks stopped successfully")
	case <-shutdownCtx.Done():
		s.Logger.Warn("Background tasks shutdown timeout, forcing stop")
	}

	stopped := make(chan struct{})
	go func() {
		grpcServer.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
		s.Logger.Info("gRPC server stopped gracefully")
	case <-shutdownCtx.Done():
		s.Logger.Warn("Graceful shutdown timeout, forcing stop")
		grpcServer.Stop()
	}

	s.Logger.Info("Graceful shutdown completed")
	return nil
}

func (s *GRPCServer) Cleanup() {
	s.Logger.Info("Cleaning up resources...")

	for i := len(s.cleanupHooks) - 1; i >= 0; i-- {
		if err := s.cleanupHooks[i](); err != nil {
			s.Logger.Error("Failed to close application resource", zap.Error(err))
		}
	}

	if s.DB != nil {
		if err := s.DB.Close(); err != nil {
			s.Logger.Error("Failed to close database", zap.Error(err))
		} else {
			s.Logger.Info("Database connection pool closed")
		}
	}

	if s.Redis != nil {
		if err := s.Redis.Close(); err != nil {
			s.Logger.Error("Failed to close Redis connection", zap.Error(err))
		} else {
			s.Logger.Info("Redis connection closed")
		}
	}

	if s.Telemetry != nil {
		if err := s.Telemetry.Shutdown(context.Background()); err != nil {
			s.Logger.Error("Failed to shutdown telemetry", zap.Error(err))
		} else {
			s.Logger.Info("Telemetry shutdown successfully")
		}
	}

	s.Logger.Info("Cleanup completed")
}

func initPyroscope(cfg *Config) error {
	if os.Getenv("PYROSCOPE_ENABLED") != "true" {
		return nil
	}

	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: cfg.ServiceName,
		ServerAddress:   os.Getenv("PYROSCOPE_SERVER"),
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
		},
		Tags: map[string]string{
			"service": cfg.ServiceName,
			"env":     cfg.Environment,
			"version": cfg.ServiceVersion,
		},
	})
	return err
}

func initTelemetry(cfg *Config) *otel_pkg.Telemetry {
	return otel_pkg.NewTelemetry(otel_pkg.Config{
		ServiceName:            cfg.ServiceName,
		ServiceVersion:         cfg.ServiceVersion,
		Environment:            cfg.Environment,
		Endpoint:               envOr("OTEL_ENDPOINT", cfg.OtelEndpoint),
		Insecure:               true,
		EnableRuntimeMetrics:   os.Getenv("OTEL_ENABLED") != "false",
		RuntimeMetricsInterval: 15 * time.Second,
		Disabled:               os.Getenv("OTEL_ENABLED") == "false",
	})
}

// envOr returns the value of the environment variable key if set, otherwise fallback.
// Used to let native (non-container) service runs redirect telemetry to a
// host-mapped OTLP endpoint instead of the in-compose hostname default.
func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func initRedisServer(ctx context.Context, logger logger.LoggerInterface, serviceName string) (*redis.Client, error) {
	// Domain services expose Redis settings with a service suffix (for example,
	// REDIS_HOST_TOPUP). The previous generic-key lookup silently produced
	// ":0", allowing a broken cache client to reach request handling.
	suffix := strings.ToUpper(strings.TrimSuffix(serviceName, "-service"))
	getString := func(suffixed, generic, fallback string) string {
		value := viper.GetString(suffixed)
		if value == "" {
			value = viper.GetString(generic)
		}
		if value == "" {
			value = fallback
		}
		return value
	}

	host := getString("REDIS_HOST_"+suffix, "REDIS_HOST", "")
	port := getString("REDIS_PORT_"+suffix, "REDIS_PORT", "6379")
	if host == "" {
		return nil, fmt.Errorf("Redis host is not configured for %s", serviceName)
	}

	password := getString("REDIS_PASSWORD_"+suffix, "REDIS_PASSWORD", "")
	db := viper.GetInt("REDIS_DB_" + suffix)
	if db == 0 {
		db = viper.GetInt("REDIS_DB")
	}

	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		Password:     password,
		DB:           db,
		DialTimeout:  RedisDialTimeout,
		ReadTimeout:  RedisReadTimeout,
		WriteTimeout: RedisWriteTimeout,
		PoolSize:     RedisPoolSize,
		MinIdleConns: RedisMinIdleConns,
	})

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := client.Ping(pingCtx).Err(); err != nil {
		_ = client.Close()
		return nil, fmt.Errorf("failed to ping Redis for %s at %s:%s: %w", serviceName, host, port, err)
	}

	logger.Info("Redis connection established", zap.String("addr", host+":"+port), zap.Int("db", db))
	return client, nil
}

func (s *GRPCServer) spawnMonitoringTask() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		ticker := time.NewTicker(MonitoringInterval)
		defer ticker.Stop()
		for {
			select {
			case <-s.Ctx.Done():
				return
			case <-ticker.C:
				s.monitorCache()
			}
		}
	}()
	return done
}

func (s *GRPCServer) monitorCache() {
	refCount := s.CacheStore.GetRefCount()
	stats, err := s.CacheStore.GetStats(s.Ctx)
	if err != nil {
		s.Logger.Error("Failed to get cache stats", zap.Error(err))
		return
	}
	logLevel := zap.InfoLevel
	if refCount > CacheRefCountThreshold {
		logLevel = zap.WarnLevel
	}
	if ce := s.Logger.Check(logLevel, "Cache statistics"); ce != nil {
		ce.Write(
			zap.Int64("ref_count", refCount),
			zap.Int64("total_keys", stats.TotalKeys),
			zap.Float64("hit_rate", stats.HitRate),
			zap.String("memory_used", stats.MemoryUsedHuman),
		)
	}
}

func (s *GRPCServer) spawnCleanupTask() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		ticker := time.NewTicker(CleanupInterval)
		defer ticker.Stop()
		for {
			select {
			case <-s.Ctx.Done():
				return
			case <-ticker.C:
				s.CacheStore.ClearExpired(s.Ctx)
			}
		}
	}()
	return done
}
