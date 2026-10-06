package myhandlerkafka

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/IBM/sarama"
	"github.com/MamangRust/monolith-payment-gateway-merchant/service"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/repository"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ────────────────────────────────────────────────────────────────────────────
// Fakes
// ────────────────────────────────────────────────────────────────────────────

// fakeMerchantQueryService implements service.MerchantQueryService; only
// FindByApiKey is exercised by the Kafka handler.
type fakeMerchantQueryService struct {
	merchant *db.GetMerchantByApiKeyRow
	err      error
}

func (f *fakeMerchantQueryService) FindAll(context.Context, *requests.FindAllMerchants) ([]*db.GetMerchantsRow, *int, error) {
	return nil, nil, nil
}
func (f *fakeMerchantQueryService) FindById(context.Context, int) (*db.GetMerchantByIDRow, error) {
	return nil, nil
}
func (f *fakeMerchantQueryService) FindByActive(context.Context, *requests.FindAllMerchants) ([]*db.GetActiveMerchantsRow, *int, error) {
	return nil, nil, nil
}
func (f *fakeMerchantQueryService) FindByTrashed(context.Context, *requests.FindAllMerchants) ([]*db.GetTrashedMerchantsRow, *int, error) {
	return nil, nil, nil
}
func (f *fakeMerchantQueryService) FindByApiKey(context.Context, string) (*db.GetMerchantByApiKeyRow, error) {
	return f.merchant, f.err
}
func (f *fakeMerchantQueryService) FindByMerchantUserId(context.Context, int) ([]*db.GetMerchantsByUserIDRow, error) {
	return nil, nil
}

// fakeSyncProducer implements kafka.SyncProducer and records every message.
type fakeSyncProducer struct {
	sent []*sarama.ProducerMessage
}

func (f *fakeSyncProducer) SendMessage(msg *sarama.ProducerMessage) (int32, int64, error) {
	f.sent = append(f.sent, msg)
	return 0, 0, nil
}

func (f *fakeSyncProducer) Close() error { return nil }

// fakeLogger implements logger.LoggerInterface as a no-op.
type fakeLogger struct{}

func (f *fakeLogger) Info(string, ...zap.Field)                         {}
func (f *fakeLogger) Fatal(string, ...zap.Field)                        {}
func (f *fakeLogger) Debug(string, ...zap.Field)                        {}
func (f *fakeLogger) Error(string, ...zap.Field)                        {}
func (f *fakeLogger) Warn(string, ...zap.Field)                         {}
func (f *fakeLogger) Check(zapcore.Level, string) *zapcore.CheckedEntry { return nil }
func (f *fakeLogger) With(...zap.Field) logger.LoggerInterface          { return f }
func (f *fakeLogger) Sync() error                                       { return nil }

// fakeSession implements sarama.ConsumerGroupSession, recording commits.
type fakeSession struct {
	marked int
}

func (s *fakeSession) Claims() map[string][]int32                  { return nil }
func (s *fakeSession) MemberID() string                            { return "test-member" }
func (s *fakeSession) GenerationID() int32                         { return 0 }
func (s *fakeSession) MarkOffset(string, int32, int64, string)     {}
func (s *fakeSession) ResetOffset(string, int32, int64, string)    {}
func (s *fakeSession) MarkMessage(*sarama.ConsumerMessage, string) { s.marked++ }
func (s *fakeSession) Context() context.Context                    { return context.Background() }
func (s *fakeSession) Commit()                                     {}
func (s *fakeSession) AsyncClose()                                 {}

// fakeClaim implements sarama.ConsumerGroupClaim over a fixed message list.
type fakeClaim struct {
	topic string
	msgs  []*sarama.ConsumerMessage
}

func (c *fakeClaim) Topic() string              { return c.topic }
func (c *fakeClaim) Partition() int32           { return 0 }
func (c *fakeClaim) InitialOffset() int64       { return 0 }
func (c *fakeClaim) HighWaterMarkOffset() int64 { return int64(len(c.msgs)) }
func (c *fakeClaim) Messages() <-chan *sarama.ConsumerMessage {
	ch := make(chan *sarama.ConsumerMessage, len(c.msgs))
	for _, m := range c.msgs {
		ch <- m
	}
	close(ch)
	return ch
}

// ────────────────────────────────────────────────────────────────────────────
// Helpers
// ────────────────────────────────────────────────────────────────────────────

func newTestHandler(svc service.MerchantQueryService) (*merchantKafkaHandler, *fakeSyncProducer, *fakeSession) {
	producer := &fakeSyncProducer{}
	k := kafka.NewKafkaWithProducer(producer, &fakeLogger{}, nil)
	handler := NewMerchantKafkaHandler(svc, k, &fakeLogger{}).(*merchantKafkaHandler)
	session := &fakeSession{}
	return handler, producer, session
}

func requestMessage(t *testing.T, apiKey string) *sarama.ConsumerMessage {
	t.Helper()
	payload := requests.MerchantRequestPayload{
		ApiKey:        apiKey,
		CorrelationID: "corr-123",
		ReplyTopic:    "merchant.apikey.response",
	}
	b, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}
	return &sarama.ConsumerMessage{Value: b}
}

func decodeResponse(t *testing.T, msg *sarama.ProducerMessage) response.MerchantResponsePayload {
	t.Helper()
	var resp response.MerchantResponsePayload
	value, ok := msg.Value.(sarama.ByteEncoder)
	if !ok {
		t.Fatalf("unexpected value type %T", msg.Value)
	}
	if err := json.Unmarshal([]byte(value), &resp); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}
	return resp
}

// ────────────────────────────────────────────────────────────────────────────
// Tests
// ────────────────────────────────────────────────────────────────────────────

func TestMerchantKafkaHandler_ValidApiKeyRepliesTrue(t *testing.T) {
	svc := &fakeMerchantQueryService{
		merchant: &db.GetMerchantByApiKeyRow{MerchantID: 42, ApiKey: "secret-key"},
	}
	handler, producer, session := newTestHandler(svc)
	claim := &fakeClaim{msgs: []*sarama.ConsumerMessage{requestMessage(t, "secret-key")}}

	if err := handler.ConsumeClaim(session, claim); err != nil {
		t.Fatalf("ConsumeClaim: %v", err)
	}

	if len(producer.sent) != 1 {
		t.Fatalf("expected 1 reply, got %d", len(producer.sent))
	}
	if session.marked != 1 {
		t.Fatalf("expected message to be marked, got %d marks", session.marked)
	}
	resp := decodeResponse(t, producer.sent[0])
	if !resp.Valid || resp.MerchantID != 42 {
		t.Fatalf("expected Valid=true MerchantID=42, got %+v", resp)
	}
	if resp.CorrelationID != "corr-123" {
		t.Fatalf("correlation id not echoed back, got %q", resp.CorrelationID)
	}
}

func TestMerchantKafkaHandler_UnknownApiKeyRepliesFalse(t *testing.T) {
	svc := &fakeMerchantQueryService{
		err: merchant_errors.ErrMerchantNotFound.WithInternal(errors.New("no rows")),
	}
	handler, producer, session := newTestHandler(svc)
	claim := &fakeClaim{msgs: []*sarama.ConsumerMessage{requestMessage(t, "unknown-key")}}

	if err := handler.ConsumeClaim(session, claim); err != nil {
		t.Fatalf("ConsumeClaim: %v", err)
	}

	if len(producer.sent) != 1 {
		t.Fatalf("expected 1 reply, got %d", len(producer.sent))
	}
	if session.marked != 1 {
		t.Fatalf("expected not-found to be committed, got %d marks", session.marked)
	}
	resp := decodeResponse(t, producer.sent[0])
	if resp.Valid {
		t.Fatalf("expected Valid=false for unknown api key, got %+v", resp)
	}
}

func TestMerchantKafkaHandler_NilMerchantWithoutErrorRepliesFalse(t *testing.T) {
	// A service returning (nil, nil) — no merchant, no error — must be treated
	// as a deterministic not-found (Valid:false, committed), not as an infra
	// error that would leave the gateway hanging until timeout.
	svc := &fakeMerchantQueryService{}
	handler, producer, session := newTestHandler(svc)
	claim := &fakeClaim{msgs: []*sarama.ConsumerMessage{requestMessage(t, "ghost-key")}}

	if err := handler.ConsumeClaim(session, claim); err != nil {
		t.Fatalf("ConsumeClaim: %v", err)
	}

	if len(producer.sent) != 1 {
		t.Fatalf("expected 1 reply for nil-merchant, got %d", len(producer.sent))
	}
	if session.marked != 1 {
		t.Fatalf("expected nil-merchant to be committed, got %d marks", session.marked)
	}
	resp := decodeResponse(t, producer.sent[0])
	if resp.Valid {
		t.Fatalf("expected Valid=false for nil merchant, got %+v", resp)
	}
}

func TestMerchantKafkaHandler_MalformedPayloadIsAcknowledged(t *testing.T) {
	handler, producer, session := newTestHandler(&fakeMerchantQueryService{})
	if err := handler.ConsumeClaim(session, &fakeClaim{msgs: []*sarama.ConsumerMessage{{Value: []byte("not-json")}}}); err != nil {
		t.Fatalf("ConsumeClaim: %v", err)
	}
	if len(producer.sent) != 0 {
		t.Fatalf("expected no reply for malformed payload, got %d", len(producer.sent))
	}
	if session.marked != 1 {
		t.Fatalf("expected malformed payload to be acknowledged, got %d marks", session.marked)
	}
}

func TestMerchantKafkaHandler_InfraErrorLeavesMessageUncommitted(t *testing.T) {
	svc := &fakeMerchantQueryService{
		err: errors.New("database connection refused"),
	}
	handler, producer, session := newTestHandler(svc)
	claim := &fakeClaim{msgs: []*sarama.ConsumerMessage{requestMessage(t, "any-key")}}

	if err := handler.ConsumeClaim(session, claim); err == nil {
		t.Fatal("expected infrastructure error")
	}

	if len(producer.sent) != 0 {
		t.Fatalf("expected NO reply on infra error, got %d", len(producer.sent))
	}
	if session.marked != 0 {
		t.Fatalf("expected message NOT marked on infra error, got %d marks", session.marked)
	}
}
