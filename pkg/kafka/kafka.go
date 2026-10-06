package kafka

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// SyncProducer is an interface that represents a Kafka producer.
type SyncProducer interface {
	SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error)
	Close() error
}

// Kafka is a struct that represents a Kafka producer.
type consumerGroupFactory func([]string, string, *sarama.Config) (sarama.ConsumerGroup, error)

type Kafka struct {
	logger          logger.LoggerInterface
	producer        SyncProducer
	brokers         []string
	consumerFactory consumerGroupFactory

	consumersMu sync.Mutex
	consumers   []ConsumerCloser
	closeOnce   sync.Once
	closeErr    error
}

// ConsumerCloser owns a Sarama consumer group started by Kafka and can stop it
// during application shutdown.
type ConsumerCloser interface {
	Close() error
}

var ErrConsumerCloseTimeout = errors.New("timed out closing Kafka consumer")

const consumerCloseTimeout = 5 * time.Second

type managedConsumer struct {
	group      sarama.ConsumerGroup
	cancel     context.CancelFunc
	done       chan struct{}
	errorsDone chan struct{}
	closeOnce  sync.Once
	closeErr   error
}

func (c *managedConsumer) Close() error {
	c.closeOnce.Do(func() {
		c.cancel()

		closeDone := make(chan error, 1)
		go func() {
			closeDone <- c.group.Close()
		}()

		timer := time.NewTimer(consumerCloseTimeout)
		defer timer.Stop()
		select {
		case c.closeErr = <-closeDone:
		case <-timer.C:
			c.closeErr = ErrConsumerCloseTimeout
			return
		}

		select {
		case <-c.done:
		case <-timer.C:
			c.closeErr = ErrConsumerCloseTimeout
			return
		}
		select {
		case <-c.errorsDone:
		case <-timer.C:
			c.closeErr = ErrConsumerCloseTimeout
		}
	})
	return c.closeErr
}

// NewKafka initializes a new Kafka struct.
//
// It takes a logger and a list of broker addresses as inputs and returns a pointer to the Kafka struct.
// It creates a new Kafka producer with the given configuration, and logs a message indicating if the connection is successful.
// If the connection fails, it logs an error message and exits.
func NewKafka(logger logger.LoggerInterface, brokers []string) *Kafka {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	log.Println("Kafka producer connected successfully")

	return &Kafka{
		producer:        producer,
		brokers:         brokers,
		logger:          logger,
		consumerFactory: sarama.NewConsumerGroup,
	}
}

// NewKafkaWithProducer creates a Kafka client backed by a caller-provided
// SyncProducer. It is primarily intended for tests (e.g. injecting a fake or
// mock producer) and keeps the same production code paths (SendMessage,
// SendMessageWithRetry, StartConsumers) intact.
func NewKafkaWithProducer(producer SyncProducer, logger logger.LoggerInterface, brokers []string) *Kafka {
	return &Kafka{
		producer:        producer,
		brokers:         brokers,
		logger:          logger,
		consumerFactory: sarama.NewConsumerGroup,
	}
}

// GetBrokers returns a list of the Kafka broker addresses that the producer is connected to.
func (k *Kafka) GetBrokers() []string {
	return k.brokers
}

// SendMessage sends a message to the given Kafka topic with the given key and value.
//
// It uses the configured SyncProducer to send the message and logs the result of the send operation.
// If the send operation fails, it returns an error.
func (k *Kafka) SendMessage(topic string, key string, value []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(value),
	}

	partition, offset, err := k.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}

// SendMessageWithHeaders publishes a message with explicit record headers in
// addition to the payload, propagating the active trace context as
// traceparent/tracestate headers. Used by the email service to attach retry/DLQ
// metadata (Phase 4) while keeping the payload an unchanged envelope.
func (k *Kafka) SendMessageWithHeaders(ctx context.Context, topic string, key string, value []byte, headers []sarama.RecordHeader) error {
	msg := &sarama.ProducerMessage{
		Topic:   topic,
		Key:     sarama.StringEncoder(key),
		Value:   sarama.ByteEncoder(value),
		Headers: headers,
	}
	if ctx != nil {
		carrier := propagation.MapCarrier{}
		otel.GetTextMapPropagator().Inject(ctx, carrier)
		if len(carrier) > 0 {
			for kk, vv := range carrier {
				msg.Headers = append(msg.Headers, sarama.RecordHeader{Key: []byte(kk), Value: []byte(vv)})
			}
		}
	}

	partition, offset, err := k.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}

// SendMessageWithRetry sends a message synchronously with up to attempts tries
// and a short linear backoff between them. The caller decides how to handle the
// final error; this helper exists so callers can publish deterministically
// instead of fire-and-forget after a database commit. The underlying sarama
// producer already retries internally (Producer.Retry.Max), so this loop mainly
// guards against transient producer-level failures.
func (k *Kafka) SendMessageWithRetry(topic string, key string, value []byte, attempts int) error {
	var lastErr error
	for i := 0; i < attempts; i++ {
		if err := k.SendMessage(topic, key, value); err != nil {
			lastErr = err
			if i < attempts-1 {
				time.Sleep(200 * time.Millisecond * time.Duration(i+1))
			}
			continue
		}
		return nil
	}
	return lastErr
}

// StartConsumersWithContextManualCommit starts a managed consumer group with
// Sarama auto-commit disabled; the handler must call session.Commit()
// explicitly after reaching a terminal outcome. Used by the email retry
// processor (Phase 4) so a retry/DLQ publish that fails leaves the offset
// uncommitted and the message is redelivered.
func (k *Kafka) StartConsumersWithContextManualCommit(ctx context.Context, topics []string, groupID string, handler sarama.ConsumerGroupHandler) (ConsumerCloser, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = false

	consumerFactory := k.consumerFactory
	if consumerFactory == nil {
		consumerFactory = sarama.NewConsumerGroup
	}
	consumerGroup, err := consumerFactory(k.brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	consumerCtx, cancel := context.WithCancel(ctx)
	managed := &managedConsumer{
		group:      consumerGroup,
		cancel:     cancel,
		done:       make(chan struct{}),
		errorsDone: make(chan struct{}),
	}
	k.consumersMu.Lock()
	k.consumers = append(k.consumers, managed)
	k.consumersMu.Unlock()

	go func() {
		<-consumerCtx.Done()
		_ = consumerGroup.Close()
	}()

	go func() {
		defer close(managed.done)
		defer cancel()
		defer consumerGroup.Close()
		retries := 0
		const maxRetries = 5
		for {
			if consumerCtx.Err() != nil {
				return
			}

			consumeErr := consumerGroup.Consume(consumerCtx, topics, handler)
			if consumerCtx.Err() != nil {
				return
			}
			if consumeErr != nil {
				log.Printf("Error from consumer: %v", consumeErr)
				retries++
				if retries >= maxRetries {
					log.Printf("Max retries reached for consumer group; stopping consumer")
					return
				}

				timer := time.NewTimer(30 * time.Second)
				select {
				case <-consumerCtx.Done():
					if !timer.Stop() {
						select {
						case <-timer.C:
						default:
						}
					}
					return
				case <-timer.C:
				}
				continue
			}
			retries = 0
		}
	}()

	go func() {
		defer close(managed.errorsDone)
		for {
			select {
			case err, ok := <-consumerGroup.Errors():
				if !ok {
					return
				}
				log.Printf("Consumer group error: %v", err)
			case <-consumerCtx.Done():
				return
			}
		}
	}()

	return managed, nil
}

// StartConsumers starts a Kafka consumer group with a background lifetime for
// backwards compatibility. New applications should use StartConsumersWithContext
// and close the returned handle during shutdown.
func (k *Kafka) StartConsumers(topics []string, groupID string, handler sarama.ConsumerGroupHandler) error {
	_, err := k.StartConsumersWithContext(context.Background(), topics, groupID, handler)
	return err
}

// StartConsumersWithContext starts a managed Kafka consumer group. Cancelling
// ctx or closing the returned handle stops both the Consume loop and Sarama's
// error stream, preventing consumer goroutine leaks during service shutdown.
func (k *Kafka) StartConsumersWithContext(ctx context.Context, topics []string, groupID string, handler sarama.ConsumerGroupHandler) (ConsumerCloser, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	// Start new consumer groups at the oldest available offset so a fresh
	// deployment can replay events that were published while it was offline.
	// Existing committed group offsets still take precedence.
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerFactory := k.consumerFactory
	if consumerFactory == nil {
		consumerFactory = sarama.NewConsumerGroup
	}
	consumerGroup, err := consumerFactory(k.brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	consumerCtx, cancel := context.WithCancel(ctx)
	managed := &managedConsumer{
		group:      consumerGroup,
		cancel:     cancel,
		done:       make(chan struct{}),
		errorsDone: make(chan struct{}),
	}
	k.consumersMu.Lock()
	k.consumers = append(k.consumers, managed)
	k.consumersMu.Unlock()

	go func() {
		<-consumerCtx.Done()
		// Close is idempotent inside Sarama. This also closes Errors() when the
		// parent service context is cancelled externally.
		_ = consumerGroup.Close()
	}()

	go func() {
		defer close(managed.done)
		defer cancel()
		defer consumerGroup.Close()
		retries := 0
		const maxRetries = 5
		for {
			if consumerCtx.Err() != nil {
				return
			}

			consumeErr := consumerGroup.Consume(consumerCtx, topics, handler)
			if consumerCtx.Err() != nil {
				return
			}
			if consumeErr != nil {
				log.Printf("Error from consumer: %v", consumeErr)
				retries++
				if retries >= maxRetries {
					log.Printf("Max retries reached for consumer group; stopping consumer")
					return
				}

				timer := time.NewTimer(30 * time.Second)
				select {
				case <-consumerCtx.Done():
					if !timer.Stop() {
						select {
						case <-timer.C:
						default:
						}
					}
					return
				case <-timer.C:
				}
				continue
			}
			retries = 0
		}
	}()

	go func() {
		defer close(managed.errorsDone)
		for {
			select {
			case err, ok := <-consumerGroup.Errors():
				if !ok {
					return
				}
				log.Printf("Consumer group error: %v", err)
			case <-consumerCtx.Done():
				return
			}
		}
	}()

	return managed, nil
}

// Close stops all consumers and then closes the producer. It is safe to call
// more than once and is intended to be called from the owning service's
// shutdown path.
func (k *Kafka) Close() error {
	k.closeOnce.Do(func() {
		k.consumersMu.Lock()
		consumers := append([]ConsumerCloser(nil), k.consumers...)
		k.consumersMu.Unlock()

		for _, consumer := range consumers {
			if err := consumer.Close(); err != nil && k.closeErr == nil {
				k.closeErr = err
			}
		}
		if k.producer != nil {
			if err := k.producer.Close(); err != nil && k.closeErr == nil {
				k.closeErr = err
			}
		}
	})
	return k.closeErr
}
