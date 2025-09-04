package tokenservice

import (
	"time"

	userdomain "github.com/Andre-Hollis/chat-auth-service/internal/domain/user-domain"
	userrepo "github.com/Andre-Hollis/chat-auth-service/internal/infra/user-repo"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	tokenRepo      userrepo.IUserRepo
	jwtSecret      []byte
	accessTokenTTL time.Duration
}

func NewTokenService(refreshTokenRepo *models.RefreshTokenRepository, jwtSecret string, accessTokenTTL time.Duration) *TokenService {
	return &TokenService{
		refreshTokenRepo: refreshTokenRepo,
		jwtSecret:        []byte(jwtSecret),
		accessTokenTTL:   accessTokenTTL,
	}
}

func (s *TokenService) GenerateAccessToken(c *fiber.Ctx, user userdomain.User) (string, error) {
	// Set the expiration time
	expirationTime := time.Now().Add(s.accessTokenTTL)

	// Create the JWT claims
	claims := jwt.MapClaims{
		"sub":      user.ID,               // subject (user ID)
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

func (s *TokenService) ValidateToken(c *fiber.Ctx, user *userdomain.User) (*userdomain.User, error) {
}
