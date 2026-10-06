package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	_ "github.com/MamangRust/monolith-payment-gateway-apigateway/docs"
	"github.com/MamangRust/monolith-payment-gateway-apigateway/handler"
	"github.com/MamangRust/monolith-payment-gateway-apigateway/middlewares"
	"github.com/MamangRust/monolith-payment-gateway-pkg/auth"
	"github.com/MamangRust/monolith-payment-gateway-pkg/dotenv"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	otel_pkg "github.com/MamangRust/monolith-payment-gateway-pkg/otel"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"github.com/grafana/pyroscope-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
)

const (
	defaultServerPort             = ":5000"
	defaultWindowSizeClient       = 16 * 1024 * 1024
	defaultKeepaliveTimeClient    = 20 * time.Second
	defaultKeepaliveTimeoutClient = 5 * time.Second

	monitoringInterval     = 30 * time.Second
	cleanupInterval        = 5 * time.Minute
	cacheRefCountThreshold = 10000
	shutdownTimeout        = 30 * time.Second

	redisDialTimeout  = 5 * time.Second
	redisReadTimeout  = 3 * time.Second
	redisWriteTimeout = 3 * time.Second
	redisPoolSize     = 10
	redisMinIdleConns = 5
)

// @title PaymentGateway gRPC
// @version 1.0
// @description gRPC based Payment Gateway service
// @host localhost:5000
// @BasePath /api/
// @securityDefinitions.apikey BearerAuth
// @in Header
// @name Authorization

// Client represents the main application client
type Client struct {
	Logger       logger.LoggerInterface
	Echo         *echo.Echo
	GRPCConn     *grpc.ClientConn
	TokenManager *auth.Manager
	Telemetry    *otel_pkg.Telemetry
	Config       *ClientConfig
	Redis        *redis.Client
	Kafka        *kafka.Kafka
	cancelTasks  context.CancelFunc
	tasksDone    []<-chan struct{}
}

type ClientConfig struct {
	ServiceName    string   `mapstructure:"service_name"`
	ServiceVersion string   `mapstructure:"service_version"`
	Environment    string   `mapstructure:"environment"`
	OtelEndpoint   string   `mapstructure:"otel_endpoint"`
	ServerPort     string   `mapstructure:"server_port"`
	AllowedOrigins []string `mapstructure:"allowed_origins"`
}

type CacheManager struct {
	cache  *cache.CacheStore
	logger logger.LoggerInterface
}

type ServiceAddresses struct {
	Auth        string `mapstructure:"auth"`
	Role        string `mapstructure:"role"`
	Card        string `mapstructure:"card"`
	Merchant    string `mapstructure:"merchant"`
	User        string `mapstructure:"user"`
	Saldo       string `mapstructure:"saldo"`
	Topup       string `mapstructure:"topup"`
	Transaction string `mapstructure:"transaction"`
	Transfer    string `mapstructure:"transfer"`
	Withdraw    string `mapstructure:"withdraw"`
}

func NewCacheManager(cache *cache.CacheStore, logger logger.LoggerInterface) *CacheManager {
	return &CacheManager{
		cache:  cache,
		logger: logger,
	}
}

func (cm *CacheManager) StartMonitoring(ctx context.Context) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		ticker := time.NewTicker(monitoringInterval)
		defer ticker.Stop()

		cm.logger.Info("Cache monitoring task started",
			zap.Duration("interval", monitoringInterval),
		)

		for {
			select {
			case <-ctx.Done():
				cm.logger.Info("Cache monitoring task stopped")
				return
			case <-ticker.C:
				cm.monitor(ctx)
			}
		}
	}()

	return done
}

func (cm *CacheManager) monitor(ctx context.Context) {
	refCount := cm.cache.GetRefCount()

	stats, err := cm.cache.GetStats(ctx)
	if err != nil {
		cm.logger.Error("Failed to get cache stats", zap.Error(err))
		return
	}

	logLevel := zap.InfoLevel
	if refCount > cacheRefCountThreshold {
		logLevel = zap.WarnLevel
	}

	if ce := cm.logger.Check(logLevel, "Cache statistics"); ce != nil {
		ce.Write(
			zap.Int64("ref_count", refCount),
			zap.Int64("total_keys", stats.TotalKeys),
			zap.Float64("hit_rate", stats.HitRate),
			zap.String("memory_used", stats.MemoryUsedHuman),
			zap.Bool("high_ref_count", refCount > cacheRefCountThreshold),
		)
	}
}

func (cm *CacheManager) StartCleanup(ctx context.Context) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		ticker := time.NewTicker(cleanupInterval)
		defer ticker.Stop()

		cm.logger.Info("Cache cleanup task started",
			zap.Duration("interval", cleanupInterval),
		)

		for {
			select {
			case <-ctx.Done():
				cm.logger.Info("Cache cleanup task stopped")
				return
			case <-ticker.C:
				cm.cleanup(ctx)
			}
		}
	}()

	return done
}

func (cm *CacheManager) cleanup(ctx context.Context) {
	cm.logger.Info("Starting periodic cache cleanup")

	statsBefore, err := cm.cache.GetStats(ctx)
	if err != nil {
		cm.logger.Error("Failed to get cache stats before cleanup", zap.Error(err))
		statsBefore = nil
	}

	scanned, err := cm.cache.ClearExpired(ctx)
	if err != nil {
		cm.logger.Error("Cache cleanup failed", zap.Error(err))
		return
	}

	statsAfter, err := cm.cache.GetStats(ctx)
	if err != nil {
		cm.logger.Error("Failed to get cache stats after cleanup", zap.Error(err))
		statsAfter = nil
	}

	logFields := []zap.Field{
		zap.Int64("scanned_keys", scanned),
		zap.Int64("ref_count", cm.cache.GetRefCount()),
	}

	if statsBefore != nil && statsAfter != nil {
		keysRemoved := statsBefore.TotalKeys - statsAfter.TotalKeys
		logFields = append(logFields,
			zap.Int64("keys_before", statsBefore.TotalKeys),
			zap.Int64("keys_after", statsAfter.TotalKeys),
			zap.Int64("keys_removed", keysRemoved),
			zap.String("memory_before", statsBefore.MemoryUsedHuman),
			zap.String("memory_after", statsAfter.MemoryUsedHuman),
		)
	}

	cm.logger.Info("Cache cleanup completed", logFields...)
}

func NewClient(cfg *ClientConfig) (*Client, error) {
	if err := dotenv.Viper(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	if err := initPyroscope(); err != nil {
		log.Fatal("Failed to initialize pyroscope:", err)
	}

	cfg, err := loadClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	telemetry, err := initTelemetry(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize telemetry: %w", err)
	}

	cacheMetrics, err := observability.NewCacheMetrics("cache")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cache metrics: %w", err)
	}

	logger, err := logger.NewLogger(cfg.ServiceName, telemetry.GetLogger())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	tokenManager, err := auth.NewManager(viper.GetString("SECRET_KEY"))
	if err != nil {
		return nil, fmt.Errorf("failed to create token manager: %w", err)
	}

	addresses, err := loadServiceAddresses()

	if err != nil {
		return nil, fmt.Errorf("failed to load service")
	}

	conns, err := createServiceConnections(addresses, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to connect services: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	redisClient, err := initRedisClient(ctx, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Redis: %w", err)
	}

	echoServer := createEchoServer(cfg, redisClient, conns)

	myKafka := kafka.NewKafka(logger, []string{viper.GetString("KAFKA_BROKERS")})

	cacheStore := cache.NewCacheStore(redisClient, logger, cacheMetrics)

	tasksCtx, cancelTasks := context.WithCancel(context.Background())
	cacheManager := NewCacheManager(cacheStore, logger)

	tasksDone := []<-chan struct{}{
		cacheManager.StartMonitoring(tasksCtx),
		cacheManager.StartCleanup(tasksCtx),
	}

	handlerDeps := &handler.Deps{
		Kafka:              myKafka,
		ServiceConnections: conns,
		Token:              tokenManager,
		E:                  echoServer,
		Logger:             logger,
		Cache:              cacheStore,
	}
	handler.NewHandler(handlerDeps)

	client := &Client{
		Logger:       logger,
		Echo:         echoServer,
		TokenManager: tokenManager,
		Telemetry:    telemetry,
		Config:       cfg,
		Redis:        redisClient,
		Kafka:        myKafka,
		cancelTasks:  cancelTasks,
		tasksDone:    tasksDone,
	}

	logger.Info("Client initialized successfully",
		zap.String("service", cfg.ServiceName),
		zap.String("version", cfg.ServiceVersion),
		zap.String("server_port", cfg.ServerPort),
	)

	return client, nil
}

func (c *Client) Run() error {
	defer c.Cleanup()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	errChan := make(chan error, 1)
	go func() {
		c.Logger.Info("HTTP server starting",
			zap.String("port", c.Config.ServerPort),
			zap.String("swagger", "http://localhost"+c.Config.ServerPort+"/swagger/index.html"),
		)
		if err := c.Echo.Start(c.Config.ServerPort); err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("failed to start server: %w", err)
		}
	}()

	select {
	case sig := <-quit:
		c.Logger.Info("Received shutdown signal", zap.String("signal", sig.String()))
	case err := <-errChan:
		c.Logger.Error("Server error", zap.Error(err))
		return err
	}

	return c.gracefulShutdown()
}

func (c *Client) gracefulShutdown() error {
	c.Logger.Info("Starting graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := c.Echo.Shutdown(ctx); err != nil {
		c.Logger.Error("Echo shutdown error", zap.Error(err))
		return fmt.Errorf("failed to shutdown echo server: %w", err)
	}

	c.Logger.Info("HTTP server stopped gracefully")
	return nil
}

func (c *Client) Cleanup() {
	c.Logger.Info("Cleaning up resources...")

	if c.cancelTasks != nil {
		c.Logger.Info("Stopping background tasks...")
		c.cancelTasks()

		for i, done := range c.tasksDone {
			c.Logger.Debug("Waiting for background task to complete", zap.Int("task_index", i))
			<-done
		}
		c.Logger.Info("All background tasks stopped")
	}

	if c.Kafka != nil {
		if err := c.Kafka.Close(); err != nil {
			c.Logger.Error("Failed to close Kafka resources", zap.Error(err))
		} else {
			c.Logger.Info("Kafka resources closed")
		}
	}

	if c.Redis != nil {
		if err := c.Redis.Close(); err != nil {
			c.Logger.Error("Failed to close Redis connection", zap.Error(err))
		} else {
			c.Logger.Info("Redis connection closed")
		}
	}

	if c.GRPCConn != nil {
		if err := c.GRPCConn.Close(); err != nil {
			c.Logger.Error("Failed to close gRPC connection", zap.Error(err))
		} else {
			c.Logger.Info("gRPC connection closed")
		}
	}

	if c.Telemetry != nil {
		if err := c.Telemetry.Shutdown(context.Background()); err != nil {
			c.Logger.Error("Failed to shutdown telemetry", zap.Error(err))
		} else {
			c.Logger.Info("Telemetry shutdown successfully")
		}
	}

	if c.Logger != nil {
		_ = c.Logger.Sync()
	}

	c.Logger.Info("Cleanup completed")
}

func initPyroscope() error {
	if os.Getenv("PYROSCOPE_ENABLED") != "true" {
		return nil
	}

	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "apigateway",
		ServerAddress:   os.Getenv("PYROSCOPE_SERVER"),
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
		},
		Tags: map[string]string{
			"service": "apigateway",
			"env":     os.Getenv("ENV"),
			"version": os.Getenv("VERSION"),
		},
	})
	return err
}

func initTelemetry(cfg *ClientConfig) (*otel_pkg.Telemetry, error) {
	telemetry := otel_pkg.NewTelemetry(otel_pkg.Config{
		ServiceName:            cfg.ServiceName,
		ServiceVersion:         cfg.ServiceVersion,
		Environment:            cfg.Environment,
		Endpoint:               cfg.OtelEndpoint,
		Insecure:               true,
		EnableRuntimeMetrics:   os.Getenv("OTEL_ENABLED") != "false",
		RuntimeMetricsInterval: 15 * time.Second,
		Disabled:               os.Getenv("OTEL_ENABLED") == "false",
	})

	if err := telemetry.Init(context.Background()); err != nil {
		return nil, err
	}

	return telemetry, nil
}

func initRedisClient(ctx context.Context, logger logger.LoggerInterface) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", viper.GetString("REDIS_HOST_APIGATEWAY"), viper.GetString("REDIS_PORT_APIGATEWAY")),
		Password:     viper.GetString("REDIS_PASSWORD_APIGATEWAY"),
		DB:           viper.GetInt("REDIS_DB_APIGATEWAY"),
		DialTimeout:  redisDialTimeout,
		ReadTimeout:  redisReadTimeout,
		WriteTimeout: redisWriteTimeout,
		PoolSize:     redisPoolSize,
		MinIdleConns: redisMinIdleConns,
	})

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := client.Ping(pingCtx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping Redis: %w", err)
	}

	logger.Info("Redis connection established",
		zap.String("addr", fmt.Sprintf("%s:%s", viper.GetString("REDIS_HOST_APIGATEWAY"), viper.GetString("REDIS_PORT_APIGATEWAY"))),
		zap.Int("db", viper.GetInt("REDIS_DB_APIGATEWAY")),
	)

	return client, nil
}

func createEchoServer(cfg *ClientConfig, redisClient *redis.Client, conns *handler.ServiceConnections) *echo.Echo {
	e := echo.New()

	// Echo's default not-found handler can leave an unmatched request with a
	// successful empty response when no route writes a body. Return an explicit
	// JSON 404 so API tests and clients cannot mistake an unknown route for a
	// successful endpoint.
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		code := http.StatusInternalServerError
		message := "Internal server error"
		if httpErr, ok := err.(*echo.HTTPError); ok {
			code = httpErr.Code
			if text, ok := httpErr.Message.(string); ok && text != "" {
				message = text
			}
		}

		_ = c.JSON(code, map[string]interface{}{
			"status":  "error",
			"message": message,
			"code":    code,
		})
	}

	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	// Echo can leave unmatched requests with a nil handler on this runtime.
	// Convert that condition into an explicit 404 before it reaches the
	// response writer, so an unknown endpoint can never look like HTTP 200.
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil && c.Path() == "" && !c.Response().Committed {
				return echo.NewHTTPError(http.StatusNotFound)
			}
			return err
		}
	})
	e.Use(createLoggerMiddleware())
	e.Use(middlewares.PyroscopeMiddleware())
	e.Use(createCORSMiddleware(cfg.AllowedOrigins))
	e.Use(middleware.Gzip())
	e.Use(createSecureMiddleware())

	middlewares.WebSecurityConfig(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/health", createLivenessHandler(cfg))
	e.GET("/live", createLivenessHandler(cfg))
	e.GET("/ready", createReadinessHandler(cfg, redisClient, conns))

	return e
}

func createLivenessHandler(cfg *ClientConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "healthy",
			"service": cfg.ServiceName,
			"version": cfg.ServiceVersion,
			"time":    time.Now().UTC(),
		})
	}
}

func createReadinessHandler(cfg *ClientConfig, redisClient *redis.Client, conns *handler.ServiceConnections) echo.HandlerFunc {
	return func(c echo.Context) error {
		checks := map[string]string{}
		ready := true

		redisCtx, cancel := context.WithTimeout(c.Request().Context(), time.Second)
		defer cancel()
		if redisClient == nil || redisClient.Ping(redisCtx).Err() != nil {
			checks["redis"] = "unavailable"
			ready = false
		} else {
			checks["redis"] = "ready"
		}

		if conns == nil {
			ready = false
			checks["grpc"] = "unavailable"
		} else {
			connections := map[string]*grpc.ClientConn{
				"auth": conns.Auth, "role": conns.Role, "card": conns.Card,
				"merchant": conns.Merchant, "user": conns.User, "saldo": conns.Saldo,
				"topup": conns.Topup, "transaction": conns.Transaction,
				"transfer": conns.Transfer, "withdraw": conns.Withdraw,
			}
			for name, conn := range connections {
				if conn == nil {
					checks[name] = "unavailable"
					ready = false
					continue
				}

				state := conn.GetState()
				if state == connectivity.TransientFailure || state == connectivity.Shutdown {
					checks[name] = "unavailable"
					ready = false
					continue
				}

				checkCtx, checkCancel := context.WithTimeout(c.Request().Context(), time.Second)
				response, err := grpc_health_v1.NewHealthClient(conn).Check(checkCtx, &grpc_health_v1.HealthCheckRequest{})
				checkCancel()
				if err != nil || response.GetStatus() != grpc_health_v1.HealthCheckResponse_SERVING {
					checks[name] = "unavailable"
					ready = false
					continue
				}
				checks[name] = state.String()
			}
		}

		statusCode := http.StatusOK
		status := "ready"
		if !ready {
			statusCode = http.StatusServiceUnavailable
			status = "not_ready"
		}
		return c.JSON(statusCode, map[string]interface{}{
			"status":  status,
			"service": cfg.ServiceName,
			"version": cfg.ServiceVersion,
			"checks":  checks,
			"time":    time.Now().UTC(),
		})
	}
}

func createLoggerMiddleware() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","status":${status},` +
			`"error":"${error}","latency":${latency},"latency_human":"${latency_human}",` +
			`"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	})
}

func createCORSMiddleware(allowedOrigins []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"X-Request-ID",
		},
		AllowCredentials: true,
		MaxAge:           86400,
	})
}

func createSecureMiddleware() echo.MiddlewareFunc {
	return middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            31536000,
		HSTSExcludeSubdomains: false,
		HSTSPreloadEnabled:    true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		ContentSecurityPolicy: "default-src 'self'; " +
			"script-src 'self' 'unsafe-inline' https://cdnjs.cloudflare.com; " +
			"style-src 'self' 'unsafe-inline' https://cdnjs.cloudflare.com; " +
			"img-src 'self' data: https:; " +
			"font-src 'self' data: https://cdnjs.cloudflare.com; " +
			"connect-src 'self'; " +
			"frame-ancestors 'none';",
	})
}

func loadClientConfig() (*ClientConfig, error) {
	v := viper.GetViper()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.SetDefault("service_name", "apigateway")
	v.SetDefault("service_version", "1.0.0")
	v.SetDefault("environment", "production")
	v.SetDefault("otel_endpoint", "otel-collector:4317")
	v.SetDefault("server_port", defaultServerPort)
	v.SetDefault("allowed_origins", []string{"http://localhost:1420"})

	var cfg ClientConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal client config: %w", err)
	}

	return &cfg, nil
}

func loadServiceAddresses() (*ServiceAddresses, error) {
	v := viper.GetViper()

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// The local and Kubernetes configuration files expose addresses as
	// GRPC_<SERVICE>_ADDR. Reading the short keys (for example, "card")
	// silently falls back to defaults and can route the gateway to a service's
	// wrong port, so resolve the address keys explicitly.
	defaults := map[string]string{
		"auth":        "auth:50051",
		"role":        "role:50052",
		"card":        "card:50054",
		"merchant":    "merchant:50055",
		"user":        "user:50053",
		"saldo":       "saldo:50056",
		"topup":       "topup:50058",
		"transaction": "transaction:50057",
		"transfer":    "transfer:50059",
		"withdraw":    "withdraw:50060",
	}

	addresses := make(map[string]string, len(defaults))
	for service, fallback := range defaults {
		key := "GRPC_" + strings.ToUpper(service) + "_ADDR"
		address := v.GetString(key)
		if address == "" {
			address = fallback
		}
		addresses[service] = address
	}

	return &ServiceAddresses{
		Auth:        addresses["auth"],
		Role:        addresses["role"],
		Card:        addresses["card"],
		Merchant:    addresses["merchant"],
		User:        addresses["user"],
		Saldo:       addresses["saldo"],
		Topup:       addresses["topup"],
		Transaction: addresses["transaction"],
		Transfer:    addresses["transfer"],
		Withdraw:    addresses["withdraw"],
	}, nil
}

func createServiceConnections(addresses *ServiceAddresses, logger logger.LoggerInterface) (*handler.ServiceConnections, error) {
	connections := &handler.ServiceConnections{}

	serviceMap := map[string]struct {
		addr *string
		conn **grpc.ClientConn
	}{
		"Auth":        {&addresses.Auth, &connections.Auth},
		"Role":        {&addresses.Role, &connections.Role},
		"Card":        {&addresses.Card, &connections.Card},
		"Merchant":    {&addresses.Merchant, &connections.Merchant},
		"User":        {&addresses.User, &connections.User},
		"Saldo":       {&addresses.Saldo, &connections.Saldo},
		"Topup":       {&addresses.Topup, &connections.Topup},
		"Transaction": {&addresses.Transaction, &connections.Transaction},
		"Transfer":    {&addresses.Transfer, &connections.Transfer},
		"Withdraw":    {&addresses.Withdraw, &connections.Withdraw},
	}

	for name, svc := range serviceMap {
		conn, err := createConnection(*svc.addr, name, logger)
		if err != nil {
			closeConnections(connections, logger)
			return nil, err
		}
		*svc.conn = conn
	}

	return connections, nil
}

func createConnection(address, serviceName string, logger logger.LoggerInterface) (*grpc.ClientConn, error) {
	logger.Info(fmt.Sprintf("Connecting to %s service", serviceName),
		zap.String("address", address),
	)

	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithInitialConnWindowSize(defaultWindowSizeClient),
		grpc.WithInitialWindowSize(defaultWindowSizeClient),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                defaultKeepaliveTimeClient,
			Timeout:             defaultKeepaliveTimeoutClient,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to %s service", serviceName), zap.Error(err))
		return nil, fmt.Errorf("failed to connect to %s service: %w", serviceName, err)
	}

	logger.Info(fmt.Sprintf("Successfully connected to %s service", serviceName))
	return conn, nil
}

func closeConnections(conns *handler.ServiceConnections, logger logger.LoggerInterface) {
	connectionMap := map[string]*grpc.ClientConn{
		"Auth":        conns.Auth,
		"Role":        conns.Role,
		"Card":        conns.Card,
		"Merchant":    conns.Merchant,
		"User":        conns.User,
		"Saldo":       conns.Saldo,
		"Topup":       conns.Topup,
		"Transaction": conns.Transaction,
		"Transfer":    conns.Transfer,
		"Withdraw":    conns.Withdraw,
	}

	for name, conn := range connectionMap {
		if conn != nil {
			if err := conn.Close(); err != nil {
				logger.Error(fmt.Sprintf("Failed to close %s connection", name), zap.Error(err))
			} else {
				logger.Info(fmt.Sprintf("%s connection closed", name))
			}
		}
	}
}
