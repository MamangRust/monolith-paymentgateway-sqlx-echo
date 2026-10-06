package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/MamangRust/monolith-payment-gateway-email/mailer"
	"github.com/MamangRust/monolith-payment-gateway-email/metrics"
	"github.com/MamangRust/monolith-payment-gateway-pkg/emailretry"
	"github.com/MamangRust/monolith-payment-gateway-pkg/event"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-pkg/outbox"
	traceunic "github.com/MamangRust/monolith-payment-gateway-pkg/trace_unic"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// retryHandler is the Phase 4 retry processor. It drains the shared retry
// topic ordered per partition (backoff via the _retryAt header), re-attempts
// SMTP, and escalates: success → mark inbox processed + commit; transient
// failure with attempts left → re-publish to the retry topic with attempt+1;
// max attempts reached → publish to the DLQ. A retry/DLQ publish must succeed
// before the retry-topic offset is committed.
type retryHandler struct {
	ctx             context.Context
	trace           trace.Tracer
	logger          logger.LoggerInterface
	Mailer          mailer.MailerInterface
	inbox           outbox.ConsumerInbox
	consumerName    string
	producer        retryPublisher
	maxRetries      int
	retryBackoff    time.Duration
	requestCounter  metric.Int64Counter
	requestDuration metric.Float64Histogram
}

// NewRetryHandler builds the retry processor. maxRetries and retryBackoff fall
// back to the shared defaults when non-positive.
func NewRetryHandler(ctx context.Context, logger logger.LoggerInterface, mailer mailer.MailerInterface, inbox outbox.ConsumerInbox, consumerName string, producer retryPublisher, maxRetries int, retryBackoff time.Duration) *retryHandler {
	if maxRetries <= 0 {
		maxRetries = emailretry.DefaultMaxAttempts
	}
	if retryBackoff <= 0 {
		retryBackoff = emailretry.DefaultBackoff
	}
	meter := otel.Meter("email-service")

	requestCounter, err := meter.Int64Counter(
		"email_retry_requests_total",
		metric.WithDescription("Total number of retry-processor requests"),
		metric.WithUnit("1"),
	)
	if err != nil {
		panic(err)
	}
	requestDuration, err := meter.Float64Histogram(
		"email_retry_request_duration_seconds",
		metric.WithDescription("Histogram of retry-processor request durations"),
		metric.WithUnit("s"),
	)
	if err != nil {
		panic(err)
	}

	return &retryHandler{
		ctx:             ctx,
		logger:          logger,
		Mailer:          mailer,
		trace:           otel.Tracer("email-retry-handler"),
		inbox:           inbox,
		consumerName:    consumerName,
		producer:        producer,
		maxRetries:      maxRetries,
		retryBackoff:    retryBackoff,
		requestCounter:  requestCounter,
		requestDuration: requestDuration,
	}
}

func (h *retryHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *retryHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h *retryHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	start := time.Now()
	status := "success"
	defer func() {
		attrs := metric.WithAttributes(
			attribute.String("method", "ConsumeClaim"),
			attribute.String("status", status),
		)
		h.requestCounter.Add(context.Background(), 1, attrs)
		h.requestDuration.Record(context.Background(), time.Since(start).Seconds(), attrs)
	}()

	for msg := range claim.Messages() {
		ctx := otel.GetTextMapPropagator().Extract(h.ctx, kafkaHeaderCarrier(msg.Headers))
		ctx, span := h.trace.Start(ctx, "retry:"+msg.Topic)
		span.SetAttributes(
			attribute.String("messaging.system", "kafka"),
			attribute.String("messaging.destination", msg.Topic),
			attribute.Int64("messaging.kafka.partition", int64(msg.Partition)),
			attribute.Int64("messaging.kafka.offset", msg.Offset),
		)

		meta, err := emailretry.ParseMeta(msg)
		if err != nil {
			h.logger.Error("Retry message without valid retry metadata; committing as poison", zap.Error(err))
			status = "failed_invalid_message"
			metrics.EmailInvalid.Add(ctx, 1)
			span.End()
			sess.MarkMessage(msg, "")
			sess.Commit()
			continue
		}

		// Ordered per-partition backoff: an attempt is not processed before its
		// scheduled _retryAt time (capped by emailretry.MaxBackoff).
		if delay := time.Until(meta.RetryAt); delay > 0 {
			time.Sleep(delay)
		}

		var payload event.EmailEnvelope
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			traceID := traceunic.GenerateTraceID("FAILED_UNMARSHAL_RETRY_MESSAGE")
			h.logger.Error("Failed to unmarshal retry message; committing as poison", zap.Error(err))
			span.SetAttributes(attribute.String("trace.id", traceID))
			span.RecordError(err)
			status = "failed_unmarshal_message"
			metrics.EmailInvalid.Add(ctx, 1)
			span.End()
			sess.MarkMessage(msg, "")
			sess.Commit()
			continue
		}
		if !payload.IsValid() {
			h.logger.Error("Invalid email envelope in retry message; committing as poison")
			status = "failed_invalid_message"
			metrics.EmailInvalid.Add(ctx, 1)
			span.End()
			sess.MarkMessage(msg, "")
			sess.Commit()
			continue
		}

		eventKey := fmt.Sprintf("%s:%s", meta.SrcTopic, payload.EventID)
		var reservationVersion int64
		if h.inbox != nil {
			reserved, processed, version, err := h.inbox.Reserve(ctx, h.consumerName, eventKey, meta.SrcTopic, meta.SrcPartition, meta.SrcOffset)
			if err != nil {
				status = "failed_inbox_reserve"
				span.End()
				return err
			}
			if processed {
				// The email was already sent by an earlier attempt (or by the
				// main consumer before its offset commit) — nothing to do.
				span.End()
				sess.MarkMessage(msg, "")
				sess.Commit()
				continue
			}
			if !reserved {
				status = "inbox_lease_busy"
				span.End()
				return fmt.Errorf("consumer inbox lease is active for event %s", eventKey)
			}
			reservationVersion = version
		}

		if err := h.sendAndEscalate(ctx, sess, msg, meta, payload, reservationVersion); err != nil {
			status = "failed_escalation"
			span.End()
			return err
		}
		span.End()
		sess.MarkMessage(msg, "")
		sess.Commit()
	}

	h.logger.Info("Retry ConsumeClaim finished", zap.Int("messages", len(claim.Messages())))
	return nil
}

// sendAndEscalate attempts SMTP and, on failure, either re-schedules the event
// on the retry topic or dead-letters it. The caller commits the retry-topic
// offset only when this returns nil (i.e. the escalation publish succeeded).
func (h *retryHandler) sendAndEscalate(ctx context.Context, sess sarama.ConsumerGroupSession, msg *sarama.ConsumerMessage, meta emailretry.RetryMeta, payload event.EmailEnvelope, reservationVersion int64) error {
	eventKey := fmt.Sprintf("%s:%s", meta.SrcTopic, payload.EventID)

	err := h.Mailer.Send(payload.Email, payload.Subject, payload.Body)
	if err == nil {
		metrics.EmailSent.Add(ctx, 1)
		if h.inbox != nil {
			if markErr := h.inbox.MarkProcessed(ctx, h.consumerName, eventKey, reservationVersion); markErr != nil {
				return markErr
			}
		}
		return nil
	}

	h.logger.Error("Retry attempt failed", zap.Error(err), zap.Int("attempt", meta.Attempt), zap.String("event_key", eventKey))
	metrics.EmailFailed.Add(ctx, 1)
	if h.inbox != nil {
		if releaseErr := h.inbox.Release(ctx, h.consumerName, eventKey, reservationVersion, err); releaseErr != nil {
			h.logger.Error("failed to release consumer inbox lease", zap.Error(releaseErr), zap.String("event_key", eventKey))
		}
	}

	if meta.Attempt >= h.maxRetries {
		if pubErr := h.publishDLQ(ctx, payload.EventID, msg.Value, meta, err); pubErr != nil {
			return fmt.Errorf("publish to DLQ topic: %w", pubErr)
		}
		metrics.EmailDeadLetter.Add(ctx, 1)
		return nil
	}

	next := emailretry.RetryMeta{
		SrcTopic: meta.SrcTopic, SrcPartition: meta.SrcPartition, SrcOffset: meta.SrcOffset,
		Attempt: meta.Attempt + 1, RetryAt: emailretry.NextRetryAt(meta.Attempt+1, h.retryBackoff), Reason: err.Error(),
	}
	if pubErr := h.publishRetry(ctx, payload.EventID, msg.Value, next); pubErr != nil {
		return fmt.Errorf("publish to retry topic: %w", pubErr)
	}
	metrics.EmailRetried.Add(ctx, 1)
	return nil
}

func (h *retryHandler) publishRetry(ctx context.Context, key string, value []byte, meta emailretry.RetryMeta) error {
	return h.producer.SendMessageWithHeaders(ctx, emailretry.RetryTopic, key, value, emailretry.BuildHeaders(meta))
}

func (h *retryHandler) publishDLQ(ctx context.Context, key string, value []byte, meta emailretry.RetryMeta, lastErr error) error {
	meta.Attempts = meta.Attempt
	meta.Reason = lastErr.Error()
	return h.producer.SendMessageWithHeaders(ctx, emailretry.DLQTopic, key, value, emailretry.BuildDLQHeaders(meta))
}
