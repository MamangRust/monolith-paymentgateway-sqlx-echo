package kafka

import (
	"context"
	"testing"
	"time"

	"github.com/IBM/sarama"
)

type lifecycleNoopHandler struct{}

func (lifecycleNoopHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (lifecycleNoopHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (lifecycleNoopHandler) ConsumeClaim(sarama.ConsumerGroupSession, sarama.ConsumerGroupClaim) error {
	return nil
}

func TestStartConsumersWithContextClosesGroupOnCancellation(t *testing.T) {
	group := newLifecycleTestGroup()
	client := NewKafkaWithProducer(nil, lifecycleTestLogger{}, nil)
	client.consumerFactory = func([]string, string, *sarama.Config) (sarama.ConsumerGroup, error) {
		return group, nil
	}
	ctx, cancel := context.WithCancel(context.Background())
	managed, err := client.StartConsumersWithContext(ctx, []string{"test-topic"}, "test-group", lifecycleNoopHandler{})
	if err != nil {
		t.Fatalf("StartConsumersWithContext returned error: %v", err)
	}

	cancel()
	closeDone := make(chan error, 1)
	go func() { closeDone <- managed.Close() }()

	select {
	case err := <-closeDone:
		if err != nil {
			t.Fatalf("managed consumer close returned error: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("managed consumer did not close after context cancellation")
	}

	select {
	case <-group.closeDone:
	case <-time.After(time.Second):
		t.Fatal("consumer group was not closed after context cancellation")
	}
}
