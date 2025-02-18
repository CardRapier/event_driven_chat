package chat

type ChatProducer interface {
	PublishChatMessage(message *ChatMessage) error
}