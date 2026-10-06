package myhandlerkafka

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/IBM/sarama"
	"github.com/MamangRust/monolith-payment-gateway-merchant/service"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/repository"
	"go.uber.org/zap"
)

// merchantKafkaHandler is a struct that implements the sarama.ConsumerGroupHandler interface
type merchantKafkaHandler struct {
	logger          logger.LoggerInterface
	merchantService service.MerchantQueryService
	kafka           *kafka.Kafka
}

// NewMerchantKafkaHandler creates a new Kafka consumer group handler for processing merchant API key validation responses.
//
// It takes a merchant query service, a Kafka producer, and a logger as parameters.
// The handler is used to process incoming Kafka messages from the merchant API key validation response topic.
// It implements the sarama.ConsumerGroupHandler interface to handle consumer group lifecycle events.
func NewMerchantKafkaHandler(merchantService service.MerchantQueryService, kafka *kafka.Kafka, logger logger.LoggerInterface) sarama.ConsumerGroupHandler {
	return &merchantKafkaHandler{
		merchantService: merchantService,
		kafka:           kafka,
		logger:          logger,
	}
}

// Setup is called at the beginning of a new Kafka consumer group session.
//
// It can be used to initialize resources or state before message consumption begins.
// In this implementation, it performs no setup and returns nil.
func (m *merchantKafkaHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is called at the end of a Kafka consumer group session.
//
// It can be used to release resources allocated during Setup or message consumption.
// In this implementation, it performs no cleanup and returns nil.
func (m *merchantKafkaHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim processes incoming Kafka messages from the specified consumer group claim.
// It unmarshals each message into a payload map and retrieves the correlation ID.
// If a valid correlation ID is found, it checks if the API key is valid by calling
// the FindByApiKey method of the merchantService. If the API key is valid, it sends
// a valid response to the corresponding Kafka topic. Each message is marked as processed
// in the consumer group session.
func (m *merchantKafkaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for msg := range claim.Messages() {
		var payload requests.MerchantRequestPayload
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			m.logger.Error("Failed to unmarshal merchant request", zap.Error(err))
			// Malformed input is non-retryable and safe to acknowledge.
			session.MarkMessage(msg, "")
			continue
		}

		resp := response.MerchantResponsePayload{
			CorrelationID: payload.CorrelationID,
			Valid:         false,
		}

		merchant, err := m.merchantService.FindByApiKey(ctx, payload.ApiKey)
		switch {
		case err == nil && merchant != nil:
			resp.Valid = true
			resp.MerchantID = int64(merchant.MerchantID)
		case err == nil:
			// Service mengembalikan (nil, nil) tanpa error → perlakukan sebagai
			// not-found deterministik (Valid:false), bukan kegagalan infra.
		case errors.Is(err, merchant_errors.ErrMerchantNotFound):
			// API key tidak dikenal (bukan kegagalan sistem) → balas Valid:false;
			// aman di-commit karena keputusan deterministik.
		default:
			// Kegagalan infrastruktur (DB error, timeout, dsb) harus menghentikan
			// claim. Jika hanya continue, pesan berikutnya yang di-mark dapat
			// membuat offset pesan gagal ikut ter-commit.
			m.logger.Error("Merchant API key validation failed with infrastructure error; leaving message uncommitted for retry",
				zap.Error(err),
				zap.String("correlation_id", payload.CorrelationID))
			return err
		}

		respBytes, _ := json.Marshal(resp)
		sendErr := m.kafka.SendMessageWithRetry(payload.ReplyTopic, payload.CorrelationID, respBytes, 3)
		if sendErr != nil {
			m.logger.Error("Failed to send Kafka response; leaving message uncommitted for retry", zap.Error(sendErr))
			return sendErr
		}

		session.MarkMessage(msg, "")
	}
	return nil
}
