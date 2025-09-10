package user

import (
	"github.com/Andre-Hollis/chat-auth-service/internal/application/user/handler"
	"github.com/Andre-Hollis/chat-auth-service/internal/infra/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, h *handler.UserHandler) {
	api := app.Group("/api/v1", middleware.JWTMiddleware())

	api.Post("/register", h.Register)
	api.Post("/login", h.Login)
	api.Post("/", h.CreateUser)

}
