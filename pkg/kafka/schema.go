package kafka

// Payload Struct : For Any Data We Want to Sent To Broker
type Payload interface {
	any
}

// Schema Event Data
// When a producer send a message to kafka topic
// It Can Set Key as a Part Of Message Metadata
// Kafka uses this key to determine the partition to which the message will be written
// Ensuring that related messages or events are processed in the order they were produced
type Schema[T Payload] struct {
	Id        int    `json:"id"`  // You Can Use ID Or Not Use
	Key       string `json:"key"` // Key or Type of event, we will set as a part of metadata
	Timestamp string `json:"timestamp"`
	Payload   T      `json:"payload"`
}
