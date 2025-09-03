package user

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	AuthService *auth.Service // injected service from internal/auth
}

// NewHandler returns a new Handler with its dependencies
func NewHandler(authService *auth.Service) *Handler {
	return &Handler{
		AuthService: authService,
	}
}

// Register handles user registration
func (h *Handler) Register(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	user, err := h.AuthService.Register(body.Email, body.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// Login handles user login
func (h *Handler) Login(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	token, err := h.AuthService.Login(body.Email, body.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(fiber.Map{"token": token})
}
