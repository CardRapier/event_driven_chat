package chat

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ChatHandler struct {
	chatService *ChatService
}

func NewChatHandler(chatService *ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}

func (ch *ChatHandler) StartRouting(router fiber.Router) {
	chat := router.Group("/chat")
	chat.Post("", ch.CreateMessage)
}

func (ch *ChatHandler) CreateMessage(c *fiber.Ctx) error {
	var message *CreateChatMessage

	if err := c.BodyParser(&message); err != nil {
		return err
	}

	createMessage := &ChatMessage{
		UserId: message.UserId,
		Message: message.Message,
		Timestamp: message.Timestamp,
		MessageId: uuid.New(),
	}

	err := ch.chatService.CreateMessage(createMessage)
	if err != nil {
		return err
	}

	return nil;
}