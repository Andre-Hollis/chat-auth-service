package main

import (
	"log"

	"github.com/Andre-Hollis/chat-auth-service/api"
	"github.com/Andre-Hollis/chat-auth-service/internal/config"
	"github.com/Andre-Hollis/chat-auth-service/internal/infra/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err.Error())
	}

	app := fiber.New()

	user := app.Group("/user", middleware.AuthMiddleware())
	user.Post("/", api.CreateUser)
	user.Get("/:userId", api.ReadUser)

	app.Listen(":3000")
}
