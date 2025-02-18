package chat

import "github.com/gofiber/fiber/v2"

type ChatHandler struct {}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{}
}

func (ch *ChatHandler) StartRouting(router fiber.Router) {
	chat := router.Group("/chat")
	chat.Post("", ch.CreateMessage)
}

func (ch *ChatHandler) CreateMessage(c *fiber.Ctx) error {

	return nil;
}