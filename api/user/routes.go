package user

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App, h *Handler) {
	api := app.Group("/api/v1")

	api.Post("/register", h.Register)
	api.Post("/login", h.Login)
}
