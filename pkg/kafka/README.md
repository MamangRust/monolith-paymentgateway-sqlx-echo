# 📦 Package `kafka`

**Source Path:** `./pkg/kafka`

## 🧩 Types

### `Kafka`

```go
type Kafka struct {
	logger logger.LoggerInterface
	producer sarama.SyncProducer
	brokers []string
}
```

#### Methods

##### `GetBrokers`

GetBrokers returns a list of the Kafka broker addresses that the producer is connected to.

```go
func (k *Kafka) GetBrokers() []string
```

##### `SendMessage`

SendMessage sends a message to the given Kafka topic with the given key and value.

It uses the configured SyncProducer to send the message and logs the result of the send operation.
If the send operation fails, it returns an error.

```go
func (k *Kafka) SendMessage(topic string, key string, value []byte) error
```

##### `StartConsumers`

StartConsumers starts a Kafka consumer group to consume messages from the specified topics.

It takes a list of topics, a consumer group ID, and a handler implementing the sarama.ConsumerGroupHandler interface.
The method initializes a new Kafka consumer group and begins consuming messages in a background goroutine.
If an error occurs during consumption, it retries up to a maximum number of retries with a delay between attempts.
Any errors from the consumer group are logged and the function returns an error if the consumer group initialization fails.

```go
func (k *Kafka) StartConsumers(topics []string, groupID string, handler sarama.ConsumerGroupHandler) error
```

