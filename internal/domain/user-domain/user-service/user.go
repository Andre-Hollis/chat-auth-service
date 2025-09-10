package userservice

import (
	userdomain "github.com/Andre-Hollis/chat-auth-service/internal/domain/user-domain"
	userrepo "github.com/Andre-Hollis/chat-auth-service/internal/infra/user-repo"
	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	userRepo userrepo.IUserRepo
}

func NewUserService(userRepo userrepo.IUserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) ReadUser(c *fiber.Ctx, email string) (*userdomain.User, error) {
	user, err := s.userRepo.ReadUserByEmail(c.Context(), email)
	if err != nil {

	}
	return user, err
}

func (s *UserService) SaveUser(c *fiber.Ctx, user *userdomain.User) (*userdomain.User, error) {
	id, err := s.userRepo.Save(c.Context(), user)
	if err != nil {

	}
	user.SetId(id)
	return user, err
}
