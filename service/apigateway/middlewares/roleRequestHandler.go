package middlewares

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"go.uber.org/zap"
)

// roleResponseHandler handles Kafka consumer group events related to role validation responses.
//
// It implements the sarama.ConsumerGroupHandler interface, and is responsible for
// setting up and cleaning up the Kafka consumer session when consuming role validation responses.
type roleResponseHandler struct {
	// validator is the RoleValidator instance that manages Kafka-based role response handling.
	validator *RoleValidator
}

// Setup is called at the beginning of a new Kafka consumer group session.
//
// It can be used to initialize resources or state before message consumption begins.
// In this implementation, it performs no setup and returns nil.
func (h *roleResponseHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is called at the end of a Kafka consumer group session.
//
// It can be used to release resources allocated during Setup or message consumption.
// In this implementation, it performs no cleanup and returns nil.
func (h *roleResponseHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim processes incoming Kafka messages from the specified consumer group claim.
// It unmarshals each message into a payload map and retrieves the correlation ID.
// If a valid correlation ID is found, it sends the message value to the corresponding
// response channel managed by the validator. Each message is marked as processed
// in the consumer group session.
func (h *roleResponseHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		// Log penerimaan pesan
		h.validator.logger.Debug("Received Kafka response message",
			zap.String("topic", msg.Topic),
			// Key Kafka digunakan sebagai correlationID
			zap.String("key", string(msg.Key)))

		// Unmarshal payload respons
		var roleResponse response.RoleResponsePayload
		if err := json.Unmarshal(msg.Value, &roleResponse); err != nil {
			h.validator.logger.Error("Failed to unmarshal role response", zap.Error(err))
			session.MarkMessage(msg, "")
			continue
		}

		// Dapatkan correlationID dari payload
		correlationID := roleResponse.CorrelationID
		if correlationID == "" {
			h.validator.logger.Error("Missing correlation ID in response")
			session.MarkMessage(msg, "")
			continue
		}

		// Cari channel yang sesuai di map validator
		h.validator.mu.RLock()
		ch, exists := h.validator.responseChans[correlationID]
		h.validator.mu.RUnlock()

		if exists && ch != nil {
			// Kirim respons ke channel yang menunggu
			select {
			case ch <- &roleResponse:
				h.validator.logger.Debug("Response delivered to channel",
					zap.String("correlation_id", correlationID))
			default:
				// Channel penuh atau sudah ditutup (kemungkinan kecil karena buffer 1 dan defer)
				h.validator.logger.Debug("Response channel full or closed",
					zap.String("correlation_id", correlationID))
			}
		} else {
			// Tidak ada channel yang menunggu untuk correlationID ini
			// Ini adalah masalah utama yang menyebabkan timeout
			h.validator.logger.Debug("No waiting channel for correlation ID",
				zap.String("correlation_id", correlationID))
		}
		// Tandai pesan sebagai telah diproses
		session.MarkMessage(msg, "")
	}
	return nil
}
