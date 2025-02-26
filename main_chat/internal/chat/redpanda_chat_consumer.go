package chat

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type RedpandaChatConsumer struct {
	redpanda *kafka.Reader
}

func NewRedpandaChatConsumer(redpanda *kafka.Reader) *RedpandaChatConsumer {
	return &RedpandaChatConsumer{
		redpanda: redpanda,
	}
}

func (rcc *RedpandaChatConsumer) ConsumeChatMessages() {
	defer rcc.redpanda.Close()

	for {
		// Read a message from the topic
		m, err := rcc.redpanda.ReadMessage(context.Background())
		if err != nil {
			log.Printf("failed to read message: %v", err)
			break
		}
		if string(m.Key) == "chat_message" {
			fmt.Printf("Consumer - Received: %s\n", string(m.Value))
		}

		// Commit the offset after processing the message
		if err := rcc.redpanda.CommitMessages(context.Background(), m); err != nil {
			log.Printf("failed to commit message: %v", err)
		}
	}
}