package user

import (
	"github.com/Andre-Hollis/chat-auth-service/internal/application/user/dto"
	"github.com/Andre-Hollis/chat-auth-service/internal/application/user/mappers"
	userdomain "github.com/Andre-Hollis/chat-auth-service/internal/domain/user-domain"

	"github.com/gofiber/fiber/v2"
)

type TokenHandler struct {
	TokenService *tokenservice.TokenService // injected service from internal/auth
}

// NewHandler returns a new Handler with its dependencies
func NewUserHandler(tokenService *tokenservice.TokenService) *TokenHandler {
	return &TokenHandler{
		TokenService: tokenService,
	}
}

// Register handles user registration
func (h *TokenHandler) Register(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	user, err := h.TokenService.Register(body.Email, body.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// Login handles user login
func (h *TokenHandler) Login(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	token, err := h.TokenService.Login(body.Email, body.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(fiber.Map{"token": token})
}

// Login handles user login
func (h *TokenHandler) CreateUser(c *fiber.Ctx) error {
	var userCreate dto.UserCreateDTO

	if err := c.BodyParser(&userCreate); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	user := userdomain.User{
		Email:        userCreate.Email,
		Username:     userCreate.Username,
		PasswordHash: userCreate.Password,
	}

	h.TokenService.SaveUser(c, &user)

	return c.Status(fiber.StatusOK).JSON(mappers.UserToDto(&user))
}
