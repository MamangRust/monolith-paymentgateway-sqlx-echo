package kafka

import (
	"errors"
	"testing"

	"github.com/IBM/sarama"
)

type failingProducer struct {
	calls int
	err   error
}

func (p *failingProducer) SendMessage(*sarama.ProducerMessage) (int32, int64, error) {
	p.calls++
	return 0, 0, p.err
}
func (p *failingProducer) Close() error { return nil }

func TestSendMessageWithRetryReturnsFailureAfterBoundedAttempts(t *testing.T) {
	producer := &failingProducer{err: errors.New("broker unavailable")}
	client := NewKafkaWithProducer(producer, lifecycleTestLogger{}, nil)

	err := client.SendMessageWithRetry("test-topic", "test-key", []byte("payload"), 3)
	if err == nil {
		t.Fatal("expected producer failure")
	}
	if producer.calls != 3 {
		t.Fatalf("expected 3 producer attempts, got %d", producer.calls)
	}
}

func TestSendMessageWithRetryStopsAfterRecovery(t *testing.T) {
	producer := &recoveringProducer{failures: 2, err: errors.New("temporary broker failure")}
	client := NewKafkaWithProducer(producer, lifecycleTestLogger{}, nil)

	if err := client.SendMessageWithRetry("test-topic", "test-key", []byte("payload"), 3); err != nil {
		t.Fatalf("expected recovery on third attempt, got %v", err)
	}
	if producer.calls != 3 {
		t.Fatalf("expected 3 producer attempts, got %d", producer.calls)
	}
}

type recoveringProducer struct {
	calls    int
	failures int
	err      error
}

func (p *recoveringProducer) SendMessage(*sarama.ProducerMessage) (int32, int64, error) {
	p.calls++
	if p.calls <= p.failures {
		return 0, 0, p.err
	}
	return 0, 1, nil
}
func (p *recoveringProducer) Close() error { return nil }
