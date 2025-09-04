package userservice

import (
	"github.com/Andre-Hollis/chat-auth-service/internal/domain/user"
	userrepo "github.com/Andre-Hollis/chat-auth-service/internal/infra/user-repo"
	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	userRepo userrepo.IUserRepo
}

func (s *UserService) ReadUser() (*user.User, error) {
	user, err = s.userRepo.FindByID()
	return
}

func (s *UserService) SaveUser(c *fiber.Ctx, user *user.User) (*user.User, error) {
	id, err := s.userRepo.Save(c.Context(), user)
	user.SetId()
}
