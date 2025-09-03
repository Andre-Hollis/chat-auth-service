package user

import (
	"log"

	"github.com/Andre-Hollis/chat-auth-service/internal/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	userCreate := new(models.UserCreate)

	if err := c.BodyParser(userCreate); err != nil {
		log.Println(err)
	}

	return c.JSON(p)
}
