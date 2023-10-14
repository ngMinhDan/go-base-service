package kafka

import (
	"base/pkg/log"
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

// RequiredAcks enum
const (
	// fire-and-forget, do not wait for acknowledgements from the cluster
	AcksRequireNone = kafka.RequireNone
	// only wait for the leader to acknowledge - good for non-transactional data
	AcksRequireOne = kafka.RequireOne
	// wait for all brokers to acknowledge the writes (In-Sync Replicas)
	AcksRequireAll = kafka.RequireAll
)

const (
	// Define Time Sleep When Write Fail
	_defaultKafkaSleep = 250 * time.Millisecond
	// Define Maximum Of Retry When Write Message
	_defaultRetries = 5
)

type ProducerConfig struct {
	Brokers []string
	Topic   string

	//BatchSize - The total number of messages that should be buffered before writing to the brokers
	BatchSize int //Default value is 100

	// BatchTimeout - The maximum time before which messages are written to the brokers.
	// That means that even if the message batch is not full,
	// they will still be written onto the Kafka cluster once this time period has elapsed.
	BatchTimeout time.Duration //Default value is 1 seconds

	RequiredAcks kafka.RequiredAcks //Default value is AcksRequireOne(1)
}

type Producer[T Payload] struct {
	writer *kafka.Writer
	// You need define type of prodcuer or something below
}

// Producer Create Method
func CreateProducer(config ProducerConfig) *Producer[Payload] {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(config.Brokers...),
		Topic:                  config.Topic,
		BatchSize:              config.BatchSize,
		BatchTimeout:           config.BatchTimeout,
		RequiredAcks:           config.RequiredAcks,
		AllowAutoTopicCreation: true,
	}
	return &Producer[Payload]{
		writer: writer,
	}
}

// Produce Message Method
func (producer *Producer[T]) Produce(schema *Schema[T]) error {

	jsonData, err := json.Marshal(schema)
	if err != nil {
		return err
	}

	key, err := json.Marshal(schema.Key)
	if err != nil {
		return err
	}

	for i := 0; i < _defaultRetries; i++ {
		// Must have to create topic before send message or event to broker
		// Or auto.create.topics.enable = true in setting of kafka

		err = producer.writer.WriteMessages(context.Background(), kafka.Message{
			Value: jsonData,
			// Kafka uses this key to determine the partition to which the message will be written
			Key: key,
		})
		if err != nil {
			log.Println(log.LogLevelError, "kafka-write-error", err)
			time.Sleep(_defaultKafkaSleep)
			continue
		} else {
			break
		}
	}
	return nil
}
