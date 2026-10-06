package kafka

import (
	"context"
	"errors"
	"testing"

	"github.com/IBM/sarama"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type billingTestService struct {
	err       error
	calledDay int
}

func (f *billingTestService) TriggerBillingCycle(_ context.Context, day int) (int, error) {
	f.calledDay = day
	return 0, f.err
}
func (f *billingTestService) GetStatement(context.Context, string) (*db.BillingCycle, error) {
	return nil, nil
}
func (f *billingTestService) GetStatementsByCard(context.Context, string, int, int) ([]*db.BillingCycle, error) {
	return nil, nil
}
func (f *billingTestService) GetBillingCyclesByCardNumber(context.Context, string) ([]*db.BillingCycle, error) {
	return nil, nil
}

type billingTestLogger struct{}

func (billingTestLogger) Info(string, ...zap.Field)                         {}
func (billingTestLogger) Fatal(string, ...zap.Field)                        {}
func (billingTestLogger) Debug(string, ...zap.Field)                        {}
func (billingTestLogger) Error(string, ...zap.Field)                        {}
func (billingTestLogger) Warn(string, ...zap.Field)                         {}
func (billingTestLogger) Check(zapcore.Level, string) *zapcore.CheckedEntry { return nil }
func (billingTestLogger) With(...zap.Field) logger.LoggerInterface          { return billingTestLogger{} }
func (billingTestLogger) Sync() error                                       { return nil }

type billingTestSession struct{ marked int }

func (s *billingTestSession) Claims() map[string][]int32                  { return nil }
func (s *billingTestSession) MemberID() string                            { return "billing-test" }
func (s *billingTestSession) GenerationID() int32                         { return 1 }
func (s *billingTestSession) MarkOffset(string, int32, int64, string)     {}
func (s *billingTestSession) ResetOffset(string, int32, int64, string)    {}
func (s *billingTestSession) MarkMessage(*sarama.ConsumerMessage, string) { s.marked++ }
func (s *billingTestSession) Context() context.Context                    { return context.Background() }
func (s *billingTestSession) Commit()                                     {}
func (s *billingTestSession) AsyncClose()                                 {}

type billingTestClaim struct{ message *sarama.ConsumerMessage }

func (c *billingTestClaim) Topic() string              { return "card.billing.trigger" }
func (c *billingTestClaim) Partition() int32           { return 0 }
func (c *billingTestClaim) InitialOffset() int64       { return 0 }
func (c *billingTestClaim) HighWaterMarkOffset() int64 { return 1 }
func (c *billingTestClaim) Messages() <-chan *sarama.ConsumerMessage {
	ch := make(chan *sarama.ConsumerMessage, 1)
	ch <- c.message
	close(ch)
	return ch
}

const billingPayload = "{\"billing_cycle_day\":25}"

func TestBillingTriggerHandler_ServiceErrorReturnsWithoutAck(t *testing.T) {
	handler := NewBillingTriggerHandler(&billingTestService{err: errors.New("database unavailable")}, billingTestLogger{})
	session := &billingTestSession{}

	err := handler.ConsumeClaim(session, &billingTestClaim{
		message: &sarama.ConsumerMessage{Value: []byte(billingPayload)},
	})
	if err == nil {
		t.Fatal("expected billing service error")
	}
	if session.marked != 0 {
		t.Fatalf("expected failed message to remain uncommitted, got %d marks", session.marked)
	}
}

func TestBillingTriggerHandler_SuccessAcknowledges(t *testing.T) {
	service := &billingTestService{}
	handler := NewBillingTriggerHandler(service, billingTestLogger{})
	session := &billingTestSession{}

	if err := handler.ConsumeClaim(session, &billingTestClaim{
		message: &sarama.ConsumerMessage{Value: []byte(billingPayload)},
	}); err != nil {
		t.Fatalf("ConsumeClaim returned error: %v", err)
	}
	if session.marked != 1 {
		t.Fatalf("expected successful message to be acknowledged, got %d marks", session.marked)
	}
	if service.calledDay != 25 {
		t.Fatalf("expected billing day 25, got %d", service.calledDay)
	}
}
