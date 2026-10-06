package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-email/config"
	"github.com/MamangRust/monolith-payment-gateway-email/handler"
	"github.com/MamangRust/monolith-payment-gateway-email/mailer"
	"github.com/MamangRust/monolith-payment-gateway-email/metrics"
	"github.com/MamangRust/monolith-payment-gateway-pkg/database"
	"github.com/MamangRust/monolith-payment-gateway-pkg/dotenv"
	"github.com/MamangRust/monolith-payment-gateway-pkg/emailretry"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	otel_pkg "github.com/MamangRust/monolith-payment-gateway-pkg/otel"
	"github.com/MamangRust/monolith-payment-gateway-pkg/outbox"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	telemetry := otel_pkg.NewTelemetry(otel_pkg.Config{
		ServiceName:            "email-service",
		ServiceVersion:         "v1.0.0",
		Environment:            "production",
		Endpoint:               envOr("OTEL_ENDPOINT", "otel-collector:4317"),
		Insecure:               true,
		EnableRuntimeMetrics:   os.Getenv("OTEL_ENABLED") != "false",
		RuntimeMetricsInterval: 15 * time.Second,
		Disabled:               os.Getenv("OTEL_ENABLED") == "false",
	})

	if err := telemetry.Init(context.Background()); err != nil {
		return
	}

	// Register OTel metric instruments after the SDK is initialized so they
	// are bound to the real meter provider (exported via OTLP), not the noop
	// default from package init.
	metrics.Register()

	logger, err := logger.NewLogger("email-service", telemetry.GetLogger())
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}

	if err := dotenv.Viper(); err != nil {
		logger.Fatal("Failed to load .env file", zap.Error(err))
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	cfg := config.Config{
		KafkaBrokers: []string{viper.GetString("KAFKA_BROKERS")},
		SMTPServer:   viper.GetString("SMTP_SERVER"),
		SMTPPort:     viper.GetInt("SMTP_PORT"),
		SMTPUser:     viper.GetString("SMTP_USER"),
		SMTPPass:     viper.GetString("SMTP_PASS"),
		MaxRetries:   viper.GetInt("EMAIL_MAX_RETRIES"),
		RetryBackoff: viper.GetDuration("EMAIL_RETRY_BACKOFF"),
	}
	if cfg.MaxRetries <= 0 {
		cfg.MaxRetries = emailretry.DefaultMaxAttempts
	}
	if cfg.RetryBackoff <= 0 {
		cfg.RetryBackoff = emailretry.DefaultBackoff
	}

	defer func() {
		if err := telemetry.Shutdown(ctx); err != nil {
			logger.Fatal("Failed to shutdown tracer provider", zap.Error(err))
		}
	}()

	m := mailer.NewMailer(
		ctx,
		cfg.SMTPServer,
		cfg.SMTPPort,
		cfg.SMTPUser,
		cfg.SMTPPass,
		logger,
	)

	dbPool, err := database.NewClient(logger)
	if err != nil {
		logger.Fatal("Failed to connect to database for consumer inbox", zap.Error(err))
	}
	defer dbPool.Close()

	inbox, err := outbox.NewPostgresInbox(dbPool)
	if err != nil {
		logger.Fatal("Failed to initialize consumer inbox", zap.Error(err))
	}
	myKafka := kafka.NewKafka(logger, cfg.KafkaBrokers)

	h := handler.NewEmailHandlerWithInbox(ctx, logger, m, inbox, "email-service-group", myKafka, cfg.RetryBackoff)

	srvConsumer, err := myKafka.StartConsumersWithContextManualCommit(ctx, []string{
		"email-service-topic-auth-register",
		"email-service-topic-auth-forgot-password",
		"email-service-topic-auth-verify-code-success",
		"email-service-topic-saldo-create",
		"email-service-topic-topup-create",
		"email-service-topic-transaction-create",
		"email-service-topic-transfer-create",
		"email-service-topic-withdraw-create",
		"email-service-topic-withdraw-update",
		"email-service-topic-merchant-create",
		"email-service-topic-merchant-update-status",
		"email-service-topic-merchant-document-create",
		"email-service-topic-merchant-document-update-status",
	}, "email-service-group", h)

	if err != nil {
		log.Fatalf("Error starting consumer: %v", err)
	}

	retryH := handler.NewRetryHandler(ctx, logger, m, inbox, "email-service-group", myKafka, cfg.MaxRetries, cfg.RetryBackoff)
	srvRetry, err := myKafka.StartConsumersWithContextManualCommit(ctx, []string{emailretry.RetryTopic}, emailretry.RetryGroup, retryH)
	if err != nil {
		log.Fatalf("Error starting retry consumer: %v", err)
	}

	logger.Info("Email service started", zap.String("service", "email-service"), zap.String("retry_topic", emailretry.RetryTopic), zap.String("dlq_topic", emailretry.DLQTopic))
	<-ctx.Done()
	if err := srvRetry.Close(); err != nil {
		logger.Error("Failed to close email retry consumer", zap.Error(err))
	}
	if err := srvConsumer.Close(); err != nil {
		logger.Error("Failed to close email consumer", zap.Error(err))
	}
	if err := myKafka.Close(); err != nil {
		logger.Error("Failed to close Kafka resources", zap.Error(err))
	}

}

// envOr returns the value of the environment variable key if set, otherwise fallback.
func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
