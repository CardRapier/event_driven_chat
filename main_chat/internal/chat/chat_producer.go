package chat

type ChatProducer interface {
	PublishChatMessage(chatMessage *ChatMessage) error
}