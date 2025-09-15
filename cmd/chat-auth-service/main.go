package main

import (
	"github.com/Andre-Hollis/chat-auth-service/api/user"
	"github.com/Andre-Hollis/chat-auth-service/internal/application/user/handler"
	userservice "github.com/Andre-Hollis/chat-auth-service/internal/domain/user-domain/user-service"
	"github.com/Andre-Hollis/chat-auth-service/internal/infra/user-repo/impl"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// config, err := config.LoadConfig("config.json")
	// if err != nil {
	// 	log.Fatalf("Error loading configuration: %s", err.Error())
	// }

	app := fiber.New()

	userRepo := impl.NewUserRedisRepo()
	userService := userservice.NewUserService(userRepo)
	h := handler.NewUserHandler(userService)

	user.RegisterUserRoutes(app, h)

	app.Listen(":3000")
}
