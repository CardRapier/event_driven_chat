package chat

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type RedpandaChatProducer struct {
	redpanda *kafka.Writer
}

func NewRedpandaChatProducer(redpanda *kafka.Writer) *RedpandaChatProducer {
	return &RedpandaChatProducer{}
}

func (rcp *RedpandaChatProducer) PublishChatMessage(chatMessage *ChatMessage) error {
	// Convert the chatMessage to JSON bytes.
	messageBytes, err := json.Marshal(chatMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal chat message: %w", err)
	}
	err = rcp.redpanda.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(fmt.Sprintf("user-%s", chatMessage.UserId.String())),
			Value: messageBytes,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
