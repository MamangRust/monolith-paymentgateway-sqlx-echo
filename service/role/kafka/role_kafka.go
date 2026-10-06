package myhandlerkafka // Ganti dengan nama package yang sesuai

import (
	"context"
	"encoding/json"
	"time"

	"errors"
	"github.com/IBM/sarama"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-role/service" // Sesuaikan path
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	role_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/role_errors/repository"
	"go.uber.org/zap"
)

// roleKafkaHandler is a struct that implements the sarama.ConsumerGroupHandler interface
type roleKafkaHandler struct {
	logger      logger.LoggerInterface
	roleService service.RoleQueryService
	kafka       *kafka.Kafka
	ctx         context.Context
}

// NewRoleKafkaHandler creates a new Kafka consumer group handler for processing role validation responses.
//
// It takes a role query service, a Kafka producer, and a logger as parameters.
// The handler is used to process incoming Kafka messages from the role validation response topic.
// It implements the sarama.ConsumerGroupHandler interface to handle consumer group lifecycle events.
func NewRoleKafkaHandler(roleService service.RoleQueryService, kafka *kafka.Kafka, logger logger.LoggerInterface, ctx context.Context) sarama.ConsumerGroupHandler {
	return &roleKafkaHandler{
		roleService: roleService,
		kafka:       kafka,
		logger:      logger,
		ctx:         ctx,
	}
}

// Setup is a method that is called when the Kafka consumer group is set up
func (h *roleKafkaHandler) Setup(session sarama.ConsumerGroupSession) error {
	h.logger.Info("Role Kafka handler setup completed")
	return nil
}

// Cleanup is a method that is called when the Kafka consumer group is cleaned up
func (h *roleKafkaHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	h.logger.Info("Role Kafka handler cleanup completed")
	return nil
}

// ConsumeClaim is a method that is called when the Kafka consumer group has messages to process
//
// It processes incoming Kafka messages from the role validation request topic.
// It unmarshals each message into a payload map and retrieves the correlation ID.
// If a valid correlation ID is found, it sends the message value to the corresponding
// response channel managed by the validator. Each message is marked as processed
// in the consumer group session.
func (h *roleKafkaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		msgCtx, cancel := context.WithTimeout(h.ctx, 20*time.Second)
		defer cancel()

		h.logger.Info("Received role validation request",
			zap.String("topic", msg.Topic),
			zap.String("key", string(msg.Key)))

		var payload requests.RoleRequestPayload
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			h.logger.Error("Invalid role request payload", zap.Error(err))
			session.MarkMessage(msg, "")
			continue
		}

		if payload.CorrelationID == "" || payload.ReplyTopic == "" {
			h.logger.Error("Missing required fields in role request",
				zap.String("correlation_id", payload.CorrelationID),
				zap.String("reply_topic", payload.ReplyTopic))
			session.MarkMessage(msg, "")
			continue
		}

		h.logger.Info("Processing role validation request",
			zap.Int("user_id", payload.UserID),
			zap.String("correlation_id", payload.CorrelationID))

		roles, errResp := h.roleService.FindByUserId(msgCtx, payload.UserID)
		if errResp != nil && !errors.Is(errResp, role_errors.ErrRoleNotFound) {
			h.logger.Error("Role lookup failed; leaving message uncommitted for retry", zap.Error(errResp), zap.Int("user_id", payload.UserID))
			cancel()
			return errResp
		}

		resp := response.RoleResponsePayload{
			CorrelationID: payload.CorrelationID,
			Valid:         errResp == nil && len(roles) > 0,
			RoleNames:     make([]string, 0),
		}

		if errResp == nil && len(roles) > 0 {
			for _, r := range roles {
				resp.RoleNames = append(resp.RoleNames, r.RoleName)
			}
			h.logger.Info("Role validation successful",
				zap.Int("user_id", payload.UserID),
				zap.Strings("roles", resp.RoleNames),
				zap.String("correlation_id", payload.CorrelationID))
		} else {
			h.logger.Debug("Role validation failed",
				zap.Int("user_id", payload.UserID),
				zap.Any("error", errResp),
				zap.String("correlation_id", payload.CorrelationID))
		}

		respBytes, err := json.Marshal(resp)
		if err != nil {
			h.logger.Error("Failed to marshal role response",
				zap.Any("error", err),
				zap.String("correlation_id", payload.CorrelationID))
			session.MarkMessage(msg, "")
			continue
		}

		err = h.kafka.SendMessageWithRetry(payload.ReplyTopic, payload.CorrelationID, respBytes, 3)
		if err != nil {
			h.logger.Error("Failed to send Kafka role response; leaving message uncommitted for retry",
				zap.Any("error", err),
				zap.String("reply_topic", payload.ReplyTopic),
				zap.String("correlation_id", payload.CorrelationID))
			return err
		} else {
			h.logger.Info("Role response sent successfully",
				zap.String("reply_topic", payload.ReplyTopic),
				zap.String("correlation_id", payload.CorrelationID))
		}

		session.MarkMessage(msg, "")
	}
	return nil
}
