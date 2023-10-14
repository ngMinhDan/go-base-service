package consumer

import (
	"base/pkg/config"
	"base/pkg/elastic"
	"base/pkg/kafka"
	"base/pkg/log"
	"base/service/search/model"
	"fmt"
	"time"
)

var MessageConsumer *kafka.Consumer[model.Message]

func init() {
	var brokerAddresses = []string{config.Config.GetString("BROKER_ADDRESS")}
	MessageConsumer = (*kafka.Consumer[model.Message])(kafka.CreateConsumer(
		kafka.ConsumerConfig{
			Brokers:     brokerAddresses,
			Topic:       config.Config.GetString("TOPIC_SAMPLE"),
			GroupId:     config.Config.GetString("CONSUMER_GROUP_SAMPLE"),
			MinBytes:    1,
			MaxBytes:    1e6,
			MaxWait:     5 * time.Second,
			StartOffset: kafka.ConsumeNewDataOnly,
		},
	))
}

// ConsumeMessage: Consume Message from Kafka Broker, Insert into Elasticsearch
func ConsumeMessage() {
	for {
		var schema = &kafka.Schema[model.Message]{}
		schema, err := MessageConsumer.Consume()

		if err != nil {
			log.Println(log.LogLevelError, "MessageConsumer", err)
			continue
		}

		// Now Do Business Logic Here With New Message
		err = elastic.ES.Insert(config.Config.GetString("ES_SAMPLE_INDEX"), *&schema.Payload, fmt.Sprint(schema.Id))
		if err != nil {
			log.Println(log.LogLevelError, "insert-into-elastic-search", err.Error())
		}
	}
}
