package user

import (
	"github.com/Andre-Hollis/chat-auth-service/internal/domain/user"
	userrepo "github.com/Andre-Hollis/chat-auth-service/internal/infra/user-repo"
)

type UserService struct {
	userRepo userrepo.IUserRepo
}

func (s *UserService) ReadUser() (*user.User, error) {
	user, err = s.userRepo.FindByID()
	return
}

func (s *UserService) SaveUser() (*user.User, error) {

}
