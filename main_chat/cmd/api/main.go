package api

import (
	"chat/internal/chat"

	"github.com/gofiber/fiber/v2"
)

func StartApi(){
	app := fiber.New()

	startRouting(app)

	app.Listen(":3000")
}

func startRouting(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("v1")

	chatHandler := chat.NewChatHandler()

	chatHandler.StartRouting(v1)
}