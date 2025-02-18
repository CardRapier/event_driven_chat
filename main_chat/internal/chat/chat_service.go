package chat

type ChatService struct {
	chatProducer ChatProducer
}

func NewChatService(chatProducer ChatProducer) *ChatService {
	return &ChatService{chatProducer: chatProducer}
}

func (cs *ChatService) CreateMessage(message *ChatMessage) error {
	err := cs.chatProducer.PublishChatMessage(message)

	if err != nil {
		return err
	}

	return nil
}