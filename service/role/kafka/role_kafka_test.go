package myhandlerkafka

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/IBM/sarama"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-role/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	role_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/role_errors/repository"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type roleTestService struct {
	roles []*db.Role
	err   error
}

func (f *roleTestService) FindAll(context.Context, *requests.FindAllRoles) ([]*db.GetRolesRow, *int, error) {
	return nil, nil, nil
}
func (f *roleTestService) FindByActiveRole(context.Context, *requests.FindAllRoles) ([]*db.GetActiveRolesRow, *int, error) {
	return nil, nil, nil
}
func (f *roleTestService) FindByTrashedRole(context.Context, *requests.FindAllRoles) ([]*db.GetTrashedRolesRow, *int, error) {
	return nil, nil, nil
}
func (f *roleTestService) FindById(context.Context, int) (*db.Role, error) { return nil, nil }
func (f *roleTestService) FindByName(context.Context, string) (*db.Role, error) {
	return nil, nil
}
func (f *roleTestService) FindByUserId(context.Context, int) ([]*db.Role, error) {
	return f.roles, f.err
}

type roleTestProducer struct {
	sent int
}

func (p *roleTestProducer) SendMessage(*sarama.ProducerMessage) (int32, int64, error) {
	p.sent++
	return 0, 0, nil
}
func (p *roleTestProducer) Close() error { return nil }

type roleTestLogger struct{}

func (roleTestLogger) Info(string, ...zap.Field)                         {}
func (roleTestLogger) Fatal(string, ...zap.Field)                        {}
func (roleTestLogger) Debug(string, ...zap.Field)                        {}
func (roleTestLogger) Error(string, ...zap.Field)                        {}
func (roleTestLogger) Warn(string, ...zap.Field)                         {}
func (roleTestLogger) Check(zapcore.Level, string) *zapcore.CheckedEntry { return nil }
func (l roleTestLogger) With(...zap.Field) logger.LoggerInterface        { return l }
func (roleTestLogger) Sync() error                                       { return nil }

type roleTestSession struct{ marked int }

func (s *roleTestSession) Claims() map[string][]int32                  { return nil }
func (s *roleTestSession) MemberID() string                            { return "test" }
func (s *roleTestSession) GenerationID() int32                         { return 1 }
func (s *roleTestSession) MarkOffset(string, int32, int64, string)     {}
func (s *roleTestSession) ResetOffset(string, int32, int64, string)    {}
func (s *roleTestSession) MarkMessage(*sarama.ConsumerMessage, string) { s.marked++ }
func (s *roleTestSession) Context() context.Context                    { return context.Background() }
func (s *roleTestSession) Commit()                                     {}
func (s *roleTestSession) AsyncClose()                                 {}

type roleTestClaim struct{ messages []*sarama.ConsumerMessage }

func (c *roleTestClaim) Topic() string              { return "role.validation" }
func (c *roleTestClaim) Partition() int32           { return 0 }
func (c *roleTestClaim) InitialOffset() int64       { return 0 }
func (c *roleTestClaim) HighWaterMarkOffset() int64 { return int64(len(c.messages)) }
func (c *roleTestClaim) Messages() <-chan *sarama.ConsumerMessage {
	ch := make(chan *sarama.ConsumerMessage, len(c.messages))
	for _, message := range c.messages {
		ch <- message
	}
	close(ch)
	return ch
}

func roleRequestMessage(t *testing.T) *sarama.ConsumerMessage {
	t.Helper()
	payload := requests.RoleRequestPayload{UserID: 7, CorrelationID: "role-correlation", ReplyTopic: "role.response"}
	value, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	return &sarama.ConsumerMessage{Value: value}
}

func newRoleTestHandler(service service.RoleQueryService) (*roleKafkaHandler, *roleTestProducer, *roleTestSession) {
	producer := &roleTestProducer{}
	client := kafka.NewKafkaWithProducer(producer, roleTestLogger{}, nil)
	handler := NewRoleKafkaHandler(service, client, roleTestLogger{}, context.Background()).(*roleKafkaHandler)
	return handler, producer, &roleTestSession{}
}

func TestRoleKafkaHandler_NotFoundRepliesFalseAndMarks(t *testing.T) {
	handler, producer, session := newRoleTestHandler(&roleTestService{
		err: role_errors.ErrRoleNotFound.WithInternal(errors.New("no roles")),
	})

	if err := handler.ConsumeClaim(session, &roleTestClaim{messages: []*sarama.ConsumerMessage{roleRequestMessage(t)}}); err != nil {
		t.Fatalf("ConsumeClaim returned error: %v", err)
	}
	if producer.sent != 1 || session.marked != 1 {
		t.Fatalf("expected one reply and one mark, got replies=%d marks=%d", producer.sent, session.marked)
	}
}

func TestRoleKafkaHandler_InfraErrorReturnsAndDoesNotMark(t *testing.T) {
	handler, producer, session := newRoleTestHandler(&roleTestService{err: errors.New("database unavailable")})

	if err := handler.ConsumeClaim(session, &roleTestClaim{messages: []*sarama.ConsumerMessage{roleRequestMessage(t)}}); err == nil {
		t.Fatal("expected infrastructure error")
	}
	if producer.sent != 0 || session.marked != 0 {
		t.Fatalf("expected no reply and no mark, got replies=%d marks=%d", producer.sent, session.marked)
	}
}
