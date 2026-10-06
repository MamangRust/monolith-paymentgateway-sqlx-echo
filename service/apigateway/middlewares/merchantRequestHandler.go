package middlewares

import (
	"encoding/json"

	"github.com/IBM/sarama"
)

type merchantResponseHandler struct {
	validator *ApiKeyValidator
}

func NewResponseHandler(validator *ApiKeyValidator) *merchantResponseHandler {
	return &merchantResponseHandler{
		validator: validator,
	}
}

func (h *merchantResponseHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *merchantResponseHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *merchantResponseHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var payload map[string]interface{}
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			continue
		}

		correlationID, ok := payload["correlation_id"].(string)
		if !ok {
			continue
		}

		h.validator.mu.Lock()
		ch, ok := h.validator.responseChans[correlationID]
		h.validator.mu.Unlock()

		if ok {
			ch <- msg.Value
		}

		session.MarkMessage(msg, "")
	}
	return nil
}
