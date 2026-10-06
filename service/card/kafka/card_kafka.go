package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/MamangRust/monolith-payment-gateway-card/service"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"go.uber.org/zap"
)

// cardKafkaHandler is a struct that implements the sarama.ConsumerGroupHandler interface
type cardKafkaHandler struct {
	logger      logger.LoggerInterface
	cardService service.CardCommandService
}

// NewCardKafkaHandler initializes a new cardKafkaHandler with the provided cardService and logger.
// It returns an instance of the cardKafkaHandler struct.
func NewCardKafkaHandler(cardService service.CardCommandService, logger logger.LoggerInterface) *cardKafkaHandler {
	return &cardKafkaHandler{
		cardService: cardService,
		logger:      logger,
	}
}

// Setup is called when a new Kafka consumer group session begins.
//
// It can be used to initialize resources or state before message consumption begins.
// In this implementation, it performs no setup and returns nil.
func (s *cardKafkaHandler) Setup(session sarama.ConsumerGroupSession) error {
	s.logger.Info("card kafka handler setup")
	return nil
}

// Cleanup is called at the end of a Kafka consumer group session.
//
// It can be used to release resources allocated during Setup or message consumption.
// In this implementation, it performs no cleanup and returns nil.
func (s *cardKafkaHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	s.logger.Info("card kafka handler cleanup")
	return nil
}

// ConsumeClaim processes incoming Kafka messages from the specified consumer group claim.
// It unmarshals each message into a payload map and creates a new card request.
// If a valid card request is found, it calls the CreateCard method of the cardService
// with the created request. If the CreateCard method returns an error, it logs the
// error and returns an error with the message "card service error: <error message>".
// Each message is marked as processed in the consumer group session.
func (s *cardKafkaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for msg := range claim.Messages() {
		var payload struct {
			UserID       int       `json:"user_id"`
			CardType     string    `json:"card_type"`
			ExpireDate   time.Time `json:"expire_date"`
			CVV          string    `json:"cvv"`
			CardProvider string    `json:"card_provider"`
		}
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			s.logger.Error("invalid card Kafka payload", zap.Error(err))
			session.MarkMessage(msg, "")
			continue
		}
		if payload.UserID <= 0 || payload.CardType == "" || payload.ExpireDate.IsZero() || payload.CVV == "" || payload.CardProvider == "" {
			s.logger.Error("card Kafka payload missing required fields")
			session.MarkMessage(msg, "")
			continue
		}

		card := &requests.CreateCardRequest{
			UserID:       payload.UserID,
			CardType:     payload.CardType,
			ExpireDate:   payload.ExpireDate,
			CVV:          payload.CVV,
			CardProvider: payload.CardProvider,
		}

		_, errRes := s.cardService.CreateCard(ctx, card)

		if errRes != nil {
			s.logger.Error("card service error", zap.Any("error", errRes))

			return fmt.Errorf("card service error: %v", errRes)
		}
		session.MarkMessage(msg, "")
	}

	return nil
}
