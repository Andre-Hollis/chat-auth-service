package user

import (
	"github.com/Andre-Hollis/chat-auth-service/internal/application/user"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, h *user.UserHandler) {
	api := app.Group("/api/v1")

	api.Post("/register", h.Register)
	api.Post("/login", h.Login)
	api.Post("/", h.CreateUser)

}
