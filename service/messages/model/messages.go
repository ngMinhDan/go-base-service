package model

import (
	"base/pkg/db"
	"base/pkg/kafka"
	"base/pkg/utils"
	"base/service/messages/producer"
)

const KeyMessageToBroker = "lastest"

// User represents the data structure for a user.
type Message struct {
	ID        int     `json:"id"`
	Content   *string `json:"content"`
	CreatedAt *string `json:"createdAt"`
}

// CreateNewMessage : Create message, save into DB
func (mess *Message) CreateNewMessage() error {
	query := `INSERT INTO messages (content, created_at) VALUES ($1, $2) RETURNING id`
	err := db.PSQL.QueryRow(query, *mess.Content, *mess.CreatedAt).Scan(&mess.ID)
	return err
}

// ProducerMessage:  Send Message To Message Broker
func (mess Message) ProducerMessage() error {
	var messForm = producer.MessForm{
		Content:   mess.Content,
		CreatedAt: mess.CreatedAt,
	}

	var schema = &kafka.Schema[producer.MessForm]{
		Id:        mess.ID,
		Timestamp: utils.TimeNow(),
		Payload:   messForm,
		Key:       KeyMessageToBroker,
	}

	err := producer.MessageProducer.Produce(schema)
	if err != nil {
		return err
	}
	return nil
}
