package producer

import (
	"base/pkg/config"
	"base/pkg/kafka"
	"time"
)

type MessForm struct {
	Content   *string `json:"content"`
	CreatedAt *string `json:"createdAt"`
}

var MessageProducer *kafka.Producer[MessForm]

func init() {
	MessageProducer = (*kafka.Producer[MessForm])(kafka.CreateProducer(
		kafka.ProducerConfig{
			Brokers:      []string{config.Config.GetString("BROKER_ADDRESS")},
			Topic:        config.Config.GetString("TOPIC_SAMPLE"),
			BatchSize:    1,
			BatchTimeout: 5 * time.Second,
			RequiredAcks: kafka.AcksRequireOne,
		},
	))
}
