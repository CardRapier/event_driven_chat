package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService *UserService
}

func NewUserHandler(userService *UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (ch *UserHandler) StartRouting(router fiber.Router) {
	chat := router.Group("/user")
	chat.Get("/:userId", ch.CheckIfUserExists)
	chat.Post("", ch.CreateUser)
}

func (ch *UserHandler) CheckIfUserExists(c *fiber.Ctx) error {
	userIdStr := c.Params("userId")
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	err = ch.userService.CheckIfUserExists(&userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User Found",
	})
}

func (ch *UserHandler) CreateUser(c *fiber.Ctx) error {
	var newUser *CreateUser
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "cannot parse JSON",
        })
	}


	user, err := ch.userService.CreateUser(newUser)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}