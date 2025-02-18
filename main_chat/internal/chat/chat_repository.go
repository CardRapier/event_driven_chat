package chat

type ChatRepository interface {
	SaveMessage(message *ChatMessage) error
}