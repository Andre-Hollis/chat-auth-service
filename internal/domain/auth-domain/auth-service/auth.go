package authservice

import (
	"context"
	"database/sql"
	"errors"
	"time"

	userdomain "github.com/Andre-Hollis/chat-auth-service/internal/domain/user-domain"
	userrepo "github.com/Andre-Hollis/chat-auth-service/internal/infra/user-repo"
	"github.com/Andre-Hollis/chat-auth-service/internal/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
	ErrExpiredToken       = errors.New("token has expired")
	ErrEmailInUse         = errors.New("email already in use")
)

type AuthService struct {
	userRepo       userrepo.IUserRepo
	jwtSecret      []byte
	accessTokenTTL time.Duration
}

func NewAuthService(
	userRepo userrepo.IUserRepo,
	jwtSecret string,
	accessTokenTTL time.Duration,
) *AuthService {
	return &AuthService{
		userRepo:       userRepo,
		jwtSecret:      []byte(jwtSecret),
		accessTokenTTL: accessTokenTTL,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	// Get the user from the database
	user, err := s.userRepo.ReadUserByEmail(ctx, email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	// Verify the password
	if err := utils.VerifyPassword(user.PasswordHash, password); err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := s.GenerateAccessToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) GenerateAccessToken(user *userdomain.User) (string, error) {
	// Set the expiration time
	expirationTime := time.Now().Add(s.accessTokenTTL)

	// Create the JWT claims
	claims := jwt.MapClaims{
		"sub":      user.Email,            // subject (user ID)
		"username": user.Username,         // custom claim
		"email":    user.Email,            // custom claim
		"exp":      expirationTime.Unix(), // expiration time
		"iat":      time.Now().Unix(),     // issued at time
	}
	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with our secret key
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *AuthService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return s.jwtSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	// Extract and validate claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

// Register creates a new user with the provided credentials
func (s *AuthService) Register(c *fiber.Ctx, email, password string) (*userdomain.User, error) {
	// Check if user already exists
	_, err := s.userRepo.ReadUserByEmail(c.Context(), email)
	if err == nil {
		return nil, ErrEmailInUse
	}
	// Only proceed if the error was "user not found"
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	// Hash the password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser := userdomain.User{
		Email:        email,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
		LastLogin:    time.Now(),
	}

	// Create the user
	_, err = s.userRepo.Save(c.Context(), &newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
