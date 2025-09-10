package handler

import (
	"github.com/Andre-Hollis/chat-auth-service/internal/application/user/dto"
	"github.com/Andre-Hollis/chat-auth-service/internal/application/user/mappers"
	authservice "github.com/Andre-Hollis/chat-auth-service/internal/domain/auth-domain/auth-service"
	userdomain "github.com/Andre-Hollis/chat-auth-service/internal/domain/user-domain"
	userservice "github.com/Andre-Hollis/chat-auth-service/internal/domain/user-domain/user-service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *userservice.UserService // injected service from internal/auth
	AuthService *authservice.AuthService // injected service from internal/auth
}

// NewHandler returns a new Handler with its dependencies
func NewUserHandler(userService *userservice.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// Register handles user registration
func (h *UserHandler) Register(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	user, err := h.AuthService.Register(c, body.Email, body.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// Login handles user login
func (h *UserHandler) Login(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	token, err := h.AuthService.Login(c.Context(), body.Email, body.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(fiber.Map{"token": token})
}

// Login handles user login
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var userCreate dto.UserCreateDTO

	if err := c.BodyParser(&userCreate); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	user := userdomain.User{
		Email:        userCreate.Email,
		Username:     userCreate.Username,
		PasswordHash: userCreate.Password,
	}

	h.UserService.SaveUser(c, &user)

	return c.Status(fiber.StatusOK).JSON(mappers.UserToDto(&user))
}
