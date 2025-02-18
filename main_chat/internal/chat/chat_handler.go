package chat

import (
	"github.com/gofiber/fiber/v2"
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
	var message *ChatMessage

	if err := c.BodyParser(&message); err != nil {
		return err
	}

	err := ch.chatService.CreateMessage(message)
	if err != nil {
		return err
	}

	return nil;
}