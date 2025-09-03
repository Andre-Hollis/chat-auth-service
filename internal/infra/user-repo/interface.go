package userrepo

import (
	"context"

	"github.com/Andre-Hollis/chat-auth-service/internal/domain/user"
)

type IUserRepo interface {
	Save(ctx context.Context, user *user.User) error
	FindByID(ctx context.Context, id string) (*user.User, error)
}
