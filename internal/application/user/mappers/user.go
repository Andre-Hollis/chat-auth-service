package mappers

import (
	"github.com/Andre-Hollis/chat-auth-service/internal/application/user/dto"
	userdomain "github.com/Andre-Hollis/chat-auth-service/internal/domain/user-domain"
)

func UserToDto(u *userdomain.User) dto.UserDTO {
	return dto.UserDTO{
		Email:    u.Email,
		Username: u.Username,
	}
}
