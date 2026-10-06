package middlewares

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	mencache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ApiKeyValidator struct {
	kafka *kafka.Kafka

	logger logger.LoggerInterface

	requestTopic string

	responseTopic string

	timeout time.Duration

	responseChans map[string]chan []byte

	mu sync.Mutex

	cache mencache.MerchantCache
}

func NewApiKeyValidator(k *kafka.Kafka, requestTopic, responseTopic string, timeout time.Duration, logger logger.LoggerInterface, cache mencache.MerchantCache) *ApiKeyValidator {
	v := &ApiKeyValidator{
		kafka:         k,
		requestTopic:  requestTopic,
		responseTopic: responseTopic,
		timeout:       timeout,
		responseChans: make(map[string]chan []byte),
		logger:        logger,
		cache:         cache,
	}

	handler := &merchantResponseHandler{validator: v}
	if k != nil {
		go func() {
			err := k.StartConsumers([]string{responseTopic}, "api-gateway-group", handler)
			if err != nil {
				v.logger.Fatal("Failed to start kafka consumer", zap.Error(err))
				panic("failed to start kafka c	onsumer: " + err.Error())
			}
		}()
	}

	return v
}

func (v *ApiKeyValidator) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiKey := c.Request().Header.Get("X-Api-Key")
			if apiKey == "" {
				v.logger.Error("Missing API Key in request header")
				return echo.NewHTTPError(http.StatusUnauthorized, "API Key is required")
			}

			if merchantID, found := v.cache.GetMerchantCache(c.Request().Context(), apiKey); found {
				v.logger.Info("Merchant ID found in cache", zap.String("apiKey", observability.MaskIdentifier(apiKey)), zap.String("merchant_id", merchantID))
				c.Set("merchant_id", merchantID)
				c.Set("apiKey", apiKey)
				return next(c)
			}

			correlationID := uuid.NewString()
			v.logger.Info("Cache miss, sending Kafka request", zap.String("apiKey", observability.MaskIdentifier(apiKey)), zap.String("correlationID", correlationID))

			payload := requests.MerchantRequestPayload{
				ApiKey:        apiKey,
				CorrelationID: correlationID,
				ReplyTopic:    v.responseTopic,
			}

			data, err := json.Marshal(payload)
			if err != nil {
				v.logger.Error("Failed to encode payload", zap.Error(err), zap.String("correlationID", correlationID))
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to encode payload")
			}

			if v.kafka != nil {
				err = v.kafka.SendMessage(v.requestTopic, correlationID, data)
				if err != nil {
					v.logger.Error("Failed to send Kafka message", zap.Error(err), zap.String("correlationID", correlationID))
					return echo.NewHTTPError(http.StatusInternalServerError, "Failed to send Kafka message")
				}
			} else {
				v.logger.Warn("Kafka is nil, skipping Kafka-based API key validation", zap.String("correlationID", correlationID))
				// In test environment without Kafka, we might want to just proceed if it's already in cache or mock it.
				// For now, if it's not in cache and no Kafka, we can't validate.
				return echo.NewHTTPError(http.StatusServiceUnavailable, "Kafka is required for API key validation but is nil")
			}

			respChan := make(chan []byte, 1)

			v.mu.Lock()
			v.responseChans[correlationID] = respChan
			v.mu.Unlock()

			defer func() {
				v.mu.Lock()
				delete(v.responseChans, correlationID)
				v.mu.Unlock()
			}()

			select {
			case msg := <-respChan:
				var response response.MerchantResponsePayload
				if err := json.Unmarshal(msg, &response); err != nil {
					v.logger.Error("Failed to decode Kafka response", zap.Error(err), zap.String("correlationID", correlationID))
					return echo.NewHTTPError(http.StatusInternalServerError, "Invalid response format")
				}

				if !response.Valid || response.MerchantID == 0 {
					v.logger.Error("Invalid API Key validation result", zap.String("correlationID", correlationID), zap.String("apiKey", observability.MaskIdentifier(apiKey)))
					return echo.NewHTTPError(http.StatusUnauthorized, "Invalid API Key")
				}

				merchantIDStr := strconv.Itoa(int(response.MerchantID))

				v.cache.SetMerchantCache(c.Request().Context(), merchantIDStr, apiKey)

				v.logger.Info("API Key validated successfully", zap.String("merchant_id", merchantIDStr), zap.String("correlationID", correlationID))

				c.Set("merchant_id", merchantIDStr)
				c.Set("apiKey", apiKey)
				return next(c)

			case <-time.After(v.timeout):
				v.logger.Error("Timeout waiting for Kafka response", zap.String("correlationID", correlationID), zap.Duration("timeout", v.timeout))
				return echo.NewHTTPError(http.StatusRequestTimeout, "Timeout waiting for API key validation")
			}
		}
	}
}
