package userrepo

import (
	"context"

	userdomain "github.com/Andre-Hollis/chat-auth-service/internal/domain/user-domain"
)

type IUserRepo interface {
	Save(ctx context.Context, user *userdomain.User) error
	FindByID(ctx context.Context, id string) (*userdomain.User, error)
}
