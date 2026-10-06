package tests

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/IBM/sarama"
	kafkamodule "github.com/testcontainers/testcontainers-go/modules/kafka"
)

func TestKafkaTestcontainerProducerConsumer(t *testing.T) {
	if os.Getenv("RUN_KAFKA_CONTAINER_TESTS") != "1" {
		t.Skip("set RUN_KAFKA_CONTAINER_TESTS=1 to run Kafka testcontainer coverage")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	container, err := kafkamodule.RunContainer(ctx)
	if err != nil {
		t.Fatalf("start Kafka testcontainer: %v", err)
	}
	defer container.Terminate(context.Background())

	brokers, err := container.Brokers(ctx)
	if err != nil {
		t.Fatalf("resolve Kafka brokers: %v", err)
	}

	config := sarama.NewConfig()
	config.Version = sarama.V3_6_0_0
	config.Producer.Return.Successes = true
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		t.Fatalf("create Kafka producer: %v", err)
	}
	defer producer.Close()

	consumer, err := sarama.NewConsumerGroup(brokers, "testcontainer-group", config)
	if err != nil {
		t.Fatalf("create Kafka consumer group: %v", err)
	}
	defer consumer.Close()

	topic := "testcontainer-events"
	payload := map[string]string{"event": "producer-consumer"}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	consumed := make(chan struct{})
	handler := &kafkaTestHandler{payload: payloadBytes, consumed: consumed}
	consumeCtx, consumeCancel := context.WithTimeout(ctx, 30*time.Second)
	defer consumeCancel()
	go func() {
		_ = consumer.Consume(consumeCtx, []string{topic}, handler)
	}()

	if _, _, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder("event-1"),
		Value: sarama.ByteEncoder(payloadBytes),
	}); err != nil {
		t.Fatalf("publish Kafka event: %v", err)
	}

	select {
	case <-consumed:
	case <-consumeCtx.Done():
		t.Fatalf("consumer did not receive event: %v", consumeCtx.Err())
	}
}

type kafkaTestHandler struct {
	payload  []byte
	consumed chan struct{}
}

func (h *kafkaTestHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h *kafkaTestHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (h *kafkaTestHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if string(message.Value) == string(h.payload) {
			session.MarkMessage(message, "")
			select {
			case <-h.consumed:
			default:
				close(h.consumed)
			}
			return nil
		}
	}
	return nil
}
