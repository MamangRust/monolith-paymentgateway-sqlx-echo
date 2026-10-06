package kafka

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type lifecycleTestLogger struct{}

func (lifecycleTestLogger) Info(string, ...zap.Field)                         {}
func (lifecycleTestLogger) Fatal(string, ...zap.Field)                        {}
func (lifecycleTestLogger) Debug(string, ...zap.Field)                        {}
func (lifecycleTestLogger) Error(string, ...zap.Field)                        {}
func (lifecycleTestLogger) Warn(string, ...zap.Field)                         {}
func (lifecycleTestLogger) Check(zapcore.Level, string) *zapcore.CheckedEntry { return nil }
func (lifecycleTestLogger) With(...zap.Field) logger.LoggerInterface          { return lifecycleTestLogger{} }
func (lifecycleTestLogger) Sync() error                                       { return nil }

type lifecycleTestProducer struct {
	mu     sync.Mutex
	closed int
}

func (p *lifecycleTestProducer) SendMessage(*sarama.ProducerMessage) (int32, int64, error) {
	return 0, 0, nil
}
func (p *lifecycleTestProducer) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.closed++
	return nil
}

type lifecycleTestGroup struct {
	closeOnce sync.Once
	closed    chan struct{}
	errors    chan error
	closeDone chan struct{}
}

func newLifecycleTestGroup() *lifecycleTestGroup {
	return &lifecycleTestGroup{
		closed:    make(chan struct{}),
		errors:    make(chan error),
		closeDone: make(chan struct{}),
	}
}

func (g *lifecycleTestGroup) Consume(ctx context.Context, _ []string, _ sarama.ConsumerGroupHandler) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-g.closed:
		return context.Canceled
	}
}
func (g *lifecycleTestGroup) Errors() <-chan error { return g.errors }
func (g *lifecycleTestGroup) Close() error {
	g.closeOnce.Do(func() {
		close(g.closed)
		close(g.errors)
		close(g.closeDone)
	})
	return nil
}
func (g *lifecycleTestGroup) Pause(map[string][]int32)  {}
func (g *lifecycleTestGroup) Resume(map[string][]int32) {}
func (g *lifecycleTestGroup) PauseAll()                 {}
func (g *lifecycleTestGroup) ResumeAll()                {}

func TestManagedConsumerCloseStopsConsumeAndErrors(t *testing.T) {
	group := newLifecycleTestGroup()
	ctx, cancel := context.WithCancel(context.Background())
	managed := &managedConsumer{
		group:      group,
		cancel:     cancel,
		done:       make(chan struct{}),
		errorsDone: make(chan struct{}),
	}

	go func() {
		defer close(managed.done)
		<-ctx.Done()
	}()
	go func() {
		defer close(managed.errorsDone)
		for range group.Errors() {
		}
	}()

	if err := managed.Close(); err != nil {
		t.Fatalf("managed.Close returned error: %v", err)
	}
	select {
	case <-group.closeDone:
	case <-time.After(time.Second):
		t.Fatal("consumer group was not closed")
	}
	if err := managed.Close(); err != nil {
		t.Fatalf("second managed.Close returned error: %v", err)
	}
}

func TestKafkaCloseClosesProducerOnce(t *testing.T) {
	producer := &lifecycleTestProducer{}
	client := NewKafkaWithProducer(producer, lifecycleTestLogger{}, nil)

	if err := client.Close(); err != nil {
		t.Fatalf("Kafka.Close returned error: %v", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("second Kafka.Close returned error: %v", err)
	}
	producer.mu.Lock()
	closed := producer.closed
	producer.mu.Unlock()
	if closed != 1 {
		t.Fatalf("expected producer to close once, got %d closes", closed)
	}
}
